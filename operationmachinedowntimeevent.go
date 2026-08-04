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

// Log and review machine stoppages. Downtime is the source of OEE availability and
// changeover time.
//
// OperationMachineDowntimeEventService contains methods and other services that
// help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationMachineDowntimeEventService] method instead.
type OperationMachineDowntimeEventService struct {
	options []option.RequestOption
}

// NewOperationMachineDowntimeEventService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationMachineDowntimeEventService(opts ...option.RequestOption) (r OperationMachineDowntimeEventService) {
	r = OperationMachineDowntimeEventService{}
	r.options = opts
	return
}

// Logs a machine downtime event.
//
// Omit `ended_at` while the machine is still down. A machine can only have one
// open event at a time, so logging a second open stoppage against a machine that
// is already down is rejected until the first is closed.
//
// The department is taken from the machine, the business day is taken from
// `started_at`, the event is attributed to the credentials that made the request,
// and the duration is calculated when the event is closed.
//
// This endpoint requires the permission: `machine_downtime:create`.
func (r *OperationMachineDowntimeEventService) New(ctx context.Context, params OperationMachineDowntimeEventNewParams, opts ...option.RequestOption) (res *MachineDowntimeEvent, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/machine-downtime-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single machine downtime event.
//
// This endpoint requires the permission: `machine_downtime:read`.
func (r *OperationMachineDowntimeEventService) Get(ctx context.Context, id string, query OperationMachineDowntimeEventGetParams, opts ...option.RequestOption) (res *MachineDowntimeEvent, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/machine-downtime-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Closes or corrects a machine downtime event.
//
// Only the fields provided in the request are changed. Setting `ended_at` closes
// the event and calculates its duration; sending it as null reopens an event
// closed by mistake, which is rejected when the machine already has another open
// stoppage. The machine an event belongs to cannot be changed.
//
// This endpoint requires the permission: `machine_downtime:update`.
func (r *OperationMachineDowntimeEventService) Update(ctx context.Context, id string, params OperationMachineDowntimeEventUpdateParams, opts ...option.RequestOption) (res *MachineDowntimeEvent, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/machine-downtime-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of machine downtime events, most recently started
// first.
//
// The search term matches text in the event note. Filters combine, so a machine, a
// reason and a date range narrow the list together.
//
// This endpoint requires the permission: `machine_downtime:read`.
func (r *OperationMachineDowntimeEventService) List(ctx context.Context, query OperationMachineDowntimeEventListParams, opts ...option.RequestOption) (res *ListMachineDowntimeEvent, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/machine-downtime-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a machine downtime event.
//
// Meant for a stoppage that was logged by mistake: the event is removed
// permanently and stops counting against the machine's availability. To correct a
// real stoppage, update it instead so the record of the downtime survives.
//
// This endpoint requires the permission: `machine_downtime:delete`.
func (r *OperationMachineDowntimeEventService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationMachineDowntimeEventDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/machine-downtime-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to log a machine downtime event.
//
// The properties MachineID, Reason, StartedAt are required.
type CreateMachineDowntimeEventRequestParam struct {
	// ID of the machine that stopped.
	MachineID string `json:"machine_id" api:"required"`
	// Why the machine stopped.
	//
	// The reason decides which OEE term the stoppage charges, so it does more than
	// label the event. Retrieve the available reasons and the term each one charges
	// from the downtime reasons list.
	//
	// Any of "breakdown", "changeover", "material_shortage", "no_operator",
	// "planned_maintenance", "minor_stop", "quality_hold", "no_schedule".
	Reason CreateMachineDowntimeEventRequestReason `json:"reason,omitzero" api:"required"`
	// When the machine stopped.
	//
	// Cannot be in the future beyond a few minutes of clock skew, which is allowed so
	// a shop-floor tablet running fast can still log "just now". The business day the
	// stoppage counts against is taken from this timestamp.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// ID of the batch in progress when the machine stopped.
	BatchID param.Opt[string] `json:"batch_id,omitzero"`
	// When the machine started running again.
	//
	// Omit it while the machine is still down; that leaves the event open, and the
	// duration is filled in once the event is closed. It must be later than
	// `started_at`.
	EndedAt param.Opt[time.Time] `json:"ended_at,omitzero" format:"date-time"`
	// ID of the item the machine was running when it stopped.
	ItemID param.Opt[string] `json:"item_id,omitzero"`
	// Free-form notes about the stoppage.
	//
	// Searchable from the downtime events list. Maximum 2000 characters.
	Note param.Opt[string] `json:"note,omitzero"`
	// ID of the production run in progress when the machine stopped.
	ProductionRunID param.Opt[string] `json:"production_run_id,omitzero"`
	// How the event was recorded.
	//
	// Records the stoppage as manually logged unless you say otherwise, so an
	// integration or shop-floor station should send its own source to keep
	// hand-entered downtime distinguishable.
	//
	// Any of "manual", "scanner", "inferred", "api".
	Source CreateMachineDowntimeEventRequestSource `json:"source,omitzero"`
	paramObj
}

func (r CreateMachineDowntimeEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMachineDowntimeEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMachineDowntimeEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the machine stopped.
//
// The reason decides which OEE term the stoppage charges, so it does more than
// label the event. Retrieve the available reasons and the term each one charges
// from the downtime reasons list.
type CreateMachineDowntimeEventRequestReason string

const (
	CreateMachineDowntimeEventRequestReasonBreakdown          CreateMachineDowntimeEventRequestReason = "breakdown"
	CreateMachineDowntimeEventRequestReasonChangeover         CreateMachineDowntimeEventRequestReason = "changeover"
	CreateMachineDowntimeEventRequestReasonMaterialShortage   CreateMachineDowntimeEventRequestReason = "material_shortage"
	CreateMachineDowntimeEventRequestReasonNoOperator         CreateMachineDowntimeEventRequestReason = "no_operator"
	CreateMachineDowntimeEventRequestReasonPlannedMaintenance CreateMachineDowntimeEventRequestReason = "planned_maintenance"
	CreateMachineDowntimeEventRequestReasonMinorStop          CreateMachineDowntimeEventRequestReason = "minor_stop"
	CreateMachineDowntimeEventRequestReasonQualityHold        CreateMachineDowntimeEventRequestReason = "quality_hold"
	CreateMachineDowntimeEventRequestReasonNoSchedule         CreateMachineDowntimeEventRequestReason = "no_schedule"
)

// How the event was recorded.
//
// Records the stoppage as manually logged unless you say otherwise, so an
// integration or shop-floor station should send its own source to keep
// hand-entered downtime distinguishable.
type CreateMachineDowntimeEventRequestSource string

const (
	CreateMachineDowntimeEventRequestSourceManual   CreateMachineDowntimeEventRequestSource = "manual"
	CreateMachineDowntimeEventRequestSourceScanner  CreateMachineDowntimeEventRequestSource = "scanner"
	CreateMachineDowntimeEventRequestSourceInferred CreateMachineDowntimeEventRequestSource = "inferred"
	CreateMachineDowntimeEventRequestSourceAPI      CreateMachineDowntimeEventRequestSource = "api"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMachineDowntimeEvent struct {
	// Resources in this page.
	Data []MachineDowntimeEvent `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMachineDowntimeEventObject `json:"object" api:"required"`
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
func (r ListMachineDowntimeEvent) RawJSON() string { return r.JSON.raw }
func (r *ListMachineDowntimeEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMachineDowntimeEventObject string

const (
	ListMachineDowntimeEventObjectList ListMachineDowntimeEventObject = "list"
)

// A period during which a machine was not running.
//
// Downtime is what makes OEE Availability a measurement rather than an estimate.
// An event with no `ended_at` is still open, meaning the machine is down right
// now; a machine can only have one open event at a time.
type MachineDowntimeEvent struct {
	// Downtime event ID.
	ID string `json:"id" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Batch Entity `json:"batch" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department Department `json:"department" api:"required"`
	// How long the machine was down, in seconds.
	//
	// Calculated when the event is closed, and recalculated whenever its start or end
	// time changes.
	DurationSeconds int64 `json:"duration_seconds" api:"required"`
	// When the machine started running again.
	EndedAt time.Time `json:"ended_at" api:"required" format:"date-time"`
	// An entry in your catalog: something you sell, consume, or build with.
	Item Item `json:"item" api:"required"`
	// A piece of production equipment, such as a CNC router or press, assigned to a
	// department.
	Machine Machine `json:"machine" api:"required"`
	// Free-form notes about the stoppage.
	Note string `json:"note" api:"required"`
	// Resource type identifier.
	//
	// Any of "machine_downtime_event".
	Object MachineDowntimeEventObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionRun Entity `json:"production_run" api:"required"`
	// The reason for a stoppage, as carried on a downtime event.
	//
	// A denormalized view of the reason taxonomy: the stable code plus the display
	// name and OEE bucket resolved from it at read time.
	Reason MachineDowntimeReasonSummary `json:"reason" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	ReportedBy Actor `json:"reported_by" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ScheduleLine Entity `json:"schedule_line" api:"required"`
	// The business day the stoppage is counted against.
	//
	// Taken from the calendar date of `started_at`, so correcting the start time can
	// move the stoppage onto a different day's totals.
	ShiftAt time.Time `json:"shift_at" api:"required" format:"date-time"`
	// The shift the stoppage is counted against.
	ShiftCode string `json:"shift_code" api:"required"`
	// How the event was recorded.
	//
	// - `manual`: a person logged the stoppage.
	// - `scanner`: a shop-floor station logged it.
	// - `inferred`: the system derived it from a gap in activity.
	// - `api`: an integration reported it.
	//
	// Any of "manual", "scanner", "inferred", "api".
	Source MachineDowntimeEventSource `json:"source" api:"required"`
	// When the machine stopped.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Batch           respjson.Field
		CreatedAt       respjson.Field
		Department      respjson.Field
		DurationSeconds respjson.Field
		EndedAt         respjson.Field
		Item            respjson.Field
		Machine         respjson.Field
		Note            respjson.Field
		Object          respjson.Field
		ProductionRun   respjson.Field
		Reason          respjson.Field
		ReportedBy      respjson.Field
		ScheduleLine    respjson.Field
		ShiftAt         respjson.Field
		ShiftCode       respjson.Field
		Source          respjson.Field
		StartedAt       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineDowntimeEvent) RawJSON() string { return r.JSON.raw }
func (r *MachineDowntimeEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MachineDowntimeEventObject string

const (
	MachineDowntimeEventObjectMachineDowntimeEvent MachineDowntimeEventObject = "machine_downtime_event"
)

// How the event was recorded.
//
// - `manual`: a person logged the stoppage.
// - `scanner`: a shop-floor station logged it.
// - `inferred`: the system derived it from a gap in activity.
// - `api`: an integration reported it.
type MachineDowntimeEventSource string

const (
	MachineDowntimeEventSourceManual   MachineDowntimeEventSource = "manual"
	MachineDowntimeEventSourceScanner  MachineDowntimeEventSource = "scanner"
	MachineDowntimeEventSourceInferred MachineDowntimeEventSource = "inferred"
	MachineDowntimeEventSourceAPI      MachineDowntimeEventSource = "api"
)

// Request to update a machine downtime event.
type UpdateMachineDowntimeEventRequestParam struct {
	// ID of the batch in progress when the machine stopped.
	//
	// Send null to detach the batch.
	BatchID param.Opt[string] `json:"batch_id,omitzero"`
	// When the machine started running again.
	//
	// Setting it closes the event and records the duration. Send null to reopen an
	// event that was closed by mistake, which is rejected if the machine has since had
	// another stoppage logged that is still open.
	EndedAt param.Opt[time.Time] `json:"ended_at,omitzero" format:"date-time"`
	// ID of the item the machine was running when it stopped.
	//
	// Send null to detach the item.
	ItemID param.Opt[string] `json:"item_id,omitzero"`
	// Free-form notes about the stoppage.
	//
	// Send null to remove the note. Maximum 2000 characters.
	Note param.Opt[string] `json:"note,omitzero"`
	// ID of the production run in progress when the machine stopped.
	//
	// Send null to detach the run.
	ProductionRunID param.Opt[string] `json:"production_run_id,omitzero"`
	// When the machine stopped.
	//
	// Correcting it recalculates the duration and can move the stoppage onto a
	// different business day.
	StartedAt param.Opt[time.Time] `json:"started_at,omitzero" format:"date-time"`
	// Why the machine stopped.
	//
	// Reclassifying a stoppage moves it to the OEE term the new reason charges, so
	// past availability figures change with it.
	//
	// Any of "breakdown", "changeover", "material_shortage", "no_operator",
	// "planned_maintenance", "minor_stop", "quality_hold", "no_schedule".
	Reason UpdateMachineDowntimeEventRequestReason `json:"reason,omitzero"`
	paramObj
}

func (r UpdateMachineDowntimeEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMachineDowntimeEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMachineDowntimeEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the machine stopped.
//
// Reclassifying a stoppage moves it to the OEE term the new reason charges, so
// past availability figures change with it.
type UpdateMachineDowntimeEventRequestReason string

const (
	UpdateMachineDowntimeEventRequestReasonBreakdown          UpdateMachineDowntimeEventRequestReason = "breakdown"
	UpdateMachineDowntimeEventRequestReasonChangeover         UpdateMachineDowntimeEventRequestReason = "changeover"
	UpdateMachineDowntimeEventRequestReasonMaterialShortage   UpdateMachineDowntimeEventRequestReason = "material_shortage"
	UpdateMachineDowntimeEventRequestReasonNoOperator         UpdateMachineDowntimeEventRequestReason = "no_operator"
	UpdateMachineDowntimeEventRequestReasonPlannedMaintenance UpdateMachineDowntimeEventRequestReason = "planned_maintenance"
	UpdateMachineDowntimeEventRequestReasonMinorStop          UpdateMachineDowntimeEventRequestReason = "minor_stop"
	UpdateMachineDowntimeEventRequestReasonQualityHold        UpdateMachineDowntimeEventRequestReason = "quality_hold"
	UpdateMachineDowntimeEventRequestReasonNoSchedule         UpdateMachineDowntimeEventRequestReason = "no_schedule"
)

type OperationMachineDowntimeEventDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationMachineDowntimeEventDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationMachineDowntimeEventDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationMachineDowntimeEventNewParams struct {
	// Request to log a machine downtime event.
	CreateMachineDowntimeEventRequest CreateMachineDowntimeEventRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "machine", "department", "item", "reported_by".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationMachineDowntimeEventNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMachineDowntimeEventRequest)
}
func (r *OperationMachineDowntimeEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationMachineDowntimeEventNewParams]'s query parameters
// as `url.Values`.
func (r OperationMachineDowntimeEventNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationMachineDowntimeEventGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "machine", "department", "item", "reported_by".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationMachineDowntimeEventGetParams]'s query parameters
// as `url.Values`.
func (r OperationMachineDowntimeEventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationMachineDowntimeEventUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "machine", "department", "item", "reported_by".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a machine downtime event.
	UpdateMachineDowntimeEventRequest UpdateMachineDowntimeEventRequestParam
	paramObj
}

func (r OperationMachineDowntimeEventUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateMachineDowntimeEventRequest)
}
func (r *OperationMachineDowntimeEventUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationMachineDowntimeEventUpdateParams]'s query
// parameters as `url.Values`.
func (r OperationMachineDowntimeEventUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationMachineDowntimeEventListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return events that started on or before this timestamp, formatted as
	// RFC3339.
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return events that are still open, meaning the machine is down right now.
	//
	// Sending `false` is the same as leaving it out: both open and closed events come
	// back.
	Open param.Opt[bool] `query:"open,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Only return events that started on or after this timestamp, formatted as
	// RFC3339.
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// Only return events for machines in these departments.
	DepartmentIDs []string `query:"department_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "machine", "department", "item", "reported_by".
	Include []string `query:"include,omitzero" json:"-"`
	// Only return events for these machines.
	MachineIDs []string `query:"machine_ids,omitzero" json:"-"`
	// Only return events logged against these reasons.
	//
	// Any of "breakdown", "changeover", "material_shortage", "no_operator",
	// "planned_maintenance", "minor_stop", "quality_hold", "no_schedule".
	Reasons []string `query:"reasons,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationMachineDowntimeEventListParams]'s query parameters
// as `url.Values`.
func (r OperationMachineDowntimeEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
