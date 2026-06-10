// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// CatalogService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogService] method instead.
type CatalogService struct {
	options []option.RequestOption
	// List and manage units.
	Units CatalogUnitService
	// List and manage unit groups and their associated units.
	UnitGroups CatalogUnitGroupService
	// List and manage properties and their attributes.
	Properties CatalogPropertyService
	// List and manage inventory items.
	Items CatalogItemService
	// List and manage item categories.
	ItemCategories CatalogItemCategoryService
	// List and manage materials.
	Materials CatalogMaterialService
	// List and manage parts.
	Parts CatalogPartService
	// List and manage product lines.
	ProductLines CatalogProductLineService
	// List and manage products.
	Products CatalogProductService
}

// NewCatalogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCatalogService(opts ...option.RequestOption) (r CatalogService) {
	r = CatalogService{}
	r.options = opts
	r.Units = NewCatalogUnitService(opts...)
	r.UnitGroups = NewCatalogUnitGroupService(opts...)
	r.Properties = NewCatalogPropertyService(opts...)
	r.Items = NewCatalogItemService(opts...)
	r.ItemCategories = NewCatalogItemCategoryService(opts...)
	r.Materials = NewCatalogMaterialService(opts...)
	r.Parts = NewCatalogPartService(opts...)
	r.ProductLines = NewCatalogProductLineService(opts...)
	r.Products = NewCatalogProductService(opts...)
	return
}
