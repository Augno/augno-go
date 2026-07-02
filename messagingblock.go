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

// Block and unblock users from direct messaging.
//
// MessagingBlockService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingBlockService] method instead.
type MessagingBlockService struct {
	options []option.RequestOption
}

// NewMessagingBlockService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagingBlockService(opts ...option.RequestOption) (r MessagingBlockService) {
	r = MessagingBlockService{}
	r.options = opts
	return
}

// Blocks an account user (prevents DMs in both directions).
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingBlockService) New(ctx context.Context, params MessagingBlockNewParams, opts ...option.RequestOption) (res *MessagingBlock, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/blocks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Lists the caller's messaging blocks.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingBlockService) List(ctx context.Context, query MessagingBlockListParams, opts ...option.RequestOption) (res *ListMessagingBlock, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/blocks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes a block.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingBlockService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingBlockDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/blocks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A user's membership in an account, carrying the account-specific status, role,
// and department.
//
// Profile fields (name, email, username, image URL) live on the expandable `user`
// sub-resource, which is shared across every account the user belongs to.
type AccountUser struct {
	// Account user ID.
	ID string `json:"id" api:"required"`
	// When the account user was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department Department `json:"department" api:"required"`
	// When the user last accessed this account.
	LastUsedAt time.Time `json:"last_used_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "account_user".
	Object AccountUserObject `json:"object" api:"required"`
	// A named set of permissions that can be assigned to users to control what they
	// can access.
	Role Role `json:"role" api:"required"`
	// Account user status.
	//
	// - `active`: the user can access the account.
	// - `disabled`: the user is locked out of the account.
	// - `removed`: the user has been removed (soft-deleted) from the account.
	//
	// Any of "active", "disabled", "removed".
	Status AccountUserStatus `json:"status" api:"required"`
	// When the account user was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A user's global profile, shared across every account they belong to.
	//
	// Account-specific settings (status, role, department) live on the account user
	// resource that links the user to each account.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Department  respjson.Field
		LastUsedAt  respjson.Field
		Object      respjson.Field
		Role        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountUser) RawJSON() string { return r.JSON.raw }
func (r *AccountUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountUserObject string

const (
	AccountUserObjectAccountUser AccountUserObject = "account_user"
)

// Account user status.
//
// - `active`: the user can access the account.
// - `disabled`: the user is locked out of the account.
// - `removed`: the user has been removed (soft-deleted) from the account.
type AccountUserStatus string

const (
	AccountUserStatusActive   AccountUserStatus = "active"
	AccountUserStatusDisabled AccountUserStatus = "disabled"
	AccountUserStatusRemoved  AccountUserStatus = "removed"
)

// Request to block another account user from messaging the caller.
//
// The property BlockedAccountUserID is required.
type BlockRequestParam struct {
	// The account user to block.
	BlockedAccountUserID string `json:"blocked_account_user_id" api:"required"`
	paramObj
}

func (r BlockRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BlockRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlockRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Material consumed by a production step.
//
// Each consumption records one input item and how much of it the step uses.
// Consumptions also determine the production flow: when another step produces the
// consumed item, the two steps are linked upstream/downstream automatically.
type Consumption struct {
	// Consumption ID.
	ID string `json:"id" api:"required"`
	// Item is an inventory item (product, material, or part).
	ConsumedItem Item `json:"consumed_item" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Instructions for how this material is consumed.
	Instructions string `json:"instructions" api:"required"`
	// Resource type identifier.
	//
	// Any of "consumption".
	Object ConsumptionObject `json:"object" api:"required"`
	// Value with an associated unit.
	Quantity Quantity `json:"quantity" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Value with an associated unit.
	WasteQuantity Quantity `json:"waste_quantity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ConsumedItem  respjson.Field
		CreatedAt     respjson.Field
		Instructions  respjson.Field
		Object        respjson.Field
		Quantity      respjson.Field
		UpdatedAt     respjson.Field
		WasteQuantity respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Consumption) RawJSON() string { return r.JSON.raw }
func (r *Consumption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ConsumptionObject string

const (
	ConsumptionObjectConsumption ConsumptionObject = "consumption"
)

// A functional area of a production operation, such as fabrication or packaging,
// that groups scanning stations and machines.
type Department struct {
	// Department ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A physical storage location, such as a warehouse, aisle, or bin, arranged in a
	// parent-child hierarchy.
	Location Location `json:"location" api:"required"`
	// List represents a paginated list of resources.
	Machines *ListMachine `json:"machines" api:"required"`
	// Display name of the department.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Free-form notes about the department.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "department".
	Object DepartmentObject `json:"object" api:"required"`
	// List represents a paginated list of resources.
	ScanningStations *ListScanningStation `json:"scanning_stations" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Location         respjson.Field
		Machines         respjson.Field
		Name             respjson.Field
		Notes            respjson.Field
		Object           respjson.Field
		ScanningStations respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Department) RawJSON() string { return r.JSON.raw }
func (r *Department) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type DepartmentObject string

const (
	DepartmentObjectDepartment DepartmentObject = "department"
)

// List represents a paginated list of resources.
type ListConsumption struct {
	// Resources in this page.
	Data []Consumption `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListConsumptionObject `json:"object" api:"required"`
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
func (r ListConsumption) RawJSON() string { return r.JSON.raw }
func (r *ListConsumption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListConsumptionObject string

const (
	ListConsumptionObjectList ListConsumptionObject = "list"
)

// List represents a paginated list of resources.
type ListLocation struct {
	// Resources in this page.
	Data []Location `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListLocationObject `json:"object" api:"required"`
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
func (r ListLocation) RawJSON() string { return r.JSON.raw }
func (r *ListLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListLocationObject string

const (
	ListLocationObjectList ListLocationObject = "list"
)

// List represents a paginated list of resources.
type ListMachine struct {
	// Resources in this page.
	Data []Machine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMachineObject `json:"object" api:"required"`
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
func (r ListMachine) RawJSON() string { return r.JSON.raw }
func (r *ListMachine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMachineObject string

const (
	ListMachineObjectList ListMachineObject = "list"
)

// List represents a paginated list of resources.
type ListMessagingBlock struct {
	// Resources in this page.
	Data []MessagingBlock `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessagingBlockObject `json:"object" api:"required"`
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
func (r ListMessagingBlock) RawJSON() string { return r.JSON.raw }
func (r *ListMessagingBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessagingBlockObject string

const (
	ListMessagingBlockObjectList ListMessagingBlockObject = "list"
)

// List represents a paginated list of resources.
type ListProductionStep struct {
	// Resources in this page.
	Data []ProductionStep `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionStepObject `json:"object" api:"required"`
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
func (r ListProductionStep) RawJSON() string { return r.JSON.raw }
func (r *ListProductionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionStepObject string

const (
	ListProductionStepObjectList ListProductionStepObject = "list"
)

// List represents a paginated list of resources.
type ListScanningStation struct {
	// Resources in this page.
	Data []ScanningStation `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScanningStationObject `json:"object" api:"required"`
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
func (r ListScanningStation) RawJSON() string { return r.JSON.raw }
func (r *ListScanningStation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScanningStationObject string

const (
	ListScanningStationObjectList ListScanningStationObject = "list"
)

// A physical storage location, such as a warehouse, aisle, or bin, arranged in a
// parent-child hierarchy.
type Location struct {
	// Location ID.
	ID string `json:"id" api:"required"`
	// List represents a paginated list of resources.
	Children *ListLocation `json:"children" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the location.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "location".
	Object LocationObject `json:"object" api:"required"`
	// A physical storage location, such as a warehouse, aisle, or bin, arranged in a
	// parent-child hierarchy.
	Parent *Location `json:"parent" api:"required"`
	// Location type code, identifying this location's level in the storage hierarchy.
	//
	// - `building`: a building-level location.
	// - `section`: a section within a building.
	// - `aisle`: an aisle within a section.
	// - `rack`: a rack within an aisle.
	// - `shelf`: a shelf within a rack.
	// - `bin`: a bin within a shelf.
	//
	// Any of "building", "section", "aisle", "rack", "shelf", "bin".
	Type LocationTypeCode `json:"type" api:"required"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Parent      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type LocationObject string

const (
	LocationObjectLocation LocationObject = "location"
)

type LocationTypeCode string

const (
	LocationTypeCodeBuilding LocationTypeCode = "building"
	LocationTypeCodeSection  LocationTypeCode = "section"
	LocationTypeCodeAisle    LocationTypeCode = "aisle"
	LocationTypeCodeRack     LocationTypeCode = "rack"
	LocationTypeCodeShelf    LocationTypeCode = "shelf"
	LocationTypeCodeBin      LocationTypeCode = "bin"
)

// A piece of production equipment, such as a CNC router or press, assigned to a
// department.
type Machine struct {
	// Machine ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department *Department `json:"department" api:"required"`
	// Display name of the machine.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Free-form notes about the machine.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "machine".
	Object MachineObject `json:"object" api:"required"`
	// Serial number of the machine.
	SerialNumber string `json:"serial_number" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Department   respjson.Field
		Name         respjson.Field
		Notes        respjson.Field
		Object       respjson.Field
		SerialNumber respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Machine) RawJSON() string { return r.JSON.raw }
func (r *Machine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MachineObject string

const (
	MachineObjectMachine MachineObject = "machine"
)

// A 1:1 messaging block: the caller has blocked another account user from
// messaging them.
type MessagingBlock struct {
	// Block ID.
	ID string `json:"id" api:"required"`
	// A user's membership in an account, carrying the account-specific status, role,
	// and department.
	//
	// Profile fields (name, email, username, image URL) live on the expandable `user`
	// sub-resource, which is shared across every account the user belongs to.
	BlockedUser AccountUser `json:"blocked_user" api:"required"`
	// When the block was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "messaging_block".
	Object MessagingBlockObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BlockedUser respjson.Field
		CreatedAt   respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingBlock) RawJSON() string { return r.JSON.raw }
func (r *MessagingBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MessagingBlockObject string

const (
	MessagingBlockObjectMessagingBlock MessagingBlockObject = "messaging_block"
)

// The output of a production step: the item it produces and the quantity produced.
type ProductionOutput struct {
	// Production ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "production".
	Object ProductionOutputObject `json:"object" api:"required"`
	// Item is an inventory item (product, material, or part).
	ProducedItem Item `json:"produced_item" api:"required"`
	// Value with an associated unit.
	Quantity Quantity `json:"quantity" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Object       respjson.Field
		ProducedItem respjson.Field
		Quantity     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionOutput) RawJSON() string { return r.JSON.raw }
func (r *ProductionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionOutputObject string

const (
	ProductionOutputObjectProduction ProductionOutputObject = "production"
)

// A single stage of work in an item's production flow, with its output, material
// inputs, cost rates, and graph connections.
type ProductionStep struct {
	// Production step ID.
	ID string `json:"id" api:"required"`
	// Allowance correction factor applied to labor time in cost calculations, as a
	// decimal string.
	//
	// Effective labor time per unit is
	// `labor_time × (1 + leveling_factor) × (1 + allowances)`.
	Allowances string `json:"allowances" api:"required" format:"decimal"`
	// List represents a paginated list of resources.
	Consumptions ListConsumption `json:"consumptions" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department *Department `json:"department" api:"required"`
	// List represents a paginated list of resources.
	InSteps *ListProductionStep `json:"in_steps" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	LaborRate Rate `json:"labor_rate" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	LaborTime Rate `json:"labor_time" api:"required"`
	// Leveling correction factor applied to labor time in cost calculations, as a
	// decimal string.
	//
	// Effective labor time per unit is
	// `labor_time × (1 + leveling_factor) × (1 + allowances)`.
	LevelingFactor string `json:"leveling_factor" api:"required" format:"decimal"`
	// List represents a paginated list of resources.
	Machines *ListMachine `json:"machines" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Free-form notes about the step.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_step".
	Object ProductionStepObject `json:"object" api:"required"`
	// List represents a paginated list of resources.
	OutSteps *ListProductionStep `json:"out_steps" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	OverheadRate Rate `json:"overhead_rate" api:"required"`
	// The output of a production step: the item it produces and the quantity produced.
	Production ProductionOutput `json:"production" api:"required"`
	// A station on the production floor where operators scan batches to perform a
	// batch operation, such as initializing or moving a batch.
	ScanningStation *ScanningStation `json:"scanning_station" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Allowances      respjson.Field
		Consumptions    respjson.Field
		CreatedAt       respjson.Field
		Department      respjson.Field
		InSteps         respjson.Field
		LaborRate       respjson.Field
		LaborTime       respjson.Field
		LevelingFactor  respjson.Field
		Machines        respjson.Field
		Name            respjson.Field
		Notes           respjson.Field
		Object          respjson.Field
		OutSteps        respjson.Field
		OverheadRate    respjson.Field
		Production      respjson.Field
		ScanningStation respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionStep) RawJSON() string { return r.JSON.raw }
func (r *ProductionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionStepObject string

const (
	ProductionStepObjectProductionStep ProductionStepObject = "production_step"
)

// A station on the production floor where operators scan batches to perform a
// batch operation, such as initializing or moving a batch.
type ScanningStation struct {
	// Scanning station ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department *Department `json:"department" api:"required"`
	// Size of the labels printed at this station, given as width-by-height (for
	// example, `1x1`).
	//
	// Any of "1x1", "1x3", "1x4", "2x4".
	LabelSize ScanningStationLabelSize `json:"label_size" api:"required"`
	// Type of label printed at this station.
	//
	//   - `tag`: a label attached to the physical product.
	//   - `traveler`: a routing sheet that accompanies the batch through every
	//     production step.
	//
	// Any of "tag", "traveler".
	LabelType ScanningStationLabelType `json:"label_type" api:"required"`
	// Display name of the scanning station.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Free-form notes about the scanning station.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "scanning_station".
	Object ScanningStationObject `json:"object" api:"required"`
	// Whether operators must perform a material check at this station.
	//
	// - `none`: no additional operator check is required.
	// - `material_check`: a material check is expected before the operation.
	//
	// Any of "none", "material_check".
	OperatorRequirement ScanningStationOperatorRequirement `json:"operator_requirement" api:"required"`
	// List represents a paginated list of resources.
	ProductionSteps *ListProductionStep `json:"production_steps" api:"required"`
	// Scanning station type, determining which batch operation the station performs.
	//
	// - `init_batch`: initializes a new batch.
	// - `merge_batch`: merges multiple batches into one.
	// - `move_batch`: moves a batch to another location or step.
	// - `split_batch`: splits a batch into multiple batches.
	//
	// Any of "init_batch", "merge_batch", "move_batch", "split_batch".
	Type ScanningStationType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		Department          respjson.Field
		LabelSize           respjson.Field
		LabelType           respjson.Field
		Name                respjson.Field
		Notes               respjson.Field
		Object              respjson.Field
		OperatorRequirement respjson.Field
		ProductionSteps     respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScanningStation) RawJSON() string { return r.JSON.raw }
func (r *ScanningStation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Size of the labels printed at this station, given as width-by-height (for
// example, `1x1`).
type ScanningStationLabelSize string

const (
	ScanningStationLabelSize1x1 ScanningStationLabelSize = "1x1"
	ScanningStationLabelSize1x3 ScanningStationLabelSize = "1x3"
	ScanningStationLabelSize1x4 ScanningStationLabelSize = "1x4"
	ScanningStationLabelSize2x4 ScanningStationLabelSize = "2x4"
)

// Type of label printed at this station.
//
//   - `tag`: a label attached to the physical product.
//   - `traveler`: a routing sheet that accompanies the batch through every
//     production step.
type ScanningStationLabelType string

const (
	ScanningStationLabelTypeTag      ScanningStationLabelType = "tag"
	ScanningStationLabelTypeTraveler ScanningStationLabelType = "traveler"
)

// Resource type identifier.
type ScanningStationObject string

const (
	ScanningStationObjectScanningStation ScanningStationObject = "scanning_station"
)

// Whether operators must perform a material check at this station.
//
// - `none`: no additional operator check is required.
// - `material_check`: a material check is expected before the operation.
type ScanningStationOperatorRequirement string

const (
	ScanningStationOperatorRequirementNone          ScanningStationOperatorRequirement = "none"
	ScanningStationOperatorRequirementMaterialCheck ScanningStationOperatorRequirement = "material_check"
)

// Scanning station type, determining which batch operation the station performs.
//
// - `init_batch`: initializes a new batch.
// - `merge_batch`: merges multiple batches into one.
// - `move_batch`: moves a batch to another location or step.
// - `split_batch`: splits a batch into multiple batches.
type ScanningStationType string

const (
	ScanningStationTypeInitBatch  ScanningStationType = "init_batch"
	ScanningStationTypeMergeBatch ScanningStationType = "merge_batch"
	ScanningStationTypeMoveBatch  ScanningStationType = "move_batch"
	ScanningStationTypeSplitBatch ScanningStationType = "split_batch"
)

// A user's global profile, shared across every account they belong to.
//
// Account-specific settings (status, role, department) live on the account user
// resource that links the user to each account.
type User struct {
	// User ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address.
	Email string `json:"email" api:"required"`
	// When the user verified their email address.
	EmailVerifiedAt time.Time `json:"email_verified_at" api:"required" format:"date-time"`
	// URL of the user's profile image.
	ImageURL string `json:"image_url" api:"required"`
	// User's full display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "user".
	Object UserObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Username.
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		EmailVerifiedAt respjson.Field
		ImageURL        respjson.Field
		Name            respjson.Field
		Object          respjson.Field
		UpdatedAt       respjson.Field
		Username        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type UserObject string

const (
	UserObjectUser UserObject = "user"
)

type MessagingBlockDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingBlockDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingBlockDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingBlockNewParams struct {
	// Request to block another account user from messaging the caller.
	BlockRequest BlockRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "blocked_user", "blocked_user.user", "blocked_user.role",
	// "blocked_user.department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingBlockNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlockRequest)
}
func (r *MessagingBlockNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingBlockNewParams]'s query parameters as
// `url.Values`.
func (r MessagingBlockNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingBlockListParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "blocked_user", "blocked_user.user", "blocked_user.role",
	// "blocked_user.department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingBlockListParams]'s query parameters as
// `url.Values`.
func (r MessagingBlockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
