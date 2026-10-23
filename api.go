package lemak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"mime/multipart"
	"net/http"
	"net/url"
)

// Every method takes the generated XxxParams of its route as its last argument, and nil is
// always allowed. Its IdempotencyKey, when set, replaces the generated one; its ApiVersion
// replaces Client.APIVersion for that call.

func (c *Client) CreatePaymentIntent(ctx context.Context, in PaymentIntentCreate, p *CreatePaymentIntentParams) (*PaymentIntent, error) {
	return call[PaymentIntent](ctx, c, http.MethodPost, "/v1/payment_intents", p, in)
}

// ListPaymentIntents iterates over every payment intent the filter matches, newest first,
// fetching pages as it goes. p.Limit sets the page size.
func (c *Client) ListPaymentIntents(ctx context.Context, p *ListPaymentIntentsParams) iter.Seq2[PaymentIntent, error] {
	return pages(ctx, c, "/v1/payment_intents", p, func(v PaymentIntent) string { return v.Id })
}

func (c *Client) PaymentIntent(ctx context.Context, id string, p *PaymentIntentParams) (*PaymentIntent, error) {
	return call[PaymentIntent](ctx, c, http.MethodGet, "/v1/payment_intents/"+url.PathEscape(id), p, nil)
}

func (c *Client) CancelPaymentIntent(ctx context.Context, id string, p *CancelPaymentIntentParams) (*PaymentIntent, error) {
	return call[PaymentIntent](ctx, c, http.MethodPost, "/v1/payment_intents/"+url.PathEscape(id)+"/cancel", p, nil)
}

func (c *Client) CreatePaymentAttempt(ctx context.Context, id string, in CreatePaymentAttemptJSONRequestBody, p *CreatePaymentAttemptParams) (*PaymentAttempt, error) {
	return call[PaymentAttempt](ctx, c, http.MethodPost, "/v1/payment_intents/"+url.PathEscape(id)+"/attempts", p, in)
}

// ListPaymentAttempts returns every attempt on one intent, oldest first. The list has no cursor.
func (c *Client) ListPaymentAttempts(ctx context.Context, id string, p *ListPaymentAttemptsParams) (*PaymentAttemptList, error) {
	return call[PaymentAttemptList](ctx, c, http.MethodGet, "/v1/payment_intents/"+url.PathEscape(id)+"/attempts", p, nil)
}

// SimulatePaymentIntent reports an outcome for a test-mode intent. A live key is refused.
// An intent with no pending attempt returns 409 no_pending_attempt: create an attempt
// with CreatePaymentAttempt, or let the payer pick a method on the hosted page, first.
func (c *Client) SimulatePaymentIntent(ctx context.Context, id string, in PaymentIntentSimulate, p *SimulatePaymentIntentParams) (*PaymentIntent, error) {
	return call[PaymentIntent](ctx, c, http.MethodPost, "/v1/test/payment_intents/"+url.PathEscape(id)+"/simulate", p, in)
}

// AdvancePaymentIntent moves one timer of a test-mode intent to now.
func (c *Client) AdvancePaymentIntent(ctx context.Context, id string, in PaymentIntentAdvance, p *AdvancePaymentIntentParams) (*PaymentIntent, error) {
	return call[PaymentIntent](ctx, c, http.MethodPost, "/v1/test/payment_intents/"+url.PathEscape(id)+"/advance", p, in)
}

func (c *Client) ListDisputes(ctx context.Context, p *ListDisputesParams) iter.Seq2[Dispute, error] {
	return pages(ctx, c, "/v1/disputes", p, func(v Dispute) string { return v.Id })
}

func (c *Client) Dispute(ctx context.Context, id string, p *DisputeParams) (*Dispute, error) {
	return call[Dispute](ctx, c, http.MethodGet, "/v1/disputes/"+url.PathEscape(id), p, nil)
}

// CloseDispute accepts the dispute: it is closed as lost and no further evidence is taken.
func (c *Client) CloseDispute(ctx context.Context, id string, p *CloseDisputeParams) (*Dispute, error) {
	return call[Dispute](ctx, c, http.MethodPost, "/v1/disputes/"+url.PathEscape(id)+"/close", p, nil)
}

