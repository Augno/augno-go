// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List and manage shipping terms.
//
// OperationShippingTermService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationShippingTermService] method instead.
type OperationShippingTermService struct {
	options []option.RequestOption
}

// NewOperationShippingTermService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationShippingTermService(opts ...option.RequestOption) (r OperationShippingTermService) {
	r = OperationShippingTermService{}
	r.options = opts
	return
}

// Creates an account-owned shipping term.
//
// This endpoint requires the permission: `shipping_terms:create`.
func (r *OperationShippingTermService) New(ctx context.Context, params OperationShippingTermNewParams, opts ...option.RequestOption) (res *ShippingTerm, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/shipping-terms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a shipping term by ID.
//
// This endpoint requires the permission: `shipping_terms:read`.
func (r *OperationShippingTermService) Get(ctx context.Context, id string, query OperationShippingTermGetParams, opts ...option.RequestOption) (res *ShippingTerm, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/shipping-terms/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an account-owned shipping term.
//
// System-provided default shipping terms cannot be updated.
//
// This endpoint requires the permission: `shipping_terms:update`.
func (r *OperationShippingTermService) Update(ctx context.Context, id string, params OperationShippingTermUpdateParams, opts ...option.RequestOption) (res *ShippingTerm, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/shipping-terms/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of shipping terms for the account, including default
// system shipping terms.
//
// This endpoint requires the permission: `shipping_terms:read`.
func (r *OperationShippingTermService) List(ctx context.Context, query OperationShippingTermListParams, opts ...option.RequestOption) (res *ListShippingTerm, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/shipping-terms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an account-owned shipping term.
//
// System-provided default shipping terms cannot be deleted.
//
// This endpoint requires the permission: `shipping_terms:delete`.
func (r *OperationShippingTermService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationShippingTermDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/shipping-terms/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a shipping term.
//
// The properties Name, Type are required.
type CreateShippingTermRequestParam struct {
	// Human-readable name for the shipping term, used to identify it when assigning
	// shipping terms to customers and orders.
	Name string `json:"name" api:"required"`
	// Freight pricing model applied by this shipping term.
	//
	//   - `free_freight`: no shipping cost to the buyer.
	//   - `flat_rate_freight`: a fixed shipping cost regardless of order details (see
	//     `flat_rate`).
	//   - `carrier_rate_freight`: shipping cost is determined by the carrier's quoted
	//     rate.
	//
	// Any of "free_freight", "flat_rate_freight", "carrier_rate_freight".
	Type CreateShippingTermRequestType `json:"type,omitzero" api:"required"`
	// A value with an associated unit, used in create and update requests.
	FlatRate QuantityInputParam `json:"flat_rate,omitzero"`
	// IDs of service levels that ship for free under this term (typically once
	// `minimum_order_value` is met).
	FreeShippingServiceLevelIDs []string `json:"free_shipping_service_level_ids,omitzero"`
	// A value with an associated unit, used in create and update requests.
	MinimumOrderValue QuantityInputParam `json:"minimum_order_value,omitzero"`
	paramObj
}

func (r CreateShippingTermRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateShippingTermRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateShippingTermRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Freight pricing model applied by this shipping term.
//
//   - `free_freight`: no shipping cost to the buyer.
//   - `flat_rate_freight`: a fixed shipping cost regardless of order details (see
//     `flat_rate`).
//   - `carrier_rate_freight`: shipping cost is determined by the carrier's quoted
//     rate.
type CreateShippingTermRequestType string

const (
	CreateShippingTermRequestTypeFreeFreight        CreateShippingTermRequestType = "free_freight"
	CreateShippingTermRequestTypeFlatRateFreight    CreateShippingTermRequestType = "flat_rate_freight"
	CreateShippingTermRequestTypeCarrierRateFreight CreateShippingTermRequestType = "carrier_rate_freight"
)

// List represents a paginated list of resources.
type ListShippingTerm struct {
	// Resources in this page.
	Data []ShippingTerm `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListShippingTermObject `json:"object" api:"required"`
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
func (r ListShippingTerm) RawJSON() string { return r.JSON.raw }
func (r *ListShippingTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListShippingTermObject string

const (
	ListShippingTermObjectList ListShippingTermObject = "list"
)

// Request to partially update a shipping term.
//
// All fields are optional and absent fields are left unchanged. Send an explicit
// JSON `null` for `flat_rate`, `minimum_order_value`, or
// `free_shipping_service_level_ids` to clear the existing value.
type UpdateShippingTermRequestParam struct {
	// Human-readable name for the shipping term, used to identify it when assigning
	// shipping terms to customers and orders.
	Name param.Opt[string] `json:"name,omitzero"`
	// IDs of service levels that ship for free under this term (typically once
	// `minimum_order_value` is met).
	//
	// Replaces the existing list. Send `null` to clear.
	FreeShippingServiceLevelIDs []string `json:"free_shipping_service_level_ids,omitzero"`
	// A value with an associated unit, used in create and update requests.
	FlatRate QuantityInputParam `json:"flat_rate,omitzero"`
	// A value with an associated unit, used in create and update requests.
	MinimumOrderValue QuantityInputParam `json:"minimum_order_value,omitzero"`
	// Freight pricing model applied by this shipping term.
	//
	//   - `free_freight`: no shipping cost to the buyer.
	//   - `flat_rate_freight`: a fixed shipping cost regardless of order details (see
	//     `flat_rate`).
	//   - `carrier_rate_freight`: shipping cost is determined by the carrier's quoted
	//     rate.
	//
	// Any of "free_freight", "flat_rate_freight", "carrier_rate_freight".
	Type UpdateShippingTermRequestType `json:"type,omitzero"`
	paramObj
}

func (r UpdateShippingTermRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateShippingTermRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateShippingTermRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Freight pricing model applied by this shipping term.
//
//   - `free_freight`: no shipping cost to the buyer.
//   - `flat_rate_freight`: a fixed shipping cost regardless of order details (see
//     `flat_rate`).
//   - `carrier_rate_freight`: shipping cost is determined by the carrier's quoted
//     rate.
type UpdateShippingTermRequestType string

const (
	UpdateShippingTermRequestTypeFreeFreight        UpdateShippingTermRequestType = "free_freight"
	UpdateShippingTermRequestTypeFlatRateFreight    UpdateShippingTermRequestType = "flat_rate_freight"
	UpdateShippingTermRequestTypeCarrierRateFreight UpdateShippingTermRequestType = "carrier_rate_freight"
)

type OperationShippingTermDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationShippingTermDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationShippingTermDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationShippingTermNewParams struct {
	// Request to create a shipping term.
	CreateShippingTermRequest CreateShippingTermRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "flat_rate.unit", "minimum_order_value.unit",
	// "free_shipping_service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationShippingTermNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateShippingTermRequest)
}
func (r *OperationShippingTermNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationShippingTermNewParams]'s query parameters as
// `url.Values`.
func (r OperationShippingTermNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationShippingTermGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "flat_rate.unit", "minimum_order_value.unit",
	// "free_shipping_service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationShippingTermGetParams]'s query parameters as
// `url.Values`.
func (r OperationShippingTermGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationShippingTermUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "flat_rate.unit", "minimum_order_value.unit",
	// "free_shipping_service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a shipping term.
	//
	// All fields are optional and absent fields are left unchanged. Send an explicit
	// JSON `null` for `flat_rate`, `minimum_order_value`, or
	// `free_shipping_service_level_ids` to clear the existing value.
	UpdateShippingTermRequest UpdateShippingTermRequestParam
	paramObj
}

func (r OperationShippingTermUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateShippingTermRequest)
}
func (r *OperationShippingTermUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationShippingTermUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationShippingTermUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationShippingTermListParams struct {
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
	// Any of "owner", "owner.account", "flat_rate.unit", "minimum_order_value.unit",
	// "free_shipping_service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationShippingTermListParams]'s query parameters as
// `url.Values`.
func (r OperationShippingTermListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
