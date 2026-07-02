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
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List, read, and manage broadcast announcements.
//
// MessagingAnnouncementService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingAnnouncementService] method instead.
type MessagingAnnouncementService struct {
	options []option.RequestOption
	// List, read, and manage broadcast announcements.
	Actions MessagingAnnouncementActionService
}

// NewMessagingAnnouncementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingAnnouncementService(opts ...option.RequestOption) (r MessagingAnnouncementService) {
	r = MessagingAnnouncementService{}
	r.options = opts
	r.Actions = NewMessagingAnnouncementActionService(opts...)
	return
}

// Returns one active announcement by ID.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingAnnouncementService) Get(ctx context.Context, id string, query MessagingAnnouncementGetParams, opts ...option.RequestOption) (res *Announcement, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/announcements/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the broadcast announcements currently active for the caller, most recent
// first.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingAnnouncementService) List(ctx context.Context, query MessagingAnnouncementListParams, opts ...option.RequestOption) (res *ListAnnouncement, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/announcements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A broadcast announcement shown in the bell feed, with the caller's per-user read
// state.
type Announcement struct {
	// Announcement ID.
	ID string `json:"id" api:"required"`
	// Preview/body text.
	Body string `json:"body" api:"required"`
	// Category of the announcement.
	//
	// Any of "chat.message", "chat.mention", "chat.added", "order.updated",
	// "agent.run_completed", "agent.alert", "system.broadcast".
	Category AnnouncementCategory `json:"category" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the calling actor dismissed the announcement.
	DismissedAt time.Time `json:"dismissed_at" api:"required" format:"date-time"`
	// When the announcement stops being shown.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "announcement".
	Object AnnouncementObject `json:"object" api:"required"`
	// Delivery priority.
	//
	// Any of "low", "normal", "high", "urgent".
	Priority AnnouncementPriority `json:"priority" api:"required"`
	// When the announcement becomes visible in the feed.
	PublishAt time.Time `json:"publish_at" api:"required" format:"date-time"`
	// When the calling actor opened the announcement.
	ReadAt time.Time `json:"read_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// Reach of the announcement.
	//
	// - `account`: shown only to users within this account.
	// - `platform`: shown to every user across all accounts.
	//
	// Any of "account", "platform".
	Scope AnnouncementScope `json:"scope" api:"required"`
	// When the calling actor first saw the announcement.
	SeenAt time.Time `json:"seen_at" api:"required" format:"date-time"`
	// Lifecycle status of the announcement for the calling actor, derived from their
	// seen/read/dismissed receipt.
	//
	// - `unseen`: not yet surfaced in the caller's feed.
	// - `seen`: surfaced in the feed but not yet opened.
	// - `read`: opened by the caller.
	// - `dismissed`: dismissed by the caller.
	//
	// Any of "unseen", "seen", "read", "dismissed".
	Status AnnouncementStatus `json:"status" api:"required"`
	// Human-readable title.
	Title string `json:"title" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Body        respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		DismissedAt respjson.Field
		ExpiresAt   respjson.Field
		Object      respjson.Field
		Priority    respjson.Field
		PublishAt   respjson.Field
		ReadAt      respjson.Field
		Resource    respjson.Field
		Scope       respjson.Field
		SeenAt      respjson.Field
		Status      respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Announcement) RawJSON() string { return r.JSON.raw }
func (r *Announcement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Category of the announcement.
type AnnouncementCategory string

const (
	AnnouncementCategoryChatMessage       AnnouncementCategory = "chat.message"
	AnnouncementCategoryChatMention       AnnouncementCategory = "chat.mention"
	AnnouncementCategoryChatAdded         AnnouncementCategory = "chat.added"
	AnnouncementCategoryOrderUpdated      AnnouncementCategory = "order.updated"
	AnnouncementCategoryAgentRunCompleted AnnouncementCategory = "agent.run_completed"
	AnnouncementCategoryAgentAlert        AnnouncementCategory = "agent.alert"
	AnnouncementCategorySystemBroadcast   AnnouncementCategory = "system.broadcast"
)

// Resource type identifier.
type AnnouncementObject string

const (
	AnnouncementObjectAnnouncement AnnouncementObject = "announcement"
)

// Delivery priority.
type AnnouncementPriority string

const (
	AnnouncementPriorityLow    AnnouncementPriority = "low"
	AnnouncementPriorityNormal AnnouncementPriority = "normal"
	AnnouncementPriorityHigh   AnnouncementPriority = "high"
	AnnouncementPriorityUrgent AnnouncementPriority = "urgent"
)

// Reach of the announcement.
//
// - `account`: shown only to users within this account.
// - `platform`: shown to every user across all accounts.
type AnnouncementScope string

const (
	AnnouncementScopeAccount  AnnouncementScope = "account"
	AnnouncementScopePlatform AnnouncementScope = "platform"
)

// Lifecycle status of the announcement for the calling actor, derived from their
// seen/read/dismissed receipt.
//
// - `unseen`: not yet surfaced in the caller's feed.
// - `seen`: surfaced in the feed but not yet opened.
// - `read`: opened by the caller.
// - `dismissed`: dismissed by the caller.
type AnnouncementStatus string

const (
	AnnouncementStatusUnseen    AnnouncementStatus = "unseen"
	AnnouncementStatusSeen      AnnouncementStatus = "seen"
	AnnouncementStatusRead      AnnouncementStatus = "read"
	AnnouncementStatusDismissed AnnouncementStatus = "dismissed"
)

// List represents a paginated list of resources.
type ListAnnouncement struct {
	// Resources in this page.
	Data []Announcement `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAnnouncementObject `json:"object" api:"required"`
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
func (r ListAnnouncement) RawJSON() string { return r.JSON.raw }
func (r *ListAnnouncement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAnnouncementObject string

const (
	ListAnnouncementObjectList ListAnnouncementObject = "list"
)

type MessagingAnnouncementGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingAnnouncementGetParams]'s query parameters as
// `url.Values`.
func (r MessagingAnnouncementGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingAnnouncementListParams struct {
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
	// Any of "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingAnnouncementListParams]'s query parameters as
// `url.Values`.
func (r MessagingAnnouncementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
