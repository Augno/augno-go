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

// List and manage units.
//
// CatalogUnitService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogUnitService] method instead.
type CatalogUnitService struct {
	options []option.RequestOption
}

// NewCatalogUnitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogUnitService(opts ...option.RequestOption) (r CatalogUnitService) {
	r = CatalogUnitService{}
	r.options = opts
	return
}

// Creates an account-owned unit.
func (r *CatalogUnitService) New(ctx context.Context, params CatalogUnitNewParams, opts ...option.RequestOption) (res *Unit, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/units"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a unit by ID, including both account-owned and global system units.
func (r *CatalogUnitService) Get(ctx context.Context, id string, query CatalogUnitGetParams, opts ...option.RequestOption) (res *Unit, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/units/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an account-owned unit; system units cannot be updated.
func (r *CatalogUnitService) Update(ctx context.Context, id string, params CatalogUnitUpdateParams, opts ...option.RequestOption) (res *Unit, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/units/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of units for the current account, including both
// account-owned and global system units.
func (r *CatalogUnitService) List(ctx context.Context, query CatalogUnitListParams, opts ...option.RequestOption) (res *ListUnit, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/units"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an account-owned unit. Associated unit group memberships are also
// removed, and system units cannot be deleted.
func (r *CatalogUnitService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CatalogUnitDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/units/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a unit.
//
// The properties Abbreviation, Name, OffsetDenominator, OffsetNumerator,
// RatioDenominator, RatioNumerator, Type are required.
type CreateUnitRequestParam struct {
	// Short abbreviation for the unit (e.g. "g").
	//
	// Must be unique within the account.
	Abbreviation string `json:"abbreviation" api:"required"`
	// Display name of the unit (e.g. "Gram").
	//
	// Must be unique within the account.
	Name string `json:"name" api:"required"`
	// Conversion offset denominator, as a decimal string. Must not be zero.
	OffsetDenominator string `json:"offset_denominator" api:"required" format:"decimal"`
	// Conversion offset numerator, as a decimal string.
	OffsetNumerator string `json:"offset_numerator" api:"required" format:"decimal"`
	// Conversion ratio denominator relative to the base unit, as a decimal string.
	// Must not be zero.
	RatioDenominator string `json:"ratio_denominator" api:"required" format:"decimal"`
	// Conversion ratio numerator relative to the base unit, as a decimal string.
	RatioNumerator string `json:"ratio_numerator" api:"required" format:"decimal"`
	// Unit dimension (e.g. `mass`, `volume`, `currency`).
	//
	// Units can only be converted to other units of the same dimension.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CreateUnitRequestType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r CreateUnitRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateUnitRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateUnitRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unit dimension (e.g. `mass`, `volume`, `currency`).
//
// Units can only be converted to other units of the same dimension.
type CreateUnitRequestType string

const (
	CreateUnitRequestTypeCurrency    CreateUnitRequestType = "currency"
	CreateUnitRequestTypeQuantity    CreateUnitRequestType = "quantity"
	CreateUnitRequestTypeTime        CreateUnitRequestType = "time"
	CreateUnitRequestTypeMass        CreateUnitRequestType = "mass"
	CreateUnitRequestTypeVolume      CreateUnitRequestType = "volume"
	CreateUnitRequestTypeLength      CreateUnitRequestType = "length"
	CreateUnitRequestTypeTemperature CreateUnitRequestType = "temperature"
	CreateUnitRequestTypeArea        CreateUnitRequestType = "area"
)

// List represents a paginated list of resources.
type ListUnit struct {
	// Resources in this page.
	Data []Unit `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListUnitObject `json:"object" api:"required"`
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
func (r ListUnit) RawJSON() string { return r.JSON.raw }
func (r *ListUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListUnitObject string

const (
	ListUnitObjectList ListUnitObject = "list"
)

// Unit of measurement used for conversions and product quantities.
type Unit struct {
	// Unit ID.
	ID string `json:"id" api:"required"`
	// Short abbreviation for the unit (e.g. "g", "kg").
	Abbreviation string `json:"abbreviation" api:"required"`
	// When this unit was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this is the base unit for its dimension.
	//
	// Conversion ratios are relative to this unit. Base units are platform-defined;
	// account-created units always have this set to `false`.
	IsBaseUnit bool `json:"is_base_unit" api:"required"`
	// Display name of the unit (e.g. "Gram", "Kilogram").
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "unit".
	Object UnitObject `json:"object" api:"required"`
	// Conversion offset denominator.
	//
	// Typically 1. Cannot be zero.
	OffsetDenominator string `json:"offset_denominator" api:"required" format:"decimal"`
	// Conversion offset numerator, used for temperature-like conversions.
	//
	// Zero for most unit types.
	OffsetNumerator string `json:"offset_numerator" api:"required" format:"decimal"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Conversion ratio denominator relative to the base unit in the same dimension.
	//
	// Cannot be zero.
	RatioDenominator string `json:"ratio_denominator" api:"required" format:"decimal"`
	// Conversion ratio numerator relative to the base unit in the same dimension.
	RatioNumerator string `json:"ratio_numerator" api:"required" format:"decimal"`
	// Physical dimension the unit measures, such as mass, volume, or currency.
	//
	// A unit can only be converted to another unit of the same dimension. The
	// `quantity` dimension is for discrete countable items rather than a physical
	// measure.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type UnitType `json:"type" api:"required"`
	// When this unit was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Abbreviation      respjson.Field
		CreatedAt         respjson.Field
		IsBaseUnit        respjson.Field
		Name              respjson.Field
		Object            respjson.Field
		OffsetDenominator respjson.Field
		OffsetNumerator   respjson.Field
		Owner             respjson.Field
		RatioDenominator  respjson.Field
		RatioNumerator    respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Unit) RawJSON() string { return r.JSON.raw }
func (r *Unit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type UnitObject string

const (
	UnitObjectUnit UnitObject = "unit"
)

// Physical dimension the unit measures, such as mass, volume, or currency.
//
// A unit can only be converted to another unit of the same dimension. The
// `quantity` dimension is for discrete countable items rather than a physical
// measure.
type UnitType string

const (
	UnitTypeCurrency    UnitType = "currency"
	UnitTypeQuantity    UnitType = "quantity"
	UnitTypeTime        UnitType = "time"
	UnitTypeMass        UnitType = "mass"
	UnitTypeVolume      UnitType = "volume"
	UnitTypeLength      UnitType = "length"
	UnitTypeTemperature UnitType = "temperature"
	UnitTypeArea        UnitType = "area"
)

// Request to partially update a unit.
type UpdateUnitRequestParam struct {
	// Short abbreviation for the unit.
	//
	// Must be unique within the account.
	Abbreviation param.Opt[string] `json:"abbreviation,omitzero"`
	// Display name of the unit.
	//
	// Must be unique within the account.
	Name param.Opt[string] `json:"name,omitzero"`
	// Conversion offset denominator, as a decimal string. Must not be zero.
	OffsetDenominator param.Opt[string] `json:"offset_denominator,omitzero" format:"decimal"`
	// Conversion offset numerator, as a decimal string.
	OffsetNumerator param.Opt[string] `json:"offset_numerator,omitzero" format:"decimal"`
	// Conversion ratio denominator, as a decimal string. Must not be zero.
	RatioDenominator param.Opt[string] `json:"ratio_denominator,omitzero" format:"decimal"`
	// Conversion ratio numerator, as a decimal string.
	RatioNumerator param.Opt[string] `json:"ratio_numerator,omitzero" format:"decimal"`
	paramObj
}

func (r UpdateUnitRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateUnitRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateUnitRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogUnitDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogUnitDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogUnitDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogUnitNewParams struct {
	// Request to create a unit.
	CreateUnitRequest CreateUnitRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogUnitNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateUnitRequest)
}
func (r *CatalogUnitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitNewParams]'s query parameters as `url.Values`.
func (r CatalogUnitNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitGetParams]'s query parameters as `url.Values`.
func (r CatalogUnitGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a unit.
	UpdateUnitRequest UpdateUnitRequestParam
	paramObj
}

func (r CatalogUnitUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateUnitRequest)
}
func (r *CatalogUnitUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitListParams struct {
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
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by unit dimension (e.g. `mass`).
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CatalogUnitListParamsType `query:"type,omitzero" json:"-"`
	// Return only units that belong to at least one of the given unit groups.
	UnitGroupIDs []string `query:"unit_group_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitListParams]'s query parameters as `url.Values`.
func (r CatalogUnitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by unit dimension (e.g. `mass`).
type CatalogUnitListParamsType string

const (
	CatalogUnitListParamsTypeCurrency    CatalogUnitListParamsType = "currency"
	CatalogUnitListParamsTypeQuantity    CatalogUnitListParamsType = "quantity"
	CatalogUnitListParamsTypeTime        CatalogUnitListParamsType = "time"
	CatalogUnitListParamsTypeMass        CatalogUnitListParamsType = "mass"
	CatalogUnitListParamsTypeVolume      CatalogUnitListParamsType = "volume"
	CatalogUnitListParamsTypeLength      CatalogUnitListParamsType = "length"
	CatalogUnitListParamsTypeTemperature CatalogUnitListParamsType = "temperature"
	CatalogUnitListParamsTypeArea        CatalogUnitListParamsType = "area"
)
