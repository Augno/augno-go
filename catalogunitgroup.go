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

// List and manage unit groups and their associated units.
//
// CatalogUnitGroupService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogUnitGroupService] method instead.
type CatalogUnitGroupService struct {
	options []option.RequestOption
	// List and manage unit groups and their associated units.
	Units CatalogUnitGroupUnitService
}

// NewCatalogUnitGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogUnitGroupService(opts ...option.RequestOption) (r CatalogUnitGroupService) {
	r = CatalogUnitGroupService{}
	r.options = opts
	r.Units = NewCatalogUnitGroupUnitService(opts...)
	return
}

// Creates a unit group with optional associated units.
//
// This endpoint requires the permission: `unit_groups:create`.
func (r *CatalogUnitGroupService) New(ctx context.Context, params CatalogUnitGroupNewParams, opts ...option.RequestOption) (res *UnitGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/unit-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a unit group by ID.
//
// This endpoint requires the permission: `unit_groups:read`.
func (r *CatalogUnitGroupService) Get(ctx context.Context, id string, query CatalogUnitGroupGetParams, opts ...option.RequestOption) (res *UnitGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a unit group. System unit groups cannot be updated.
//
// This endpoint requires the permission: `unit_groups:update`.
func (r *CatalogUnitGroupService) Update(ctx context.Context, id string, params CatalogUnitGroupUpdateParams, opts ...option.RequestOption) (res *UnitGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of unit groups, including system unit groups.
//
// This endpoint requires the permission: `unit_groups:read`.
func (r *CatalogUnitGroupService) List(ctx context.Context, query CatalogUnitGroupListParams, opts ...option.RequestOption) (res *ListUnitGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/unit-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a unit group and all of its associated units. System unit groups cannot
// be deleted.
//
// This endpoint requires the permission: `unit_groups:delete`.
func (r *CatalogUnitGroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CatalogUnitGroupDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a unit group.
//
// The properties BaseUnitID, Name, Type are required.
type CreateUnitGroupRequestParam struct {
	// ID of the unit to designate as the group's reference unit.
	BaseUnitID string `json:"base_unit_id" api:"required"`
	// Display name of the unit group.
	//
	// Must be unique within the account.
	Name string `json:"name" api:"required"`
	// Dimension shared by every unit in this group.
	//
	// All associated units must be of this dimension.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CreateUnitGroupRequestType `json:"type,omitzero" api:"required"`
	// Free-form notes about the unit group.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Associated units to create with the group.
	AssociatedUnits []CreateUnitGroupUnitParam `json:"associated_units,omitzero"`
	paramObj
}

func (r CreateUnitGroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateUnitGroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateUnitGroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dimension shared by every unit in this group.
//
// All associated units must be of this dimension.
type CreateUnitGroupRequestType string

const (
	CreateUnitGroupRequestTypeCurrency    CreateUnitGroupRequestType = "currency"
	CreateUnitGroupRequestTypeQuantity    CreateUnitGroupRequestType = "quantity"
	CreateUnitGroupRequestTypeTime        CreateUnitGroupRequestType = "time"
	CreateUnitGroupRequestTypeMass        CreateUnitGroupRequestType = "mass"
	CreateUnitGroupRequestTypeVolume      CreateUnitGroupRequestType = "volume"
	CreateUnitGroupRequestTypeLength      CreateUnitGroupRequestType = "length"
	CreateUnitGroupRequestTypeTemperature CreateUnitGroupRequestType = "temperature"
	CreateUnitGroupRequestTypeArea        CreateUnitGroupRequestType = "area"
)

// Parameters for associating a unit with a unit group.
//
// The property UnitID is required.
type CreateUnitGroupUnitParam struct {
	// ID of the unit to associate with the group.
	//
	// The unit's dimension must match the group's `type`.
	UnitID string `json:"unit_id" api:"required"`
	// Flat amount subtracted from the unit's price when an order is placed in this
	// unit.
	DiscountFixed param.Opt[float64] `json:"discount_fixed,omitzero"`
	// Percentage discount applied to the unit's price when an order is placed in this
	// unit (e.g. `10` is a 10% discount).
	DiscountPercentage param.Opt[float64] `json:"discount_percentage,omitzero"`
	// Whether the unit is shown to customers in the customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CreateUnitGroupUnitParamCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r CreateUnitGroupUnitParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateUnitGroupUnitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateUnitGroupUnitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the unit is shown to customers in the customer portal.
type CreateUnitGroupUnitParamCustomerPortalVisibility string

const (
	CreateUnitGroupUnitParamCustomerPortalVisibilityVisible CreateUnitGroupUnitParamCustomerPortalVisibility = "visible"
	CreateUnitGroupUnitParamCustomerPortalVisibilityHidden  CreateUnitGroupUnitParamCustomerPortalVisibility = "hidden"
)

// List represents a paginated list of resources.
type ListUnitGroup struct {
	// Resources in this page.
	Data []UnitGroup `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListUnitGroupObject `json:"object" api:"required"`
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
func (r ListUnitGroup) RawJSON() string { return r.JSON.raw }
func (r *ListUnitGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListUnitGroupObject string

const (
	ListUnitGroupObjectList ListUnitGroupObject = "list"
)

// List represents a paginated list of resources.
type ListUnitGroupUnit struct {
	// Resources in this page.
	Data []UnitGroupUnit `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListUnitGroupUnitObject `json:"object" api:"required"`
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
func (r ListUnitGroupUnit) RawJSON() string { return r.JSON.raw }
func (r *ListUnitGroupUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListUnitGroupUnitObject string

const (
	ListUnitGroupUnitObjectList ListUnitGroupUnitObject = "list"
)

// Named collection of units sharing one dimension, defining which units products
// can be ordered in along with per-unit discounts and customer portal visibility.
type UnitGroup struct {
	// Unit group ID.
	ID string `json:"id" api:"required"`
	// List represents a paginated list of resources.
	AssociatedUnits ListUnitGroupUnit `json:"associated_units" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	BaseUnit Unit `json:"base_unit" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the unit group.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Free-form notes about the unit group.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "unit_group".
	Object UnitGroupObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Physical dimension shared by every unit in this group, such as mass, volume, or
	// currency.
	//
	// Only units of this dimension can belong to the group.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type UnitGroupType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AssociatedUnits respjson.Field
		BaseUnit        respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Notes           respjson.Field
		Object          respjson.Field
		Owner           respjson.Field
		Type            respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitGroup) RawJSON() string { return r.JSON.raw }
func (r *UnitGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type UnitGroupObject string

const (
	UnitGroupObjectUnitGroup UnitGroupObject = "unit_group"
)

// Physical dimension shared by every unit in this group, such as mass, volume, or
// currency.
//
// Only units of this dimension can belong to the group.
type UnitGroupType string

const (
	UnitGroupTypeCurrency    UnitGroupType = "currency"
	UnitGroupTypeQuantity    UnitGroupType = "quantity"
	UnitGroupTypeTime        UnitGroupType = "time"
	UnitGroupTypeMass        UnitGroupType = "mass"
	UnitGroupTypeVolume      UnitGroupType = "volume"
	UnitGroupTypeLength      UnitGroupType = "length"
	UnitGroupTypeTemperature UnitGroupType = "temperature"
	UnitGroupTypeArea        UnitGroupType = "area"
)

// Membership of a unit in a unit group, carrying the discount and customer portal
// visibility settings applied when ordering in that unit.
type UnitGroupUnit struct {
	// Unit group unit ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this unit is shown to customers in the customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility UnitGroupUnitCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Flat amount subtracted from the unit's price when an order is placed in this
	// unit.
	DiscountFixed float64 `json:"discount_fixed" api:"required"`
	// Percentage discount applied to the unit's price when an order is placed in this
	// unit (e.g. `10` is a 10% discount).
	DiscountPercentage float64 `json:"discount_percentage" api:"required"`
	// Resource type identifier.
	//
	// Any of "unit_group_unit".
	Object UnitGroupUnitObject `json:"object" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	Unit Unit `json:"unit" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		CustomerPortalVisibility respjson.Field
		DiscountFixed            respjson.Field
		DiscountPercentage       respjson.Field
		Object                   respjson.Field
		Unit                     respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitGroupUnit) RawJSON() string { return r.JSON.raw }
func (r *UnitGroupUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this unit is shown to customers in the customer portal.
type UnitGroupUnitCustomerPortalVisibility string

const (
	UnitGroupUnitCustomerPortalVisibilityVisible UnitGroupUnitCustomerPortalVisibility = "visible"
	UnitGroupUnitCustomerPortalVisibilityHidden  UnitGroupUnitCustomerPortalVisibility = "hidden"
)

// Resource type identifier.
type UnitGroupUnitObject string

const (
	UnitGroupUnitObjectUnitGroupUnit UnitGroupUnitObject = "unit_group_unit"
)

// Request to partially update a unit group.
type UpdateUnitGroupRequestParam struct {
	// Free-form notes about the unit group.
	//
	// Set to `null` to clear.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// ID of the group's base unit.
	BaseUnitID param.Opt[string] `json:"base_unit_id,omitzero"`
	// Display name of the unit group.
	//
	// Must be unique within the account.
	Name param.Opt[string] `json:"name,omitzero"`
	// Associated units to add or update in the group.
	//
	// Upserted by unit: a listed unit already in the group has its association
	// updated, otherwise it is added. Existing units not in the list are preserved.
	AssociatedUnits []CreateUnitGroupUnitParam `json:"associated_units,omitzero"`
	paramObj
}

func (r UpdateUnitGroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateUnitGroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateUnitGroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogUnitGroupDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogUnitGroupDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogUnitGroupDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogUnitGroupNewParams struct {
	// Request to create a unit group.
	CreateUnitGroupRequest CreateUnitGroupRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "base_unit", "associated_units".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogUnitGroupNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateUnitGroupRequest)
}
func (r *CatalogUnitGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitGroupNewParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGroupGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "base_unit", "associated_units".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitGroupGetParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGroupUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "base_unit", "associated_units".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a unit group.
	UpdateUnitGroupRequest UpdateUnitGroupRequestParam
	paramObj
}

func (r CatalogUnitGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateUnitGroupRequest)
}
func (r *CatalogUnitGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitGroupUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGroupListParams struct {
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
	// Any of "owner", "owner.account", "base_unit", "associated_units".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by unit dimension.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CatalogUnitGroupListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitGroupListParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by unit dimension.
type CatalogUnitGroupListParamsType string

const (
	CatalogUnitGroupListParamsTypeCurrency    CatalogUnitGroupListParamsType = "currency"
	CatalogUnitGroupListParamsTypeQuantity    CatalogUnitGroupListParamsType = "quantity"
	CatalogUnitGroupListParamsTypeTime        CatalogUnitGroupListParamsType = "time"
	CatalogUnitGroupListParamsTypeMass        CatalogUnitGroupListParamsType = "mass"
	CatalogUnitGroupListParamsTypeVolume      CatalogUnitGroupListParamsType = "volume"
	CatalogUnitGroupListParamsTypeLength      CatalogUnitGroupListParamsType = "length"
	CatalogUnitGroupListParamsTypeTemperature CatalogUnitGroupListParamsType = "temperature"
	CatalogUnitGroupListParamsTypeArea        CatalogUnitGroupListParamsType = "area"
)
