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
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// Create and manage reusable rosters (named member sets) that seed many
// conversations.
//
// MessagingGroupService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingGroupService] method instead.
type MessagingGroupService struct {
	options []option.RequestOption
	// Create and manage reusable rosters (named member sets) that seed many
	// conversations.
	Members MessagingGroupMemberService
}

// NewMessagingGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagingGroupService(opts ...option.RequestOption) (r MessagingGroupService) {
	r = MessagingGroupService{}
	r.options = opts
	r.Members = NewMessagingGroupMemberService(opts...)
	return
}

// Creates a reusable roster of members (users and/or agents) that can seed many
// conversations.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingGroupService) New(ctx context.Context, body MessagingGroupNewParams, opts ...option.RequestOption) (res *MessagingGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a reusable roster (with its members).
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingGroupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Renames a reusable roster.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingGroupService) Update(ctx context.Context, id string, body MessagingGroupUpdateParams, opts ...option.RequestOption) (res *MessagingGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists the reusable rosters in the caller's account (most-recently-updated
// first).
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingGroupService) List(ctx context.Context, opts ...option.RequestOption) (res *ListMessagingGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a reusable roster.
//
// Conversations already started from it are unaffected (their members were
// snapshotted); they simply lose the roster reference.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingGroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingGroupDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a reusable roster.
//
// The property Name is required.
type CreateMessagingGroupRequestParam struct {
	// The roster's display name.
	Name string `json:"name" api:"required"`
	// The account users to include in the roster.
	MemberAccountUserIDs []string `json:"member_account_user_ids,omitzero"`
	// The agents to include in the roster.
	MemberAgentConfigIDs []string `json:"member_agent_config_ids,omitzero"`
	paramObj
}

func (r CreateMessagingGroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMessagingGroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMessagingGroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List represents a paginated list of resources.
type ListMessagingGroup struct {
	// Resources in this page.
	Data []MessagingGroup `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessagingGroupObject `json:"object" api:"required"`
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
func (r ListMessagingGroup) RawJSON() string { return r.JSON.raw }
func (r *ListMessagingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessagingGroupObject string

const (
	ListMessagingGroupObjectList ListMessagingGroupObject = "list"
)

// Request to rename a reusable roster.
//
// The property Name is required.
type UpdateMessagingGroupRequestParam struct {
	// The roster's new display name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r UpdateMessagingGroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMessagingGroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMessagingGroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingGroupDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingGroupDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingGroupDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingGroupNewParams struct {
	// Request to create a reusable roster.
	CreateMessagingGroupRequest CreateMessagingGroupRequestParam
	paramObj
}

func (r MessagingGroupNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMessagingGroupRequest)
}
func (r *MessagingGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingGroupUpdateParams struct {
	// Request to rename a reusable roster.
	UpdateMessagingGroupRequest UpdateMessagingGroupRequestParam
	paramObj
}

func (r MessagingGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateMessagingGroupRequest)
}
func (r *MessagingGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
