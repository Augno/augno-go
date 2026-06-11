// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List and retrieve account statuses.
//
// SaleAccountStatusService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleAccountStatusService] method instead.
type SaleAccountStatusService struct {
	options []option.RequestOption
}

// NewSaleAccountStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleAccountStatusService(opts ...option.RequestOption) (r SaleAccountStatusService) {
	r = SaleAccountStatusService{}
	r.options = opts
	return
}

// Returns an account status by ID or code.
func (r *SaleAccountStatusService) Get(ctx context.Context, id string, query SaleAccountStatusGetParams, opts ...option.RequestOption) (res *AccountStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-statuses/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of account statuses. Global lookup values for setting
// account relationship statuses.
func (r *SaleAccountStatusService) List(ctx context.Context, query SaleAccountStatusListParams, opts ...option.RequestOption) (res *ListAccountStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/account-statuses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// AccountStatus is an account status lookup value.
type AccountStatus struct {
	// Account status ID.
	ID string `json:"id" api:"required"`
	// Machine-readable status code.
	//
	//   - `normal`: standard account with no restrictions.
	//   - `preferred`: account flagged as preferred (e.g. for prioritized handling).
	//   - `hold_shipment`: shipments to this account are held; orders may still be
	//     placed.
	//   - `hold_all`: all activity for this account is held.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Code AccountStatusCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_status".
	Object AccountStatusObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountStatus) RawJSON() string { return r.JSON.raw }
func (r *AccountStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable status code.
//
//   - `normal`: standard account with no restrictions.
//   - `preferred`: account flagged as preferred (e.g. for prioritized handling).
//   - `hold_shipment`: shipments to this account are held; orders may still be
//     placed.
//   - `hold_all`: all activity for this account is held.
type AccountStatusCode string

const (
	AccountStatusCodeNormal       AccountStatusCode = "normal"
	AccountStatusCodePreferred    AccountStatusCode = "preferred"
	AccountStatusCodeHoldShipment AccountStatusCode = "hold_shipment"
	AccountStatusCodeHoldAll      AccountStatusCode = "hold_all"
)

// Resource type identifier.
type AccountStatusObject string

const (
	AccountStatusObjectAccountStatus AccountStatusObject = "account_status"
)

// List represents a paginated list of resources.
type ListAccountStatus struct {
	// Resources in this page.
	Data []AccountStatus `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAccountStatusObject `json:"object" api:"required"`
	// PageInfo contains URL-based pagination metadata.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListAccountStatus) RawJSON() string { return r.JSON.raw }
func (r *ListAccountStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAccountStatusObject string

const (
	ListAccountStatusObjectList ListAccountStatusObject = "list"
)

type SaleAccountStatusGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleAccountStatusGetParams]'s query parameters as
// `url.Values`.
func (r SaleAccountStatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleAccountStatusListParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleAccountStatusListParams]'s query parameters as
// `url.Values`.
func (r SaleAccountStatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
