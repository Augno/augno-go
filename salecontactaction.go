// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// Look up the people you do business with.
//
// SaleContactActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleContactActionService] method instead.
type SaleContactActionService struct {
	options []option.RequestOption
}

// NewSaleContactActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleContactActionService(opts ...option.RequestOption) (r SaleContactActionService) {
	r = SaleContactActionService{}
	r.options = opts
	return
}

// Finds the contacts that match an email address.
//
// Only people on accounts you have a relationship with are returned — your
// customers, your suppliers, or your own account. A match's `relationship` says
// how you relate to the account it belongs to. Several accounts can share an
// email, so this can return more than one match.
//
// This endpoint requires the permission: `customers:read`.
func (r *SaleContactActionService) FindByEmail(ctx context.Context, params SaleContactActionFindByEmailParams, opts ...option.RequestOption) (res *ListContactMatch, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/contacts/actions/find-by-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// A contact found by email on an account you have a relationship with — one of
// your customers, your suppliers, or your own account.
//
// The same email can be a contact on many accounts across the platform; only
// accounts you relate to are returned. The matched person is available through
// `account_user` (and the shared profile through `account_user.user`), and the
// account they belong to through `account`.
type ContactMatch struct {
	// Resource ID.
	//
	// This is the matched account user's ID, so the same value also appears as
	// `account_user.id`.
	ID string `json:"id" api:"required"`
	// A customer account, including its branding and customer portal sub-resources.
	Account Account `json:"account" api:"required"`
	// A user's membership in an account, carrying the account-specific status, role,
	// and department.
	//
	// Profile fields (name, email, username, image URL) live on the expandable `user`
	// sub-resource, which is shared across every account the user belongs to.
	AccountUser AccountUser `json:"account_user" api:"required"`
	// The email address that was matched.
	Email string `json:"email" api:"required"`
	// Resource type identifier.
	//
	// Any of "contact_match".
	Object ContactMatchObject `json:"object" api:"required"`
	// How you relate to the account this contact belongs to.
	//
	// - `customer` — the account is one of your customers.
	// - `supplier` — the account is one of your suppliers.
	// - `self` — the account is your own.
	//
	// Any of "customer", "supplier", "self".
	Relationship ContactMatchRelationship `json:"relationship" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Account      respjson.Field
		AccountUser  respjson.Field
		Email        respjson.Field
		Object       respjson.Field
		Relationship respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactMatch) RawJSON() string { return r.JSON.raw }
func (r *ContactMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ContactMatchObject string

const (
	ContactMatchObjectContactMatch ContactMatchObject = "contact_match"
)

// How you relate to the account this contact belongs to.
//
// - `customer` — the account is one of your customers.
// - `supplier` — the account is one of your suppliers.
// - `self` — the account is your own.
type ContactMatchRelationship string

const (
	ContactMatchRelationshipCustomer ContactMatchRelationship = "customer"
	ContactMatchRelationshipSupplier ContactMatchRelationship = "supplier"
	ContactMatchRelationshipSelf     ContactMatchRelationship = "self"
)

// Request to find contacts by email.
//
// The property Email is required.
type FindContactByEmailRequestParam struct {
	// The email address to look up.
	Email string `json:"email" api:"required"`
	paramObj
}

func (r FindContactByEmailRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FindContactByEmailRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FindContactByEmailRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List represents a paginated list of resources.
type ListContactMatch struct {
	// Resources in this page.
	Data []ContactMatch `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListContactMatchObject `json:"object" api:"required"`
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
func (r ListContactMatch) RawJSON() string { return r.JSON.raw }
func (r *ListContactMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListContactMatchObject string

const (
	ListContactMatchObjectList ListContactMatchObject = "list"
)

type SaleContactActionFindByEmailParams struct {
	// Request to find contacts by email.
	FindContactByEmailRequest FindContactByEmailRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account_user", "account_user.user", "account_user.role",
	// "account_user.department", "account".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter to contacts whose relationship to you is one of these.
	//
	// Any of "customer", "supplier", "self".
	Relationships []string `query:"relationships,omitzero" json:"-"`
	paramObj
}

func (r SaleContactActionFindByEmailParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FindContactByEmailRequest)
}
func (r *SaleContactActionFindByEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleContactActionFindByEmailParams]'s query parameters as
// `url.Values`.
func (r SaleContactActionFindByEmailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
