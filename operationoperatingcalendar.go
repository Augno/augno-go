// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// The days a plant tenders freight and a customer's dock accepts it, less the
// holidays and shutdowns either side is closed for. Every ship-by date is resolved
// against them, so an order is never committed to a day nobody can act on.
//
// OperationOperatingCalendarService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationOperatingCalendarService] method instead.
type OperationOperatingCalendarService struct {
	options []option.RequestOption
	// The days a plant tenders freight and a customer's dock accepts it, less the
	// holidays and shutdowns either side is closed for. Every ship-by date is resolved
	// against them, so an order is never committed to a day nobody can act on.
	Closures OperationOperatingCalendarClosureService
}

// NewOperationOperatingCalendarService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationOperatingCalendarService(opts ...option.RequestOption) (r OperationOperatingCalendarService) {
	r = OperationOperatingCalendarService{}
	r.options = opts
	r.Closures = NewOperationOperatingCalendarClosureService(opts...)
	return
}

// Creates an operating calendar.
//
// The days govern every ship-by date resolved against this calendar: a plant that
// tenders freight Monday to Thursday never gets committed to a Friday shipment,
// and a customer's promised delivery date is worked back from a day they can
// actually receive on.
//
// A calendar starts with no closures. Add holidays and shutdowns to it separately.
//
// This endpoint requires the permission: `production_schedules:create`.
func (r *OperationOperatingCalendarService) New(ctx context.Context, body OperationOperatingCalendarNewParams, opts ...option.RequestOption) (res *OperatingCalendar, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/operating-calendars"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves one operating calendar by ID.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationOperatingCalendarService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OperatingCalendar, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/operating-calendars/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an operating calendar.
//
// A calendar's kind cannot change: a shipping calendar that became a receiving one
// would silently drop the pickup cutoff every commitment resolved against it
// depends on. Create a second calendar instead.
//
// Changes apply to commitments made from now on. Orders already issued keep the
// dates they were stamped with, so adding a holiday never retroactively makes a
// past order late.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationOperatingCalendarService) Update(ctx context.Context, id string, body OperationOperatingCalendarUpdateParams, opts ...option.RequestOption) (res *OperatingCalendar, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/operating-calendars/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists the operating calendars configured for the account.
//
// Both kinds are returned unless `kind` narrows it, ordered with each kind's
// default first.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationOperatingCalendarService) List(ctx context.Context, query OperationOperatingCalendarListParams, opts ...option.RequestOption) (res *ListOperatingCalendar, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/operating-calendars"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an operating calendar.
//
// Refused while any address, customer, customer group, or account setting still
// points at it. Deleting a calendar out from under its references would quietly
// return every affected order to a plain Monday-to-Friday week, which reads as the
// feature breaking rather than as a decision anybody made — so re-point them
// first.
//
// This endpoint requires the permission: `production_schedules:delete`.
func (r *OperationOperatingCalendarService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationOperatingCalendarDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/operating-calendars/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create an operating calendar.
//
// The properties Code, DaysOfWeek, Kind, Name are required.
type CreateOperatingCalendarRequestParam struct {
	// Short stable identifier, unique per account.
	Code string `json:"code" api:"required"`
	// Open weekdays as seven characters of '0' or '1', Monday first. "1111100" is
	// Monday to Friday; "1111000" is a Monday-to-Thursday plant. At least one day must
	// be open.
	DaysOfWeek string `json:"days_of_week" api:"required"`
	// Which side of a shipment this calendar describes.
	//
	// Any of "ship", "receive".
	Kind CreateOperatingCalendarRequestKind `json:"kind,omitzero" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Local time freight has to be tendered by, as "15:00". Only a shipping calendar
	// accepts one.
	CutoffAt param.Opt[string] `json:"cutoff_at,omitzero"`
	// Make this the calendar used when nothing more specific is linked. Setting it
	// demotes whichever calendar of the same kind held the role.
	IsDefault param.Opt[bool] `json:"is_default,omitzero"`
	// IANA zone the cutoff is read in, such as "America/Chicago". On a receiving
	// calendar, leave it unset to take the zone from the ship-to address.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	paramObj
}

func (r CreateOperatingCalendarRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateOperatingCalendarRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateOperatingCalendarRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which side of a shipment this calendar describes.
type CreateOperatingCalendarRequestKind string

const (
	CreateOperatingCalendarRequestKindShip    CreateOperatingCalendarRequestKind = "ship"
	CreateOperatingCalendarRequestKindReceive CreateOperatingCalendarRequestKind = "receive"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOperatingCalendar struct {
	// Resources in this page.
	Data []OperatingCalendar `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOperatingCalendarObject `json:"object" api:"required"`
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
func (r ListOperatingCalendar) RawJSON() string { return r.JSON.raw }
func (r *ListOperatingCalendar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOperatingCalendarObject string

const (
	ListOperatingCalendarObjectList ListOperatingCalendarObject = "list"
)

// OperatingCalendar is the set of days one party to a shipment operates.
//
// A `ship` calendar is the plant tendering freight to a carrier; a `receive`
// calendar is a customer's dock accepting it. Ship-by dates are resolved against
// both, so an order is never committed to a day nobody can act on.
type OperatingCalendar struct {
	// Unique identifier.
	ID string `json:"id" api:"required"`
	// Short stable identifier, unique per account.
	Code string `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Local time freight has to be tendered by, as "15:00". Only a shipping calendar
	// carries one.
	CutoffAt string `json:"cutoff_at" api:"required"`
	// Open weekdays as seven characters of '0' or '1', Monday first. "1111100" is
	// Monday to Friday; "1111000" is a Monday-to-Thursday plant.
	DaysOfWeek string `json:"days_of_week" api:"required"`
	// Whether this is the calendar used when nothing more specific is linked. Exactly
	// one per kind.
	IsDefault bool `json:"is_default" api:"required"`
	// Which side of a shipment this calendar describes.
	//
	// Any of "ship", "receive".
	Kind OperatingCalendarKind `json:"kind" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "operating_calendar".
	Object OperatingCalendarObject `json:"object" api:"required"`
	// IANA zone the cutoff is read in. Null on a receiving calendar means it is taken
	// from the ship-to address.
	Timezone string `json:"timezone" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		CutoffAt    respjson.Field
		DaysOfWeek  respjson.Field
		IsDefault   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Timezone    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatingCalendar) RawJSON() string { return r.JSON.raw }
func (r *OperatingCalendar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which side of a shipment this calendar describes.
type OperatingCalendarKind string

const (
	OperatingCalendarKindShip    OperatingCalendarKind = "ship"
	OperatingCalendarKindReceive OperatingCalendarKind = "receive"
)

// Resource type identifier.
type OperatingCalendarObject string

const (
	OperatingCalendarObjectOperatingCalendar OperatingCalendarObject = "operating_calendar"
)

// Request to update an operating calendar.
type UpdateOperatingCalendarRequestParam struct {
	// Local time freight has to be tendered by. Clearing it leaves the ship-by date a
	// day with no time of day attached.
	CutoffAt param.Opt[string] `json:"cutoff_at,omitzero"`
	// IANA zone the cutoff is read in. Clearing it on a receiving calendar returns to
	// taking the zone from the ship-to address.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// Open weekdays as seven characters of '0' or '1', Monday first. At least one day
	// must be open.
	DaysOfWeek param.Opt[string] `json:"days_of_week,omitzero"`
	// Make this the calendar used when nothing more specific is linked. Setting it
	// demotes whichever calendar of the same kind held the role.
	IsDefault param.Opt[bool] `json:"is_default,omitzero"`
	// Human-readable name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r UpdateOperatingCalendarRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateOperatingCalendarRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateOperatingCalendarRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationOperatingCalendarDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationOperatingCalendarDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationOperatingCalendarDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationOperatingCalendarNewParams struct {
	// Request to create an operating calendar.
	CreateOperatingCalendarRequest CreateOperatingCalendarRequestParam
	paramObj
}

func (r OperationOperatingCalendarNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateOperatingCalendarRequest)
}
func (r *OperationOperatingCalendarNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationOperatingCalendarUpdateParams struct {
	// Request to update an operating calendar.
	UpdateOperatingCalendarRequest UpdateOperatingCalendarRequestParam
	paramObj
}

func (r OperationOperatingCalendarUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateOperatingCalendarRequest)
}
func (r *OperationOperatingCalendarUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationOperatingCalendarListParams struct {
	// Return only shipping or only receiving calendars.
	//
	// Any of "ship", "receive".
	Kind OperationOperatingCalendarListParamsKind `query:"kind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationOperatingCalendarListParams]'s query parameters as
// `url.Values`.
func (r OperationOperatingCalendarListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Return only shipping or only receiving calendars.
type OperationOperatingCalendarListParamsKind string

const (
	OperationOperatingCalendarListParamsKindShip    OperationOperatingCalendarListParamsKind = "ship"
	OperationOperatingCalendarListParamsKindReceive OperationOperatingCalendarListParamsKind = "receive"
)
