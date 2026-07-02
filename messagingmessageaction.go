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

// Send, list, edit, and delete chat messages.
//
// MessagingMessageActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingMessageActionService] method instead.
type MessagingMessageActionService struct {
	options []option.RequestOption
}

// NewMessagingMessageActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingMessageActionService(opts ...option.RequestOption) (r MessagingMessageActionService) {
	r = MessagingMessageActionService{}
	r.options = opts
	return
}

// Approves a customer-reply draft and sends it to the customer, promoting the
// draft to a sent customer-visible message in place.
//
// Idempotent on `client_message_id`.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageActionService) ApproveSend(ctx context.Context, id string, params MessagingMessageActionApproveSendParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s/actions/approve-send", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Cancels a scheduled message the caller created (status becomes `canceled`).
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageActionService) Cancel(ctx context.Context, id string, body MessagingMessageActionCancelParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s/actions/cancel", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Discards an open customer-reply draft without sending it (status becomes
// `rejected`).
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageActionService) Reject(ctx context.Context, id string, body MessagingMessageActionRejectParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s/actions/reject", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to approve a customer-reply draft and send it to the customer.
//
// The property ClientMessageID is required.
type ApproveSendDraftRequestParam struct {
	// Client-supplied dedupe key for the resulting customer-visible message.
	ClientMessageID string `json:"client_message_id" api:"required"`
	paramObj
}

func (r ApproveSendDraftRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ApproveSendDraftRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApproveSendDraftRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingMessageActionApproveSendParams struct {
	// Request to approve a customer-reply draft and send it to the customer.
	ApproveSendDraftRequest ApproveSendDraftRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingMessageActionApproveSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ApproveSendDraftRequest)
}
func (r *MessagingMessageActionApproveSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingMessageActionApproveSendParams]'s query parameters
// as `url.Values`.
func (r MessagingMessageActionApproveSendParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingMessageActionCancelParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingMessageActionCancelParams]'s query parameters as
// `url.Values`.
func (r MessagingMessageActionCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingMessageActionRejectParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingMessageActionRejectParams]'s query parameters as
// `url.Values`.
func (r MessagingMessageActionRejectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
