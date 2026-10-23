# Changelog

All notable changes to this module are documented in this file. The format is based on
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this module follows
[Semantic Versioning](https://semver.org/spec/v2.0.0.html). Before 1.0.0, a minor version can
change an exported signature; each such change is listed under **Changed** with how to migrate.

## [0.7.0] - 2026-10-23

### Added

- `ListIps` returns the addresses webhook deliveries leave from, for a firewall allow-list. The
  webhook signature stays the control; the list is empty until published.

## [0.6.0] - 2026-10-23

### Added

- Refunds: `CreateRefund`, `Refund` and `ListRefunds`.
- Dispute evidence files: `CreateFile` uploads a PDF, JPEG or PNG of at most 5 MB as
  `multipart/form-data` and signs the multipart bytes as sent; `File` reads the file object.
- `CloseDispute` accepts a dispute. The dispute moves to `lost`.
- `EvidenceReceipt`, `EvidenceCustomerCommunication`, `EvidenceShippingDocumentation` and
  `EvidenceUncategorizedFile`: the `DisputeEvidence` keys that take a file ID.
- `ListBalanceTransactions` iterates over every movement of the balance.
- `Statement` returns the CSV or HTML statement of one month as raw bytes.
- Withdrawals: `CreateSettlement` takes a `SettlementCreate` body (`Amount`, `Currency`,
  `PayoutMethod`) and returns a `MerchantRequest` when an approval policy holds the withdrawal.
- Withdrawal keys (`wk_test_...`, `wk_live_...`) work as `Client.Key`. In production,
  `CreateSettlement` needs one.
- `APIError.DeclineCode`, set on `payment_blocked`: `velocity`, `blocklist`, `geo`, `amount` or
  `generic`.
- The event types `EventRefundCreated`, `EventRefundSucceeded` and `EventRefundFailed`.

### Changed

- **Breaking:** `CreateSettlement` has a new signature.

  ```go
  // 0.5.0
  func (c *Client) CreateSettlement(ctx context.Context, p *CreateSettlementParams) (*Settlement, error)
  // 0.6.0
  func (c *Client) CreateSettlement(ctx context.Context, in *SettlementCreate, p *CreateSettlementParams) (*Settlement, *MerchantRequest, error)
  ```

  To keep the old behaviour, pay every released payment to the default destination, pass `nil`
  as `in`. When the account has no approval policy, the `*MerchantRequest` is always `nil`:

  ```go
  st, _, err := c.CreateSettlement(ctx, nil, nil)
  ```

  When the policy holds the withdrawal for a second approver, the `*Settlement` is `nil` and the
  `*MerchantRequest` is the pending request.
- `APIVersion` is `2026-10-23`, and every request sends it as `Api-Version` unless
  `Client.APIVersion` or the request names another version. At `2026-10-23` the evidence keys
  that take a file must name a file uploaded with `FilePurposeDisputeEvidence`, or the request
  fails with `invalid_evidence_file`. To keep sending free text under those keys, set
  `Client.APIVersion` to `2026-09-01`.

## [0.5.0] - 2026-10-22

### Added

- Invoices and receipts.

## [0.4.0] - 2026-10-22

### Added

- Transaction limits and presentment details.

## [0.3.1] - 2026-10-21

### Removed

- The unused channel field of `CreatePaymentAttempt`.

## [0.3.0] - 2026-10-21

### Added

- `ListSettlements` and `Settlement`.
- `ListPaymentMethods` and `ListCurrencies`.
- The `no_pending_attempt` error of `SimulatePaymentIntent` is documented.

## [0.2.0] - 2026-10-21

### Added

- Automatic currency conversion.

## [0.1.0] - 2026-10-20

- Initial release.

[0.6.0]: https://github.com/lemakpay/client-go/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/lemakpay/client-go/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/lemakpay/client-go/compare/v0.3.1...v0.4.0
[0.3.1]: https://github.com/lemakpay/client-go/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/lemakpay/client-go/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/lemakpay/client-go/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/lemakpay/client-go/releases/tag/v0.1.0
