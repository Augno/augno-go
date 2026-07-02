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
)

// Add, remove, and manage participants (including agents) in a conversation.
//
// MessagingConversationParticipantActionService contains methods and other
// services that help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationParticipantActionService] method instead.
type MessagingConversationParticipantActionService struct {
	options []option.RequestOption
}

// NewMessagingConversationParticipantActionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewMessagingConversationParticipantActionService(opts ...option.RequestOption) (r MessagingConversationParticipantActionService) {
	r = MessagingConversationParticipantActionService{}
	r.options = opts
	return
}

// Changes a participant's role within a group conversation.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationParticipantActionService) SetRole(ctx context.Context, pid string, params MessagingConversationParticipantActionSetRoleParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/participants/%s/actions/set-role", url.PathEscape(params.ID), url.PathEscape(pid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to change a participant's role (owner only).
//
// The property Role is required.
type UpdateParticipantRoleRequestParam struct {
	// The role to assign to the participant.
	//
	// - `owner`: can rename or delete the conversation and manage members and roles.
	// - `admin`: can add and remove members and rename the conversation.
	// - `member`: can post, leave, mute, and react.
	// - `viewer`: read-only access.
	//
	// Any of "owner", "admin", "member", "viewer".
	Role UpdateParticipantRoleRequestRole `json:"role,omitzero" api:"required"`
	paramObj
}

func (r UpdateParticipantRoleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateParticipantRoleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateParticipantRoleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The role to assign to the participant.
//
// - `owner`: can rename or delete the conversation and manage members and roles.
// - `admin`: can add and remove members and rename the conversation.
// - `member`: can post, leave, mute, and react.
// - `viewer`: read-only access.
type UpdateParticipantRoleRequestRole string

const (
	UpdateParticipantRoleRequestRoleOwner  UpdateParticipantRoleRequestRole = "owner"
	UpdateParticipantRoleRequestRoleAdmin  UpdateParticipantRoleRequestRole = "admin"
	UpdateParticipantRoleRequestRoleMember UpdateParticipantRoleRequestRole = "member"
	UpdateParticipantRoleRequestRoleViewer UpdateParticipantRoleRequestRole = "viewer"
)

type MessagingConversationParticipantActionSetRoleParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Request to change a participant's role (owner only).
	UpdateParticipantRoleRequest UpdateParticipantRoleRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "participants", "topic", "last_message", "last_message.sender",
	// "last_message.author", "last_message.resource", "last_message.attachments".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationParticipantActionSetRoleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateParticipantRoleRequest)
}
func (r *MessagingConversationParticipantActionSetRoleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationParticipantActionSetRoleParams]'s
// query parameters as `url.Values`.
func (r MessagingConversationParticipantActionSetRoleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
