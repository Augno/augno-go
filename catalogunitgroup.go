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
func (r *CatalogUnitGroupService) New(ctx context.Context, params CatalogUnitGroupNewParams, opts ...option.RequestOption) (res *UnitGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/unit-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a unit group by ID.
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
func (r *CatalogUnitGroupService) List(ctx context.Context, query CatalogUnitGroupListParams, opts ...option.RequestOption) (res *ListUnitGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/unit-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a unit group and all associated unit conversions. System unit groups
// cannot be deleted.
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

// CreateUnitGroupRequest is a request to create a unit group.
//
// The properties BaseUnitID, Name, Type are required.
type CreateUnitGroupRequestParam struct {
	// Base unit ID.
	BaseUnitID string `json:"base_unit_id" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Unit type.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CreateUnitGroupRequestType `json:"type,omitzero" api:"required"`
	// Notes.
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

// Unit type.
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

// CreateUnitGroupUnitParam contains parameters for an associated unit.
//
// The property UnitID is required.
type CreateUnitGroupUnitParam struct {
	// Unit ID.
	UnitID string `json:"unit_id" api:"required"`
	// Fixed discount amount.
	DiscountFixed param.Opt[float64] `json:"discount_fixed,omitzero"`
	// Discount percentage.
	DiscountPercentage param.Opt[float64] `json:"discount_percentage,omitzero"`
	// Customer portal visibility.
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

// Customer portal visibility.
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

// UnitGroup is a unit group resource.
type UnitGroup struct {
	// Unit group ID.
	ID string `json:"id" api:"required"`
	// List represents a paginated list of resources.
	AssociatedUnits ListUnitGroupUnit `json:"associated_units" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	BaseUnit Unit `json:"base_unit" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "unit_group".
	Object UnitGroupObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Unit type.
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

// Unit type.
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

// UnitGroupUnit is an associated unit within a unit group.
type UnitGroupUnit struct {
	// Unit group unit ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Customer portal visibility.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility UnitGroupUnitCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Fixed discount amount.
	DiscountFixed float64 `json:"discount_fixed" api:"required"`
	// Discount percentage.
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

// Customer portal visibility.
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

// UpdateUnitGroupRequest is a request to partially update a unit group.
type UpdateUnitGroupRequestParam struct {
	// Notes. Set to null to clear.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Base unit ID.
	BaseUnitID param.Opt[string] `json:"base_unit_id,omitzero"`
	// Display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Upserts associated units when provided. Existing units not in the list are
	// preserved.
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
	// CreateUnitGroupRequest is a request to create a unit group.
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
	// UpdateUnitGroupRequest is a request to partially update a unit group.
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
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "base_unit", "associated_units".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by the unit type.
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

// Filter by the unit type.
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
