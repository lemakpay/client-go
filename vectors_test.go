package lemak

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"slices"
	"testing"
	"time"
)

// vectors is testdata/vectors.json, written by the Lemak API server's own signing code. The
// TypeScript client reads the same file.
type vectors struct {
	APIVersion  string `json:"api_version"`
	Routes      []struct{ Operation, Method, Path string }
	EventTypes  []EventType `json:"event_types"`
	HMACSecret  string      `json:"hmac_secret"`
	Ed25519Seed string      `json:"ed25519_seed"`
	KeyID       string      `json:"key_id"`
	Requests    []struct {
		Name, Timestamp, Method, Path, Body, Canonical, V1, V2 string
	}
	Webhooks []struct {
		Name, Header, Body, Secret, Want string
		PublicKeys                       map[string]string `json:"public_keys"`
		Now                              int64
	}
}

func load(t *testing.T) vectors {
	t.Helper()
	b, err := os.ReadFile("testdata/vectors.json")
	if err != nil {
		t.Fatal(err)
	}
	var v vectors
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	return v
}

func (v vectors) privateKey(t *testing.T) ed25519.PrivateKey {
	seed, err := hex.DecodeString(v.Ed25519Seed)
	if err != nil {
		t.Fatal(err)
	}
	return ed25519.NewKeyFromSeed(seed)
}

func TestRequestVectors(t *testing.T) {
	v := load(t)
	for _, tt := range v.Requests {
		t.Run(tt.Name, func(t *testing.T) {
			c := Canonical(tt.Timestamp, tt.Method, tt.Path, []byte(tt.Body))
			if string(c) != tt.Canonical {
				t.Fatalf("canonical = %q, want %q", c, tt.Canonical)
			}
			if got, _, _ := (HMACSigner{Secret: v.HMACSecret}).Sign(t.Context(), c); got != tt.V1 {
				t.Errorf("hmac = %s, want %s", got, tt.V1)
			}
			got, kid, _ := Ed25519Signer{KeyID: v.KeyID, Key: v.privateKey(t)}.Sign(t.Context(), c)
			if got != tt.V2 || kid != v.KeyID {
				t.Errorf("ed25519 = %s %s, want %s %s", got, kid, tt.V2, v.KeyID)
			}
			key := v.privateKey(t)
			remote := Ed25519Func{KeyID: v.KeyID, Func: func(_ context.Context, msg []byte) ([]byte, error) {
				return ed25519.Sign(key, msg), nil
			}}
			if got, kid, _ := remote.Sign(t.Context(), c); got != tt.V2 || kid != v.KeyID {
				t.Errorf("ed25519 func = %s %s, want %s %s", got, kid, tt.V2, v.KeyID)
			}
		})
	}
}

func TestWebhookVectors(t *testing.T) {
	for _, tt := range load(t).Webhooks {
		t.Run(tt.Name, func(t *testing.T) {
			err := VerifyWebhook(tt.Header, []byte(tt.Body), WebhookKeys{tt.Secret, tt.PublicKeys}, time.Unix(tt.Now, 0))
			got := "ok"
			switch {
			case errors.Is(err, ErrWebhookStale):
				got = "stale"
			case errors.Is(err, ErrWebhookSignature):
				got = "signature"
			case err != nil:
				t.Fatalf("err = %v, want one of the two sentinels", err)
			}
			if got != tt.Want {
				t.Errorf("verify = %s (%v), want %s", got, err, tt.Want)
			}
		})
	}
}

func TestParseEvent(t *testing.T) {
	e, err := ParseEvent([]byte(load(t).Webhooks[0].Body))
	if err != nil {
		t.Fatal(err)
	}
	if e.Type != EventPaymentIntentSucceeded || e.Data.Object != "payment_intent" || e.Livemode || e.Created.IsZero() {
		t.Errorf("event = %+v, want a test payment_intent.succeeded", e)
	}
	if _, err := ParseEvent([]byte(`{"hello":1}`)); err == nil {
		t.Error("a body that is not an event parsed")
	}
}

func TestConstantsMatchTheServer(t *testing.T) {
	v := load(t)
	if APIVersion != v.APIVersion {
		t.Errorf("APIVersion = %s, want %s, the version the spec pins", APIVersion, v.APIVersion)
	}
	if !slices.Equal(EventTypes(), v.EventTypes) {
		t.Errorf("EventTypes = %v, want %v, the types the server delivers", EventTypes(), v.EventTypes)
	}
}
