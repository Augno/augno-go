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

// Create and manage API keys for programmatic access.
//
// AuthAPIKeyService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthAPIKeyService] method instead.
type AuthAPIKeyService struct {
	options []option.RequestOption
	// Create and manage API keys for programmatic access.
	Actions AuthAPIKeyActionService
}

// NewAuthAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthAPIKeyService(opts ...option.RequestOption) (r AuthAPIKeyService) {
	r = AuthAPIKeyService{}
	r.options = opts
	r.Actions = NewAuthAPIKeyActionService(opts...)
	return
}

// Creates an [API key](https://docs.augno.com/api/api-keys) to authenticate API
// requests.
//
// The secret key is returned once and cannot be retrieved later, so you should
// store it securely. We provide some
// [recommendations](https://docs.augno.com/api/managing-api-keys) on how you can
// manage your API keys.
func (r *AuthAPIKeyService) New(ctx context.Context, params AuthAPIKeyNewParams, opts ...option.RequestOption) (res *CreatedAPIKey, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/auth/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns [API key](https://docs.augno.com/api/api-keys) metadata by ID.
func (r *AuthAPIKeyService) Get(ctx context.Context, id string, query AuthAPIKeyGetParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/auth/api-keys/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of [API keys](https://docs.augno.com/api/api-keys).
func (r *AuthAPIKeyService) List(ctx context.Context, query AuthAPIKeyListParams, opts ...option.RequestOption) (res *ListAPIKey, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/auth/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Revokes an [API key](https://docs.augno.com/api/api-keys).
//
// Revoked API keys will be unable to be used to authenticate requests.
func (r *AuthAPIKeyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AuthAPIKeyDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/auth/api-keys/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Account with optional branding and portal sub-resources.
type Account struct {
	// Account ID.
	ID string `json:"id" api:"required"`
	// Branding metadata for an account.
	Branding AccountBranding `json:"branding" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Address with associated geolocation.
	DefaultBillingAddress Address `json:"default_billing_address" api:"required"`
	// Address with associated geolocation.
	DefaultShippingAddress Address `json:"default_shipping_address" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "account".
	Object AccountObject `json:"object" api:"required"`
	// Portal metadata for an account.
	Portal AccountPortal `json:"portal" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Branding               respjson.Field
		CreatedAt              respjson.Field
		DefaultBillingAddress  respjson.Field
		DefaultShippingAddress respjson.Field
		Name                   respjson.Field
		Object                 respjson.Field
		Portal                 respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountObject string

const (
	AccountObjectAccount AccountObject = "account"
)

// Branding metadata for an account.
type AccountBranding struct {
	// Branding ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Facebook handle.
	FacebookHandle string `json:"facebook_handle" api:"required"`
	// Instagram handle.
	InstagramHandle string `json:"instagram_handle" api:"required"`
	// LinkedIn handle.
	LinkedinHandle string `json:"linkedin_handle" api:"required"`
	// Logo URL.
	LogoURL string `json:"logo_url" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_branding".
	Object AccountBrandingObject `json:"object" api:"required"`
	// Support phone number.
	PhoneNumber string `json:"phone_number" api:"required"`
	// Support email address.
	SupportEmail string `json:"support_email" api:"required"`
	// Twitter handle.
	TwitterHandle string `json:"twitter_handle" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Website URL.
	WebsiteURL string `json:"website_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		FacebookHandle  respjson.Field
		InstagramHandle respjson.Field
		LinkedinHandle  respjson.Field
		LogoURL         respjson.Field
		Object          respjson.Field
		PhoneNumber     respjson.Field
		SupportEmail    respjson.Field
		TwitterHandle   respjson.Field
		UpdatedAt       respjson.Field
		WebsiteURL      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBranding) RawJSON() string { return r.JSON.raw }
func (r *AccountBranding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountBrandingObject string

const (
	AccountBrandingObjectAccountBranding AccountBrandingObject = "account_branding"
)

// Portal metadata for an account.
type AccountPortal struct {
	// Portal ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "account_portal".
	Object AccountPortalObject `json:"object" api:"required"`
	// Portal slug.
	Slug string `json:"slug" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Object      respjson.Field
		Slug        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPortal) RawJSON() string { return r.JSON.raw }
func (r *AccountPortal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountPortalObject string

const (
	AccountPortalObjectAccountPortal AccountPortalObject = "account_portal"
)

// Address with associated geolocation.
type Address struct {
	// Address ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address associated with the address.
	Email string `json:"email" api:"required"`
	// Geolocation sub-resource.
	Geolocation Geolocation `json:"geolocation" api:"required"`
	// Display name of the address.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "address".
	Object AddressObject `json:"object" api:"required"`
	// Phone number associated with the address.
	Phone string `json:"phone" api:"required"`
	// Address type.
	//
	//   - `standard`: a normal shipping or billing address.
	//   - `drop_ship`: an address an order is shipped to directly, typically a third
	//     party or end customer rather than the account itself.
	//
	// Any of "standard", "drop_ship".
	Type AddressType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		Geolocation respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Phone       respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AddressObject string

const (
	AddressObjectAddress AddressObject = "address"
)

// Address type.
//
//   - `standard`: a normal shipping or billing address.
//   - `drop_ship`: an address an order is shipped to directly, typically a third
//     party or end customer rather than the account itself.
type AddressType string

const (
	AddressTypeStandard AddressType = "standard"
	AddressTypeDropShip AddressType = "drop_ship"
)

// API key resource.
type APIKey struct {
	// API key ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the key expires and stops authenticating.
	//
	// `null` if the key never expires.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// When the key was last used to authenticate a request.
	//
	// `null` if it has never been used.
	LastUsedAt time.Time `json:"last_used_at" api:"required" format:"date-time"`
	// Human-readable name for the API key.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "api_key".
	Object APIKeyObject `json:"object" api:"required"`
	// Redacted key value safe for display.
	RedactedValue string `json:"redacted_value" api:"required"`
	// When the key was revoked.
	//
	// `null` if the key has not been revoked.
	RevokedAt time.Time `json:"revoked_at" api:"required" format:"date-time"`
	// Role resource.
	Role Role `json:"role" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		ExpiresAt     respjson.Field
		LastUsedAt    respjson.Field
		Name          respjson.Field
		Object        respjson.Field
		RedactedValue respjson.Field
		RevokedAt     respjson.Field
		Role          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type APIKeyObject string

const (
	APIKeyObjectAPIKey APIKeyObject = "api_key"
)

// Request to create an API key.
//
// The properties Name, RoleID are required.
type CreateAPIKeyRequestParam struct {
	// Human-readable name for the API key.
	Name string `json:"name" api:"required"`
	// Role ID assigned to the API key.
	RoleID string `json:"role_id" api:"required"`
	// Expiration timestamp. If not set, the key does not expire.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	paramObj
}

func (r CreateAPIKeyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAPIKeyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAPIKeyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of creating an API key, with the full secret value.
type CreatedAPIKey struct {
	// API key resource.
	APIKeyInfo APIKey `json:"api_key_info" api:"required"`
	// Full secret value.
	//
	// Returned once and cannot be retrieved later. Learn more about
	// [managing your API keys](https://docs.augno.com/api/managing-api-keys).
	APIKeySecret string `json:"api_key_secret" api:"required"`
	// Resource type identifier.
	//
	// Any of "created_api_key".
	Object CreatedAPIKeyObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyInfo   respjson.Field
		APIKeySecret respjson.Field
		Object       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedAPIKey) RawJSON() string { return r.JSON.raw }
func (r *CreatedAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CreatedAPIKeyObject string

const (
	CreatedAPIKeyObjectCreatedAPIKey CreatedAPIKeyObject = "created_api_key"
)

// Geolocation sub-resource.
type Geolocation struct {
	// Geolocation ID.
	ID string `json:"id" api:"required"`
	// Two-letter country code.
	Country string `json:"country" api:"required"`
	// City or locality.
	Locality string `json:"locality" api:"required"`
	// Resource type identifier.
	//
	// Any of "geolocation".
	Object GeolocationObject `json:"object" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// State or administrative area.
	State string `json:"state" api:"required"`
	// First line of the street address.
	StreetLine1 string `json:"street_line_1" api:"required"`
	// Second line of the street address.
	StreetLine2 string `json:"street_line_2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Country     respjson.Field
		Locality    respjson.Field
		Object      respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		StreetLine1 respjson.Field
		StreetLine2 respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Geolocation) RawJSON() string { return r.JSON.raw }
func (r *Geolocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type GeolocationObject string

const (
	GeolocationObjectGeolocation GeolocationObject = "geolocation"
)

// List represents a paginated list of resources.
type ListAPIKey struct {
	// Resources in this page.
	Data []APIKey `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAPIKeyObject `json:"object" api:"required"`
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
func (r ListAPIKey) RawJSON() string { return r.JSON.raw }
func (r *ListAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAPIKeyObject string

const (
	ListAPIKeyObjectList ListAPIKeyObject = "list"
)

// Owner describes the provenance of a resource.
type Owner struct {
	// Account with optional branding and portal sub-resources.
	Account Account `json:"account" api:"required"`
	// Resource type identifier.
	//
	// Any of "owner".
	Object OwnerObject `json:"object" api:"required"`
	// Owner type, identifying where the resource came from.
	//
	//   - `system`: a platform-provided default shared across all accounts; not
	//     editable.
	//   - `account`: created and owned by a specific account; the `account` field
	//     identifies which.
	//
	// Any of "system", "account".
	Type OwnerType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Object      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Owner) RawJSON() string { return r.JSON.raw }
func (r *Owner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type OwnerObject string

const (
	OwnerObjectOwner OwnerObject = "owner"
)

// Owner type, identifying where the resource came from.
//
//   - `system`: a platform-provided default shared across all accounts; not
//     editable.
//   - `account`: created and owned by a specific account; the `account` field
//     identifies which.
type OwnerType string

const (
	OwnerTypeSystem  OwnerType = "system"
	OwnerTypeAccount OwnerType = "account"
)

// PageInfo contains URL-based pagination metadata.
type PageInfo struct {
	// Whether more results exist after this page.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Whether results exist before this page.
	HasPrevPage bool `json:"has_prev_page" api:"required"`
	// URL to fetch the next page, `null` if no more pages.
	NextPageURL string `json:"next_page_url" api:"required"`
	// URL to fetch the previous page, `null` if on the first page.
	PreviousPageURL string `json:"previous_page_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage     respjson.Field
		HasPrevPage     respjson.Field
		NextPageURL     respjson.Field
		PreviousPageURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageInfo) RawJSON() string { return r.JSON.raw }
func (r *PageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role resource.
type Role struct {
	// Role ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "role".
	Object RoleObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Permissions in `{domain}:{action}` format.
	Permissions []string `json:"permissions" api:"required"`
	// Role type code.
	//
	// The role's type is sometimes used to gate special behaviors in the frontend and
	// to restrict some actions to only certain types of roles. For example, only roles
	// with the type `admin` can create and manage API keys.
	//
	//   - `admin`: full administrative access, including managing API keys.
	//   - `user`: a custom role tailored to a specific need (its permissions are defined
	//     explicitly).
	//   - `scanner`: a role for scanning-station operators.
	//   - `sales_rep`: a role for sales representatives.
	//   - `agent`: a role assigned to an automated agent rather than a person.
	//
	// Any of "admin", "user", "scanner", "sales_rep", "agent".
	Type RoleType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		Permissions respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Role) RawJSON() string { return r.JSON.raw }
func (r *Role) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type RoleObject string

const (
	RoleObjectRole RoleObject = "role"
)

// Role type code.
//
// The role's type is sometimes used to gate special behaviors in the frontend and
// to restrict some actions to only certain types of roles. For example, only roles
// with the type `admin` can create and manage API keys.
//
//   - `admin`: full administrative access, including managing API keys.
//   - `user`: a custom role tailored to a specific need (its permissions are defined
//     explicitly).
//   - `scanner`: a role for scanning-station operators.
//   - `sales_rep`: a role for sales representatives.
//   - `agent`: a role assigned to an automated agent rather than a person.
type RoleType string

const (
	RoleTypeAdmin    RoleType = "admin"
	RoleTypeUser     RoleType = "user"
	RoleTypeScanner  RoleType = "scanner"
	RoleTypeSalesRep RoleType = "sales_rep"
	RoleTypeAgent    RoleType = "agent"
)

type AuthAPIKeyDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthAPIKeyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthAPIKeyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthAPIKeyNewParams struct {
	// Request to create an API key.
	CreateAPIKeyRequest CreateAPIKeyRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r AuthAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAPIKeyRequest)
}
func (r *AuthAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AuthAPIKeyNewParams]'s query parameters as `url.Values`.
func (r AuthAPIKeyNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthAPIKeyGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuthAPIKeyGetParams]'s query parameters as `url.Values`.
func (r AuthAPIKeyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthAPIKeyListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	// API key statuses to filter by.
	//
	// Any of "active", "expired", "revoked".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuthAPIKeyListParams]'s query parameters as `url.Values`.
func (r AuthAPIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