// The DisputeEvidence keys that take the ID of a file uploaded with FilePurposeDisputeEvidence.
const (
	EvidenceReceipt               = "receipt"
	EvidenceCustomerCommunication = "customer_communication"
	EvidenceShippingDocumentation = "shipping_documentation"
	EvidenceUncategorizedFile     = "uncategorized_file"
)

func (c *Client) SubmitDisputeEvidence(ctx context.Context, id string, in DisputeEvidence, p *SubmitDisputeEvidenceParams) (*Dispute, error) {
	return call[Dispute](ctx, c, http.MethodPost, "/v1/disputes/"+url.PathEscape(id)+"/evidence", p, in)
}

// CreateSettlement withdraws from the available balance. A nil in pays every released payment
// to the default destination. When the account's policy holds the withdrawal for a second
// approver, the API answers 202 with the request: then the Settlement is nil and the
// MerchantRequest is not. In production it takes a withdrawal key, wk_live_….
func (c *Client) CreateSettlement(ctx context.Context, in *SettlementCreate, p *CreateSettlementParams) (*Settlement, *MerchantRequest, error) {
	var body any
	if in != nil {
		body = in
	}
	var data []byte
	if err := c.do(ctx, http.MethodPost, "/v1/settlements", p, body, &data); err != nil {
		return nil, nil, err
	}
	var head struct {
		Object string `json:"object"`
	}
	if err := json.Unmarshal(data, &head); err != nil {
		return nil, nil, fmt.Errorf("lemak: decode POST /v1/settlements: %w", err)
	}
	if head.Object == "merchant_request" {
		var r MerchantRequest
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, nil, fmt.Errorf("lemak: decode POST /v1/settlements: %w", err)
		}
		return nil, &r, nil
	}
	var st Settlement
	if err := json.Unmarshal(data, &st); err != nil {
		return nil, nil, fmt.Errorf("lemak: decode POST /v1/settlements: %w", err)
	}
	return &st, nil, nil
}

// ListSettlements iterates over the settlements in the key's mode, newest first. Each one
// leaves out its items; Settlement returns them.
func (c *Client) ListSettlements(ctx context.Context, p *ListSettlementsParams) iter.Seq2[Settlement, error] {
	return pages(ctx, c, "/v1/settlements", p, func(v Settlement) string { return v.Id })
}

// Settlement returns one settlement, with the payment intents it pays in Items.
func (c *Client) Settlement(ctx context.Context, id string, p *SettlementParams) (*Settlement, error) {
	return call[Settlement](ctx, c, http.MethodGet, "/v1/settlements/"+url.PathEscape(id), p, nil)
}

// CreateFile uploads the contents of r, a PDF, a JPEG or a PNG of at most 5 MB, as filename.
// r is read to the end before the request is sent, so a retry sends the same bytes.
func (c *Client) CreateFile(ctx context.Context, purpose FilePurpose, filename string, r io.Reader, p *CreateFileParams) (*File, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	if err := w.WriteField("purpose", string(purpose)); err != nil {
		return nil, fmt.Errorf("lemak: encode file: %w", err)
	}
	part, err := w.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("lemak: encode file: %w", err)
	}
	if _, err := io.Copy(part, r); err != nil {
		return nil, fmt.Errorf("lemak: read %s: %w", filename, err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("lemak: encode file: %w", err)
	}
	return call[File](ctx, c, http.MethodPost, "/v1/files", p, rawBody{buf.Bytes(), w.FormDataContentType()})
}

func (c *Client) File(ctx context.Context, id string, p *FileParams) (*File, error) {
	return call[File](ctx, c, http.MethodGet, "/v1/files/"+url.PathEscape(id), p, nil)
}

// CreateRefund refunds a succeeded payment intent, in whole or in part.
func (c *Client) CreateRefund(ctx context.Context, in RefundCreate, p *CreateRefundParams) (*Refund, error) {
	return call[Refund](ctx, c, http.MethodPost, "/v1/refunds", p, in)
}

