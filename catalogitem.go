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
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List and manage inventory items.
//
// CatalogItemService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemService] method instead.
type CatalogItemService struct {
	options []option.RequestOption
	// List and manage inventory items.
	Attributes CatalogItemAttributeService
}

// NewCatalogItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogItemService(opts ...option.RequestOption) (r CatalogItemService) {
	r = CatalogItemService{}
	r.options = opts
	r.Attributes = NewCatalogItemAttributeService(opts...)
	return
}

// Returns an item by ID.
func (r *CatalogItemService) Get(ctx context.Context, id string, query CatalogItemGetParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of items.
func (r *CatalogItemService) List(ctx context.Context, query CatalogItemListParams, opts ...option.RequestOption) (res *ListItem, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Moves an item to a different category.
//
// The item's rate units (unit value, unit cost, burn rate) and any related
// order-point, consumption, and production quantity units are updated to the new
// category's base unit. Re-assigning the item's current category is a no-op.
func (r *CatalogItemService) ChangeCategory(ctx context.Context, categoryID string, params CatalogItemChangeCategoryParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/category/%s", url.PathEscape(params.ID), url.PathEscape(categoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns inventory quantities for an item, including on-hand, reserved,
// available-to-promise, and short amounts.
func (r *CatalogItemService) GetInventory(ctx context.Context, id string, query CatalogItemGetInventoryParams, opts ...option.RequestOption) (res *ItemInventory, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/inventory", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Item is an inventory item (product, material, or part).
type Item struct {
	// Item ID.
	ID string `json:"id" api:"required"`
	// List represents a paginated list of resources.
	Attributes ListAttribute `json:"attributes" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	BurnRate Rate `json:"burn_rate" api:"required"`
	// A grouping of related catalog items that defines the unit group and properties
	// available to the items within it.
	Category ItemCategory `json:"category" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Item description.
	Description string `json:"description" api:"required"`
	// Free-form notes about the item.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "item".
	Object ItemObject `json:"object" api:"required"`
	// Stock keeping unit code, unique within the account.
	SKU string `json:"sku" api:"required"`
	// What kind of item this is.
	//
	// - `product`: a finished product.
	// - `material`: a raw material or component consumed in production.
	// - `part`: a part used in production.
	//
	// Any of "product", "material", "part".
	Type ItemType `json:"type" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitCost Rate `json:"unit_cost" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitValue Rate `json:"unit_value" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		BurnRate    respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Notes       respjson.Field
		Object      respjson.Field
		SKU         respjson.Field
		Type        respjson.Field
		UnitCost    respjson.Field
		UnitValue   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Item) RawJSON() string { return r.JSON.raw }
func (r *Item) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ItemObject string

const (
	ItemObjectItem ItemObject = "item"
)

// What kind of item this is.
//
// - `product`: a finished product.
// - `material`: a raw material or component consumed in production.
// - `part`: a part used in production.
type ItemType string

const (
	ItemTypeProduct  ItemType = "product"
	ItemTypeMaterial ItemType = "material"
	ItemTypePart     ItemType = "part"
)

// A grouping of related catalog items that defines the unit group and properties
// available to the items within it.
type ItemCategory struct {
	// Item category ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the item category.
	Name string `json:"name" api:"required"`
	// Free-form notes about the item category.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "item_category".
	Object ItemCategoryObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// List represents a paginated list of resources.
	Properties ListProperty `json:"properties" api:"required"`
	// What kind of items this category groups.
	//
	// An item can only be assigned to a category whose type matches the item's `type`.
	//
	//   - `material_category`: groups raw materials and components (items of type
	//     `material`).
	//   - `product_category`: groups finished products and parts (items of type
	//     `product` or `part`).
	//
	// Any of "material_category", "product_category".
	Type ItemCategoryType `json:"type" api:"required"`
	// Named collection of units sharing one dimension, defining which units products
	// can be ordered in along with per-unit discounts and customer portal visibility.
	UnitGroup UnitGroup `json:"unit_group" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		UnitGroup   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemCategory) RawJSON() string { return r.JSON.raw }
func (r *ItemCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ItemCategoryObject string

const (
	ItemCategoryObjectItemCategory ItemCategoryObject = "item_category"
)

// What kind of items this category groups.
//
// An item can only be assigned to a category whose type matches the item's `type`.
//
//   - `material_category`: groups raw materials and components (items of type
//     `material`).
//   - `product_category`: groups finished products and parts (items of type
//     `product` or `part`).
type ItemCategoryType string

const (
	ItemCategoryTypeMaterialCategory ItemCategoryType = "material_category"
	ItemCategoryTypeProductCategory  ItemCategoryType = "product_category"
)

// ItemInventory contains inventory quantities for an item.
type ItemInventory struct {
	// Value with an associated unit.
	AvailableToPromise Quantity `json:"available_to_promise" api:"required"`
	// Resource type identifier.
	//
	// Any of "item_inventory".
	Object ItemInventoryObject `json:"object" api:"required"`
	// Value with an associated unit.
	OnHand Quantity `json:"on_hand" api:"required"`
	// Value with an associated unit.
	Reserved Quantity `json:"reserved" api:"required"`
	// Value with an associated unit.
	Short Quantity `json:"short" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableToPromise respjson.Field
		Object             respjson.Field
		OnHand             respjson.Field
		Reserved           respjson.Field
		Short              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemInventory) RawJSON() string { return r.JSON.raw }
func (r *ItemInventory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ItemInventoryObject string

const (
	ItemInventoryObjectItemInventory ItemInventoryObject = "item_inventory"
)

// List represents a paginated list of resources.
type ListItem struct {
	// Resources in this page.
	Data []Item `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListItemObject `json:"object" api:"required"`
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
func (r ListItem) RawJSON() string { return r.JSON.raw }
func (r *ListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListItemObject string

const (
	ListItemObjectList ListItemObject = "list"
)

// Value with an associated unit.
type Quantity struct {
	// Quantity ID.
	ID string `json:"id" api:"required"`
	// Formatted value with unit abbreviation (e.g. "$1,234.56" or "100 kg").
	DisplayValue string `json:"display_value" api:"required"`
	// Resource type identifier.
	//
	// Any of "quantity".
	Object QuantityObject `json:"object" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	Unit Unit `json:"unit" api:"required"`
	// Raw decimal value of the quantity, as a string to preserve precision.
	//
	// This is the unformatted machine value; see `display_value` for the
	// human-readable rendering with unit and thousands separators.
	Value string `json:"value" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DisplayValue respjson.Field
		Object       respjson.Field
		Unit         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Quantity) RawJSON() string { return r.JSON.raw }
func (r *Quantity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuantityObject string

const (
	QuantityObjectQuantity QuantityObject = "quantity"
)

// Value expressed as a ratio of two units, such as a price per kilogram or a
// throughput per hour.
type Rate struct {
	// Rate ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unit of measurement used for conversions and product quantities.
	DenominatorUnit Unit `json:"denominator_unit" api:"required"`
	// Human-readable formatted value (e.g. "$25.50 / kg" or "100 kg / hr").
	DisplayValue string `json:"display_value" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	NumeratorUnit Unit `json:"numerator_unit" api:"required"`
	// Resource type identifier.
	//
	// Any of "rate".
	Object RateObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Decimal value of the rate, as a string to preserve precision.
	//
	// Expressed as the amount of the numerator unit per one denominator unit.
	Value string `json:"value" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		DenominatorUnit respjson.Field
		DisplayValue    respjson.Field
		NumeratorUnit   respjson.Field
		Object          respjson.Field
		UpdatedAt       respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rate) RawJSON() string { return r.JSON.raw }
func (r *Rate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type RateObject string

const (
	RateObjectRate RateObject = "rate"
)

type CatalogItemGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemGetParams]'s query parameters as `url.Values`.
func (r CatalogItemGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter items created on or before this date.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter items created on or after this date.
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	// Filter by supplier ID.
	SupplierID param.Opt[string] `query:"supplier_id,omitzero" json:"-"`
	// Filter by attribute IDs.
	AttributeIDs []string `query:"attribute_ids,omitzero" json:"-"`
	// Filter by category IDs.
	CategoryIDs []string `query:"category_ids,omitzero" json:"-"`
	// Filter by customer account IDs (only items whose product line is accessible to
	// any of these customers).
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by product line IDs (only items whose product belongs to one of these
	// lines).
	ProductLineIDs []string `query:"product_line_ids,omitzero" json:"-"`
	// Restricts results based on where the item is produced in its production flow.
	//
	//   - `all`: no restriction.
	//   - `initial_only`: only items produced by an initial production step, i.e. a step
	//     with no upstream steps feeding into it.
	//
	// Any of "all", "initial_only".
	SubassemblyFilter CatalogItemListParamsSubassemblyFilter `query:"subassembly_filter,omitzero" json:"-"`
	// Filter to items of these types (`product`, `material`, `part`).
	Types []string `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemListParams]'s query parameters as `url.Values`.
func (r CatalogItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restricts results based on where the item is produced in its production flow.
//
//   - `all`: no restriction.
//   - `initial_only`: only items produced by an initial production step, i.e. a step
//     with no upstream steps feeding into it.
type CatalogItemListParamsSubassemblyFilter string

const (
	CatalogItemListParamsSubassemblyFilterAll         CatalogItemListParamsSubassemblyFilter = "all"
	CatalogItemListParamsSubassemblyFilterInitialOnly CatalogItemListParamsSubassemblyFilter = "initial_only"
)

type CatalogItemChangeCategoryParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemChangeCategoryParams]'s query parameters as
// `url.Values`.
func (r CatalogItemChangeCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemGetInventoryParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "on_hand", "reserved", "available_to_promise", "short".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemGetInventoryParams]'s query parameters as
// `url.Values`.
func (r CatalogItemGetInventoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
