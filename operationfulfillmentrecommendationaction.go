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

// The planning assumptions production schedules are solved against, and the
// per-resource overrides that mark which machines constrain the plan.
//
// OperationFulfillmentRecommendationActionService contains methods and other
// services that help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationFulfillmentRecommendationActionService] method instead.
type OperationFulfillmentRecommendationActionService struct {
	options []option.RequestOption
}

// NewOperationFulfillmentRecommendationActionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOperationFulfillmentRecommendationActionService(opts ...option.RequestOption) (r OperationFulfillmentRecommendationActionService) {
	r = OperationFulfillmentRecommendationActionService{}
	r.options = opts
	return
}

// Adopts the recommended fulfillment policy for the named items, writing it as a
// per-item planning override.
//
// The recommendation is recomputed as part of applying it, rather than taken from
// the request. Advice read minutes ago may no longer be the advice — demand moves
// — and writing a stale verdict would set a policy the engine would not give
// today. What comes back is what was actually written.
//
// Takes effect on the next generated schedule; versions already generated keep the
// assumptions they were solved under.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationFulfillmentRecommendationActionService) Apply(ctx context.Context, body OperationFulfillmentRecommendationActionApplyParams, opts ...option.RequestOption) (res *ListFulfillmentRecommendation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/fulfillment-recommendations/actions/apply"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to adopt fulfillment recommendations for specific items.
//
// The property ItemIDs is required.
type ApplyFulfillmentRecommendationsRequestParam struct {
	// Items whose recommendation should be adopted.
	//
	// Named explicitly rather than applied wholesale: adopting advice in bulk without
	// saying what is being adopted is how a plant changes what it builds by accident.
	// Items not named here are left exactly as they are.
	ItemIDs []string `json:"item_ids,omitzero" api:"required"`
	paramObj
}

func (r ApplyFulfillmentRecommendationsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ApplyFulfillmentRecommendationsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplyFulfillmentRecommendationsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationFulfillmentRecommendationActionApplyParams struct {
	// Request to adopt fulfillment recommendations for specific items.
	ApplyFulfillmentRecommendationsRequest ApplyFulfillmentRecommendationsRequestParam
	paramObj
}

func (r OperationFulfillmentRecommendationActionApplyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ApplyFulfillmentRecommendationsRequest)
}
func (r *OperationFulfillmentRecommendationActionApplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
