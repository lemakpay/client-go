package lemak

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

// recorder answers every request with the next scripted reply, or 200 {} once the script is
// spent, and keeps each request it saw.
type recorder struct {
	mu      sync.Mutex
	replies []func(w http.ResponseWriter)
	seen    []*http.Request
	bodies  []string
}

func (r *recorder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b, _ := io.ReadAll(req.Body)
	r.mu.Lock()
	r.seen = append(r.seen, req)
	r.bodies = append(r.bodies, string(b))
	var reply func(http.ResponseWriter)
	if len(r.replies) > 0 {
		reply, r.replies = r.replies[0], r.replies[1:]
	}
	r.mu.Unlock()
	if reply == nil {
		w.Write([]byte(`{}`))
		return
	}
	reply(w)
}

func serve(t *testing.T, replies ...func(w http.ResponseWriter)) (*Client, *recorder) {
	r := &recorder{replies: replies}
	srv := httptest.NewServer(r)
	t.Cleanup(srv.Close)
	return &Client{Key: "sk_test_X", BaseURL: srv.URL}, r
}

func status(code int, body string, header ...string) func(w http.ResponseWriter) {
	return func(w http.ResponseWriter) {
		for i := 0; i+1 < len(header); i += 2 {
			w.Header().Set(header[i], header[i+1])
		}
		w.WriteHeader(code)
		w.Write([]byte(body))
	}
}

// Every merchant route of the spec has a method, and the method reaches that route.
func TestEveryRouteIsCovered(t *testing.T) {
	ctx := context.Background()
	c, r := serve(t)
	calls := map[string]func(){
		"createPaymentIntent": func() { c.CreatePaymentIntent(ctx, PaymentIntentCreate{}, nil) },
		"listPaymentIntents": func() {
			for range c.ListPaymentIntents(ctx, nil) {
			}
		},
		"paymentIntent":         func() { c.PaymentIntent(ctx, "pi_A", nil) },
		"cancelPaymentIntent":   func() { c.CancelPaymentIntent(ctx, "pi_A", nil) },
		"createPaymentAttempt":  func() { c.CreatePaymentAttempt(ctx, "pi_A", CreatePaymentAttemptJSONRequestBody{}, nil) },
		"listPaymentAttempts":   func() { c.ListPaymentAttempts(ctx, "pi_A", nil) },
		"simulatePaymentIntent": func() { c.SimulatePaymentIntent(ctx, "pi_A", PaymentIntentSimulate{}, nil) },
		"advancePaymentIntent":  func() { c.AdvancePaymentIntent(ctx, "pi_A", PaymentIntentAdvance{}, nil) },
		"listDisputes": func() {
			for range c.ListDisputes(ctx, nil) {
			}
		},
		"dispute":               func() { c.Dispute(ctx, "dp_A", nil) },
		"submitDisputeEvidence": func() { c.SubmitDisputeEvidence(ctx, "dp_A", DisputeEvidence{}, nil) },
		"createSettlement":      func() { c.CreateSettlement(ctx, nil) },
		"balance":               func() { c.Balance(ctx, nil) },
		"fxQuote":               func() { c.FxQuote(ctx, FxQuoteParams{From: "THB", To: "USD"}) },
		"events": func() {
			for range c.Events(ctx, nil) {
			}
		},
		"webhookPublicKeys": func() { c.WebhookPublicKeys(ctx, nil) },
	}
	for _, route := range load(t).Routes {
		t.Run(route.Operation, func(t *testing.T) {
			call, ok := calls[route.Operation]
			if !ok {
				t.Fatalf("no method calls %s %s", route.Method, route.Path)
			}
			delete(calls, route.Operation)
			n := len(r.seen)
			call()
			if len(r.seen) != n+1 {
				t.Fatalf("%d requests, want 1", len(r.seen)-n)
			}
			req := r.seen[n]
			want := "^" + regexp.MustCompile(`\\\{[a-z_]+\\\}`).ReplaceAllString(regexp.QuoteMeta(route.Path), `[^/]+`) + "$"
			if req.Method != route.Method || !regexp.MustCompile(want).MatchString(req.URL.Path) {
				t.Errorf("%s %s, want %s %s", req.Method, req.URL.Path, route.Method, route.Path)
			}
		})
	}
	for op := range calls {
		t.Errorf("%s is not a merchant route of the spec", op)
	}
}

