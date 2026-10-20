# Lemak Go Library

The Lemak Go library provides access to the Lemak API from applications written in Go. It covers
every `/v1` endpoint, signs requests, retries safely and verifies webhook signatures. It has no
dependencies outside the Go standard library.

## Requirements

Go 1.24 or later.

## Installation

```sh
go get github.com/lemakpay/client-go@v0.1.0
```

```go
import lemak "github.com/lemakpay/client-go"
```

## Usage

```go
c := &lemak.Client{Key: "sk_test_..."}

method := "promptpay"
pi, err := c.CreatePaymentIntent(ctx, lemak.PaymentIntentCreate{
	Amount:          100000, // 1,000.00 THB, in minor units
	Currency:        "THB",
	MerchantOrderId: "order-1001",
	PaymentMethod:   &method,
}, nil)
if err != nil {
	return err
}
```

List endpoints return an iterator that fetches pages as needed:

```go
for pi, err := range c.ListPaymentIntents(ctx, nil) {
	if err != nil {
		return err
	}
	fmt.Println(pi.Id)
}
```

### Request signing

If your account requires signed requests, set a `Signer`:

```go
c := &lemak.Client{
	Key:    "sk_live_...",
	Signer: lemak.Ed25519Signer{KeyID: "key_...", Key: privateKey},
}
```

`lemak.HMACSigner` is available for accounts that use `hmac_sha256`.

To keep the private key in AWS KMS or a hardware security module, use `lemak.Ed25519Func`,
which calls your function to sign. With an AWS KMS key of spec `ECC_NIST_EDWARDS25519`:

```go
signer := lemak.Ed25519Func{
	KeyID: "key_...",
	Func: func(ctx context.Context, msg []byte) ([]byte, error) {
		out, err := kmsClient.Sign(ctx, &kms.SignInput{
			KeyId:            aws.String(kmsKeyID),
			Message:          msg,
			MessageType:      types.MessageTypeRaw,
			SigningAlgorithm: types.SigningAlgorithmSpecEd25519Sha512,
		})
		if err != nil {
			return nil, err
		}
		return out.Signature, nil
	},
}
```

Register the key's public key in the dashboard. `GetPublicKey` returns it DER-encoded; the
last 32 bytes are the raw Ed25519 public key.

Any type with a `Sign(ctx, canonical)` method can be used as a `Signer`.

### Idempotency and retries

Every POST request is sent with an `Idempotency-Key`. To set your own, pass it in the request
parameters. Network errors, `429` and `5xx` responses are retried twice with the same key.

### Errors

API errors are returned as `*lemak.APIError`, with `Status`, `Type`, `Code`, `Param` and
`RequestID`. Use `Code` to handle a specific error.

### Webhooks

Verify the signature against the raw request body before you use an event:

```go
err := lemak.VerifyWebhook(r.Header.Get(lemak.WebhookHeader), body,
	lemak.WebhookKeys{Secret: endpointSecret}, time.Now())
if err != nil {
	http.Error(w, "invalid signature", http.StatusBadRequest)
	return
}
event, err := lemak.ParseEvent(body)
```

## Documentation

See the [API reference](https://lemak.cash/docs/) and the
[SDK guide](https://lemak.cash/docs/sdks.html).

## License

MIT
