// Code generated from the Lemak merchant API specification. DO NOT EDIT.

package lemak

import (
	"time"
)

// Defines values for BalanceTransactionBalanceType.
const (
	BalanceTransactionBalanceTypeAvailable   BalanceTransactionBalanceType = "available"
	BalanceTransactionBalanceTypeDisputeHold BalanceTransactionBalanceType = "dispute_hold"
	BalanceTransactionBalanceTypePending     BalanceTransactionBalanceType = "pending"
	BalanceTransactionBalanceTypeReviewHold  BalanceTransactionBalanceType = "review_hold"
)

// Valid indicates whether the value is a known member of the BalanceTransactionBalanceType enum.
func (e BalanceTransactionBalanceType) Valid() bool {
	switch e {
	case BalanceTransactionBalanceTypeAvailable:
		return true
	case BalanceTransactionBalanceTypeDisputeHold:
		return true
	case BalanceTransactionBalanceTypePending:
		return true
	case BalanceTransactionBalanceTypeReviewHold:
		return true
	default:
		return false
	}
}

// Defines values for BalanceTransactionType.
const (
	BalanceTransactionTypeAdjustment     BalanceTransactionType = "adjustment"
	BalanceTransactionTypeConversion     BalanceTransactionType = "conversion"
	BalanceTransactionTypeDisputeFee     BalanceTransactionType = "dispute_fee"
	BalanceTransactionTypeDisputeHold    BalanceTransactionType = "dispute_hold"
	BalanceTransactionTypeDisputeLoss    BalanceTransactionType = "dispute_loss"
	BalanceTransactionTypeDisputeRelease BalanceTransactionType = "dispute_release"
	BalanceTransactionTypePayment        BalanceTransactionType = "payment"
	BalanceTransactionTypeRelease        BalanceTransactionType = "release"
	BalanceTransactionTypeSettlement     BalanceTransactionType = "settlement"
)

// Valid indicates whether the value is a known member of the BalanceTransactionType enum.
func (e BalanceTransactionType) Valid() bool {
	switch e {
	case BalanceTransactionTypeAdjustment:
		return true
	case BalanceTransactionTypeConversion:
		return true
	case BalanceTransactionTypeDisputeFee:
		return true
	case BalanceTransactionTypeDisputeHold:
		return true
	case BalanceTransactionTypeDisputeLoss:
		return true
	case BalanceTransactionTypeDisputeRelease:
		return true
	case BalanceTransactionTypePayment:
		return true
	case BalanceTransactionTypeRelease:
		return true
	case BalanceTransactionTypeSettlement:
		return true
	default:
		return false
	}
}

// Defines values for DestinationType.
const (
	DestinationTypeSimulated   DestinationType = "simulated"
	DestinationTypeSwift       DestinationType = "swift"
	DestinationTypeUsdcOnchain DestinationType = "usdc_onchain"
)

// Valid indicates whether the value is a known member of the DestinationType enum.
func (e DestinationType) Valid() bool {
	switch e {
	case DestinationTypeSimulated:
		return true
	case DestinationTypeSwift:
		return true
	case DestinationTypeUsdcOnchain:
		return true
	default:
		return false
	}
}

// Defines values for DisputeStatus.
const (
	DisputeStatusLost        DisputeStatus = "lost"
	DisputeStatusOpened      DisputeStatus = "opened"
	DisputeStatusUnderReview DisputeStatus = "under_review"
	DisputeStatusWon         DisputeStatus = "won"
)

// Valid indicates whether the value is a known member of the DisputeStatus enum.
func (e DisputeStatus) Valid() bool {
	switch e {
	case DisputeStatusLost:
		return true
	case DisputeStatusOpened:
		return true
	case DisputeStatusUnderReview:
		return true
	case DisputeStatusWon:
		return true
	default:
		return false
	}
}

// Defines values for ErrorDeclineCode.
const (
	ErrorDeclineCodeAmount    ErrorDeclineCode = "amount"
	ErrorDeclineCodeBlocklist ErrorDeclineCode = "blocklist"
	ErrorDeclineCodeGeneric   ErrorDeclineCode = "generic"
	ErrorDeclineCodeGeo       ErrorDeclineCode = "geo"
	ErrorDeclineCodeVelocity  ErrorDeclineCode = "velocity"
)

// Valid indicates whether the value is a known member of the ErrorDeclineCode enum.
func (e ErrorDeclineCode) Valid() bool {
	switch e {
	case ErrorDeclineCodeAmount:
		return true
	case ErrorDeclineCodeBlocklist:
		return true
	case ErrorDeclineCodeGeneric:
		return true
	case ErrorDeclineCodeGeo:
		return true
	case ErrorDeclineCodeVelocity:
		return true
	default:
		return false
	}
}

// Defines values for ErrorType.
const (
	ErrorTypeApiError            ErrorType = "api_error"
	ErrorTypeAuthenticationError ErrorType = "authentication_error"
	ErrorTypeIdempotencyError    ErrorType = "idempotency_error"
	ErrorTypeInvalidRequestError ErrorType = "invalid_request_error"
	ErrorTypeRateLimitError      ErrorType = "rate_limit_error"
)

// Valid indicates whether the value is a known member of the ErrorType enum.
func (e ErrorType) Valid() bool {
	switch e {
	case ErrorTypeApiError:
		return true
	case ErrorTypeAuthenticationError:
		return true
	case ErrorTypeIdempotencyError:
		return true
	case ErrorTypeInvalidRequestError:
		return true
	case ErrorTypeRateLimitError:
		return true
	default:
		return false
	}
}

// Defines values for FilePurpose.
const (
	FilePurposeDisputeEvidence FilePurpose = "dispute_evidence"
)

// Valid indicates whether the value is a known member of the FilePurpose enum.
func (e FilePurpose) Valid() bool {
	switch e {
	case FilePurposeDisputeEvidence:
		return true
	default:
		return false
	}
}

// Defines values for FileType.
const (
	FileTypeJpg FileType = "jpg"
	FileTypePdf FileType = "pdf"
	FileTypePng FileType = "png"
)

// Valid indicates whether the value is a known member of the FileType enum.
func (e FileType) Valid() bool {
	switch e {
	case FileTypeJpg:
		return true
	case FileTypePdf:
		return true
	case FileTypePng:
		return true
	default:
		return false
	}
}

// Defines values for FxLockLockPoint.
const (
	FxLockLockPointOnChannelSettled FxLockLockPoint = "on_channel_settled"
	FxLockLockPointOnCreate         FxLockLockPoint = "on_create"
	FxLockLockPointOnPaid           FxLockLockPoint = "on_paid"
)

// Valid indicates whether the value is a known member of the FxLockLockPoint enum.
func (e FxLockLockPoint) Valid() bool {
	switch e {
	case FxLockLockPointOnChannelSettled:
		return true
	case FxLockLockPointOnCreate:
		return true
	case FxLockLockPointOnPaid:
		return true
	default:
		return false
	}
}

// Defines values for MerchantRequestKind.
const (
	MerchantRequestKindPayoutMethodCreate  MerchantRequestKind = "payout_method.create"
	MerchantRequestKindPayoutMethodDefault MerchantRequestKind = "payout_method.default"
	MerchantRequestKindPayoutMethodRemove  MerchantRequestKind = "payout_method.remove"
	MerchantRequestKindWithdrawal          MerchantRequestKind = "withdrawal"
)

// Valid indicates whether the value is a known member of the MerchantRequestKind enum.
func (e MerchantRequestKind) Valid() bool {
	switch e {
	case MerchantRequestKindPayoutMethodCreate:
		return true
	case MerchantRequestKindPayoutMethodDefault:
		return true
	case MerchantRequestKindPayoutMethodRemove:
		return true
	case MerchantRequestKindWithdrawal:
		return true
	default:
		return false
	}
}

