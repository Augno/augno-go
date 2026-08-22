// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// The planning assumptions production schedules are solved against, and the
// per-resource overrides that mark which machines constrain the plan.
//
// OperationProductionScheduleSettingItemService contains methods and other
// services that help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationProductionScheduleSettingItemService] method instead.
type OperationProductionScheduleSettingItemService struct {
	options []option.RequestOption
}

// NewOperationProductionScheduleSettingItemService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOperationProductionScheduleSettingItemService(opts ...option.RequestOption) (r OperationProductionScheduleSettingItemService) {
	r = OperationProductionScheduleSettingItemService{}
	r.options = opts
	return
}

// Returns the planning overrides for one item.
//
// Fails with a not-found error when the item has none, rather than returning an
// empty set of overrides: an item with no overrides is planned on the account
// defaults and its product line's conventions, and reporting that as a resource
// would suggest there is something here to edit.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleSettingItemService) Get(ctx context.Context, itemID string, opts ...option.RequestOption) (res *ProductionScheduleItemSetting, err error) {
	opts = slices.Concat(r.options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedule-settings/items/%s", url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Writes the planning overrides for one item.
//
// An item has at most one set of overrides, so this replaces the existing entry
// rather than adding a second, and the entry keeps the ID it already had.
//
// The fulfillment policy is the most consequential of these. A `make_to_order`
// item contributes no forecast demand and holds no safety stock, so it is built
// only against orders already on the book — which is what stops a slow mover
// accumulating inventory nobody asked for. It also propagates: an intermediate
// item is planned to order only when every finished good it becomes is, so one
// stocked sibling keeps the whole family buffered.
//
// Overrides are read when a plan is generated, so a change takes effect on the
// next generated version and leaves existing ones untouched.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleSettingItemService) Update(ctx context.Context, itemID string, body OperationProductionScheduleSettingItemUpdateParams, opts ...option.RequestOption) (res *ProductionScheduleItemSetting, err error) {
	opts = slices.Concat(r.options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedule-settings/items/%s", url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns every per-item planning override in the account.
//
// Only items that have been given an override appear here. An item with none is
// planned on the account defaults and its product line's conventions, which is the
// normal case — this is the list of exceptions, not a list of every item.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleSettingItemService) List(ctx context.Context, opts ...option.RequestOption) (res *ListProductionScheduleItemSetting, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedule-settings/items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes one item's planning overrides, returning it to the account defaults and
// its product line's conventions.
//
// Fails with a not-found error when the item has no overrides, rather than
// reporting success: a mistyped item ID would otherwise read as a change that
// never happened.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleSettingItemService) Delete(ctx context.Context, itemID string, opts ...option.RequestOption) (res *OperationProductionScheduleSettingItemDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedule-settings/items/%s", url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionScheduleItemSetting struct {
	// Resources in this page.
	Data []ProductionScheduleItemSetting `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleItemSettingObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
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
func (r ListProductionScheduleItemSetting) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleItemSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleItemSettingObject string

const (
	ListProductionScheduleItemSettingObjectList ListProductionScheduleItemSettingObject = "list"
)

// Planning overrides for one item, on top of the account-wide assumptions.
type ProductionScheduleItemSetting struct {
	// Item setting ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// How this item is produced.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer.
	//
	// Null inherits from the item's product line, then from the account default.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy ProductionScheduleItemSettingFulfillmentPolicy `json:"fulfillment_policy" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Units in one production lot for this item, overriding the lot its product line
	// would supply.
	LotMultipleUnits float64 `json:"lot_multiple_units" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_item_setting".
	Object ProductionScheduleItemSettingObject `json:"object" api:"required"`
	// Whether this item takes part in planning.
	//
	// An excluded item is left out of the plan entirely: no campaigns, no policy, no
	// capacity.
	//
	// Any of "included", "excluded".
	ParticipationStatus ProductionScheduleItemSettingParticipationStatus `json:"participation_status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		FulfillmentPolicy   respjson.Field
		Item                respjson.Field
		LotMultipleUnits    respjson.Field
		Object              respjson.Field
		ParticipationStatus respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleItemSetting) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleItemSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How this item is produced.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer.
//
// Null inherits from the item's product line, then from the account default.
type ProductionScheduleItemSettingFulfillmentPolicy string

const (
	ProductionScheduleItemSettingFulfillmentPolicyMakeToStock ProductionScheduleItemSettingFulfillmentPolicy = "make_to_stock"
	ProductionScheduleItemSettingFulfillmentPolicyMakeToOrder ProductionScheduleItemSettingFulfillmentPolicy = "make_to_order"
)

// Resource type identifier.
type ProductionScheduleItemSettingObject string

const (
	ProductionScheduleItemSettingObjectProductionScheduleItemSetting ProductionScheduleItemSettingObject = "production_schedule_item_setting"
)

// Whether this item takes part in planning.
//
// An excluded item is left out of the plan entirely: no campaigns, no policy, no
// capacity.
type ProductionScheduleItemSettingParticipationStatus string

const (
	ProductionScheduleItemSettingParticipationStatusIncluded ProductionScheduleItemSettingParticipationStatus = "included"
	ProductionScheduleItemSettingParticipationStatusExcluded ProductionScheduleItemSettingParticipationStatus = "excluded"
)

// Request to write one item's planning overrides.
//
// The property ParticipationStatus is required.
type UpsertItemSettingRequestParam struct {
	// Whether this item takes part in planning.
	//
	// An excluded item is left out of the plan entirely: no campaigns, no policy, no
	// capacity.
	//
	// Any of "included", "excluded".
	ParticipationStatus UpsertItemSettingRequestParticipationStatus `json:"participation_status,omitzero" api:"required"`
	// Units in one production lot for this item, overriding the lot its product line
	// would supply.
	LotMultipleUnits param.Opt[float64] `json:"lot_multiple_units,omitzero"`
	// How this item is produced.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer.
	//
	// Clearing it returns the item to its product line's policy, then to the account
	// default.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy UpsertItemSettingRequestFulfillmentPolicy `json:"fulfillment_policy,omitzero"`
	paramObj
}

func (r UpsertItemSettingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertItemSettingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertItemSettingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this item takes part in planning.
//
// An excluded item is left out of the plan entirely: no campaigns, no policy, no
// capacity.
type UpsertItemSettingRequestParticipationStatus string

const (
	UpsertItemSettingRequestParticipationStatusIncluded UpsertItemSettingRequestParticipationStatus = "included"
	UpsertItemSettingRequestParticipationStatusExcluded UpsertItemSettingRequestParticipationStatus = "excluded"
)

// How this item is produced.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer.
//
// Clearing it returns the item to its product line's policy, then to the account
// default.
type UpsertItemSettingRequestFulfillmentPolicy string

const (
	UpsertItemSettingRequestFulfillmentPolicyMakeToStock UpsertItemSettingRequestFulfillmentPolicy = "make_to_stock"
	UpsertItemSettingRequestFulfillmentPolicyMakeToOrder UpsertItemSettingRequestFulfillmentPolicy = "make_to_order"
)

type OperationProductionScheduleSettingItemDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationProductionScheduleSettingItemDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationProductionScheduleSettingItemDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleSettingItemUpdateParams struct {
	// Request to write one item's planning overrides.
	UpsertItemSettingRequest UpsertItemSettingRequestParam
	paramObj
}

func (r OperationProductionScheduleSettingItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpsertItemSettingRequest)
}
func (r *OperationProductionScheduleSettingItemUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
