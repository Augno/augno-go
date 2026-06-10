// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
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

// Create, view, update, and delete transactions.
//
// FinanceService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinanceService] method instead.
type FinanceService struct {
	options []option.RequestOption
	// List and manage payment terms.
	PaymentTerms FinancePaymentTermService
}

// NewFinanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFinanceService(opts ...option.RequestOption) (r FinanceService) {
	r = FinanceService{}
	r.options = opts
	r.PaymentTerms = NewFinancePaymentTermService(opts...)
	return
}

// Returns a paginated list of adjustment types.
func (r *FinanceService) GetAdjustmentTypes(ctx context.Context, query FinanceGetAdjustmentTypesParams, opts ...option.RequestOption) (res *ListAdjustmentType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/finance/adjustment-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of transaction methods.
func (r *FinanceService) GetTransactionMethods(ctx context.Context, query FinanceGetTransactionMethodsParams, opts ...option.RequestOption) (res *ListTransactionMethod, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/finance/transaction-methods"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of transaction types.
func (r *FinanceService) GetTransactionTypes(ctx context.Context, query FinanceGetTransactionTypesParams, opts ...option.RequestOption) (res *ListTransactionType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/finance/transaction-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Adjustment type resource.
type AdjustmentType struct {
	// Adjustment ID.
	ID string `json:"id" api:"required"`
	// Machine-readable code.
	//
	// Any of "discount", "shipping_discrepancy", "short_payment", "write_off", "fee",
	// "refund".
	Code AdjustmentTypeCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "adjustment_type".
	Object AdjustmentTypeObject `json:"object" api:"required"`
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
func (r AdjustmentType) RawJSON() string { return r.JSON.raw }
func (r *AdjustmentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable code.
type AdjustmentTypeCode string

const (
	AdjustmentTypeCodeDiscount            AdjustmentTypeCode = "discount"
	AdjustmentTypeCodeShippingDiscrepancy AdjustmentTypeCode = "shipping_discrepancy"
	AdjustmentTypeCodeShortPayment        AdjustmentTypeCode = "short_payment"
	AdjustmentTypeCodeWriteOff            AdjustmentTypeCode = "write_off"
	AdjustmentTypeCodeFee                 AdjustmentTypeCode = "fee"
	AdjustmentTypeCodeRefund              AdjustmentTypeCode = "refund"
)

// Resource type identifier.
type AdjustmentTypeObject string

const (
	AdjustmentTypeObjectAdjustmentType AdjustmentTypeObject = "adjustment_type"
)

// List represents a paginated list of resources.
type ListAdjustmentType struct {
	// Resources in this page.
	Data []AdjustmentType `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAdjustmentTypeObject `json:"object" api:"required"`
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
func (r ListAdjustmentType) RawJSON() string { return r.JSON.raw }
func (r *ListAdjustmentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAdjustmentTypeObject string

const (
	ListAdjustmentTypeObjectList ListAdjustmentTypeObject = "list"
)

// List represents a paginated list of resources.
type ListTransactionMethod struct {
	// Resources in this page.
	Data []TransactionMethod `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListTransactionMethodObject `json:"object" api:"required"`
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
func (r ListTransactionMethod) RawJSON() string { return r.JSON.raw }
func (r *ListTransactionMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListTransactionMethodObject string

const (
	ListTransactionMethodObjectList ListTransactionMethodObject = "list"
)

// List represents a paginated list of resources.
type ListTransactionType struct {
	// Resources in this page.
	Data []TransactionType `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListTransactionTypeObject `json:"object" api:"required"`
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
func (r ListTransactionType) RawJSON() string { return r.JSON.raw }
func (r *ListTransactionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListTransactionTypeObject string

const (
	ListTransactionTypeObjectList ListTransactionTypeObject = "list"
)

// Transaction method resource.
type TransactionMethod struct {
	// Transaction method ID.
	ID string `json:"id" api:"required"`
	// Machine-readable code.
	//
	// Any of "cash", "check", "credit_card", "gift_card", "ach".
	Code TransactionMethodCode `json:"code" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "transaction_method".
	Object TransactionMethodObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionMethod) RawJSON() string { return r.JSON.raw }
func (r *TransactionMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable code.
type TransactionMethodCode string

const (
	TransactionMethodCodeCash       TransactionMethodCode = "cash"
	TransactionMethodCodeCheck      TransactionMethodCode = "check"
	TransactionMethodCodeCreditCard TransactionMethodCode = "credit_card"
	TransactionMethodCodeGiftCard   TransactionMethodCode = "gift_card"
	TransactionMethodCodeACH        TransactionMethodCode = "ach"
)

// Resource type identifier.
type TransactionMethodObject string

const (
	TransactionMethodObjectTransactionMethod TransactionMethodObject = "transaction_method"
)

// Transaction type resource.
type TransactionType struct {
	// Transaction ID.
	ID string `json:"id" api:"required"`
	// Machine-readable code.
	//
	// Any of "payment", "credit_memo", "adjustment", "rebate".
	Code TransactionTypeCode `json:"code" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "transaction_type".
	Object TransactionTypeObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionType) RawJSON() string { return r.JSON.raw }
func (r *TransactionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable code.
type TransactionTypeCode string

const (
	TransactionTypeCodePayment    TransactionTypeCode = "payment"
	TransactionTypeCodeCreditMemo TransactionTypeCode = "credit_memo"
	TransactionTypeCodeAdjustment TransactionTypeCode = "adjustment"
	TransactionTypeCodeRebate     TransactionTypeCode = "rebate"
)

// Resource type identifier.
type TransactionTypeObject string

const (
	TransactionTypeObjectTransactionType TransactionTypeObject = "transaction_type"
)

type FinanceGetAdjustmentTypesParams struct {
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

// URLQuery serializes [FinanceGetAdjustmentTypesParams]'s query parameters as
// `url.Values`.
func (r FinanceGetAdjustmentTypesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FinanceGetTransactionMethodsParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FinanceGetTransactionMethodsParams]'s query parameters as
// `url.Values`.
func (r FinanceGetTransactionMethodsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FinanceGetTransactionTypesParams struct {
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FinanceGetTransactionTypesParams]'s query parameters as
// `url.Values`.
func (r FinanceGetTransactionTypesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
