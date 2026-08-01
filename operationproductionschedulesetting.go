// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/augno/augno-go/internal/apijson"
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// The planning assumptions production schedules are solved against, and the
// per-resource overrides that mark which machines constrain the plan.
//
// OperationProductionScheduleSettingService contains methods and other services
// that help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationProductionScheduleSettingService] method instead.
type OperationProductionScheduleSettingService struct {
	options []option.RequestOption
	// The planning assumptions production schedules are solved against, and the
	// per-resource overrides that mark which machines constrain the plan.
	Resources OperationProductionScheduleSettingResourceService
}

// NewOperationProductionScheduleSettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOperationProductionScheduleSettingService(opts ...option.RequestOption) (r OperationProductionScheduleSettingService) {
	r = OperationProductionScheduleSettingService{}
	r.options = opts
	r.Resources = NewOperationProductionScheduleSettingResourceService(opts...)
	return
}

// Replaces the planning assumptions production schedules are solved against.
//
// Settings are replaced wholesale rather than patched, because they are read as
// one coherent set: a horizon that no longer matches the frozen window, or a
// capacity headroom that no longer matches the shift pattern, would produce a plan
// nobody intended.
//
// Existing schedule versions are unaffected — each one snapshots the assumptions
// it was solved under, so changing settings changes future plans only.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleSettingService) Update(ctx context.Context, body OperationProductionScheduleSettingUpdateParams, opts ...option.RequestOption) (res *ProductionScheduleSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedule-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the planning assumptions production schedules are solved against.
//
// Always fully populated: an account that has never saved settings gets the
// solver's own defaults rather than nulls, so a caller never has to know which
// values would otherwise be assumed. `settings_status` distinguishes the two.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ProductionScheduleSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedule-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The planning assumptions a production schedule is solved against.
//
// Every value here was a hardcoded constant in the scheduling script this feature
// replaced. The resource is always fully populated: an account that has never
// saved settings gets the solver's own defaults, so a caller never has to know
// which values would otherwise be assumed. `settings_status` says which of the two
// it is looking at.
type ProductionScheduleSettings struct {
	// Whether a generated version is published automatically.
	//
	// Any of "active", "inactive".
	AutoPublishStatus ProductionScheduleSettingsAutoPublishStatus `json:"auto_publish_status" api:"required"`
	// Whether schedules are generated on a timer.
	//
	// Any of "active", "inactive".
	CadenceStatus ProductionScheduleSettingsCadenceStatus `json:"cadence_status" api:"required"`
	// Share of machine time a plan may fill. The remainder absorbs changeovers, which
	// are not scheduled as explicit blocks.
	CapacityHeadroomPct float64 `json:"capacity_headroom_pct" api:"required"`
	// Typical changeover duration, used to calibrate the changeover model.
	ChangeoverAvgMinutes float64 `json:"changeover_avg_minutes" api:"required"`
	// Hourly labour rate charged to a changeover.
	ChangeoverLaborRate float64 `json:"changeover_labor_rate" api:"required"`
	// Longest plausible changeover.
	ChangeoverMaxMinutes float64 `json:"changeover_max_minutes" api:"required"`
	// Shortest plausible changeover.
	ChangeoverMinMinutes float64 `json:"changeover_min_minutes" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ConstraintDepartment Entity `json:"constraint_department" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Default weeks of lead time at the constraint when an item has no measurement.
	DefaultConstraintLeadTimeWeeks float64 `json:"default_constraint_lead_time_weeks" api:"required"`
	// Units in a default production lot.
	DefaultLotUnits float64 `json:"default_lot_units" api:"required"`
	// How demand is derived.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis ProductionScheduleSettingsDemandBasis `json:"demand_basis" api:"required"`
	// Months of order history the demand baseline is drawn from.
	DemandWindowMonths int64 `json:"demand_window_months" api:"required"`
	// Weeks between finishing at the constraint and being sellable.
	FinishLeadTimeWeeks float64 `json:"finish_lead_time_weeks" api:"required"`
	// Months of history the forecast is fitted to.
	ForecastHistoryMonths int64 `json:"forecast_history_months" api:"required"`
	// Months the forecast projects forward.
	ForecastMonths int64 `json:"forecast_months" api:"required"`
	// Z-score applied to forecast variability.
	ForecastZ float64 `json:"forecast_z" api:"required"`
	// How many leading weeks become a commitment when a version is published.
	FrozenWeeks int64 `json:"frozen_weeks" api:"required"`
	// Cron expression driving the generation cadence.
	GenerationCron string `json:"generation_cron" api:"required"`
	// Timezone the cadence is interpreted in.
	GenerationTimezone string `json:"generation_timezone" api:"required"`
	// Annual cost of holding stock, as a share of item value.
	HoldingRatePct float64 `json:"holding_rate_pct" api:"required"`
	// Hours in a shift.
	HoursPerShift float64 `json:"hours_per_shift" api:"required"`
	// When the cadence last fired.
	LastGeneratedAt time.Time `json:"last_generated_at" api:"required" format:"date-time"`
	// How many steps downstream department work is derived for.
	MaxFlowDepth int64 `json:"max_flow_depth" api:"required"`
	// Ceiling on how far ahead any item is built.
	MaxWeeksSupply float64 `json:"max_weeks_supply" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_settings".
	Object ProductionScheduleSettingsObject `json:"object" api:"required"`
	// How many weeks a generated plan covers.
	PlanningHorizonWeeks int64 `json:"planning_horizon_weeks" api:"required"`
	// Z-score for the service level safety stock targets.
	ServiceLevelZ float64 `json:"service_level_z" api:"required"`
	// Whether these are the merchant's saved values or the solver's defaults.
	//
	// Any of "stored", "default".
	SettingsStatus ProductionScheduleSettingsSettingsStatus `json:"settings_status" api:"required"`
	// Shifts worked per day.
	ShiftsPerDay int64 `json:"shifts_per_day" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Day a planning week starts, where 0 is Sunday.
	WeekStartDay int64 `json:"week_start_day" api:"required"`
	// Weeks worked per year.
	WeeksPerYear int64 `json:"weeks_per_year" api:"required"`
	// Days worked per week.
	WorkDaysPerWeek int64 `json:"work_days_per_week" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoPublishStatus              respjson.Field
		CadenceStatus                  respjson.Field
		CapacityHeadroomPct            respjson.Field
		ChangeoverAvgMinutes           respjson.Field
		ChangeoverLaborRate            respjson.Field
		ChangeoverMaxMinutes           respjson.Field
		ChangeoverMinMinutes           respjson.Field
		ConstraintDepartment           respjson.Field
		CreatedAt                      respjson.Field
		DefaultConstraintLeadTimeWeeks respjson.Field
		DefaultLotUnits                respjson.Field
		DemandBasis                    respjson.Field
		DemandWindowMonths             respjson.Field
		FinishLeadTimeWeeks            respjson.Field
		ForecastHistoryMonths          respjson.Field
		ForecastMonths                 respjson.Field
		ForecastZ                      respjson.Field
		FrozenWeeks                    respjson.Field
		GenerationCron                 respjson.Field
		GenerationTimezone             respjson.Field
		HoldingRatePct                 respjson.Field
		HoursPerShift                  respjson.Field
		LastGeneratedAt                respjson.Field
		MaxFlowDepth                   respjson.Field
		MaxWeeksSupply                 respjson.Field
		Object                         respjson.Field
		PlanningHorizonWeeks           respjson.Field
		ServiceLevelZ                  respjson.Field
		SettingsStatus                 respjson.Field
		ShiftsPerDay                   respjson.Field
		UpdatedAt                      respjson.Field
		WeekStartDay                   respjson.Field
		WeeksPerYear                   respjson.Field
		WorkDaysPerWeek                respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleSettings) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether a generated version is published automatically.
type ProductionScheduleSettingsAutoPublishStatus string

const (
	ProductionScheduleSettingsAutoPublishStatusActive   ProductionScheduleSettingsAutoPublishStatus = "active"
	ProductionScheduleSettingsAutoPublishStatusInactive ProductionScheduleSettingsAutoPublishStatus = "inactive"
)

// Whether schedules are generated on a timer.
type ProductionScheduleSettingsCadenceStatus string

const (
	ProductionScheduleSettingsCadenceStatusActive   ProductionScheduleSettingsCadenceStatus = "active"
	ProductionScheduleSettingsCadenceStatusInactive ProductionScheduleSettingsCadenceStatus = "inactive"
)

// How demand is derived.
type ProductionScheduleSettingsDemandBasis string

const (
	ProductionScheduleSettingsDemandBasisTrailing12  ProductionScheduleSettingsDemandBasis = "trailing_12"
	ProductionScheduleSettingsDemandBasisSeasonalEma ProductionScheduleSettingsDemandBasis = "seasonal_ema"
)

// Resource type identifier.
type ProductionScheduleSettingsObject string

const (
	ProductionScheduleSettingsObjectProductionScheduleSettings ProductionScheduleSettingsObject = "production_schedule_settings"
)

// Whether these are the merchant's saved values or the solver's defaults.
type ProductionScheduleSettingsSettingsStatus string

const (
	ProductionScheduleSettingsSettingsStatusStored  ProductionScheduleSettingsSettingsStatus = "stored"
	ProductionScheduleSettingsSettingsStatusDefault ProductionScheduleSettingsSettingsStatus = "default"
)

// Request to replace the account's planning assumptions.
//
// The properties AutoPublishStatus, CadenceStatus, CapacityHeadroomPct,
// ChangeoverAvgMinutes, ChangeoverLaborRate, ChangeoverMaxMinutes,
// ChangeoverMinMinutes, DefaultConstraintLeadTimeWeeks, DefaultLotUnits,
// DemandBasis, DemandWindowMonths, FinishLeadTimeWeeks, ForecastHistoryMonths,
// ForecastMonths, ForecastZ, FrozenWeeks, GenerationTimezone, HoldingRatePct,
// HoursPerShift, MaxFlowDepth, MaxWeeksSupply, PlanningHorizonWeeks,
// ServiceLevelZ, ShiftsPerDay, WeekStartDay, WeeksPerYear, WorkDaysPerWeek are
// required.
type UpdateProductionScheduleSettingsRequestParam struct {
	// Whether a generated version is published automatically.
	//
	// Any of "active", "inactive".
	AutoPublishStatus UpdateProductionScheduleSettingsRequestAutoPublishStatus `json:"auto_publish_status,omitzero" api:"required"`
	// Whether schedules are generated on a timer.
	//
	// Any of "active", "inactive".
	CadenceStatus UpdateProductionScheduleSettingsRequestCadenceStatus `json:"cadence_status,omitzero" api:"required"`
	// Share of machine time a plan may fill.
	CapacityHeadroomPct float64 `json:"capacity_headroom_pct" api:"required"`
	// Typical changeover duration.
	ChangeoverAvgMinutes float64 `json:"changeover_avg_minutes" api:"required"`
	// Hourly labour rate charged to a changeover.
	ChangeoverLaborRate float64 `json:"changeover_labor_rate" api:"required"`
	// Longest plausible changeover.
	ChangeoverMaxMinutes float64 `json:"changeover_max_minutes" api:"required"`
	// Shortest plausible changeover.
	ChangeoverMinMinutes float64 `json:"changeover_min_minutes" api:"required"`
	// Default weeks of lead time at the constraint.
	DefaultConstraintLeadTimeWeeks float64 `json:"default_constraint_lead_time_weeks" api:"required"`
	// Units in a default production lot.
	DefaultLotUnits float64 `json:"default_lot_units" api:"required"`
	// How demand is derived.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis UpdateProductionScheduleSettingsRequestDemandBasis `json:"demand_basis,omitzero" api:"required"`
	// Months of order history the demand baseline is drawn from.
	DemandWindowMonths int64 `json:"demand_window_months" api:"required"`
	// Weeks between finishing at the constraint and being sellable.
	FinishLeadTimeWeeks float64 `json:"finish_lead_time_weeks" api:"required"`
	// Months of history the forecast is fitted to.
	ForecastHistoryMonths int64 `json:"forecast_history_months" api:"required"`
	// Months the forecast projects forward.
	ForecastMonths int64 `json:"forecast_months" api:"required"`
	// Z-score applied to forecast variability.
	ForecastZ float64 `json:"forecast_z" api:"required"`
	// How many leading weeks become a commitment when a version is published.
	FrozenWeeks int64 `json:"frozen_weeks" api:"required"`
	// Timezone the cadence is interpreted in.
	GenerationTimezone string `json:"generation_timezone" api:"required"`
	// Annual cost of holding stock, as a share of item value.
	HoldingRatePct float64 `json:"holding_rate_pct" api:"required"`
	// Hours in a shift.
	HoursPerShift float64 `json:"hours_per_shift" api:"required"`
	// How many steps downstream department work is derived for.
	MaxFlowDepth int64 `json:"max_flow_depth" api:"required"`
	// Ceiling on how far ahead any item is built.
	MaxWeeksSupply float64 `json:"max_weeks_supply" api:"required"`
	// How many weeks a generated plan covers.
	PlanningHorizonWeeks int64 `json:"planning_horizon_weeks" api:"required"`
	// Z-score for service level safety stock targets.
	ServiceLevelZ float64 `json:"service_level_z" api:"required"`
	// Shifts worked per day.
	ShiftsPerDay int64 `json:"shifts_per_day" api:"required"`
	// Day a planning week starts, where 0 is Sunday.
	WeekStartDay int64 `json:"week_start_day" api:"required"`
	// Weeks worked per year.
	WeeksPerYear int64 `json:"weeks_per_year" api:"required"`
	// Days worked per week.
	WorkDaysPerWeek int64 `json:"work_days_per_week" api:"required"`
	// ID of the department that sets the pace of the factory. Every machine in it is
	// planned, and every step downstream responds.
	ConstraintDepartmentID param.Opt[string] `json:"constraint_department_id,omitzero"`
	// Cron expression driving the generation cadence.
	GenerationCron param.Opt[string] `json:"generation_cron,omitzero"`
	paramObj
}

func (r UpdateProductionScheduleSettingsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateProductionScheduleSettingsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateProductionScheduleSettingsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether a generated version is published automatically.
type UpdateProductionScheduleSettingsRequestAutoPublishStatus string

const (
	UpdateProductionScheduleSettingsRequestAutoPublishStatusActive   UpdateProductionScheduleSettingsRequestAutoPublishStatus = "active"
	UpdateProductionScheduleSettingsRequestAutoPublishStatusInactive UpdateProductionScheduleSettingsRequestAutoPublishStatus = "inactive"
)

// Whether schedules are generated on a timer.
type UpdateProductionScheduleSettingsRequestCadenceStatus string

const (
	UpdateProductionScheduleSettingsRequestCadenceStatusActive   UpdateProductionScheduleSettingsRequestCadenceStatus = "active"
	UpdateProductionScheduleSettingsRequestCadenceStatusInactive UpdateProductionScheduleSettingsRequestCadenceStatus = "inactive"
)

// How demand is derived.
type UpdateProductionScheduleSettingsRequestDemandBasis string

const (
	UpdateProductionScheduleSettingsRequestDemandBasisTrailing12  UpdateProductionScheduleSettingsRequestDemandBasis = "trailing_12"
	UpdateProductionScheduleSettingsRequestDemandBasisSeasonalEma UpdateProductionScheduleSettingsRequestDemandBasis = "seasonal_ema"
)

type OperationProductionScheduleSettingUpdateParams struct {
	// Request to replace the account's planning assumptions.
	UpdateProductionScheduleSettingsRequest UpdateProductionScheduleSettingsRequestParam
	paramObj
}

func (r OperationProductionScheduleSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateProductionScheduleSettingsRequest)
}
func (r *OperationProductionScheduleSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
