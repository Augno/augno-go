// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and retrieve priorities.
//
// SalePriorityService contains methods and other services that help with
// interacting with the openmrp API.
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

// Retrieves a single priority level by ID or by code.
//
// Looking one up by code is usually more convenient, because other resources refer
// to a priority by code rather than by ID.
//
// This endpoint requires the permission: `priorities:read`.
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

// Lists the priority levels that can be set on a sales order or purchase order.
//
// The levels are platform-provided and the same for every account, so the result
// is small and stable enough to cache. Results are ordered newest first rather
// than by urgency.
//
// This endpoint requires the permission: `priorities:read`.
func (r *SalePriorityService) List(ctx context.Context, query SalePriorityListParams, opts ...option.RequestOption) (res *ListPriority, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/priorities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListPriority struct {
	// Resources in this page.
	Data []Priority `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPriorityObject `json:"object" api:"required"`
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
func (r ListPriority) RawJSON() string { return r.JSON.raw }
func (r *ListPriority) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPriorityObject string

const (
	ListPriorityObjectList ListPriorityObject = "list"
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
