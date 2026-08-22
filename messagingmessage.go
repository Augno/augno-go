// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
)

// Send, list, edit, and delete chat messages.
//
// MessagingMessageService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingMessageService] method instead.
type MessagingMessageService struct {
	options []option.RequestOption
	// Send, list, edit, and delete chat messages.
	Actions MessagingMessageActionService
}

// NewMessagingMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingMessageService(opts ...option.RequestOption) (r MessagingMessageService) {
	r = MessagingMessageService{}
	r.options = opts
	r.Actions = NewMessagingMessageActionService(opts...)
	return
}

// Revises a reply draft before it is sent to the customer.
//
// Only a draft that is still awaiting approval can be edited; once it has been
// approved, rejected, or superseded the request fails. Nothing reaches the
// customer until the draft is approved.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageService) Update(ctx context.Context, id string, params MessagingMessageUpdateParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Request to edit a still-open customer-reply draft message.
//
// The property Body is required.
type UpdateDraftRequestParam struct {
	// The revised reply body, replacing what the draft said before.
	Body string `json:"body" api:"required"`
	// The revised subject line for a draft that will be sent by email.
	//
	// Leaving it out keeps the draft's current subject.
	Subject param.Opt[string] `json:"subject,omitzero"`
	paramObj
}

func (r UpdateDraftRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateDraftRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateDraftRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingMessageUpdateParams struct {
	// Request to edit a still-open customer-reply draft message.
	UpdateDraftRequest UpdateDraftRequestParam
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

func (r MessagingMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateDraftRequest)
}
func (r *MessagingMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingMessageUpdateParams]'s query parameters as
// `url.Values`.
func (r MessagingMessageUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
