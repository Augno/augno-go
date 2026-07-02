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
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// List and manage account groups.
//
// SaleAccountGroupService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleAccountGroupService] method instead.
type SaleAccountGroupService struct {
	options []option.RequestOption
}

// NewSaleAccountGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleAccountGroupService(opts ...option.RequestOption) (r SaleAccountGroupService) {
	r = SaleAccountGroupService{}
	r.options = opts
	return
}

// Creates an account group.
//
// Returns a conflict error if an account group with the same name already exists.
//
// This endpoint requires the permission: `customer_groups:create`.
func (r *SaleAccountGroupService) New(ctx context.Context, body SaleAccountGroupNewParams, opts ...option.RequestOption) (res *AccountGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/account-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns an account group by ID.
//
// This endpoint requires the permission: `customer_groups:read`.
func (r *SaleAccountGroupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially updates an account group.
//
// Only the provided fields are changed. The account group's `type` cannot be
// changed after creation.
//
// This endpoint requires the permission: `customer_groups:update`.
func (r *SaleAccountGroupService) Update(ctx context.Context, id string, body SaleAccountGroupUpdateParams, opts ...option.RequestOption) (res *AccountGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of account groups.
//
// This endpoint requires the permission: `customer_groups:read`.
func (r *SaleAccountGroupService) List(ctx context.Context, query SaleAccountGroupListParams, opts ...option.RequestOption) (res *ListAccountGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/account-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an account group.
//
// Deletion fails with a validation error while the account group is still in use —
// for example by customer records, product line access, volume discounts, pricing
// assignments, or an active registration flow.
//
// This endpoint requires the permission: `customer_groups:delete`.
func (r *SaleAccountGroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleAccountGroupDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A named grouping of customer accounts, used for pricing rules or to categorize
// accounts.
type AccountGroup struct {
	// Account group ID.
	ID string `json:"id" api:"required"`
	// How sales commission applies to accounts in this group.
	//
	//   - `commission_applied`: sales commission is calculated on orders from accounts
	//     in this group.
	//   - `commission_exempt`: orders from accounts in this group are exempt from
	//     commission.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy AccountGroupCommissionPolicy `json:"commission_policy" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Free-form description of the account group.
	Description string `json:"description" api:"required"`
	// How freight charges apply to orders from accounts in this group.
	//
	//   - `free_freight`: customers within this group will not have to pay for freight.
	//   - `billed_freight`: freight will be applied to any order within this account
	//     group, unless overridden elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy AccountGroupFreightPolicy `json:"freight_policy" api:"required"`
	// Display name of the account group.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_group".
	Object AccountGroupObject `json:"object" api:"required"`
	// How this account group is used.
	//
	//   - `pricing_group`: used for pricing rules, such as a "Preferred" group that
	//     receives a special discount.
	//   - `type_group`: used to categorize accounts, such as "Consumers" or
	//     "Distributors".
	//
	// Any of "pricing_group", "type_group".
	Type AccountGroupType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CommissionPolicy respjson.Field
		CreatedAt        respjson.Field
		Description      respjson.Field
		FreightPolicy    respjson.Field
		Name             respjson.Field
		Object           respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGroup) RawJSON() string { return r.JSON.raw }
func (r *AccountGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How sales commission applies to accounts in this group.
//
//   - `commission_applied`: sales commission is calculated on orders from accounts
//     in this group.
//   - `commission_exempt`: orders from accounts in this group are exempt from
//     commission.
type AccountGroupCommissionPolicy string

const (
	AccountGroupCommissionPolicyCommissionApplied AccountGroupCommissionPolicy = "commission_applied"
	AccountGroupCommissionPolicyCommissionExempt  AccountGroupCommissionPolicy = "commission_exempt"
)

// How freight charges apply to orders from accounts in this group.
//
//   - `free_freight`: customers within this group will not have to pay for freight.
//   - `billed_freight`: freight will be applied to any order within this account
//     group, unless overridden elsewhere.
type AccountGroupFreightPolicy string

const (
	AccountGroupFreightPolicyFreeFreight   AccountGroupFreightPolicy = "free_freight"
	AccountGroupFreightPolicyBilledFreight AccountGroupFreightPolicy = "billed_freight"
)

// Resource type identifier.
type AccountGroupObject string

const (
	AccountGroupObjectAccountGroup AccountGroupObject = "account_group"
)

// How this account group is used.
//
//   - `pricing_group`: used for pricing rules, such as a "Preferred" group that
//     receives a special discount.
//   - `type_group`: used to categorize accounts, such as "Consumers" or
//     "Distributors".
type AccountGroupType string

const (
	AccountGroupTypePricingGroup AccountGroupType = "pricing_group"
	AccountGroupTypeTypeGroup    AccountGroupType = "type_group"
)

// Request to create an account group.
//
// The properties Name, Type are required.
type CreateAccountGroupRequestParam struct {
	// Display name of the account group.
	//
	// Must be unique within your account; maximum 255 characters.
	Name string `json:"name" api:"required"`
	// How this account group will be used.
	//
	//   - `pricing_group`: used for pricing rules, such as a "Preferred" group that
	//     receives a special discount.
	//   - `type_group`: used to categorize accounts, such as "Consumers" or
	//     "Distributors".
	//
	// The type cannot be changed after creation.
	//
	// Any of "pricing_group", "type_group".
	Type CreateAccountGroupRequestType `json:"type,omitzero" api:"required"`
	// Free-form description of the account group.
	Description param.Opt[string] `json:"description,omitzero"`
	// How sales commission applies to accounts in this group.
	//
	//   - `commission_applied`: sales commission is calculated on orders from accounts
	//     in this group.
	//   - `commission_exempt`: orders from accounts in this group are exempt from
	//     commission.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CreateAccountGroupRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// How freight charges apply to orders from accounts in this group.
	//
	//   - `free_freight`: customers within this group will not have to pay for freight.
	//   - `billed_freight`: freight will be applied to any order within this account
	//     group, unless overridden elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy CreateAccountGroupRequestFreightPolicy `json:"freight_policy,omitzero"`
	paramObj
}

func (r CreateAccountGroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAccountGroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAccountGroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How this account group will be used.
//
//   - `pricing_group`: used for pricing rules, such as a "Preferred" group that
//     receives a special discount.
//   - `type_group`: used to categorize accounts, such as "Consumers" or
//     "Distributors".
//
// The type cannot be changed after creation.
type CreateAccountGroupRequestType string

const (
	CreateAccountGroupRequestTypePricingGroup CreateAccountGroupRequestType = "pricing_group"
	CreateAccountGroupRequestTypeTypeGroup    CreateAccountGroupRequestType = "type_group"
)

// How sales commission applies to accounts in this group.
//
//   - `commission_applied`: sales commission is calculated on orders from accounts
//     in this group.
//   - `commission_exempt`: orders from accounts in this group are exempt from
//     commission.
type CreateAccountGroupRequestCommissionPolicy string

const (
	CreateAccountGroupRequestCommissionPolicyCommissionApplied CreateAccountGroupRequestCommissionPolicy = "commission_applied"
	CreateAccountGroupRequestCommissionPolicyCommissionExempt  CreateAccountGroupRequestCommissionPolicy = "commission_exempt"
)

// How freight charges apply to orders from accounts in this group.
//
//   - `free_freight`: customers within this group will not have to pay for freight.
//   - `billed_freight`: freight will be applied to any order within this account
//     group, unless overridden elsewhere.
type CreateAccountGroupRequestFreightPolicy string

const (
	CreateAccountGroupRequestFreightPolicyFreeFreight   CreateAccountGroupRequestFreightPolicy = "free_freight"
	CreateAccountGroupRequestFreightPolicyBilledFreight CreateAccountGroupRequestFreightPolicy = "billed_freight"
)

// List represents a paginated list of resources.
type ListAccountGroup struct {
	// Resources in this page.
	Data []AccountGroup `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAccountGroupObject `json:"object" api:"required"`
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
func (r ListAccountGroup) RawJSON() string { return r.JSON.raw }
func (r *ListAccountGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAccountGroupObject string

const (
	ListAccountGroupObjectList ListAccountGroupObject = "list"
)

// Request to partially update an account group.
type UpdateAccountGroupRequestParam struct {
	// Free-form description of the account group.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of the account group.
	//
	// Must be unique within your account; maximum 255 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// How sales commission applies to accounts in this group.
	//
	//   - `commission_applied`: sales commission is calculated on orders from accounts
	//     in this group.
	//   - `commission_exempt`: orders from accounts in this group are exempt from
	//     commission.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy UpdateAccountGroupRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// How freight charges apply to orders from accounts in this group.
	//
	//   - `free_freight`: customers within this group will not have to pay for freight.
	//   - `billed_freight`: freight will be applied to any order within this account
	//     group, unless overridden elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy UpdateAccountGroupRequestFreightPolicy `json:"freight_policy,omitzero"`
	paramObj
}

func (r UpdateAccountGroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAccountGroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAccountGroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How sales commission applies to accounts in this group.
//
//   - `commission_applied`: sales commission is calculated on orders from accounts
//     in this group.
//   - `commission_exempt`: orders from accounts in this group are exempt from
//     commission.
type UpdateAccountGroupRequestCommissionPolicy string

const (
	UpdateAccountGroupRequestCommissionPolicyCommissionApplied UpdateAccountGroupRequestCommissionPolicy = "commission_applied"
	UpdateAccountGroupRequestCommissionPolicyCommissionExempt  UpdateAccountGroupRequestCommissionPolicy = "commission_exempt"
)

// How freight charges apply to orders from accounts in this group.
//
//   - `free_freight`: customers within this group will not have to pay for freight.
//   - `billed_freight`: freight will be applied to any order within this account
//     group, unless overridden elsewhere.
type UpdateAccountGroupRequestFreightPolicy string

const (
	UpdateAccountGroupRequestFreightPolicyFreeFreight   UpdateAccountGroupRequestFreightPolicy = "free_freight"
	UpdateAccountGroupRequestFreightPolicyBilledFreight UpdateAccountGroupRequestFreightPolicy = "billed_freight"
)

type SaleAccountGroupDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleAccountGroupDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleAccountGroupDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountGroupNewParams struct {
	// Request to create an account group.
	CreateAccountGroupRequest CreateAccountGroupRequestParam
	paramObj
}

func (r SaleAccountGroupNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAccountGroupRequest)
}
func (r *SaleAccountGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountGroupUpdateParams struct {
	// Request to partially update an account group.
	UpdateAccountGroupRequest UpdateAccountGroupRequestParam
	paramObj
}

func (r SaleAccountGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAccountGroupRequest)
}
func (r *SaleAccountGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountGroupListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filters results to account groups of the given type.
	//
	// Any of "pricing_group", "type_group".
	Type SaleAccountGroupListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleAccountGroupListParams]'s query parameters as
// `url.Values`.
func (r SaleAccountGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filters results to account groups of the given type.
type SaleAccountGroupListParamsType string

const (
	SaleAccountGroupListParamsTypePricingGroup SaleAccountGroupListParamsType = "pricing_group"
	SaleAccountGroupListParamsTypeTypeGroup    SaleAccountGroupListParamsType = "type_group"
)
