// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
)

// List and manage account prices.
//
// SaleAccountPriceActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleAccountPriceActionService] method instead.
type SaleAccountPriceActionService struct {
	options []option.RequestOption
}

// NewSaleAccountPriceActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleAccountPriceActionService(opts ...option.RequestOption) (r SaleAccountPriceActionService) {
	r = SaleAccountPriceActionService{}
	r.options = opts
	return
}

// Starts a customer's price list and returns the job that tracks it.
//
// The document covers every product the customer may order, grouped by product
// line and then by the SKUs that share a price, with the attributes that vary
// shown as columns. Prices are calculated by the same engine that prices a sales
// order, so they include the customer's contracted prices and any volume discount
// they qualify for; a volume break becomes its own price column only where it
// actually changes a price.
//
// Pricing a whole catalog takes too long to hold a request open for, so the PDF is
// rendered in the background. Poll the returned job and download the file it names
// once it completes.
//
// This endpoint requires the permission: `discounts:read`.
func (r *SaleAccountPriceActionService) ExportPriceList(ctx context.Context, params SaleAccountPriceActionExportPriceListParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/account-prices/actions/export-price-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to export a customer's price list.
//
// The property CustomerID is required.
type ExportPriceListRequestParam struct {
	// ID of the customer whose prices are listed.
	CustomerID string `json:"customer_id" api:"required"`
	paramObj
}

func (r ExportPriceListRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportPriceListRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPriceListRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountPriceActionExportPriceListParams struct {
	// Request to export a customer's price list.
	ExportPriceListRequest ExportPriceListRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleAccountPriceActionExportPriceListParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExportPriceListRequest)
}
func (r *SaleAccountPriceActionExportPriceListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleAccountPriceActionExportPriceListParams]'s query
// parameters as `url.Values`.
func (r SaleAccountPriceActionExportPriceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