// Defines values for MerchantRequestMode.
const (
	MerchantRequestModeLive MerchantRequestMode = "live"
	MerchantRequestModeTest MerchantRequestMode = "test"
)

// Valid indicates whether the value is a known member of the MerchantRequestMode enum.
func (e MerchantRequestMode) Valid() bool {
	switch e {
	case MerchantRequestModeLive:
		return true
	case MerchantRequestModeTest:
		return true
	default:
		return false
	}
}

// Defines values for MerchantRequestStatus.
const (
	MerchantRequestStatusApproved MerchantRequestStatus = "approved"
	MerchantRequestStatusExpired  MerchantRequestStatus = "expired"
	MerchantRequestStatusPending  MerchantRequestStatus = "pending"
	MerchantRequestStatusRejected MerchantRequestStatus = "rejected"
)

// Valid indicates whether the value is a known member of the MerchantRequestStatus enum.
func (e MerchantRequestStatus) Valid() bool {
	switch e {
	case MerchantRequestStatusApproved:
		return true
	case MerchantRequestStatusExpired:
		return true
	case MerchantRequestStatusPending:
		return true
	case MerchantRequestStatusRejected:
		return true
	default:
		return false
	}
}

// Defines values for NextActionType.
const (
	NextActionTypeCryptoAddress NextActionType = "crypto_address"
	NextActionTypeQrCode        NextActionType = "qr_code"
	NextActionTypeRedirect      NextActionType = "redirect"
)

// Valid indicates whether the value is a known member of the NextActionType enum.
func (e NextActionType) Valid() bool {
	switch e {
	case NextActionTypeCryptoAddress:
		return true
	case NextActionTypeQrCode:
		return true
	case NextActionTypeRedirect:
		return true
	default:
		return false
	}
}

