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

// The days a plant tenders freight and a customer's dock accepts it, less the
// holidays and shutdowns either side is closed for. Every ship-by date is resolved
// against them, so an order is never committed to a day nobody can act on.
//
// OperationOperatingCalendarClosureService contains methods and other services
// that help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationOperatingCalendarClosureService] method instead.
type OperationOperatingCalendarClosureService struct {
	options []option.RequestOption
}

// NewOperationOperatingCalendarClosureService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationOperatingCalendarClosureService(opts ...option.RequestOption) (r OperationOperatingCalendarClosureService) {
	r = OperationOperatingCalendarClosureService{}
	r.options = opts
	return
}

// Closes a calendar on a date.
//
// Every ship-by date resolved against this calendar afterwards walks past the
// closure: a carrier that does not move on Thanksgiving pushes the day an order
// has to leave earlier, and a plant shutdown does the same.
//
// Closing the same date twice is a no-op rather than an error, so re-seeding a
// year is safe and never renames a closure somebody has relabelled.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationOperatingCalendarClosureService) New(ctx context.Context, id string, body OperationOperatingCalendarClosureNewParams, opts ...option.RequestOption) (res *OperatingCalendarClosure, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/operating-calendars/%s/closures", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists the dates a calendar is shut, within a date window.
//
// Bounded rather than exhaustive: a calendar accumulates closures indefinitely,
// and the useful answer is the year either side of today. Widen it with
// `from_date` and `to_date`.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationOperatingCalendarClosureService) List(ctx context.Context, id string, query OperationOperatingCalendarClosureListParams, opts ...option.RequestOption) (res *ListOperatingCalendarClosure, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/operating-calendars/%s/closures", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Reopens a date the calendar was closed on.
//
// Used to drop a seeded holiday a plant actually works through. Orders already
// issued keep the dates they were stamped with.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationOperatingCalendarClosureService) Delete(ctx context.Context, closureID string, body OperationOperatingCalendarClosureDeleteParams, opts ...option.RequestOption) (res *OperationOperatingCalendarClosureDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if closureID == "" {
		err = errors.New("missing required closure_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/operating-calendars/%s/closures/%s", url.PathEscape(body.ID), url.PathEscape(closureID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to close a calendar on a date.
//
// The properties ClosedOn, Name are required.
type CreateOperatingCalendarClosureRequestParam struct {
	// The date nothing operates. Truncated to a day.
	ClosedOn time.Time `json:"closed_on" api:"required" format:"date-time"`
	// What the closure is, such as "Thanksgiving Day" or "Summer shutdown".
	Name string `json:"name" api:"required"`
	paramObj
}

func (r CreateOperatingCalendarClosureRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateOperatingCalendarClosureRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateOperatingCalendarClosureRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOperatingCalendarClosure struct {
	// Resources in this page.
	Data []OperatingCalendarClosure `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOperatingCalendarClosureObject `json:"object" api:"required"`
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
func (r ListOperatingCalendarClosure) RawJSON() string { return r.JSON.raw }
func (r *ListOperatingCalendarClosure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOperatingCalendarClosureObject string

const (
	ListOperatingCalendarClosureObjectList ListOperatingCalendarClosureObject = "list"
)

// OperatingCalendarClosure is one date a calendar is shut — a holiday, or a day of
// a shutdown week.
type OperatingCalendarClosure struct {
	// Unique identifier.
	ID string `json:"id" api:"required"`
	// The date nothing operates.
	ClosedOn time.Time `json:"closed_on" api:"required" format:"date-time"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// What the closure is, such as "Thanksgiving Day" or "Summer shutdown".
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "operating_calendar_closure".
	Object OperatingCalendarClosureObject `json:"object" api:"required"`
	// The calendar this closure belongs to.
	OperatingCalendarID string `json:"operating_calendar_id" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ClosedOn            respjson.Field
		CreatedAt           respjson.Field
		Name                respjson.Field
		Object              respjson.Field
		OperatingCalendarID respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatingCalendarClosure) RawJSON() string { return r.JSON.raw }
func (r *OperatingCalendarClosure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type OperatingCalendarClosureObject string

const (
	OperatingCalendarClosureObjectOperatingCalendarClosure OperatingCalendarClosureObject = "operating_calendar_closure"
)

type OperationOperatingCalendarClosureDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationOperatingCalendarClosureDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationOperatingCalendarClosureDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationOperatingCalendarClosureNewParams struct {
	// Request to close a calendar on a date.
	CreateOperatingCalendarClosureRequest CreateOperatingCalendarClosureRequestParam
	paramObj
}

func (r OperationOperatingCalendarClosureNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateOperatingCalendarClosureRequest)
}
func (r *OperationOperatingCalendarClosureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationOperatingCalendarClosureListParams struct {
	// Earliest closure date to return. Defaults to a year ago.
	FromDate param.Opt[time.Time] `query:"from_date,omitzero" format:"date-time" json:"-"`
	// Latest closure date to return. Defaults to a year ahead.
	ToDate param.Opt[time.Time] `query:"to_date,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [OperationOperatingCalendarClosureListParams]'s query
// parameters as `url.Values`.
func (r OperationOperatingCalendarClosureListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationOperatingCalendarClosureDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
