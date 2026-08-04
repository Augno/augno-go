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

// Analyze sales, orders, manufacturing, materials, and other business metrics.
//
// CoreAnalyticsService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreAnalyticsService] method instead.
type CoreAnalyticsService struct {
	options []option.RequestOption
}

// NewCoreAnalyticsService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreAnalyticsService(opts ...option.RequestOption) (r CoreAnalyticsService) {
	r = CoreAnalyticsService{}
	r.options = opts
	return
}

// Returns Overall Equipment Effectiveness (OEE) metrics by department.
//
// Availability is measured from logged machine downtime rather than inferred, so
// it requires both `planned_time` for the department and downtime events in the
// period. Departments with `has_downtime_data` false have no availability
// measurement, and their ratios are returned as null rather than as 100%.
//
// This endpoint requires the permission: `machine_downtime:read`.
func (r *CoreAnalyticsService) UpdateOee(ctx context.Context, body CoreAnalyticsUpdateOeeParams, opts ...option.RequestOption) (res *AnalyzeOeeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/analytics/oee"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns actual production measured against the plan that was live at the time.
//
// The baseline for each week is the schedule version published on or before that
// week began, so republishing mid-horizon cannot rewrite a week the floor has
// already worked. `baseline_schedules` names the versions used.
//
// Two ratios are returned because either alone misleads: `attainment_pct` caps
// each campaign at what was asked for, so over-building one SKU cannot hide a miss
// on another, while `output_ratio_pct` is uncapped and is what reveals
// over-production. Production with no matching planned campaign is reported as
// `unplanned_quantity` rather than discarded — that number is the clearest signal
// a schedule is being worked around.
//
// Every ratio is null rather than zero when nothing was planned, and
// `has_baseline` is false when nothing was ever published over the period.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *CoreAnalyticsService) UpdateScheduleAttainment(ctx context.Context, body CoreAnalyticsUpdateScheduleAttainmentParams, opts ...option.RequestOption) (res *AnalyzeScheduleAttainmentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/analytics/schedule-attainment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// AnalyzeOeeRequest is the request to analyze Overall Equipment Effectiveness
// (OEE).
//
// The properties EndDate, StartDate are required.
type AnalyzeOeeRequestParam struct {
	// The end date for the analysis period.
	EndDate time.Time `json:"end_date" api:"required" format:"date-time"`
	// The start date for the analysis period.
	StartDate time.Time `json:"start_date" api:"required" format:"date-time"`
	// Optional department IDs to filter by.
	DepartmentIDs []string `json:"department_ids,omitzero"`
	// Scheduled production time per department for the period. Availability,
	// performance and OEE are only returned for departments this covers.
	PlannedTime []OeeDepartmentPlannedTimeParam `json:"planned_time,omitzero"`
	paramObj
}

func (r AnalyzeOeeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AnalyzeOeeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyzeOeeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AnalyzeOeeResponse represents the response from the analyze OEE endpoint.
type AnalyzeOeeResponse struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Departments ListOeeDepartment `json:"departments" api:"required"`
	// Resource type identifier.
	//
	// Any of "analyze_oee_response".
	Object AnalyzeOeeResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Departments respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyzeOeeResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyzeOeeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AnalyzeOeeResponseObject string

const (
	AnalyzeOeeResponseObjectAnalyzeOeeResponse AnalyzeOeeResponseObject = "analyze_oee_response"
)

// AnalyzeScheduleAttainmentRequest is the request to measure production against
// plan.
//
// The properties EndDate, StartDate are required.
type AnalyzeScheduleAttainmentRequestParam struct {
	// The end date for the analysis period.
	EndDate time.Time `json:"end_date" api:"required" format:"date-time"`
	// The start date for the analysis period.
	StartDate time.Time `json:"start_date" api:"required" format:"date-time"`
	// Only measure production in these departments.
	DepartmentIDs []string `json:"department_ids,omitzero"`
	// The dimension to break the results down by. Defaults to `week`.
	//
	// Any of "week", "machine", "department", "item".
	GroupBy AnalyzeScheduleAttainmentRequestGroupBy `json:"group_by,omitzero"`
	// Only measure production on these machines.
	MachineIDs []string `json:"machine_ids,omitzero"`
	paramObj
}

func (r AnalyzeScheduleAttainmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AnalyzeScheduleAttainmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyzeScheduleAttainmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The dimension to break the results down by. Defaults to `week`.
type AnalyzeScheduleAttainmentRequestGroupBy string

const (
	AnalyzeScheduleAttainmentRequestGroupByWeek       AnalyzeScheduleAttainmentRequestGroupBy = "week"
	AnalyzeScheduleAttainmentRequestGroupByMachine    AnalyzeScheduleAttainmentRequestGroupBy = "machine"
	AnalyzeScheduleAttainmentRequestGroupByDepartment AnalyzeScheduleAttainmentRequestGroupBy = "department"
	AnalyzeScheduleAttainmentRequestGroupByItem       AnalyzeScheduleAttainmentRequestGroupBy = "item"
)

// Actual production measured against the plan that was live at the time.
//
// The baseline for each week is the version that was published on or before that
// week began, so republishing mid-horizon cannot rewrite a week the floor has
// already worked. `baseline_schedules` names the versions used, so any number here
// can be traced back to the plan that produced it.
type AnalyzeScheduleAttainmentResponse struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	BaselineSchedules ListEntity `json:"baseline_schedules" api:"required"`
	// Whether the period had a plan to measure against. When `no_baseline`, every
	// ratio is null and the period has no plan rather than a missed one.
	//
	// Any of "measured", "no_baseline".
	BaselineStatus AnalyzeScheduleAttainmentResponseBaselineStatus `json:"baseline_status" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Buckets ListAttainmentBucket `json:"buckets" api:"required"`
	// End of the measured period.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	FrozenAdherence ListFrozenAdherence `json:"frozen_adherence" api:"required"`
	// The dimension the breakdown is grouped by.
	//
	// Any of "week", "machine", "department", "item".
	GroupBy AnalyzeScheduleAttainmentResponseGroupBy `json:"group_by" api:"required"`
	// Resource type identifier.
	//
	// Any of "analyze_schedule_attainment_response".
	Object AnalyzeScheduleAttainmentResponseObject `json:"object" api:"required"`
	// Start of the measured period.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	// One row of a schedule-attainment breakdown.
	//
	// Both ratios are reported because either alone misleads. `attainment_pct` caps
	// each SKU at what was asked for, so over-building one easy item cannot paper over
	// a total miss on another; `output_ratio_pct` does not cap, so it is the only one
	// that reveals over-production.
	Totals AttainmentBucket `json:"totals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineSchedules respjson.Field
		BaselineStatus    respjson.Field
		Buckets           respjson.Field
		EndsAt            respjson.Field
		FrozenAdherence   respjson.Field
		GroupBy           respjson.Field
		Object            respjson.Field
		StartsAt          respjson.Field
		Totals            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyzeScheduleAttainmentResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyzeScheduleAttainmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the period had a plan to measure against. When `no_baseline`, every
// ratio is null and the period has no plan rather than a missed one.
type AnalyzeScheduleAttainmentResponseBaselineStatus string

const (
	AnalyzeScheduleAttainmentResponseBaselineStatusMeasured   AnalyzeScheduleAttainmentResponseBaselineStatus = "measured"
	AnalyzeScheduleAttainmentResponseBaselineStatusNoBaseline AnalyzeScheduleAttainmentResponseBaselineStatus = "no_baseline"
)

// The dimension the breakdown is grouped by.
type AnalyzeScheduleAttainmentResponseGroupBy string

const (
	AnalyzeScheduleAttainmentResponseGroupByWeek       AnalyzeScheduleAttainmentResponseGroupBy = "week"
	AnalyzeScheduleAttainmentResponseGroupByMachine    AnalyzeScheduleAttainmentResponseGroupBy = "machine"
	AnalyzeScheduleAttainmentResponseGroupByDepartment AnalyzeScheduleAttainmentResponseGroupBy = "department"
	AnalyzeScheduleAttainmentResponseGroupByItem       AnalyzeScheduleAttainmentResponseGroupBy = "item"
)

// Resource type identifier.
type AnalyzeScheduleAttainmentResponseObject string

const (
	AnalyzeScheduleAttainmentResponseObjectAnalyzeScheduleAttainmentResponse AnalyzeScheduleAttainmentResponseObject = "analyze_schedule_attainment_response"
)

// One row of a schedule-attainment breakdown.
//
// Both ratios are reported because either alone misleads. `attainment_pct` caps
// each SKU at what was asked for, so over-building one easy item cannot paper over
// a total miss on another; `output_ratio_pct` does not cap, so it is the only one
// that reveals over-production.
type AttainmentBucket struct {
	// Units actually produced.
	ActualQuantity float64 `json:"actual_quantity" api:"required"`
	// Share of the plan that was met. Null when nothing was planned.
	AttainmentPct float64 `json:"attainment_pct" api:"required"`
	// Batches scanned in this bucket.
	BatchCount int64 `json:"batch_count" api:"required"`
	// Identifies the bucket within the chosen grouping — a week start, machine ID,
	// department ID or item ID.
	Key string `json:"key" api:"required"`
	// Display label for the bucket.
	Label string `json:"label" api:"required"`
	// Units produced that were planned for, capped per campaign at what was asked.
	MatchedQuantity float64 `json:"matched_quantity" api:"required"`
	// Output as a share of plan, uncapped. Null when nothing was planned.
	OutputRatioPct float64 `json:"output_ratio_pct" api:"required"`
	// Planned campaigns in this bucket.
	PlannedLines int64 `json:"planned_lines" api:"required"`
	// Units the live plan called for.
	PlannedQuantity float64 `json:"planned_quantity" api:"required"`
	// Machine hours the plan called for.
	PlannedRunHours float64 `json:"planned_run_hours" api:"required"`
	// Units produced with no matching planned campaign.
	UnplannedQuantity float64 `json:"unplanned_quantity" api:"required"`
	// Units scrapped.
	WasteQuantity float64 `json:"waste_quantity" api:"required"`
	// First day of the week, when grouping by week.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualQuantity    respjson.Field
		AttainmentPct     respjson.Field
		BatchCount        respjson.Field
		Key               respjson.Field
		Label             respjson.Field
		MatchedQuantity   respjson.Field
		OutputRatioPct    respjson.Field
		PlannedLines      respjson.Field
		PlannedQuantity   respjson.Field
		PlannedRunHours   respjson.Field
		UnplannedQuantity respjson.Field
		WasteQuantity     respjson.Field
		WeekStartsAt      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttainmentBucket) RawJSON() string { return r.JSON.raw }
func (r *AttainmentBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How well a published commitment survived the week it covered.
type FrozenAdherence struct {
	// Total absolute unit change across frozen-week deviations.
	AbsDeltaUnits float64 `json:"abs_delta_units" api:"required"`
	// Campaigns added into the frozen window after publish.
	AddedLines int64 `json:"added_lines" api:"required"`
	// Frozen campaigns that were changed after publish.
	DeviatedLines int64 `json:"deviated_lines" api:"required"`
	// Campaigns frozen at publish.
	FrozenLineCount int64 `json:"frozen_line_count" api:"required"`
	// Units frozen at publish.
	FrozenPlannedQuantity float64 `json:"frozen_planned_quantity" api:"required"`
	// Last day of the frozen window.
	FrozenThroughAt time.Time `json:"frozen_through_at" api:"required" format:"date-time"`
	// Share of frozen campaigns that survived untouched. Null when nothing was frozen.
	LineAdherencePct float64 `json:"line_adherence_pct" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Schedule Entity `json:"schedule" api:"required"`
	// Share of frozen units that survived untouched. Null when nothing was frozen.
	UnitsAdherencePct float64 `json:"units_adherence_pct" api:"required"`
	// Version number of that schedule.
	Version int64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AbsDeltaUnits         respjson.Field
		AddedLines            respjson.Field
		DeviatedLines         respjson.Field
		FrozenLineCount       respjson.Field
		FrozenPlannedQuantity respjson.Field
		FrozenThroughAt       respjson.Field
		LineAdherencePct      respjson.Field
		Schedule              respjson.Field
		UnitsAdherencePct     respjson.Field
		Version               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FrozenAdherence) RawJSON() string { return r.JSON.raw }
func (r *FrozenAdherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAttainmentBucket struct {
	// Resources in this page.
	Data []AttainmentBucket `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAttainmentBucketObject `json:"object" api:"required"`
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
func (r ListAttainmentBucket) RawJSON() string { return r.JSON.raw }
func (r *ListAttainmentBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAttainmentBucketObject string

const (
	ListAttainmentBucketObjectList ListAttainmentBucketObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListFrozenAdherence struct {
	// Resources in this page.
	Data []FrozenAdherence `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListFrozenAdherenceObject `json:"object" api:"required"`
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
func (r ListFrozenAdherence) RawJSON() string { return r.JSON.raw }
func (r *ListFrozenAdherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListFrozenAdherenceObject string

const (
	ListFrozenAdherenceObjectList ListFrozenAdherenceObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOeeDepartment struct {
	// Resources in this page.
	Data []OeeDepartment `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOeeDepartmentObject `json:"object" api:"required"`
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
func (r ListOeeDepartment) RawJSON() string { return r.JSON.raw }
func (r *ListOeeDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOeeDepartmentObject string

const (
	ListOeeDepartmentObjectList ListOeeDepartmentObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOeeDowntimeReason struct {
	// Resources in this page.
	Data []OeeDowntimeReason `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOeeDowntimeReasonObject `json:"object" api:"required"`
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
func (r ListOeeDowntimeReason) RawJSON() string { return r.JSON.raw }
func (r *ListOeeDowntimeReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOeeDowntimeReasonObject string

const (
	ListOeeDowntimeReasonObjectList ListOeeDowntimeReasonObject = "list"
)

// OeeDepartment represents OEE metrics for a single department.
type OeeDepartment struct {
	// Data-quality warnings for this grouping. Empty when the numbers can be taken at
	// face value.
	//
	// Any of "performance_above_capacity".
	Anomalies []string `json:"anomalies" api:"required"`
	// Logged downtime charged against availability, in seconds.
	AvailabilityLossSeconds float64 `json:"availability_loss_seconds" api:"required"`
	// Run time divided by scheduled time.
	AvailabilityPct float64 `json:"availability_pct" api:"required"`
	// Time spent changing over between products, in seconds.
	ChangeoverSeconds float64 `json:"changeover_seconds" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Department Entity `json:"department" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	DowntimeBreakdown ListOeeDowntimeReason `json:"downtime_breakdown" api:"required"`
	// Number of downtime events logged in the period.
	DowntimeEventCount int64 `json:"downtime_event_count" api:"required"`
	// The estimated runtime in hours.
	EstimatedRuntimeHours float64 `json:"estimated_runtime_hours" api:"required"`
	// The number of good units produced.
	GoodUnits float64 `json:"good_units" api:"required"`
	// Whether availability was measured from logged downtime or estimated from
	// runtime. A department with no logged downtime computes as perfectly available,
	// so an estimate is labelled rather than presented as a measurement.
	//
	// Any of "measured", "estimated".
	MeasurementStatus OeeDepartmentMeasurementStatus `json:"measurement_status" api:"required"`
	// Time nobody planned to run, removed from the OEE denominator rather than counted
	// as a loss.
	NotScheduledSeconds float64 `json:"not_scheduled_seconds" api:"required"`
	// Availability multiplied by performance multiplied by quality.
	OeePct float64 `json:"oee_pct" api:"required"`
	// Logged downtime charged against performance, in seconds.
	PerformanceLossSeconds float64 `json:"performance_loss_seconds" api:"required"`
	// Standard seconds earned divided by run time: how fast the department ran against
	// the designed speed of its production steps.
	PerformancePct float64 `json:"performance_pct" api:"required"`
	// Logged downtime charged against quality, in seconds.
	QualityLossSeconds float64 `json:"quality_loss_seconds" api:"required"`
	// Good units divided by total units produced.
	QualityPct float64 `json:"quality_pct" api:"required"`
	// Scheduled time net of availability losses, in seconds.
	RunTimeSeconds float64 `json:"run_time_seconds" api:"required"`
	// Planned time net of not-scheduled downtime, in seconds.
	ScheduledSeconds float64 `json:"scheduled_seconds" api:"required"`
	// The number of seconds units.
	SecondsUnits float64 `json:"seconds_units" api:"required"`
	// The time this output should have taken at each production step's own labor rate:
	// ideal cycle time multiplied by the units produced. This is the numerator of
	// Performance.
	StandardSecondsEarned float64 `json:"standard_seconds_earned" api:"required"`
	// The number of waste units.
	WasteUnits float64 `json:"waste_units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Anomalies               respjson.Field
		AvailabilityLossSeconds respjson.Field
		AvailabilityPct         respjson.Field
		ChangeoverSeconds       respjson.Field
		Department              respjson.Field
		DowntimeBreakdown       respjson.Field
		DowntimeEventCount      respjson.Field
		EstimatedRuntimeHours   respjson.Field
		GoodUnits               respjson.Field
		MeasurementStatus       respjson.Field
		NotScheduledSeconds     respjson.Field
		OeePct                  respjson.Field
		PerformanceLossSeconds  respjson.Field
		PerformancePct          respjson.Field
		QualityLossSeconds      respjson.Field
		QualityPct              respjson.Field
		RunTimeSeconds          respjson.Field
		ScheduledSeconds        respjson.Field
		SecondsUnits            respjson.Field
		StandardSecondsEarned   respjson.Field
		WasteUnits              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OeeDepartment) RawJSON() string { return r.JSON.raw }
func (r *OeeDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether availability was measured from logged downtime or estimated from
// runtime. A department with no logged downtime computes as perfectly available,
// so an estimate is labelled rather than presented as a measurement.
type OeeDepartmentMeasurementStatus string

const (
	OeeDepartmentMeasurementStatusMeasured  OeeDepartmentMeasurementStatus = "measured"
	OeeDepartmentMeasurementStatusEstimated OeeDepartmentMeasurementStatus = "estimated"
)

// OeeDepartmentPlannedTime supplies the scheduled production time for one
// department.
//
// The properties DepartmentID, PlannedHours are required.
type OeeDepartmentPlannedTimeParam struct {
	// The department ID.
	DepartmentID string `json:"department_id" api:"required"`
	// Scheduled production hours for the period.
	PlannedHours float64 `json:"planned_hours" api:"required"`
	paramObj
}

func (r OeeDepartmentPlannedTimeParam) MarshalJSON() (data []byte, err error) {
	type shadow OeeDepartmentPlannedTimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OeeDepartmentPlannedTimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OeeDowntimeReason represents one reason's contribution to a department's
// downtime.
type OeeDowntimeReason struct {
	// Downtime attributed to this reason, in seconds.
	DowntimeSeconds float64 `json:"downtime_seconds" api:"required"`
	// Number of events logged against this reason.
	EventCount int64 `json:"event_count" api:"required"`
	// Which OEE term this reason charges.
	//
	// Any of "availability", "performance", "quality", "not_scheduled".
	OeeBucket OeeDowntimeReasonOeeBucket `json:"oee_bucket" api:"required"`
	// Why the machine stopped.
	//
	// Any of "breakdown", "changeover", "material_shortage", "no_operator",
	// "planned_maintenance", "minor_stop", "quality_hold", "no_schedule".
	Reason OeeDowntimeReasonReason `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DowntimeSeconds respjson.Field
		EventCount      respjson.Field
		OeeBucket       respjson.Field
		Reason          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OeeDowntimeReason) RawJSON() string { return r.JSON.raw }
func (r *OeeDowntimeReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which OEE term this reason charges.
type OeeDowntimeReasonOeeBucket string

const (
	OeeDowntimeReasonOeeBucketAvailability OeeDowntimeReasonOeeBucket = "availability"
	OeeDowntimeReasonOeeBucketPerformance  OeeDowntimeReasonOeeBucket = "performance"
	OeeDowntimeReasonOeeBucketQuality      OeeDowntimeReasonOeeBucket = "quality"
	OeeDowntimeReasonOeeBucketNotScheduled OeeDowntimeReasonOeeBucket = "not_scheduled"
)

// Why the machine stopped.
type OeeDowntimeReasonReason string

const (
	OeeDowntimeReasonReasonBreakdown          OeeDowntimeReasonReason = "breakdown"
	OeeDowntimeReasonReasonChangeover         OeeDowntimeReasonReason = "changeover"
	OeeDowntimeReasonReasonMaterialShortage   OeeDowntimeReasonReason = "material_shortage"
	OeeDowntimeReasonReasonNoOperator         OeeDowntimeReasonReason = "no_operator"
	OeeDowntimeReasonReasonPlannedMaintenance OeeDowntimeReasonReason = "planned_maintenance"
	OeeDowntimeReasonReasonMinorStop          OeeDowntimeReasonReason = "minor_stop"
	OeeDowntimeReasonReasonQualityHold        OeeDowntimeReasonReason = "quality_hold"
	OeeDowntimeReasonReasonNoSchedule         OeeDowntimeReasonReason = "no_schedule"
)

type CoreAnalyticsUpdateOeeParams struct {
	// AnalyzeOeeRequest is the request to analyze Overall Equipment Effectiveness
	// (OEE).
	AnalyzeOeeRequest AnalyzeOeeRequestParam
	paramObj
}

func (r CoreAnalyticsUpdateOeeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AnalyzeOeeRequest)
}
func (r *CoreAnalyticsUpdateOeeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CoreAnalyticsUpdateScheduleAttainmentParams struct {
	// AnalyzeScheduleAttainmentRequest is the request to measure production against
	// plan.
	AnalyzeScheduleAttainmentRequest AnalyzeScheduleAttainmentRequestParam
	paramObj
}

func (r CoreAnalyticsUpdateScheduleAttainmentParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AnalyzeScheduleAttainmentRequest)
}
func (r *CoreAnalyticsUpdateScheduleAttainmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
