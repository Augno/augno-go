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

// List and manage carriers and their Shippo integrations.
//
// OperationCarrierService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationCarrierService] method instead.
type OperationCarrierService struct {
	options []option.RequestOption
	// List and manage service levels (shipping service levels).
	ServiceLevels OperationCarrierServiceLevelService
}

// NewOperationCarrierService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationCarrierService(opts ...option.RequestOption) (r OperationCarrierService) {
	r = OperationCarrierService{}
	r.options = opts
	r.ServiceLevels = NewOperationCarrierServiceLevelService(opts...)
	return
}

// Creates a carrier.
//
// If a Shippo-supported code (`fedex`, `ups`, `usps`) is provided, the carrier is
// connected through Shippo and its service levels are auto-synced, initially
// hidden from the customer portal. Sandbox accounts skip the Shippo connection.
//
// This endpoint requires the permission: `carriers:create`.
func (r *OperationCarrierService) New(ctx context.Context, params OperationCarrierNewParams, opts ...option.RequestOption) (res *Carrier, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/carriers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a carrier by ID.
//
// This endpoint requires the permissions: `carriers:read`, `customers:read`,
// `suppliers:read`.
func (r *OperationCarrierService) Get(ctx context.Context, id string, query OperationCarrierGetParams, opts ...option.RequestOption) (res *Carrier, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a carrier's name and portal visibility.
//
// This endpoint requires the permission: `carriers:update`.
func (r *OperationCarrierService) Update(ctx context.Context, id string, params OperationCarrierUpdateParams, opts ...option.RequestOption) (res *Carrier, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of carriers for the current account.
//
// This endpoint requires the permissions: `carriers:read`, `customers:read`,
// `suppliers:read`.
func (r *OperationCarrierService) List(ctx context.Context, query OperationCarrierListParams, opts ...option.RequestOption) (res *ListCarrier, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/carriers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a carrier and all of its service levels.
//
// If the carrier is connected through Shippo, its Shippo carrier account is
// deactivated. System-owned carriers cannot be deleted.
//
// This endpoint requires the permission: `carriers:delete`.
func (r *OperationCarrierService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationCarrierDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a carrier.
//
// The property Name is required.
type CreateCarrierRequestParam struct {
	// Human-readable name for the carrier.
	//
	// Must be unique among your account's carriers.
	Name string `json:"name" api:"required"`
	// Your account number with this carrier.
	//
	// Required when `code` is `ups` or `usps`, which connect to Shippo using this
	// number; FedEx connects via OAuth instead.
	AccountNumber param.Opt[string] `json:"account_number,omitzero"`
	// Well-known carrier code.
	//
	// Omit for a custom carrier. Providing a Shippo-supported code (`fedex`, `ups`,
	// `usps`) connects the carrier through Shippo and auto-syncs its service levels.
	//
	// Any of "fedex", "ups", "usps", "will_call", "delivery", "ltl", "ltl1",
	// "freight_collect".
	Code CreateCarrierRequestCode `json:"code,omitzero"`
	// Carrier visibility in the customer portal.
	//
	// A `visible` carrier can be selected by your customers at checkout; a `hidden`
	// carrier is not offered there. New carriers are visible unless set to `hidden`.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CreateCarrierRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r CreateCarrierRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCarrierRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCarrierRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Well-known carrier code.
//
// Omit for a custom carrier. Providing a Shippo-supported code (`fedex`, `ups`,
// `usps`) connects the carrier through Shippo and auto-syncs its service levels.
type CreateCarrierRequestCode string

const (
	CreateCarrierRequestCodeFedex          CreateCarrierRequestCode = "fedex"
	CreateCarrierRequestCodeUps            CreateCarrierRequestCode = "ups"
	CreateCarrierRequestCodeUsps           CreateCarrierRequestCode = "usps"
	CreateCarrierRequestCodeWillCall       CreateCarrierRequestCode = "will_call"
	CreateCarrierRequestCodeDelivery       CreateCarrierRequestCode = "delivery"
	CreateCarrierRequestCodeLtl            CreateCarrierRequestCode = "ltl"
	CreateCarrierRequestCodeLtl1           CreateCarrierRequestCode = "ltl1"
	CreateCarrierRequestCodeFreightCollect CreateCarrierRequestCode = "freight_collect"
)

// Carrier visibility in the customer portal.
//
// A `visible` carrier can be selected by your customers at checkout; a `hidden`
// carrier is not offered there. New carriers are visible unless set to `hidden`.
type CreateCarrierRequestCustomerPortalVisibility string

const (
	CreateCarrierRequestCustomerPortalVisibilityVisible CreateCarrierRequestCustomerPortalVisibility = "visible"
	CreateCarrierRequestCustomerPortalVisibilityHidden  CreateCarrierRequestCustomerPortalVisibility = "hidden"
)

// List represents a paginated list of resources.
type ListCarrier struct {
	// Resources in this page.
	Data []Carrier `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListCarrierObject `json:"object" api:"required"`
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
func (r ListCarrier) RawJSON() string { return r.JSON.raw }
func (r *ListCarrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListCarrierObject string

const (
	ListCarrierObjectList ListCarrierObject = "list"
)

// Request to update a carrier.
type UpdateCarrierRequestParam struct {
	// Human-readable name for the carrier, unique among your account's carriers.
	Name param.Opt[string] `json:"name,omitzero"`
	// Carrier visibility in the customer portal.
	//
	// A `visible` carrier can be selected by your customers at checkout; a `hidden`
	// carrier is not offered there.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility UpdateCarrierRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r UpdateCarrierRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCarrierRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCarrierRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Carrier visibility in the customer portal.
//
// A `visible` carrier can be selected by your customers at checkout; a `hidden`
// carrier is not offered there.
type UpdateCarrierRequestCustomerPortalVisibility string

const (
	UpdateCarrierRequestCustomerPortalVisibilityVisible UpdateCarrierRequestCustomerPortalVisibility = "visible"
	UpdateCarrierRequestCustomerPortalVisibilityHidden  UpdateCarrierRequestCustomerPortalVisibility = "hidden"
)

type OperationCarrierDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationCarrierDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationCarrierDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationCarrierNewParams struct {
	// Request to create a carrier.
	CreateCarrierRequest CreateCarrierRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationCarrierNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCarrierRequest)
}
func (r *OperationCarrierNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierNewParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierGetParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a carrier.
	UpdateCarrierRequest UpdateCarrierRequestParam
	paramObj
}

func (r OperationCarrierUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCarrierRequest)
}
func (r *OperationCarrierUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierListParams struct {
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
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierListParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
