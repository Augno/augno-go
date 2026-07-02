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

// Adds a property to an item category, making the property available to items in
// that category.
//
// Each property name can appear only once per category; adding a property whose
// name duplicates one already in the category returns a conflict error. Default
// system categories cannot be modified.
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

// Removes a property from an item category.
//
// Default system categories cannot be modified.
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
