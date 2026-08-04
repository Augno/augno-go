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

// Creates a demand override, telling the planner about demand the sales history
// cannot see.
//
// The scope reference is validated against the account's items and product lines,
// so an override can never silently match nothing. An `account`-scoped override
// takes no scope reference and must be a delta rather than an absolute value,
// since one number fanned out across every item would flatten the whole plan.
//
// Schedules that have already been generated are unaffected; the override is
// picked up by the next one.
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
// Only the fields sent are changed. The adjustment and value are validated as a
// pair against the resulting override, so switching a stored unit adjustment to
// `delta_percent` is checked as a percentage even when only the adjustment is
// sent; the period is checked the same way.
//
// What an override targets cannot be changed — create a new override to adjust a
// different item, product line, or the account as a whole. Schedules that have
// already been generated are unaffected; the change is picked up by the next one.
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
// spanning a quarter is returned when querying a single month inside it. The `q`
// search term matches the override's note.
//
// This endpoint requires the permission: `demand_overrides:read`.
func (r *OperationDemandOverrideService) List(ctx context.Context, query OperationDemandOverrideListParams, opts ...option.RequestOption) (res *ListDemandOverride, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/demand-overrides"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a demand override permanently.
//
// Schedules that have already been generated are unaffected: each one records the
// overrides it applied, so deleting an override changes only schedules generated
// from now on. To stop an override applying while keeping it on file, deactivate
// it instead.
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
	// - `absolute`: replaces the forecast for each month in the period.
	// - `delta_units`: adds the value to each month in the period.
	// - `delta_percent`: scales each month in the period by the value as a percentage.
	//
	// When several overrides land on the same month they are applied in that order, so
	// a percentage always acts on the already-adjusted number.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Adjustment CreateDemandOverrideRequestAdjustment `json:"adjustment,omitzero" api:"required"`
	// Last day of the demand period the override applies to.
	//
	// Must fall on or after `period_starts_at`.
	PeriodEndsAt time.Time `json:"period_ends_at" api:"required" format:"date-time"`
	// First day of the demand period the override applies to.
	//
	// Overrides are applied month by month, so every calendar month the period touches
	// is adjusted and any time of day is ignored.
	PeriodStartsAt time.Time `json:"period_starts_at" api:"required" format:"date-time"`
	// ID of the item or product line the override targets.
	//
	// Omit it for an `account`-wide override, which targets every planned item rather
	// than one thing. The ID is checked against the account's items and product lines,
	// so an override cannot be created against something that does not exist.
	ScopeRefID string `json:"scope_ref_id" api:"required"`
	// What the override targets.
	//
	//   - `item`: a single item.
	//   - `product_line`: every item sold under one product line.
	//   - `account`: every item in the plan, which is how a blanket assumption such as
	//     "plan for double demand" is expressed.
	//
	// Any of "item", "product_line", "account".
	ScopeType CreateDemandOverrideRequestScopeType `json:"scope_type,omitzero" api:"required"`
	// The amount of the adjustment, interpreted according to `adjustment`.
	//
	// A `delta_percent` value is a number of percent, so `-25` plans a quarter less
	// than the forecast; it cannot go below `-100`. An `absolute` value cannot be
	// negative, while a `delta_units` value can, so that a cancelled program removes
	// demand.
	Value float64 `json:"value" api:"required"`
	// Whether the override is taken into account when a schedule is generated.
	//
	// Send `false` to stage an adjustment that should not affect schedules yet; an
	// override is otherwise created ready to apply.
	Active param.Opt[bool] `json:"active,omitzero"`
	// When the override starts being applied to newly generated schedules.
	//
	// When omitted, the override starts applying straight away.
	EffectiveAt param.Opt[time.Time] `json:"effective_at,omitzero" format:"date-time"`
	// When the override stops being applied to newly generated schedules.
	//
	// When omitted, the override keeps applying until it is deactivated or deleted.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Free-form notes about the adjustment.
	//
	// This is the text the free-text search on the list endpoint matches against.
	Note param.Opt[string] `json:"note,omitzero"`
	// ID of the unit the value is expressed in.
	//
	// Recorded for context only: the value is applied to the planned demand without
	// unit conversion, so a unit adjustment should be stated in the unit the item is
	// planned in.
	UnitID param.Opt[string] `json:"unit_id,omitzero"`
	// Why the adjustment was made.
	//
	// The reason is carried into each schedule the override changes, so a plan can
	// explain why a month departs from history.
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
//
// - `absolute`: replaces the forecast for each month in the period.
// - `delta_units`: adds the value to each month in the period.
// - `delta_percent`: scales each month in the period by the value as a percentage.
//
// When several overrides land on the same month they are applied in that order, so
// a percentage always acts on the already-adjusted number.
type CreateDemandOverrideRequestAdjustment string

const (
	CreateDemandOverrideRequestAdjustmentAbsolute     CreateDemandOverrideRequestAdjustment = "absolute"
	CreateDemandOverrideRequestAdjustmentDeltaUnits   CreateDemandOverrideRequestAdjustment = "delta_units"
	CreateDemandOverrideRequestAdjustmentDeltaPercent CreateDemandOverrideRequestAdjustment = "delta_percent"
)

// What the override targets.
//
//   - `item`: a single item.
//   - `product_line`: every item sold under one product line.
//   - `account`: every item in the plan, which is how a blanket assumption such as
//     "plan for double demand" is expressed.
type CreateDemandOverrideRequestScopeType string

const (
	CreateDemandOverrideRequestScopeTypeItem        CreateDemandOverrideRequestScopeType = "item"
	CreateDemandOverrideRequestScopeTypeProductLine CreateDemandOverrideRequestScopeType = "product_line"
	CreateDemandOverrideRequestScopeTypeAccount     CreateDemandOverrideRequestScopeType = "account"
)

// Why the adjustment was made.
//
// The reason is carried into each schedule the override changes, so a plan can
// explain why a month departs from history.
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

// An adjustment to the demand a production schedule is planned against.
//
// Sales history cannot see a large customer that is about to order, a promotion,
// or a line that is being discontinued. An override is how management tells the
// planner about it. The period names the months the demand will occur in, and only
// months of the coming planning year are adjusted — a period entirely in the past
// changes nothing, because the plan covers the year ahead. `effective_at` and
// `expires_at` answer a different question: how long the override is consulted at
// all, so an adjustment can be retired on a date without deleting it.
type DemandOverride struct {
	// Demand override ID.
	ID string `json:"id" api:"required"`
	// How the value adjusts the forecast.
	//
	// - `absolute`: replaces the forecast for each month in the period.
	// - `delta_units`: adds the value to each month in the period.
	// - `delta_percent`: scales each month in the period by the value as a percentage.
	//
	// When several overrides land on the same month they are applied in that order, so
	// a percentage always acts on the already-adjusted number. An adjusted month is
	// never taken below zero.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Adjustment DemandOverrideAdjustment `json:"adjustment" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	CreatedBy Actor `json:"created_by" api:"required"`
	// When the override starts being applied to newly generated schedules.
	EffectiveAt time.Time `json:"effective_at" api:"required" format:"date-time"`
	// When the override stops being applied to newly generated schedules.
	//
	// An override with no expiry keeps applying until it is deactivated or deleted.
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
	//
	// Overrides are applied month by month, so every calendar month the period touches
	// is adjusted and any time of day is ignored.
	PeriodStartsAt time.Time `json:"period_starts_at" api:"required" format:"date-time"`
	// Why the adjustment was made.
	//
	// The reason is carried into each schedule the override changes, so a plan can
	// explain why a month departs from history.
	//
	// Any of "new_customer", "lost_account", "promotion", "seasonal_shift",
	// "new_product", "discontinued", "market_intelligence", "other".
	Reason DemandOverrideReason `json:"reason" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Scope Entity `json:"scope" api:"required"`
	// What the override targets.
	//
	//   - `item`: a single item.
	//   - `product_line`: every item sold under one product line.
	//   - `account`: every item in the plan, which is how a blanket assumption such as
	//     "plan for double demand" is expressed.
	//
	// Any of "item", "product_line", "account".
	ScopeType DemandOverrideScopeType `json:"scope_type" api:"required"`
	// Whether the override is taken into account when a schedule is generated.
	//
	// An inactive override is skipped whatever its effective window says, which is how
	// a prepared adjustment is parked without losing it.
	//
	// Any of "active", "inactive".
	Status DemandOverrideStatus `json:"status" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	Unit Unit `json:"unit" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The amount of the adjustment, interpreted according to `adjustment`.
	//
	// A `delta_percent` value is a number of percent, so `-25` plans a quarter less
	// than the forecast.
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
//
// - `absolute`: replaces the forecast for each month in the period.
// - `delta_units`: adds the value to each month in the period.
// - `delta_percent`: scales each month in the period by the value as a percentage.
//
// When several overrides land on the same month they are applied in that order, so
// a percentage always acts on the already-adjusted number. An adjusted month is
// never taken below zero.
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
//
// The reason is carried into each schedule the override changes, so a plan can
// explain why a month departs from history.
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

// What the override targets.
//
//   - `item`: a single item.
//   - `product_line`: every item sold under one product line.
//   - `account`: every item in the plan, which is how a blanket assumption such as
//     "plan for double demand" is expressed.
type DemandOverrideScopeType string

const (
	DemandOverrideScopeTypeItem        DemandOverrideScopeType = "item"
	DemandOverrideScopeTypeProductLine DemandOverrideScopeType = "product_line"
	DemandOverrideScopeTypeAccount     DemandOverrideScopeType = "account"
)

// Whether the override is taken into account when a schedule is generated.
//
// An inactive override is skipped whatever its effective window says, which is how
// a prepared adjustment is parked without losing it.
type DemandOverrideStatus string

const (
	DemandOverrideStatusActive   DemandOverrideStatus = "active"
	DemandOverrideStatusInactive DemandOverrideStatus = "inactive"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDemandOverride struct {
	// Resources in this page.
	Data []DemandOverride `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDemandOverrideObject `json:"object" api:"required"`
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
	// When the override stops being applied to newly generated schedules.
	//
	// Clear it to keep the override applying until it is deactivated or deleted.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Free-form notes about the adjustment.
	Note param.Opt[string] `json:"note,omitzero"`
	// ID of the unit the value is expressed in.
	//
	// Recorded for context only: the value is applied to the planned demand without
	// unit conversion.
	UnitID param.Opt[string] `json:"unit_id,omitzero"`
	// Whether the override is taken into account when a schedule is generated.
	//
	// Deactivating parks the override without losing it; it is skipped whatever its
	// effective window says, and can be reactivated later.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Last day of the demand period the override applies to.
	//
	// Must fall on or after the override's start, whether that is sent here or already
	// stored.
	PeriodEndsAt param.Opt[time.Time] `json:"period_ends_at,omitzero" format:"date-time"`
	// First day of the demand period the override applies to.
	//
	// Overrides are applied month by month, so every calendar month the period touches
	// is adjusted and any time of day is ignored.
	PeriodStartsAt param.Opt[time.Time] `json:"period_starts_at,omitzero" format:"date-time"`
	// The amount of the adjustment, interpreted according to `adjustment`.
	//
	// It is validated against the adjustment the override ends up with, so switching a
	// stored unit delta to `delta_percent` without sending a new value requires the
	// existing value to be a legal percentage.
	Value param.Opt[float64] `json:"value,omitzero"`
	// Why the adjustment was made.
	//
	// The reason is carried into each schedule the override changes, so a plan can
	// explain why a month departs from history.
	//
	// Any of "new_customer", "lost_account", "promotion", "seasonal_shift",
	// "new_product", "discontinued", "market_intelligence", "other".
	Reason UpdateDemandOverrideRequestReason `json:"reason,omitzero"`
	// How the value adjusts the forecast.
	//
	// - `absolute`: replaces the forecast for each month in the period.
	// - `delta_units`: adds the value to each month in the period.
	// - `delta_percent`: scales each month in the period by the value as a percentage.
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
//
// - `absolute`: replaces the forecast for each month in the period.
// - `delta_units`: adds the value to each month in the period.
// - `delta_percent`: scales each month in the period by the value as a percentage.
type UpdateDemandOverrideRequestAdjustment string

const (
	UpdateDemandOverrideRequestAdjustmentAbsolute     UpdateDemandOverrideRequestAdjustment = "absolute"
	UpdateDemandOverrideRequestAdjustmentDeltaUnits   UpdateDemandOverrideRequestAdjustment = "delta_units"
	UpdateDemandOverrideRequestAdjustmentDeltaPercent UpdateDemandOverrideRequestAdjustment = "delta_percent"
)

// Why the adjustment was made.
//
// The reason is carried into each schedule the override changes, so a plan can
// explain why a month departs from history.
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
	// Only return overrides with these kinds of target.
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
