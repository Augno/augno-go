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
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/respjson"
)

// List, read, and manage in-app notifications.
//
// MessagingNotificationActionService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingNotificationActionService] method instead.
type MessagingNotificationActionService struct {
	options []option.RequestOption
}

// NewMessagingNotificationActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingNotificationActionService(opts ...option.RequestOption) (r MessagingNotificationActionService) {
	r = MessagingNotificationActionService{}
	r.options = opts
	return
}

// Dismisses a notification, removing it from the active feed.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingNotificationActionService) Dismiss(ctx context.Context, id string, body MessagingNotificationActionDismissParams, opts ...option.RequestOption) (res *Notification, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/notifications/%s/actions/dismiss", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks every one of the caller's unseen notifications as seen.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingNotificationActionService) MarkAllSeen(ctx context.Context, opts ...option.RequestOption) (res *MessagingNotificationActionMarkAllSeenResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/notifications/actions/mark-all-seen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Marks a notification as read.
//
// Reading also marks the notification seen if it was not already.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingNotificationActionService) Read(ctx context.Context, id string, body MessagingNotificationActionReadParams, opts ...option.RequestOption) (res *Notification, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/notifications/%s/actions/read", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks a notification as seen.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingNotificationActionService) Seen(ctx context.Context, id string, body MessagingNotificationActionSeenParams, opts ...option.RequestOption) (res *Notification, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/notifications/%s/actions/seen", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MessagingNotificationActionMarkAllSeenResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingNotificationActionMarkAllSeenResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingNotificationActionMarkAllSeenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingNotificationActionDismissParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingNotificationActionDismissParams]'s query
// parameters as `url.Values`.
func (r MessagingNotificationActionDismissParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingNotificationActionReadParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingNotificationActionReadParams]'s query parameters
// as `url.Values`.
func (r MessagingNotificationActionReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingNotificationActionSeenParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingNotificationActionSeenParams]'s query parameters
// as `url.Values`.
func (r MessagingNotificationActionSeenParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
