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

// Generate and review machine-level production schedules.
//
// OperationProductionScheduleLineService contains methods and other services that
// help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationProductionScheduleLineService] method instead.
type OperationProductionScheduleLineService struct {
	options []option.RequestOption
}

// NewOperationProductionScheduleLineService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationProductionScheduleLineService(opts ...option.RequestOption) (r OperationProductionScheduleLineService) {
	r = OperationProductionScheduleLineService{}
	r.options = opts
	return
}

// Adds a campaign to a schedule by hand.
//
// The line is recorded as manual, so a later regenerate can tell it apart from
// what the solver produced, and the change is written to the deviation log. Adding
// into a frozen week requires a `reason`.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleLineService) New(ctx context.Context, id string, body OperationProductionScheduleLineNewParams, opts ...option.RequestOption) (res *ProductionScheduleLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/lines", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Edits a campaign on a schedule.
//
// Every change is written to the deviation log with a full before-and-after
// snapshot, and the line becomes manual so a regenerate can tell it apart from
// solver output. A change that touches a frozen week — including moving a campaign
// out of one — requires a `reason`.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleLineService) Update(ctx context.Context, lineID string, params OperationProductionScheduleLineUpdateParams, opts ...option.RequestOption) (res *ProductionScheduleLine, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if lineID == "" {
		err = errors.New("missing required line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/lines/%s", url.PathEscape(params.ID), url.PathEscape(lineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns the planned campaigns for a schedule version, in the order they run.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleLineService) List(ctx context.Context, id string, query OperationProductionScheduleLineListParams, opts ...option.RequestOption) (res *ListProductionScheduleLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/lines", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes a campaign from a schedule.
//
// The deviation log keeps a full snapshot of the removed line, so the change stays
// readable after the line itself is gone. Removing from a frozen week requires a
// `reason`.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleLineService) Delete(ctx context.Context, lineID string, params OperationProductionScheduleLineDeleteParams, opts ...option.RequestOption) (res *OperationProductionScheduleLineDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if lineID == "" {
		err = errors.New("missing required line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/lines/%s", url.PathEscape(params.ID), url.PathEscape(lineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Request to add a campaign to a schedule by hand.
//
// The properties ItemID, MachineID, Quantity, WeekIndex are required.
type CreateProductionScheduleLineRequestParam struct {
	// ID of the item to build.
	ItemID string `json:"item_id" api:"required"`
	// ID of the machine that will run it.
	MachineID string `json:"machine_id" api:"required"`
	// Units to build.
	Quantity float64 `json:"quantity" api:"required"`
	// Horizon week to plan the campaign in, zero-based.
	WeekIndex int64 `json:"week_index" api:"required"`
	// Lots the quantity represents.
	Lots param.Opt[int64] `json:"lots,omitzero"`
	// Free-form explanation of the change.
	ReasonNote param.Opt[string] `json:"reason_note,omitzero"`
	// Machine hours the campaign will take.
	RunHours param.Opt[float64] `json:"run_hours,omitzero"`
	// Why the campaign was added. Required when it lands inside a frozen week.
	//
	// Any of "machine_down", "material_shortage", "rush_order", "quality_hold",
	// "over_run", "under_run", "capacity_change", "other".
	Reason CreateProductionScheduleLineRequestReason `json:"reason,omitzero"`
	paramObj
}

func (r CreateProductionScheduleLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateProductionScheduleLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateProductionScheduleLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the campaign was added. Required when it lands inside a frozen week.
type CreateProductionScheduleLineRequestReason string

const (
	CreateProductionScheduleLineRequestReasonMachineDown      CreateProductionScheduleLineRequestReason = "machine_down"
	CreateProductionScheduleLineRequestReasonMaterialShortage CreateProductionScheduleLineRequestReason = "material_shortage"
	CreateProductionScheduleLineRequestReasonRushOrder        CreateProductionScheduleLineRequestReason = "rush_order"
	CreateProductionScheduleLineRequestReasonQualityHold      CreateProductionScheduleLineRequestReason = "quality_hold"
	CreateProductionScheduleLineRequestReasonOverRun          CreateProductionScheduleLineRequestReason = "over_run"
	CreateProductionScheduleLineRequestReasonUnderRun         CreateProductionScheduleLineRequestReason = "under_run"
	CreateProductionScheduleLineRequestReasonCapacityChange   CreateProductionScheduleLineRequestReason = "capacity_change"
	CreateProductionScheduleLineRequestReasonOther            CreateProductionScheduleLineRequestReason = "other"
)

// List represents a paginated list of resources.
type ListProductionScheduleLine struct {
	// Resources in this page.
	Data []ProductionScheduleLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleLineObject `json:"object" api:"required"`
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
func (r ListProductionScheduleLine) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleLineObject string

const (
	ListProductionScheduleLineObjectList ListProductionScheduleLineObject = "list"
)

// A saved campaign on a production schedule.
type ProductionScheduleLine struct {
	// Line ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Department Entity `json:"department" api:"required"`
	// Whether the line is inside the frozen window and can no longer be changed
	// without recording a deviation.
	//
	// Any of "frozen", "flexible".
	FreezeStatus ProductionScheduleLineFreezeStatus `json:"freeze_status" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Machine Entity `json:"machine" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_line".
	Object ProductionScheduleLineObject `json:"object" api:"required"`
	// Modelled changeover time before the campaign.
	PlannedChangeoverMinutes float64 `json:"planned_changeover_minutes" api:"required"`
	// Units in one lot, which is the batch size the week is released to the floor in.
	PlannedLotUnits float64 `json:"planned_lot_units" api:"required"`
	// Whole lots the quantity rounds to.
	PlannedLots int64 `json:"planned_lots" api:"required"`
	// Quantity to produce.
	PlannedQuantity float64 `json:"planned_quantity" api:"required"`
	// Constraint hours the campaign consumes.
	PlannedRunHours float64 `json:"planned_run_hours" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	PlannedUnit Entity `json:"planned_unit" api:"required"`
	// Abbreviation of the unit every quantity on this line is counted in, for display.
	//
	// A campaign of 360 means 360 pairs or 360 eaches depending on this, so the two
	// are never meaningful apart.
	PlannedUnitAbbreviation string `json:"planned_unit_abbreviation" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionRun Entity `json:"production_run" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionSchedule Entity `json:"production_schedule" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionStep Entity `json:"production_step" api:"required"`
	// Projected stock after the campaign lands and the week's demand is drawn down.
	ProjectedOnHandAfter float64 `json:"projected_on_hand_after" api:"required"`
	// Projected stock before the campaign lands.
	ProjectedOnHandBefore float64 `json:"projected_on_hand_before" api:"required"`
	// Why the campaign was scheduled.
	//
	// Any of "machine_down", "material_shortage", "rush_order", "quality_hold",
	// "over_run", "under_run", "capacity_change", "other".
	Reason ProductionScheduleLineReason `json:"reason" api:"required"`
	// Batches this campaign issued to the floor when its week was released.
	//
	// Zero until the week is released.
	ReleasedBatchCount int64 `json:"released_batch_count" api:"required"`
	// Batches of this campaign the floor has scanned.
	ScannedBatchCount int64 `json:"scanned_batch_count" api:"required"`
	// Quantity scanned so far, in the planned unit.
	//
	// Measured from the run the week was released as, matched on this campaign's item,
	// so a run holding several SKUs credits each campaign with only its own work.
	ScannedQuantity float64 `json:"scanned_quantity" api:"required"`
	// Order the campaign runs within its week.
	SequenceIndex int64 `json:"sequence_index" api:"required"`
	// Whether the solver or a person created the line.
	//
	// Any of "solver", "manual".
	Source ProductionScheduleLineSource `json:"source" api:"required"`
	// Where the line is in its lifecycle.
	//
	// Any of "planned", "released", "in_progress", "complete", "cancelled".
	Status ProductionScheduleLineStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex int64 `json:"week_index" api:"required"`
	// First instant of the week this campaign runs in.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		Department               respjson.Field
		FreezeStatus             respjson.Field
		Item                     respjson.Field
		Machine                  respjson.Field
		Object                   respjson.Field
		PlannedChangeoverMinutes respjson.Field
		PlannedLotUnits          respjson.Field
		PlannedLots              respjson.Field
		PlannedQuantity          respjson.Field
		PlannedRunHours          respjson.Field
		PlannedUnit              respjson.Field
		PlannedUnitAbbreviation  respjson.Field
		ProductionRun            respjson.Field
		ProductionSchedule       respjson.Field
		ProductionStep           respjson.Field
		ProjectedOnHandAfter     respjson.Field
		ProjectedOnHandBefore    respjson.Field
		Reason                   respjson.Field
		ReleasedBatchCount       respjson.Field
		ScannedBatchCount        respjson.Field
		ScannedQuantity          respjson.Field
		SequenceIndex            respjson.Field
		Source                   respjson.Field
		Status                   respjson.Field
		UpdatedAt                respjson.Field
		WeekIndex                respjson.Field
		WeekStartsAt             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleLine) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the line is inside the frozen window and can no longer be changed
// without recording a deviation.
type ProductionScheduleLineFreezeStatus string

const (
	ProductionScheduleLineFreezeStatusFrozen   ProductionScheduleLineFreezeStatus = "frozen"
	ProductionScheduleLineFreezeStatusFlexible ProductionScheduleLineFreezeStatus = "flexible"
)

// Resource type identifier.
type ProductionScheduleLineObject string

const (
	ProductionScheduleLineObjectProductionScheduleLine ProductionScheduleLineObject = "production_schedule_line"
)

// Why the campaign was scheduled.
type ProductionScheduleLineReason string

const (
	ProductionScheduleLineReasonMachineDown      ProductionScheduleLineReason = "machine_down"
	ProductionScheduleLineReasonMaterialShortage ProductionScheduleLineReason = "material_shortage"
	ProductionScheduleLineReasonRushOrder        ProductionScheduleLineReason = "rush_order"
	ProductionScheduleLineReasonQualityHold      ProductionScheduleLineReason = "quality_hold"
	ProductionScheduleLineReasonOverRun          ProductionScheduleLineReason = "over_run"
	ProductionScheduleLineReasonUnderRun         ProductionScheduleLineReason = "under_run"
	ProductionScheduleLineReasonCapacityChange   ProductionScheduleLineReason = "capacity_change"
	ProductionScheduleLineReasonOther            ProductionScheduleLineReason = "other"
)

// Whether the solver or a person created the line.
type ProductionScheduleLineSource string

const (
	ProductionScheduleLineSourceSolver ProductionScheduleLineSource = "solver"
	ProductionScheduleLineSourceManual ProductionScheduleLineSource = "manual"
)

// Where the line is in its lifecycle.
type ProductionScheduleLineStatus string

const (
	ProductionScheduleLineStatusPlanned    ProductionScheduleLineStatus = "planned"
	ProductionScheduleLineStatusReleased   ProductionScheduleLineStatus = "released"
	ProductionScheduleLineStatusInProgress ProductionScheduleLineStatus = "in_progress"
	ProductionScheduleLineStatusComplete   ProductionScheduleLineStatus = "complete"
	ProductionScheduleLineStatusCancelled  ProductionScheduleLineStatus = "cancelled"
)

// Request to edit a campaign on a schedule.
type UpdateProductionScheduleLineRequestParam struct {
	// Lots the quantity represents.
	Lots param.Opt[int64] `json:"lots,omitzero"`
	// ID of the machine to move the campaign to.
	MachineID param.Opt[string] `json:"machine_id,omitzero"`
	// Units to build.
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Free-form explanation of the change.
	ReasonNote param.Opt[string] `json:"reason_note,omitzero"`
	// Machine hours the campaign will take.
	RunHours param.Opt[float64] `json:"run_hours,omitzero"`
	// Position within the week's run order.
	SequenceIndex param.Opt[int64] `json:"sequence_index,omitzero"`
	// Horizon week to move the campaign to, zero-based.
	WeekIndex param.Opt[int64] `json:"week_index,omitzero"`
	// Why the campaign changed. Required when the change touches a frozen week.
	//
	// Any of "machine_down", "material_shortage", "rush_order", "quality_hold",
	// "over_run", "under_run", "capacity_change", "other".
	Reason UpdateProductionScheduleLineRequestReason `json:"reason,omitzero"`
	// Lifecycle state of the campaign.
	//
	// Any of "planned", "released", "in_progress", "complete", "cancelled".
	Status UpdateProductionScheduleLineRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r UpdateProductionScheduleLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateProductionScheduleLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateProductionScheduleLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the campaign changed. Required when the change touches a frozen week.
type UpdateProductionScheduleLineRequestReason string

const (
	UpdateProductionScheduleLineRequestReasonMachineDown      UpdateProductionScheduleLineRequestReason = "machine_down"
	UpdateProductionScheduleLineRequestReasonMaterialShortage UpdateProductionScheduleLineRequestReason = "material_shortage"
	UpdateProductionScheduleLineRequestReasonRushOrder        UpdateProductionScheduleLineRequestReason = "rush_order"
	UpdateProductionScheduleLineRequestReasonQualityHold      UpdateProductionScheduleLineRequestReason = "quality_hold"
	UpdateProductionScheduleLineRequestReasonOverRun          UpdateProductionScheduleLineRequestReason = "over_run"
	UpdateProductionScheduleLineRequestReasonUnderRun         UpdateProductionScheduleLineRequestReason = "under_run"
	UpdateProductionScheduleLineRequestReasonCapacityChange   UpdateProductionScheduleLineRequestReason = "capacity_change"
	UpdateProductionScheduleLineRequestReasonOther            UpdateProductionScheduleLineRequestReason = "other"
)

// Lifecycle state of the campaign.
type UpdateProductionScheduleLineRequestStatus string

const (
	UpdateProductionScheduleLineRequestStatusPlanned    UpdateProductionScheduleLineRequestStatus = "planned"
	UpdateProductionScheduleLineRequestStatusReleased   UpdateProductionScheduleLineRequestStatus = "released"
	UpdateProductionScheduleLineRequestStatusInProgress UpdateProductionScheduleLineRequestStatus = "in_progress"
	UpdateProductionScheduleLineRequestStatusComplete   UpdateProductionScheduleLineRequestStatus = "complete"
	UpdateProductionScheduleLineRequestStatusCancelled  UpdateProductionScheduleLineRequestStatus = "cancelled"
)

type OperationProductionScheduleLineDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationProductionScheduleLineDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationProductionScheduleLineDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleLineNewParams struct {
	// Request to add a campaign to a schedule by hand.
	CreateProductionScheduleLineRequest CreateProductionScheduleLineRequestParam
	paramObj
}

func (r OperationProductionScheduleLineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateProductionScheduleLineRequest)
}
func (r *OperationProductionScheduleLineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleLineUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Request to edit a campaign on a schedule.
	UpdateProductionScheduleLineRequest UpdateProductionScheduleLineRequestParam
	paramObj
}

func (r OperationProductionScheduleLineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateProductionScheduleLineRequest)
}
func (r *OperationProductionScheduleLineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleLineListParams struct {
	// Only return campaigns in this horizon week, zero-based.
	WeekIndex param.Opt[int64] `query:"week_index,omitzero" json:"-"`
	// Only return campaigns on these machines.
	MachineIDs []string `query:"machine_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationProductionScheduleLineListParams]'s query
// parameters as `url.Values`.
func (r OperationProductionScheduleLineListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationProductionScheduleLineDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Free-form explanation of the change.
	ReasonNote param.Opt[string] `query:"reason_note,omitzero" json:"-"`
	// Why the campaign was removed. Required when it sits in a frozen week.
	//
	// Any of "machine_down", "material_shortage", "rush_order", "quality_hold",
	// "over_run", "under_run", "capacity_change", "other".
	Reason OperationProductionScheduleLineDeleteParamsReason `query:"reason,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationProductionScheduleLineDeleteParams]'s query
// parameters as `url.Values`.
func (r OperationProductionScheduleLineDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Why the campaign was removed. Required when it sits in a frozen week.
type OperationProductionScheduleLineDeleteParamsReason string

const (
	OperationProductionScheduleLineDeleteParamsReasonMachineDown      OperationProductionScheduleLineDeleteParamsReason = "machine_down"
	OperationProductionScheduleLineDeleteParamsReasonMaterialShortage OperationProductionScheduleLineDeleteParamsReason = "material_shortage"
	OperationProductionScheduleLineDeleteParamsReasonRushOrder        OperationProductionScheduleLineDeleteParamsReason = "rush_order"
	OperationProductionScheduleLineDeleteParamsReasonQualityHold      OperationProductionScheduleLineDeleteParamsReason = "quality_hold"
	OperationProductionScheduleLineDeleteParamsReasonOverRun          OperationProductionScheduleLineDeleteParamsReason = "over_run"
	OperationProductionScheduleLineDeleteParamsReasonUnderRun         OperationProductionScheduleLineDeleteParamsReason = "under_run"
	OperationProductionScheduleLineDeleteParamsReasonCapacityChange   OperationProductionScheduleLineDeleteParamsReason = "capacity_change"
	OperationProductionScheduleLineDeleteParamsReasonOther            OperationProductionScheduleLineDeleteParamsReason = "other"
)
