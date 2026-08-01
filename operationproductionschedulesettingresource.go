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
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// The planning assumptions production schedules are solved against, and the
// per-resource overrides that mark which machines constrain the plan.
//
// OperationProductionScheduleSettingResourceService contains methods and other
// services that help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationProductionScheduleSettingResourceService] method instead.
type OperationProductionScheduleSettingResourceService struct {
	options []option.RequestOption
}

// NewOperationProductionScheduleSettingResourceService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewOperationProductionScheduleSettingResourceService(opts ...option.RequestOption) (r OperationProductionScheduleSettingResourceService) {
	r = OperationProductionScheduleSettingResourceService{}
	r.options = opts
	return
}

// Writes a per-resource planning override.
//
// One override exists per resource, so this replaces any existing entry for the
// same scope rather than adding a second. Machines are selected by the constraint
// department, so this is where one is taken _out_ of planning — a machine down for
// a rebuild — and where a department or step declares how many weeks after the
// constraint its work starts.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleSettingResourceService) Update(ctx context.Context, body OperationProductionScheduleSettingResourceUpdateParams, opts ...option.RequestOption) (res *ProductionScheduleResourceSetting, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedule-settings/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the per-machine, per-department and per-step planning overrides.
//
// This is where machines are marked as the planning constraint, and where a
// department or step declares how many weeks after the constraint its work starts.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleSettingResourceService) List(ctx context.Context, opts ...option.RequestOption) (res *ListProductionScheduleResourceSetting, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedule-settings/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes a per-resource planning override, returning that resource to the account
// defaults.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleSettingResourceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationProductionScheduleSettingResourceDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedule-settings/resources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List represents a paginated list of resources.
type ListProductionScheduleResourceSetting struct {
	// Resources in this page.
	Data []ProductionScheduleResourceSetting `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleResourceSettingObject `json:"object" api:"required"`
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
func (r ListProductionScheduleResourceSetting) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleResourceSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleResourceSettingObject string

const (
	ListProductionScheduleResourceSettingObjectList ListProductionScheduleResourceSettingObject = "list"
)

// A per-resource override of the account's planning assumptions.
//
// This is where a machine is marked as the planning constraint, and where a
// department or step declares how many weeks after the constraint its work
// actually starts.
type ProductionScheduleResourceSetting struct {
	// Resource setting ID.
	ID string `json:"id" api:"required"`
	// Weeks after the constraint campaign this resource's work starts.
	LeadTimeOffsetWeeks float64 `json:"lead_time_offset_weeks" api:"required"`
	// Weeks of lead time at this resource.
	LeadTimeWeeks float64 `json:"lead_time_weeks" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_resource_setting".
	Object ProductionScheduleResourceSettingObject `json:"object" api:"required"`
	// Whether this resource takes part in planning. Machines are selected by
	// department, so this takes one out rather than opting one in — for a machine down
	// for a rebuild that should not be planned against.
	//
	// Any of "included", "excluded".
	ParticipationStatus ProductionScheduleResourceSettingParticipationStatus `json:"participation_status" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Scope Entity `json:"scope" api:"required"`
	// What kind of resource this overrides.
	//
	// Any of "machine", "department", "production_step".
	ScopeType ProductionScheduleResourceSettingScopeType `json:"scope_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		LeadTimeOffsetWeeks respjson.Field
		LeadTimeWeeks       respjson.Field
		Object              respjson.Field
		ParticipationStatus respjson.Field
		Scope               respjson.Field
		ScopeType           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleResourceSetting) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleResourceSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionScheduleResourceSettingObject string

const (
	ProductionScheduleResourceSettingObjectProductionScheduleResourceSetting ProductionScheduleResourceSettingObject = "production_schedule_resource_setting"
)

// Whether this resource takes part in planning. Machines are selected by
// department, so this takes one out rather than opting one in — for a machine down
// for a rebuild that should not be planned against.
type ProductionScheduleResourceSettingParticipationStatus string

const (
	ProductionScheduleResourceSettingParticipationStatusIncluded ProductionScheduleResourceSettingParticipationStatus = "included"
	ProductionScheduleResourceSettingParticipationStatusExcluded ProductionScheduleResourceSettingParticipationStatus = "excluded"
)

// What kind of resource this overrides.
type ProductionScheduleResourceSettingScopeType string

const (
	ProductionScheduleResourceSettingScopeTypeMachine        ProductionScheduleResourceSettingScopeType = "machine"
	ProductionScheduleResourceSettingScopeTypeDepartment     ProductionScheduleResourceSettingScopeType = "department"
	ProductionScheduleResourceSettingScopeTypeProductionStep ProductionScheduleResourceSettingScopeType = "production_step"
)

// Request to write a per-resource planning override.
//
// The properties LeadTimeOffsetWeeks, ParticipationStatus, ScopeRefID, ScopeType
// are required.
type UpsertResourceSettingRequestParam struct {
	// Weeks after the constraint campaign this resource's work starts.
	LeadTimeOffsetWeeks float64 `json:"lead_time_offset_weeks" api:"required"`
	// Whether this resource takes part in planning. Machines are selected by
	// department, so this excludes one rather than opting one in.
	//
	// Any of "included", "excluded".
	ParticipationStatus UpsertResourceSettingRequestParticipationStatus `json:"participation_status,omitzero" api:"required"`
	// ID of the machine, department or production step.
	ScopeRefID string `json:"scope_ref_id" api:"required"`
	// What kind of resource this overrides.
	//
	// Any of "machine", "department", "production_step".
	ScopeType UpsertResourceSettingRequestScopeType `json:"scope_type,omitzero" api:"required"`
	// Weeks of lead time at this resource.
	LeadTimeWeeks param.Opt[float64] `json:"lead_time_weeks,omitzero"`
	paramObj
}

func (r UpsertResourceSettingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertResourceSettingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertResourceSettingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this resource takes part in planning. Machines are selected by
// department, so this excludes one rather than opting one in.
type UpsertResourceSettingRequestParticipationStatus string

const (
	UpsertResourceSettingRequestParticipationStatusIncluded UpsertResourceSettingRequestParticipationStatus = "included"
	UpsertResourceSettingRequestParticipationStatusExcluded UpsertResourceSettingRequestParticipationStatus = "excluded"
)

// What kind of resource this overrides.
type UpsertResourceSettingRequestScopeType string

const (
	UpsertResourceSettingRequestScopeTypeMachine        UpsertResourceSettingRequestScopeType = "machine"
	UpsertResourceSettingRequestScopeTypeDepartment     UpsertResourceSettingRequestScopeType = "department"
	UpsertResourceSettingRequestScopeTypeProductionStep UpsertResourceSettingRequestScopeType = "production_step"
)

type OperationProductionScheduleSettingResourceDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationProductionScheduleSettingResourceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationProductionScheduleSettingResourceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleSettingResourceUpdateParams struct {
	// Request to write a per-resource planning override.
	UpsertResourceSettingRequest UpsertResourceSettingRequestParam
	paramObj
}

func (r OperationProductionScheduleSettingResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpsertResourceSettingRequest)
}
func (r *OperationProductionScheduleSettingResourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
