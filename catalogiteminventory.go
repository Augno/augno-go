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

// List and manage inventory items.
//
// CatalogItemInventoryService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemInventoryService] method instead.
type CatalogItemInventoryService struct {
	options []option.RequestOption
}

// NewCatalogItemInventoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogItemInventoryService(opts ...option.RequestOption) (r CatalogItemInventoryService) {
	r = CatalogItemInventoryService{}
	r.options = opts
	return
}

// Adjusts or reconciles the quantity of an item you hold.
//
// With `operation` set to `adjust` (the behavior when it is omitted), `quantity`
// is added to the current quantity; with `reconcile`, the current quantity is set
// to exactly `quantity`. Either way it is the resulting difference that gets
// written, so a difference of zero moves no stock.
//
// The figure a `reconcile` measures against is what is on hand net of demand
// nothing has covered — the same figure `available_to_promise` is derived from,
// not the raw on-hand total. Reconciling to the quantity already reported
// therefore writes nothing.
//
// Stock that arrives is allocated against unfilled demand for the item, so an
// adjustment can settle a shortfall instead of raising the quantity free to
// promise. That allocation happens just after the request rather than inside it,
// because it walks every open issue for the item. The change is recorded in the
// item's inventory audit trail as a user correction attributed to the caller.
//
// This endpoint requires the permission: `items:update`.
func (r *CatalogItemInventoryService) Update(ctx context.Context, id string, body CatalogItemInventoryUpdateParams, opts ...option.RequestOption) (res *CatalogItemInventoryUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/inventory", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns the stock position for an item: what is on hand, what is reserved
// against existing orders, what is free to promise, and what is short.
//
// Stock your account either owns or holds counts toward the on-hand figure, so
// customer-supplied material sitting in your facility is included. All four
// quantities are reported in the base unit of the item's category.
//
// This endpoint requires the permission: `items:read`.
func (r *CatalogItemInventoryService) List(ctx context.Context, id string, query CatalogItemInventoryListParams, opts ...option.RequestOption) (res *ItemInventory, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/inventory", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// An amount calculated on demand rather than stored.
//
// The same shape as a quantity minus the ID, because nothing was written: it is
// derived per request, such as a total rolled up across invoiced lines for one
// analysis.
type ComputedQuantity struct {
	// Formatted value with unit abbreviation (e.g. "1,200 pr").
	DisplayValue string `json:"display_value" api:"required"`
	// Resource type identifier.
	//
	// Any of "computed_quantity".
	Object ComputedQuantityObject `json:"object" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	Unit Unit `json:"unit" api:"required"`
	// Raw decimal value, as a string to preserve precision.
	//
	// This is the unformatted machine value; see `display_value` for the
	// human-readable rendering.
	Value string `json:"value" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayValue respjson.Field
		Object       respjson.Field
		Unit         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputedQuantity) RawJSON() string { return r.JSON.raw }
func (r *ComputedQuantity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ComputedQuantityObject string

const (
	ComputedQuantityObjectComputedQuantity ComputedQuantityObject = "computed_quantity"
)

// The stock position for an item: what is in stock, what is already committed, and
// what is still free to sell.
//
// All four quantities are reported in the same unit — the base unit of the item's
// category. Derived figures, not stored rows: each is netted out of the ledger at
// read time, so none of them carries a quantity id.
type ItemInventory struct {
	// An amount calculated on demand rather than stored.
	//
	// The same shape as a quantity minus the ID, because nothing was written: it is
	// derived per request, such as a total rolled up across invoiced lines for one
	// analysis.
	AvailableToPromise ComputedQuantity `json:"available_to_promise" api:"required"`
	// Resource type identifier.
	//
	// Any of "item_inventory".
	Object ItemInventoryObject `json:"object" api:"required"`
	// An amount calculated on demand rather than stored.
	//
	// The same shape as a quantity minus the ID, because nothing was written: it is
	// derived per request, such as a total rolled up across invoiced lines for one
	// analysis.
	OnHand ComputedQuantity `json:"on_hand" api:"required"`
	// An amount calculated on demand rather than stored.
	//
	// The same shape as a quantity minus the ID, because nothing was written: it is
	// derived per request, such as a total rolled up across invoiced lines for one
	// analysis.
	Reserved ComputedQuantity `json:"reserved" api:"required"`
	// An amount calculated on demand rather than stored.
	//
	// The same shape as a quantity minus the ID, because nothing was written: it is
	// derived per request, such as a total rolled up across invoiced lines for one
	// analysis.
	Short ComputedQuantity `json:"short" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableToPromise respjson.Field
		Object             respjson.Field
		OnHand             respjson.Field
		Reserved           respjson.Field
		Short              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemInventory) RawJSON() string { return r.JSON.raw }
func (r *ItemInventory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ItemInventoryObject string

const (
	ItemInventoryObjectItemInventory ItemInventoryObject = "item_inventory"
)

// An amount together with the unit it is expressed in.
//
// The unit may be a currency, so money amounts such as a credit limit are written
// the same way as physical amounts like weights or counts.
//
// The properties UnitID, Value are required.
type QuantityInputParam struct {
	// ID of the unit of measure for the value.
	UnitID string `json:"unit_id" api:"required"`
	// Decimal value, as a string to preserve precision.
	Value string `json:"value" api:"required" format:"decimal"`
	paramObj
}

func (r QuantityInputParam) MarshalJSON() (data []byte, err error) {
	type shadow QuantityInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuantityInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to adjust or reconcile inventory for an item.
//
// The property Quantity is required.
type UpdateItemInventoryRequestParam struct {
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	Quantity QuantityInputParam `json:"quantity,omitzero" api:"required"`
	// ID of the customer account that owns the resulting inventory.
	//
	// Use this for stock you hold but do not own, such as customer-supplied material.
	// It only affects quantity being added: your account stays the holder, the
	// customer becomes the owner, and the current quantity a `reconcile` measures
	// against is still your account's. Requires edit access to that customer.
	CustomerID param.Opt[string] `json:"customer_id,omitzero"`
	// ID of the location to record the inventory change against.
	//
	// Must be a location in your account.
	LocationID param.Opt[string] `json:"location_id,omitzero"`
	// Lot number to record the inventory change against.
	//
	// The lot is created for the item if it does not already exist.
	LotNumber param.Opt[string] `json:"lot_number,omitzero"`
	// How `quantity` is applied.
	//
	// - `adjust`: adds `quantity` to the current quantity.
	// - `reconcile`: sets the current quantity to exactly `quantity`.
	//
	// Any of "adjust", "reconcile".
	Operation UpdateItemInventoryRequestOperation `json:"operation,omitzero"`
	paramObj
}

func (r UpdateItemInventoryRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateItemInventoryRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateItemInventoryRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How `quantity` is applied.
//
// - `adjust`: adds `quantity` to the current quantity.
// - `reconcile`: sets the current quantity to exactly `quantity`.
type UpdateItemInventoryRequestOperation string

const (
	UpdateItemInventoryRequestOperationAdjust    UpdateItemInventoryRequestOperation = "adjust"
	UpdateItemInventoryRequestOperationReconcile UpdateItemInventoryRequestOperation = "reconcile"
)

type CatalogItemInventoryUpdateResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogItemInventoryUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogItemInventoryUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemInventoryUpdateParams struct {
	// Request to adjust or reconcile inventory for an item.
	UpdateItemInventoryRequest UpdateItemInventoryRequestParam
	paramObj
}

func (r CatalogItemInventoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateItemInventoryRequest)
}
func (r *CatalogItemInventoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemInventoryListParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "on_hand", "reserved", "available_to_promise", "short".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemInventoryListParams]'s query parameters as
// `url.Values`.
func (r CatalogItemInventoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
