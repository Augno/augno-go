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

// List and manage properties and their attributes.
//
// CatalogPropertyAttributeService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogPropertyAttributeService] method instead.
type CatalogPropertyAttributeService struct {
	options []option.RequestOption
}

// NewCatalogPropertyAttributeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCatalogPropertyAttributeService(opts ...option.RequestOption) (r CatalogPropertyAttributeService) {
	r = CatalogPropertyAttributeService{}
	r.options = opts
	return
}

// Creates an attribute under a property.
func (r *CatalogPropertyAttributeService) New(ctx context.Context, propertyID string, body CatalogPropertyAttributeNewParams, opts ...option.RequestOption) (res *Attribute, err error) {
	opts = slices.Concat(r.options, opts)
	if propertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s/attributes", url.PathEscape(propertyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns an attribute by ID within a property.
func (r *CatalogPropertyAttributeService) Get(ctx context.Context, id string, query CatalogPropertyAttributeGetParams, opts ...option.RequestOption) (res *Attribute, err error) {
	opts = slices.Concat(r.options, opts)
	if query.PropertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s/attributes/%s", url.PathEscape(query.PropertyID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially updates an attribute.
func (r *CatalogPropertyAttributeService) Update(ctx context.Context, id string, params CatalogPropertyAttributeUpdateParams, opts ...option.RequestOption) (res *Attribute, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PropertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s/attributes/%s", url.PathEscape(params.PropertyID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of attributes for a property.
func (r *CatalogPropertyAttributeService) List(ctx context.Context, propertyID string, query CatalogPropertyAttributeListParams, opts ...option.RequestOption) (res *ListAttribute, err error) {
	opts = slices.Concat(r.options, opts)
	if propertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s/attributes", url.PathEscape(propertyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an attribute from a property.
//
// Remaining attributes in the property are shifted so their sort orders stay
// contiguous.
func (r *CatalogPropertyAttributeService) Delete(ctx context.Context, id string, body CatalogPropertyAttributeDeleteParams, opts ...option.RequestOption) (res *CatalogPropertyAttributeDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.PropertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/properties/%s/attributes/%s", url.PathEscape(body.PropertyID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create an attribute.
//
// The property Value is required.
type CreateAttributeRequestParam struct {
	// The selectable value this attribute represents, such as `Red`.
	//
	// Must be unique across all attributes in the account, not just within the
	// property. Leading and trailing whitespace is trimmed.
	Value string `json:"value" api:"required"`
	// Position of the new attribute relative to its siblings within the property,
	// starting at `1`.
	//
	// Must be at most the property's current attribute count plus one; siblings at or
	// after this position are shifted one position later. Defaults to the last
	// position if not provided.
	SortOrder param.Opt[int64] `json:"sort_order,omitzero"`
	// Swatch color used to display this attribute in the UI.
	//
	// When omitted, one of the nine named colors (everything except `default`) is
	// assigned at random.
	//
	// Any of "blue", "brown", "default", "gray", "green", "orange", "pink", "purple",
	// "red", "yellow".
	Color CreateAttributeRequestColor `json:"color,omitzero"`
	paramObj
}

func (r CreateAttributeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAttributeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAttributeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swatch color used to display this attribute in the UI.
//
// When omitted, one of the nine named colors (everything except `default`) is
// assigned at random.
type CreateAttributeRequestColor string

const (
	CreateAttributeRequestColorBlue    CreateAttributeRequestColor = "blue"
	CreateAttributeRequestColorBrown   CreateAttributeRequestColor = "brown"
	CreateAttributeRequestColorDefault CreateAttributeRequestColor = "default"
	CreateAttributeRequestColorGray    CreateAttributeRequestColor = "gray"
	CreateAttributeRequestColorGreen   CreateAttributeRequestColor = "green"
	CreateAttributeRequestColorOrange  CreateAttributeRequestColor = "orange"
	CreateAttributeRequestColorPink    CreateAttributeRequestColor = "pink"
	CreateAttributeRequestColorPurple  CreateAttributeRequestColor = "purple"
	CreateAttributeRequestColorRed     CreateAttributeRequestColor = "red"
	CreateAttributeRequestColorYellow  CreateAttributeRequestColor = "yellow"
)

// Request to update an attribute.
type UpdateAttributeRequestParam struct {
	// New position of this attribute relative to its siblings within the property,
	// starting at `1`.
	//
	// Must be at most the property's current attribute count; the attributes between
	// the old and new positions shift to make room.
	SortOrder param.Opt[int64] `json:"sort_order,omitzero"`
	// The selectable value this attribute represents, such as `Red`.
	//
	// Must be non-blank and unique across all attributes in the account, not just
	// within the property.
	Value param.Opt[string] `json:"value,omitzero"`
	// Swatch color used to display this attribute in the UI.
	//
	// Any of "blue", "brown", "default", "gray", "green", "orange", "pink", "purple",
	// "red", "yellow".
	Color UpdateAttributeRequestColor `json:"color,omitzero"`
	paramObj
}

func (r UpdateAttributeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAttributeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAttributeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swatch color used to display this attribute in the UI.
type UpdateAttributeRequestColor string

const (
	UpdateAttributeRequestColorBlue    UpdateAttributeRequestColor = "blue"
	UpdateAttributeRequestColorBrown   UpdateAttributeRequestColor = "brown"
	UpdateAttributeRequestColorDefault UpdateAttributeRequestColor = "default"
	UpdateAttributeRequestColorGray    UpdateAttributeRequestColor = "gray"
	UpdateAttributeRequestColorGreen   UpdateAttributeRequestColor = "green"
	UpdateAttributeRequestColorOrange  UpdateAttributeRequestColor = "orange"
	UpdateAttributeRequestColorPink    UpdateAttributeRequestColor = "pink"
	UpdateAttributeRequestColorPurple  UpdateAttributeRequestColor = "purple"
	UpdateAttributeRequestColorRed     UpdateAttributeRequestColor = "red"
	UpdateAttributeRequestColorYellow  UpdateAttributeRequestColor = "yellow"
)

type CatalogPropertyAttributeDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogPropertyAttributeDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogPropertyAttributeDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPropertyAttributeNewParams struct {
	// Request to create an attribute.
	CreateAttributeRequest CreateAttributeRequestParam
	paramObj
}

func (r CatalogPropertyAttributeNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAttributeRequest)
}
func (r *CatalogPropertyAttributeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPropertyAttributeGetParams struct {
	PropertyID string `path:"property_id" api:"required" json:"-"`
	paramObj
}

type CatalogPropertyAttributeUpdateParams struct {
	PropertyID string `path:"property_id" api:"required" json:"-"`
	// Request to update an attribute.
	UpdateAttributeRequest UpdateAttributeRequestParam
	paramObj
}

func (r CatalogPropertyAttributeUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAttributeRequest)
}
func (r *CatalogPropertyAttributeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPropertyAttributeListParams struct {
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

// URLQuery serializes [CatalogPropertyAttributeListParams]'s query parameters as
// `url.Values`.
func (r CatalogPropertyAttributeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPropertyAttributeDeleteParams struct {
	PropertyID string `path:"property_id" api:"required" json:"-"`
	paramObj
}
