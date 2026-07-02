// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List messageable contacts (the messaging directory).
//
// MessagingService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingService] method instead.
type MessagingService struct {
	options []option.RequestOption
	// List, read, and manage in-app notifications.
	Notifications MessagingNotificationService
	// List, read, and manage broadcast announcements.
	Announcements MessagingAnnouncementService
	// Create conversations, send and read messages (1:1 direct messages).
	Conversations MessagingConversationService
	// Send, list, edit, and delete chat messages.
	Messages MessagingMessageService
	// Create and manage reusable rosters (named member sets) that seed many
	// conversations.
	Groups MessagingGroupService
	// Block and unblock users from direct messaging.
	Blocks MessagingBlockService
	// Manage per-category notification channel preferences (in-app, email, push).
	Preferences MessagingPreferenceService
	// Register customer-owned domains with the email bridge and verify them for
	// sending and receiving mail.
	EmailDomains MessagingEmailDomainService
	// Provision and manage routable email inboxes that bind inbound mail to chat
	// conversations and send agent replies.
	EmailInboxes MessagingEmailInboxService
}

// NewMessagingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagingService(opts ...option.RequestOption) (r MessagingService) {
	r = MessagingService{}
	r.options = opts
	r.Notifications = NewMessagingNotificationService(opts...)
	r.Announcements = NewMessagingAnnouncementService(opts...)
	r.Conversations = NewMessagingConversationService(opts...)
	r.Messages = NewMessagingMessageService(opts...)
	r.Groups = NewMessagingGroupService(opts...)
	r.Blocks = NewMessagingBlockService(opts...)
	r.Preferences = NewMessagingPreferenceService(opts...)
	r.EmailDomains = NewMessagingEmailDomainService(opts...)
	r.EmailInboxes = NewMessagingEmailInboxService(opts...)
	return
}

// Lists the caller's messageable contacts.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingService) GetContacts(ctx context.Context, query MessagingGetContactsParams, opts ...option.RequestOption) (res *ListActor, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List represents a paginated list of resources.
type ListActor struct {
	// Resources in this page.
	Data []Actor `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListActorObject `json:"object" api:"required"`
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
func (r ListActor) RawJSON() string { return r.JSON.raw }
func (r *ListActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListActorObject string

const (
	ListActorObjectList ListActorObject = "list"
)

type MessagingGetContactsParams struct {
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
	// Any of "role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingGetContactsParams]'s query parameters as
// `url.Values`.
func (r MessagingGetContactsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
