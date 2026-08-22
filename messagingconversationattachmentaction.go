// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Create presigned upload targets for message attachments.
//
// MessagingConversationAttachmentActionService contains methods and other services
// that help with interacting with the openmrp API.
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

// Creates a short-lived URL for uploading a chat attachment straight to object
// storage.
//
// Upload the file to the returned URL, then send a message in the same
// conversation carrying the returned key as an attachment — the file only becomes
// part of the conversation at that point, and an upload that is never sent is
// discarded automatically. You must be an active participant of the conversation
// to stage an upload for it.
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
// PUT the file to `upload_url`, then send a message carrying an attachment whose
// `s3_key` is the key returned here. An upload that is never sent with a message
// is discarded automatically, so abandoning a target costs nothing.
type AttachmentUploadTarget struct {
	// A file, image, link, or resource attached to a message.
	Attachment MessageAttachment `json:"attachment" api:"required"`
	// When the upload URL stops working.
	//
	// Targets are short-lived (about fifteen minutes); request a new one if the upload
	// has not finished by then.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "attachment_upload_target".
	Object AttachmentUploadTargetObject `json:"object" api:"required"`
	// The object-storage key identifying the uploaded file.
	//
	// Pass it back as an attachment's `s3_key` when sending a message. It is bound to
	// the conversation it was minted for and cannot be attached in another one.
	S3Key string `json:"s3_key" api:"required"`
	// The presigned URL to PUT the file to.
	//
	// Send the file with the same content type used to mint the target, or the upload
	// is rejected.
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
	// The MIME content type of the file.
	//
	// The file must then be uploaded with this same content type, or object storage
	// rejects it. It also decides how the attachment preview returned here is
	// classified: `image/…` becomes an inline image, anything else a file.
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
