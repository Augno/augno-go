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
)

// Create conversations, send and read messages (1:1 direct messages).
//
// MessagingConversationActionService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationActionService] method instead.
type MessagingConversationActionService struct {
	options []option.RequestOption
}

// NewMessagingConversationActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingConversationActionService(opts ...option.RequestOption) (r MessagingConversationActionService) {
	r = MessagingConversationActionService{}
	r.options = opts
	return
}

// Archives a conversation at the account level so it drops out of active lists for
// everyone until it is unarchived.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Archive(ctx context.Context, id string, body MessagingConversationActionArchiveParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/archive", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Assigns an external customer-service case to an owner — a user or a team — or
// clears the assignment.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Assign(ctx context.Context, id string, params MessagingConversationActionAssignParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/assign", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Hides a conversation from the caller's own list without affecting other
// participants.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Hide(ctx context.Context, id string, body MessagingConversationActionHideParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/hide", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Removes the caller from a conversation, marking their membership as left.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Leave(ctx context.Context, id string, body MessagingConversationActionLeaveParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/leave", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Mutes a conversation for the calling actor.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Mute(ctx context.Context, id string, params MessagingConversationActionMuteParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/mute", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Advances the caller's read cursor and returns the refreshed conversation (with
// new unread count).
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Read(ctx context.Context, id string, params MessagingConversationActionReadParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/read", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Permanently redacts the content of every message in a conversation (GDPR
// right-to-erasure).
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingConversationActionService) Redact(ctx context.Context, id string, body MessagingConversationActionRedactParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/redact", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Files an abuse report against a conversation (optionally a specific message) and
// returns the conversation.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingConversationActionService) Report(ctx context.Context, id string, params MessagingConversationActionReportParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/report", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Places a conversation under legal hold or releases it.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) SetLegalHold(ctx context.Context, id string, params MessagingConversationActionSetLegalHoldParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/set-legal-hold", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sets the triage lane of an external customer-service case.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) SetStatus(ctx context.Context, id string, params MessagingConversationActionSetStatusParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/set-status", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns an archived conversation to the active state so it appears in active
// lists again.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Unarchive(ctx context.Context, id string, body MessagingConversationActionUnarchiveParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/unarchive", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Restores a conversation the caller had hidden back to their own list.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Unhide(ctx context.Context, id string, body MessagingConversationActionUnhideParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/unhide", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Restores notifications for a conversation the caller had muted.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationActionService) Unmute(ctx context.Context, id string, body MessagingConversationActionUnmuteParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/actions/unmute", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to assign a customer-service case to a single owner — a user or a team.
//
// The owner is a polymorphic (`assignee_resource_type`, `assignee_resource_id`)
// reference; omit both fields to clear the assignment.
type AssignConversationRequestParam struct {
	// The owner's id, an `account_user` or `account_group` matching
	// `assignee_resource_type`.
	//
	// Omit this and `assignee_resource_type` to clear the assignment.
	AssigneeResourceID param.Opt[string] `json:"assignee_resource_id,omitzero"`
	// The owner's resource type: `account_user` (a teammate) or `account_group` (a
	// team).
	AssigneeResourceType param.Opt[string] `json:"assignee_resource_type,omitzero"`
	paramObj
}

func (r AssignConversationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AssignConversationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssignConversationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to advance the caller's read cursor in a conversation.
//
// The property UpToSequence is required.
type MarkConversationReadRequestParam struct {
	// Mark all messages up to and including this sequence as read.
	UpToSequence int64 `json:"up_to_sequence" api:"required"`
	paramObj
}

func (r MarkConversationReadRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MarkConversationReadRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarkConversationReadRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to mute a conversation for the caller.
type MuteConversationRequestParam struct {
	// When the mute expires.
	//
	// Omit to mute indefinitely.
	MutedUntil param.Opt[time.Time] `json:"muted_until,omitzero" format:"date-time"`
	paramObj
}

func (r MuteConversationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MuteConversationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MuteConversationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to report a conversation (optionally a specific message) for abuse.
//
// The property Reason is required.
type ReportConversationRequestParam struct {
	// The reason the conversation/message is being reported.
	Reason string `json:"reason" api:"required"`
	// The specific message being reported.
	//
	// Omit to report the conversation as a whole.
	MessageID param.Opt[string] `json:"message_id,omitzero"`
	paramObj
}

func (r ReportConversationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ReportConversationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReportConversationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to place a conversation under legal hold or release it.
//
// While held, the conversation is exempt from automatic retention purging and from
// GDPR redaction.
//
// The property LegalHold is required.
type SetLegalHoldRequestParam struct {
	// The legal-hold status to set.
	//
	// Any of "released", "held".
	LegalHold SetLegalHoldRequestLegalHold `json:"legal_hold,omitzero" api:"required"`
	paramObj
}

func (r SetLegalHoldRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SetLegalHoldRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetLegalHoldRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The legal-hold status to set.
type SetLegalHoldRequestLegalHold string

const (
	SetLegalHoldRequestLegalHoldReleased SetLegalHoldRequestLegalHold = "released"
	SetLegalHoldRequestLegalHoldHeld     SetLegalHoldRequestLegalHold = "held"
)

// Request to set the triage lane of a customer-service case.
//
// The property WorkflowStatus is required.
type SetWorkflowStatusRequestParam struct {
	// The triage lane to move the case to.
	//
	// Any of "new", "open", "waiting_internal", "waiting_external", "needs_approval",
	// "resolved".
	WorkflowStatus SetWorkflowStatusRequestWorkflowStatus `json:"workflow_status,omitzero" api:"required"`
	paramObj
}

func (r SetWorkflowStatusRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SetWorkflowStatusRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetWorkflowStatusRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The triage lane to move the case to.
type SetWorkflowStatusRequestWorkflowStatus string

const (
	SetWorkflowStatusRequestWorkflowStatusNew             SetWorkflowStatusRequestWorkflowStatus = "new"
	SetWorkflowStatusRequestWorkflowStatusOpen            SetWorkflowStatusRequestWorkflowStatus = "open"
	SetWorkflowStatusRequestWorkflowStatusWaitingInternal SetWorkflowStatusRequestWorkflowStatus = "waiting_internal"
	SetWorkflowStatusRequestWorkflowStatusWaitingExternal SetWorkflowStatusRequestWorkflowStatus = "waiting_external"
	SetWorkflowStatusRequestWorkflowStatusNeedsApproval   SetWorkflowStatusRequestWorkflowStatus = "needs_approval"
	SetWorkflowStatusRequestWorkflowStatusResolved        SetWorkflowStatusRequestWorkflowStatus = "resolved"
)

type MessagingConversationActionArchiveParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationActionArchiveParams]'s query
// parameters as `url.Values`.
func (r MessagingConversationActionArchiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionAssignParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to assign a customer-service case to a single owner — a user or a team.
	//
	// The owner is a polymorphic (`assignee_resource_type`, `assignee_resource_id`)
	// reference; omit both fields to clear the assignment.
	AssignConversationRequest AssignConversationRequestParam
	paramObj
}

func (r MessagingConversationActionAssignParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AssignConversationRequest)
}
func (r *MessagingConversationActionAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationActionAssignParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionAssignParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionHideParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationActionHideParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionHideParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionLeaveParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationActionLeaveParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionLeaveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionMuteParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to mute a conversation for the caller.
	MuteConversationRequest MuteConversationRequestParam
	paramObj
}

func (r MessagingConversationActionMuteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MuteConversationRequest)
}
func (r *MessagingConversationActionMuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationActionMuteParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionMuteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionReadParams struct {
	// Request to advance the caller's read cursor in a conversation.
	MarkConversationReadRequest MarkConversationReadRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationActionReadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarkConversationReadRequest)
}
func (r *MessagingConversationActionReadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationActionReadParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionRedactParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationActionRedactParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionRedactParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionReportParams struct {
	// Request to report a conversation (optionally a specific message) for abuse.
	ReportConversationRequest ReportConversationRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationActionReportParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ReportConversationRequest)
}
func (r *MessagingConversationActionReportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationActionReportParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionReportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionSetLegalHoldParams struct {
	// Request to place a conversation under legal hold or release it.
	//
	// While held, the conversation is exempt from automatic retention purging and from
	// GDPR redaction.
	SetLegalHoldRequest SetLegalHoldRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationActionSetLegalHoldParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetLegalHoldRequest)
}
func (r *MessagingConversationActionSetLegalHoldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationActionSetLegalHoldParams]'s query
// parameters as `url.Values`.
func (r MessagingConversationActionSetLegalHoldParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionSetStatusParams struct {
	// Request to set the triage lane of a customer-service case.
	SetWorkflowStatusRequest SetWorkflowStatusRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationActionSetStatusParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetWorkflowStatusRequest)
}
func (r *MessagingConversationActionSetStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationActionSetStatusParams]'s query
// parameters as `url.Values`.
func (r MessagingConversationActionSetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionUnarchiveParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationActionUnarchiveParams]'s query
// parameters as `url.Values`.
func (r MessagingConversationActionUnarchiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionUnhideParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationActionUnhideParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionUnhideParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationActionUnmuteParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationActionUnmuteParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationActionUnmuteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
