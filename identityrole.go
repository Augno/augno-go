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

// List and manage roles.
//
// IdentityRoleService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityRoleService] method instead.
type IdentityRoleService struct {
	options []option.RequestOption
}

// NewIdentityRoleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityRoleService(opts ...option.RequestOption) (r IdentityRoleService) {
	r = IdentityRoleService{}
	r.options = opts
	return
}

// Creates a custom role with the specified permissions. Roles created through the
// API always have type `user`.
func (r *IdentityRoleService) New(ctx context.Context, params IdentityRoleNewParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/identity/roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a role by ID, including its permissions.
func (r *IdentityRoleService) Get(ctx context.Context, id string, query IdentityRoleGetParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/roles/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a custom role's name or permissions. Provided permissions
// replace all existing ones; global roles cannot be modified.
func (r *IdentityRoleService) Update(ctx context.Context, id string, params IdentityRoleUpdateParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/roles/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of roles for the target account, including global
// roles.
func (r *IdentityRoleService) List(ctx context.Context, query IdentityRoleListParams, opts ...option.RequestOption) (res *ListRole, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/identity/roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a role and its associated permissions. Global roles and roles currently
// assigned to one or more users cannot be deleted.
func (r *IdentityRoleService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityRoleDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/roles/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// CreateRoleRequest is a request to create a role.
//
// The properties Name, Permissions are required.
type CreateRoleRequestParam struct {
	// Display name for the role, unique within the account.
	Name string `json:"name" api:"required"`
	// Permissions to grant, in `{domain}:{action}` format, such as `customers:read`.
	//
	// The action must be one of `create`, `read`, `update`, or `delete`. Omit to
	// create a role with no permissions.
	Permissions []string `json:"permissions,omitzero" api:"required"`
	paramObj
}

func (r CreateRoleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateRoleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateRoleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List represents a paginated list of resources.
type ListRole struct {
	// Resources in this page.
	Data []Role `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListRoleObject `json:"object" api:"required"`
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
func (r ListRole) RawJSON() string { return r.JSON.raw }
func (r *ListRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListRoleObject string

const (
	ListRoleObjectList ListRoleObject = "list"
)

// UpdateRoleRequest is a request to update a role.
type UpdateRoleRequestParam struct {
	// New display name for the role, unique within the account. Omit to leave
	// unchanged.
	Name param.Opt[string] `json:"name,omitzero"`
	// Full replacement set of permissions, in `{domain}:{action}` format, such as
	// `customers:read`.
	//
	// Replaces all existing permissions on the role. Pass an empty array to remove all
	// permissions, or omit to leave them unchanged.
	Permissions []string `json:"permissions,omitzero"`
	paramObj
}

func (r UpdateRoleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateRoleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateRoleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityRoleDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityRoleDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityRoleDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityRoleNewParams struct {
	// CreateRoleRequest is a request to create a role.
	CreateRoleRequest CreateRoleRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r IdentityRoleNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateRoleRequest)
}
func (r *IdentityRoleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [IdentityRoleNewParams]'s query parameters as `url.Values`.
func (r IdentityRoleNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityRoleGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IdentityRoleGetParams]'s query parameters as `url.Values`.
func (r IdentityRoleGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityRoleUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "permissions".
	Include []string `query:"include,omitzero" json:"-"`
	// UpdateRoleRequest is a request to update a role.
	UpdateRoleRequest UpdateRoleRequestParam
	paramObj
}

func (r IdentityRoleUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateRoleRequest)
}
func (r *IdentityRoleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [IdentityRoleUpdateParams]'s query parameters as
// `url.Values`.
func (r IdentityRoleUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityRoleListParams struct {
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
	// Any of "owner", "owner.account", "permissions".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter results to roles whose type matches any of the given values.
	//
	// Any of "admin", "user", "scanner", "sales_rep", "agent".
	Types []string `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IdentityRoleListParams]'s query parameters as `url.Values`.
func (r IdentityRoleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