// Defines values for PaymentAttemptStatus.
const (
	PaymentAttemptStatusFailed    PaymentAttemptStatus = "failed"
	PaymentAttemptStatusPending   PaymentAttemptStatus = "pending"
	PaymentAttemptStatusSucceeded PaymentAttemptStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the PaymentAttemptStatus enum.
func (e PaymentAttemptStatus) Valid() bool {
	switch e {
	case PaymentAttemptStatusFailed:
		return true
	case PaymentAttemptStatusPending:
		return true
	case PaymentAttemptStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for PaymentIntentMarkupBearer.
const (
	PaymentIntentMarkupBearerMerchant PaymentIntentMarkupBearer = "merchant"
	PaymentIntentMarkupBearerPayer    PaymentIntentMarkupBearer = "payer"
)

// Valid indicates whether the value is a known member of the PaymentIntentMarkupBearer enum.
func (e PaymentIntentMarkupBearer) Valid() bool {
	switch e {
	case PaymentIntentMarkupBearerMerchant:
		return true
	case PaymentIntentMarkupBearerPayer:
		return true
	default:
		return false
	}
}

// Defines values for PaymentIntentMode.
const (
	PaymentIntentModeLive PaymentIntentMode = "live"
	PaymentIntentModeTest PaymentIntentMode = "test"
)

// Valid indicates whether the value is a known member of the PaymentIntentMode enum.
func (e PaymentIntentMode) Valid() bool {
	switch e {
	case PaymentIntentModeLive:
		return true
	case PaymentIntentModeTest:
		return true
	default:
		return false
	}
}

// Defines values for PaymentIntentAdvanceClock.
const (
	PaymentIntentAdvanceClockExpiry  PaymentIntentAdvanceClock = "expiry"
	PaymentIntentAdvanceClockRelease PaymentIntentAdvanceClock = "release"
)

// Valid indicates whether the value is a known member of the PaymentIntentAdvanceClock enum.
func (e PaymentIntentAdvanceClock) Valid() bool {
	switch e {
	case PaymentIntentAdvanceClockExpiry:
		return true
	case PaymentIntentAdvanceClockRelease:
		return true
	default:
		return false
	}
}

// Defines values for PaymentIntentCreateMarkupBearer.
const (
	PaymentIntentCreateMarkupBearerMerchant PaymentIntentCreateMarkupBearer = "merchant"
	PaymentIntentCreateMarkupBearerPayer    PaymentIntentCreateMarkupBearer = "payer"
)

// Valid indicates whether the value is a known member of the PaymentIntentCreateMarkupBearer enum.
func (e PaymentIntentCreateMarkupBearer) Valid() bool {
	switch e {
	case PaymentIntentCreateMarkupBearerMerchant:
		return true
	case PaymentIntentCreateMarkupBearerPayer:
		return true
	default:
		return false
	}
}

// Defines values for PaymentIntentSimulateOutcome.
const (
	PaymentIntentSimulateOutcomeDisputeLost   PaymentIntentSimulateOutcome = "dispute_lost"
	PaymentIntentSimulateOutcomeDisputeOpened PaymentIntentSimulateOutcome = "dispute_opened"
	PaymentIntentSimulateOutcomeDisputeWon    PaymentIntentSimulateOutcome = "dispute_won"
	PaymentIntentSimulateOutcomeFailed        PaymentIntentSimulateOutcome = "failed"
	PaymentIntentSimulateOutcomePending       PaymentIntentSimulateOutcome = "pending"
	PaymentIntentSimulateOutcomeSucceeded     PaymentIntentSimulateOutcome = "succeeded"
)

// Valid indicates whether the value is a known member of the PaymentIntentSimulateOutcome enum.
func (e PaymentIntentSimulateOutcome) Valid() bool {
	switch e {
	case PaymentIntentSimulateOutcomeDisputeLost:
		return true
	case PaymentIntentSimulateOutcomeDisputeOpened:
		return true
	case PaymentIntentSimulateOutcomeDisputeWon:
		return true
	case PaymentIntentSimulateOutcomeFailed:
		return true
	case PaymentIntentSimulateOutcomePending:
		return true
	case PaymentIntentSimulateOutcomeSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for PaymentIntentStatus.
const (
	PaymentIntentStatusCanceled        PaymentIntentStatus = "canceled"
	PaymentIntentStatusChannelSettled  PaymentIntentStatus = "channel_settled"
	PaymentIntentStatusCreated         PaymentIntentStatus = "created"
	PaymentIntentStatusExpired         PaymentIntentStatus = "expired"
	PaymentIntentStatusFailed          PaymentIntentStatus = "failed"
	PaymentIntentStatusProcessing      PaymentIntentStatus = "processing"
	PaymentIntentStatusRequiresPayment PaymentIntentStatus = "requires_payment"
	PaymentIntentStatusSucceeded       PaymentIntentStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the PaymentIntentStatus enum.
func (e PaymentIntentStatus) Valid() bool {
	switch e {
	case PaymentIntentStatusCanceled:
		return true
	case PaymentIntentStatusChannelSettled:
		return true
	case PaymentIntentStatusCreated:
		return true
	case PaymentIntentStatusExpired:
		return true
	case PaymentIntentStatusFailed:
		return true
	case PaymentIntentStatusProcessing:
		return true
	case PaymentIntentStatusRequiresPayment:
		return true
	case PaymentIntentStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for RefundReason.
const (
	RefundReasonDuplicate           RefundReason = "duplicate"
	RefundReasonFraudulent          RefundReason = "fraudulent"
	RefundReasonRequestedByCustomer RefundReason = "requested_by_customer"
)

// Valid indicates whether the value is a known member of the RefundReason enum.
func (e RefundReason) Valid() bool {
	switch e {
	case RefundReasonDuplicate:
		return true
	case RefundReasonFraudulent:
		return true
	case RefundReasonRequestedByCustomer:
		return true
	default:
		return false
	}
}

// Defines values for RefundStatus.
const (
	RefundStatusFailed    RefundStatus = "failed"
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusSucceeded RefundStatus = "succeeded"
)

// Valid indicates whether the value is a known member of the RefundStatus enum.
func (e RefundStatus) Valid() bool {
	switch e {
	case RefundStatusFailed:
		return true
	case RefundStatusPending:
		return true
	case RefundStatusSucceeded:
		return true
	default:
		return false
	}
}

// Defines values for SettlementMode.
const (
	SettlementModeLive SettlementMode = "live"
	SettlementModeTest SettlementMode = "test"
)

// Valid indicates whether the value is a known member of the SettlementMode enum.
func (e SettlementMode) Valid() bool {
	switch e {
	case SettlementModeLive:
		return true
	case SettlementModeTest:
		return true
	default:
		return false
	}
}

// Defines values for SettlementCreateCurrency.
const (
	SettlementCreateCurrencyUSD SettlementCreateCurrency = "USD"
)

// Valid indicates whether the value is a known member of the SettlementCreateCurrency enum.
func (e SettlementCreateCurrency) Valid() bool {
	switch e {
	case SettlementCreateCurrencyUSD:
		return true
	default:
		return false
	}
}

// Defines values for SettlementStatus.
const (
	SettlementStatusFailed     SettlementStatus = "failed"
	SettlementStatusPaid       SettlementStatus = "paid"
	SettlementStatusPending    SettlementStatus = "pending"
	SettlementStatusProcessing SettlementStatus = "processing"
)

// Valid indicates whether the value is a known member of the SettlementStatus enum.
func (e SettlementStatus) Valid() bool {
	switch e {
	case SettlementStatusFailed:
		return true
	case SettlementStatusPaid:
		return true
	case SettlementStatusPending:
		return true
	case SettlementStatusProcessing:
		return true
	default:
		return false
	}
}

// Defines values for SimulatedDecline.
const (
	SimulatedDeclineInsufficientFunds SimulatedDecline = "insufficient_funds"
	SimulatedDeclinePayerCancelled    SimulatedDecline = "payer_cancelled"
	SimulatedDeclineRiskBlocked       SimulatedDecline = "risk_blocked"
)

// Valid indicates whether the value is a known member of the SimulatedDecline enum.
func (e SimulatedDecline) Valid() bool {
	switch e {
	case SimulatedDeclineInsufficientFunds:
		return true
	case SimulatedDeclinePayerCancelled:
		return true
	case SimulatedDeclineRiskBlocked:
		return true
	default:
		return false
	}
}

// Defines values for TransactionLimitsRoundingMode.
const (
	TransactionLimitsRoundingModeUp TransactionLimitsRoundingMode = "up"
)

// Valid indicates whether the value is a known member of the TransactionLimitsRoundingMode enum.
func (e TransactionLimitsRoundingMode) Valid() bool {
	switch e {
	case TransactionLimitsRoundingModeUp:
		return true
	default:
		return false
	}
}

// Defines values for StatementParamsFormat.
const (
	StatementParamsFormatCsv  StatementParamsFormat = "csv"
	StatementParamsFormatHtml StatementParamsFormat = "html"
)

// Valid indicates whether the value is a known member of the StatementParamsFormat enum.
func (e StatementParamsFormat) Valid() bool {
	switch e {
	case StatementParamsFormatCsv:
		return true
	case StatementParamsFormatHtml:
		return true
	default:
		return false
	}
}

// Defines values for StatementParamsLang.
const (
	StatementParamsLangEn     StatementParamsLang = "en"
	StatementParamsLangZhHans StatementParamsLang = "zh-Hans"
	StatementParamsLangZhHant StatementParamsLang = "zh-Hant"
)

// Valid indicates whether the value is a known member of the StatementParamsLang enum.
func (e StatementParamsLang) Valid() bool {
	switch e {
	case StatementParamsLangEn:
		return true
	case StatementParamsLangZhHans:
		return true
	case StatementParamsLangZhHant:
		return true
	default:
		return false
	}
}

// Balance Your funds by state. Each list holds one amount per currency.
type Balance struct {
	// Available Funds that can be paid out to you.
	Available []Money `json:"available"`

	// DisputeHold Funds held for open disputes.
	DisputeHold []Money `json:"dispute_hold"`

	// Pending Funds from payments that are not yet available.
	Pending []Money `json:"pending"`

	// ReviewHold Crypto payments held for manual review. Each is usually released within one business day. Empty if you do not accept crypto.
	ReviewHold []Money `json:"review_hold"`
}

// BalanceTransaction defines model for BalanceTransaction.
type BalanceTransaction struct {
	// Amount The gross amount. Negative when money leaves this balance type. It equals `net` plus `fee`.
	Amount Money `json:"amount"`

	// AvailableOn When the funds become available. Present on the available balance, and on a pending entry of a payment once its release time is known.
	AvailableOn *time.Time `json:"available_on,omitempty"`

	// BalanceType The balance this entry moves, as GET /v1/balance names it.
	BalanceType BalanceTransactionBalanceType `json:"balance_type"`

	// Created The time the entry was booked.
	Created time.Time `json:"created"`

	// ExchangeRate The rate of a `conversion`: units of the original currency for one unit of the settlement currency, markup included. Absent on every other type.
	ExchangeRate *string `json:"exchange_rate,omitempty"`

	// Fee The fee taken from `amount`. Non-zero on a `payment` alone.
	Fee Money `json:"fee"`

	// Id Unique identifier for the balance transaction.
	Id string `json:"id"`

	// Net The amount this entry adds to the balance type. Negative when it takes.
	Net    Money       `json:"net"`
	Object interface{} `json:"object"`

	// RunningBalance The balance of this balance type and currency after this entry.
	RunningBalance Money `json:"running_balance"`

	// Source The ID of the object that caused the entry: a payment intent (`pi_`), a settlement (`st_`) or a dispute (`dp_`). Absent on an adjustment.
	Source *string `json:"source,omitempty"`

	// Type What moved the money. `payment` credits a payment net of its fee. `conversion` is the exchange to your settlement currency, one entry per currency. `release` moves funds to the next balance type, such as pending to available. `settlement` takes available funds into a settlement. `dispute_hold`, `dispute_release` and `dispute_loss` are the stages of a dispute, and `dispute_fee` is the fee a lost dispute costs. `adjustment` is every other movement.
	Type BalanceTransactionType `json:"type"`
}

// BalanceTransactionBalanceType The balance this entry moves, as GET /v1/balance names it.
type BalanceTransactionBalanceType string

// BalanceTransactionList defines model for BalanceTransactionList.
type BalanceTransactionList struct {
	Data []BalanceTransaction `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// BalanceTransactionType What moved the money. `payment` credits a payment net of its fee. `conversion` is the exchange to your settlement currency, one entry per currency. `release` moves funds to the next balance type, such as pending to available. `settlement` takes available funds into a settlement. `dispute_hold`, `dispute_release` and `dispute_loss` are the stages of a dispute, and `dispute_fee` is the fee a lost dispute costs. `adjustment` is every other movement.
type BalanceTransactionType string

// CryptoFinality Whether a crypto payment is final. Present only on crypto payment intents. On a layer 2 network, a payment succeeds when the transfer is seen and becomes final when its batch is confirmed on layer 1. Until then the chain can reverse it: the payment intent moves to `failed`, and you receive the events `payment_intent.failed` and `payment_intent.reversed`. When the payment becomes final, you receive `payment_intent.finalized` once. For a high-value order, wait for this event before you deliver.
type CryptoFinality struct {
	// Final Whether the chain can no longer reverse the payment.
	Final bool `json:"final"`

	// FinalAt The time the payment became final.
	FinalAt *time.Time `json:"final_at,omitempty"`
}

// Currency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
type Currency = string

// CurrencyDetails A currency you can charge in, get a quote for, or settle in.
type CurrencyDetails struct {
	// Chargeable Whether a payment method serves it, so you can create a payment intent in it.
	Chargeable bool `json:"chargeable"`

	// Code A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	Code Currency `json:"code"`

	// Exponent The number of decimal places of the smallest unit. An amount of `1000` is 10.00 at exponent 2 and 1000 at exponent 0.
	Exponent int         `json:"exponent"`
	Object   interface{} `json:"object"`

	// Quotable Whether GET /v1/fx/quotes prices it against the settlement currency, so a payer can pay in it when you price in the settlement currency.
	Quotable bool `json:"quotable"`

	// Settlement Whether it is the settlement currency, which is the one you price in for automatic conversion.
	Settlement bool `json:"settlement"`
}

// CurrencyList defines model for CurrencyList.
type CurrencyList struct {
	Data []CurrencyDetails `json:"data"`

	// HasMore Always `false`.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// DestinationType The kind of payout destination.
type DestinationType string

// Dispute defines model for Dispute.
type Dispute struct {
	// Amount The amount the channel claims, in the smallest unit of the currency the payer paid in.
	Amount int64 `json:"amount"`

	// ClosedAt The time the dispute was decided.
	ClosedAt *time.Time `json:"closed_at,omitempty"`

	// Currency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	Currency Currency `json:"currency"`

	// DueBy The deadline for submitting evidence.
	DueBy *time.Time `json:"due_by,omitempty"`

	// Evidence The evidence you have submitted, as key-value pairs.
	Evidence *map[string]interface{} `json:"evidence,omitempty"`

	// Fee The fee you pay if you lose the dispute, in addition to the claimed amount.
	Fee *Money `json:"fee,omitempty"`

	// Held The amount held from your balance, in the settlement currency, while the dispute is open.
	Held *Money `json:"held,omitempty"`

	// Id Unique identifier for the dispute.
	Id     string      `json:"id"`
	Object interface{} `json:"object"`

	// OpenedAt The time the dispute was opened.
	OpenedAt time.Time `json:"opened_at"`

	// PaymentIntent The ID of the disputed payment intent.
	PaymentIntent string `json:"payment_intent"`

	// Reason The reason given for the dispute.
	Reason *string `json:"reason,omitempty"`

	// Status The status of the dispute.
	Status DisputeStatus `json:"status"`
}

// DisputeEvidence defines model for DisputeEvidence.
type DisputeEvidence struct {
	// Evidence The evidence to add, as key-value pairs. Each key replaces the same key on the dispute. The keys `receipt`, `customer_communication`, `shipping_documentation` and `uncategorized_file` each take the ID of a file you uploaded with purpose `dispute_evidence`, in the same mode. Every other key takes any value.
	Evidence map[string]interface{} `json:"evidence"`
}

// DisputeList defines model for DisputeList.
type DisputeList struct {
	Data []Dispute `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// DisputeStatus The status of the dispute.
type DisputeStatus string

// Error defines model for Error.
type Error struct {
	// Code A short, stable string that identifies the error, such as `resource_missing`. Use it to handle errors in code.
	Code string `json:"code"`

	// DeclineCode Present on `payment_blocked` alone: the kind of risk rule that declined the payment. `velocity` is too many payments in a short time, `blocklist` is a payer on a block list, `geo` is the payer's country, `amount` is the amount, and `generic` is any other rule.
	DeclineCode *ErrorDeclineCode `json:"decline_code,omitempty"`

	// Message A readable explanation of the error. The text can change.
	Message string `json:"message"`

	// Param The parameter the error relates to, if any.
	Param *string `json:"param,omitempty"`

	// RequestId The ID of the request, also sent in the `Request-Id` header.
	RequestId *string `json:"request_id,omitempty"`

	// Type The category of the error.
	Type ErrorType `json:"type"`
}

// ErrorDeclineCode Present on `payment_blocked` alone: the kind of risk rule that declined the payment. `velocity` is too many payments in a short time, `blocklist` is a payer on a block list, `geo` is the payer's country, `amount` is the amount, and `generic` is any other rule.
type ErrorDeclineCode string

// ErrorType The category of the error.
type ErrorType string

// ErrorEnvelope The body of every error response.
type ErrorEnvelope struct {
	Error Error `json:"error"`
}

// Event An event reports a change to one of your objects, such as a successful payment. Events are sent to your webhook endpoints.
type Event struct {
	// Created The time the event occurred.
	Created time.Time `json:"created"`

	// Data The data of the event, such as the object that changed.
	Data map[string]interface{} `json:"data"`

	// Id Unique identifier for the event.
	Id string `json:"id"`

	// Livemode `true` for a production event, `false` for a test event. Test events are sent only to test endpoints.
	Livemode bool        `json:"livemode"`
	Object   interface{} `json:"object"`

	// Type The event type, such as `payment_intent.succeeded`.
	Type string `json:"type"`
}

// EventList defines model for EventList.
type EventList struct {
	Data []Event `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// File defines model for File.
type File struct {
	// CreatedAt The time the file was uploaded.
	CreatedAt time.Time `json:"created_at"`

	// Filename The name the file was uploaded with.
	Filename string `json:"filename"`

	// Id Unique identifier for the file.
	Id     string      `json:"id"`
	Object interface{} `json:"object"`

	// Purpose Why the file was uploaded.
	Purpose FilePurpose `json:"purpose"`

	// Size The size of the file in bytes.
	Size int64 `json:"size"`

	// Type The type of the file, read from its contents.
	Type FileType `json:"type"`
}

// FilePurpose Why the file was uploaded.
type FilePurpose string

// FileType The type of the file, read from its contents.
type FileType string

// FileUpload defines model for FileUpload.
type FileUpload struct {
	// File The file. A PDF, a JPEG or a PNG of at most 5 MB.
	File string `json:"file"`

	// Purpose Why the file was uploaded.
	Purpose FilePurpose `json:"purpose"`
}

// FxLock The exchange rate locked for this payment.
type FxLock struct {
	// LockPoint When the rate is locked: when the payment intent is created, when it is paid, or when the channel settles it. A payment intent with automatic conversion is always `on_create`.
	LockPoint *FxLockLockPoint `json:"lock_point,omitempty"`

	// Locked Whether a rate is locked.
	Locked bool `json:"locked"`

	// LockedAt The time the rate was locked.
	LockedAt *time.Time `json:"locked_at,omitempty"`

	// QuoteCurrency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	QuoteCurrency *Currency `json:"quote_currency,omitempty"`

	// Rate The applied rate, as a decimal string: units of the payment currency for one unit of `quote_currency`, markup included.
	Rate *string `json:"rate,omitempty"`
}

// FxLockLockPoint When the rate is locked: when the payment intent is created, when it is paid, or when the channel settles it. A payment intent with automatic conversion is always `on_create`.
type FxLockLockPoint string

// FxQuote defines model for FxQuote.
type FxQuote struct {
	Estimated *Money `json:"estimated,omitempty"`

	// From A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	From Currency `json:"from"`

	// Indicative Always `true`. A quote locks no rate.
	Indicative interface{} `json:"indicative"`

	// Rate The rate that applies, markup included, as a decimal string.
	Rate string `json:"rate"`

	// To A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	To Currency `json:"to"`
}

// MerchantRequest defines model for MerchantRequest.
type MerchantRequest struct {
	CreatedAt time.Time           `json:"created_at"`
	DecidedAt *time.Time          `json:"decided_at,omitempty"`
	DecidedBy *string             `json:"decided_by,omitempty"`
	ExpiresAt time.Time           `json:"expires_at"`
	Id        string              `json:"id"`
	Kind      MerchantRequestKind `json:"kind"`
	Mode      MerchantRequestMode `json:"mode"`
	Note      *string             `json:"note,omitempty"`
	Object    interface{}         `json:"object"`

	// RequestedBy The dashboard user who asked.
	RequestedBy *string `json:"requested_by,omitempty"`

	// RequestedByKey The withdrawal key that asked.
	RequestedByKey *string `json:"requested_by_key,omitempty"`

	// Result The settlement or the payout destination the approval made.
	Result *string               `json:"result,omitempty"`
	Status MerchantRequestStatus `json:"status"`

	// Summary What the change does, masked: the amount and the destination of a withdrawal, the type, label and last four characters of a destination.
	Summary map[string]string `json:"summary"`
}

// MerchantRequestKind defines model for MerchantRequest.Kind.
type MerchantRequestKind string

// MerchantRequestMode defines model for MerchantRequest.Mode.
type MerchantRequestMode string

// MerchantRequestStatus defines model for MerchantRequestStatus.
type MerchantRequestStatus string

// Money defines model for Money.
type Money struct {
	// Amount The amount in the smallest currency unit, as an integer. For example, `1000` is 10.00 USD.
	Amount int64 `json:"amount"`

	// Currency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	Currency Currency `json:"currency"`
}

// NextAction What the payer must do to complete the payment.
type NextAction struct {
	// Address The deposit address of a `crypto_address` action. Show it as text next to the QR code.
	Address *string `json:"address,omitempty"`

	// QrData The data to show as a QR code. For `crypto_address`, the wallet payment URI.
	QrData *string `json:"qr_data,omitempty"`

	// Type How to present the payment: `redirect` sends the payer to `url`, `qr_code` shows `qr_data` as a QR code, and `crypto_address` shows a deposit address.
	Type NextActionType `json:"type"`

	// Url The URL to send the payer to. For `crypto_address`, the wallet payment URI.
	Url *string `json:"url,omitempty"`
}

// NextActionType How to present the payment: `redirect` sends the payer to `url`, `qr_code` shows `qr_data` as a QR code, and `crypto_address` shows a deposit address.
type NextActionType string

// Payer Details about the payer. Some payment channels require some of these fields, and a channel is not used when a field it requires is missing. The details are sent to the channel for this payment only. They are never stored or returned.
type Payer struct {
	// BankCode The payer's bank, as one of the bank codes the payment method supports.
	BankCode *string `json:"bank_code,omitempty"`

	// Country Two-letter ISO 3166-1 country code. When set, routing keeps to payment methods of this country and to methods tied to no country, such as crypto. The hosted checkout page offers only those.
	Country *string `json:"country,omitempty"`

	// Email The payer's email address.
	Email *string `json:"email,omitempty"`

	// Ip The payer's IP address. Some channels, such as `cashapp`, require it. Send the payer's address, not your server's. The hosted checkout page ignores this field and uses the payer's own address.
	Ip *string `json:"ip,omitempty"`

	// Name The payer's full name.
	Name *string `json:"name,omitempty"`

	// NationalId The payer's national identity number, such as a 13-digit CNIC.
	NationalId *string `json:"national_id,omitempty"`

	// Phone The payer's phone number in E.164 format, such as `+66812345678`.
	Phone *string `json:"phone,omitempty"`
}

// PaymentAttempt An attempt to pay a payment intent through a payment channel. A payment intent has at most one pending attempt. Failed attempts are kept.
type PaymentAttempt struct {
	// Amount The amount in the smallest currency unit.
	Amount int64 `json:"amount"`

	// Created The time the attempt was created.
	Created time.Time `json:"created"`

	// Currency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	Currency Currency `json:"currency"`

	// FailureCode Why the attempt failed, such as `insufficient_funds`.
	FailureCode *string `json:"failure_code,omitempty"`

	// FailureMessage A readable explanation of the failure.
	FailureMessage *string `json:"failure_message,omitempty"`

	// Id Unique identifier for the attempt.
	Id string `json:"id"`

	// NextAction What the payer must do to complete the payment.
	NextAction *NextAction `json:"next_action,omitempty"`
	Object     interface{} `json:"object"`

	// PaymentMethod The payment method code.
	PaymentMethod *string `json:"payment_method,omitempty"`

	// Status The status of the attempt.
	Status PaymentAttemptStatus `json:"status"`

	// SucceededAt The time the attempt succeeded.
	SucceededAt *time.Time `json:"succeeded_at,omitempty"`
}

// PaymentAttemptStatus The status of the attempt.
type PaymentAttemptStatus string

// PaymentAttemptList defines model for PaymentAttemptList.
type PaymentAttemptList struct {
	Data []PaymentAttempt `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// PaymentIntent A payment intent tracks one payment from creation to settlement.
type PaymentIntent struct {
	// Amount The amount to collect, in the smallest currency unit. With automatic conversion, the price you set, and the payer pays `payment_amount`.
	Amount int64 `json:"amount"`

	// AvailableAt The time the funds become available in your balance.
	AvailableAt *time.Time `json:"available_at,omitempty"`

	// CheckoutUrl The URL of the hosted checkout page. It is returned only in the response that creates the payment intent.
	CheckoutUrl *string `json:"checkout_url,omitempty"`

	// Created The time the payment intent was created.
	Created time.Time `json:"created"`

	// Crypto Whether a crypto payment is final. Present only on crypto payment intents. On a layer 2 network, a payment succeeds when the transfer is seen and becomes final when its batch is confirmed on layer 1. Until then the chain can reverse it: the payment intent moves to `failed`, and you receive the events `payment_intent.failed` and `payment_intent.reversed`. When the payment becomes final, you receive `payment_intent.finalized` once. For a high-value order, wait for this event before you deliver.
	Crypto *CryptoFinality `json:"crypto,omitempty"`

	// Currency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	Currency Currency `json:"currency"`

	// ExpiresAt The time the payment intent expires if it is not paid.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// Fee The whole fee for the payment, in the payment currency, once a payment method is chosen. It includes the cost of the payment channel.
	Fee *Money `json:"fee,omitempty"`

	// Fx The exchange rate locked for this payment.
	Fx *FxLock `json:"fx,omitempty"`

	// Id Unique identifier for the payment intent.
	Id string `json:"id"`

	// Invoice The invoice issued with the payment intent. Its hosted page shows the amount due, the due date and a pay button while the payment intent is open, and the outcome after. The page prints to an A4 PDF from the browser.
	Invoice *PaymentIntentInvoice `json:"invoice,omitempty"`

	// MarkupBearer Who pays the conversion markup. Present only with automatic conversion.
	MarkupBearer *PaymentIntentMarkupBearer `json:"markup_bearer,omitempty"`

	// MerchantOrderId Your ID for the order.
	MerchantOrderId string `json:"merchant_order_id"`

	// Metadata Key-value pairs of strings stored on the payment intent.
	Metadata *map[string]string `json:"metadata,omitempty"`

	// Mode `live` for production, `test` for test mode.
	Mode PaymentIntentMode `json:"mode"`

	// Net The amount you receive, in the settlement currency.
	Net *Money `json:"net,omitempty"`

	// NextAction What the payer must do to complete the payment.
	NextAction *NextAction `json:"next_action,omitempty"`
	Object     interface{} `json:"object"`

	// PaymentAmount What the payer pays. Present only with automatic conversion. The fee is charged on this amount.
	PaymentAmount *Money `json:"payment_amount,omitempty"`

	// PaymentMethod The payment method code, once one is chosen.
	PaymentMethod *string `json:"payment_method,omitempty"`

	// PresentmentDetails What the payer is shown and charged. Present once a payment method is chosen.
	PresentmentDetails *PresentmentDetails `json:"presentment_details,omitempty"`

	// Receipt The receipt issued when the payment intent succeeds. Present from `succeeded` on.
	Receipt *PaymentIntentReceipt `json:"receipt,omitempty"`

	// Status The status of the payment intent.
	Status PaymentIntentStatus `json:"status"`
}

// PaymentIntentMarkupBearer Who pays the conversion markup. Present only with automatic conversion.
type PaymentIntentMarkupBearer string

// PaymentIntentMode `live` for production, `test` for test mode.
type PaymentIntentMode string

// PaymentIntentAdvance defines model for PaymentIntentAdvance.
type PaymentIntentAdvance struct {
	// Clock The timer to move. `expiry` moves `expires_at` and `release` moves `available_at`.
	Clock PaymentIntentAdvanceClock `json:"clock"`
}

// PaymentIntentAdvanceClock The timer to move. `expiry` moves `expires_at` and `release` moves `available_at`.
type PaymentIntentAdvanceClock string

// PaymentIntentCreate defines model for PaymentIntentCreate.
type PaymentIntentCreate struct {
	// Amount The amount to collect, in the smallest currency unit. With `payment_currency`, the price of the order in `currency`, which must then be your settlement currency, or the request fails with 400 `unsupported_pricing_currency`.
	Amount int64 `json:"amount"`

	// ChannelPreference The code of a payment channel you prefer.
	ChannelPreference *string `json:"channel_preference,omitempty"`

	// Currency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	Currency Currency `json:"currency"`

	// Customer Your ID for the customer.
	Customer *string `json:"customer,omitempty"`

	// Description A description of the payment.
	Description *string `json:"description,omitempty"`

	// ExpiresIn The number of seconds until the payment intent expires. When omitted, the payment method's default applies.
	ExpiresIn *int `json:"expires_in,omitempty"`

	// MarkupBearer Who pays the conversion markup when `payment_currency` is set. Defaults to `payer`.
	//
	// `payer`: the payer pays the conversion and your fee, so you receive exactly `amount`. The payment amount is `amount` at the applied rate (`fx.rate`), grossed up by your fee rate and fixed fee, and rounded up to the payment currency's smallest unit.
	//
	// `merchant`: the payer pays `amount` converted at the market rate without markup, rounded up. You receive the payment amount less your fee, converted at the applied rate (`fx.rate`) and rounded down.
	MarkupBearer *PaymentIntentCreateMarkupBearer `json:"markup_bearer,omitempty"`

	// MerchantOrderId Your ID for the order. It must be unique across all payment intents in your account.
	MerchantOrderId string `json:"merchant_order_id"`

	// Metadata Key-value pairs of strings to store on the payment intent.
	Metadata *map[string]string `json:"metadata,omitempty"`

	// Payer Details about the payer. Some payment channels require some of these fields, and a channel is not used when a field it requires is missing. The details are sent to the channel for this payment only. They are never stored or returned.
	Payer *Payer `json:"payer,omitempty"`

	// PaymentCurrency The currency the payer pays in, when it differs from `currency`. The amount is converted automatically and returned in `payment_amount`. The rate is locked at creation and holds until the payment intent expires. Requires `payment_method`; a crypto method returns 400 `conversion_unsupported`. When no rate is available, the request fails with 400 `rate_unavailable`.
	PaymentCurrency *Currency `json:"payment_currency,omitempty"`

	// PaymentMethod The payment method code. When set, the order is placed with a payment channel at once. When omitted, the payer chooses a method on the hosted checkout page. For a crypto method such as `tron` or `base`, `currency` is the token, `USDT` or `USDC`, and the default expiry is 30 minutes.
	PaymentMethod *string `json:"payment_method,omitempty"`

	// ReceiptEmail The payer's email address. When the payment succeeds, the receipt is sent to it, in the language of the payment's country, else English. Any value that is not one email address returns 400 `invalid_receipt_email`. It is not returned.
	ReceiptEmail *string `json:"receipt_email,omitempty"`

	// ReturnUrl The page the payer returns to after a successful payment. It must be an absolute `https` URL with a host name and no user info. IP addresses and `localhost` are not accepted, in test mode too. Any other value returns 400 `invalid_return_url`.
	ReturnUrl *string `json:"return_url,omitempty"`
}

// PaymentIntentCreateMarkupBearer Who pays the conversion markup when `payment_currency` is set. Defaults to `payer`.
//
// `payer`: the payer pays the conversion and your fee, so you receive exactly `amount`. The payment amount is `amount` at the applied rate (`fx.rate`), grossed up by your fee rate and fixed fee, and rounded up to the payment currency's smallest unit.
//
// `merchant`: the payer pays `amount` converted at the market rate without markup, rounded up. You receive the payment amount less your fee, converted at the applied rate (`fx.rate`) and rounded down.
type PaymentIntentCreateMarkupBearer string

// PaymentIntentInvoice The invoice issued with the payment intent. Its hosted page shows the amount due, the due date and a pay button while the payment intent is open, and the outcome after. The page prints to an A4 PDF from the browser.
type PaymentIntentInvoice struct {
	// HostedInvoiceUrl The URL of the hosted invoice page. Share it with the payer.
	HostedInvoiceUrl string `json:"hosted_invoice_url"`

	// Number The invoice number, sequential per account and mode, for example `3F9A1C07-0042`. A sandbox number starts with `TEST-`.
	Number string `json:"number"`
}

// PaymentIntentList defines model for PaymentIntentList.
type PaymentIntentList struct {
	Data []PaymentIntent `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// PaymentIntentReceipt The receipt issued when the payment intent succeeds. Present from `succeeded` on.
type PaymentIntentReceipt struct {
	// Number The receipt number, sequential per account and mode, for example `3F9A1C07-R0042`. A sandbox number starts with `TEST-`.
	Number string `json:"number"`

	// ReceiptUrl The URL of the hosted receipt page.
	ReceiptUrl string `json:"receipt_url"`
}

// PaymentIntentSimulate defines model for PaymentIntentSimulate.
type PaymentIntentSimulate struct {
	// Amount The amount the channel reports, in the smallest currency unit. Defaults to the payment intent amount. Send a different value to test a payment whose amount does not match.
	Amount *int64 `json:"amount,omitempty"`

	// Currency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	Currency *Currency `json:"currency,omitempty"`

	// FailureCode The decline reason for a `failed` outcome. It becomes the attempt's `failure_code`. Allowed only with the outcome `failed`.
	FailureCode *SimulatedDecline `json:"failure_code,omitempty"`

	// Outcome The outcome to simulate. `succeeded`, `failed` and `pending` are payment outcomes, and they replace any outcome a test amount set. `dispute_opened` opens a dispute for the full amount of a paid payment intent, and `dispute_won` or `dispute_lost` decides that dispute.
	Outcome PaymentIntentSimulateOutcome `json:"outcome"`
}

// PaymentIntentSimulateOutcome The outcome to simulate. `succeeded`, `failed` and `pending` are payment outcomes, and they replace any outcome a test amount set. `dispute_opened` opens a dispute for the full amount of a paid payment intent, and `dispute_won` or `dispute_lost` decides that dispute.
type PaymentIntentSimulateOutcome string

// PaymentIntentStatus The status of the payment intent.
type PaymentIntentStatus string

// PaymentMethod A payment method you can use, and the currencies it serves.
type PaymentMethod struct {
	// Code The code you send as `payment_method`.
	Code string `json:"code"`

	// Currencies The currencies the method serves, in alphabetical order.
	Currencies []Currency `json:"currencies"`

	// IconUrl An icon for the method, when one is set.
	IconUrl *string `json:"icon_url,omitempty"`

	// Kind The kind of method, such as `qr`, `wallet`, `bank`, `card` or `crypto`.
	Kind string `json:"kind"`

	// Names The name of the method by language tag, such as `en`, `zh-Hant` or `zh-Hans`. `en` is always present.
	Names  map[string]string `json:"names"`
	Object interface{}       `json:"object"`

	// TransactionLimits The charge rules of the method, per currency it serves.
	TransactionLimits map[string]TransactionLimits `json:"transaction_limits"`
}

// PaymentMethodList defines model for PaymentMethodList.
type PaymentMethodList struct {
	Data []PaymentMethod `json:"data"`

	// HasMore Always `false`.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// PresentmentDetails The amount the payer is shown and charged.
type PresentmentDetails struct {
	// PresentmentAmount The amount charged, in the smallest currency unit: `amount` rounded up to the method's `rounding_increment`, and at least its `minimum_amount`.
	PresentmentAmount int64 `json:"presentment_amount"`

	// PresentmentCurrency A currency code in uppercase: an ISO 4217 code such as `THB`, or a token such as `USDT`.
	PresentmentCurrency Currency `json:"presentment_currency"`
}

// Refund defines model for Refund.
type Refund struct {
	// Amount The amount refunded to the payer, in the currency the payer paid in.
	Amount Money `json:"amount"`

	// CreatedAt The time the refund was created.
	CreatedAt time.Time `json:"created_at"`

	// FailureCode Why the refund failed. Present on a `failed` refund alone.
	FailureCode *string `json:"failure_code,omitempty"`

	// Fee The refund fee of the payment channel, in your settlement currency. It is debited from your balance with the refund.
	Fee Money `json:"fee"`

	// Id Unique identifier for the refund.
	Id string `json:"id"`

	// Metadata The key-value pairs stored on the refund.
	Metadata map[string]string `json:"metadata"`
	Object   interface{}       `json:"object"`

	// PaymentIntent The ID of the refunded payment intent.
	PaymentIntent string `json:"payment_intent"`

	// Reason Why you refund the payment.
	Reason *RefundReason `json:"reason,omitempty"`

	// Status The status of the refund. `pending` waits for the payment channel to confirm it, and `succeeded` and `failed` are final.
	Status RefundStatus `json:"status"`
}

// RefundCreate defines model for RefundCreate.
type RefundCreate struct {
	// Amount The amount to refund, in the smallest unit of the currency the payer paid in. When omitted, the whole amount that is left is refunded.
	Amount *int64 `json:"amount,omitempty"`

	// Metadata Key-value pairs of strings to store on the refund.
	Metadata *map[string]string `json:"metadata,omitempty"`

	// PaymentIntent The ID of the payment intent to refund.
	PaymentIntent string `json:"payment_intent"`

	// Reason Why you refund the payment.
	Reason *RefundReason `json:"reason,omitempty"`
}

// RefundList defines model for RefundList.
type RefundList struct {
	Data []Refund `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// RefundReason Why you refund the payment.
type RefundReason string

// RefundStatus The status of the refund. `pending` waits for the payment channel to confirm it, and `succeeded` and `failed` are final.
type RefundStatus string

// Settlement defines model for Settlement.
type Settlement struct {
	// Amount The amount paid to you. It equals the total of the settled payments.
	Amount Money `json:"amount"`

	// Created The time the settlement was created.
	Created time.Time `json:"created"`

	// Destination Where the settlement is paid. Absent when the settlement names no destination.
	Destination *SettlementDestination `json:"destination,omitempty"`

	// Id Unique identifier for the settlement.
	Id string `json:"id"`

	// Items The payment intents this settlement pays. Present only when you retrieve one settlement. The item amounts add up to `amount`.
	Items *[]SettlementItem `json:"items,omitempty"`

	// Mode `live` for production, `test` for test mode.
	Mode   SettlementMode `json:"mode"`
	Object interface{}    `json:"object"`

	// PaidAt The time the settlement was paid. Present only when `status` is `paid`.
	PaidAt *time.Time `json:"paid_at,omitempty"`

	// Payments The number of payment intents in this settlement.
	Payments *int `json:"payments,omitempty"`

	// PeriodEnd The end of the period this settlement covers.
	PeriodEnd time.Time `json:"period_end"`

	// PeriodStart The start of the period this settlement covers.
	PeriodStart time.Time `json:"period_start"`

	// Status The status of the settlement.
	Status SettlementStatus `json:"status"`
}

// SettlementMode `live` for production, `test` for test mode.
type SettlementMode string

// SettlementCreate defines model for SettlementCreate.
type SettlementCreate struct {
	// Amount Minor units to withdraw. Absent pays every released payment. The oldest payments are paid first, and the last one in part.
	Amount *int64 `json:"amount,omitempty"`

	// Currency Required with `amount`.
	Currency *SettlementCreateCurrency `json:"currency,omitempty"`

	// PayoutMethod An address book entry. Absent is the default.
	PayoutMethod *string `json:"payout_method,omitempty"`
}

// SettlementCreateCurrency Required with `amount`.
type SettlementCreateCurrency string

// SettlementDestination Where the settlement is paid. Absent when the settlement names no destination.
type SettlementDestination struct {
	// Label Your own name for the destination
	Label *string `json:"label,omitempty"`

	// Masked The last four characters of the destination, and nothing else. An empty string when the destination has four characters or fewer, or none.
	Masked string `json:"masked"`

	// Type The kind of payout destination.
	Type DestinationType `json:"type"`
}

// SettlementItem defines model for SettlementItem.
type SettlementItem struct {
	// Amount The amount paid for this payment intent.
	Amount Money       `json:"amount"`
	Object interface{} `json:"object"`

	// PaymentIntent The ID of the payment intent this item pays.
	PaymentIntent string `json:"payment_intent"`
}

// SettlementList defines model for SettlementList.
type SettlementList struct {
	Data []Settlement `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// SettlementStatus The status of the settlement.
type SettlementStatus string

// SimulatedDecline The decline reason for a `failed` outcome. It becomes the attempt's `failure_code`. Allowed only with the outcome `failed`.
type SimulatedDecline string

// TransactionLimits How the payer's charge is computed from a payment intent's `amount`, which never changes.
type TransactionLimits struct {
	// MinimumAmount The least the payer is charged, in the smallest currency unit, or null for none.
	MinimumAmount *int64 `json:"minimum_amount"`

	// RoundingIncrement The payer's charge is a multiple of this, in the smallest currency unit; 1 takes any amount.
	RoundingIncrement int64 `json:"rounding_increment"`

	// RoundingMode The charge is rounded up, never down.
	RoundingMode TransactionLimitsRoundingMode `json:"rounding_mode"`
}

// TransactionLimitsRoundingMode The charge is rounded up, never down.
type TransactionLimitsRoundingMode string

// WebhookPublicKey defines model for WebhookPublicKey.
type WebhookPublicKey struct {
	Algorithm interface{} `json:"algorithm"`

	// Id The ID of the key. The `v2_kid` value in the `X-Webhook-Signature` header names it.
	Id     string      `json:"id"`
	Object interface{} `json:"object"`

	// PublicKey The raw 32-byte public key, base64-encoded.
	PublicKey string `json:"public_key"`
}

// WebhookPublicKeyList defines model for WebhookPublicKeyList.
type WebhookPublicKeyList struct {
	Data []WebhookPublicKey `json:"data"`

	// HasMore Whether more objects exist after this page.
	HasMore bool        `json:"has_more"`
	Object  interface{} `json:"object"`
}

// ApiVersion defines model for ApiVersion.
type ApiVersion = string

// DisputeID defines model for DisputeID.
type DisputeID = string

// FileID defines model for FileID.
type FileID = string

// IdempotencyKey defines model for IdempotencyKey.
type IdempotencyKey = string

// Limit defines model for Limit.
type Limit = int

// PaymentIntentID defines model for PaymentIntentID.
type PaymentIntentID = string

// RefundID defines model for RefundID.
type RefundID = string

// SettlementID defines model for SettlementID.
type SettlementID = string

// StartingAfter defines model for StartingAfter.
type StartingAfter = string

// ErrorResponse The body of every error response.
type ErrorResponse = ErrorEnvelope

// BalanceParams defines parameters for Balance.
type BalanceParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListBalanceTransactionsParams defines parameters for ListBalanceTransactions.
type ListBalanceTransactionsParams struct {
	// Limit The maximum number of objects to return, from 1 to 100.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// StartingAfter A pagination cursor. Pass the ID of the last object on the previous page to get the next page.
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListCurrenciesParams defines parameters for ListCurrencies.
type ListCurrenciesParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListDisputesParams defines parameters for ListDisputes.
type ListDisputesParams struct {
	// Limit The maximum number of objects to return, from 1 to 100.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// StartingAfter A pagination cursor. Pass the ID of the last object on the previous page to get the next page.
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// Status Only return disputes with this status.
	Status *DisputeStatus `form:"status,omitempty" json:"status,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// DisputeParams defines parameters for Dispute.
type DisputeParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// CloseDisputeParams defines parameters for CloseDispute.
type CloseDisputeParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// SubmitDisputeEvidenceParams defines parameters for SubmitDisputeEvidence.
type SubmitDisputeEvidenceParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// EventsParams defines parameters for Events.
type EventsParams struct {
	// Limit The maximum number of objects to return, from 1 to 100.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// StartingAfter The ID of the last event on the previous page.
	StartingAfter *string `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// CreateFileParams defines parameters for CreateFile.
type CreateFileParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// FileParams defines parameters for File.
type FileParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// FxQuoteParams defines parameters for FxQuote.
type FxQuoteParams struct {
	// From The currency to convert from.
	From Currency `form:"from" json:"from"`

	// To The currency to convert to. Only `USD` is supported. Any other value returns 400 `unsupported_pair`.
	To Currency `form:"to" json:"to"`

	// PaymentMethod A payment method code. The rate can differ by payment method.
	PaymentMethod *string `form:"payment_method,omitempty" json:"payment_method,omitempty"`

	// Amount The amount to convert, in the smallest unit of the `from` currency.
	Amount *int64 `form:"amount,omitempty" json:"amount,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListPaymentIntentsParams defines parameters for ListPaymentIntents.
type ListPaymentIntentsParams struct {
	// Limit The maximum number of objects to return, from 1 to 100.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// StartingAfter A pagination cursor. Pass the ID of the last object on the previous page to get the next page.
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// Status Only return payment intents with this status.
	Status *PaymentIntentStatus `form:"status,omitempty" json:"status,omitempty"`

	// MerchantOrderId Only return the payment intent with this order ID.
	MerchantOrderId *string `form:"merchant_order_id,omitempty" json:"merchant_order_id,omitempty"`

	// CreatedGte Only return payment intents created at or after this time, in RFC 3339 format.
	CreatedGte *time.Time `form:"created[gte],omitempty" json:"created[gte],omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// CreatePaymentIntentParams defines parameters for CreatePaymentIntent.
type CreatePaymentIntentParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// PaymentIntentParams defines parameters for PaymentIntent.
type PaymentIntentParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListPaymentAttemptsParams defines parameters for ListPaymentAttempts.
type ListPaymentAttemptsParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// CreatePaymentAttemptJSONBody defines parameters for CreatePaymentAttempt.
type CreatePaymentAttemptJSONBody struct {
	// PaymentMethod The payment method code. Required if the payment intent has no order yet. Without it, the request returns 400 `payment_method_required`.
	PaymentMethod string `json:"payment_method"`
}

// CreatePaymentAttemptParams defines parameters for CreatePaymentAttempt.
type CreatePaymentAttemptParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// CancelPaymentIntentParams defines parameters for CancelPaymentIntent.
type CancelPaymentIntentParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListPaymentMethodsParams defines parameters for ListPaymentMethods.
type ListPaymentMethodsParams struct {
	// Currency List only the methods that serve this currency.
	Currency *Currency `form:"currency,omitempty" json:"currency,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListRefundsParams defines parameters for ListRefunds.
type ListRefundsParams struct {
	// Limit The maximum number of objects to return, from 1 to 100.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// StartingAfter A pagination cursor. Pass the ID of the last object on the previous page to get the next page.
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// PaymentIntent Only return refunds of this payment intent.
	PaymentIntent *string `form:"payment_intent,omitempty" json:"payment_intent,omitempty"`

	// Status Only return refunds with this status.
	Status *RefundStatus `form:"status,omitempty" json:"status,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// CreateRefundParams defines parameters for CreateRefund.
type CreateRefundParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// RefundParams defines parameters for Refund.
type RefundParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// ListSettlementsParams defines parameters for ListSettlements.
type ListSettlementsParams struct {
	// Limit The maximum number of objects to return, from 1 to 100.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// StartingAfter A pagination cursor. Pass the ID of the last object on the previous page to get the next page.
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// Status Only return settlements with this status. Any other value returns 400 `invalid_status`.
	Status *SettlementStatus `form:"status,omitempty" json:"status,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// CreateSettlementParams defines parameters for CreateSettlement.
type CreateSettlementParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// SettlementParams defines parameters for Settlement.
type SettlementParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// StatementParams defines parameters for Statement.
type StatementParams struct {
	// Month The calendar month in UTC, as YYYY-MM. The current month is served to the time of the request and marked provisional. A month that has not begun is 400 invalid_month.
	Month string `form:"month" json:"month"`

	// Format csv for a file Excel opens, or html for a printable A4 page. Print the page to save a PDF. Any other value is 400 invalid_format.
	Format StatementParamsFormat `form:"format" json:"format"`

	// Lang The language of the labels. Defaults to en.
	Lang *StatementParamsLang `form:"lang,omitempty" json:"lang,omitempty"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// StatementParamsFormat defines parameters for Statement.
type StatementParamsFormat string

// StatementParamsLang defines parameters for Statement.
type StatementParamsLang string

// AdvancePaymentIntentParams defines parameters for AdvancePaymentIntent.
type AdvancePaymentIntentParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// SimulatePaymentIntentParams defines parameters for SimulatePaymentIntent.
type SimulatePaymentIntentParams struct {
	// IdempotencyKey A unique key you generate for this request, up to 255 characters. A retry with the same key and body returns the original response. Without it, the request returns 400 `idempotency_key_required`.
	IdempotencyKey IdempotencyKey `json:"Idempotency-Key"`

	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// WebhookPublicKeysParams defines parameters for WebhookPublicKeys.
type WebhookPublicKeysParams struct {
	// ApiVersion The API version for this request, as a date in `YYYY-MM-DD` format. When omitted, your account's version applies, which is set to the latest version on your first request. An unknown version returns 400 `invalid_api_version`.
	ApiVersion *ApiVersion `json:"Api-Version,omitempty"`
}

// SubmitDisputeEvidenceJSONRequestBody defines body for SubmitDisputeEvidence for application/json ContentType.
type SubmitDisputeEvidenceJSONRequestBody = DisputeEvidence

// CreateFileMultipartRequestBody defines body for CreateFile for multipart/form-data ContentType.
type CreateFileMultipartRequestBody = FileUpload

// CreatePaymentIntentJSONRequestBody defines body for CreatePaymentIntent for application/json ContentType.
type CreatePaymentIntentJSONRequestBody = PaymentIntentCreate

// CreatePaymentAttemptJSONRequestBody defines body for CreatePaymentAttempt for application/json ContentType.
type CreatePaymentAttemptJSONRequestBody CreatePaymentAttemptJSONBody

// CreateRefundJSONRequestBody defines body for CreateRefund for application/json ContentType.
type CreateRefundJSONRequestBody = RefundCreate

// CreateSettlementJSONRequestBody defines body for CreateSettlement for application/json ContentType.
type CreateSettlementJSONRequestBody = SettlementCreate

// AdvancePaymentIntentJSONRequestBody defines body for AdvancePaymentIntent for application/json ContentType.
type AdvancePaymentIntentJSONRequestBody = PaymentIntentAdvance

// SimulatePaymentIntentJSONRequestBody defines body for SimulatePaymentIntent for application/json ContentType.
type SimulatePaymentIntentJSONRequestBody = PaymentIntentSimulate
