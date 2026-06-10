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

// List and manage parts.
//
// CatalogPartService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogPartService] method instead.
type CatalogPartService struct {
	options []option.RequestOption
}

// NewCatalogPartService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogPartService(opts ...option.RequestOption) (r CatalogPartService) {
	r = CatalogPartService{}
	r.options = opts
	return
}

// Creates a part with the specified SKU and category.
func (r *CatalogPartService) New(ctx context.Context, params CatalogPartNewParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/parts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a part by ID.
func (r *CatalogPartService) Get(ctx context.Context, id string, query CatalogPartGetParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/parts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a part. Fields not provided retain their current values.
func (r *CatalogPartService) Update(ctx context.Context, id string, params CatalogPartUpdateParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/parts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of parts for the current account.
func (r *CatalogPartService) List(ctx context.Context, query CatalogPartListParams, opts ...option.RequestOption) (res *ListPart, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/parts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a part. Sets the deleted_at timestamp rather than removing the record.
func (r *CatalogPartService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/parts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a part.
//
// The properties CategoryID, SKU are required.
type CreatePartRequestParam struct {
	// Category ID.
	CategoryID string `json:"category_id" api:"required"`
	// SKU.
	SKU string `json:"sku" api:"required"`
	// Description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Attribute IDs to connect to the part at creation time.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// RateInput represents the input for creating or updating a rate.
	UnitCost RateInputParam `json:"unit_cost,omitzero"`
	// RateInput represents the input for creating or updating a rate.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r CreatePartRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePartRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePartRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List represents a paginated list of resources.
type ListPart struct {
	// Resources in this page.
	Data []Part `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPartObject `json:"object" api:"required"`
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
func (r ListPart) RawJSON() string { return r.JSON.raw }
func (r *ListPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPartObject string

const (
	ListPartObjectList ListPartObject = "list"
)

// Part resource.
type Part struct {
	// Part ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Item is an inventory item (product, material, or part).
	Item Item `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "part".
	Object PartObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Item        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Part) RawJSON() string { return r.JSON.raw }
func (r *Part) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PartObject string

const (
	PartObjectPart PartObject = "part"
)

// Request to partially update a part.
type UpdatePartRequestParam struct {
	// Description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// SKU.
	SKU param.Opt[string] `json:"sku,omitzero"`
	paramObj
}

func (r UpdatePartRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdatePartRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdatePartRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPartNewParams struct {
	// Request to create a part.
	CreatePartRequest CreatePartRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogPartNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePartRequest)
}
func (r *CatalogPartNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogPartNewParams]'s query parameters as `url.Values`.
func (r CatalogPartNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPartGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.category.properties",
	// "item.category.unit_group", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogPartGetParams]'s query parameters as `url.Values`.
func (r CatalogPartGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPartUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a part.
	UpdatePartRequest UpdatePartRequestParam
	paramObj
}

func (r CatalogPartUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdatePartRequest)
}
func (r *CatalogPartUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogPartUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogPartUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPartListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter parts created on or before this date.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter parts created on or after this date.
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	// Filter by attribute IDs.
	AttributeIDs []string `query:"attribute_ids,omitzero" json:"-"`
	// Filter by category IDs.
	CategoryIDs []string `query:"category_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.category.properties",
	// "item.category.unit_group", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogPartListParams]'s query parameters as `url.Values`.
func (r CatalogPartListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
