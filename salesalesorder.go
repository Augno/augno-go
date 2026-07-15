// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// SaleSalesOrderService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderService] method instead.
type SaleSalesOrderService struct {
	options []option.RequestOption
	// List, view, create, update, and delete sales orders.
	Actions SaleSalesOrderActionService
	// List, view, create, update, and delete sales orders.
	Lines SaleSalesOrderLineService
}

// NewSaleSalesOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleSalesOrderService(opts ...option.RequestOption) (r SaleSalesOrderService) {
	r = SaleSalesOrderService{}
	r.options = opts
	r.Actions = NewSaleSalesOrderActionService(opts...)
	r.Lines = NewSaleSalesOrderLineService(opts...)
	return
}

// Creates a sales order in `estimate` status.
//
// The order number is assigned automatically, and a sales rep is auto-assigned
// when none is provided. A shipping line is always added to the order, plus a
// discount line when an order discount is supplied.
//
// This endpoint requires the permission: `sales_orders:create`.
func (r *SaleSalesOrderService) New(ctx context.Context, params SaleSalesOrderNewParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a sales order by ID.
//
// This endpoint requires the permissions: `customers:read`, `suppliers:read`,
// `sales_orders:read`.
func (r *SaleSalesOrderService) Get(ctx context.Context, id string, query SaleSalesOrderGetParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a sales order.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderService) Update(ctx context.Context, id string, params SaleSalesOrderUpdateParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of sales orders for the current account.
//
// This endpoint requires the permissions: `sales_orders:read`, `customers:read`,
// `suppliers:read`.
func (r *SaleSalesOrderService) List(ctx context.Context, query SaleSalesOrderListParams, opts ...option.RequestOption) (res *ListSalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a sales order and all its related records.
//
// Fulfilled orders cannot be deleted.
//
// This endpoint requires the permission: `sales_orders:delete`.
func (r *SaleSalesOrderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleSalesOrderDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Creates a hosted payment checkout session for a sales order.
//
// Requires an active Stripe integration on the account. The checkout is built from
// the order's lines, and the checkout link is emailed to the provided address.
// Fails with a conflict if the order already has a payment.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderService) Checkout(ctx context.Context, id string, body SaleSalesOrderCheckoutParams, opts ...option.RequestOption) (res *CheckoutSalesOrderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/checkout", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Calculates the unit price for each line without creating an order.
//
// Use this to display prices to users as they build an order. Prices are computed
// server-side from the product's list price, contracted account prices, and
// applicable discounts — the same logic used when an order is created. Internal
// price overrides are not accepted here; the calculated price is always returned.
//
// This endpoint requires the permission: `sales_orders:read`.
func (r *SaleSalesOrderService) PriceQuote(ctx context.Context, body SaleSalesOrderPriceQuoteParams, opts ...option.RequestOption) (res *QuoteSalesOrderPricesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/price-quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of sales order statuses.
func (r *SaleSalesOrderService) GetStatuses(ctx context.Context, query SaleSalesOrderGetStatusesParams, opts ...option.RequestOption) (res *ListSalesOrderStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/statuses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request to create a checkout session for a sales order.
//
// The property Email is required.
type CheckoutSalesOrderRequestParam struct {
	// Email address to send the checkout link to.
	//
	// Also set as the customer email on the payment provider's checkout session.
	Email string `json:"email" api:"required"`
	paramObj
}

func (r CheckoutSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Checkout session result.
type CheckoutSalesOrderResponse struct {
	// URL of the hosted payment page where the customer completes the checkout.
	CheckoutURL string `json:"checkout_url" api:"required"`
	// Resource type identifier.
	//
	// Any of "checkout_sales_order".
	Object CheckoutSalesOrderResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckoutURL respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckoutSalesOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckoutSalesOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CheckoutSalesOrderResponseObject string

const (
	CheckoutSalesOrderResponseObjectCheckoutSalesOrder CheckoutSalesOrderResponseObject = "checkout_sales_order"
)

// Line item input for a create sales order request.
//
// The item, unit cost, and (unless an internal user supplies a `unit_price`
// override) the unit price are resolved server-side from the product. The quantity
// unit must belong to the product's unit group.
//
// The properties ProductID, Quantity are required.
type CreateSalesOrderLineInputParam struct {
	// ID of the product being ordered.
	ProductID string `json:"product_id" api:"required"`
	// A value with an associated unit, used in create and update requests.
	Quantity QuantityInputParam `json:"quantity,omitzero" api:"required"`
	// Description recorded on the line.
	//
	// Defaults to the product's description when omitted.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// SKU recorded on the line.
	//
	// Defaults to the product's SKU when omitted.
	ProductSKU param.Opt[string] `json:"product_sku,omitzero"`
	// A rate value with its numerator and denominator units, used in create and update
	// requests.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r CreateSalesOrderLineInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSalesOrderLineInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSalesOrderLineInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to create a sales order.
//
// The properties BillToAddressID, BuyerAccountID, Lines, PriorityCode,
// ShipToAddressID are required.
type CreateSalesOrderRequestParam struct {
	// Bill-to address ID.
	//
	// Must reference an existing address on the order's owner or buyer account.
	BillToAddressID string `json:"bill_to_address_id" api:"required"`
	// ID of the customer account the order is for.
	BuyerAccountID string `json:"buyer_account_id" api:"required"`
	// Order lines to create.
	Lines []CreateSalesOrderLineInputParam `json:"lines,omitzero" api:"required"`
	// Fulfillment priority used to rank the order on the shop floor.
	PriorityCode string `json:"priority_code" api:"required"`
	// Ship-to address ID.
	//
	// Must reference an existing address on the order's owner or buyer account.
	ShipToAddressID string `json:"ship_to_address_id" api:"required"`
	// Carrier billing account number.
	CarrierBillingAccountNumber param.Opt[string] `json:"carrier_billing_account_number,omitzero"`
	// Carrier ID.
	CarrierID param.Opt[string] `json:"carrier_id,omitzero"`
	// The customer's own purchase order number, for cross-referencing.
	//
	// Must be unique among your orders for this customer.
	CustomerPurchaseOrderNumber param.Opt[string] `json:"customer_purchase_order_number,omitzero"`
	// Order note.
	Note param.Opt[string] `json:"note,omitzero"`
	// Order discount ID.
	//
	// When supplied, a discount line is added to the order automatically.
	OrderDiscountID param.Opt[string] `json:"order_discount_id,omitzero"`
	// Payment term ID.
	PaymentTermID param.Opt[string] `json:"payment_term_id,omitzero"`
	// Promised delivery date.
	PromisedAt param.Opt[time.Time] `json:"promised_at,omitzero" format:"date-time"`
	// Sales rep ID.
	//
	// When omitted, a rep is assigned automatically: the customer's default sales rep
	// first, then the sales territory matching the ship-to postal code, then the
	// ship-to state.
	SalesRepID param.Opt[string] `json:"sales_rep_id,omitzero"`
	// Service level ID.
	ServiceLevelID param.Opt[string] `json:"service_level_id,omitzero"`
	// Shipping term ID.
	ShippingTermID param.Opt[string] `json:"shipping_term_id,omitzero"`
	// Account users who should receive order acknowledgement emails.
	AcknowledgementEmailContacts []SalesOrderEmailContactInputParam `json:"acknowledgement_email_contacts,omitzero"`
	// Who is billed for freight.
	//
	//   - `sender`: the sender pays for shipping.
	//   - `third_party`: a third party pays for shipping, using the carrier billing
	//     account number.
	//
	// Any of "sender", "third_party".
	CarrierBillingType CreateSalesOrderRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// Account users who should receive invoice emails.
	InvoiceEmailContacts []SalesOrderEmailContactInputParam `json:"invoice_email_contacts,omitzero"`
	paramObj
}

func (r CreateSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who is billed for freight.
//
//   - `sender`: the sender pays for shipping.
//   - `third_party`: a third party pays for shipping, using the carrier billing
//     account number.
type CreateSalesOrderRequestCarrierBillingType string

const (
	CreateSalesOrderRequestCarrierBillingTypeSender     CreateSalesOrderRequestCarrierBillingType = "sender"
	CreateSalesOrderRequestCarrierBillingTypeThirdParty CreateSalesOrderRequestCarrierBillingType = "third_party"
)

// CreatedBy describes who created a resource and their relationship to the account
// that owns it.
//
// It is resolved from the resource's create audit event.
type CreatedBy struct {
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// Resource type identifier.
	//
	// Any of "created_by".
	Object CreatedByObject `json:"object" api:"required"`
	// The creator's relationship to the account that owns the resource.
	//
	// - `internal`: created by a user of the owning account.
	// - `customer`: created by a customer of the owning account.
	// - `system`: created automatically with no human actor (e.g. an EDI import).
	//
	// Any of "internal", "customer", "system".
	Relation CreatedByRelation `json:"relation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actor       respjson.Field
		Object      respjson.Field
		Relation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedBy) RawJSON() string { return r.JSON.raw }
func (r *CreatedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CreatedByObject string

const (
	CreatedByObjectCreatedBy CreatedByObject = "created_by"
)

// The creator's relationship to the account that owns the resource.
//
// - `internal`: created by a user of the owning account.
// - `customer`: created by a customer of the owning account.
// - `system`: created automatically with no human actor (e.g. an EDI import).
type CreatedByRelation string

const (
	CreatedByRelationInternal CreatedByRelation = "internal"
	CreatedByRelationCustomer CreatedByRelation = "customer"
	CreatedByRelationSystem   CreatedByRelation = "system"
)

// Freight describes the carrier selection and freight billing for a record.
//
// It is a generic, reusable sub-resource shared by anything that carries shipping
// configuration — for example a sales order's chosen freight, or a customer's
// default freight preferences.
type Freight struct {
	// Carrier account number to bill, used when `billing_type` is `third_party`.
	BillingAccountNumber string `json:"billing_account_number" api:"required"`
	// Which party the carrier bills for the shipment.
	//
	// - `sender`: the shipper (your account) is billed.
	// - `third_party`: a third party is billed via `billing_account_number`.
	//
	// Any of "sender", "third_party".
	BillingType FreightBillingType `json:"billing_type" api:"required"`
	// A shipping carrier configured for fulfilling orders.
	//
	// Carriers with a Shippo-supported `code` (`fedex`, `ups`, `usps`) are connected
	// through Shippo for live rating and label purchase; other carriers represent
	// self-managed shipping methods such as will call or local delivery.
	Carrier Carrier `json:"carrier" api:"required"`
	// Resource type identifier.
	//
	// Any of "freight".
	Object FreightObject `json:"object" api:"required"`
	// How freight is arranged and billed for the record.
	//
	// Populated where a freight policy applies, such as a customer's default
	// preferences.
	//
	// - `free_freight`: no shipping cost to the buyer.
	// - `billed_freight`: freight is billed to the buyer.
	//
	// Any of "free_freight", "billed_freight".
	Policy FreightPolicy `json:"policy" api:"required"`
	// Shipping service level for a carrier.
	ServiceLevel ServiceLevel `json:"service_level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAccountNumber respjson.Field
		BillingType          respjson.Field
		Carrier              respjson.Field
		Object               respjson.Field
		Policy               respjson.Field
		ServiceLevel         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Freight) RawJSON() string { return r.JSON.raw }
func (r *Freight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which party the carrier bills for the shipment.
//
// - `sender`: the shipper (your account) is billed.
// - `third_party`: a third party is billed via `billing_account_number`.
type FreightBillingType string

const (
	FreightBillingTypeSender     FreightBillingType = "sender"
	FreightBillingTypeThirdParty FreightBillingType = "third_party"
)

// Resource type identifier.
type FreightObject string

const (
	FreightObjectFreight FreightObject = "freight"
)

// How freight is arranged and billed for the record.
//
// Populated where a freight policy applies, such as a customer's default
// preferences.
//
// - `free_freight`: no shipping cost to the buyer.
// - `billed_freight`: freight is billed to the buyer.
type FreightPolicy string

const (
	FreightPolicyFreeFreight   FreightPolicy = "free_freight"
	FreightPolicyBilledFreight FreightPolicy = "billed_freight"
)

// List represents a paginated list of resources.
type ListQuotedSalesOrderLine struct {
	// Resources in this page.
	Data []QuotedSalesOrderLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListQuotedSalesOrderLineObject `json:"object" api:"required"`
	// PageInfo contains URL-based pagination metadata.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListQuotedSalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *ListQuotedSalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListQuotedSalesOrderLineObject string

const (
	ListQuotedSalesOrderLineObjectList ListQuotedSalesOrderLineObject = "list"
)

// List represents a paginated list of resources.
type ListRecord struct {
	// Resources in this page.
	Data []Record `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListRecordObject `json:"object" api:"required"`
	// PageInfo contains URL-based pagination metadata.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListRecord) RawJSON() string { return r.JSON.raw }
func (r *ListRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListRecordObject string

const (
	ListRecordObjectList ListRecordObject = "list"
)

// List represents a paginated list of resources.
type ListSalesOrder struct {
	// Resources in this page.
	Data []SalesOrder `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSalesOrderObject `json:"object" api:"required"`
	// PageInfo contains URL-based pagination metadata.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSalesOrder) RawJSON() string { return r.JSON.raw }
func (r *ListSalesOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSalesOrderObject string

const (
	ListSalesOrderObjectList ListSalesOrderObject = "list"
)

// List represents a paginated list of resources.
type ListSalesOrderLine struct {
	// Resources in this page.
	Data []SalesOrderLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSalesOrderLineObject `json:"object" api:"required"`
	// PageInfo contains URL-based pagination metadata.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *ListSalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSalesOrderLineObject string

const (
	ListSalesOrderLineObjectList ListSalesOrderLineObject = "list"
)

// List represents a paginated list of resources.
type ListSalesOrderStatus struct {
	// Resources in this page.
	Data []SalesOrderStatus `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSalesOrderStatusObject `json:"object" api:"required"`
	// PageInfo contains URL-based pagination metadata.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSalesOrderStatus) RawJSON() string { return r.JSON.raw }
func (r *ListSalesOrderStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSalesOrderStatusObject string

const (
	ListSalesOrderStatusObjectList ListSalesOrderStatusObject = "list"
)

// OrderContact groups a sales order's email recipients by notification purpose.
type OrderContact struct {
	// Email addresses that receive order acknowledgements for this order.
	Acknowledgement []string `json:"acknowledgement" api:"required"`
	// Email addresses that receive invoices for this order.
	Invoice []string `json:"invoice" api:"required"`
	// Resource type identifier.
	//
	// Any of "order_contact".
	Object OrderContactObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Acknowledgement respjson.Field
		Invoice         respjson.Field
		Object          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderContact) RawJSON() string { return r.JSON.raw }
func (r *OrderContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type OrderContactObject string

const (
	OrderContactObjectOrderContact OrderContactObject = "order_contact"
)

// A discount code that can be applied to a sales order.
//
// An order discount reduces the order total by either a percentage or a fixed
// amount, depending on `discount_type`.
type OrderDiscount struct {
	// Order discount ID.
	ID string `json:"id" api:"required"`
	// Fixed amount off as a decimal string.
	//
	// Applies when `discount_type` is `amount`; otherwise `0`.
	Amount string `json:"amount" api:"required" format:"decimal"`
	// The code entered to apply this discount to an order.
	//
	// Must be unique within the account.
	Code string `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// How the discount is calculated, determining whether `percentage` or `amount` is
	// used.
	//
	// - `percentage`: the discount is a percent off, taken from `percentage`.
	// - `amount`: the discount is a fixed amount off, taken from `amount`.
	//
	// Any of "percentage", "amount".
	DiscountType OrderDiscountDiscountType `json:"discount_type" api:"required"`
	// Display name of the discount.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "order_discount".
	Object OrderDiscountObject `json:"object" api:"required"`
	// Number of orders currently using this discount.
	OrderCount int64 `json:"order_count" api:"required"`
	// Percent off as a decimal string (e.g. `10` for 10%).
	//
	// Applies when `discount_type` is `percentage`; otherwise `0`.
	Percentage string `json:"percentage" api:"required" format:"decimal"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		Code         respjson.Field
		CreatedAt    respjson.Field
		DiscountType respjson.Field
		Name         respjson.Field
		Object       respjson.Field
		OrderCount   respjson.Field
		Percentage   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderDiscount) RawJSON() string { return r.JSON.raw }
func (r *OrderDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the discount is calculated, determining whether `percentage` or `amount` is
// used.
//
// - `percentage`: the discount is a percent off, taken from `percentage`.
// - `amount`: the discount is a fixed amount off, taken from `amount`.
type OrderDiscountDiscountType string

const (
	OrderDiscountDiscountTypePercentage OrderDiscountDiscountType = "percentage"
	OrderDiscountDiscountTypeAmount     OrderDiscountDiscountType = "amount"
)

// Resource type identifier.
type OrderDiscountObject string

const (
	OrderDiscountObjectOrderDiscount OrderDiscountObject = "order_discount"
)

// A line to price in a quote request.
//
// The properties ProductID, Quantity are required.
type QuoteSalesOrderLineInputParam struct {
	// ID of the product to price.
	ProductID string `json:"product_id" api:"required"`
	// A value with an associated unit, used in create and update requests.
	Quantity QuantityInputParam `json:"quantity,omitzero" api:"required"`
	paramObj
}

func (r QuoteSalesOrderLineInputParam) MarshalJSON() (data []byte, err error) {
	type shadow QuoteSalesOrderLineInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuoteSalesOrderLineInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to quote sales-order line prices without creating an order.
//
// The properties BuyerAccountID, Lines are required.
type QuoteSalesOrderPricesRequestParam struct {
	// ID of the customer account the prices are for.
	BuyerAccountID string `json:"buyer_account_id" api:"required"`
	// Lines to price.
	Lines []QuoteSalesOrderLineInputParam `json:"lines,omitzero" api:"required"`
	paramObj
}

func (r QuoteSalesOrderPricesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow QuoteSalesOrderPricesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuoteSalesOrderPricesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quoted unit prices for the requested lines, in request order.
type QuoteSalesOrderPricesResponse struct {
	// List represents a paginated list of resources.
	Lines ListQuotedSalesOrderLine `json:"lines" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_price_quote".
	Object QuoteSalesOrderPricesResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lines       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteSalesOrderPricesResponse) RawJSON() string { return r.JSON.raw }
func (r *QuoteSalesOrderPricesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuoteSalesOrderPricesResponseObject string

const (
	QuoteSalesOrderPricesResponseObjectSalesOrderPriceQuote QuoteSalesOrderPricesResponseObject = "sales_order_price_quote"
)

// One priced line in a quote response.
type QuotedSalesOrderLine struct {
	// Resource type identifier.
	//
	// Any of "sales_order_price_quote_line".
	Object QuotedSalesOrderLineObject `json:"object" api:"required"`
	// Product pairs an inventory item with how it is sold: its product type, optional
	// product line, and customer portal visibility.
	Product Product `json:"product" api:"required"`
	// A per-unit rate on a sales-order quote.
	//
	// A lightweight, unpersisted variant of a rate: it carries no ID or timestamps
	// because a quote is computed on demand and never stored.
	UnitPrice SalesOrderQuoteRate `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		Product     respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotedSalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *QuotedSalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuotedSalesOrderLineObject string

const (
	QuotedSalesOrderLineObjectSalesOrderPriceQuoteLine QuotedSalesOrderLineObject = "sales_order_price_quote_line"
)

// Record is a lightweight reference to a business record — a sales order, purchase
// order, pick, shipment, production run, invoice, etc.
//
// Like Actor and Entity, it carries just enough to identify and label the
// referenced record without embedding its full resource. The optional status and
// metadata fields hold type-specific detail that varies by the kind of record
// referenced.
type Record struct {
	// Unique identifier for the record.
	ID string `json:"id" api:"required"`
	// Type-specific metadata.
	//
	// The set of keys varies by record type.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Human-readable record number, when the record has one.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "record".
	Object RecordObject `json:"object" api:"required"`
	// Type-specific status code, when applicable.
	Status string `json:"status" api:"required"`
	// The kind of business record referenced.
	//
	// Determines how to resolve the record and which `status` and `metadata` keys may
	// appear.
	//
	// - `sales_order`: a customer order.
	// - `purchase_order`: an order placed with a supplier.
	// - `receiving_order`: an inbound order being received into inventory.
	// - `pick`: a warehouse pick task.
	// - `shipment`: an outbound shipment.
	// - `delivery`: a delivery of one or more shipments to a destination.
	// - `production_run`: a manufacturing production run.
	// - `invoice`: a customer invoice.
	// - `transaction`: a payment or financial transaction.
	// - `settlement`: a settlement reconciling transactions against invoices.
	//
	// Any of "sales_order", "purchase_order", "receiving_order", "pick", "shipment",
	// "delivery", "production_run", "invoice", "transaction", "settlement".
	Type RecordType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Metadata    respjson.Field
		Number      respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Record) RawJSON() string { return r.JSON.raw }
func (r *Record) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type RecordObject string

const (
	RecordObjectRecord RecordObject = "record"
)

// The kind of business record referenced.
//
// Determines how to resolve the record and which `status` and `metadata` keys may
// appear.
//
// - `sales_order`: a customer order.
// - `purchase_order`: an order placed with a supplier.
// - `receiving_order`: an inbound order being received into inventory.
// - `pick`: a warehouse pick task.
// - `shipment`: an outbound shipment.
// - `delivery`: a delivery of one or more shipments to a destination.
// - `production_run`: a manufacturing production run.
// - `invoice`: a customer invoice.
// - `transaction`: a payment or financial transaction.
// - `settlement`: a settlement reconciling transactions against invoices.
type RecordType string

const (
	RecordTypeSalesOrder     RecordType = "sales_order"
	RecordTypePurchaseOrder  RecordType = "purchase_order"
	RecordTypeReceivingOrder RecordType = "receiving_order"
	RecordTypePick           RecordType = "pick"
	RecordTypeShipment       RecordType = "shipment"
	RecordTypeDelivery       RecordType = "delivery"
	RecordTypeProductionRun  RecordType = "production_run"
	RecordTypeInvoice        RecordType = "invoice"
	RecordTypeTransaction    RecordType = "transaction"
	RecordTypeSettlement     RecordType = "settlement"
)

// Full sales order resource.
type SalesOrder struct {
	// Sales order ID.
	ID string `json:"id" api:"required"`
	// Whether an order acknowledgment has been sent to the customer.
	//
	// Any of "not_sent", "sent".
	AcknowledgmentStatus SalesOrderAcknowledgmentStatus `json:"acknowledgment_status" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	BillToAddress Address `json:"bill_to_address" api:"required"`
	// When the order was fulfilled and closed.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// OrderContact groups a sales order's email recipients by notification purpose.
	Contacts OrderContact `json:"contacts" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// CreatedBy describes who created a resource and their relationship to the account
	// that owns it.
	//
	// It is resolved from the resource's create audit event.
	CreatedBy CreatedBy `json:"created_by" api:"required"`
	// A business you sell to, with its contact details, default fulfillment settings,
	// and order policies.
	Customer Customer `json:"customer" api:"required"`
	// The customer's own purchase order number, for cross-referencing.
	//
	// Unique among this customer's orders.
	CustomerPurchaseOrderNumber string `json:"customer_purchase_order_number" api:"required"`
	// When this estimate expires, if an expiration was set.
	ExpiredAt time.Time `json:"expired_at" api:"required" format:"date-time"`
	// When the first shipment against this order went out.
	FirstShipAt time.Time `json:"first_ship_at" api:"required" format:"date-time"`
	// Freight describes the carrier selection and freight billing for a record.
	//
	// It is a generic, reusable sub-resource shared by anything that carries shipping
	// configuration — for example a sales order's chosen freight, or a customer's
	// default freight preferences.
	Freight Freight `json:"freight" api:"required"`
	// When the order was issued (moved out of `estimate`).
	IssuedAt time.Time `json:"issued_at" api:"required" format:"date-time"`
	// Number of order lines on this order, returned even when the `lines` list itself
	// is not expanded.
	LineCount int64 `json:"line_count" api:"required"`
	// List represents a paginated list of resources.
	Lines ListSalesOrderLine `json:"lines" api:"required"`
	// Order note.
	Note string `json:"note" api:"required"`
	// Human-readable order number, e.g. `SO-001`.
	//
	// Assigned automatically when the order is created; unique within your account.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order".
	Object SalesOrderObject `json:"object" api:"required"`
	// A discount code that can be applied to a sales order.
	//
	// An order discount reduces the order total by either a percentage or a fixed
	// amount, depending on `discount_type`.
	OrderDiscount OrderDiscount `json:"order_discount" api:"required"`
	// Stripe payment intent IDs recorded against this order.
	PaymentIntentIDs []string `json:"payment_intent_ids" api:"required"`
	// Payment state of the order, derived from settlement allocations, invoices, and
	// Stripe payments.
	//
	// Any of "unpaid", "partially_paid", "paid".
	PaymentStatus SalesOrderPaymentStatus `json:"payment_status" api:"required"`
	// A payment term describing when payment is due (e.g. `Net 30`), assignable to
	// customers, sales orders, purchase orders, and invoices.
	PaymentTerm PaymentTerm `json:"payment_term" api:"required"`
	// Fulfillment priority, used to rank orders on the shop floor.
	//
	// Any of "low", "normal", "high".
	Priority SalesOrderPriority `json:"priority" api:"required"`
	// Date promised to the customer for delivery, if one was committed.
	PromisedAt time.Time `json:"promised_at" api:"required" format:"date-time"`
	// SalesOrderRelated groups the records related to a sales order.
	//
	// The members are individually expandable (e.g. include[]=related.pick). The group
	// is null unless at least one of its members is expanded.
	Related SalesOrderRelated `json:"related" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	SalesRep Actor `json:"sales_rep" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	ShipToAddress Address `json:"ship_to_address" api:"required"`
	// A shipping term defining how freight charges are calculated for an order.
	ShippingTerm ShippingTerm `json:"shipping_term" api:"required"`
	// Order lifecycle status.
	//
	//   - `estimate`: a draft quote that has not yet been committed; not counted as a
	//     real order.
	//   - `issued`: the order has been issued and is being fulfilled.
	//   - `fulfilled`: the order has been completed and closed.
	//
	// Status changes are made through the issue, unissue, close, and reopen action
	// endpoints rather than by updating this field.
	//
	// Any of "estimate", "issued", "fulfilled".
	Status SalesOrderStatus `json:"status" api:"required"`
	// SalesOrderTotals holds the derived monetary totals for a sales order or one of
	// its lines, following the lifecycle ordered -> picked -> packed -> invoiced. Each
	// downstream stage carries both its monetary amount and its completion progress
	// against the ordered baseline.
	Totals SalesOrderTotals `json:"totals" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AcknowledgmentStatus        respjson.Field
		BillToAddress               respjson.Field
		CompletedAt                 respjson.Field
		Contacts                    respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		Customer                    respjson.Field
		CustomerPurchaseOrderNumber respjson.Field
		ExpiredAt                   respjson.Field
		FirstShipAt                 respjson.Field
		Freight                     respjson.Field
		IssuedAt                    respjson.Field
		LineCount                   respjson.Field
		Lines                       respjson.Field
		Note                        respjson.Field
		Number                      respjson.Field
		Object                      respjson.Field
		OrderDiscount               respjson.Field
		PaymentIntentIDs            respjson.Field
		PaymentStatus               respjson.Field
		PaymentTerm                 respjson.Field
		Priority                    respjson.Field
		PromisedAt                  respjson.Field
		Related                     respjson.Field
		SalesRep                    respjson.Field
		ShipToAddress               respjson.Field
		ShippingTerm                respjson.Field
		Status                      respjson.Field
		Totals                      respjson.Field
		UpdatedAt                   respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrder) RawJSON() string { return r.JSON.raw }
func (r *SalesOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether an order acknowledgment has been sent to the customer.
type SalesOrderAcknowledgmentStatus string

const (
	SalesOrderAcknowledgmentStatusNotSent SalesOrderAcknowledgmentStatus = "not_sent"
	SalesOrderAcknowledgmentStatusSent    SalesOrderAcknowledgmentStatus = "sent"
)

// Resource type identifier.
type SalesOrderObject string

const (
	SalesOrderObjectSalesOrder SalesOrderObject = "sales_order"
)

// Payment state of the order, derived from settlement allocations, invoices, and
// Stripe payments.
type SalesOrderPaymentStatus string

const (
	SalesOrderPaymentStatusUnpaid        SalesOrderPaymentStatus = "unpaid"
	SalesOrderPaymentStatusPartiallyPaid SalesOrderPaymentStatus = "partially_paid"
	SalesOrderPaymentStatusPaid          SalesOrderPaymentStatus = "paid"
)

// Fulfillment priority, used to rank orders on the shop floor.
type SalesOrderPriority string

const (
	SalesOrderPriorityLow    SalesOrderPriority = "low"
	SalesOrderPriorityNormal SalesOrderPriority = "normal"
	SalesOrderPriorityHigh   SalesOrderPriority = "high"
)

// Order lifecycle status.
//
//   - `estimate`: a draft quote that has not yet been committed; not counted as a
//     real order.
//   - `issued`: the order has been issued and is being fulfilled.
//   - `fulfilled`: the order has been completed and closed.
//
// Status changes are made through the issue, unissue, close, and reopen action
// endpoints rather than by updating this field.
type SalesOrderStatus string

const (
	SalesOrderStatusEstimate  SalesOrderStatus = "estimate"
	SalesOrderStatusIssued    SalesOrderStatus = "issued"
	SalesOrderStatusFulfilled SalesOrderStatus = "fulfilled"
)

// SalesOrderEmailContactInput represents an account user subscribed to a
// sales-order email notification type.
//
// The property AccountUserID is required.
type SalesOrderEmailContactInputParam struct {
	// Account user ID to receive the notification.
	AccountUserID string `json:"account_user_id" api:"required"`
	paramObj
}

func (r SalesOrderEmailContactInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SalesOrderEmailContactInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SalesOrderEmailContactInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full sales order line resource.
type SalesOrderLine struct {
	// Sales order line ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Position of the line on the order.
	//
	// Assigned automatically in sequence, starting at `1`.
	LineItemNumber int64 `json:"line_item_number" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_line".
	Object SalesOrderLineObject `json:"object" api:"required"`
	// Product pairs an inventory item with how it is sold: its product type, optional
	// product line, and customer portal visibility.
	Product Product `json:"product" api:"required"`
	// Product description.
	ProductDescription string `json:"product_description" api:"required"`
	// Product SKU.
	ProductSKU string `json:"product_sku" api:"required"`
	// Value with an associated unit.
	QuantityOrdered Quantity `json:"quantity_ordered" api:"required"`
	// SalesOrderTotals holds the derived monetary totals for a sales order or one of
	// its lines, following the lifecycle ordered -> picked -> packed -> invoiced. Each
	// downstream stage carries both its monetary amount and its completion progress
	// against the ordered baseline.
	Totals SalesOrderTotals `json:"totals" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitCost Rate `json:"unit_cost" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitPrice Rate `json:"unit_price" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		LineItemNumber     respjson.Field
		Object             respjson.Field
		Product            respjson.Field
		ProductDescription respjson.Field
		ProductSKU         respjson.Field
		QuantityOrdered    respjson.Field
		Totals             respjson.Field
		UnitCost           respjson.Field
		UnitPrice          respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderLineObject string

const (
	SalesOrderLineObjectSalesOrderLine SalesOrderLineObject = "sales_order_line"
)

// A per-unit rate on a sales-order quote.
//
// A lightweight, unpersisted variant of a rate: it carries no ID or timestamps
// because a quote is computed on demand and never stored.
type SalesOrderQuoteRate struct {
	// Unit of measurement used for conversions and product quantities.
	DenominatorUnit Unit `json:"denominator_unit" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	NumeratorUnit Unit `json:"numerator_unit" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_quote_rate".
	Object SalesOrderQuoteRateObject `json:"object" api:"required"`
	// Decimal value of the rate, expressed as the amount of the numerator unit per one
	// denominator unit.
	Value string `json:"value" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DenominatorUnit respjson.Field
		NumeratorUnit   respjson.Field
		Object          respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderQuoteRate) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderQuoteRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderQuoteRateObject string

const (
	SalesOrderQuoteRateObjectSalesOrderQuoteRate SalesOrderQuoteRateObject = "sales_order_quote_rate"
)

// SalesOrderRelated groups the records related to a sales order.
//
// The members are individually expandable (e.g. include[]=related.pick). The group
// is null unless at least one of its members is expanded.
type SalesOrderRelated struct {
	// List represents a paginated list of resources.
	Invoices ListRecord `json:"invoices" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_related".
	Object SalesOrderRelatedObject `json:"object" api:"required"`
	// Record is a lightweight reference to a business record — a sales order, purchase
	// order, pick, shipment, production run, invoice, etc.
	//
	// Like Actor and Entity, it carries just enough to identify and label the
	// referenced record without embedding its full resource. The optional status and
	// metadata fields hold type-specific detail that varies by the kind of record
	// referenced.
	Pick Record `json:"pick" api:"required"`
	// Record is a lightweight reference to a business record — a sales order, purchase
	// order, pick, shipment, production run, invoice, etc.
	//
	// Like Actor and Entity, it carries just enough to identify and label the
	// referenced record without embedding its full resource. The optional status and
	// metadata fields hold type-specific detail that varies by the kind of record
	// referenced.
	ProductionRun Record `json:"production_run" api:"required"`
	// List represents a paginated list of resources.
	Shipments ListRecord `json:"shipments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoices      respjson.Field
		Object        respjson.Field
		Pick          respjson.Field
		ProductionRun respjson.Field
		Shipments     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderRelated) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderRelated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderRelatedObject string

const (
	SalesOrderRelatedObjectSalesOrderRelated SalesOrderRelatedObject = "sales_order_related"
)

// SalesOrderStageTotal pairs a fulfillment stage's monetary amount with its
// completion progress.
type SalesOrderStageTotal struct {
	// Amount for this stage as a decimal string (unit price x quantity at this stage).
	Amount string `json:"amount" api:"required" format:"decimal"`
	// Progress to completion for this stage, as a fraction between 0 and 1: quantity
	// at this stage divided by quantity ordered. `0` when nothing has reached this
	// stage yet.
	Completion float64 `json:"completion" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_stage_total".
	Object SalesOrderStageTotalObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Completion  respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderStageTotal) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderStageTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderStageTotalObject string

const (
	SalesOrderStageTotalObjectSalesOrderStageTotal SalesOrderStageTotalObject = "sales_order_stage_total"
)

// SalesOrderTotals holds the derived monetary totals for a sales order or one of
// its lines, following the lifecycle ordered -> picked -> packed -> invoiced. Each
// downstream stage carries both its monetary amount and its completion progress
// against the ordered baseline.
type SalesOrderTotals struct {
	// SalesOrderStageTotal pairs a fulfillment stage's monetary amount with its
	// completion progress.
	Invoiced SalesOrderStageTotal `json:"invoiced" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_totals".
	Object SalesOrderTotalsObject `json:"object" api:"required"`
	// Total ordered amount as a decimal string (unit price x quantity ordered). This
	// is the baseline the stage completions are measured against.
	Ordered string `json:"ordered" api:"required" format:"decimal"`
	// SalesOrderStageTotal pairs a fulfillment stage's monetary amount with its
	// completion progress.
	Packed SalesOrderStageTotal `json:"packed" api:"required"`
	// SalesOrderStageTotal pairs a fulfillment stage's monetary amount with its
	// completion progress.
	Picked SalesOrderStageTotal `json:"picked" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoiced    respjson.Field
		Object      respjson.Field
		Ordered     respjson.Field
		Packed      respjson.Field
		Picked      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderTotals) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderTotalsObject string

const (
	SalesOrderTotalsObjectSalesOrderTotals SalesOrderTotalsObject = "sales_order_totals"
)

// Request to update a sales order.
type UpdateSalesOrderRequestParam struct {
	// Carrier billing account number. Send `null` to clear.
	CarrierBillingAccountNumber param.Opt[string] `json:"carrier_billing_account_number,omitzero"`
	// Customer's purchase order number. Send `null` to clear.
	CustomerPurchaseOrderNumber param.Opt[string] `json:"customer_purchase_order_number,omitzero"`
	// Order note. Send `null` to clear.
	Note param.Opt[string] `json:"note,omitzero"`
	// Order discount ID. Send `null` to clear.
	OrderDiscountID param.Opt[string] `json:"order_discount_id,omitzero"`
	// Promised delivery date. Send `null` to clear.
	PromisedAt param.Opt[time.Time] `json:"promised_at,omitzero" format:"date-time"`
	// Sales rep ID. Send `null` to clear.
	SalesRepID param.Opt[string] `json:"sales_rep_id,omitzero"`
	// Service level ID. Send `null` to clear.
	ServiceLevelID param.Opt[string] `json:"service_level_id,omitzero"`
	// Billing address ID.
	//
	// Re-points the order to an existing address. To change an address's contents, use
	// the update-address endpoint.
	BillingAddressID param.Opt[string] `json:"billing_address_id,omitzero"`
	// Carrier ID.
	CarrierID param.Opt[string] `json:"carrier_id,omitzero"`
	// Customer ID.
	CustomerID param.Opt[string] `json:"customer_id,omitzero"`
	// Payment term ID.
	PaymentTermID param.Opt[string] `json:"payment_term_id,omitzero"`
	// New fulfillment priority for the order.
	PriorityCode param.Opt[string] `json:"priority_code,omitzero"`
	// Shipping address ID.
	//
	// Re-points the order to an existing address. To change an address's contents, use
	// the update-address endpoint.
	ShippingAddressID param.Opt[string] `json:"shipping_address_id,omitzero"`
	// Shipping term ID.
	ShippingTermID param.Opt[string] `json:"shipping_term_id,omitzero"`
	// Who is billed for freight. Send `null` to clear.
	//
	//   - `sender`: the sender pays for shipping.
	//   - `third_party`: a third party pays for shipping, using the carrier billing
	//     account number.
	//
	// Any of "sender", "third_party".
	CarrierBillingType UpdateSalesOrderRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// Replaces the acknowledgement email contacts on the order.
	//
	// An empty list clears all contacts; omitting the field leaves existing contacts
	// untouched.
	AcknowledgementEmailContacts []SalesOrderEmailContactInputParam `json:"acknowledgement_email_contacts,omitzero"`
	// Acknowledgment status of the order.
	//
	// Set to `sent` to mark the acknowledgement as sent without emailing the customer,
	// or `not_sent` to reset it.
	//
	// Any of "not_sent", "sent".
	AcknowledgmentStatus UpdateSalesOrderRequestAcknowledgmentStatus `json:"acknowledgment_status,omitzero"`
	// Replaces the invoice email contacts on the order.
	//
	// An empty list clears all contacts; omitting the field leaves existing contacts
	// untouched.
	InvoiceEmailContacts []SalesOrderEmailContactInputParam `json:"invoice_email_contacts,omitzero"`
	paramObj
}

func (r UpdateSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Acknowledgment status of the order.
//
// Set to `sent` to mark the acknowledgement as sent without emailing the customer,
// or `not_sent` to reset it.
type UpdateSalesOrderRequestAcknowledgmentStatus string

const (
	UpdateSalesOrderRequestAcknowledgmentStatusNotSent UpdateSalesOrderRequestAcknowledgmentStatus = "not_sent"
	UpdateSalesOrderRequestAcknowledgmentStatusSent    UpdateSalesOrderRequestAcknowledgmentStatus = "sent"
)

// Who is billed for freight. Send `null` to clear.
//
//   - `sender`: the sender pays for shipping.
//   - `third_party`: a third party pays for shipping, using the carrier billing
//     account number.
type UpdateSalesOrderRequestCarrierBillingType string

const (
	UpdateSalesOrderRequestCarrierBillingTypeSender     UpdateSalesOrderRequestCarrierBillingType = "sender"
	UpdateSalesOrderRequestCarrierBillingTypeThirdParty UpdateSalesOrderRequestCarrierBillingType = "third_party"
)

type SaleSalesOrderDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleSalesOrderDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleSalesOrderDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderNewParams struct {
	// Request to create a sales order.
	CreateSalesOrderRequest CreateSalesOrderRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "bill_to_address", "ship_to_address", "freight",
	// "payment_term", "shipping_term", "order_discount", "totals", "contacts",
	// "related.pick", "related.production_run", "related.shipments",
	// "related.invoices", "lines", "lines.product", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleSalesOrderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSalesOrderRequest)
}
func (r *SaleSalesOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderNewParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "created_by", "bill_to_address",
	// "ship_to_address", "freight", "payment_term", "shipping_term", "order_discount",
	// "totals", "contacts", "related.pick", "related.production_run",
	// "related.shipments", "related.invoices", "lines", "lines.product",
	// "lines.product.item", "lines.product.product_line", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderGetParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "bill_to_address", "ship_to_address", "freight",
	// "payment_term", "shipping_term", "order_discount", "totals", "contacts",
	// "related.pick", "related.production_run", "related.shipments",
	// "related.invoices", "lines", "lines.product", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a sales order.
	UpdateSalesOrderRequest UpdateSalesOrderRequestParam
	paramObj
}

func (r SaleSalesOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateSalesOrderRequest)
}
func (r *SaleSalesOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Latest order creation date to include, in `YYYY-MM-DD` format (inclusive).
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Earliest order creation date to include, in `YYYY-MM-DD` format (inclusive).
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// Filter by customer group IDs.
	CustomerGroupIDs []string `query:"customer_group_ids,omitzero" json:"-"`
	// Filter by customer IDs.
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "created_by", "bill_to_address",
	// "ship_to_address", "freight", "payment_term", "shipping_term", "order_discount",
	// "totals", "contacts", "related.pick", "related.production_run",
	// "related.shipments", "related.invoices", "lines", "lines.product",
	// "lines.product.item", "lines.product.product_line", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by item IDs.
	ItemIDs []string `query:"item_ids,omitzero" json:"-"`
	// Filter by product line IDs.
	ProductLineIDs []string `query:"product_line_ids,omitzero" json:"-"`
	// Filter by sales rep IDs.
	SalesRepIDs []string `query:"sales_rep_ids,omitzero" json:"-"`
	// Filter by status codes.
	StatusCodes []string `query:"status_codes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderListParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderCheckoutParams struct {
	// Request to create a checkout session for a sales order.
	CheckoutSalesOrderRequest CheckoutSalesOrderRequestParam
	paramObj
}

func (r SaleSalesOrderCheckoutParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CheckoutSalesOrderRequest)
}
func (r *SaleSalesOrderCheckoutParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderPriceQuoteParams struct {
	// Request to quote sales-order line prices without creating an order.
	QuoteSalesOrderPricesRequest QuoteSalesOrderPricesRequestParam
	paramObj
}

func (r SaleSalesOrderPriceQuoteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.QuoteSalesOrderPricesRequest)
}
func (r *SaleSalesOrderPriceQuoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderGetStatusesParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderGetStatusesParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderGetStatusesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
