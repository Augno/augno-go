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
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/respjson"
)

// List and manage item categories.
//
// CatalogItemCategoryPropertyService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemCategoryPropertyService] method instead.
type CatalogItemCategoryPropertyService struct {
	options []option.RequestOption
}

// NewCatalogItemCategoryPropertyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCatalogItemCategoryPropertyService(opts ...option.RequestOption) (r CatalogItemCategoryPropertyService) {
	r = CatalogItemCategoryPropertyService{}
	r.options = opts
	return
}

// Attaches one of your account's properties to an item category.
//
// The property then appears among the category's properties, including in the
// customer-facing catalog, describing a dimension along which the category's items
// vary. Each property name can appear only once per category, so attaching a
// property whose name duplicates one already there returns a conflict error.
//
// This endpoint requires the permission: `item_categories:update`.
func (r *CatalogItemCategoryPropertyService) Update(ctx context.Context, propertyID string, body CatalogItemCategoryPropertyUpdateParams, opts ...option.RequestOption) (res *CatalogItemCategoryPropertyUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if propertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s/properties/%s", url.PathEscape(body.ID), url.PathEscape(propertyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Detaches a property from an item category.
//
// Only the link between the property and the category is removed; the property
// itself and its attributes are left intact and stay available to other
// categories. The property must belong to your account.
//
// This endpoint requires the permission: `item_categories:update`.
func (r *CatalogItemCategoryPropertyService) Delete(ctx context.Context, propertyID string, body CatalogItemCategoryPropertyDeleteParams, opts ...option.RequestOption) (res *CatalogItemCategoryPropertyDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if propertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s/properties/%s", url.PathEscape(body.ID), url.PathEscape(propertyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type CatalogItemCategoryPropertyUpdateResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogItemCategoryPropertyUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogItemCategoryPropertyUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryPropertyDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogItemCategoryPropertyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogItemCategoryPropertyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryPropertyUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type CatalogItemCategoryPropertyDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
