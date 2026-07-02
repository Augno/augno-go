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

// List and manage unit groups and their associated units.
//
// CatalogUnitGroupUnitService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogUnitGroupUnitService] method instead.
type CatalogUnitGroupUnitService struct {
	options []option.RequestOption
}

// NewCatalogUnitGroupUnitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogUnitGroupUnitService(opts ...option.RequestOption) (r CatalogUnitGroupUnitService) {
	r = CatalogUnitGroupUnitService{}
	r.options = opts
	return
}

// Adds a unit to a unit group. If the unit is already in the group, its existing
// association is updated with the provided settings instead.
//
// This endpoint requires the permission: `unit_groups:update`.
func (r *CatalogUnitGroupUnitService) New(ctx context.Context, unitGroupID string, params CatalogUnitGroupUnitNewParams, opts ...option.RequestOption) (res *UnitGroupUnit, err error) {
	opts = slices.Concat(r.options, opts)
	if unitGroupID == "" {
		err = errors.New("missing required unit_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s/units", url.PathEscape(unitGroupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns an associated unit within a unit group by ID.
//
// This endpoint requires the permission: `unit_groups:read`.
func (r *CatalogUnitGroupUnitService) Get(ctx context.Context, id string, params CatalogUnitGroupUnitGetParams, opts ...option.RequestOption) (res *UnitGroupUnit, err error) {
	opts = slices.Concat(r.options, opts)
	if params.UnitGroupID == "" {
		err = errors.New("missing required unit_group_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s/units/%s", url.PathEscape(params.UnitGroupID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Partially updates an associated unit within a unit group.
//
// This endpoint requires the permission: `unit_groups:update`.
func (r *CatalogUnitGroupUnitService) Update(ctx context.Context, id string, params CatalogUnitGroupUnitUpdateParams, opts ...option.RequestOption) (res *UnitGroupUnit, err error) {
	opts = slices.Concat(r.options, opts)
	if params.UnitGroupID == "" {
		err = errors.New("missing required unit_group_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s/units/%s", url.PathEscape(params.UnitGroupID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a list of associated units within a unit group.
//
// This endpoint requires the permission: `unit_groups:read`.
func (r *CatalogUnitGroupUnitService) List(ctx context.Context, unitGroupID string, query CatalogUnitGroupUnitListParams, opts ...option.RequestOption) (res *ListUnitGroupUnit, err error) {
	opts = slices.Concat(r.options, opts)
	if unitGroupID == "" {
		err = errors.New("missing required unit_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s/units", url.PathEscape(unitGroupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes a unit from a unit group. The unit itself is not deleted.
//
// This endpoint requires the permission: `unit_groups:delete`.
func (r *CatalogUnitGroupUnitService) Delete(ctx context.Context, id string, body CatalogUnitGroupUnitDeleteParams, opts ...option.RequestOption) (res *CatalogUnitGroupUnitDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.UnitGroupID == "" {
		err = errors.New("missing required unit_group_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/unit-groups/%s/units/%s", url.PathEscape(body.UnitGroupID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to add a unit to a unit group.
//
// The property UnitID is required.
type CreateUnitGroupUnitRequestParam struct {
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
	CustomerPortalVisibility CreateUnitGroupUnitRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r CreateUnitGroupUnitRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateUnitGroupUnitRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateUnitGroupUnitRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the unit is shown to customers in the customer portal.
type CreateUnitGroupUnitRequestCustomerPortalVisibility string

const (
	CreateUnitGroupUnitRequestCustomerPortalVisibilityVisible CreateUnitGroupUnitRequestCustomerPortalVisibility = "visible"
	CreateUnitGroupUnitRequestCustomerPortalVisibilityHidden  CreateUnitGroupUnitRequestCustomerPortalVisibility = "hidden"
)

// Request to partially update an associated unit within a unit group.
type UpdateUnitGroupUnitRequestParam struct {
	// Flat amount subtracted from the unit's price when an order is placed in this
	// unit.
	DiscountFixed param.Opt[float64] `json:"discount_fixed,omitzero"`
	// Percentage discount applied to the unit's price when an order is placed in this
	// unit (e.g. `10` is a 10% discount).
	DiscountPercentage param.Opt[float64] `json:"discount_percentage,omitzero"`
	// ID of the unit this association refers to.
	//
	// The unit's dimension must match the group's `type`.
	UnitID param.Opt[string] `json:"unit_id,omitzero"`
	// Whether the unit is shown to customers in the customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility UpdateUnitGroupUnitRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r UpdateUnitGroupUnitRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateUnitGroupUnitRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateUnitGroupUnitRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the unit is shown to customers in the customer portal.
type UpdateUnitGroupUnitRequestCustomerPortalVisibility string

const (
	UpdateUnitGroupUnitRequestCustomerPortalVisibilityVisible UpdateUnitGroupUnitRequestCustomerPortalVisibility = "visible"
	UpdateUnitGroupUnitRequestCustomerPortalVisibilityHidden  UpdateUnitGroupUnitRequestCustomerPortalVisibility = "hidden"
)

type CatalogUnitGroupUnitDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogUnitGroupUnitDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogUnitGroupUnitDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogUnitGroupUnitNewParams struct {
	// Request to add a unit to a unit group.
	CreateUnitGroupUnitRequest CreateUnitGroupUnitRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogUnitGroupUnitNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateUnitGroupUnitRequest)
}
func (r *CatalogUnitGroupUnitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitGroupUnitNewParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupUnitNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGroupUnitGetParams struct {
	UnitGroupID string `path:"unit_group_id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitGroupUnitGetParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupUnitGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGroupUnitUpdateParams struct {
	UnitGroupID string `path:"unit_group_id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update an associated unit within a unit group.
	UpdateUnitGroupUnitRequest UpdateUnitGroupUnitRequestParam
	paramObj
}

func (r CatalogUnitGroupUnitUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateUnitGroupUnitRequest)
}
func (r *CatalogUnitGroupUnitUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitGroupUnitUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupUnitUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGroupUnitListParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitGroupUnitListParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitGroupUnitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGroupUnitDeleteParams struct {
	UnitGroupID string `path:"unit_group_id" api:"required" json:"-"`
	paramObj
}
