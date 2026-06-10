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

// List and manage products.
//
// CatalogProductService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogProductService] method instead.
type CatalogProductService struct {
	options []option.RequestOption
}

// NewCatalogProductService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogProductService(opts ...option.RequestOption) (r CatalogProductService) {
	r = CatalogProductService{}
	r.options = opts
	return
}

// Creates a product.
func (r *CatalogProductService) New(ctx context.Context, params CatalogProductNewParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a product by ID.
func (r *CatalogProductService) Get(ctx context.Context, id string, query CatalogProductGetParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a product.
func (r *CatalogProductService) Update(ctx context.Context, id string, params CatalogProductUpdateParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of products for the target account.
func (r *CatalogProductService) List(ctx context.Context, query CatalogProductListParams, opts ...option.RequestOption) (res *ListProduct, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Soft-deletes a product and returns the deleted product.
func (r *CatalogProductService) Delete(ctx context.Context, id string, body CatalogProductDeleteParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Changes the product line assignment for a product.
func (r *CatalogProductService) ChangeProductLine(ctx context.Context, productLineID string, params CatalogProductChangeProductLineParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if productLineID == "" {
		err = errors.New("missing required product_line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s/product-line/%s", url.PathEscape(params.ID), url.PathEscape(productLineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// CreateProductRequest is the request to create a product.
//
// The properties CategoryID, SKU, Type are required.
type CreateProductRequestParam struct {
	// Category ID.
	CategoryID string `json:"category_id" api:"required"`
	// SKU.
	SKU string `json:"sku" api:"required"`
	// Product type code.
	//
	// Any of "sale", "service", "shipping", "credit", "return", "tax".
	Type CreateProductRequestType `json:"type,omitzero" api:"required"`
	// Description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Product line ID.
	ProductLineID param.Opt[string] `json:"product_line_id,omitzero"`
	// Attribute IDs to connect to the product at creation time.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// Whether visible in the customer portal.
	//
	// Any of "visible", "hidden".
	PortalVisibility CreateProductRequestPortalVisibility `json:"portal_visibility,omitzero"`
	// RateInput represents the input for creating or updating a rate.
	UnitCost RateInputParam `json:"unit_cost,omitzero"`
	// RateInput represents the input for creating or updating a rate.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r CreateProductRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateProductRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateProductRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product type code.
type CreateProductRequestType string

const (
	CreateProductRequestTypeSale     CreateProductRequestType = "sale"
	CreateProductRequestTypeService  CreateProductRequestType = "service"
	CreateProductRequestTypeShipping CreateProductRequestType = "shipping"
	CreateProductRequestTypeCredit   CreateProductRequestType = "credit"
	CreateProductRequestTypeReturn   CreateProductRequestType = "return"
	CreateProductRequestTypeTax      CreateProductRequestType = "tax"
)

// Whether visible in the customer portal.
type CreateProductRequestPortalVisibility string

const (
	CreateProductRequestPortalVisibilityVisible CreateProductRequestPortalVisibility = "visible"
	CreateProductRequestPortalVisibilityHidden  CreateProductRequestPortalVisibility = "hidden"
)

// List represents a paginated list of resources.
type ListProduct struct {
	// Resources in this page.
	Data []Product `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductObject `json:"object" api:"required"`
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
func (r ListProduct) RawJSON() string { return r.JSON.raw }
func (r *ListProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductObject string

const (
	ListProductObjectList ListProductObject = "list"
)

// Product with expandable item, product line, and product type.
type Product struct {
	// Product ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Item is an inventory item (product, material, or part).
	Item Item `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "product".
	Object ProductObject `json:"object" api:"required"`
	// Product portal visibility.
	//
	// Any of "visible", "hidden".
	PortalVisibility ProductPortalVisibility `json:"portal_visibility" api:"required"`
	// Product line resource.
	ProductLine ProductLine `json:"product_line" api:"required"`
	// Product type code.
	//
	// Any of "sale", "service", "shipping", "credit", "return", "tax".
	Type ProductType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Item             respjson.Field
		Object           respjson.Field
		PortalVisibility respjson.Field
		ProductLine      respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Product) RawJSON() string { return r.JSON.raw }
func (r *Product) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductObject string

const (
	ProductObjectProduct ProductObject = "product"
)

// Product portal visibility.
type ProductPortalVisibility string

const (
	ProductPortalVisibilityVisible ProductPortalVisibility = "visible"
	ProductPortalVisibilityHidden  ProductPortalVisibility = "hidden"
)

// Product type code.
type ProductType string

const (
	ProductTypeSale     ProductType = "sale"
	ProductTypeService  ProductType = "service"
	ProductTypeShipping ProductType = "shipping"
	ProductTypeCredit   ProductType = "credit"
	ProductTypeReturn   ProductType = "return"
	ProductTypeTax      ProductType = "tax"
)

// UpdateProductRequest is the request to partially update a product.
type UpdateProductRequestParam struct {
	// Description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// SKU.
	SKU param.Opt[string] `json:"sku,omitzero"`
	// Whether visible in the customer portal.
	//
	// Any of "visible", "hidden".
	PortalVisibility UpdateProductRequestPortalVisibility `json:"portal_visibility,omitzero"`
	// RateInput represents the input for creating or updating a rate.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r UpdateProductRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateProductRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateProductRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether visible in the customer portal.
type UpdateProductRequestPortalVisibility string

const (
	UpdateProductRequestPortalVisibilityVisible UpdateProductRequestPortalVisibility = "visible"
	UpdateProductRequestPortalVisibilityHidden  UpdateProductRequestPortalVisibility = "hidden"
)

type CatalogProductNewParams struct {
	// CreateProductRequest is the request to create a product.
	CreateProductRequest CreateProductRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogProductNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateProductRequest)
}
func (r *CatalogProductNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductNewParams]'s query parameters as
// `url.Values`.
func (r CatalogProductNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductGetParams]'s query parameters as
// `url.Values`.
func (r CatalogProductGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// UpdateProductRequest is the request to partially update a product.
	UpdateProductRequest UpdateProductRequestParam
	paramObj
}

func (r CatalogProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateProductRequest)
}
func (r *CatalogProductUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogProductUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// End of creation date range.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Start of creation date range.
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	// Filter by attribute IDs.
	AttributeIDs []string `query:"attribute_ids,omitzero" json:"-"`
	// Filter by category IDs.
	CategoryIDs []string `query:"category_ids,omitzero" json:"-"`
	// Filter by customer IDs.
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by customer portal visibility.
	//
	// Any of "visible", "hidden".
	PortalVisibility CatalogProductListParamsPortalVisibility `query:"portal_visibility,omitzero" json:"-"`
	// Filter by product line IDs.
	ProductLineIDs []string `query:"product_line_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductListParams]'s query parameters as
// `url.Values`.
func (r CatalogProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by customer portal visibility.
type CatalogProductListParamsPortalVisibility string

const (
	CatalogProductListParamsPortalVisibilityVisible CatalogProductListParamsPortalVisibility = "visible"
	CatalogProductListParamsPortalVisibilityHidden  CatalogProductListParamsPortalVisibility = "hidden"
)

type CatalogProductDeleteParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductDeleteParams]'s query parameters as
// `url.Values`.
func (r CatalogProductDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductChangeProductLineParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductChangeProductLineParams]'s query parameters
// as `url.Values`.
func (r CatalogProductChangeProductLineParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
