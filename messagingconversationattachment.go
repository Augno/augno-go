// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// MessagingConversationAttachmentService contains methods and other services that
// help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationAttachmentService] method instead.
type MessagingConversationAttachmentService struct {
	options []option.RequestOption
	// Create presigned upload targets for message attachments.
	Actions MessagingConversationAttachmentActionService
}

// NewMessagingConversationAttachmentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingConversationAttachmentService(opts ...option.RequestOption) (r MessagingConversationAttachmentService) {
	r = MessagingConversationAttachmentService{}
	r.options = opts
	r.Actions = NewMessagingConversationAttachmentActionService(opts...)
	return
}