func TestHeaders(t *testing.T) {
	c, r := serve(t)
	ctx := context.Background()
	c.CreatePaymentIntent(ctx, PaymentIntentCreate{Amount: 1}, nil)
	v := "2027-01-01"
	c.CreatePaymentIntent(ctx, PaymentIntentCreate{Amount: 1}, &CreatePaymentIntentParams{IdempotencyKey: "mine", ApiVersion: &v})
	c.Balance(ctx, nil)

	first, second, get := r.seen[0].Header, r.seen[1].Header, r.seen[2].Header
	if got := first.Get("Authorization"); got != "Bearer sk_test_X" {
		t.Errorf("Authorization = %q", got)
	}
	if got := first.Get("Idempotency-Key"); len(got) < 20 {
		t.Errorf("generated Idempotency-Key = %q, want a random one", got)
	}
	if got := first.Get("Api-Version"); got != APIVersion {
		t.Errorf("Api-Version = %q, want %q", got, APIVersion)
	}
	if got := second.Get("Idempotency-Key"); got != "mine" {
		t.Errorf("Idempotency-Key = %q, want the caller's", got)
	}
	if got := second.Get("Api-Version"); got != v {
		t.Errorf("Api-Version = %q, want the caller's %q", got, v)
	}
	if got := get.Get("Idempotency-Key"); got != "" {
		t.Errorf("a GET carries Idempotency-Key %q", got)
	}
	if got := get.Get("X-Signature"); got != "" {
		t.Errorf("an unsigned client sent X-Signature %q", got)
	}
}

func TestQuery(t *testing.T) {
	c, r := serve(t)
	limit, status := 50, PaymentIntentStatusSucceeded
	since := time.Date(2026, 10, 1, 7, 0, 0, 0, time.FixedZone("ICT", 7*3600))
	for range c.ListPaymentIntents(context.Background(), &ListPaymentIntentsParams{Limit: &limit, Status: &status, CreatedGte: &since}) {
	}
	if got, want := r.seen[0].URL.RawQuery, "created%5Bgte%5D=2026-10-01T00%3A00%3A00Z&limit=50&status=succeeded"; got != want {
		t.Errorf("query = %s, want %s", got, want)
	}
}

