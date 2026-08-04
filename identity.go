// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
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

// List permission groups and their permissions.
//
// IdentityService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityService] method instead.
type IdentityService struct {
	options []option.RequestOption
	// List and manage account users.
	AccountUsers IdentityAccountUserService
	// Manage account details, branding, portal, logo, and favicon.
	Accounts IdentityAccountService
	// List and manage roles.
	Roles IdentityRoleService
}

// NewIdentityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityService(opts ...option.RequestOption) (r IdentityService) {
	r = IdentityService{}
	r.options = opts
	r.AccountUsers = NewIdentityAccountUserService(opts...)
	r.Accounts = NewIdentityAccountService(opts...)
	r.Roles = NewIdentityRoleService(opts...)
	return
}

// Lists the permission catalog, organized into groups of related permissions.
//
// Each group carries the individual permissions it covers; pair a permission's
// code with an action (`create`, `read`, `update`, or `delete`) to build the
// permission strings accepted when creating or updating a role. The catalog is
// platform-defined and identical for every account.
//
// This endpoint requires the permission: `permissions:read`.
func (r *IdentityService) GetPermissionGroups(ctx context.Context, query IdentityGetPermissionGroupsParams, opts ...option.RequestOption) (res *ListPermissionGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/identity/permission-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListPermission struct {
	// Resources in this page.
	Data []Permission `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPermissionObject `json:"object" api:"required"`
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
func (r ListPermission) RawJSON() string { return r.JSON.raw }
func (r *ListPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPermissionObject string

const (
	ListPermissionObjectList ListPermissionObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListPermissionGroup struct {
	// Resources in this page.
	Data []PermissionGroup `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPermissionGroupObject `json:"object" api:"required"`
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
func (r ListPermissionGroup) RawJSON() string { return r.JSON.raw }
func (r *ListPermissionGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPermissionGroupObject string

const (
	ListPermissionGroupObjectList ListPermissionGroupObject = "list"
)

// One area of the product that access can be granted for, such as customers,
// invoices, or production runs.
//
// A role never grants a permission outright; it grants specific actions on it,
// written as `{code}:{action}` — for example `customers:read`.
type Permission struct {
	// Permission ID.
	ID string `json:"id" api:"required"`
	// Stable code identifying the area this permission controls, such as `customers`
	// or `sales_orders`.
	//
	// Pair the code with an action (`create`, `read`, `update`, or `delete`) to form
	// the permission strings used when creating or updating a role.
	Code string `json:"code" api:"required"`
	// When the permission was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable description of what this permission controls.
	Description string `json:"description" api:"required"`
	// Code of the permission group this permission is listed under, such as
	// `inventory`.
	Group string `json:"group" api:"required"`
	// Human-readable name for the permission.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "permission".
	Object PermissionObject `json:"object" api:"required"`
	// When the permission was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Group       respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Permission) RawJSON() string { return r.JSON.raw }
func (r *Permission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PermissionObject string

const (
	PermissionObjectPermission PermissionObject = "permission"
)

// A category of the permission catalog that collects related permissions, such as
// inventory or invoices.
//
// Groups exist to organize the catalog for display; access is always granted by
// the individual permissions inside a group, never by the group itself.
type PermissionGroup struct {
	// Permission group ID.
	ID string `json:"id" api:"required"`
	// Unique code identifying the permission group, such as `customers`.
	Code string `json:"code" api:"required"`
	// When the permission group was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Free-form description of the permission group.
	Description string `json:"description" api:"required"`
	// Human-readable name for the permission group.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "permission_group".
	Object PermissionGroupObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Permissions ListPermission `json:"permissions" api:"required"`
	// When the permission group was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		Permissions respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PermissionGroup) RawJSON() string { return r.JSON.raw }
func (r *PermissionGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PermissionGroupObject string

const (
	PermissionGroupObjectPermissionGroup PermissionGroupObject = "permission_group"
)

type IdentityGetPermissionGroupsParams struct {
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

// URLQuery serializes [IdentityGetPermissionGroupsParams]'s query parameters as
// `url.Values`.
func (r IdentityGetPermissionGroupsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
