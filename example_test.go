package lemak_test

import (
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	lemak "github.com/lemakpay/client-go"
)

func Example() {
	ctx := context.Background()
	c := &lemak.Client{Key: "sk_test_MFWZLFDEM5X3WDBPM7ULRVDKCE"}

	method := "promptpay"
	pi, err := c.CreatePaymentIntent(ctx, lemak.PaymentIntentCreate{
		Amount:          100000, // 1,000.00 THB in minor units
		Currency:        "THB",
		MerchantOrderId: "order-1001",
		PaymentMethod:   &method,
	}, nil)
	var apiErr *lemak.APIError
	if errors.As(err, &apiErr) && apiErr.Code == "merchant_order_id_taken" {
		log.Fatal("this order already has a payment intent")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pi.Id, pi.Status, *pi.CheckoutUrl)
}

func ExampleClient_ListPaymentIntents() {
	c := &lemak.Client{Key: "sk_test_MFWZLFDEM5X3WDBPM7ULRVDKCE"}
	status := lemak.PaymentIntentStatusSucceeded
	for pi, err := range c.ListPaymentIntents(context.Background(), &lemak.ListPaymentIntentsParams{Status: &status}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pi.Id, pi.Amount, pi.Currency)
	}
}

func ExampleEd25519Signer() {
	var priv ed25519.PrivateKey // your key; register its public half in the dashboard
	c := &lemak.Client{
		Key:    "sk_live_…",
		Signer: lemak.Ed25519Signer{KeyID: "key_…", Key: priv},
	}
	_ = c
}

func ExampleVerifyWebhook() {
	keys := lemak.WebhookKeys{Secret: "<the endpoint secret, 64 hex characters>"}
	http.HandleFunc("/lemak/webhook", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
		if err != nil {
			http.Error(w, "read", http.StatusBadRequest)
			return
		}
		if err := lemak.VerifyWebhook(r.Header.Get(lemak.WebhookHeader), body, keys, time.Now()); err != nil {
			http.Error(w, "signature", http.StatusBadRequest)
			return
		}
		e, err := lemak.ParseEvent(body)
		if err != nil {
			http.Error(w, "event", http.StatusBadRequest)
			return
		}
		if e.Type == lemak.EventPaymentIntentSucceeded {
			// Store e.ID to deduplicate, answer 200, then read the intent with the API.
		}
		w.WriteHeader(http.StatusOK)
	})
}
