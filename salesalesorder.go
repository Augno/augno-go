// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List sales order statuses.
//
// SaleSalesOrderService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderService] method instead.
type SaleSalesOrderService struct {
	options []option.RequestOption
}

// NewSaleSalesOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleSalesOrderService(opts ...option.RequestOption) (r SaleSalesOrderService) {
	r = SaleSalesOrderService{}
	r.options = opts
	return
}

// Returns a paginated list of sales order statuses.
func (r *SaleSalesOrderService) GetStatuses(ctx context.Context, query SaleSalesOrderGetStatusesParams, opts ...option.RequestOption) (res *ListSalesOrderStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/statuses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

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

// Sales order status lookup value.
type SalesOrderStatus struct {
	// Sales order status ID.
	ID string `json:"id" api:"required"`
	// Machine-readable status code.
	//
	// Any of "estimate", "issued", "fulfilled".
	Code SalesOrderStatusCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_status".
	Object SalesOrderStatusObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderStatus) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable status code.
type SalesOrderStatusCode string

const (
	SalesOrderStatusCodeEstimate  SalesOrderStatusCode = "estimate"
	SalesOrderStatusCodeIssued    SalesOrderStatusCode = "issued"
	SalesOrderStatusCodeFulfilled SalesOrderStatusCode = "fulfilled"
)

// Resource type identifier.
type SalesOrderStatusObject string

const (
	SalesOrderStatusObjectSalesOrderStatus SalesOrderStatusObject = "sales_order_status"
)

type SaleSalesOrderGetStatusesParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
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
