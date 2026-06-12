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

// List and manage payment terms.
//
// FinancePaymentTermService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancePaymentTermService] method instead.
type FinancePaymentTermService struct {
	options []option.RequestOption
}

// NewFinancePaymentTermService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFinancePaymentTermService(opts ...option.RequestOption) (r FinancePaymentTermService) {
	r = FinancePaymentTermService{}
	r.options = opts
	return
}

// Creates a payment term.
//
// The new term is owned by your account and starts with status `active`.
func (r *FinancePaymentTermService) New(ctx context.Context, params FinancePaymentTermNewParams, opts ...option.RequestOption) (res *PaymentTerm, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/finance/payment-terms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a payment term by ID.
func (r *FinancePaymentTermService) Get(ctx context.Context, id string, query FinancePaymentTermGetParams, opts ...option.RequestOption) (res *PaymentTerm, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/finance/payment-terms/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a payment term.
//
// Only payment terms created by your account can be updated; system-owned default
// terms cannot be.
func (r *FinancePaymentTermService) Update(ctx context.Context, id string, params FinancePaymentTermUpdateParams, opts ...option.RequestOption) (res *PaymentTerm, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/finance/payment-terms/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of payment terms.
//
// The list includes both payment terms created by your account and Augno-provided
// system defaults.
func (r *FinancePaymentTermService) List(ctx context.Context, query FinancePaymentTermListParams, opts ...option.RequestOption) (res *ListPaymentTerm, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/finance/payment-terms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a payment term.
//
// Only payment terms created by your account can be deleted; system-owned default
// terms cannot be.
func (r *FinancePaymentTermService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FinancePaymentTermDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/finance/payment-terms/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a payment term.
//
// The property Name is required.
type CreatePaymentTermRequestParam struct {
	// Display name (e.g. `Net 30`).
	//
	// Must be unique among the payment terms visible to your account, including system
	// defaults.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r CreatePaymentTermRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePaymentTermRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePaymentTermRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List represents a paginated list of resources.
type ListPaymentTerm struct {
	// Resources in this page.
	Data []PaymentTerm `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPaymentTermObject `json:"object" api:"required"`
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
func (r ListPaymentTerm) RawJSON() string { return r.JSON.raw }
func (r *ListPaymentTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPaymentTermObject string

const (
	ListPaymentTermObjectList ListPaymentTermObject = "list"
)

// Request to partially update a payment term.
type UpdatePaymentTermRequestParam struct {
	// New display name for the payment term.
	//
	// Must be unique among the payment terms visible to your account, including system
	// defaults.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r UpdatePaymentTermRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdatePaymentTermRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdatePaymentTermRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinancePaymentTermDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FinancePaymentTermDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FinancePaymentTermDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinancePaymentTermNewParams struct {
	// Request to create a payment term.
	CreatePaymentTermRequest CreatePaymentTermRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r FinancePaymentTermNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePaymentTermRequest)
}
func (r *FinancePaymentTermNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FinancePaymentTermNewParams]'s query parameters as
// `url.Values`.
func (r FinancePaymentTermNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FinancePaymentTermGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FinancePaymentTermGetParams]'s query parameters as
// `url.Values`.
func (r FinancePaymentTermGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FinancePaymentTermUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a payment term.
	UpdatePaymentTermRequest UpdatePaymentTermRequestParam
	paramObj
}

func (r FinancePaymentTermUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdatePaymentTermRequest)
}
func (r *FinancePaymentTermUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FinancePaymentTermUpdateParams]'s query parameters as
// `url.Values`.
func (r FinancePaymentTermUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FinancePaymentTermListParams struct {
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
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FinancePaymentTermListParams]'s query parameters as
// `url.Values`.
func (r FinancePaymentTermListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
