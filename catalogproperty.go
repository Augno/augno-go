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

// List and manage properties and their attributes.
//
// CatalogPropertyService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogPropertyService] method instead.
type CatalogPropertyService struct {
	options []option.RequestOption
	// List and manage properties and their attributes.
	Attributes CatalogPropertyAttributeService
}

// NewCatalogPropertyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogPropertyService(opts ...option.RequestOption) (r CatalogPropertyService) {
	r = CatalogPropertyService{}
	r.options = opts
	r.Attributes = NewCatalogPropertyAttributeService(opts...)
	return
}

// Creates a property.
func (r *CatalogPropertyService) New(ctx context.Context, params CatalogPropertyNewParams, opts ...option.RequestOption) (res *Property, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/properties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a property by ID.
func (r *CatalogPropertyService) Get(ctx context.Context, id string, query CatalogPropertyGetParams, opts ...option.RequestOption) (res *Property, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a property.
func (r *CatalogPropertyService) Update(ctx context.Context, id string, params CatalogPropertyUpdateParams, opts ...option.RequestOption) (res *Property, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of properties for the target account.
func (r *CatalogPropertyService) List(ctx context.Context, query CatalogPropertyListParams, opts ...option.RequestOption) (res *ListProperty, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/properties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a property and all associated attributes.
func (r *CatalogPropertyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CatalogPropertyDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Value option within a property.
type Attribute struct {
	// Attribute ID.
	ID string `json:"id" api:"required"`
	// Color code.
	//
	// Any of "blue", "brown", "default", "gray", "green", "orange", "pink", "purple",
	// "red", "yellow".
	Color AttributeColor `json:"color" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "attribute".
	Object AttributeObject `json:"object" api:"required"`
	// Property that groups attributes.
	Property *Property `json:"property" api:"required"`
	// Display order.
	SortOrder int64 `json:"sort_order" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attribute value.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Color       respjson.Field
		CreatedAt   respjson.Field
		Object      respjson.Field
		Property    respjson.Field
		SortOrder   respjson.Field
		UpdatedAt   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Attribute) RawJSON() string { return r.JSON.raw }
func (r *Attribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Color code.
type AttributeColor string

const (
	AttributeColorBlue    AttributeColor = "blue"
	AttributeColorBrown   AttributeColor = "brown"
	AttributeColorDefault AttributeColor = "default"
	AttributeColorGray    AttributeColor = "gray"
	AttributeColorGreen   AttributeColor = "green"
	AttributeColorOrange  AttributeColor = "orange"
	AttributeColorPink    AttributeColor = "pink"
	AttributeColorPurple  AttributeColor = "purple"
	AttributeColorRed     AttributeColor = "red"
	AttributeColorYellow  AttributeColor = "yellow"
)

// Resource type identifier.
type AttributeObject string

const (
	AttributeObjectAttribute AttributeObject = "attribute"
)

// Request to create a property.
//
// The property Name is required.
type CreatePropertyRequestParam struct {
	// Name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r CreatePropertyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePropertyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePropertyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List represents a paginated list of resources.
type ListAttribute struct {
	// Resources in this page.
	Data []Attribute `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAttributeObject `json:"object" api:"required"`
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
func (r ListAttribute) RawJSON() string { return r.JSON.raw }
func (r *ListAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAttributeObject string

const (
	ListAttributeObjectList ListAttributeObject = "list"
)

// List represents a paginated list of resources.
type ListProperty struct {
	// Resources in this page.
	Data []Property `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPropertyObject `json:"object" api:"required"`
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
func (r ListProperty) RawJSON() string { return r.JSON.raw }
func (r *ListProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPropertyObject string

const (
	ListPropertyObjectList ListPropertyObject = "list"
)

// Property that groups attributes.
type Property struct {
	// Property ID.
	ID string `json:"id" api:"required"`
	// List represents a paginated list of resources.
	Attributes *ListAttribute `json:"attributes" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "property".
	Object PropertyObject `json:"object" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Property) RawJSON() string { return r.JSON.raw }
func (r *Property) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PropertyObject string

const (
	PropertyObjectProperty PropertyObject = "property"
)

// Request to update a property.
type UpdatePropertyRequestParam struct {
	// Name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r UpdatePropertyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdatePropertyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdatePropertyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPropertyDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogPropertyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogPropertyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPropertyNewParams struct {
	// Request to create a property.
	CreatePropertyRequest CreatePropertyRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogPropertyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePropertyRequest)
}
func (r *CatalogPropertyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogPropertyNewParams]'s query parameters as
// `url.Values`.
func (r CatalogPropertyNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPropertyGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogPropertyGetParams]'s query parameters as
// `url.Values`.
func (r CatalogPropertyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPropertyUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a property.
	UpdatePropertyRequest UpdatePropertyRequestParam
	paramObj
}

func (r CatalogPropertyUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdatePropertyRequest)
}
func (r *CatalogPropertyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogPropertyUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogPropertyUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPropertyListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogPropertyListParams]'s query parameters as
// `url.Values`.
func (r CatalogPropertyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
