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
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List and manage locations.
//
// OperationLocationTypeService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationLocationTypeService] method instead.
type OperationLocationTypeService struct {
	options []option.RequestOption
}

// NewOperationLocationTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationLocationTypeService(opts ...option.RequestOption) (r OperationLocationTypeService) {
	r = OperationLocationTypeService{}
	r.options = opts
	return
}

// Returns a location type by ID or code.
func (r *OperationLocationTypeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LocationType, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/location-types/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of location types.
func (r *OperationLocationTypeService) List(ctx context.Context, query OperationLocationTypeListParams, opts ...option.RequestOption) (res *ListLocationType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/location-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List represents a paginated list of resources.
type ListLocationType struct {
	// Resources in this page.
	Data []LocationType `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListLocationTypeObject `json:"object" api:"required"`
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
func (r ListLocationType) RawJSON() string { return r.JSON.raw }
func (r *ListLocationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListLocationTypeObject string

const (
	ListLocationTypeObjectList ListLocationTypeObject = "list"
)

// A level in the storage location hierarchy, such as a building or a bin.
type LocationType struct {
	// Location type ID.
	ID string `json:"id" api:"required"`
	// Location type code, identifying the level of the storage hierarchy this type
	// represents.
	//
	// - `building`: a building-level location.
	// - `section`: a section within a building.
	// - `aisle`: an aisle within a section.
	// - `rack`: a rack within an aisle.
	// - `shelf`: a shelf within a rack.
	// - `bin`: a bin within a shelf.
	//
	// Any of "building", "section", "aisle", "rack", "shelf", "bin".
	Code LocationTypeCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the location type.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "location_type".
	Object LocationTypeObject `json:"object" api:"required"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationType) RawJSON() string { return r.JSON.raw }
func (r *LocationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type LocationTypeObject string

const (
	LocationTypeObjectLocationType LocationTypeObject = "location_type"
)

type OperationLocationTypeListParams struct {
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
	paramObj
}

// URLQuery serializes [OperationLocationTypeListParams]'s query parameters as
// `url.Values`.
func (r OperationLocationTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
