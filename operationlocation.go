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

// List and manage locations.
//
// OperationLocationService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationLocationService] method instead.
type OperationLocationService struct {
	options []option.RequestOption
}

// NewOperationLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationLocationService(opts ...option.RequestOption) (r OperationLocationService) {
	r = OperationLocationService{}
	r.options = opts
	return
}

// Creates a storage location, optionally placing it in the location hierarchy.
func (r *OperationLocationService) New(ctx context.Context, params OperationLocationNewParams, opts ...option.RequestOption) (res *Location, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a location by ID.
func (r *OperationLocationService) Get(ctx context.Context, id string, query OperationLocationGetParams, opts ...option.RequestOption) (res *Location, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/locations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a location.
func (r *OperationLocationService) Update(ctx context.Context, id string, params OperationLocationUpdateParams, opts ...option.RequestOption) (res *Location, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/locations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of locations in your account.
func (r *OperationLocationService) List(ctx context.Context, query OperationLocationListParams, opts ...option.RequestOption) (res *ListLocation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a location.
//
// Fails if the location has child locations; remove or reassign the children
// first.
func (r *OperationLocationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationLocationDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/locations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a location.
//
// The properties Name, Type are required.
type CreateLocationRequestParam struct {
	// Display name of the location.
	//
	// Maximum 255 characters.
	Name string `json:"name" api:"required"`
	// Location type code, identifying this location's level in the storage hierarchy.
	//
	// - `building`: a building-level location.
	// - `section`: a section within a building.
	// - `aisle`: an aisle within a section.
	// - `rack`: a rack within an aisle.
	// - `shelf`: a shelf within a rack.
	// - `bin`: a bin within a shelf.
	//
	// Any of "building", "section", "aisle", "rack", "shelf", "bin".
	Type LocationTypeCode `json:"type,omitzero" api:"required"`
	// ID of the parent location.
	//
	// Omit for top-level locations.
	ParentID param.Opt[string] `json:"parent_id,omitzero"`
	// IDs of existing locations to attach as children of the new location.
	//
	// Listed locations are moved from their current parent, if any.
	ChildIDs []string `json:"child_ids,omitzero"`
	paramObj
}

func (r CreateLocationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateLocationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateLocationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to partially update a location.
type UpdateLocationRequestParam struct {
	// ID of the parent location.
	//
	// Send `null` to clear the parent and make this a top-level location.
	ParentID param.Opt[string] `json:"parent_id,omitzero"`
	// Display name of the location.
	//
	// Maximum 255 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// IDs of locations to set as this location's children.
	//
	// When provided, replaces the full set of children: current children not listed
	// are detached, and listed locations are moved from their current parent. Send
	// `null` to detach all children.
	ChildIDs []string `json:"child_ids,omitzero"`
	// Location type code, identifying this location's level in the storage hierarchy.
	//
	// - `building`: a building-level location.
	// - `section`: a section within a building.
	// - `aisle`: an aisle within a section.
	// - `rack`: a rack within an aisle.
	// - `shelf`: a shelf within a rack.
	// - `bin`: a bin within a shelf.
	//
	// Any of "building", "section", "aisle", "rack", "shelf", "bin".
	Type LocationTypeCode `json:"type,omitzero"`
	paramObj
}

func (r UpdateLocationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateLocationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateLocationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationLocationDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationLocationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationLocationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationLocationNewParams struct {
	// Request to create a location.
	CreateLocationRequest CreateLocationRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "parent", "children".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationLocationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateLocationRequest)
}
func (r *OperationLocationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationLocationNewParams]'s query parameters as
// `url.Values`.
func (r OperationLocationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationLocationGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "parent", "children".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationLocationGetParams]'s query parameters as
// `url.Values`.
func (r OperationLocationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationLocationUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "parent", "children".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a location.
	UpdateLocationRequest UpdateLocationRequestParam
	paramObj
}

func (r OperationLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLocationRequest)
}
func (r *OperationLocationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationLocationUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationLocationUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationLocationListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "parent", "children".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationLocationListParams]'s query parameters as
// `url.Values`.
func (r OperationLocationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
