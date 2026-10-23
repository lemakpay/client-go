# Lemak Go Library

The Lemak Go library provides access to the Lemak API from applications written in Go. It covers
every `/v1` endpoint, signs requests, retries safely and verifies webhook signatures. It has no
dependencies outside the Go standard library.

## Requirements

Go 1.24 or later.

## Installation

```sh
go get github.com/lemakpay/client-go@v0.6.0
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

### Automatic currency conversion

To price an order in your settlement currency and let the payer pay in theirs, set
`PaymentCurrency`:

```go
method, thb := "promptpay", lemak.Currency("THB")
bearer := lemak.PaymentIntentCreateMarkupBearerPayer
pi, err := c.CreatePaymentIntent(ctx, lemak.PaymentIntentCreate{
	Amount:          1000, // 10.00 USD, your price
	Currency:        "USD",
	PaymentCurrency: &thb,
	MarkupBearer:    &bearer,
	MerchantOrderId: "order-1002",
	PaymentMethod:   &method,
}, nil)
// pi.PaymentAmount is what the payer pays, in THB.
```

The payer pays in another currency and you are paid in yours. The exchange rate is always locked
when the payment intent is created, whatever your lock point setting, and it holds until the
payment intent expires. `currency` must be your settlement currency, and `payment_method` is
required. The response keeps your price in `amount` and `currency`; the payer's side is in
`payment_amount`, and the applied rate, markup included, is in `fx.rate`.

`markup_bearer` decides who pays the conversion markup and your fee:

- `payer` (the default): the payer pays them, so you receive exactly `amount`. The payment
  amount is rounded up to the payment currency's smallest unit.
- `merchant`: the payer pays `amount` converted without markup, rounded up. You receive the
  payment amount less your fee, converted at `fx.rate` and rounded down.

### What the payer is charged

A payment method can charge only in steps and above a minimum. Each method lists these rules per
currency in `TransactionLimits`:

```go
methods, err := c.ListPaymentMethods(ctx, nil)
if err != nil {
	return err
}
for _, m := range methods.Data {
	if l, ok := m.TransactionLimits["THB"]; ok {
		fmt.Println(m.Code, l.RoundingIncrement, l.MinimumAmount) // MinimumAmount is nil for none
	}
}
```

The payer's charge is `amount` rounded up to the method's `rounding_increment`, and at least its
`minimum_amount`. `amount` never changes. Once a payment method is chosen, the payment intent
shows the charge in `PresentmentDetails`:

```go
if d := pi.PresentmentDetails; d != nil {
	fmt.Println(d.PresentmentAmount, d.PresentmentCurrency)
}
```

A method that is not enabled for your account in that currency returns 400
`payment_method_unavailable`.

List endpoints return an iterator that fetches pages as needed:

```go
for pi, err := range c.ListPaymentIntents(ctx, nil) {
	if err != nil {
		return err
	}
	fmt.Println(pi.Id)
}
```

### Invoices and receipts

Every payment intent carries an invoice. Pass `ReceiptEmail` on create, and the payer gets a
receipt by email when the payment succeeds:

```go
email := "payer@example.com"
pi, err := c.CreatePaymentIntent(ctx, lemak.PaymentIntentCreate{
	Amount:          100000,
	Currency:        "THB",
	MerchantOrderId: "order-1001",
	ReceiptEmail:    &email,
}, nil)
if err != nil {
	return err
}
if inv := pi.Invoice; inv != nil {
	fmt.Println(inv.Number, inv.HostedInvoiceUrl) // share the URL with the payer
}
if r := pi.Receipt; r != nil { // present from succeeded on
	fmt.Println(r.Number, r.ReceiptUrl)
}
```

`ReceiptEmail` is not returned. A value that is not one email address returns 400
`invalid_receipt_email`. The hosted invoice and receipt pages have a Download PDF button.

### Refunds

```go
amount := int64(50000)
re, err := c.CreateRefund(ctx, lemak.RefundCreate{PaymentIntent: pi.Id, Amount: &amount}, nil)
```

Leave `Amount` nil to refund all that is left. `ListRefunds` and `Refund` read refunds back, and
the `refund.created`, `refund.succeeded` and `refund.failed` events report their progress.

### Disputes and evidence files

Upload a file, then name it under one of the file keys of the evidence:

```go
f, err := os.Open("receipt.pdf")
if err != nil {
	return err
}
defer f.Close()
file, err := c.CreateFile(ctx, lemak.FilePurposeDisputeEvidence, "receipt.pdf", f, nil)
if err != nil {
	return err
}
dp, err := c.SubmitDisputeEvidence(ctx, disputeID, lemak.DisputeEvidence{
	Evidence: map[string]any{lemak.EvidenceReceipt: file.Id},
}, nil)
```

A file is a PDF, a JPEG or a PNG of at most 5 MB. To accept a dispute instead, call
`CloseDispute`.

### Balance transactions and statements

`ListBalanceTransactions` iterates over every movement of your balance. `Statement` returns the
statement of one month as CSV or as a printable HTML page:

```go
csv, err := c.Statement(ctx, lemak.StatementParams{Month: "2026-09", Format: lemak.StatementParamsFormatCsv})
```

### Withdrawals

On a manual settlement schedule, `CreateSettlement` withdraws from the available balance. In
production it needs a withdrawal key (`wk_live_...`) as the client's `Key`. If your approval
policy holds the withdrawal for a second approver, you get the pending request instead of a
settlement:

```go
amount, usd := int64(100000), lemak.SettlementCreateCurrencyUSD
st, req, err := c.CreateSettlement(ctx, &lemak.SettlementCreate{Amount: &amount, Currency: &usd}, nil)
if err != nil {
	return err
}
if req != nil {
	fmt.Println("waiting for approval:", req.Id)
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
`RequestID`. Use `Code` to handle a specific error. A `payment_blocked` error also carries
`DeclineCode`, the kind of risk rule that declined the payment.

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
