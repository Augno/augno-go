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

// List and manage service levels (shipping service levels).
//
// OperationCarrierServiceLevelService contains methods and other services that
// help with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationCarrierServiceLevelService] method instead.
type OperationCarrierServiceLevelService struct {
	options []option.RequestOption
}

// NewOperationCarrierServiceLevelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationCarrierServiceLevelService(opts ...option.RequestOption) (r OperationCarrierServiceLevelService) {
	r = OperationCarrierServiceLevelService{}
	r.options = opts
	return
}

// Creates a service level for a carrier.
func (r *OperationCarrierServiceLevelService) New(ctx context.Context, carrierID string, params OperationCarrierServiceLevelNewParams, opts ...option.RequestOption) (res *ServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if carrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels", url.PathEscape(carrierID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a service level by ID.
func (r *OperationCarrierServiceLevelService) Get(ctx context.Context, id string, params OperationCarrierServiceLevelGetParams, opts ...option.RequestOption) (res *ServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CarrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels/%s", url.PathEscape(params.CarrierID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Partially updates a service level.
func (r *OperationCarrierServiceLevelService) Update(ctx context.Context, id string, params OperationCarrierServiceLevelUpdateParams, opts ...option.RequestOption) (res *ServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CarrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels/%s", url.PathEscape(params.CarrierID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of service levels for a carrier.
func (r *OperationCarrierServiceLevelService) List(ctx context.Context, carrierID string, query OperationCarrierServiceLevelListParams, opts ...option.RequestOption) (res *ListServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if carrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels", url.PathEscape(carrierID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a service level. Fails if the service level is a default
// (system-synced) level.
func (r *OperationCarrierServiceLevelService) Delete(ctx context.Context, id string, body OperationCarrierServiceLevelDeleteParams, opts ...option.RequestOption) (res *OperationCarrierServiceLevelDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CarrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels/%s", url.PathEscape(body.CarrierID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a service level.
//
// The properties Code, IsDefault, Name are required.
type CreateServiceLevelRequestParam struct {
	// Service level code.
	Code string `json:"code" api:"required"`
	// Default service levels are the default-selected service level for that carrier.
	IsDefault bool `json:"is_default" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Whether this service level will be available for customers to select in the
	// customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CreateServiceLevelRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r CreateServiceLevelRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateServiceLevelRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateServiceLevelRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this service level will be available for customers to select in the
// customer portal.
type CreateServiceLevelRequestCustomerPortalVisibility string

const (
	CreateServiceLevelRequestCustomerPortalVisibilityVisible CreateServiceLevelRequestCustomerPortalVisibility = "visible"
	CreateServiceLevelRequestCustomerPortalVisibilityHidden  CreateServiceLevelRequestCustomerPortalVisibility = "hidden"
)

// Request to update a service level.
type UpdateServiceLevelRequestParam struct {
	// Service level code.
	Code param.Opt[string] `json:"code,omitzero"`
	// Default service levels are the default-selected service level for that carrier.
	IsDefault param.Opt[bool] `json:"is_default,omitzero"`
	// Display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether this service level will be available for customers to select in the
	// customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility UpdateServiceLevelRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r UpdateServiceLevelRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateServiceLevelRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateServiceLevelRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this service level will be available for customers to select in the
// customer portal.
type UpdateServiceLevelRequestCustomerPortalVisibility string

const (
	UpdateServiceLevelRequestCustomerPortalVisibilityVisible UpdateServiceLevelRequestCustomerPortalVisibility = "visible"
	UpdateServiceLevelRequestCustomerPortalVisibilityHidden  UpdateServiceLevelRequestCustomerPortalVisibility = "hidden"
)

type OperationCarrierServiceLevelDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationCarrierServiceLevelDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationCarrierServiceLevelDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationCarrierServiceLevelNewParams struct {
	// Request to create a service level.
	CreateServiceLevelRequest CreateServiceLevelRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationCarrierServiceLevelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateServiceLevelRequest)
}
func (r *OperationCarrierServiceLevelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierServiceLevelNewParams]'s query parameters
// as `url.Values`.
func (r OperationCarrierServiceLevelNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelGetParams struct {
	CarrierID string `path:"carrier_id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierServiceLevelGetParams]'s query parameters
// as `url.Values`.
func (r OperationCarrierServiceLevelGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelUpdateParams struct {
	CarrierID string `path:"carrier_id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a service level.
	UpdateServiceLevelRequest UpdateServiceLevelRequestParam
	paramObj
}

func (r OperationCarrierServiceLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateServiceLevelRequest)
}
func (r *OperationCarrierServiceLevelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierServiceLevelUpdateParams]'s query
// parameters as `url.Values`.
func (r OperationCarrierServiceLevelUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierServiceLevelListParams]'s query parameters
// as `url.Values`.
func (r OperationCarrierServiceLevelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelDeleteParams struct {
	CarrierID string `path:"carrier_id" api:"required" json:"-"`
	paramObj
}
