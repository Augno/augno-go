// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apiquery"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
)

// List, read, and manage broadcast announcements.
//
// MessagingAnnouncementActionService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingAnnouncementActionService] method instead.
type MessagingAnnouncementActionService struct {
	options []option.RequestOption
}

// NewMessagingAnnouncementActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingAnnouncementActionService(opts ...option.RequestOption) (r MessagingAnnouncementActionService) {
	r = MessagingAnnouncementActionService{}
	r.options = opts
	return
}

// Marks an announcement as dismissed for the calling actor.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingAnnouncementActionService) Dismiss(ctx context.Context, id string, body MessagingAnnouncementActionDismissParams, opts ...option.RequestOption) (res *Announcement, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/announcements/%s/actions/dismiss", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks an announcement as read for the calling actor.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingAnnouncementActionService) Read(ctx context.Context, id string, body MessagingAnnouncementActionReadParams, opts ...option.RequestOption) (res *Announcement, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/announcements/%s/actions/read", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks an announcement as seen for the calling actor.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingAnnouncementActionService) Seen(ctx context.Context, id string, body MessagingAnnouncementActionSeenParams, opts ...option.RequestOption) (res *Announcement, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/announcements/%s/actions/seen", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MessagingAnnouncementActionDismissParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingAnnouncementActionDismissParams]'s query
// parameters as `url.Values`.
func (r MessagingAnnouncementActionDismissParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingAnnouncementActionReadParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingAnnouncementActionReadParams]'s query parameters
// as `url.Values`.
func (r MessagingAnnouncementActionReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingAnnouncementActionSeenParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingAnnouncementActionSeenParams]'s query parameters
// as `url.Values`.
func (r MessagingAnnouncementActionSeenParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
