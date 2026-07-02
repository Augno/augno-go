// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
)

// Register customer-owned domains with the email bridge and verify them for
// sending and receiving mail.
//
// MessagingEmailDomainActionService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingEmailDomainActionService] method instead.
type MessagingEmailDomainActionService struct {
	options []option.RequestOption
}

// NewMessagingEmailDomainActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingEmailDomainActionService(opts ...option.RequestOption) (r MessagingEmailDomainActionService) {
	r = MessagingEmailDomainActionService{}
	r.options = opts
	return
}

// Re-polls the provider and flips the domain to `verified` once its DKIM records
// are confirmed.
//
// Returns the updated domain (still `pending` if not yet confirmed).
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingEmailDomainActionService) Verify(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailDomain, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-domains/%s/actions/verify", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}
