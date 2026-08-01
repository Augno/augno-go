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

// Adjust the demand a production schedule plans against. Overrides are how
// management accounts for demand that sales history cannot see.
//
// OperationDemandOverrideService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationDemandOverrideService] method instead.
type OperationDemandOverrideService struct {
	options []option.RequestOption
}

// NewOperationDemandOverrideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationDemandOverrideService(opts ...option.RequestOption) (r OperationDemandOverrideService) {
	r = OperationDemandOverrideService{}
	r.options = opts
	return
}

// Creates a demand override.
//
// The scope reference is validated against the account's items or product lines,
// so an override can never silently match nothing. An `account`-scoped override
// applies to every planned item and takes no scope reference; it must be a delta,
// not an absolute value. An `absolute` value replaces the forecast for the period,
// `delta_units` adds to it, and `delta_percent` scales it; a percent override
// cannot reduce demand by more than 100%.
//
// This endpoint requires the permission: `demand_overrides:create`.
func (r *OperationDemandOverrideService) New(ctx context.Context, params OperationDemandOverrideNewParams, opts ...option.RequestOption) (res *DemandOverride, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/demand-overrides"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a single demand override by ID.
//
// This endpoint requires the permission: `demand_overrides:read`.
func (r *OperationDemandOverrideService) Get(ctx context.Context, id string, query OperationDemandOverrideGetParams, opts ...option.RequestOption) (res *DemandOverride, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/demand-overrides/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a demand override.
//
// The type and value are validated as a pair against the resulting override, so
// switching an existing units adjustment to `delta_percent` is checked as a
// percent even when only the type is sent.
//
// This endpoint requires the permission: `demand_overrides:update`.
func (r *OperationDemandOverrideService) Update(ctx context.Context, id string, params OperationDemandOverrideUpdateParams, opts ...option.RequestOption) (res *DemandOverride, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/demand-overrides/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of demand overrides, most recently created first.
//
// The period filters match on overlap rather than containment, so an override
// spanning a quarter is returned when querying a single month inside it.
//
// This endpoint requires the permission: `demand_overrides:read`.
func (r *OperationDemandOverrideService) List(ctx context.Context, query OperationDemandOverrideListParams, opts ...option.RequestOption) (res *ListDemandOverride, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/demand-overrides"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a demand override.
//
// Schedules already generated are unaffected: a version snapshots the overrides it
// applied, so deleting one changes future solves only.
//
// This endpoint requires the permission: `demand_overrides:delete`.
func (r *OperationDemandOverrideService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationDemandOverrideDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/demand-overrides/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a demand override.
//
// The properties Adjustment, PeriodEndsAt, PeriodStartsAt, ScopeRefID, ScopeType,
// Value are required.
type CreateDemandOverrideRequestParam struct {
	// How the value adjusts the forecast.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Adjustment CreateDemandOverrideRequestAdjustment `json:"adjustment,omitzero" api:"required"`
	// Last day of the demand period the override applies to.
	PeriodEndsAt time.Time `json:"period_ends_at" api:"required" format:"date-time"`
	// First day of the demand period the override applies to.
	PeriodStartsAt time.Time `json:"period_starts_at" api:"required" format:"date-time"`
	// ID of the item or product line the override targets. Omit for an `account`-wide
	// override, which targets every planned item.
	ScopeRefID string `json:"scope_ref_id" api:"required"`
	// What the override targets.
	//
	// Any of "item", "product_line", "account".
	ScopeType CreateDemandOverrideRequestScopeType `json:"scope_type,omitzero" api:"required"`
	// The adjustment, interpreted according to `adjustment`.
	Value float64 `json:"value" api:"required"`
	// Whether the override is applied to solves at all. Defaults to true.
	Active param.Opt[bool] `json:"active,omitzero"`
	// When the override starts being applied to solves. Defaults to now.
	EffectiveAt param.Opt[time.Time] `json:"effective_at,omitzero" format:"date-time"`
	// When the override stops being applied to solves. Omit for an override with no
	// end.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Free-form notes about the adjustment.
	Note param.Opt[string] `json:"note,omitzero"`
	// ID of the unit the value is expressed in.
	UnitID param.Opt[string] `json:"unit_id,omitzero"`
	// Why the adjustment was made.
	//
	// Any of "new_customer", "lost_account", "promotion", "seasonal_shift",
	// "new_product", "discontinued", "market_intelligence", "other".
	Reason CreateDemandOverrideRequestReason `json:"reason,omitzero"`
	paramObj
}

func (r CreateDemandOverrideRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateDemandOverrideRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateDemandOverrideRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the value adjusts the forecast.
type CreateDemandOverrideRequestAdjustment string

const (
	CreateDemandOverrideRequestAdjustmentAbsolute     CreateDemandOverrideRequestAdjustment = "absolute"
	CreateDemandOverrideRequestAdjustmentDeltaUnits   CreateDemandOverrideRequestAdjustment = "delta_units"
	CreateDemandOverrideRequestAdjustmentDeltaPercent CreateDemandOverrideRequestAdjustment = "delta_percent"
)

// What the override targets.
type CreateDemandOverrideRequestScopeType string

const (
	CreateDemandOverrideRequestScopeTypeItem        CreateDemandOverrideRequestScopeType = "item"
	CreateDemandOverrideRequestScopeTypeProductLine CreateDemandOverrideRequestScopeType = "product_line"
	CreateDemandOverrideRequestScopeTypeAccount     CreateDemandOverrideRequestScopeType = "account"
)

// Why the adjustment was made.
type CreateDemandOverrideRequestReason string

const (
	CreateDemandOverrideRequestReasonNewCustomer        CreateDemandOverrideRequestReason = "new_customer"
	CreateDemandOverrideRequestReasonLostAccount        CreateDemandOverrideRequestReason = "lost_account"
	CreateDemandOverrideRequestReasonPromotion          CreateDemandOverrideRequestReason = "promotion"
	CreateDemandOverrideRequestReasonSeasonalShift      CreateDemandOverrideRequestReason = "seasonal_shift"
	CreateDemandOverrideRequestReasonNewProduct         CreateDemandOverrideRequestReason = "new_product"
	CreateDemandOverrideRequestReasonDiscontinued       CreateDemandOverrideRequestReason = "discontinued"
	CreateDemandOverrideRequestReasonMarketIntelligence CreateDemandOverrideRequestReason = "market_intelligence"
	CreateDemandOverrideRequestReasonOther              CreateDemandOverrideRequestReason = "other"
)

// An adjustment to the demand a production schedule plans against.
//
// Sales history cannot see a large customer that is about to order, a promotion,
// or a line that is being discontinued. An override is how management tells the
// planner about it. The period bounds the demand months the adjustment applies to;
// `effective_from` and `expires_at` bound when the override is consulted at all,
// which is a different question — an override for next quarter typically stops
// applying once the real orders arrive.
//
// A product-line override applies to each of the line's items; an account-wide
// override applies to every planned item, which is how a global growth assumption
// (e.g. "plan for double demand") is expressed.
type DemandOverride struct {
	// Demand override ID.
	ID string `json:"id" api:"required"`
	// How the value adjusts the forecast.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Adjustment DemandOverrideAdjustment `json:"adjustment" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	CreatedBy Actor `json:"created_by" api:"required"`
	// When the override starts being applied to solves.
	EffectiveAt time.Time `json:"effective_at" api:"required" format:"date-time"`
	// When the override stops being applied to solves.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Free-form notes about the adjustment.
	Note string `json:"note" api:"required"`
	// Resource type identifier.
	//
	// Any of "demand_override".
	Object DemandOverrideObject `json:"object" api:"required"`
	// Last day of the demand period the override applies to.
	PeriodEndsAt time.Time `json:"period_ends_at" api:"required" format:"date-time"`
	// First day of the demand period the override applies to.
	PeriodStartsAt time.Time `json:"period_starts_at" api:"required" format:"date-time"`
	// Why the adjustment was made.
	//
	// Any of "new_customer", "lost_account", "promotion", "seasonal_shift",
	// "new_product", "discontinued", "market_intelligence", "other".
	Reason DemandOverrideReason `json:"reason" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Scope Entity `json:"scope" api:"required"`
	// What kind of resource the override targets. Mirrors `scope.type`, which is only
	// present when the scope is expanded.
	//
	// Any of "item", "product_line", "account".
	ScopeType DemandOverrideScopeType `json:"scope_type" api:"required"`
	// Whether the override is applied to solves at all.
	//
	// Any of "active", "inactive".
	Status DemandOverrideStatus `json:"status" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	Unit Unit `json:"unit" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The adjustment, interpreted according to `adjustment`.
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Adjustment     respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		EffectiveAt    respjson.Field
		ExpiresAt      respjson.Field
		Note           respjson.Field
		Object         respjson.Field
		PeriodEndsAt   respjson.Field
		PeriodStartsAt respjson.Field
		Reason         respjson.Field
		Scope          respjson.Field
		ScopeType      respjson.Field
		Status         respjson.Field
		Unit           respjson.Field
		UpdatedAt      respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DemandOverride) RawJSON() string { return r.JSON.raw }
func (r *DemandOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the value adjusts the forecast.
type DemandOverrideAdjustment string

const (
	DemandOverrideAdjustmentAbsolute     DemandOverrideAdjustment = "absolute"
	DemandOverrideAdjustmentDeltaUnits   DemandOverrideAdjustment = "delta_units"
	DemandOverrideAdjustmentDeltaPercent DemandOverrideAdjustment = "delta_percent"
)

// Resource type identifier.
type DemandOverrideObject string

const (
	DemandOverrideObjectDemandOverride DemandOverrideObject = "demand_override"
)

// Why the adjustment was made.
type DemandOverrideReason string

const (
	DemandOverrideReasonNewCustomer        DemandOverrideReason = "new_customer"
	DemandOverrideReasonLostAccount        DemandOverrideReason = "lost_account"
	DemandOverrideReasonPromotion          DemandOverrideReason = "promotion"
	DemandOverrideReasonSeasonalShift      DemandOverrideReason = "seasonal_shift"
	DemandOverrideReasonNewProduct         DemandOverrideReason = "new_product"
	DemandOverrideReasonDiscontinued       DemandOverrideReason = "discontinued"
	DemandOverrideReasonMarketIntelligence DemandOverrideReason = "market_intelligence"
	DemandOverrideReasonOther              DemandOverrideReason = "other"
)

// What kind of resource the override targets. Mirrors `scope.type`, which is only
// present when the scope is expanded.
type DemandOverrideScopeType string

const (
	DemandOverrideScopeTypeItem        DemandOverrideScopeType = "item"
	DemandOverrideScopeTypeProductLine DemandOverrideScopeType = "product_line"
	DemandOverrideScopeTypeAccount     DemandOverrideScopeType = "account"
)

// Whether the override is applied to solves at all.
type DemandOverrideStatus string

const (
	DemandOverrideStatusActive   DemandOverrideStatus = "active"
	DemandOverrideStatusInactive DemandOverrideStatus = "inactive"
)

// List represents a paginated list of resources.
type ListDemandOverride struct {
	// Resources in this page.
	Data []DemandOverride `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDemandOverrideObject `json:"object" api:"required"`
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
func (r ListDemandOverride) RawJSON() string { return r.JSON.raw }
func (r *ListDemandOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDemandOverrideObject string

const (
	ListDemandOverrideObjectList ListDemandOverrideObject = "list"
)

// Request to update a demand override.
type UpdateDemandOverrideRequestParam struct {
	// When the override stops being applied to solves. Clear it to make the override
	// permanent.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Free-form notes about the adjustment.
	Note param.Opt[string] `json:"note,omitzero"`
	// ID of the unit the value is expressed in.
	UnitID param.Opt[string] `json:"unit_id,omitzero"`
	// Whether the override is applied to solves at all.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Last day of the demand period the override applies to.
	PeriodEndsAt param.Opt[time.Time] `json:"period_ends_at,omitzero" format:"date-time"`
	// First day of the demand period the override applies to.
	PeriodStartsAt param.Opt[time.Time] `json:"period_starts_at,omitzero" format:"date-time"`
	// The adjustment, interpreted according to `adjustment`.
	Value param.Opt[float64] `json:"value,omitzero"`
	// Why the adjustment was made.
	//
	// Any of "new_customer", "lost_account", "promotion", "seasonal_shift",
	// "new_product", "discontinued", "market_intelligence", "other".
	Reason UpdateDemandOverrideRequestReason `json:"reason,omitzero"`
	// How the value adjusts the forecast.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Adjustment UpdateDemandOverrideRequestAdjustment `json:"adjustment,omitzero"`
	paramObj
}

func (r UpdateDemandOverrideRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateDemandOverrideRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateDemandOverrideRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the value adjusts the forecast.
type UpdateDemandOverrideRequestAdjustment string

const (
	UpdateDemandOverrideRequestAdjustmentAbsolute     UpdateDemandOverrideRequestAdjustment = "absolute"
	UpdateDemandOverrideRequestAdjustmentDeltaUnits   UpdateDemandOverrideRequestAdjustment = "delta_units"
	UpdateDemandOverrideRequestAdjustmentDeltaPercent UpdateDemandOverrideRequestAdjustment = "delta_percent"
)

// Why the adjustment was made.
type UpdateDemandOverrideRequestReason string

const (
	UpdateDemandOverrideRequestReasonNewCustomer        UpdateDemandOverrideRequestReason = "new_customer"
	UpdateDemandOverrideRequestReasonLostAccount        UpdateDemandOverrideRequestReason = "lost_account"
	UpdateDemandOverrideRequestReasonPromotion          UpdateDemandOverrideRequestReason = "promotion"
	UpdateDemandOverrideRequestReasonSeasonalShift      UpdateDemandOverrideRequestReason = "seasonal_shift"
	UpdateDemandOverrideRequestReasonNewProduct         UpdateDemandOverrideRequestReason = "new_product"
	UpdateDemandOverrideRequestReasonDiscontinued       UpdateDemandOverrideRequestReason = "discontinued"
	UpdateDemandOverrideRequestReasonMarketIntelligence UpdateDemandOverrideRequestReason = "market_intelligence"
	UpdateDemandOverrideRequestReasonOther              UpdateDemandOverrideRequestReason = "other"
)

type OperationDemandOverrideDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationDemandOverrideDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationDemandOverrideDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationDemandOverrideNewParams struct {
	// Request to create a demand override.
	CreateDemandOverrideRequest CreateDemandOverrideRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "scope".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationDemandOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDemandOverrideRequest)
}
func (r *OperationDemandOverrideNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationDemandOverrideNewParams]'s query parameters as
// `url.Values`.
func (r OperationDemandOverrideNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationDemandOverrideGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "scope".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationDemandOverrideGetParams]'s query parameters as
// `url.Values`.
func (r OperationDemandOverrideGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationDemandOverrideUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "scope".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a demand override.
	UpdateDemandOverrideRequest UpdateDemandOverrideRequestParam
	paramObj
}

func (r OperationDemandOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateDemandOverrideRequest)
}
func (r *OperationDemandOverrideUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationDemandOverrideUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationDemandOverrideUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationDemandOverrideListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return overrides whose period starts on or before this timestamp, formatted
	// as RFC3339.
	PeriodEnd param.Opt[string] `query:"period_end,omitzero" json:"-"`
	// Only return overrides whose period ends on or after this timestamp, formatted as
	// RFC3339.
	PeriodStart param.Opt[string] `query:"period_start,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Only return overrides making these kinds of adjustment.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Adjustments []string `query:"adjustments,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "scope".
	Include []string `query:"include,omitzero" json:"-"`
	// Only return overrides targeting these items or product lines.
	ScopeRefIDs []string `query:"scope_ref_ids,omitzero" json:"-"`
	// Only return overrides targeting these kinds of resource.
	//
	// Any of "item", "product_line", "account".
	ScopeTypes []string `query:"scope_types,omitzero" json:"-"`
	// Only return overrides in these activation states.
	//
	// Any of "active", "inactive".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationDemandOverrideListParams]'s query parameters as
// `url.Values`.
func (r OperationDemandOverrideListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
