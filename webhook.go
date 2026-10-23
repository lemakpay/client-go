package lemak

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// WebhookHeader carries both signatures of a delivery.
const WebhookHeader = "X-Webhook-Signature"

// WebhookTolerance is the maximum difference between a delivery's timestamp and the current time.
const WebhookTolerance = 300 * time.Second

var (
	// ErrWebhookSignature covers a missing, malformed or wrong signature, and a call with no key.
	ErrWebhookSignature = errors.New("lemak: webhook signature rejected")
	// ErrWebhookStale means the timestamp is outside WebhookTolerance.
	ErrWebhookStale = errors.New("lemak: webhook timestamp outside the tolerance")
)

// WebhookKeys is what a receiver verifies with. Secret is the endpoint secret, 64 hex
// characters, and checks v1. PublicKeys maps a key id to its base64 public key, as
// WebhookPublicKeys returns them, and checks v2 against the key the header names in v2_kid.
// Each half that is given must match, and at least one must be given. A v2_kid that is not in
// PublicKeys skips v2: fetch the keys again when you meet one.
type WebhookKeys struct {
	Secret     string
	PublicKeys map[string]string
}

// PublicKeyMap turns the answer of WebhookPublicKeys into WebhookKeys.PublicKeys.
func PublicKeyMap(l *WebhookPublicKeyList) map[string]string {
	m := make(map[string]string, len(l.Data))
	for _, k := range l.Data {
		m[k.Id] = k.PublicKey
	}
	return m
}

// VerifyWebhook checks the X-Webhook-Signature header of one delivery against the raw body
// exactly as it arrived. Unknown header fields are ignored.
func VerifyWebhook(header string, body []byte, keys WebhookKeys, now time.Time) error {
	f := make(map[string]string)
	for _, field := range strings.Split(header, ",") {
		k, v, ok := strings.Cut(strings.TrimSpace(field), "=")
		if !ok {
			return fmt.Errorf("field %q is not k=v: %w", field, ErrWebhookSignature)
		}
		f[k] = v
	}
	sec, err := strconv.ParseInt(f["t"], 10, 64)
	if err != nil {
		return fmt.Errorf("timestamp %q: %w", f["t"], ErrWebhookSignature)
	}
	if d := now.Sub(time.Unix(sec, 0)); d > WebhookTolerance || d < -WebhookTolerance {
		return fmt.Errorf("timestamp is %s from now: %w", d.Round(time.Second), ErrWebhookStale)
	}
	msg := append([]byte(f["t"]+"."), body...)

	checked := false
	if keys.Secret != "" {
		secret, err := hex.DecodeString(keys.Secret)
		if err != nil {
			return fmt.Errorf("endpoint secret is not hex: %w", ErrWebhookSignature)
		}
		want, err := hex.DecodeString(f["v1"])
		if err != nil {
			return fmt.Errorf("v1 is not hex: %w", ErrWebhookSignature)
		}
		mac := hmac.New(sha256.New, secret)
		mac.Write(msg)
		if !hmac.Equal(mac.Sum(nil), want) {
			return fmt.Errorf("v1 does not match: %w", ErrWebhookSignature)
		}
		checked = true
	}
	if b64, ok := keys.PublicKeys[f["v2_kid"]]; ok {
		pub, err := base64.StdEncoding.DecodeString(b64)
		if err != nil || len(pub) != ed25519.PublicKeySize {
			return fmt.Errorf("public key %s is not 32 bytes of base64: %w", f["v2_kid"], ErrWebhookSignature)
		}
		sig, err := base64.StdEncoding.DecodeString(f["v2"])
		if err != nil {
			return fmt.Errorf("v2 is not base64: %w", ErrWebhookSignature)
		}
		if !ed25519.Verify(pub, msg, sig) {
			return fmt.Errorf("v2 does not match: %w", ErrWebhookSignature)
		}
		checked = true
	}
	if !checked {
		return fmt.Errorf("no key to verify with: %w", ErrWebhookSignature)
	}
	return nil
}

// EventType names what happened, the object and its new status joined by a dot.
type EventType string

const (
	EventPaymentIntentCreated         EventType = "payment_intent.created"
	EventPaymentIntentRequiresPayment EventType = "payment_intent.requires_payment"
	EventPaymentIntentProcessing      EventType = "payment_intent.processing"
	EventPaymentIntentSucceeded       EventType = "payment_intent.succeeded"
	EventPaymentIntentChannelSettled  EventType = "payment_intent.channel_settled"
	EventPaymentIntentFailed          EventType = "payment_intent.failed"
	EventPaymentIntentExpired         EventType = "payment_intent.expired"
	EventPaymentIntentCanceled        EventType = "payment_intent.canceled"
	// EventPaymentIntentFinalized: a crypto payment can no longer be taken back by the chain.
	EventPaymentIntentFinalized EventType = "payment_intent.finalized"
	// EventPaymentIntentReversed: the chain took a credited crypto payment back before it was final.
	EventPaymentIntentReversed EventType = "payment_intent.reversed"
	EventSettlementPending     EventType = "settlement.pending"
	EventSettlementProcessing  EventType = "settlement.processing"
	EventSettlementPaid        EventType = "settlement.paid"
	EventSettlementFailed      EventType = "settlement.failed"
	EventDisputeOpened         EventType = "dispute.opened"
	EventDisputeUpdated        EventType = "dispute.updated"
	EventDisputeClosed         EventType = "dispute.closed"
	EventRefundCreated         EventType = "refund.created"
	EventRefundSucceeded       EventType = "refund.succeeded"
	EventRefundFailed          EventType = "refund.failed"
)

// EventTypes lists every type the platform delivers.
func EventTypes() []EventType {
	return []EventType{
		EventPaymentIntentCreated, EventPaymentIntentRequiresPayment, EventPaymentIntentProcessing,
		EventPaymentIntentSucceeded, EventPaymentIntentChannelSettled, EventPaymentIntentFailed,
		EventPaymentIntentExpired, EventPaymentIntentCanceled, EventPaymentIntentFinalized, EventPaymentIntentReversed,
		EventSettlementPending, EventSettlementProcessing, EventSettlementPaid, EventSettlementFailed,
		EventDisputeOpened, EventDisputeUpdated, EventDisputeClosed,
		EventRefundCreated, EventRefundSucceeded, EventRefundFailed,
	}
}

// WebhookEvent is one delivered event. Data is a notice, not the object: read the object with
// the API before you act on it.
type WebhookEvent struct {
	ID       string    `json:"id"`
	Type     EventType `json:"type"`
	Created  time.Time `json:"created"`
	Livemode bool      `json:"livemode"`
	Data     EventData `json:"data"`
}

// EventData identifies the object an event is about. Amount is set only on settlement events.
type EventData struct {
	ID       string `json:"id"`
	Object   string `json:"object"`
	Status   string `json:"status"`
	Merchant string `json:"merchant"`
	Mode     string `json:"mode"`
	Amount   *Money `json:"amount,omitempty"`
}

// ParseEvent decodes a verified delivery body. Unknown fields are ignored, and so is an
// unknown type: compare Type with the constants and skip what you do not handle.
func ParseEvent(body []byte) (*WebhookEvent, error) {
	var e WebhookEvent
	if err := json.Unmarshal(body, &e); err != nil {
		return nil, fmt.Errorf("lemak: decode event: %w", err)
	}
	if e.ID == "" || e.Type == "" {
		return nil, errors.New("lemak: the body is not an event")
	}
	return &e, nil
}
