// Package lemak is the Go client of the Lemak merchant API.
//
// A Client holds the API key and, optionally, a request Signer. Every POST carries an
// Idempotency-Key, generated when the caller sets none, and a failed attempt is retried under
// the same key with a fresh timestamp and signature. VerifyWebhook checks the
// X-Webhook-Signature header of a delivery, and ParseEvent decodes its body.
//
// The request and response types in types.go are generated from the API specification.
package lemak

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	mrand "math/rand/v2"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// DefaultBaseURL is where a Client sends requests when BaseURL is empty.
const DefaultBaseURL = "https://api.lemak.cash"

// APIVersion is the version these types describe. A Client sends it as Api-Version unless
// its own APIVersion field names another.
const APIVersion = "2026-09-01"

const (
	defaultRetries = 2
	defaultTimeout = 30 * time.Second
	// The backoff doubles from backoffBase to backoffCap, and each wait is drawn from its
	// upper half so that clients that failed together do not retry together.
	backoffBase = 500 * time.Millisecond
	backoffCap  = 8 * time.Second
	// maxBody limits the size of a response body.
	maxBody = 16 << 20
)

// A Client calls the merchant API. Key is required; every other field has a working zero value.
// A Client is safe for concurrent use.
type Client struct {
	// Key is the secret API key, sk_test_… or sk_live_….
	Key string
	// BaseURL defaults to DefaultBaseURL.
	BaseURL string
	// APIVersion defaults to the APIVersion constant. A request's own ApiVersion parameter wins.
	APIVersion string
	// Signer signs every request when the account requires a signature_scheme.
	Signer Signer
	// HTTPClient defaults to http.DefaultClient.
	HTTPClient *http.Client
	// MaxRetries is how many times a failed attempt is retried. Zero means 2; a negative
	// value turns retries off. Only a network error, a 429 and a 5xx are retried.
	MaxRetries int
	// Timeout bounds each attempt. Zero means 30 seconds.
	Timeout time.Duration
}

// A Signer signs the canonical string of one request attempt. It returns the X-Signature
// value and the X-Key-Id value, which is empty for a scheme that has none.
type Signer interface {
	Sign(ctx context.Context, canonical []byte) (signature, keyID string, err error)
}

// HMACSigner signs requests with the hmac_sha256 scheme. Secret is the 64-character hex
// secret shown in the dashboard.
type HMACSigner struct {
	Secret string
}

func (s HMACSigner) Sign(_ context.Context, canonical []byte) (string, string, error) {
	key, err := hex.DecodeString(s.Secret)
	if err != nil {
		return "", "", fmt.Errorf("signing secret is not hex: %w", err)
	}
	mac := hmac.New(sha256.New, key)
	mac.Write(canonical)
	return "v1=" + base64.StdEncoding.EncodeToString(mac.Sum(nil)), "", nil
}

// Ed25519Signer signs requests with the ed25519 scheme. KeyID is the key_… ID of the public
// key registered in the dashboard.
type Ed25519Signer struct {
	KeyID string
	Key   ed25519.PrivateKey
}

func (s Ed25519Signer) Sign(_ context.Context, canonical []byte) (string, string, error) {
	if len(s.Key) != ed25519.PrivateKeySize {
		return "", "", errors.New("ed25519 private key has the wrong length")
	}
	return "v2=" + base64.StdEncoding.EncodeToString(ed25519.Sign(s.Key, canonical)), s.KeyID, nil
}

// Ed25519Func signs requests with the ed25519 scheme using a key held outside the process,
// such as an AWS KMS key or a hardware security module. Func returns the 64-byte Ed25519
// signature of msg.
type Ed25519Func struct {
	KeyID string
	Func  func(ctx context.Context, msg []byte) ([]byte, error)
}

func (s Ed25519Func) Sign(ctx context.Context, canonical []byte) (string, string, error) {
	sig, err := s.Func(ctx, canonical)
	if err != nil {
		return "", "", fmt.Errorf("sign request: %w", err)
	}
	if len(sig) != ed25519.SignatureSize {
		return "", "", fmt.Errorf("ed25519 signature has %d bytes, want %d", len(sig), ed25519.SignatureSize)
	}
	return "v2=" + base64.StdEncoding.EncodeToString(sig), s.KeyID, nil
}

// Canonical is the string both request schemes sign: the timestamp, the method, the path and
// query as sent, and the hex SHA-256 of the body, joined by newlines.
func Canonical(timestamp, method, path string, body []byte) []byte {
	sum := sha256.Sum256(body)
	return []byte(timestamp + "\n" + method + "\n" + path + "\n" + hex.EncodeToString(sum[:]))
}

// APIError is an error response from the API. Use Code to handle a specific error.
type APIError struct {
	Status    int
	Type      ErrorType
	Code      string
	Message   string
	Param     string
	RequestID string
}

func (e *APIError) Error() string {
	s := fmt.Sprintf("lemak: %d %s", e.Status, e.Code)
	if e.Message != "" {
		s += ": " + e.Message
	}
	if e.RequestID != "" {
		s += " (" + e.RequestID + ")"
	}
	return s
}

func apiError(res *http.Response, body []byte) *APIError {
	e := &APIError{Status: res.StatusCode, RequestID: res.Header.Get("Request-Id")}
	var env ErrorEnvelope
	if json.Unmarshal(body, &env) == nil && env.Error.Code != "" {
		e.Type, e.Code, e.Message = env.Error.Type, env.Error.Code, env.Error.Message
		if env.Error.Param != nil {
			e.Param = *env.Error.Param
		}
		if env.Error.RequestId != nil {
			e.RequestID = *env.Error.RequestId
		}
	} else {
		e.Code = http.StatusText(res.StatusCode)
	}
	return e
}

