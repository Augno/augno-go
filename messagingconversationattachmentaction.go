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

// Create presigned upload targets for message attachments.
//
// MessagingConversationAttachmentActionService contains methods and other services
// that help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationAttachmentActionService] method instead.
type MessagingConversationAttachmentActionService struct {
	options []option.RequestOption
}

// NewMessagingConversationAttachmentActionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewMessagingConversationAttachmentActionService(opts ...option.RequestOption) (r MessagingConversationAttachmentActionService) {
	r = MessagingConversationAttachmentActionService{}
	r.options = opts
	return
}

// Mints a presigned URL for uploading a chat attachment directly to object
// storage.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingConversationAttachmentActionService) UploadURL(ctx context.Context, id string, params MessagingConversationAttachmentActionUploadURLParams, opts ...option.RequestOption) (res *AttachmentUploadTarget, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/attachments/actions/upload-url", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// A presigned target for uploading a chat attachment directly to object storage.
//
// The client PUTs the file to `upload_url`, then references `s3_key` when sending
// the message. Request `?include=attachment` to expand the pre-allocated
// attachment metadata.
type AttachmentUploadTarget struct {
	// A file, image, link, or resource attached to a message.
	Attachment MessageAttachment `json:"attachment" api:"required"`
	// When the upload URL expires.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "attachment_upload_target".
	Object AttachmentUploadTargetObject `json:"object" api:"required"`
	// The object-storage key to echo back when sending the message.
	S3Key string `json:"s3_key" api:"required"`
	// The presigned URL to PUT the file to.
	UploadURL string `json:"upload_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachment  respjson.Field
		ExpiresAt   respjson.Field
		Object      respjson.Field
		S3Key       respjson.Field
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentUploadTarget) RawJSON() string { return r.JSON.raw }
func (r *AttachmentUploadTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AttachmentUploadTargetObject string

const (
	AttachmentUploadTargetObjectAttachmentUploadTarget AttachmentUploadTargetObject = "attachment_upload_target"
)

// Request to mint a presigned upload target for a chat attachment.
//
// The property Filename is required.
type CreateAttachmentUploadURLRequestParam struct {
	// The original filename of the file to upload.
	Filename string `json:"filename" api:"required"`
	// The MIME content type of the file (sent as the Content-Type on the upload).
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	paramObj
}

func (r CreateAttachmentUploadURLRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAttachmentUploadURLRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAttachmentUploadURLRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingConversationAttachmentActionUploadURLParams struct {
	// Request to mint a presigned upload target for a chat attachment.
	CreateAttachmentUploadURLRequest CreateAttachmentUploadURLRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "attachment", "attachment.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationAttachmentActionUploadURLParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAttachmentUploadURLRequest)
}
func (r *MessagingConversationAttachmentActionUploadURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationAttachmentActionUploadURLParams]'s
// query parameters as `url.Values`.
func (r MessagingConversationAttachmentActionUploadURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
