// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
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

// List and manage addresses for accounts.
//
// SaleAddressService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleAddressService] method instead.
type SaleAddressService struct {
	options []option.RequestOption
}

// NewSaleAddressService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleAddressService(opts ...option.RequestOption) (r SaleAddressService) {
	r = SaleAddressService{}
	r.options = opts
	return
}

// Creates an address for the targeted account.
func (r *SaleAddressService) New(ctx context.Context, body SaleAddressNewParams, opts ...option.RequestOption) (res *Address, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an address by ID.
func (r *SaleAddressService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Address, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/addresses/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially updates an address.
func (r *SaleAddressService) Update(ctx context.Context, id string, body SaleAddressUpdateParams, opts ...option.RequestOption) (res *Address, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/addresses/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of addresses.
func (r *SaleAddressService) List(ctx context.Context, query SaleAddressListParams, opts ...option.RequestOption) (res *ListAddress, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an address. Fails if the address is in use as a billing or shipping
// address on a sales order, invoice, or shipment, or as a default account address.
func (r *SaleAddressService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleAddressDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/addresses/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create an address.
//
// The properties Country, Name are required.
type AddressInputParam struct {
	// Two-letter country code.
	Country string `json:"country" api:"required"`
	// Display name of the address.
	Name string `json:"name" api:"required"`
	// Email address associated with the address.
	Email param.Opt[string] `json:"email,omitzero"`
	// City or locality.
	Locality param.Opt[string] `json:"locality,omitzero"`
	// Phone number associated with the address.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// State or administrative area.
	State param.Opt[string] `json:"state,omitzero"`
	// First line of the street address.
	StreetLine1 param.Opt[string] `json:"street_line_1,omitzero"`
	// Second line of the street address.
	StreetLine2 param.Opt[string] `json:"street_line_2,omitzero"`
	// Address type.
	//
	// Any of "standard", "drop_ship".
	Type AddressInputType `json:"type,omitzero"`
	paramObj
}

func (r AddressInputParam) MarshalJSON() (data []byte, err error) {
	type shadow AddressInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Address type.
type AddressInputType string

const (
	AddressInputTypeStandard AddressInputType = "standard"
	AddressInputTypeDropShip AddressInputType = "drop_ship"
)

// List represents a paginated list of resources.
type ListAddress struct {
	// Resources in this page.
	Data []Address `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAddressObject `json:"object" api:"required"`
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
func (r ListAddress) RawJSON() string { return r.JSON.raw }
func (r *ListAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAddressObject string

const (
	ListAddressObjectList ListAddressObject = "list"
)

// Request to partially update an address.
type UpdateAddressRequestParam struct {
	// Email address associated with the address.
	Email param.Opt[string] `json:"email,omitzero"`
	// Phone number associated with the address.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Second line of the street address.
	StreetLine2 param.Opt[string] `json:"street_line_2,omitzero"`
	// Two-letter country code.
	Country param.Opt[string] `json:"country,omitzero"`
	// City or locality.
	Locality param.Opt[string] `json:"locality,omitzero"`
	// Display name of the address.
	Name param.Opt[string] `json:"name,omitzero"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// State or administrative area.
	State param.Opt[string] `json:"state,omitzero"`
	// First line of the street address.
	StreetLine1 param.Opt[string] `json:"street_line_1,omitzero"`
	// Address type.
	//
	// Any of "standard", "drop_ship".
	Type UpdateAddressRequestType `json:"type,omitzero"`
	paramObj
}

func (r UpdateAddressRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAddressRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAddressRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Address type.
type UpdateAddressRequestType string

const (
	UpdateAddressRequestTypeStandard UpdateAddressRequestType = "standard"
	UpdateAddressRequestTypeDropShip UpdateAddressRequestType = "drop_ship"
)

type SaleAddressDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleAddressDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleAddressDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAddressNewParams struct {
	// Request to create an address.
	AddressInput AddressInputParam
	paramObj
}

func (r SaleAddressNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AddressInput)
}
func (r *SaleAddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAddressUpdateParams struct {
	// Request to partially update an address.
	UpdateAddressRequest UpdateAddressRequestParam
	paramObj
}

func (r SaleAddressUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAddressRequest)
}
func (r *SaleAddressUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAddressListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter by address type.
	//
	// Any of "standard", "drop_ship".
	Type SaleAddressListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleAddressListParams]'s query parameters as `url.Values`.
func (r SaleAddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by address type.
type SaleAddressListParamsType string

const (
	SaleAddressListParamsTypeStandard SaleAddressListParamsType = "standard"
	SaleAddressListParamsTypeDropShip SaleAddressListParamsType = "drop_ship"
)
