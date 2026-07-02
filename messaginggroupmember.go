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
)

// Create and manage reusable rosters (named member sets) that seed many
// conversations.
//
// MessagingGroupMemberService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingGroupMemberService] method instead.
type MessagingGroupMemberService struct {
	options []option.RequestOption
}

// NewMessagingGroupMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingGroupMemberService(opts ...option.RequestOption) (r MessagingGroupMemberService) {
	r = MessagingGroupMemberService{}
	r.options = opts
	return
}

// Adds a member (a user or an agent) to a reusable roster.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingGroupMemberService) New(ctx context.Context, id string, body MessagingGroupMemberNewParams, opts ...option.RequestOption) (res *MessagingGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/groups/%s/members", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Removes a member from a reusable roster.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingGroupMemberService) Delete(ctx context.Context, memberID string, body MessagingGroupMemberDeleteParams, opts ...option.RequestOption) (res *MessagingGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/groups/%s/members/%s", url.PathEscape(body.ID), url.PathEscape(memberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to add a member to a reusable roster.
//
// The property MemberType is required.
type AddMessagingGroupMemberRequestParam struct {
	// The kind of member being added.
	//
	// Any of "user", "agent".
	MemberType AddMessagingGroupMemberRequestMemberType `json:"member_type,omitzero" api:"required"`
	// The account user to add (required when `member_type` is `user`).
	AccountUserID param.Opt[string] `json:"account_user_id,omitzero"`
	// The agent to add (required when `member_type` is `agent`).
	AgentConfigID param.Opt[string] `json:"agent_config_id,omitzero"`
	paramObj
}

func (r AddMessagingGroupMemberRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AddMessagingGroupMemberRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddMessagingGroupMemberRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of member being added.
type AddMessagingGroupMemberRequestMemberType string

const (
	AddMessagingGroupMemberRequestMemberTypeUser  AddMessagingGroupMemberRequestMemberType = "user"
	AddMessagingGroupMemberRequestMemberTypeAgent AddMessagingGroupMemberRequestMemberType = "agent"
)

type MessagingGroupMemberNewParams struct {
	// Request to add a member to a reusable roster.
	AddMessagingGroupMemberRequest AddMessagingGroupMemberRequestParam
	paramObj
}

func (r MessagingGroupMemberNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AddMessagingGroupMemberRequest)
}
func (r *MessagingGroupMemberNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingGroupMemberDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