// A retry keeps the Idempotency-Key and the body, and signs again with a later timestamp: the
// server accepts one timestamp and signature once.
func TestRetryResigns(t *testing.T) {
	c, r := serve(t,
		status(503, `{"error":{"type":"api_error","code":"service_unavailable","message":"x"}}`, "Retry-After", "0"),
		status(429, `{"error":{"type":"rate_limit_error","code":"rate_limit_exceeded","message":"x"}}`, "Retry-After", "0"),
		status(201, `{"id":"pi_A","status":"requires_payment"}`))
	_, priv, _ := ed25519.GenerateKey(nil)
	c.Signer = Ed25519Signer{KeyID: "key_A", Key: priv}

	pi, err := c.CreatePaymentIntent(context.Background(), PaymentIntentCreate{Amount: 1, Currency: "THB", MerchantOrderId: "o"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if pi.Id != "pi_A" {
		t.Errorf("id = %s", pi.Id)
	}
	if len(r.seen) != 3 {
		t.Fatalf("%d attempts, want 3", len(r.seen))
	}
	var prev string
	for i, req := range r.seen {
		h := req.Header
		if h.Get("Idempotency-Key") != r.seen[0].Header.Get("Idempotency-Key") || r.bodies[i] != r.bodies[0] {
			t.Errorf("attempt %d changed the key or the body", i)
		}
		ts := h.Get("X-Timestamp")
		if ts <= prev {
			t.Errorf("attempt %d timestamp %s, want later than %s", i, ts, prev)
		}
		prev = ts
		sig, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(h.Get("X-Signature"), "v2="))
		if !ed25519.Verify(priv.Public().(ed25519.PublicKey), Canonical(ts, req.Method, req.URL.RequestURI(), []byte(r.bodies[i])), sig) ||
			h.Get("X-Key-Id") != "key_A" {
			t.Errorf("attempt %d signature does not verify", i)
		}
	}
}

func TestNetworkErrorIsRetried(t *testing.T) {
	var once sync.Once
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dropped := false
		once.Do(func() {
			conn, _, _ := w.(http.Hijacker).Hijack()
			conn.Close()
			dropped = true
		})
		if !dropped {
			w.Write([]byte(`{"available":[]}`))
		}
	}))
	defer srv.Close()
	c := &Client{Key: "k", BaseURL: srv.URL}
	if _, err := c.Balance(context.Background(), nil); err != nil {
		t.Fatalf("balance after one dropped connection: %v", err)
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		name    string
		reply   func(http.ResponseWriter)
		retries int
		want    APIError
		tries   int
	}{
		{"a 4xx is not retried", status(422,
			`{"error":{"type":"idempotency_error","code":"idempotency_key_reused","message":"m","param":"Idempotency-Key","request_id":"req_A"}}`),
			2, APIError{422, ErrorTypeIdempotencyError, "idempotency_key_reused", "m", "Idempotency-Key", "req_A"}, 1},
		{"a 5xx stops after the retries", status(500, `{"error":{"type":"api_error","code":"api_error","message":"m"}}`, "Request-Id", "req_B", "Retry-After", "0"),
			1, APIError{500, ErrorTypeApiError, "api_error", "m", "", "req_B"}, 2},
		{"a body that is not an envelope", status(502, `<html>`), -1, APIError{Status: 502, Code: "Bad Gateway"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, r := serve(t, tt.reply, tt.reply, tt.reply)
			c.MaxRetries = tt.retries
			_, err := c.Balance(context.Background(), nil)
			var got *APIError
			if !errors.As(err, &got) {
				t.Fatalf("err = %v, want an *APIError", err)
			}
			if *got != tt.want {
				t.Errorf("err = %+v, want %+v", *got, tt.want)
			}
			if len(r.seen) != tt.tries {
				t.Errorf("%d attempts, want %d", len(r.seen), tt.tries)
			}
		})
	}
}

func TestTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-time.After(time.Second):
		}
	}))
	defer srv.Close()
	c := &Client{Key: "k", BaseURL: srv.URL, Timeout: 20 * time.Millisecond, MaxRetries: -1}
	if _, err := c.Balance(context.Background(), nil); !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want a deadline", err)
	}
}

func TestPagination(t *testing.T) {
	c, r := serve(t,
		status(200, `{"object":"list","data":[{"id":"evt_A"},{"id":"evt_B"}],"has_more":true}`),
		status(200, `{"object":"list","data":[{"id":"evt_C"}],"has_more":false}`))
	limit := 2
	var got []string
	for e, err := range c.Events(context.Background(), &EventsParams{Limit: &limit}) {
		if err != nil {
			t.Fatal(err)
		}
		got = append(got, e.Id)
	}
	if fmt.Sprint(got) != "[evt_A evt_B evt_C]" {
		t.Errorf("events = %v", got)
	}
	if q := r.seen[1].URL.Query(); q.Get("starting_after") != "evt_B" || q.Get("limit") != "2" {
		t.Errorf("second page query = %v, want starting_after=evt_B and limit=2", q)
	}
}
