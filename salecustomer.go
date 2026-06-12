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

// Manage customer accounts.
//
// SaleCustomerService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleCustomerService] method instead.
type SaleCustomerService struct {
	options []option.RequestOption
	// Manage customer accounts.
	Actions SaleCustomerActionService
}

// NewSaleCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleCustomerService(opts ...option.RequestOption) (r SaleCustomerService) {
	r = SaleCustomerService{}
	r.options = opts
	r.Actions = NewSaleCustomerActionService(opts...)
	return
}

// Creates a customer account with its default addresses, fulfillment settings, and
// order policies.
//
// If `number` is omitted, the next sequential customer number is assigned
// automatically.
func (r *SaleCustomerService) New(ctx context.Context, params SaleCustomerNewParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a customer by ID.
func (r *SaleCustomerService) Get(ctx context.Context, id string, query SaleCustomerGetParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a customer account.
//
// Only the fields provided in the request are changed. Nullable fields can be set
// to `null` to clear their current value.
func (r *SaleCustomerService) Update(ctx context.Context, id string, params SaleCustomerUpdateParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of customers for the current account.
func (r *SaleCustomerService) List(ctx context.Context, query SaleCustomerListParams, opts ...option.RequestOption) (res *ListCustomer, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a customer.
//
// Fails with a conflict error if any sales orders still reference the customer;
// delete or reassign those orders, or merge the customer into another first.
func (r *SaleCustomerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleCustomerDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s", url.PathEscape(id))
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

// A shipping carrier configured for fulfilling orders.
//
// Carriers with a Shippo-supported `code` (`fedex`, `ups`, `usps`) are connected
// through Shippo for live rating and label purchase; other carriers represent
// self-managed shipping methods such as will call or local delivery.
type Carrier struct {
	// Carrier ID.
	ID string `json:"id" api:"required"`
	// Your account number with this carrier, used to connect UPS and USPS accounts.
	AccountNumber string `json:"account_number" api:"required"`
	// Well-known carrier identifier, set only for recognized carriers and absent for
	// custom ones.
	//
	//   - `fedex`, `ups`, `usps`: integrated carriers managed through Shippo (live
	//     rating and labels).
	//   - `will_call`: customer picks the order up; no carrier shipment.
	//   - `delivery`: delivered by your own vehicles/drivers.
	//   - `ltl`, `ltl1`: less-than-truckload freight carriers.
	//   - `freight_collect`: freight billed to and arranged by the receiver.
	//
	// Any of "fedex", "ups", "usps", "will_call", "delivery", "ltl", "ltl1",
	// "freight_collect".
	Code CarrierCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether customers can see and select this carrier at checkout in the customer
	// portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CarrierCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Soft-delete timestamp.
	DeletedAt time.Time `json:"deleted_at" api:"required" format:"date-time"`
	// Human-readable name for the carrier, unique among your account's carriers.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "carrier".
	Object CarrierObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// List represents a paginated list of resources.
	ServiceLevels ListServiceLevel `json:"service_levels" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		AccountNumber            respjson.Field
		Code                     respjson.Field
		CreatedAt                respjson.Field
		CustomerPortalVisibility respjson.Field
		DeletedAt                respjson.Field
		Name                     respjson.Field
		Object                   respjson.Field
		Owner                    respjson.Field
		ServiceLevels            respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Carrier) RawJSON() string { return r.JSON.raw }
func (r *Carrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Well-known carrier identifier, set only for recognized carriers and absent for
// custom ones.
//
//   - `fedex`, `ups`, `usps`: integrated carriers managed through Shippo (live
//     rating and labels).
//   - `will_call`: customer picks the order up; no carrier shipment.
//   - `delivery`: delivered by your own vehicles/drivers.
//   - `ltl`, `ltl1`: less-than-truckload freight carriers.
//   - `freight_collect`: freight billed to and arranged by the receiver.
type CarrierCode string

const (
	CarrierCodeFedex          CarrierCode = "fedex"
	CarrierCodeUps            CarrierCode = "ups"
	CarrierCodeUsps           CarrierCode = "usps"
	CarrierCodeWillCall       CarrierCode = "will_call"
	CarrierCodeDelivery       CarrierCode = "delivery"
	CarrierCodeLtl            CarrierCode = "ltl"
	CarrierCodeLtl1           CarrierCode = "ltl1"
	CarrierCodeFreightCollect CarrierCode = "freight_collect"
)

// Whether customers can see and select this carrier at checkout in the customer
// portal.
type CarrierCustomerPortalVisibility string

const (
	CarrierCustomerPortalVisibilityVisible CarrierCustomerPortalVisibility = "visible"
	CarrierCustomerPortalVisibilityHidden  CarrierCustomerPortalVisibility = "hidden"
)

// Resource type identifier.
type CarrierObject string

const (
	CarrierObjectCarrier CarrierObject = "carrier"
)

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

// Request to create a customer.
//
// The properties BillToAddress, CustomerTypeGroupID, DefaultCarrierID,
// DefaultPaymentTermID, DefaultShippingTermID, Name, ShipToAddress are required.
type CreateCustomerRequestParam struct {
	// Address details used to create an address, either directly or inline on another
	// resource.
	BillToAddress AddressInputParam `json:"bill_to_address,omitzero" api:"required"`
	// ID of the account group of type `type_group` that categorizes this customer (for
	// example "Distributors").
	CustomerTypeGroupID string `json:"customer_type_group_id" api:"required"`
	// ID of the default carrier for this customer's shipments.
	DefaultCarrierID string `json:"default_carrier_id" api:"required"`
	// Default payment term ID.
	DefaultPaymentTermID string `json:"default_payment_term_id" api:"required"`
	// Default shipping term ID.
	DefaultShippingTermID string `json:"default_shipping_term_id" api:"required"`
	// The customer's business name, as shown throughout the app and on documents.
	Name string `json:"name" api:"required"`
	// Address details used to create an address, either directly or inline on another
	// resource.
	ShipToAddress AddressInputParam `json:"ship_to_address,omitzero" api:"required"`
	// Carrier billing account number charged when `carrier_billing_type` is
	// `third_party`.
	CarrierBillingAccount param.Opt[string] `json:"carrier_billing_account,omitzero"`
	// The ID of the account user to assign as the default sales rep.
	DefaultSalesRepID param.Opt[string] `json:"default_sales_rep_id,omitzero"`
	// ID of the default carrier service level.
	DefaultServiceLevelID param.Opt[string] `json:"default_service_level_id,omitzero"`
	// Email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// Free-form note about the customer.
	Note param.Opt[string] `json:"note,omitzero"`
	// Human-readable customer number used to identify the account, distinct from the
	// `id`.
	//
	// Must be unique within your account. If omitted, the next sequential number is
	// assigned automatically.
	Number param.Opt[string] `json:"number,omitzero"`
	// Phone number.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Website URL.
	URL param.Opt[string] `json:"url,omitzero"`
	// Who pays the carrier for shipments.
	//
	// - `sender`: the shipper (you) pays the carrier.
	// - `third_party`: a third party is billed, using `carrier_billing_account`.
	//
	// Any of "sender", "third_party".
	CarrierBillingType CreateCustomerRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// How sales commission applies to this customer's orders.
	//
	//   - `commission_exempt`: this customer's orders are exempt from sales commission.
	//   - `commission_applied`: sales commission is calculated on this customer's
	//     orders.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CreateCustomerRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// A value with an associated unit, used in create and update requests.
	CreditLimit QuantityInputParam `json:"credit_limit,omitzero"`
	// IDs of the account groups of type `pricing_group` to assign to this customer,
	// used to apply pricing rules.
	CustomerPriceGroupIDs []string `json:"customer_price_group_ids,omitzero"`
	// Priority applied to new orders for this customer.
	//
	// Any of "low", "normal", "high".
	DefaultPriority CreateCustomerRequestDefaultPriority `json:"default_priority,omitzero"`
	// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
	// documents with this customer.
	//
	// Any of "enabled", "disabled".
	EdiStatus CreateCustomerRequestEdiStatus `json:"edi_status,omitzero"`
	// Whether this customer is billed for freight on their orders.
	//
	//   - `free_freight`: the customer is not billed for freight.
	//   - `billed_freight`: freight is billed to the customer, unless overridden on the
	//     order.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy CreateCustomerRequestFreightPolicy `json:"freight_policy,omitzero"`
	// Account status code, controlling whether the customer can transact.
	//
	// - `normal`: standard active account with no restrictions.
	// - `preferred`: active account flagged as preferred.
	// - `hold_shipment`: orders can be placed, but shipments are held.
	// - `hold_all`: all activity is on hold.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Status CreateCustomerRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r CreateCustomerRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCustomerRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCustomerRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who pays the carrier for shipments.
//
// - `sender`: the shipper (you) pays the carrier.
// - `third_party`: a third party is billed, using `carrier_billing_account`.
type CreateCustomerRequestCarrierBillingType string

const (
	CreateCustomerRequestCarrierBillingTypeSender     CreateCustomerRequestCarrierBillingType = "sender"
	CreateCustomerRequestCarrierBillingTypeThirdParty CreateCustomerRequestCarrierBillingType = "third_party"
)

// How sales commission applies to this customer's orders.
//
//   - `commission_exempt`: this customer's orders are exempt from sales commission.
//   - `commission_applied`: sales commission is calculated on this customer's
//     orders.
type CreateCustomerRequestCommissionPolicy string

const (
	CreateCustomerRequestCommissionPolicyCommissionApplied CreateCustomerRequestCommissionPolicy = "commission_applied"
	CreateCustomerRequestCommissionPolicyCommissionExempt  CreateCustomerRequestCommissionPolicy = "commission_exempt"
)

// Priority applied to new orders for this customer.
type CreateCustomerRequestDefaultPriority string

const (
	CreateCustomerRequestDefaultPriorityLow    CreateCustomerRequestDefaultPriority = "low"
	CreateCustomerRequestDefaultPriorityNormal CreateCustomerRequestDefaultPriority = "normal"
	CreateCustomerRequestDefaultPriorityHigh   CreateCustomerRequestDefaultPriority = "high"
)

// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
// documents with this customer.
type CreateCustomerRequestEdiStatus string

const (
	CreateCustomerRequestEdiStatusEnabled  CreateCustomerRequestEdiStatus = "enabled"
	CreateCustomerRequestEdiStatusDisabled CreateCustomerRequestEdiStatus = "disabled"
)

// Whether this customer is billed for freight on their orders.
//
//   - `free_freight`: the customer is not billed for freight.
//   - `billed_freight`: freight is billed to the customer, unless overridden on the
//     order.
type CreateCustomerRequestFreightPolicy string

const (
	CreateCustomerRequestFreightPolicyFreeFreight   CreateCustomerRequestFreightPolicy = "free_freight"
	CreateCustomerRequestFreightPolicyBilledFreight CreateCustomerRequestFreightPolicy = "billed_freight"
)

// Account status code, controlling whether the customer can transact.
//
// - `normal`: standard active account with no restrictions.
// - `preferred`: active account flagged as preferred.
// - `hold_shipment`: orders can be placed, but shipments are held.
// - `hold_all`: all activity is on hold.
type CreateCustomerRequestStatus string

const (
	CreateCustomerRequestStatusNormal       CreateCustomerRequestStatus = "normal"
	CreateCustomerRequestStatusPreferred    CreateCustomerRequestStatus = "preferred"
	CreateCustomerRequestStatusHoldShipment CreateCustomerRequestStatus = "hold_shipment"
	CreateCustomerRequestStatusHoldAll      CreateCustomerRequestStatus = "hold_all"
)

// A business you sell to, with its contact details, default fulfillment settings,
// and order policies.
type Customer struct {
	// Customer ID.
	ID string `json:"id" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	BillToAddress Address `json:"bill_to_address" api:"required"`
	// List represents a paginated list of resources.
	ChildAccounts *ListCustomer `json:"child_accounts" api:"required"`
	// How sales commission applies to this customer's orders.
	//
	//   - `commission_exempt`: this customer's orders are exempt from sales commission.
	//   - `commission_applied`: sales commission is calculated on this customer's
	//     orders.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CustomerCommissionPolicy `json:"commission_policy" api:"required"`
	// Customer contact information.
	ContactInfo CustomerContactInfo `json:"contact_info" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Value with an associated unit.
	CreditLimit Quantity `json:"credit_limit" api:"required"`
	// Customer default configuration.
	Defaults CustomerDefaults `json:"defaults" api:"required"`
	// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
	// documents with this customer.
	//
	// Any of "enabled", "disabled".
	EdiStatus CustomerEdiStatus `json:"edi_status" api:"required"`
	// Customer freight and carrier settings.
	FreightPreferences CustomerFreightPreferences `json:"freight_preferences" api:"required"`
	// The customer's business name, as shown throughout the app and on documents.
	Name string `json:"name" api:"required"`
	// Free-form note about the customer.
	Note string `json:"note" api:"required"`
	// Customer notification settings.
	NotificationPreferences CustomerNotificationPreferences `json:"notification_preferences" api:"required"`
	// Human-readable customer number used to identify the account, distinct from the
	// `id`.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer".
	Object CustomerObject `json:"object" api:"required"`
	// A business you sell to, with its contact details, default fulfillment settings,
	// and order policies.
	ParentAccount *Customer `json:"parent_account" api:"required"`
	// List represents a paginated list of resources.
	PriceGroups ListAccountGroup `json:"price_groups" api:"required"`
	// The customer's position in the account hierarchy.
	//
	// - `standalone`: no parent or child accounts.
	// - `parent`: has one or more child accounts (see `child_accounts`).
	// - `child`: belongs to a parent account (see `parent_account`).
	//
	// Any of "standalone", "parent", "child".
	RelationshipType CustomerRelationshipType `json:"relationship_type" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	ShipToAddress Address `json:"ship_to_address" api:"required"`
	// Account status code, controlling whether the customer can transact.
	//
	// - `normal`: standard active account with no restrictions.
	// - `preferred`: active account flagged as preferred.
	// - `hold_shipment`: orders can be placed, but shipments are held.
	// - `hold_all`: all activity is on hold.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Status CustomerStatus `json:"status" api:"required"`
	// A named grouping of customer accounts, used for pricing rules or to categorize
	// accounts.
	Type AccountGroup `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BillToAddress           respjson.Field
		ChildAccounts           respjson.Field
		CommissionPolicy        respjson.Field
		ContactInfo             respjson.Field
		CreatedAt               respjson.Field
		CreditLimit             respjson.Field
		Defaults                respjson.Field
		EdiStatus               respjson.Field
		FreightPreferences      respjson.Field
		Name                    respjson.Field
		Note                    respjson.Field
		NotificationPreferences respjson.Field
		Number                  respjson.Field
		Object                  respjson.Field
		ParentAccount           respjson.Field
		PriceGroups             respjson.Field
		RelationshipType        respjson.Field
		ShipToAddress           respjson.Field
		Status                  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Customer) RawJSON() string { return r.JSON.raw }
func (r *Customer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How sales commission applies to this customer's orders.
//
//   - `commission_exempt`: this customer's orders are exempt from sales commission.
//   - `commission_applied`: sales commission is calculated on this customer's
//     orders.
type CustomerCommissionPolicy string

const (
	CustomerCommissionPolicyCommissionApplied CustomerCommissionPolicy = "commission_applied"
	CustomerCommissionPolicyCommissionExempt  CustomerCommissionPolicy = "commission_exempt"
)

// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
// documents with this customer.
type CustomerEdiStatus string

const (
	CustomerEdiStatusEnabled  CustomerEdiStatus = "enabled"
	CustomerEdiStatusDisabled CustomerEdiStatus = "disabled"
)

// Resource type identifier.
type CustomerObject string

const (
	CustomerObjectCustomer CustomerObject = "customer"
)

// The customer's position in the account hierarchy.
//
// - `standalone`: no parent or child accounts.
// - `parent`: has one or more child accounts (see `child_accounts`).
// - `child`: belongs to a parent account (see `parent_account`).
type CustomerRelationshipType string

const (
	CustomerRelationshipTypeStandalone CustomerRelationshipType = "standalone"
	CustomerRelationshipTypeParent     CustomerRelationshipType = "parent"
	CustomerRelationshipTypeChild      CustomerRelationshipType = "child"
)

// Account status code, controlling whether the customer can transact.
//
// - `normal`: standard active account with no restrictions.
// - `preferred`: active account flagged as preferred.
// - `hold_shipment`: orders can be placed, but shipments are held.
// - `hold_all`: all activity is on hold.
type CustomerStatus string

const (
	CustomerStatusNormal       CustomerStatus = "normal"
	CustomerStatusPreferred    CustomerStatus = "preferred"
	CustomerStatusHoldShipment CustomerStatus = "hold_shipment"
	CustomerStatusHoldAll      CustomerStatus = "hold_all"
)

// Customer contact information.
type CustomerContactInfo struct {
	// Email address.
	Email string `json:"email" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_contact_info".
	Object CustomerContactInfoObject `json:"object" api:"required"`
	// Phone number.
	Phone string `json:"phone" api:"required"`
	// Website URL.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Object      respjson.Field
		Phone       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerContactInfo) RawJSON() string { return r.JSON.raw }
func (r *CustomerContactInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CustomerContactInfoObject string

const (
	CustomerContactInfoObjectCustomerContactInfo CustomerContactInfoObject = "customer_contact_info"
)

// Customer default configuration.
type CustomerDefaults struct {
	// Resource type identifier.
	//
	// Any of "customer_defaults".
	Object CustomerDefaultsObject `json:"object" api:"required"`
	// A payment term describing when payment is due (e.g. `Net 30`), assignable to
	// customers, sales orders, purchase orders, and invoices.
	PaymentTerm PaymentTerm `json:"payment_term" api:"required"`
	// Priority level used to order work on sales orders, purchase orders, and picks.
	Priority Priority `json:"priority" api:"required"`
	// A user's membership in an account, carrying the account-specific status, role,
	// and department.
	//
	// Profile fields (name, email, username, image URL) live on the expandable `user`
	// sub-resource, which is shared across every account the user belongs to.
	SalesRep AccountUser `json:"sales_rep" api:"required"`
	// A shipping term defining how freight charges are calculated for an order.
	ShippingTerm ShippingTerm `json:"shipping_term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object       respjson.Field
		PaymentTerm  respjson.Field
		Priority     respjson.Field
		SalesRep     respjson.Field
		ShippingTerm respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerDefaults) RawJSON() string { return r.JSON.raw }
func (r *CustomerDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CustomerDefaultsObject string

const (
	CustomerDefaultsObjectCustomerDefaults CustomerDefaultsObject = "customer_defaults"
)

// Customer freight and carrier settings.
type CustomerFreightPreferences struct {
	// Carrier billing account number charged when `billing_type` is `third_party`.
	BillingAccount string `json:"billing_account" api:"required"`
	// Who pays the carrier for shipments.
	//
	// - `sender`: the shipper (you) pays the carrier.
	// - `third_party`: a third party is billed, using `billing_account`.
	//
	// Any of "sender", "third_party".
	BillingType CustomerFreightPreferencesBillingType `json:"billing_type" api:"required"`
	// A shipping carrier configured for fulfilling orders.
	//
	// Carriers with a Shippo-supported `code` (`fedex`, `ups`, `usps`) are connected
	// through Shippo for live rating and label purchase; other carriers represent
	// self-managed shipping methods such as will call or local delivery.
	Carrier Carrier `json:"carrier" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_freight_preferences".
	Object CustomerFreightPreferencesObject `json:"object" api:"required"`
	// Shipping service level for a carrier.
	ServiceLevel ServiceLevel `json:"service_level" api:"required"`
	// Freight policy applied to this customer's orders.
	//
	//   - `free_freight`: the customer is not billed for freight.
	//   - `billed_freight`: freight is billed to the customer, unless overridden
	//     elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	Status CustomerFreightPreferencesStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAccount respjson.Field
		BillingType    respjson.Field
		Carrier        respjson.Field
		Object         respjson.Field
		ServiceLevel   respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerFreightPreferences) RawJSON() string { return r.JSON.raw }
func (r *CustomerFreightPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who pays the carrier for shipments.
//
// - `sender`: the shipper (you) pays the carrier.
// - `third_party`: a third party is billed, using `billing_account`.
type CustomerFreightPreferencesBillingType string

const (
	CustomerFreightPreferencesBillingTypeSender     CustomerFreightPreferencesBillingType = "sender"
	CustomerFreightPreferencesBillingTypeThirdParty CustomerFreightPreferencesBillingType = "third_party"
)

// Resource type identifier.
type CustomerFreightPreferencesObject string

const (
	CustomerFreightPreferencesObjectCustomerFreightPreferences CustomerFreightPreferencesObject = "customer_freight_preferences"
)

// Freight policy applied to this customer's orders.
//
//   - `free_freight`: the customer is not billed for freight.
//   - `billed_freight`: freight is billed to the customer, unless overridden
//     elsewhere.
type CustomerFreightPreferencesStatus string

const (
	CustomerFreightPreferencesStatusFreeFreight   CustomerFreightPreferencesStatus = "free_freight"
	CustomerFreightPreferencesStatusBilledFreight CustomerFreightPreferencesStatus = "billed_freight"
)

// Customer notification settings.
type CustomerNotificationPreferences struct {
	// Whether invoice emails are accepted.
	AcceptsInvoiceEmails bool `json:"accepts_invoice_emails" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_notification_preferences".
	Object CustomerNotificationPreferencesObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptsInvoiceEmails respjson.Field
		Object               respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerNotificationPreferences) RawJSON() string { return r.JSON.raw }
func (r *CustomerNotificationPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CustomerNotificationPreferencesObject string

const (
	CustomerNotificationPreferencesObjectCustomerNotificationPreferences CustomerNotificationPreferencesObject = "customer_notification_preferences"
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
type ListCustomer struct {
	// Resources in this page.
	Data []Customer `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListCustomerObject `json:"object" api:"required"`
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
func (r ListCustomer) RawJSON() string { return r.JSON.raw }
func (r *ListCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListCustomerObject string

const (
	ListCustomerObjectList ListCustomerObject = "list"
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

// List represents a paginated list of resources.
type ListServiceLevel struct {
	// Resources in this page.
	Data []ServiceLevel `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListServiceLevelObject `json:"object" api:"required"`
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
func (r ListServiceLevel) RawJSON() string { return r.JSON.raw }
func (r *ListServiceLevel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListServiceLevelObject string

const (
	ListServiceLevelObjectList ListServiceLevelObject = "list"
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

// A payment term describing when payment is due (e.g. `Net 30`), assignable to
// customers, sales orders, purchase orders, and invoices.
type PaymentTerm struct {
	// Payment term ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name (e.g. `Net 30`).
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "payment_term".
	Object PaymentTermObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Lifecycle status of the payment term.
	//
	// Any of "active", "inactive".
	Status PaymentTermStatus `json:"status" api:"required"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentTerm) RawJSON() string { return r.JSON.raw }
func (r *PaymentTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PaymentTermObject string

const (
	PaymentTermObjectPaymentTerm PaymentTermObject = "payment_term"
)

// Lifecycle status of the payment term.
type PaymentTermStatus string

const (
	PaymentTermStatusActive   PaymentTermStatus = "active"
	PaymentTermStatusInactive PaymentTermStatus = "inactive"
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

// A value with an associated unit, used in create and update requests.
//
// The properties UnitID, Value are required.
type QuantityInputParam struct {
	// ID of the unit of measure for the value.
	UnitID string `json:"unit_id" api:"required"`
	// Decimal value, as a string to preserve precision.
	Value string `json:"value" api:"required" format:"decimal"`
	paramObj
}

func (r QuantityInputParam) MarshalJSON() (data []byte, err error) {
	type shadow QuantityInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuantityInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// Shipping service level for a carrier.
type ServiceLevel struct {
	// Service level ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether customers can see and select this service level at checkout in the
	// customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility ServiceLevelCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Whether this is the carrier's default service level, pre-selected when the
	// carrier is chosen.
	//
	// Each carrier has at most one default; setting a new default clears the previous
	// one.
	IsDefault bool `json:"is_default" api:"required"`
	// Human-readable name for the service level, shown to customers at checkout when
	// the service level is visible.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "service_level".
	Object ServiceLevelObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Carrier-specific code identifying this service level (e.g. `fedex_ground`,
	// `ups_next_day_air`).
	//
	// Values are carrier-defined, so any non-empty string is accepted.
	ServiceLevelToken string `json:"service_level_token" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		CustomerPortalVisibility respjson.Field
		IsDefault                respjson.Field
		Name                     respjson.Field
		Object                   respjson.Field
		Owner                    respjson.Field
		ServiceLevelToken        respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceLevel) RawJSON() string { return r.JSON.raw }
func (r *ServiceLevel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether customers can see and select this service level at checkout in the
// customer portal.
type ServiceLevelCustomerPortalVisibility string

const (
	ServiceLevelCustomerPortalVisibilityVisible ServiceLevelCustomerPortalVisibility = "visible"
	ServiceLevelCustomerPortalVisibilityHidden  ServiceLevelCustomerPortalVisibility = "hidden"
)

// Resource type identifier.
type ServiceLevelObject string

const (
	ServiceLevelObjectServiceLevel ServiceLevelObject = "service_level"
)

// A shipping term defining how freight charges are calculated for an order.
type ShippingTerm struct {
	// Shipping term ID.
	ID string `json:"id" api:"required"`
	// When this shipping term was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Value with an associated unit.
	FlatRate Quantity `json:"flat_rate" api:"required"`
	// List represents a paginated list of resources.
	FreeShippingServiceLevels ListServiceLevel `json:"free_shipping_service_levels" api:"required"`
	// Value with an associated unit.
	MinimumOrderValue Quantity `json:"minimum_order_value" api:"required"`
	// Human-readable name for the shipping term, used to identify it when assigning
	// shipping terms to customers and orders.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "shipping_term".
	Object ShippingTermObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Freight pricing model applied by this shipping term.
	//
	//   - `free_freight`: no shipping cost to the buyer.
	//   - `flat_rate_freight`: a fixed shipping cost regardless of order details (see
	//     `flat_rate`).
	//   - `carrier_rate_freight`: shipping cost is determined by the carrier's quoted
	//     rate.
	//
	// Any of "free_freight", "flat_rate_freight", "carrier_rate_freight".
	Type ShippingTermType `json:"type" api:"required"`
	// When this shipping term was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CreatedAt                 respjson.Field
		FlatRate                  respjson.Field
		FreeShippingServiceLevels respjson.Field
		MinimumOrderValue         respjson.Field
		Name                      respjson.Field
		Object                    respjson.Field
		Owner                     respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShippingTerm) RawJSON() string { return r.JSON.raw }
func (r *ShippingTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ShippingTermObject string

const (
	ShippingTermObjectShippingTerm ShippingTermObject = "shipping_term"
)

// Freight pricing model applied by this shipping term.
//
//   - `free_freight`: no shipping cost to the buyer.
//   - `flat_rate_freight`: a fixed shipping cost regardless of order details (see
//     `flat_rate`).
//   - `carrier_rate_freight`: shipping cost is determined by the carrier's quoted
//     rate.
type ShippingTermType string

const (
	ShippingTermTypeFreeFreight        ShippingTermType = "free_freight"
	ShippingTermTypeFlatRateFreight    ShippingTermType = "flat_rate_freight"
	ShippingTermTypeCarrierRateFreight ShippingTermType = "carrier_rate_freight"
)

// Request to partially update a customer.
type UpdateCustomerRequestParam struct {
	// ID of an existing address to use as the default billing address.
	BillToAddressID param.Opt[string] `json:"bill_to_address_id,omitzero"`
	// Carrier billing account number charged when `carrier_billing_type` is
	// `third_party`.
	CarrierBillingAccount param.Opt[string] `json:"carrier_billing_account,omitzero"`
	// The ID of the account user to assign as the default sales rep.
	DefaultSalesRepID param.Opt[string] `json:"default_sales_rep_id,omitzero"`
	// ID of the default carrier service level.
	DefaultServiceLevelID param.Opt[string] `json:"default_service_level_id,omitzero"`
	// Email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// Free-form note about the customer.
	Note param.Opt[string] `json:"note,omitzero"`
	// Phone number.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// ID of an existing address to use as the default shipping address.
	ShipToAddressID param.Opt[string] `json:"ship_to_address_id,omitzero"`
	// Website URL.
	URL param.Opt[string] `json:"url,omitzero"`
	// ID of the account group of type `type_group` that categorizes this customer (for
	// example "Distributors").
	CustomerTypeGroupID param.Opt[string] `json:"customer_type_group_id,omitzero"`
	// ID of the default carrier for this customer's shipments.
	DefaultCarrierID param.Opt[string] `json:"default_carrier_id,omitzero"`
	// Default payment term ID.
	DefaultPaymentTermID param.Opt[string] `json:"default_payment_term_id,omitzero"`
	// Default shipping term ID.
	DefaultShippingTermID param.Opt[string] `json:"default_shipping_term_id,omitzero"`
	// The customer's business name, as shown throughout the app and on documents.
	Name param.Opt[string] `json:"name,omitzero"`
	// Human-readable customer number used to identify the account, distinct from the
	// `id`.
	//
	// Must be unique within your account.
	Number param.Opt[string] `json:"number,omitzero"`
	// Who pays the carrier for shipments.
	//
	// - `sender`: the shipper (you) pays the carrier.
	// - `third_party`: a third party is billed, using `carrier_billing_account`.
	//
	// Any of "sender", "third_party".
	CarrierBillingType UpdateCustomerRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// How sales commission applies to this customer's orders.
	//
	//   - `commission_exempt`: this customer's orders are exempt from sales commission.
	//   - `commission_applied`: sales commission is calculated on this customer's
	//     orders.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy UpdateCustomerRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// A value with an associated unit, used in create and update requests.
	CreditLimit QuantityInputParam `json:"credit_limit,omitzero"`
	// IDs of the account groups of type `pricing_group` to assign to this customer,
	// used to apply pricing rules.
	//
	// When provided, replaces the customer's full set of existing price groups.
	CustomerPriceGroupIDs []string `json:"customer_price_group_ids,omitzero"`
	// Priority applied to new orders for this customer.
	//
	// Any of "low", "normal", "high".
	DefaultPriority UpdateCustomerRequestDefaultPriority `json:"default_priority,omitzero"`
	// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
	// documents with this customer.
	//
	// Any of "enabled", "disabled".
	EdiStatus UpdateCustomerRequestEdiStatus `json:"edi_status,omitzero"`
	// Whether this customer is billed for freight on their orders.
	//
	//   - `free_freight`: the customer is not billed for freight.
	//   - `billed_freight`: freight is billed to the customer, unless overridden on the
	//     order.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy UpdateCustomerRequestFreightPolicy `json:"freight_policy,omitzero"`
	// Account status code, controlling whether the customer can transact.
	//
	// - `normal`: standard active account with no restrictions.
	// - `preferred`: active account flagged as preferred.
	// - `hold_shipment`: orders can be placed, but shipments are held.
	// - `hold_all`: all activity is on hold.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Status UpdateCustomerRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r UpdateCustomerRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCustomerRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCustomerRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who pays the carrier for shipments.
//
// - `sender`: the shipper (you) pays the carrier.
// - `third_party`: a third party is billed, using `carrier_billing_account`.
type UpdateCustomerRequestCarrierBillingType string

const (
	UpdateCustomerRequestCarrierBillingTypeSender     UpdateCustomerRequestCarrierBillingType = "sender"
	UpdateCustomerRequestCarrierBillingTypeThirdParty UpdateCustomerRequestCarrierBillingType = "third_party"
)

// How sales commission applies to this customer's orders.
//
//   - `commission_exempt`: this customer's orders are exempt from sales commission.
//   - `commission_applied`: sales commission is calculated on this customer's
//     orders.
type UpdateCustomerRequestCommissionPolicy string

const (
	UpdateCustomerRequestCommissionPolicyCommissionApplied UpdateCustomerRequestCommissionPolicy = "commission_applied"
	UpdateCustomerRequestCommissionPolicyCommissionExempt  UpdateCustomerRequestCommissionPolicy = "commission_exempt"
)

// Priority applied to new orders for this customer.
type UpdateCustomerRequestDefaultPriority string

const (
	UpdateCustomerRequestDefaultPriorityLow    UpdateCustomerRequestDefaultPriority = "low"
	UpdateCustomerRequestDefaultPriorityNormal UpdateCustomerRequestDefaultPriority = "normal"
	UpdateCustomerRequestDefaultPriorityHigh   UpdateCustomerRequestDefaultPriority = "high"
)

// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
// documents with this customer.
type UpdateCustomerRequestEdiStatus string

const (
	UpdateCustomerRequestEdiStatusEnabled  UpdateCustomerRequestEdiStatus = "enabled"
	UpdateCustomerRequestEdiStatusDisabled UpdateCustomerRequestEdiStatus = "disabled"
)

// Whether this customer is billed for freight on their orders.
//
//   - `free_freight`: the customer is not billed for freight.
//   - `billed_freight`: freight is billed to the customer, unless overridden on the
//     order.
type UpdateCustomerRequestFreightPolicy string

const (
	UpdateCustomerRequestFreightPolicyFreeFreight   UpdateCustomerRequestFreightPolicy = "free_freight"
	UpdateCustomerRequestFreightPolicyBilledFreight UpdateCustomerRequestFreightPolicy = "billed_freight"
)

// Account status code, controlling whether the customer can transact.
//
// - `normal`: standard active account with no restrictions.
// - `preferred`: active account flagged as preferred.
// - `hold_shipment`: orders can be placed, but shipments are held.
// - `hold_all`: all activity is on hold.
type UpdateCustomerRequestStatus string

const (
	UpdateCustomerRequestStatusNormal       UpdateCustomerRequestStatus = "normal"
	UpdateCustomerRequestStatusPreferred    UpdateCustomerRequestStatus = "preferred"
	UpdateCustomerRequestStatusHoldShipment UpdateCustomerRequestStatus = "hold_shipment"
	UpdateCustomerRequestStatusHoldAll      UpdateCustomerRequestStatus = "hold_all"
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

type SaleCustomerDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleCustomerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleCustomerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleCustomerNewParams struct {
	// Request to create a customer.
	CreateCustomerRequest CreateCustomerRequestParam
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

func (r SaleCustomerNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCustomerRequest)
}
func (r *SaleCustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleCustomerNewParams]'s query parameters as `url.Values`.
func (r SaleCustomerNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleCustomerGetParams struct {
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

// URLQuery serializes [SaleCustomerGetParams]'s query parameters as `url.Values`.
func (r SaleCustomerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleCustomerUpdateParams struct {
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
	// Request to partially update a customer.
	UpdateCustomerRequest UpdateCustomerRequestParam
	paramObj
}

func (r SaleCustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCustomerRequest)
}
func (r *SaleCustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleCustomerUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleCustomerUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleCustomerListParams struct {
	// Filter to customers with any address in this city (exact match).
	//
	// When combined with `state` or `postal_code`, a single address must match all
	// provided values.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter to customers created at or before this timestamp (inclusive).
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter to customers with any address in this postal code (exact match).
	PostalCode param.Opt[string] `query:"postal_code,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter to customers created at or after this timestamp (inclusive).
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	// Filter to customers with any address in this state (exact match).
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	// Filter by default carrier IDs.
	CarrierIDs []string `query:"carrier_ids,omitzero" json:"-"`
	// Filter by commission policy.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionStatusCodes []string `query:"commission_status_codes,omitzero" json:"-"`
	// Filter by customer type group IDs (the account group of type `type_group`
	// returned in the customer's `type` field).
	CustomerGroupIDs []string `query:"customer_group_ids,omitzero" json:"-"`
	// Filter by freight policy.
	//
	// Any of "free_freight", "billed_freight".
	FreightStatusCodes []string `query:"freight_status_codes,omitzero" json:"-"`
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
	// Filter by whether the customer has child accounts.
	//
	// Any of "parent", "non_parent".
	ParentAccountStatus SaleCustomerListParamsParentAccountStatus `query:"parent_account_status,omitzero" json:"-"`
	// Filter by default payment term IDs.
	PaymentTermIDs []string `query:"payment_term_ids,omitzero" json:"-"`
	// Filter to customers that belong to any of these pricing groups.
	PricingGroupIDs []string `query:"pricing_group_ids,omitzero" json:"-"`
	// Filter by default sales rep IDs.
	SalesRepIDs []string `query:"sales_rep_ids,omitzero" json:"-"`
	// Filter by default service level IDs.
	ServiceLevelIDs []string `query:"service_level_ids,omitzero" json:"-"`
	// Filter by default shipping term IDs.
	ShippingTermIDs []string `query:"shipping_term_ids,omitzero" json:"-"`
	// Filter by account status codes.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	StatusCodes []string `query:"status_codes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleCustomerListParams]'s query parameters as `url.Values`.
func (r SaleCustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by whether the customer has child accounts.
type SaleCustomerListParamsParentAccountStatus string

const (
	SaleCustomerListParamsParentAccountStatusParent    SaleCustomerListParamsParentAccountStatus = "parent"
	SaleCustomerListParamsParentAccountStatusNonParent SaleCustomerListParamsParentAccountStatus = "non_parent"
)