// ListRefunds iterates over the refunds the filter matches, newest first.
func (c *Client) ListRefunds(ctx context.Context, p *ListRefundsParams) iter.Seq2[Refund, error] {
	return pages(ctx, c, "/v1/refunds", p, func(v Refund) string { return v.Id })
}

func (c *Client) Refund(ctx context.Context, id string, p *RefundParams) (*Refund, error) {
	return call[Refund](ctx, c, http.MethodGet, "/v1/refunds/"+url.PathEscape(id), p, nil)
}

// ListBalanceTransactions iterates over every movement of the balance, newest first.
func (c *Client) ListBalanceTransactions(ctx context.Context, p *ListBalanceTransactionsParams) iter.Seq2[BalanceTransaction, error] {
	return pages(ctx, c, "/v1/balance_transactions", p, func(v BalanceTransaction) string { return v.Id })
}

// Statement returns the statement of one month as it is served: CSV or HTML, by p.Format.
func (c *Client) Statement(ctx context.Context, p StatementParams) ([]byte, error) {
	var data []byte
	if err := c.do(ctx, http.MethodGet, "/v1/statements", p, nil, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) Balance(ctx context.Context, p *BalanceParams) (*Balance, error) {
	return call[Balance](ctx, c, http.MethodGet, "/v1/balance", p, nil)
}

// FxQuote returns an indicative rate. The rate a payment uses is fixed at its lock point.
func (c *Client) FxQuote(ctx context.Context, p FxQuoteParams) (*FxQuote, error) {
	return call[FxQuote](ctx, c, http.MethodGet, "/v1/fx/quotes", p, nil)
}

// ListIps lists the addresses webhook deliveries leave from. The list is empty until published.
func (c *Client) ListIps(ctx context.Context) (*IpList, error) {
	return call[IpList](ctx, c, http.MethodGet, "/v1/ips", nil, nil)
}

// ListPaymentMethods lists the payment methods the key's mode can use, each with the currencies it
// serves. Set Currency to keep the methods that serve one currency.
func (c *Client) ListPaymentMethods(ctx context.Context, p *ListPaymentMethodsParams) (*PaymentMethodList, error) {
	return call[PaymentMethodList](ctx, c, http.MethodGet, "/v1/payment_methods", p, nil)
}

// ListCurrencies lists the currencies you can charge in, quote and settle in, with their exponents.
func (c *Client) ListCurrencies(ctx context.Context, p *ListCurrenciesParams) (*CurrencyList, error) {
	return call[CurrencyList](ctx, c, http.MethodGet, "/v1/currencies", p, nil)
}

// Events iterates over the events behind the webhook deliveries, newest first.
func (c *Client) Events(ctx context.Context, p *EventsParams) iter.Seq2[Event, error] {
	return pages(ctx, c, "/v1/events", p, func(v Event) string { return v.Id })
}

// WebhookPublicKeys lists the platform keys the v2 webhook signature verifies against.
func (c *Client) WebhookPublicKeys(ctx context.Context, p *WebhookPublicKeysParams) (*WebhookPublicKeyList, error) {
	return call[WebhookPublicKeyList](ctx, c, http.MethodGet, "/v1/webhook_public_keys", p, nil)
}

func call[T any](ctx context.Context, c *Client, method, path string, params, in any) (*T, error) {
	var out T
	if err := c.do(ctx, method, path, params, in, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// pages walks a cursor list. The cursor goes into the query after the caller's own parameters,
// so it replaces any starting_after the caller set once the first page is read.
func pages[T any](ctx context.Context, c *Client, path string, params any, id func(T) string) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		query, header := encode(params)
		for {
			var page struct {
				Data    []T  `json:"data"`
				HasMore bool `json:"has_more"`
			}
			if err := c.send(ctx, http.MethodGet, path, query, header, nil, &page); err != nil {
				var zero T
				yield(zero, err)
				return
			}
			for _, v := range page.Data {
				if !yield(v, nil) {
					return
				}
			}
			if !page.HasMore || len(page.Data) == 0 {
				return
			}
			query.Set("starting_after", id(page.Data[len(page.Data)-1]))
		}
	}
}
