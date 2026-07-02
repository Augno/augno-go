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

// List and manage sandbox environments.
//
// CoreSandboxService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreSandboxService] method instead.
type CoreSandboxService struct {
	options []option.RequestOption
}

// NewCoreSandboxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreSandboxService(opts ...option.RequestOption) (r CoreSandboxService) {
	r = CoreSandboxService{}
	r.options = opts
	return
}

// Creates a sandbox account owned by your production account.
//
// When `mode` is `seeded`, sample data is populated asynchronously and may not be
// available immediately after the sandbox is created. Sandboxes cannot be created
// while acting in a sandbox.
//
// This endpoint requires the permission: `sandboxes:create`.
func (r *CoreSandboxService) New(ctx context.Context, params CoreSandboxNewParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a sandbox by ID.
//
// This endpoint requires the permission: `sandboxes:read`.
func (r *CoreSandboxService) Get(ctx context.Context, id string, query CoreSandboxGetParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/sandboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of sandboxes.
//
// This endpoint requires the permission: `sandboxes:read`.
func (r *CoreSandboxService) List(ctx context.Context, query CoreSandboxListParams, opts ...option.RequestOption) (res *ListSandbox, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a sandbox account.
//
// The sandbox's data is purged asynchronously, so it may persist briefly after
// this call returns.
//
// This endpoint requires the permission: `sandboxes:delete`.
func (r *CoreSandboxService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CoreSandboxDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/sandboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a sandbox.
//
// The property Name is required.
type CreateSandboxRequestParam struct {
	// Display name of the sandbox.
	Name string `json:"name" api:"required"`
	// Controls how the sandbox is initialized.
	//
	//   - `blank`: starts empty, with no pre-populated data.
	//   - `seeded`: starts with sample data, populated asynchronously after the sandbox
	//     is created.
	//
	// Any of "blank", "seeded".
	Mode CreateSandboxRequestMode `json:"mode,omitzero"`
	paramObj
}

func (r CreateSandboxRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSandboxRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSandboxRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls how the sandbox is initialized.
//
//   - `blank`: starts empty, with no pre-populated data.
//   - `seeded`: starts with sample data, populated asynchronously after the sandbox
//     is created.
type CreateSandboxRequestMode string

const (
	CreateSandboxRequestModeBlank  CreateSandboxRequestMode = "blank"
	CreateSandboxRequestModeSeeded CreateSandboxRequestMode = "seeded"
)

// List represents a paginated list of resources.
type ListSandbox struct {
	// Resources in this page.
	Data []Sandbox `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSandboxObject `json:"object" api:"required"`
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
func (r ListSandbox) RawJSON() string { return r.JSON.raw }
func (r *ListSandbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSandboxObject string

const (
	ListSandboxObjectList ListSandboxObject = "list"
)

// Sandbox account for isolated testing.
type Sandbox struct {
	// Sandbox ID.
	ID string `json:"id" api:"required"`
	// When this sandbox was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the sandbox.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "sandbox".
	Object SandboxObject `json:"object" api:"required"`
	// A customer account, including its branding and customer portal sub-resources.
	OwnerAccount Account `json:"owner_account" api:"required"`
	// When this sandbox was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Object       respjson.Field
		OwnerAccount respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Sandbox) RawJSON() string { return r.JSON.raw }
func (r *Sandbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SandboxObject string

const (
	SandboxObjectSandbox SandboxObject = "sandbox"
)

type CoreSandboxDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CoreSandboxDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CoreSandboxDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CoreSandboxNewParams struct {
	// Request to create a sandbox.
	CreateSandboxRequest CreateSandboxRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner_account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CoreSandboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSandboxRequest)
}
func (r *CoreSandboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CoreSandboxNewParams]'s query parameters as `url.Values`.
func (r CoreSandboxNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CoreSandboxGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner_account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreSandboxGetParams]'s query parameters as `url.Values`.
func (r CoreSandboxGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CoreSandboxListParams struct {
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
	// Any of "owner_account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreSandboxListParams]'s query parameters as `url.Values`.
func (r CoreSandboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
