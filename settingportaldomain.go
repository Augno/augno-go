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
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Connect a custom domain to the account's customer portal, verify its DNS, and
// resolve custom hosts to portal accounts.
//
// SettingPortalDomainService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingPortalDomainService] method instead.
type SettingPortalDomainService struct {
	options []option.RequestOption
	// Connect a custom domain to the account's customer portal, verify its DNS, and
	// resolve custom hosts to portal accounts.
	Actions SettingPortalDomainActionService
}

// NewSettingPortalDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingPortalDomainService(opts ...option.RequestOption) (r SettingPortalDomainService) {
	r = SettingPortalDomainService{}
	r.options = opts
	r.Actions = NewSettingPortalDomainActionService(opts...)
	return
}

// Connects a custom domain to the account's customer portal and returns the DNS
// records to publish.
//
// An account can only have one custom domain at a time: adding a second one — or
// claiming a domain another account already uses — returns a conflict error. The
// new domain starts in `pending`; publish the returned records at your DNS
// provider, then run the verify action to move it towards serving.
//
// This endpoint requires the permission: `self:update`.
func (r *SettingPortalDomainService) New(ctx context.Context, body SettingPortalDomainNewParams, opts ...option.RequestOption) (res *PortalDomain, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/portal-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single portal domain, including its current status and the DNS records
// that must be published for it.
//
// Reading a domain never re-checks it with the serving provider — the status is
// the one recorded when the domain was connected or last verified — so run the
// verify action to move a `pending` or `securing` domain forward.
//
// This endpoint requires the permission: `self:read`.
func (r *SettingPortalDomainService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PortalDomain, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/settings/portal-domains/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the account's portal domains.
//
// An account can only hold one custom portal domain, so this returns either zero
// or one entry. Reading it is the usual way to discover whether a domain is
// connected and what state it is in.
//
// This endpoint requires the permission: `self:read`.
func (r *SettingPortalDomainService) List(ctx context.Context, opts ...option.RequestOption) (res *ListPortalDomain, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/portal-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disconnects the custom domain from the account's customer portal.
//
// The domain is detached from the serving infrastructure and immediately stops
// serving the portal; buyers must go back to the account's default slug-based
// portal address. Because an account may only hold one custom domain, this is how
// you free it up to connect a different one. The DNS records you published can
// then be removed.
//
// This endpoint requires the permission: `self:update`.
func (r *SettingPortalDomainService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SettingPortalDomainDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/settings/portal-domains/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to connect a custom domain to the account's customer portal.
//
// The property Domain is required.
type CreatePortalDomainRequestParam struct {
	// The fully-qualified domain name to connect (e.g. `shop.acme.com`).
	//
	// A subdomain such as `shop.acme.com` is routed with a CNAME record and an apex
	// domain such as `acme.com` with an A record; either way the records to publish
	// come back on the response. The value is lowercased and any trailing dot is
	// stripped before it is stored, and OpenMRP-owned hostnames are rejected.
	Domain string `json:"domain" api:"required"`
	paramObj
}

func (r CreatePortalDomainRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePortalDomainRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePortalDomainRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A DNS record that must be published at your DNS provider before a portal domain
// can be verified and serve traffic.
type DNSRecord struct {
	// Record name (host) to publish.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "dns_record".
	Object DNSRecordObject `json:"object" api:"required"`
	// Why the record must be published.
	//
	//   - `routing`: the record points traffic at the portal's serving infrastructure.
	//   - `ownership`: the record proves control of a domain that is already claimed
	//     elsewhere.
	//
	// Any of "routing", "ownership".
	Reason DNSRecordReason `json:"reason" api:"required"`
	// The kind of DNS record to publish.
	//
	// - `CNAME`: points a subdomain at the portal's serving infrastructure.
	// - `A`: points an apex domain at the portal's serving infrastructure.
	// - `TXT`: carries an ownership-verification challenge.
	//
	// Any of "CNAME", "A", "TXT".
	Type DNSRecordType `json:"type" api:"required"`
	// Record value to publish.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Object      respjson.Field
		Reason      respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSRecord) RawJSON() string { return r.JSON.raw }
func (r *DNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type DNSRecordObject string

const (
	DNSRecordObjectDNSRecord DNSRecordObject = "dns_record"
)

// Why the record must be published.
//
//   - `routing`: the record points traffic at the portal's serving infrastructure.
//   - `ownership`: the record proves control of a domain that is already claimed
//     elsewhere.
type DNSRecordReason string

const (
	DNSRecordReasonRouting   DNSRecordReason = "routing"
	DNSRecordReasonOwnership DNSRecordReason = "ownership"
)

// The kind of DNS record to publish.
//
// - `CNAME`: points a subdomain at the portal's serving infrastructure.
// - `A`: points an apex domain at the portal's serving infrastructure.
// - `TXT`: carries an ownership-verification challenge.
type DNSRecordType string

const (
	DNSRecordTypeCname DNSRecordType = "CNAME"
	DNSRecordTypeA     DNSRecordType = "A"
	DNSRecordTypeTxt   DNSRecordType = "TXT"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDNSRecord struct {
	// Resources in this page.
	Data []DNSRecord `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDNSRecordObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
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
func (r ListDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *ListDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDNSRecordObject string

const (
	ListDNSRecordObjectList ListDNSRecordObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListPortalDomain struct {
	// Resources in this page.
	Data []PortalDomain `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPortalDomainObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
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
func (r ListPortalDomain) RawJSON() string { return r.JSON.raw }
func (r *ListPortalDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPortalDomainObject string

const (
	ListPortalDomainObjectList ListPortalDomainObject = "list"
)

// A custom domain that serves the account's customer portal (e.g.
// `shop.acme.com`).
//
// After creation the domain starts in `pending`; publish the returned DNS records,
// then poll the verify action. Once DNS is correct the domain moves to `securing`
// while its TLS certificate is issued — it is not yet reachable over HTTPS during
// this window — and finally to `verified` once the certificate is live and the
// portal is served on the domain.
type PortalDomain struct {
	// Portal domain ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	DNSRecords ListDNSRecord `json:"dns_records" api:"required"`
	// The fully-qualified domain name (e.g. `shop.acme.com`).
	Domain string `json:"domain" api:"required"`
	// Resource type identifier.
	//
	// Any of "portal_domain".
	Object PortalDomainObject `json:"object" api:"required"`
	// How far the domain has progressed towards serving the portal.
	//
	//   - `pending`: the domain is waiting on DNS. Publish the listed records, then run
	//     the verify action.
	//   - `securing`: DNS is correct and the TLS certificate is being issued. The portal
	//     is not yet reachable over HTTPS.
	//   - `verified`: the certificate is live and the portal is served on the domain.
	//   - `failed`: the domain was rejected and cannot be used.
	//
	// Any of "pending", "securing", "verified", "failed".
	Status PortalDomainStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// When the domain became fully verified — its TLS certificate live and the portal
	// serving on it.
	VerifiedAt time.Time `json:"verified_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DNSRecords  respjson.Field
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
func (r PortalDomain) RawJSON() string { return r.JSON.raw }
func (r *PortalDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PortalDomainObject string

const (
	PortalDomainObjectPortalDomain PortalDomainObject = "portal_domain"
)

// How far the domain has progressed towards serving the portal.
//
//   - `pending`: the domain is waiting on DNS. Publish the listed records, then run
//     the verify action.
//   - `securing`: DNS is correct and the TLS certificate is being issued. The portal
//     is not yet reachable over HTTPS.
//   - `verified`: the certificate is live and the portal is served on the domain.
//   - `failed`: the domain was rejected and cannot be used.
type PortalDomainStatus string

const (
	PortalDomainStatusPending  PortalDomainStatus = "pending"
	PortalDomainStatusSecuring PortalDomainStatus = "securing"
	PortalDomainStatusVerified PortalDomainStatus = "verified"
	PortalDomainStatusFailed   PortalDomainStatus = "failed"
)

type SettingPortalDomainDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SettingPortalDomainDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SettingPortalDomainDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPortalDomainNewParams struct {
	// Request to connect a custom domain to the account's customer portal.
	CreatePortalDomainRequest CreatePortalDomainRequestParam
	paramObj
}

func (r SettingPortalDomainNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePortalDomainRequest)
}
func (r *SettingPortalDomainNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
