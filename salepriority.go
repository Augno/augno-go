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

// List and retrieve priorities.
//
// SalePriorityService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSalePriorityService] method instead.
type SalePriorityService struct {
	options []option.RequestOption
}

// NewSalePriorityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSalePriorityService(opts ...option.RequestOption) (r SalePriorityService) {
	r = SalePriorityService{}
	r.options = opts
	return
}

// Returns a priority by ID or code.
func (r *SalePriorityService) Get(ctx context.Context, id string, query SalePriorityGetParams, opts ...option.RequestOption) (res *Priority, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/priorities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of priorities.
func (r *SalePriorityService) List(ctx context.Context, query SalePriorityListParams, opts ...option.RequestOption) (res *ListPriority, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/priorities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List represents a paginated list of resources.
type ListPriority struct {
	// Resources in this page.
	Data []Priority `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPriorityObject `json:"object" api:"required"`
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
func (r ListPriority) RawJSON() string { return r.JSON.raw }
func (r *ListPriority) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPriorityObject string

const (
	ListPriorityObjectList ListPriorityObject = "list"
)

// Priority level used to order work on sales orders, purchase orders, and picks.
type Priority struct {
	// Priority ID.
	ID string `json:"id" api:"required"`
	// Machine-readable code identifying the priority level.
	//
	// - `low`: lowest urgency; worked after normal and high.
	// - `normal`: default urgency for most orders and picks.
	// - `high`: highest urgency; worked ahead of normal and low.
	//
	// Any of "low", "normal", "high".
	Code PriorityCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "priority".
	Object PriorityObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Priority) RawJSON() string { return r.JSON.raw }
func (r *Priority) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable code identifying the priority level.
//
// - `low`: lowest urgency; worked after normal and high.
// - `normal`: default urgency for most orders and picks.
// - `high`: highest urgency; worked ahead of normal and low.
type PriorityCode string

const (
	PriorityCodeLow    PriorityCode = "low"
	PriorityCodeNormal PriorityCode = "normal"
	PriorityCodeHigh   PriorityCode = "high"
)

// Resource type identifier.
type PriorityObject string

const (
	PriorityObjectPriority PriorityObject = "priority"
)

type SalePriorityGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SalePriorityGetParams]'s query parameters as `url.Values`.
func (r SalePriorityGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SalePriorityListParams struct {
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
	// Any of "owner".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SalePriorityListParams]'s query parameters as `url.Values`.
func (r SalePriorityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
