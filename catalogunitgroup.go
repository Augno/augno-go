// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and manage unit groups and their associated units.
//
// CatalogUnitGroupService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogUnitGroupService] method instead.
type CatalogUnitGroupService struct {
	options []option.RequestOption
	// List and manage unit groups and their associated units.
	Units CatalogUnitGroupUnitService
	// List and manage unit groups and their associated units.
	Actions CatalogUnitGroupActionService
}

// NewCatalogUnitGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogUnitGroupService(opts ...option.RequestOption) (r CatalogUnitGroupService) {
	r = CatalogUnitGroupService{}
	r.options = opts
	r.Units = NewCatalogUnitGroupUnitService(opts...)
	r.Actions = NewCatalogUnitGroupActionService(opts...)
	return
}

// Creates a unit group, optionally associating units with it in the same request.
//
// The name must be unique within the account, and the base unit and every
// associated unit must share the group's dimension.
//
// This endpoint requires the permission: `unit_groups:create`.
func (r *CatalogUnitGroupService) New(ctx context.Context, params CatalogUnitGroupNewParams, opts ...option.RequestOption) (res *UnitGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/unit-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a unit group by ID, including the system unit groups shared across all
// accounts.
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

// Partially updates a unit group.
//
// System unit groups cannot be modified, and a group's dimension is fixed once it
// is created.
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

// Deletes a unit group along with every unit association it contains.
//
// The units themselves are not deleted and remain available to other groups.
// System unit groups, which are shared across all accounts, cannot be deleted.
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
	//
	// Must be a unit of the group's `type`.
	BaseUnitID string `json:"base_unit_id" api:"required"`
	// Display name of the unit group.
	//
	// Must be unique within the account.
	Name string `json:"name" api:"required"`
	// The dimension shared by every unit in this group, such as mass, volume, or
	// currency.
	//
	// The base unit and all associated units must be of this dimension, and the
	// dimension cannot be changed after the group is created.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CreateUnitGroupRequestType `json:"type,omitzero" api:"required"`
	// Free-form notes about the unit group.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Units to associate with the group, each with its own discount and customer
	// portal visibility.
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

// The dimension shared by every unit in this group, such as mass, volume, or
// currency.
//
// The base unit and all associated units must be of this dimension, and the
// dimension cannot be changed after the group is created.
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
	//
	// Subtracted before `discount_percentage` is applied.
	DiscountFixed param.Opt[float64] `json:"discount_fixed,omitzero"`
	// Share of the unit's price removed when an order is placed in this unit.
	//
	// Expressed as a decimal fraction rather than a whole number, so `0.1` is a 10%
	// discount. Send `0` explicitly for no discount — omitting the field stores a
	// discount of `1`, which removes the entire price.
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

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListUnitGroup struct {
	// Resources in this page.
	Data []UnitGroup `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListUnitGroupObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
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

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListUnitGroupUnit struct {
	// Resources in this page.
	Data []UnitGroupUnit `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListUnitGroupUnitObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
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

// A named collection of units that share one dimension, defining which units a
// product can be ordered in.
//
// Each associated unit carries its own discount and customer portal visibility,
// applied when an order line is priced in that unit. A product takes its unit
// group from its product line, falling back to its item category.
type UnitGroup struct {
	// Unit group ID.
	ID string `json:"id" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
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
	// The dimension shared by every unit in this group, such as mass, volume, or
	// currency.
	//
	// Only units of this dimension can belong to the group, and the dimension is fixed
	// once the group is created.
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

// The dimension shared by every unit in this group, such as mass, volume, or
// currency.
//
// Only units of this dimension can belong to the group, and the dimension is fixed
// once the group is created.
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
	//
	// Subtracted before `discount_percentage` is applied.
	DiscountFixed float64 `json:"discount_fixed" api:"required"`
	// Share of the unit's price removed when an order is placed in this unit.
	//
	// Expressed as a decimal fraction rather than a whole number, so `0.1` is a 10%
	// discount and `0` is no discount.
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
	// ID of the unit to designate as the group's reference unit.
	//
	// Must be a unit of the group's dimension, which cannot itself be changed.
	BaseUnitID param.Opt[string] `json:"base_unit_id,omitzero"`
	// Display name of the unit group.
	//
	// Must be unique within the account.
	Name param.Opt[string] `json:"name,omitzero"`
	// Units to add to the group.
	//
	// Only units that are not already in the group can be listed here; use the
	// associated-unit update and delete endpoints to change or remove an existing
	// association. Associations left out of the list are untouched.
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
