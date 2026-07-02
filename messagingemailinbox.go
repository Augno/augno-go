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
	"github.com/augno/augno-go/packages/respjson"
)

// Provision and manage routable email inboxes that bind inbound mail to chat
// conversations and send agent replies.
//
// MessagingEmailInboxService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingEmailInboxService] method instead.
type MessagingEmailInboxService struct {
	options []option.RequestOption
}

// NewMessagingEmailInboxService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingEmailInboxService(opts ...option.RequestOption) (r MessagingEmailInboxService) {
	r = MessagingEmailInboxService{}
	r.options = opts
	return
}

// Provisions a routable inbox address on a verified domain.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingEmailInboxService) New(ctx context.Context, params MessagingEmailInboxNewParams, opts ...option.RequestOption) (res *EmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single email inbox owned by the account.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailInboxService) Get(ctx context.Context, id string, query MessagingEmailInboxGetParams, opts ...option.RequestOption) (res *EmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-inboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Edits an email inbox's from-name, status, and default agent trigger config.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingEmailInboxService) Update(ctx context.Context, id string, params MessagingEmailInboxUpdateParams, opts ...option.RequestOption) (res *EmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-inboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns the account's email inboxes.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailInboxService) List(ctx context.Context, query MessagingEmailInboxListParams, opts ...option.RequestOption) (res *ListEmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes an email inbox.
//
// Inbound mail to its address is no longer routed.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingEmailInboxService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingEmailInboxDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-inboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to provision a routable inbox on a verified domain.
//
// The properties Address, EmailDomainID are required.
type CreateEmailInboxRequestParam struct {
	// The full inbox address (e.g. `support@acme.com`).
	//
	// Its domain part must match the selected domain, which must already be verified.
	Address string `json:"address" api:"required"`
	// The verified domain this inbox belongs to.
	EmailDomainID string `json:"email_domain_id" api:"required"`
	// The agent to bind to this inbox to handle incoming mail.
	AgentConfigID param.Opt[string] `json:"agent_config_id,omitzero"`
	// How the bound agent decides whether to run on incoming mail.
	//
	// - `mention`: runs only when the agent is @mentioned in the message.
	// - `keyword`: runs when the message contains any of the trigger keywords.
	// - `always`: runs on every incoming message.
	AgentTriggerPolicy param.Opt[string] `json:"agent_trigger_policy,omitzero"`
	// Display name for the `From` header of outbound mail.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// Keywords that fire the agent when the trigger policy is `keyword`.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords,omitzero"`
	paramObj
}

func (r CreateEmailInboxRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateEmailInboxRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateEmailInboxRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A routable email inbox on a verified domain.
//
// Inbound mail to this address is threaded into a chat conversation, and outbound
// replies may be sent from this identity. The optional agent trigger config
// controls whether the bound agent runs automatically on incoming mail.
type EmailInbox struct {
	// Email inbox ID.
	ID string `json:"id" api:"required"`
	// The full inbox address (e.g. `support@acme.com`).
	Address string `json:"address" api:"required"`
	// An AI agent available to the account.
	//
	// The definition describes what the agent does, how its runs are triggered, the
	// tools it can use, and whether it is currently enabled for the account.
	AgentConfig AgentDefinition `json:"agent_config" api:"required"`
	// Keywords that fire the agent when `agent_trigger_policy` is `keyword`.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords" api:"required"`
	// When the bound agent runs on incoming mail.
	//
	//   - `mention`: only when the agent is @mentioned, matched against its trigger
	//     keywords.
	//   - `keyword`: when the mail contains any of the configured trigger keywords.
	//   - `always`: on every incoming message.
	AgentTriggerPolicy string `json:"agent_trigger_policy" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A domain registered with the email bridge for sending and receiving mail.
	//
	// After registration the domain starts in `pending`; publish the returned DKIM
	// records, then poll the verify action until it flips to `verified`.
	EmailDomain EmailDomain `json:"email_domain" api:"required"`
	// The display name used in the `From` header of outbound mail.
	FromName string `json:"from_name" api:"required"`
	// Resource type identifier.
	//
	// Any of "email_inbox".
	Object EmailInboxObject `json:"object" api:"required"`
	// Whether the inbox is currently routing mail.
	//
	//   - `active`: inbound mail is threaded and outbound replies are allowed.
	//   - `disabled`: the inbox is provisioned but drops inbound mail and does not send
	//     replies.
	Status string `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Address              respjson.Field
		AgentConfig          respjson.Field
		AgentTriggerKeywords respjson.Field
		AgentTriggerPolicy   respjson.Field
		CreatedAt            respjson.Field
		EmailDomain          respjson.Field
		FromName             respjson.Field
		Object               respjson.Field
		Status               respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInbox) RawJSON() string { return r.JSON.raw }
func (r *EmailInbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type EmailInboxObject string

const (
	EmailInboxObjectEmailInbox EmailInboxObject = "email_inbox"
)

// List represents a paginated list of resources.
type ListEmailInbox struct {
	// Resources in this page.
	Data []EmailInbox `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListEmailInboxObject `json:"object" api:"required"`
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
func (r ListEmailInbox) RawJSON() string { return r.JSON.raw }
func (r *ListEmailInbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListEmailInboxObject string

const (
	ListEmailInboxObjectList ListEmailInboxObject = "list"
)

// Request to edit an email inbox's from-name, status, and default agent trigger
// config.
//
// The property Status is required.
type UpdateEmailInboxRequestParam struct {
	// Whether the inbox routes mail.
	//
	// - `active`: inbound mail is threaded and outbound replies are allowed.
	// - `disabled`: the inbox stays provisioned but does not route mail.
	Status string `json:"status" api:"required"`
	// The agent to bind to this inbox to handle incoming mail.
	AgentConfigID param.Opt[string] `json:"agent_config_id,omitzero"`
	// How the bound agent decides whether to run on incoming mail.
	//
	// - `mention`: runs only when the agent is @mentioned in the message.
	// - `keyword`: runs when the message contains any of the trigger keywords.
	// - `always`: runs on every incoming message.
	AgentTriggerPolicy param.Opt[string] `json:"agent_trigger_policy,omitzero"`
	// Display name for the `From` header of outbound mail.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// Keywords that fire the agent when the trigger policy is `keyword`.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords,omitzero"`
	paramObj
}

func (r UpdateEmailInboxRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateEmailInboxRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateEmailInboxRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingEmailInboxDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingEmailInboxDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingEmailInboxDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingEmailInboxNewParams struct {
	// Request to provision a routable inbox on a verified domain.
	CreateEmailInboxRequest CreateEmailInboxRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingEmailInboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateEmailInboxRequest)
}
func (r *MessagingEmailInboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingEmailInboxNewParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingEmailInboxGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingEmailInboxGetParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingEmailInboxUpdateParams struct {
	// Request to edit an email inbox's from-name, status, and default agent trigger
	// config.
	UpdateEmailInboxRequest UpdateEmailInboxRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingEmailInboxUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateEmailInboxRequest)
}
func (r *MessagingEmailInboxUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingEmailInboxUpdateParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingEmailInboxListParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingEmailInboxListParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
