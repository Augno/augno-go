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

// List and manage account users.
//
// IdentityAccountUserService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityAccountUserService] method instead.
type IdentityAccountUserService struct {
	options []option.RequestOption
	// List and manage account users.
	Actions IdentityAccountUserActionService
}

// NewIdentityAccountUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIdentityAccountUserService(opts ...option.RequestOption) (r IdentityAccountUserService) {
	r = IdentityAccountUserService{}
	r.options = opts
	r.Actions = NewIdentityAccountUserActionService(opts...)
	return
}

// Creates a new account user and invites them to the target account.
func (r *IdentityAccountUserService) New(ctx context.Context, params IdentityAccountUserNewParams, opts ...option.RequestOption) (res *AccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/identity/account-users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns an account user by ID.
func (r *IdentityAccountUserService) Get(ctx context.Context, id string, query IdentityAccountUserGetParams, opts ...option.RequestOption) (res *AccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an account user.
func (r *IdentityAccountUserService) Update(ctx context.Context, id string, params IdentityAccountUserUpdateParams, opts ...option.RequestOption) (res *AccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of account users for the current account.
func (r *IdentityAccountUserService) List(ctx context.Context, query IdentityAccountUserListParams, opts ...option.RequestOption) (res *ListAccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/identity/account-users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request to create an account user.
type CreateAccountUserRequestParam struct {
	// Department assigned to the user.
	DepartmentID param.Opt[string] `json:"department_id,omitzero"`
	// User email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// User display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Password. Only used for scanner-role users (scanning stations). Must be 8–72
	// chars and include upper, lower, number, and special character.
	Password param.Opt[string] `json:"password,omitzero"`
	// Role assigned to the user.
	RoleID param.Opt[string] `json:"role_id,omitzero"`
	// Unique username (3–255 chars; letters, numbers, underscores, hyphens).
	Username param.Opt[string] `json:"username,omitzero"`
	// Notification preferences for the user (external targets only).
	Preferences []NotificationPreferenceItemParam `json:"preferences,omitzero"`
	paramObj
}

func (r CreateAccountUserRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAccountUserRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAccountUserRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List represents a paginated list of resources.
type ListAccountUser struct {
	// Resources in this page.
	Data []AccountUser `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAccountUserObject `json:"object" api:"required"`
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
func (r ListAccountUser) RawJSON() string { return r.JSON.raw }
func (r *ListAccountUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAccountUserObject string

const (
	ListAccountUserObjectList ListAccountUserObject = "list"
)

// NotificationPreferenceItem toggles a single account-relation notification type.
//
// The properties Enabled, NotificationType are required.
type NotificationPreferenceItemParam struct {
	// Whether this notification type is enabled for the account user.
	Enabled bool `json:"enabled" api:"required"`
	// Notification type.
	//
	// Any of "invoice", "order_acknowledgement", "purchase_order_submission".
	NotificationType NotificationPreferenceItemNotificationType `json:"notification_type,omitzero" api:"required"`
	paramObj
}

func (r NotificationPreferenceItemParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationPreferenceItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationPreferenceItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification type.
type NotificationPreferenceItemNotificationType string

const (
	NotificationPreferenceItemNotificationTypeInvoice                 NotificationPreferenceItemNotificationType = "invoice"
	NotificationPreferenceItemNotificationTypeOrderAcknowledgement    NotificationPreferenceItemNotificationType = "order_acknowledgement"
	NotificationPreferenceItemNotificationTypePurchaseOrderSubmission NotificationPreferenceItemNotificationType = "purchase_order_submission"
)

// Request to partially update an account user.
type UpdateAccountUserRequestParam struct {
	// Department assigned to the user.
	DepartmentID param.Opt[string] `json:"department_id,omitzero"`
	// Role assigned to the user.
	RoleID param.Opt[string] `json:"role_id,omitzero"`
	// User email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// User display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Unique username (3–255 chars; letters, numbers, underscores, hyphens).
	Username param.Opt[string] `json:"username,omitzero"`
	// Notification preferences to update (external targets only).
	Preferences []NotificationPreferenceItemParam `json:"preferences,omitzero"`
	paramObj
}

func (r UpdateAccountUserRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAccountUserRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAccountUserRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityAccountUserNewParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to create an account user.
	CreateAccountUserRequest CreateAccountUserRequestParam
	paramObj
}

func (r IdentityAccountUserNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAccountUserRequest)
}
func (r *IdentityAccountUserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [IdentityAccountUserNewParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityAccountUserGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IdentityAccountUserGetParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityAccountUserUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update an account user.
	UpdateAccountUserRequest UpdateAccountUserRequestParam
	paramObj
}

func (r IdentityAccountUserUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAccountUserRequest)
}
func (r *IdentityAccountUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [IdentityAccountUserUpdateParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityAccountUserListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	// Controls whether removed account users are included.
	//
	// Any of "excluded", "included".
	RemovedScope IdentityAccountUserListParamsRemovedScope `query:"removed_scope,omitzero" json:"-"`
	// Filter by role type code.
	//
	// Any of "admin", "user", "scanner", "sales_rep", "agent".
	RoleType IdentityAccountUserListParamsRoleType `query:"role_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IdentityAccountUserListParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Controls whether removed account users are included.
type IdentityAccountUserListParamsRemovedScope string

const (
	IdentityAccountUserListParamsRemovedScopeExcluded IdentityAccountUserListParamsRemovedScope = "excluded"
	IdentityAccountUserListParamsRemovedScopeIncluded IdentityAccountUserListParamsRemovedScope = "included"
)

// Filter by role type code.
type IdentityAccountUserListParamsRoleType string

const (
	IdentityAccountUserListParamsRoleTypeAdmin    IdentityAccountUserListParamsRoleType = "admin"
	IdentityAccountUserListParamsRoleTypeUser     IdentityAccountUserListParamsRoleType = "user"
	IdentityAccountUserListParamsRoleTypeScanner  IdentityAccountUserListParamsRoleType = "scanner"
	IdentityAccountUserListParamsRoleTypeSalesRep IdentityAccountUserListParamsRoleType = "sales_rep"
	IdentityAccountUserListParamsRoleTypeAgent    IdentityAccountUserListParamsRoleType = "agent"
)
