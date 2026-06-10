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
)

// Manage customer accounts.
//
// SaleCustomerActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleCustomerActionService] method instead.
type SaleCustomerActionService struct {
	options []option.RequestOption
}

// NewSaleCustomerActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleCustomerActionService(opts ...option.RequestOption) (r SaleCustomerActionService) {
	r = SaleCustomerActionService{}
	r.options = opts
	return
}

// Merges one or more source customers into a target customer, reassigning all
// associated records and deleting the source accounts.
func (r *SaleCustomerActionService) Merge(ctx context.Context, id string, params SaleCustomerActionMergeParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s/actions/merge", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to merge source customers into a target customer.
//
// The property SourceCustomerIDs is required.
type MergeCustomersRequestParam struct {
	// Source customer IDs.
	SourceCustomerIDs []string `json:"source_customer_ids,omitzero" api:"required"`
	paramObj
}

func (r MergeCustomersRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MergeCustomersRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeCustomersRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleCustomerActionMergeParams struct {
	// Request to merge source customers into a target customer.
	MergeCustomersRequest MergeCustomersRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "bill_to_address", "ship_to_address", "type", "parent_account",
	// "freight_preferences.carrier", "freight_preferences.service_level",
	// "defaults.payment_term", "defaults.shipping_term", "defaults.sales_rep",
	// "defaults.sales_rep.user", "defaults.priority", "contact_info",
	// "freight_preferences", "defaults", "notification_preferences", "price_groups",
	// "child_accounts", "credit_limit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleCustomerActionMergeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MergeCustomersRequest)
}
func (r *SaleCustomerActionMergeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleCustomerActionMergeParams]'s query parameters as
// `url.Values`.
func (r SaleCustomerActionMergeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
