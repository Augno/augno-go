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

// List and manage scanning stations.
//
// OperationScanningStationService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationScanningStationService] method instead.
type OperationScanningStationService struct {
	options []option.RequestOption
}

// NewOperationScanningStationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationScanningStationService(opts ...option.RequestOption) (r OperationScanningStationService) {
	r = OperationScanningStationService{}
	r.options = opts
	return
}

// Creates a scanning station associated with a department.
func (r *OperationScanningStationService) New(ctx context.Context, params OperationScanningStationNewParams, opts ...option.RequestOption) (res *ScanningStation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/scanning-stations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a scanning station by ID.
func (r *OperationScanningStationService) Get(ctx context.Context, id string, query OperationScanningStationGetParams, opts ...option.RequestOption) (res *ScanningStation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/scanning-stations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a scanning station.
func (r *OperationScanningStationService) Update(ctx context.Context, id string, params OperationScanningStationUpdateParams, opts ...option.RequestOption) (res *ScanningStation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/scanning-stations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of scanning stations for the current account.
func (r *OperationScanningStationService) List(ctx context.Context, query OperationScanningStationListParams, opts ...option.RequestOption) (res *ListScanningStation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/scanning-stations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a scanning station.
func (r *OperationScanningStationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationScanningStationDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/scanning-stations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a scanning station.
//
// The properties DepartmentID, Name, OperatorRequirement, Type are required.
type CreateScanningStationRequestParam struct {
	// Department ID.
	DepartmentID string `json:"department_id" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Operator requirement behavior for this station.
	//
	// Any of "none", "material_check".
	OperatorRequirement CreateScanningStationRequestOperatorRequirement `json:"operator_requirement,omitzero" api:"required"`
	// Scanning station type.
	//
	// Any of "init_batch", "merge_batch", "move_batch", "split_batch".
	Type CreateScanningStationRequestType `json:"type,omitzero" api:"required"`
	// Notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Label size code.
	//
	// Any of "1x1", "1x3", "1x4", "2x4".
	LabelSize CreateScanningStationRequestLabelSize `json:"label_size,omitzero"`
	// Label type code.
	//
	// Any of "tag", "traveler".
	LabelType CreateScanningStationRequestLabelType `json:"label_type,omitzero"`
	paramObj
}

func (r CreateScanningStationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateScanningStationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateScanningStationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operator requirement behavior for this station.
type CreateScanningStationRequestOperatorRequirement string

const (
	CreateScanningStationRequestOperatorRequirementNone          CreateScanningStationRequestOperatorRequirement = "none"
	CreateScanningStationRequestOperatorRequirementMaterialCheck CreateScanningStationRequestOperatorRequirement = "material_check"
)

// Scanning station type.
type CreateScanningStationRequestType string

const (
	CreateScanningStationRequestTypeInitBatch  CreateScanningStationRequestType = "init_batch"
	CreateScanningStationRequestTypeMergeBatch CreateScanningStationRequestType = "merge_batch"
	CreateScanningStationRequestTypeMoveBatch  CreateScanningStationRequestType = "move_batch"
	CreateScanningStationRequestTypeSplitBatch CreateScanningStationRequestType = "split_batch"
)

// Label size code.
type CreateScanningStationRequestLabelSize string

const (
	CreateScanningStationRequestLabelSize1x1 CreateScanningStationRequestLabelSize = "1x1"
	CreateScanningStationRequestLabelSize1x3 CreateScanningStationRequestLabelSize = "1x3"
	CreateScanningStationRequestLabelSize1x4 CreateScanningStationRequestLabelSize = "1x4"
	CreateScanningStationRequestLabelSize2x4 CreateScanningStationRequestLabelSize = "2x4"
)

// Label type code.
type CreateScanningStationRequestLabelType string

const (
	CreateScanningStationRequestLabelTypeTag      CreateScanningStationRequestLabelType = "tag"
	CreateScanningStationRequestLabelTypeTraveler CreateScanningStationRequestLabelType = "traveler"
)

// Request to partially update a scanning station.
type UpdateScanningStationRequestParam struct {
	// Notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Label size code.
	//
	// Any of "1x1", "1x3", "1x4", "2x4".
	LabelSize UpdateScanningStationRequestLabelSize `json:"label_size,omitzero"`
	// Label type code.
	//
	// Any of "tag", "traveler".
	LabelType UpdateScanningStationRequestLabelType `json:"label_type,omitzero"`
	// Operator requirement behavior for this station.
	//
	// Any of "none", "material_check".
	OperatorRequirement UpdateScanningStationRequestOperatorRequirement `json:"operator_requirement,omitzero"`
	paramObj
}

func (r UpdateScanningStationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateScanningStationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateScanningStationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Label size code.
type UpdateScanningStationRequestLabelSize string

const (
	UpdateScanningStationRequestLabelSize1x1 UpdateScanningStationRequestLabelSize = "1x1"
	UpdateScanningStationRequestLabelSize1x3 UpdateScanningStationRequestLabelSize = "1x3"
	UpdateScanningStationRequestLabelSize1x4 UpdateScanningStationRequestLabelSize = "1x4"
	UpdateScanningStationRequestLabelSize2x4 UpdateScanningStationRequestLabelSize = "2x4"
)

// Label type code.
type UpdateScanningStationRequestLabelType string

const (
	UpdateScanningStationRequestLabelTypeTag      UpdateScanningStationRequestLabelType = "tag"
	UpdateScanningStationRequestLabelTypeTraveler UpdateScanningStationRequestLabelType = "traveler"
)

// Operator requirement behavior for this station.
type UpdateScanningStationRequestOperatorRequirement string

const (
	UpdateScanningStationRequestOperatorRequirementNone          UpdateScanningStationRequestOperatorRequirement = "none"
	UpdateScanningStationRequestOperatorRequirementMaterialCheck UpdateScanningStationRequestOperatorRequirement = "material_check"
)

type OperationScanningStationDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationScanningStationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationScanningStationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationScanningStationNewParams struct {
	// Request to create a scanning station.
	CreateScanningStationRequest CreateScanningStationRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "department", "production_steps".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationScanningStationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateScanningStationRequest)
}
func (r *OperationScanningStationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationScanningStationNewParams]'s query parameters as
// `url.Values`.
func (r OperationScanningStationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationScanningStationGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "department", "production_steps".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationScanningStationGetParams]'s query parameters as
// `url.Values`.
func (r OperationScanningStationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationScanningStationUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "department", "production_steps".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a scanning station.
	UpdateScanningStationRequest UpdateScanningStationRequestParam
	paramObj
}

func (r OperationScanningStationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateScanningStationRequest)
}
func (r *OperationScanningStationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationScanningStationUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationScanningStationUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationScanningStationListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "department", "production_steps".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationScanningStationListParams]'s query parameters as
// `url.Values`.
func (r OperationScanningStationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
