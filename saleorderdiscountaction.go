// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
)

// List and manage order discounts.
//
// SaleOrderDiscountActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleOrderDiscountActionService] method instead.
type SaleOrderDiscountActionService struct {
	options []option.RequestOption
}

// NewSaleOrderDiscountActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleOrderDiscountActionService(opts ...option.RequestOption) (r SaleOrderDiscountActionService) {
	r = SaleOrderDiscountActionService{}
	r.options = opts
	return
}

// Validates a discount code and returns the matching order discount, so a code a
// buyer typed can be attached to an order.
//
// When `buyer_account_id` is provided, or the caller is a customer user, the
// lookup also verifies that the buyer has not already redeemed the discount on
// another order, and reports an already-redeemed code as not found. Pass
// `sales_order_id` to exclude an order the buyer is currently editing from that
// check.
//
// This endpoint requires the permission: `discounts:read`.
func (r *SaleOrderDiscountActionService) FindByCode(ctx context.Context, body SaleOrderDiscountActionFindByCodeParams, opts ...option.RequestOption) (res *OrderDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/order-discounts/actions/find-by-code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to find an order discount by code.
//
// The property Code is required.
type FindOrderDiscountByCodeRequestParam struct {
	// The discount code to look up, as the buyer typed it.
	//
	// Matching ignores letter case, so `save10` finds a discount stored as `SAVE10`.
	Code string `json:"code" api:"required"`
	// The buyer account to check for prior use of this code.
	//
	// When set, the lookup returns a not-found error if that buyer has already
	// redeemed the discount on another order, so a one-use-per-customer code can be
	// rejected before it is attached to a new one. Customer callers cannot set this —
	// their own account is always used.
	BuyerAccountID param.Opt[string] `json:"buyer_account_id,omitzero"`
	// Sales order ID to exclude from the prior-usage check.
	//
	// Set this when re-validating a code on an existing order so the order's own usage
	// does not count against the buyer.
	SalesOrderID param.Opt[string] `json:"sales_order_id,omitzero"`
	paramObj
}

func (r FindOrderDiscountByCodeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FindOrderDiscountByCodeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FindOrderDiscountByCodeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleOrderDiscountActionFindByCodeParams struct {
	// Request to find an order discount by code.
	FindOrderDiscountByCodeRequest FindOrderDiscountByCodeRequestParam
	paramObj
}

func (r SaleOrderDiscountActionFindByCodeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FindOrderDiscountByCodeRequest)
}
func (r *SaleOrderDiscountActionFindByCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
