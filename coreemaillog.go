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

// View email logs for accounts.
//
// CoreEmailLogService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreEmailLogService] method instead.
type CoreEmailLogService struct {
	options []option.RequestOption
}

// NewCoreEmailLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreEmailLogService(opts ...option.RequestOption) (r CoreEmailLogService) {
	r = CoreEmailLogService{}
	r.options = opts
	return
}

// Returns an email log by ID.
func (r *CoreEmailLogService) Get(ctx context.Context, id string, query CoreEmailLogGetParams, opts ...option.RequestOption) (res *EmailLog, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/email-logs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of email logs for the current account.
func (r *CoreEmailLogService) List(ctx context.Context, query CoreEmailLogListParams, opts ...option.RequestOption) (res *ListEmailLog, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/email-logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A record of an email sent on the account's behalf, such as an invoice or a user
// invitation.
type EmailLog struct {
	// Email log ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Filename of the attached document.
	Filename string `json:"filename" api:"required"`
	// Resource type identifier.
	//
	// Any of "email_log".
	Object EmailLogObject `json:"object" api:"required"`
	// Recipient email addresses.
	Recipients []string `json:"recipients" api:"required"`
	// Email send status.
	//
	// - `pending`: the email is queued and has not been sent yet.
	// - `sent`: the email has been handed off for delivery.
	//
	// Any of "sent", "pending".
	SendStatus EmailLogSendStatus `json:"send_status" api:"required"`
	// Reference to an actor (user, API key, or agent).
	SentBy Actor `json:"sent_by" api:"required"`
	// Email subject line.
	Subject string `json:"subject" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Filename    respjson.Field
		Object      respjson.Field
		Recipients  respjson.Field
		SendStatus  respjson.Field
		SentBy      respjson.Field
		Subject     respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLog) RawJSON() string { return r.JSON.raw }
func (r *EmailLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type EmailLogObject string

const (
	EmailLogObjectEmailLog EmailLogObject = "email_log"
)

// Email send status.
//
// - `pending`: the email is queued and has not been sent yet.
// - `sent`: the email has been handed off for delivery.
type EmailLogSendStatus string

const (
	EmailLogSendStatusSent    EmailLogSendStatus = "sent"
	EmailLogSendStatusPending EmailLogSendStatus = "pending"
)

// List represents a paginated list of resources.
type ListEmailLog struct {
	// Resources in this page.
	Data []EmailLog `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListEmailLogObject `json:"object" api:"required"`
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
func (r ListEmailLog) RawJSON() string { return r.JSON.raw }
func (r *ListEmailLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListEmailLogObject string

const (
	ListEmailLogObjectList ListEmailLogObject = "list"
)

type CoreEmailLogGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sent_by".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreEmailLogGetParams]'s query parameters as `url.Values`.
func (r CoreEmailLogGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CoreEmailLogListParams struct {
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
	// Any of "sent_by".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreEmailLogListParams]'s query parameters as `url.Values`.
func (r CoreEmailLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
