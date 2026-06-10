// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"

	"github.com/augno/augno-go/internal/apijson"
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// Autocomplete, look up details, and validate addresses.
//
// CoreAddressActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreAddressActionService] method instead.
type CoreAddressActionService struct {
	options []option.RequestOption
}

// NewCoreAddressActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCoreAddressActionService(opts ...option.RequestOption) (r CoreAddressActionService) {
	r = CoreAddressActionService{}
	r.options = opts
	return
}

// Validates an address and returns whether it is valid, a formatted version, and
// any validation messages.
func (r *CoreAddressActionService) Validate(ctx context.Context, body CoreAddressActionValidateParams, opts ...option.RequestOption) (res *ValidatedAddress, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/addresses/actions/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Parsed address components.
type AddressComponents struct {
	// First line of the street address.
	AddressLine1 string `json:"address_line_1" api:"required"`
	// Second line of the street address.
	AddressLine2 string `json:"address_line_2" api:"required"`
	// City or locality.
	City string `json:"city" api:"required"`
	// Country name or code.
	Country string `json:"country" api:"required"`
	// Two-letter country code.
	CountryCode string `json:"country_code" api:"required"`
	// Resource type identifier.
	//
	// Any of "address_components".
	Object AddressComponentsObject `json:"object" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// State or administrative area.
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressLine1 respjson.Field
		AddressLine2 respjson.Field
		City         respjson.Field
		Country      respjson.Field
		CountryCode  respjson.Field
		Object       respjson.Field
		PostalCode   respjson.Field
		State        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressComponents) RawJSON() string { return r.JSON.raw }
func (r *AddressComponents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AddressComponentsObject string

const (
	AddressComponentsObjectAddressComponents AddressComponentsObject = "address_components"
)

// Request to validate an address.
//
// The properties AddressLine1, City, Country, PostalCode, State are required.
type ValidateAddressRequestParam struct {
	// First line of the street address.
	AddressLine1 string `json:"address_line_1" api:"required"`
	// City or locality.
	City string `json:"city" api:"required"`
	// Country name or code.
	Country string `json:"country" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// State or administrative area.
	State string `json:"state" api:"required"`
	// Second line of the street address.
	AddressLine2 param.Opt[string] `json:"address_line_2,omitzero"`
	paramObj
}

func (r ValidateAddressRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ValidateAddressRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ValidateAddressRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of address validation.
type ValidatedAddress struct {
	// Parsed address components.
	Components AddressComponents `json:"components" api:"required"`
	// Formatted address from the validation service.
	FormattedAddress string `json:"formatted_address" api:"required"`
	// Resource type identifier.
	//
	// Any of "validated_address".
	Object ValidatedAddressObject `json:"object" api:"required"`
	// Address validation status.
	//
	// Any of "valid", "invalid".
	Status ValidatedAddressStatus `json:"status" api:"required"`
	// Validation messages for issues found.
	ValidationMessages []string `json:"validation_messages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Components         respjson.Field
		FormattedAddress   respjson.Field
		Object             respjson.Field
		Status             respjson.Field
		ValidationMessages respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValidatedAddress) RawJSON() string { return r.JSON.raw }
func (r *ValidatedAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ValidatedAddressObject string

const (
	ValidatedAddressObjectValidatedAddress ValidatedAddressObject = "validated_address"
)

// Address validation status.
type ValidatedAddressStatus string

const (
	ValidatedAddressStatusValid   ValidatedAddressStatus = "valid"
	ValidatedAddressStatusInvalid ValidatedAddressStatus = "invalid"
)

type CoreAddressActionValidateParams struct {
	// Request to validate an address.
	ValidateAddressRequest ValidateAddressRequestParam
	paramObj
}

func (r CoreAddressActionValidateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ValidateAddressRequest)
}
func (r *CoreAddressActionValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
