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

// List and manage third-party account integrations.
//
// SettingIntegrationService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingIntegrationService] method instead.
type SettingIntegrationService struct {
	options []option.RequestOption
}

// NewSettingIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingIntegrationService(opts ...option.RequestOption) (r SettingIntegrationService) {
	r = SettingIntegrationService{}
	r.options = opts
	return
}

// Creates an account integration, or updates the name and credentials of an
// existing one with the same integration code.
//
// Credentials are validated for the provider, encrypted at rest, and never
// returned in API responses. An account can have at most one integration per
// integration code.
func (r *SettingIntegrationService) New(ctx context.Context, body SettingIntegrationNewParams, opts ...option.RequestOption) (res *AccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates an account integration's name and active status.
//
// Omitted fields are left unchanged. Credentials cannot be changed with this
// endpoint; to rotate credentials, call Create Account Integration again with the
// same integration code.
func (r *SettingIntegrationService) Update(ctx context.Context, id string, body SettingIntegrationUpdateParams, opts ...option.RequestOption) (res *AccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/settings/integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of account integrations for the target account.
func (r *SettingIntegrationService) List(ctx context.Context, query SettingIntegrationListParams, opts ...option.RequestOption) (res *ListAccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an account integration and returns the deleted resource.
func (r *SettingIntegrationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/settings/integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Third-party integration connected to an account.
type AccountIntegration struct {
	// Account integration ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the integration.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_integration".
	Object AccountIntegrationObject `json:"object" api:"required"`
	// Integration provider code.
	//
	// - `stripe`: Stripe payment processing.
	// - `shippo`: Shippo shipping and label generation.
	// - `hubspot`: HubSpot CRM.
	//
	// Any of "stripe", "shippo", "hubspot".
	Provider AccountIntegrationProvider `json:"provider" api:"required"`
	// Lifecycle status of the integration.
	//
	// Integrations are created `active`. Setting an integration to `inactive` keeps
	// its stored credentials but stops it from being used (for example, the Stripe
	// publishable key cannot be retrieved while the Stripe integration is inactive).
	//
	// Any of "active", "inactive".
	Status AccountIntegrationStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Provider    respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountIntegration) RawJSON() string { return r.JSON.raw }
func (r *AccountIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountIntegrationObject string

const (
	AccountIntegrationObjectAccountIntegration AccountIntegrationObject = "account_integration"
)

// Integration provider code.
//
// - `stripe`: Stripe payment processing.
// - `shippo`: Shippo shipping and label generation.
// - `hubspot`: HubSpot CRM.
type AccountIntegrationProvider string

const (
	AccountIntegrationProviderStripe  AccountIntegrationProvider = "stripe"
	AccountIntegrationProviderShippo  AccountIntegrationProvider = "shippo"
	AccountIntegrationProviderHubspot AccountIntegrationProvider = "hubspot"
)

// Lifecycle status of the integration.
//
// Integrations are created `active`. Setting an integration to `inactive` keeps
// its stored credentials but stops it from being used (for example, the Stripe
// publishable key cannot be retrieved while the Stripe integration is inactive).
type AccountIntegrationStatus string

const (
	AccountIntegrationStatusActive   AccountIntegrationStatus = "active"
	AccountIntegrationStatusInactive AccountIntegrationStatus = "inactive"
)

// Request to create or upsert an account integration.
//
// The properties Credentials, Name, Provider are required.
type CreateAccountIntegrationRequestParam struct {
	// JSON string containing the provider's credentials.
	//
	// Required keys depend on the provider:
	//
	//   - `stripe`: `private_key` (`sk_...`), `publishable_key` (`pk_...`), and
	//     `webhook_secret` (`whsec_...`).
	//   - `shippo`: `api_key` (`shippo_live_...` or `shippo_test_...`).
	//   - `hubspot`: `access_token` (`pat-...`).
	//
	// Sandbox accounts must use test keys and production accounts must use live keys;
	// credentials that do not match are rejected.
	Credentials string `json:"credentials" api:"required"`
	// Display name of the integration.
	Name string `json:"name" api:"required"`
	// Integration provider code.
	//
	// - `stripe`: Stripe payment processing.
	// - `shippo`: Shippo shipping and label generation.
	// - `hubspot`: HubSpot CRM.
	//
	// Any of "stripe", "shippo", "hubspot".
	Provider CreateAccountIntegrationRequestProvider `json:"provider,omitzero" api:"required"`
	paramObj
}

func (r CreateAccountIntegrationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAccountIntegrationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAccountIntegrationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration provider code.
//
// - `stripe`: Stripe payment processing.
// - `shippo`: Shippo shipping and label generation.
// - `hubspot`: HubSpot CRM.
type CreateAccountIntegrationRequestProvider string

const (
	CreateAccountIntegrationRequestProviderStripe  CreateAccountIntegrationRequestProvider = "stripe"
	CreateAccountIntegrationRequestProviderShippo  CreateAccountIntegrationRequestProvider = "shippo"
	CreateAccountIntegrationRequestProviderHubspot CreateAccountIntegrationRequestProvider = "hubspot"
)

// List represents a paginated list of resources.
type ListAccountIntegration struct {
	// Resources in this page.
	Data []AccountIntegration `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAccountIntegrationObject `json:"object" api:"required"`
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
func (r ListAccountIntegration) RawJSON() string { return r.JSON.raw }
func (r *ListAccountIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAccountIntegrationObject string

const (
	ListAccountIntegrationObjectList ListAccountIntegrationObject = "list"
)

// Request to update an account integration.
type UpdateAccountIntegrationRequestParam struct {
	// Display name of the integration.
	Name param.Opt[string] `json:"name,omitzero"`
	// Lifecycle status of the integration.
	//
	// Set to `inactive` to deactivate the integration without deleting its stored
	// credentials.
	//
	// Any of "active", "inactive".
	Status UpdateAccountIntegrationRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r UpdateAccountIntegrationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAccountIntegrationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAccountIntegrationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of the integration.
//
// Set to `inactive` to deactivate the integration without deleting its stored
// credentials.
type UpdateAccountIntegrationRequestStatus string

const (
	UpdateAccountIntegrationRequestStatusActive   UpdateAccountIntegrationRequestStatus = "active"
	UpdateAccountIntegrationRequestStatusInactive UpdateAccountIntegrationRequestStatus = "inactive"
)

type SettingIntegrationNewParams struct {
	// Request to create or upsert an account integration.
	CreateAccountIntegrationRequest CreateAccountIntegrationRequestParam
	paramObj
}

func (r SettingIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAccountIntegrationRequest)
}
func (r *SettingIntegrationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIntegrationUpdateParams struct {
	// Request to update an account integration.
	UpdateAccountIntegrationRequest UpdateAccountIntegrationRequestParam
	paramObj
}

func (r SettingIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAccountIntegrationRequest)
}
func (r *SettingIntegrationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIntegrationListParams struct {
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
	paramObj
}

// URLQuery serializes [SettingIntegrationListParams]'s query parameters as
// `url.Values`.
func (r SettingIntegrationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
