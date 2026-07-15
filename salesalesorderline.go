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

// List, view, create, update, and delete sales orders.
//
// SaleSalesOrderLineService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderLineService] method instead.
type SaleSalesOrderLineService struct {
	options []option.RequestOption
	// List, view, create, update, and delete sales orders.
	Actions SaleSalesOrderLineActionService
}

// NewSaleSalesOrderLineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleSalesOrderLineService(opts ...option.RequestOption) (r SaleSalesOrderLineService) {
	r = SaleSalesOrderLineService{}
	r.options = opts
	r.Actions = NewSaleSalesOrderLineActionService(opts...)
	return
}

// Creates a line item on a sales order.
//
// This endpoint requires the permissions: `customers:update`, `suppliers:update`,
// `sales_orders:update`.
func (r *SaleSalesOrderLineService) New(ctx context.Context, id string, params SaleSalesOrderLineNewParams, opts ...option.RequestOption) (res *SalesOrderLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/lines", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Partially updates a sales order line item.
//
// This endpoint requires the permissions: `customers:update`, `suppliers:update`,
// `sales_orders:update`.
func (r *SaleSalesOrderLineService) Update(ctx context.Context, lineID string, params SaleSalesOrderLineUpdateParams, opts ...option.RequestOption) (res *SalesOrderLine, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if lineID == "" {
		err = errors.New("missing required line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/lines/%s", url.PathEscape(params.ID), url.PathEscape(lineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Deletes a sales order line and related records.
//
// This endpoint requires the permissions: `customers:update`, `suppliers:update`,
// `sales_orders:update`.
func (r *SaleSalesOrderLineService) Delete(ctx context.Context, lineID string, body SaleSalesOrderLineDeleteParams, opts ...option.RequestOption) (res *SaleSalesOrderLineDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if lineID == "" {
		err = errors.New("missing required line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/lines/%s", url.PathEscape(body.ID), url.PathEscape(lineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a line on a sales order.
//
// This mirrors the create-order line shape (not the shared purchase-order
// OrderLineInput): the unit price is optional and, when omitted, the line is
// priced server-side from the product. The unit cost is always resolved
// server-side from the product.
//
// The properties ProductID, ProductSKU, Quantity are required.
type CreateSalesOrderLineRequestParam struct {
	// ID of the product being ordered.
	ProductID string `json:"product_id" api:"required"`
	// The product SKU recorded on the line.
	ProductSKU string `json:"product_sku" api:"required"`
	// A value with an associated unit, used in create and update requests.
	Quantity QuantityInputParam `json:"quantity,omitzero" api:"required"`
	// The product description recorded on the line.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// A rate value with its numerator and denominator units, used in create and update
	// requests.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r CreateSalesOrderLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSalesOrderLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSalesOrderLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to update a sales order line.
type UpdateSalesOrderLineRequestParam struct {
	// Product description.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// Product SKU.
	ProductSKU param.Opt[string] `json:"product_sku,omitzero"`
	// A value with an associated unit, used in create and update requests.
	Quantity QuantityInputParam `json:"quantity,omitzero"`
	// A rate value with its numerator and denominator units, used in create and update
	// requests.
	UnitCost RateInputParam `json:"unit_cost,omitzero"`
	// A rate value with its numerator and denominator units, used in create and update
	// requests.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r UpdateSalesOrderLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateSalesOrderLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateSalesOrderLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderLineDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleSalesOrderLineDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleSalesOrderLineDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderLineNewParams struct {
	// Request to create a line on a sales order.
	//
	// This mirrors the create-order line shape (not the shared purchase-order
	// OrderLineInput): the unit price is optional and, when omitted, the line is
	// priced server-side from the product. The unit cost is always resolved
	// server-side from the product.
	CreateSalesOrderLineRequest CreateSalesOrderLineRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product", "quantity_ordered", "unit_price", "unit_cost", "totals".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleSalesOrderLineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSalesOrderLineRequest)
}
func (r *SaleSalesOrderLineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderLineNewParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderLineNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderLineUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product", "quantity_ordered", "unit_price", "unit_cost", "totals".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a sales order line.
	UpdateSalesOrderLineRequest UpdateSalesOrderLineRequestParam
	paramObj
}

func (r SaleSalesOrderLineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateSalesOrderLineRequest)
}
func (r *SaleSalesOrderLineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderLineUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderLineUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderLineDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
