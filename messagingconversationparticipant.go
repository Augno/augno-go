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

// Add, remove, and manage participants (including agents) in a conversation.
//
// MessagingConversationParticipantService contains methods and other services that
// help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationParticipantService] method instead.
type MessagingConversationParticipantService struct {
	options []option.RequestOption
	// Add, remove, and manage participants (including agents) in a conversation.
	Actions MessagingConversationParticipantActionService
}

// NewMessagingConversationParticipantService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingConversationParticipantService(opts ...option.RequestOption) (r MessagingConversationParticipantService) {
	r = MessagingConversationParticipantService{}
	r.options = opts
	r.Actions = NewMessagingConversationParticipantActionService(opts...)
	return
}

// Adds (or reactivates) a participant in a group conversation.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingConversationParticipantService) New(ctx context.Context, id string, params MessagingConversationParticipantNewParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/participants", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes a participant from a group conversation.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationParticipantService) Delete(ctx context.Context, pid string, body MessagingConversationParticipantDeleteParams, opts ...option.RequestOption) (res *MessagingConversationParticipantDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if pid == "" {
		err = errors.New("missing required pid parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/participants/%s", url.PathEscape(body.ID), url.PathEscape(pid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to add a participant to a group (owner/admin).
//
// The property AccountUserID is required.
type AddParticipantRequestParam struct {
	// The account user to add.
	AccountUserID string `json:"account_user_id" api:"required"`
	// Role to grant the new participant.
	//
	// - `admin`: can add and remove members and rename the conversation.
	// - `member`: can post, leave, mute, and react.
	// - `viewer`: read-only access.
	//
	// `owner` is not accepted here; ownership can only be transferred via the set-role
	// endpoint.
	//
	// Any of "owner", "admin", "member", "viewer".
	Role AddParticipantRequestRole `json:"role,omitzero"`
	paramObj
}

func (r AddParticipantRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AddParticipantRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddParticipantRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role to grant the new participant.
//
// - `admin`: can add and remove members and rename the conversation.
// - `member`: can post, leave, mute, and react.
// - `viewer`: read-only access.
//
// `owner` is not accepted here; ownership can only be transferred via the set-role
// endpoint.
type AddParticipantRequestRole string

const (
	AddParticipantRequestRoleOwner  AddParticipantRequestRole = "owner"
	AddParticipantRequestRoleAdmin  AddParticipantRequestRole = "admin"
	AddParticipantRequestRoleMember AddParticipantRequestRole = "member"
	AddParticipantRequestRoleViewer AddParticipantRequestRole = "viewer"
)

type MessagingConversationParticipantDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingConversationParticipantDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingConversationParticipantDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingConversationParticipantNewParams struct {
	// Request to add a participant to a group (owner/admin).
	AddParticipantRequest AddParticipantRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "participants", "topic", "last_message", "last_message.sender",
	// "last_message.author", "last_message.resource", "last_message.attachments".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationParticipantNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AddParticipantRequest)
}
func (r *MessagingConversationParticipantNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationParticipantNewParams]'s query
// parameters as `url.Values`.
func (r MessagingConversationParticipantNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationParticipantDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
