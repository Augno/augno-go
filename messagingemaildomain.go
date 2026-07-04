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
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// Register customer-owned domains with the email bridge and verify them for
// sending and receiving mail.
//
// MessagingEmailDomainService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingEmailDomainService] method instead.
type MessagingEmailDomainService struct {
	options []option.RequestOption
	// Register customer-owned domains with the email bridge and verify them for
	// sending and receiving mail.
	Actions MessagingEmailDomainActionService
}

// NewMessagingEmailDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingEmailDomainService(opts ...option.RequestOption) (r MessagingEmailDomainService) {
	r = MessagingEmailDomainService{}
	r.options = opts
	r.Actions = NewMessagingEmailDomainActionService(opts...)
	return
}

// Registers a customer-owned domain with the email bridge and returns the DKIM
// records to publish.
//
// The domain starts in `pending` until verified.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingEmailDomainService) New(ctx context.Context, body MessagingEmailDomainNewParams, opts ...option.RequestOption) (res *EmailDomain, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single email domain owned by the account.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailDomainService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailDomain, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-domains/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the account's registered email domains.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailDomainService) List(ctx context.Context, opts ...option.RequestOption) (res *ListEmailDomain, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deregisters a customer-owned domain from the email bridge.
//
// The domain's SES identity is removed. The domain must have no inboxes bound to
// it.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingEmailDomainService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingEmailDomainDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-domains/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to register a sending/receiving domain with the email bridge.
//
// The property Domain is required.
type CreateEmailDomainRequestParam struct {
	// The fully-qualified domain name to register (e.g. `support.acme.com`).
	Domain string `json:"domain" api:"required"`
	paramObj
}

func (r CreateEmailDomainRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateEmailDomainRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateEmailDomainRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A domain registered with the email bridge for sending and receiving mail.
//
// After registration the domain starts in `pending`; publish the returned DKIM
// records, then poll the verify action until it flips to `verified`.
type EmailDomain struct {
	// Email domain ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The DKIM CNAME tokens the customer must publish in DNS to verify the domain.
	DkimTokens []string `json:"dkim_tokens" api:"required"`
	// The fully-qualified domain name (e.g. `support.acme.com`).
	Domain string `json:"domain" api:"required"`
	// Resource type identifier.
	//
	// Any of "email_domain".
	Object EmailDomainObject `json:"object" api:"required"`
	// Verification status.
	//
	// - `pending`: registered and awaiting DKIM confirmation.
	// - `verified`: DKIM confirmed; the domain can send and receive mail.
	// - `failed`: verification could not be completed.
	Status string `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// When the domain's DKIM verification was confirmed.
	VerifiedAt time.Time `json:"verified_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DkimTokens  respjson.Field
		Domain      respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomain) RawJSON() string { return r.JSON.raw }
func (r *EmailDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type EmailDomainObject string

const (
	EmailDomainObjectEmailDomain EmailDomainObject = "email_domain"
)

// List represents a paginated list of resources.
type ListEmailDomain struct {
	// Resources in this page.
	Data []EmailDomain `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListEmailDomainObject `json:"object" api:"required"`
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
func (r ListEmailDomain) RawJSON() string { return r.JSON.raw }
func (r *ListEmailDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListEmailDomainObject string

const (
	ListEmailDomainObjectList ListEmailDomainObject = "list"
)

type MessagingEmailDomainDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingEmailDomainDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingEmailDomainDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingEmailDomainNewParams struct {
	// Request to register a sending/receiving domain with the email bridge.
	CreateEmailDomainRequest CreateEmailDomainRequestParam
	paramObj
}

func (r MessagingEmailDomainNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateEmailDomainRequest)
}
func (r *MessagingEmailDomainNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
