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

// List and manage item categories.
//
// CatalogItemCategoryService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemCategoryService] method instead.
type CatalogItemCategoryService struct {
	options []option.RequestOption
	// List and manage item categories.
	Properties CatalogItemCategoryPropertyService
}

// NewCatalogItemCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogItemCategoryService(opts ...option.RequestOption) (r CatalogItemCategoryService) {
	r = CatalogItemCategoryService{}
	r.options = opts
	r.Properties = NewCatalogItemCategoryPropertyService(opts...)
	return
}

// Creates an account-owned item category.
func (r *CatalogItemCategoryService) New(ctx context.Context, params CatalogItemCategoryNewParams, opts ...option.RequestOption) (res *ItemCategory, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/item-categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns an item category by ID. Includes account-specific and global system
// categories.
func (r *CatalogItemCategoryService) Get(ctx context.Context, id string, query CatalogItemCategoryGetParams, opts ...option.RequestOption) (res *ItemCategory, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an account-owned item category. Default system categories
// cannot be updated.
func (r *CatalogItemCategoryService) Update(ctx context.Context, id string, params CatalogItemCategoryUpdateParams, opts ...option.RequestOption) (res *ItemCategory, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of item categories for the current account, including
// account-specific and global system categories.
func (r *CatalogItemCategoryService) List(ctx context.Context, query CatalogItemCategoryListParams, opts ...option.RequestOption) (res *ListItemCategory, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/item-categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an account-owned item category. Default system categories cannot be
// deleted.
func (r *CatalogItemCategoryService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CatalogItemCategoryDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Changes the unit group associated with an item category. All items in the
// category are updated to use the new base unit asynchronously.
func (r *CatalogItemCategoryService) ChangeUnitGroup(ctx context.Context, unitGroupID string, body CatalogItemCategoryChangeUnitGroupParams, opts ...option.RequestOption) (res *CatalogItemCategoryChangeUnitGroupResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if unitGroupID == "" {
		err = errors.New("missing required unit_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s/unit-groups/%s", url.PathEscape(body.ID), url.PathEscape(unitGroupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Request to create an item category.
//
// The properties Name, Type, UnitGroupID are required.
type CreateItemCategoryRequestParam struct {
	// Display name.
	Name string `json:"name" api:"required"`
	// Item category type. Material categories are used to group materials, while
	// product categories are used to group products and parts.
	//
	// Any of "material_category", "product_category".
	Type CreateItemCategoryRequestType `json:"type,omitzero" api:"required"`
	// Unit group ID.
	UnitGroupID string `json:"unit_group_id" api:"required"`
	paramObj
}

func (r CreateItemCategoryRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateItemCategoryRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateItemCategoryRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Item category type. Material categories are used to group materials, while
// product categories are used to group products and parts.
type CreateItemCategoryRequestType string

const (
	CreateItemCategoryRequestTypeMaterialCategory CreateItemCategoryRequestType = "material_category"
	CreateItemCategoryRequestTypeProductCategory  CreateItemCategoryRequestType = "product_category"
)

// List represents a paginated list of resources.
type ListItemCategory struct {
	// Resources in this page.
	Data []ItemCategory `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListItemCategoryObject `json:"object" api:"required"`
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
func (r ListItemCategory) RawJSON() string { return r.JSON.raw }
func (r *ListItemCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListItemCategoryObject string

const (
	ListItemCategoryObjectList ListItemCategoryObject = "list"
)

// Request to partially update an item category.
type UpdateItemCategoryRequestParam struct {
	// Display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r UpdateItemCategoryRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateItemCategoryRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateItemCategoryRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogItemCategoryDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogItemCategoryDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryChangeUnitGroupResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogItemCategoryChangeUnitGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogItemCategoryChangeUnitGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryNewParams struct {
	// Request to create an item category.
	CreateItemCategoryRequest CreateItemCategoryRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "properties", "unit_group",
	// "unit_group.base_unit", "unit_group.associated_units",
	// "unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogItemCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateItemCategoryRequest)
}
func (r *CatalogItemCategoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogItemCategoryNewParams]'s query parameters as
// `url.Values`.
func (r CatalogItemCategoryNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemCategoryGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "properties", "unit_group",
	// "unit_group.base_unit", "unit_group.associated_units",
	// "unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemCategoryGetParams]'s query parameters as
// `url.Values`.
func (r CatalogItemCategoryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemCategoryUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "properties", "unit_group",
	// "unit_group.base_unit", "unit_group.associated_units",
	// "unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update an item category.
	UpdateItemCategoryRequest UpdateItemCategoryRequestParam
	paramObj
}

func (r CatalogItemCategoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateItemCategoryRequest)
}
func (r *CatalogItemCategoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogItemCategoryUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogItemCategoryUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemCategoryListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "properties", "unit_group",
	// "unit_group.base_unit", "unit_group.associated_units",
	// "unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by item category type.
	//
	// Any of "material_category", "product_category".
	Type CatalogItemCategoryListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemCategoryListParams]'s query parameters as
// `url.Values`.
func (r CatalogItemCategoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by item category type.
type CatalogItemCategoryListParamsType string

const (
	CatalogItemCategoryListParamsTypeMaterialCategory CatalogItemCategoryListParamsType = "material_category"
	CatalogItemCategoryListParamsTypeProductCategory  CatalogItemCategoryListParamsType = "product_category"
)

type CatalogItemCategoryChangeUnitGroupParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
