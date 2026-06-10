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

// List and manage product lines.
//
// CatalogProductLineService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogProductLineService] method instead.
type CatalogProductLineService struct {
	options []option.RequestOption
}

// NewCatalogProductLineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogProductLineService(opts ...option.RequestOption) (r CatalogProductLineService) {
	r = CatalogProductLineService{}
	r.options = opts
	return
}

// Creates an account-owned product line.
func (r *CatalogProductLineService) New(ctx context.Context, params CatalogProductLineNewParams, opts ...option.RequestOption) (res *ProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/product-lines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a product line by ID, including system-owned product lines accessible to
// the account.
func (r *CatalogProductLineService) Get(ctx context.Context, id string, query CatalogProductLineGetParams, opts ...option.RequestOption) (res *ProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/product-lines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an account-owned product line. Default system product lines
// cannot be updated.
func (r *CatalogProductLineService) Update(ctx context.Context, id string, params CatalogProductLineUpdateParams, opts ...option.RequestOption) (res *ProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/product-lines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of product lines, including account-owned and system
// product lines.
func (r *CatalogProductLineService) List(ctx context.Context, query CatalogProductLineListParams, opts ...option.RequestOption) (res *ListProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/product-lines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an account-owned product line. Default system product lines cannot be
// deleted.
func (r *CatalogProductLineService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CatalogProductLineDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/product-lines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a product line.
//
// The properties CommissionPolicy, FreightPolicy, Name, UnitGroupID are required.
type CreateProductLineRequestParam struct {
	// Commission policy of products in this product line.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CreateProductLineRequestCommissionPolicy `json:"commission_policy,omitzero" api:"required"`
	// Freight policy for all items in this product line.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy CreateProductLineRequestFreightPolicy `json:"freight_policy,omitzero" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Unit group ID associated with this product line. This unit group dictates the
	// units that products in this product line may be purchased in.
	UnitGroupID string `json:"unit_group_id" api:"required"`
	paramObj
}

func (r CreateProductLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateProductLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateProductLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Commission policy of products in this product line.
type CreateProductLineRequestCommissionPolicy string

const (
	CreateProductLineRequestCommissionPolicyCommissionApplied CreateProductLineRequestCommissionPolicy = "commission_applied"
	CreateProductLineRequestCommissionPolicyCommissionExempt  CreateProductLineRequestCommissionPolicy = "commission_exempt"
)

// Freight policy for all items in this product line.
type CreateProductLineRequestFreightPolicy string

const (
	CreateProductLineRequestFreightPolicyFreeFreight   CreateProductLineRequestFreightPolicy = "free_freight"
	CreateProductLineRequestFreightPolicyBilledFreight CreateProductLineRequestFreightPolicy = "billed_freight"
)

// List represents a paginated list of resources.
type ListProductLine struct {
	// Resources in this page.
	Data []ProductLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductLineObject `json:"object" api:"required"`
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
func (r ListProductLine) RawJSON() string { return r.JSON.raw }
func (r *ListProductLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductLineObject string

const (
	ListProductLineObjectList ListProductLineObject = "list"
)

// Product line resource.
type ProductLine struct {
	// Product line ID.
	ID string `json:"id" api:"required"`
	// Commission policy of products in this product line.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy ProductLineCommissionPolicy `json:"commission_policy" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Description.
	Description string `json:"description" api:"required"`
	// Freight policy for all items in this product line.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy ProductLineFreightPolicy `json:"freight_policy" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "product_line".
	Object ProductLineObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// UnitGroup is a unit group resource.
	UnitGroup UnitGroup `json:"unit_group" api:"required"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CommissionPolicy respjson.Field
		CreatedAt        respjson.Field
		Description      respjson.Field
		FreightPolicy    respjson.Field
		Name             respjson.Field
		Notes            respjson.Field
		Object           respjson.Field
		Owner            respjson.Field
		UnitGroup        respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductLine) RawJSON() string { return r.JSON.raw }
func (r *ProductLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Commission policy of products in this product line.
type ProductLineCommissionPolicy string

const (
	ProductLineCommissionPolicyCommissionApplied ProductLineCommissionPolicy = "commission_applied"
	ProductLineCommissionPolicyCommissionExempt  ProductLineCommissionPolicy = "commission_exempt"
)

// Freight policy for all items in this product line.
type ProductLineFreightPolicy string

const (
	ProductLineFreightPolicyFreeFreight   ProductLineFreightPolicy = "free_freight"
	ProductLineFreightPolicyBilledFreight ProductLineFreightPolicy = "billed_freight"
)

// Resource type identifier.
type ProductLineObject string

const (
	ProductLineObjectProductLine ProductLineObject = "product_line"
)

// Request to partially update a product line.
type UpdateProductLineRequestParam struct {
	// Display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Unit group ID associated with this product line. This unit group dictates the
	// units that products in this product line may be purchased in.
	UnitGroupID param.Opt[string] `json:"unit_group_id,omitzero"`
	// Commission policy of products in this product line.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy UpdateProductLineRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// Freight policy for all items in this product line.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy UpdateProductLineRequestFreightPolicy `json:"freight_policy,omitzero"`
	paramObj
}

func (r UpdateProductLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateProductLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateProductLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Commission policy of products in this product line.
type UpdateProductLineRequestCommissionPolicy string

const (
	UpdateProductLineRequestCommissionPolicyCommissionApplied UpdateProductLineRequestCommissionPolicy = "commission_applied"
	UpdateProductLineRequestCommissionPolicyCommissionExempt  UpdateProductLineRequestCommissionPolicy = "commission_exempt"
)

// Freight policy for all items in this product line.
type UpdateProductLineRequestFreightPolicy string

const (
	UpdateProductLineRequestFreightPolicyFreeFreight   UpdateProductLineRequestFreightPolicy = "free_freight"
	UpdateProductLineRequestFreightPolicyBilledFreight UpdateProductLineRequestFreightPolicy = "billed_freight"
)

type CatalogProductLineDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogProductLineDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogProductLineDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogProductLineNewParams struct {
	// Request to create a product line.
	CreateProductLineRequest CreateProductLineRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "unit_group".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogProductLineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateProductLineRequest)
}
func (r *CatalogProductLineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductLineNewParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductLineGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "unit_group".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductLineGetParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductLineUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "unit_group".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a product line.
	UpdateProductLineRequest UpdateProductLineRequestParam
	paramObj
}

func (r CatalogProductLineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateProductLineRequest)
}
func (r *CatalogProductLineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductLineUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductLineListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "unit_group".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductLineListParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
