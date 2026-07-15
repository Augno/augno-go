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

// List, view, create, update, and delete sales orders.
//
// SaleSalesOrderActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderActionService] method instead.
type SaleSalesOrderActionService struct {
	options []option.RequestOption
}

// NewSaleSalesOrderActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleSalesOrderActionService(opts ...option.RequestOption) (r SaleSalesOrderActionService) {
	r = SaleSalesOrderActionService{}
	r.options = opts
	return
}

// Deletes multiple sales orders in a single atomic operation.
//
// Fulfilled orders cannot be deleted; if any requested order fails this check, no
// orders are deleted.
//
// This endpoint requires the permission: `sales_orders:delete`.
func (r *SaleSalesOrderActionService) BulkDelete(ctx context.Context, body SaleSalesOrderActionBulkDeleteParams, opts ...option.RequestOption) (res *SaleSalesOrderActionBulkDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/actions/bulk-delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Closes a sales order, transitioning it from `issued` to `fulfilled`.
//
// Sets the order's completion timestamp and marks its pick as finished.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Close(ctx context.Context, id string, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/close", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Creates a production run from a sales order.
//
// Creates a batch for each of the order's item-backed lines, reserves the material
// inventory required to produce them, and links the run to the order. An order can
// have at most one production run.
//
// This endpoint requires the permission: `production_runs:create`.
func (r *SaleSalesOrderActionService) NewProductionRun(ctx context.Context, id string, body SaleSalesOrderActionNewProductionRunParams, opts ...option.RequestOption) (res *ProductionRun, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/create-production-run", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Issues a sales order, transitioning it from `estimate` to `issued`.
//
// Issuing commits the order for fulfillment: a pick is created for the order's
// sale lines and inventory is reserved for each line tied to an inventory item.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Issue(ctx context.Context, id string, body SaleSalesOrderActionIssueParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/issue", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Reopens a sales order, transitioning it from `fulfilled` back to `issued`.
//
// Clears the order's completion timestamp and marks its pick as unfinished.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Open(ctx context.Context, id string, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/open", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Re-estimates the freight (shipping) charge for an order using the latest carrier
// rates.
//
// Computes what the order's freight charge would be from its current ship-to
// address, carrier, service level, and line items — applying the same
// freight-exemption, flat-rate, and live carrier-rate logic used when the order is
// created. The order is not modified: the returned amount is a quote to review,
// and callers apply it by updating the order's shipping line. Use this to refresh
// freight after changing the address or line items, or at any time to re-price
// against current rates.
//
// This endpoint requires the permission: `sales_orders:read`.
func (r *SaleSalesOrderActionService) QuoteFreight(ctx context.Context, id string, opts ...option.RequestOption) (res *QuoteSalesOrderFreightResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/quote-freight", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Unissues a sales order, transitioning it from `issued` back to `estimate`.
//
// Deletes the order's pick and releases any inventory reserved when the order was
// issued.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Unissue(ctx context.Context, id string, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/unissue", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Request to bulk delete sales orders.
//
// The property SalesOrderIDs is required.
type BulkDeleteSalesOrdersRequestParam struct {
	// IDs of the sales orders to delete.
	SalesOrderIDs []string `json:"sales_order_ids,omitzero" api:"required"`
	paramObj
}

func (r BulkDeleteSalesOrdersRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkDeleteSalesOrdersRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkDeleteSalesOrdersRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to issue a sales order.
//
// The property NotifyCustomer is required.
type IssueSalesOrderRequestParam struct {
	// Whether to notify the customer.
	//
	// When `true`, the order acknowledgement email is sent to the contacts configured
	// on the order and the order's `acknowledgment_status` is set to `sent`.
	NotifyCustomer bool `json:"notify_customer" api:"required"`
	paramObj
}

func (r IssueSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IssueSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IssueSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Production run resource.
type ProductionRun struct {
	// Production run ID.
	ID string `json:"id" api:"required"`
	// Number of batches currently recorded against this run.
	BatchCount int64 `json:"batch_count" api:"required"`
	// Time the run was marked complete.
	//
	// Set automatically once every batch in the run has been scanned or deleted, and
	// unset while the run is still in progress. Once set, the run can no longer be
	// updated and new batches can no longer be added.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Production run number, unique per account.
	//
	// Assigned automatically at creation as the next sequential number for the
	// account; can be changed via update.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_run".
	Object ProductionRunObject `json:"object" api:"required"`
	// A user's membership in an account, carrying the account-specific status, role,
	// and department.
	//
	// Profile fields (name, email, username, image URL) live on the expandable `user`
	// sub-resource, which is shared across every account the user belongs to.
	ResponsibleUser AccountUser `json:"responsible_user" api:"required"`
	// Time the run started production.
	//
	// Set automatically when the first batch in the run is scanned, and unset until
	// then.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BatchCount      respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		Number          respjson.Field
		Object          respjson.Field
		ResponsibleUser respjson.Field
		StartedAt       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionRun) RawJSON() string { return r.JSON.raw }
func (r *ProductionRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionRunObject string

const (
	ProductionRunObjectProductionRun ProductionRunObject = "production_run"
)

// The freshly estimated freight charge for a sales order.
type QuoteSalesOrderFreightResponse struct {
	// Resource type identifier.
	//
	// Any of "sales_order_freight_quote".
	Object QuoteSalesOrderFreightResponseObject `json:"object" api:"required"`
	// A per-unit rate on a sales-order quote.
	//
	// A lightweight, unpersisted variant of a rate: it carries no ID or timestamps
	// because a quote is computed on demand and never stored.
	UnitPrice SalesOrderQuoteRate `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteSalesOrderFreightResponse) RawJSON() string { return r.JSON.raw }
func (r *QuoteSalesOrderFreightResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuoteSalesOrderFreightResponseObject string

const (
	QuoteSalesOrderFreightResponseObjectSalesOrderFreightQuote QuoteSalesOrderFreightResponseObject = "sales_order_freight_quote"
)

type SaleSalesOrderActionBulkDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleSalesOrderActionBulkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleSalesOrderActionBulkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderActionBulkDeleteParams struct {
	// Request to bulk delete sales orders.
	BulkDeleteSalesOrdersRequest BulkDeleteSalesOrdersRequestParam
	paramObj
}

func (r SaleSalesOrderActionBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkDeleteSalesOrdersRequest)
}
func (r *SaleSalesOrderActionBulkDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderActionNewProductionRunParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "responsible_user", "responsible_user.user".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderActionNewProductionRunParams]'s query
// parameters as `url.Values`.
func (r SaleSalesOrderActionNewProductionRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderActionIssueParams struct {
	// Request to issue a sales order.
	IssueSalesOrderRequest IssueSalesOrderRequestParam
	paramObj
}

func (r SaleSalesOrderActionIssueParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IssueSalesOrderRequest)
}
func (r *SaleSalesOrderActionIssueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