// do sends one call. params is a generated XxxParams value or nil: its form fields become the
// query and its other fields, Idempotency-Key and Api-Version, become headers.
func (c *Client) do(ctx context.Context, method, path string, params, in, out any) error {
	query, header := encode(params)
	return c.send(ctx, method, path, query, header, in, out)
}

func (c *Client) send(ctx context.Context, method, path string, query url.Values, header http.Header, in, out any) error {
	header = header.Clone()
	if q := query.Encode(); q != "" {
		path += "?" + q
	}
	var body []byte
	if in != nil {
		var err error
		if body, err = json.Marshal(in); err != nil {
			return fmt.Errorf("lemak: encode %s %s: %w", method, path, err)
		}
	}
	if method == http.MethodPost && header.Get("Idempotency-Key") == "" {
		header.Set("Idempotency-Key", rand.Text())
	}
	if header.Get("Api-Version") == "" {
		header.Set("Api-Version", or(c.APIVersion, APIVersion))
	}
	header.Set("Authorization", "Bearer "+c.Key)
	if in != nil {
		header.Set("Content-Type", "application/json")
	}

	retries := c.MaxRetries
	if retries == 0 {
		retries = defaultRetries
	}
	var signedAt int64
	for attempt := 0; ; attempt++ {
		// Ed25519 is deterministic, so a retry inside the same second would repeat the
		// signature and be refused as a replay. One second ahead is well inside the skew.
		signedAt = max(time.Now().Unix(), signedAt+1)
		res, data, err := c.attempt(ctx, method, path, header, body, signedAt)
		last := attempt >= retries
		switch {
		case err != nil && (last || ctx.Err() != nil):
			return err
		case err != nil:
			err = sleep(ctx, backoff(attempt))
		case res.StatusCode == http.StatusTooManyRequests || res.StatusCode >= 500:
			if last {
				return apiError(res, data)
			}
			err = sleep(ctx, retryAfter(res.Header.Get("Retry-After"), attempt))
		case res.StatusCode >= 400:
			return apiError(res, data)
		default:
			if out == nil {
				return nil
			}
			if err := json.Unmarshal(data, out); err != nil {
				return fmt.Errorf("lemak: decode %s %s: %w", method, path, err)
			}
			return nil
		}
		if err != nil {
			return err
		}
	}
}

// attempt sends one try. The timestamp and the signature are fresh on every try: the server
// accepts one (timestamp, signature) pair once, and the Idempotency-Key is what makes the
// retry safe.
func (c *Client) attempt(ctx context.Context, method, path string, header http.Header, body []byte, signedAt int64) (*http.Response, []byte, error) {
	ctx, cancel := context.WithTimeout(ctx, or(c.Timeout, defaultTimeout))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(or(c.BaseURL, DefaultBaseURL), "/")+path, bytes.NewReader(body))
	if err != nil {
		return nil, nil, fmt.Errorf("lemak: %w", err)
	}
	req.Header = header.Clone()
	if c.Signer != nil {
		ts := strconv.FormatInt(signedAt, 10)
		sig, keyID, err := c.Signer.Sign(ctx, Canonical(ts, method, path, body))
		if err != nil {
			return nil, nil, fmt.Errorf("lemak: sign: %w", err)
		}
		req.Header.Set("X-Timestamp", ts)
		req.Header.Set("X-Signature", sig)
		if keyID != "" {
			req.Header.Set("X-Key-Id", keyID)
		}
	}
	hc := c.HTTPClient
	if hc == nil {
		hc = http.DefaultClient
	}
	res, err := hc.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("lemak: %s %s: %w", method, path, err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(io.LimitReader(res.Body, maxBody))
	if err != nil {
		return nil, nil, fmt.Errorf("lemak: read %s %s: %w", method, path, err)
	}
	return res, data, nil
}

func or[T comparable](v, def T) T {
	var zero T
	if v == zero {
		return def
	}
	return v
}

func backoff(attempt int) time.Duration {
	d := min(backoffBase<<attempt, backoffCap)
	return d/2 + mrand.N(d/2+1)
}

// retryAfter honours the server's Retry-After, in seconds or as a date, over the backoff.
func retryAfter(v string, attempt int) time.Duration {
	if s, err := strconv.Atoi(v); err == nil && s >= 0 {
		return time.Duration(s) * time.Second
	}
	if t, err := http.ParseTime(v); err == nil {
		return max(time.Until(t), 0)
	}
	return backoff(attempt)
}

func sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// encode reads a generated parameter struct. A field with a form tag is a query parameter;
// any other field is a header, named by its json tag.
func encode(params any) (url.Values, http.Header) {
	q, h := url.Values{}, http.Header{}
	v := reflect.ValueOf(params)
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return q, h
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return q, h
	}
	for i := range v.NumField() {
		f, fv := v.Type().Field(i), v.Field(i)
		if fv.Kind() == reflect.Pointer {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}
		s := format(fv)
		if s == "" {
			continue
		}
		if name, _, _ := strings.Cut(f.Tag.Get("form"), ","); name != "" {
			q.Set(name, s)
		} else if name, _, _ := strings.Cut(f.Tag.Get("json"), ","); name != "" {
			h.Set(name, s)
		}
	}
	return q, h
}

func format(v reflect.Value) string {
	if t, ok := v.Interface().(time.Time); ok {
		return t.UTC().Format(time.RFC3339)
	}
	return fmt.Sprint(v.Interface())
}
