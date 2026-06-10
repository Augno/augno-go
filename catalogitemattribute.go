// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apiquery"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
)

// List and manage inventory items.
//
// CatalogItemAttributeService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemAttributeService] method instead.
type CatalogItemAttributeService struct {
	options []option.RequestOption
}

// NewCatalogItemAttributeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogItemAttributeService(opts ...option.RequestOption) (r CatalogItemAttributeService) {
	r = CatalogItemAttributeService{}
	r.options = opts
	return
}

// Adds an attribute to an item. If the attribute is already associated with the
// item, this is a no-op.
func (r *CatalogItemAttributeService) Update(ctx context.Context, attributeID string, params CatalogItemAttributeUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if attributeID == "" {
		err = errors.New("missing required attribute_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/attributes/%s", url.PathEscape(params.ID), url.PathEscape(attributeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Removes an attribute from an item.
func (r *CatalogItemAttributeService) Delete(ctx context.Context, attributeID string, params CatalogItemAttributeDeleteParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if attributeID == "" {
		err = errors.New("missing required attribute_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/attributes/%s", url.PathEscape(params.ID), url.PathEscape(attributeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

type CatalogItemAttributeUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemAttributeUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogItemAttributeUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemAttributeDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemAttributeDeleteParams]'s query parameters as
// `url.Values`.
func (r CatalogItemAttributeDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
