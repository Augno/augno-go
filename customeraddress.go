// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/augno-go/internal/apijson"
	"github.com/stainless-sdks/augno-go/internal/requestconfig"
	"github.com/stainless-sdks/augno-go/option"
	"github.com/stainless-sdks/augno-go/packages/param"
	"github.com/stainless-sdks/augno-go/packages/respjson"
)

// CustomerAddressService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerAddressService] method instead.
type CustomerAddressService struct {
	Options []option.RequestOption
}

// NewCustomerAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerAddressService(opts ...option.RequestOption) (r CustomerAddressService) {
	r = CustomerAddressService{}
	r.Options = opts
	return
}

// Creates and validates a new address for a customer using Google Address
// Validation.
func (r *CustomerAddressService) New(ctx context.Context, customerID string, body CustomerAddressNewParams, opts ...option.RequestOption) (res *Address, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/addresses", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information for a specific customer address by ID.
func (r *CustomerAddressService) Get(ctx context.Context, addressID string, query CustomerAddressGetParams, opts ...option.RequestOption) (res *GetCustomerAddress, err error) {
	opts = append(r.Options[:], opts...)
	if query.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if addressID == "" {
		err = errors.New("missing required address_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/addresses/%s", query.CustomerID, addressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates address information with automatic address validation and geolocation
// updates.
func (r *CustomerAddressService) Update(ctx context.Context, addressID string, params CustomerAddressUpdateParams, opts ...option.RequestOption) (res *Address, err error) {
	opts = append(r.Options[:], opts...)
	if params.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if addressID == "" {
		err = errors.New("missing required address_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/addresses/%s", params.CustomerID, addressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieves all addresses associated with a specific customer including
// geolocation data.
func (r *CustomerAddressService) List(ctx context.Context, customerID string, opts ...option.RequestOption) (res *GetCustomerAddress, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/addresses", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an address for a customer.
func (r *CustomerAddressService) Delete(ctx context.Context, addressID string, body CustomerAddressDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.CustomerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if addressID == "" {
		err = errors.New("missing required address_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/addresses/%s", body.CustomerID, addressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Represents a Address resource
type Address struct {
	// Unique identifier for the address. We use type ids, which are k-sortable and
	// human readable. You can read more about them
	// [here](https://github.com/jetify-com/typeid).
	ID string `json:"id,required"`
	// Timestamp when the address was created
	CreatedAt string `json:"created_at,required"`
	// Contact name associated with this address
	Name string `json:"name,required"`
	// City or locality name
	City string `json:"city"`
	// ISO country code
	Country string `json:"country"`
	// Contact email address for this address
	Email string `json:"email"`
	// Whether this address is for drop shipping
	IsDropShip bool `json:"is_drop_ship"`
	// Contact phone number for this address
	Phone string `json:"phone"`
	// ZIP or postal code
	PostalCode string `json:"postal_code"`
	// State or province code
	State string `json:"state"`
	// First line of the street address
	StreetLine1 string `json:"street_line_1"`
	// Second line of the street address
	StreetLine2 string `json:"street_line_2"`
	// Timestamp when the address was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		City        respjson.Field
		Country     respjson.Field
		Email       respjson.Field
		IsDropShip  respjson.Field
		Phone       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		StreetLine1 respjson.Field
		StreetLine2 respjson.Field
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

// Response schema for GetCustomerAddressRequest
type GetCustomerAddress struct {
	// ID of the address
	ID string `json:"id,required"`
	// Address line 1 of the address
	AddressLine1 string `json:"address_line_1,required"`
	// City of the address
	City string `json:"city,required"`
	// Country of the address
	Country string `json:"country,required"`
	// created at
	CreatedAt string `json:"created_at,required"`
	// Name of the address
	Name string `json:"name,required"`
	// Postal code of the address
	PostalCode string `json:"postal_code,required"`
	// updated at
	UpdatedAt string `json:"updated_at,required"`
	// Address line 2 of the address
	AddressLine2 string `json:"address_line_2"`
	// Email of the address
	Email string `json:"email"`
	// Is drop ship of the address
	IsDropShip bool `json:"is_drop_ship"`
	// Phone of the address
	Phone string `json:"phone"`
	// State of the address
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AddressLine1 respjson.Field
		City         respjson.Field
		Country      respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		PostalCode   respjson.Field
		UpdatedAt    respjson.Field
		AddressLine2 respjson.Field
		Email        respjson.Field
		IsDropShip   respjson.Field
		Phone        respjson.Field
		State        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetCustomerAddress) RawJSON() string { return r.JSON.raw }
func (r *GetCustomerAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAddressNewParams struct {
	// Address line 1 of the address
	AddressLine1 string `json:"address_line_1,required"`
	// Locality of the address
	City string `json:"city,required"`
	// Country of the address
	Country string `json:"country,required"`
	// Name of the address
	Name string `json:"name,required"`
	// Postal code of the address
	PostalCode string `json:"postal_code,required"`
	// Address line 2 of the address
	AddressLine2 param.Opt[string] `json:"address_line_2,omitzero"`
	// Email of the address
	Email param.Opt[string] `json:"email,omitzero"`
	// Is drop ship of the address
	IsDropShip param.Opt[bool] `json:"is_drop_ship,omitzero"`
	// Phone of the address
	Phone param.Opt[string] `json:"phone,omitzero"`
	// State of the address
	State param.Opt[string] `json:"state,omitzero"`
	// options for return granularity
	Include []any `json:"include,omitzero"`
	paramObj
}

func (r CustomerAddressNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerAddressNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerAddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAddressGetParams struct {
	CustomerID string `path:"customer_id,required" json:"-"`
	paramObj
}

type CustomerAddressUpdateParams struct {
	CustomerID string `path:"customer_id,required" json:"-"`
	// Address line 1 of the address
	AddressLine1 param.Opt[string] `json:"address_line_1,omitzero"`
	// Address line 2 of the address
	AddressLine2 param.Opt[string] `json:"address_line_2,omitzero"`
	// City of the address
	City param.Opt[string] `json:"city,omitzero"`
	// Country of the address
	Country param.Opt[string] `json:"country,omitzero"`
	// Email of the address
	Email param.Opt[string] `json:"email,omitzero"`
	// Is drop ship of the address
	IsDropShip param.Opt[bool] `json:"is_drop_ship,omitzero"`
	// Name of the address
	Name param.Opt[string] `json:"name,omitzero"`
	// Phone of the address
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Postal code of the address
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// State of the address
	State param.Opt[string] `json:"state,omitzero"`
	// options for return granularity
	Include []any `json:"include,omitzero"`
	paramObj
}

func (r CustomerAddressUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerAddressUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerAddressUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAddressDeleteParams struct {
	CustomerID string `path:"customer_id,required" json:"-"`
	paramObj
}
