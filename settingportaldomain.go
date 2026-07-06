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

// Connect a custom domain to the account's customer portal, verify its DNS, and
// resolve custom hosts to portal accounts.
//
// SettingPortalDomainService contains methods and other services that help with
// interacting with the augno API.
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
// Each account can have one custom domain. The domain starts in `pending` until
// its DNS is verified.
//
// This endpoint requires the permission: `self:update`.
func (r *SettingPortalDomainService) New(ctx context.Context, body SettingPortalDomainNewParams, opts ...option.RequestOption) (res *PortalDomain, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/portal-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single portal domain, including the DNS records the customer must
// publish.
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
// Accounts currently have at most one custom portal domain.
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
// serving the portal.
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
	// The fully-qualified domain name to connect (e.g. `shop.acme.com`). Subdomains
	// are recommended; apex domains are supported via an A record.
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

// A DNS record the customer must publish for their portal domain.
type DNSRecord struct {
	// Record name (host) to publish.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "dns_record".
	Object DNSRecordObject `json:"object" api:"required"`
	// Why the record is needed.
	//
	// Routing records point traffic at the portal's serving infrastructure; ownership
	// records prove control of a domain that is claimed elsewhere.
	//
	// Any of "routing", "ownership".
	Reason DNSRecordReason `json:"reason" api:"required"`
	// Record type.
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

// Why the record is needed.
//
// Routing records point traffic at the portal's serving infrastructure; ownership
// records prove control of a domain that is claimed elsewhere.
type DNSRecordReason string

const (
	DNSRecordReasonRouting   DNSRecordReason = "routing"
	DNSRecordReasonOwnership DNSRecordReason = "ownership"
)

// Record type.
type DNSRecordType string

const (
	DNSRecordTypeCname DNSRecordType = "CNAME"
	DNSRecordTypeA     DNSRecordType = "A"
	DNSRecordTypeTxt   DNSRecordType = "TXT"
)

// List represents a paginated list of resources.
type ListDNSRecord struct {
	// Resources in this page.
	Data []DNSRecord `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDNSRecordObject `json:"object" api:"required"`
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
func (r ListDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *ListDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDNSRecordObject string

const (
	ListDNSRecordObjectList ListDNSRecordObject = "list"
)

// List represents a paginated list of resources.
type ListPortalDomain struct {
	// Resources in this page.
	Data []PortalDomain `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPortalDomainObject `json:"object" api:"required"`
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
// then poll the verify action until it flips to `verified`. Once verified, the
// customer portal is served on the domain with TLS provisioned automatically.
type PortalDomain struct {
	// Portal domain ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// List represents a paginated list of resources.
	DNSRecords ListDNSRecord `json:"dns_records" api:"required"`
	// The fully-qualified domain name (e.g. `shop.acme.com`).
	Domain string `json:"domain" api:"required"`
	// Resource type identifier.
	//
	// Any of "portal_domain".
	Object PortalDomainObject `json:"object" api:"required"`
	// Verification status.
	//
	// - pending domains await DNS configuration
	// - verified domains serve the portal
	// - failed domains were rejected and cannot be used
	//
	// Any of "pending", "verified", "failed".
	Status PortalDomainStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// When the domain's DNS configuration was confirmed.
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

// Verification status.
//
// - pending domains await DNS configuration
// - verified domains serve the portal
// - failed domains were rejected and cannot be used
type PortalDomainStatus string

const (
	PortalDomainStatusPending  PortalDomainStatus = "pending"
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
