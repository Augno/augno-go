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
)

// List and manage properties and their attributes.
//
// CatalogPropertyActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogPropertyActionService] method instead.
type CatalogPropertyActionService struct {
	options []option.RequestOption
}

// NewCatalogPropertyActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogPropertyActionService(opts ...option.RequestOption) (r CatalogPropertyActionService) {
	r = CatalogPropertyActionService{}
	r.options = opts
	return
}

// Creates or updates multiple properties and their attributes for the account,
// matched by name (case-insensitive), then writes asynchronously — 202 with a job
// to poll.
func (r *CatalogPropertyActionService) BulkUpsert(ctx context.Context, body CatalogPropertyActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/properties/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// carries the properties to bulk upsert
//
// The property Properties is required.
type BulkUpsertPropertiesRequestParam struct {
	// Properties to create or update, matched by name (case-insensitive) within the
	// account.
	Properties []UpsertPropertyInputParam `json:"properties,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertPropertiesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertPropertiesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertPropertiesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// carries one attribute under a bulk-upserted property
//
// The property Value is required.
type UpsertPropertyAttributeInputParam struct {
	// The selectable value this attribute represents, such as `Red`.
	//
	// Must be unique across all attributes in the account, not just within the
	// property. Leading and trailing whitespace is trimmed.
	Value string `json:"value" api:"required"`
	// Swatch color used to display this attribute in the UI.
	//
	// When omitted, one of the nine named colors is assigned. Ignored for a value the
	// property already defines.
	//
	// Any of "blue", "brown", "default", "gray", "green", "orange", "pink", "purple",
	// "red", "yellow".
	Color UpsertPropertyAttributeInputColor `json:"color,omitzero"`
	paramObj
}

func (r UpsertPropertyAttributeInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertPropertyAttributeInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertPropertyAttributeInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swatch color used to display this attribute in the UI.
//
// When omitted, one of the nine named colors is assigned. Ignored for a value the
// property already defines.
type UpsertPropertyAttributeInputColor string

const (
	UpsertPropertyAttributeInputColorBlue    UpsertPropertyAttributeInputColor = "blue"
	UpsertPropertyAttributeInputColorBrown   UpsertPropertyAttributeInputColor = "brown"
	UpsertPropertyAttributeInputColorDefault UpsertPropertyAttributeInputColor = "default"
	UpsertPropertyAttributeInputColorGray    UpsertPropertyAttributeInputColor = "gray"
	UpsertPropertyAttributeInputColorGreen   UpsertPropertyAttributeInputColor = "green"
	UpsertPropertyAttributeInputColorOrange  UpsertPropertyAttributeInputColor = "orange"
	UpsertPropertyAttributeInputColorPink    UpsertPropertyAttributeInputColor = "pink"
	UpsertPropertyAttributeInputColorPurple  UpsertPropertyAttributeInputColor = "purple"
	UpsertPropertyAttributeInputColorRed     UpsertPropertyAttributeInputColor = "red"
	UpsertPropertyAttributeInputColorYellow  UpsertPropertyAttributeInputColor = "yellow"
)

// carries one property in a bulk upsert
//
// The properties Attributes, Name are required.
type UpsertPropertyInputParam struct {
	// The selectable values to define under this property, in the order they should be
	// arranged.
	//
	// Additive — values the property already defines are left as they stand, and none
	// are removed. New values are appended after the existing ones.
	Attributes []UpsertPropertyAttributeInputParam `json:"attributes,omitzero" api:"required"`
	// Display name of the property, used to match existing properties within the
	// account.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r UpsertPropertyInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertPropertyInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertPropertyInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPropertyActionBulkUpsertParams struct {
	// carries the properties to bulk upsert
	BulkUpsertPropertiesRequest BulkUpsertPropertiesRequestParam
	paramObj
}

func (r CatalogPropertyActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertPropertiesRequest)
}
func (r *CatalogPropertyActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
