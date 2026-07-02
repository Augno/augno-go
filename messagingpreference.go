// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/augno/augno-go/internal/apijson"
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// Manage per-category notification channel preferences (in-app, email, push).
//
// MessagingPreferenceService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingPreferenceService] method instead.
type MessagingPreferenceService struct {
	options []option.RequestOption
}

// NewMessagingPreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingPreferenceService(opts ...option.RequestOption) (r MessagingPreferenceService) {
	r = MessagingPreferenceService{}
	r.options = opts
	return
}

// Creates or replaces a notification channel preference for the caller.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingPreferenceService) Update(ctx context.Context, body MessagingPreferenceUpdateParams, opts ...option.RequestOption) (res *NotificationPreference, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Lists the caller's notification channel preferences (global default +
// per-category overrides).
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingPreferenceService) List(ctx context.Context, opts ...option.RequestOption) (res *ListNotificationPreference, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List represents a paginated list of resources.
type ListNotificationPreference struct {
	// Resources in this page.
	Data []NotificationPreference `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListNotificationPreferenceObject `json:"object" api:"required"`
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
func (r ListNotificationPreference) RawJSON() string { return r.JSON.raw }
func (r *ListNotificationPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListNotificationPreferenceObject string

const (
	ListNotificationPreferenceObjectList ListNotificationPreferenceObject = "list"
)

// A per-(user, category) notification channel preference.
//
// A preference with a `null` category is the user's global default; a
// category-specific preference overrides it.
type NotificationPreference struct {
	// Preference ID.
	ID string `json:"id" api:"required"`
	// The notification category this preference applies to.
	//
	// `null` for the global default that applies to all categories without a specific
	// preference.
	Category string `json:"category" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// How email delivery for this category is batched.
	//
	//   - `instant`: send an email as soon as an eligible notification occurs.
	//   - `hourly`: batch eligible notifications into a single hourly email.
	//   - `daily`: batch eligible notifications into a single daily email.
	//   - `off`: never send email for this category, even when email delivery is
	//     otherwise enabled.
	//
	// Any of "instant", "hourly", "daily", "off".
	Digest NotificationPreferenceDigest `json:"digest" api:"required"`
	// Whether email notifications are delivered for this category.
	EmailEnabled bool `json:"email_enabled" api:"required"`
	// Whether in-app (bell) notifications are delivered for this category.
	InAppEnabled bool `json:"in_app_enabled" api:"required"`
	// Resource type identifier.
	//
	// Any of "notification_preference".
	Object NotificationPreferenceObject `json:"object" api:"required"`
	// Whether push notifications are delivered for this category.
	PushEnabled bool `json:"push_enabled" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Category     respjson.Field
		CreatedAt    respjson.Field
		Digest       respjson.Field
		EmailEnabled respjson.Field
		InAppEnabled respjson.Field
		Object       respjson.Field
		PushEnabled  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPreference) RawJSON() string { return r.JSON.raw }
func (r *NotificationPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How email delivery for this category is batched.
//
//   - `instant`: send an email as soon as an eligible notification occurs.
//   - `hourly`: batch eligible notifications into a single hourly email.
//   - `daily`: batch eligible notifications into a single daily email.
//   - `off`: never send email for this category, even when email delivery is
//     otherwise enabled.
type NotificationPreferenceDigest string

const (
	NotificationPreferenceDigestInstant NotificationPreferenceDigest = "instant"
	NotificationPreferenceDigestHourly  NotificationPreferenceDigest = "hourly"
	NotificationPreferenceDigestDaily   NotificationPreferenceDigest = "daily"
	NotificationPreferenceDigestOff     NotificationPreferenceDigest = "off"
)

// Resource type identifier.
type NotificationPreferenceObject string

const (
	NotificationPreferenceObjectNotificationPreference NotificationPreferenceObject = "notification_preference"
)

// Request to create or replace a notification preference for the caller.
//
// The preference is keyed by (caller, category), so repeating the request with the
// same category replaces the existing preference.
//
// The properties EmailEnabled, InAppEnabled, PushEnabled are required.
type UpsertNotificationPreferenceRequestParam struct {
	// Whether email notifications are delivered for this category.
	EmailEnabled bool `json:"email_enabled" api:"required"`
	// Whether in-app (bell) notifications are delivered for this category.
	InAppEnabled bool `json:"in_app_enabled" api:"required"`
	// Whether push notifications are delivered for this category.
	PushEnabled bool `json:"push_enabled" api:"required"`
	// The notification category this preference applies to.
	//
	// Omit (or `null`) to set the caller's global default.
	Category param.Opt[string] `json:"category,omitzero"`
	// How email delivery for this category is batched.
	//
	//   - `instant`: send an email as soon as an eligible notification occurs.
	//   - `hourly`: batch eligible notifications into a single hourly email.
	//   - `daily`: batch eligible notifications into a single daily email.
	//   - `off`: never send email for this category, even when email delivery is
	//     otherwise enabled.
	//
	// Any of "instant", "hourly", "daily", "off".
	Digest UpsertNotificationPreferenceRequestDigest `json:"digest,omitzero"`
	paramObj
}

func (r UpsertNotificationPreferenceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertNotificationPreferenceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertNotificationPreferenceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How email delivery for this category is batched.
//
//   - `instant`: send an email as soon as an eligible notification occurs.
//   - `hourly`: batch eligible notifications into a single hourly email.
//   - `daily`: batch eligible notifications into a single daily email.
//   - `off`: never send email for this category, even when email delivery is
//     otherwise enabled.
type UpsertNotificationPreferenceRequestDigest string

const (
	UpsertNotificationPreferenceRequestDigestInstant UpsertNotificationPreferenceRequestDigest = "instant"
	UpsertNotificationPreferenceRequestDigestHourly  UpsertNotificationPreferenceRequestDigest = "hourly"
	UpsertNotificationPreferenceRequestDigestDaily   UpsertNotificationPreferenceRequestDigest = "daily"
	UpsertNotificationPreferenceRequestDigestOff     UpsertNotificationPreferenceRequestDigest = "off"
)

type MessagingPreferenceUpdateParams struct {
	// Request to create or replace a notification preference for the caller.
	//
	// The preference is keyed by (caller, category), so repeating the request with the
	// same category replaces the existing preference.
	UpsertNotificationPreferenceRequest UpsertNotificationPreferenceRequestParam
	paramObj
}

func (r MessagingPreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpsertNotificationPreferenceRequest)
}
func (r *MessagingPreferenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
