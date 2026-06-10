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

// Creates a customer account. Auto-generates a customer number if one is not
// provided.
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

// Partially updates a customer account. When a Stripe integration is active,
// customer changes are synced to Stripe.
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

// Deletes a customer and associated account relations, addresses, and account
// users.
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

// Account user with profile, role, and department.
type AccountUser struct {
	// Account user ID.
	ID string `json:"id" api:"required"`
	// When the account user was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Department resource.
	Department Department `json:"department" api:"required"`
	// Email address.
	Email string `json:"email" api:"required"`
	// Profile image URL.
	ImageURL string `json:"image_url" api:"required"`
	// When the user last used this account.
	LastUsedAt time.Time `json:"last_used_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_user".
	Object AccountUserObject `json:"object" api:"required"`
	// Role resource.
	Role Role `json:"role" api:"required"`
	// Account user status.
	//
	// Any of "active", "disabled", "removed".
	Status AccountUserStatus `json:"status" api:"required"`
	// When the account user was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	User Entity `json:"user" api:"required"`
	// Username.
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Department  respjson.Field
		Email       respjson.Field
		ImageURL    respjson.Field
		LastUsedAt  respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Role        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		User        respjson.Field
		Username    respjson.Field
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
type AccountUserStatus string

const (
	AccountUserStatusActive   AccountUserStatus = "active"
	AccountUserStatusDisabled AccountUserStatus = "disabled"
	AccountUserStatusRemoved  AccountUserStatus = "removed"
)

// Carrier resource.
type Carrier struct {
	// Carrier ID.
	ID string `json:"id" api:"required"`
	// Account number.
	AccountNumber string `json:"account_number" api:"required"`
	// Carrier code.
	//
	// Any of "fedex", "ups", "usps", "will_call", "delivery", "ltl", "ltl1",
	// "freight_collect".
	Code CarrierCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Customer portal visibility.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CarrierCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Soft-delete timestamp.
	DeletedAt time.Time `json:"deleted_at" api:"required" format:"date-time"`
	// Display name.
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

// Carrier code.
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

// Customer portal visibility.
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
	// Request to create an address.
	BillToAddress AddressInputParam `json:"bill_to_address,omitzero" api:"required"`
	// Customer type group ID.
	CustomerTypeGroupID string `json:"customer_type_group_id" api:"required"`
	// Default carrier ID.
	DefaultCarrierID string `json:"default_carrier_id" api:"required"`
	// Default payment term ID.
	DefaultPaymentTermID string `json:"default_payment_term_id" api:"required"`
	// Default shipping term ID.
	DefaultShippingTermID string `json:"default_shipping_term_id" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Request to create an address.
	ShipToAddress AddressInputParam `json:"ship_to_address,omitzero" api:"required"`
	// Carrier billing account number.
	CarrierBillingAccount param.Opt[string] `json:"carrier_billing_account,omitzero"`
	// The ID of the account user to assign as the default sales rep.
	DefaultSalesRepID param.Opt[string] `json:"default_sales_rep_id,omitzero"`
	// Default service level ID.
	DefaultServiceLevelID param.Opt[string] `json:"default_service_level_id,omitzero"`
	// Email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// Note.
	Note param.Opt[string] `json:"note,omitzero"`
	// Customer number. Auto-generated if omitted.
	Number param.Opt[string] `json:"number,omitzero"`
	// Phone number.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Website URL.
	URL param.Opt[string] `json:"url,omitzero"`
	// Carrier billing type.
	//
	// Any of "sender", "third_party".
	CarrierBillingType CreateCustomerRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// Commission policy.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CreateCustomerRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// QuantityInput represents a value with an associated unit for create/update
	// requests.
	CreditLimit QuantityInputParam `json:"credit_limit,omitzero"`
	// Price group IDs.
	CustomerPriceGroupIDs []string `json:"customer_price_group_ids,omitzero"`
	// Default priority code.
	//
	// Any of "low", "normal", "high".
	DefaultPriority CreateCustomerRequestDefaultPriority `json:"default_priority,omitzero"`
	// EDI status.
	//
	// Any of "enabled", "disabled".
	EdiStatus CreateCustomerRequestEdiStatus `json:"edi_status,omitzero"`
	// Freight policy.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy CreateCustomerRequestFreightPolicy `json:"freight_policy,omitzero"`
	// Account status code.
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

// Carrier billing type.
type CreateCustomerRequestCarrierBillingType string

const (
	CreateCustomerRequestCarrierBillingTypeSender     CreateCustomerRequestCarrierBillingType = "sender"
	CreateCustomerRequestCarrierBillingTypeThirdParty CreateCustomerRequestCarrierBillingType = "third_party"
)

// Commission policy.
type CreateCustomerRequestCommissionPolicy string

const (
	CreateCustomerRequestCommissionPolicyCommissionApplied CreateCustomerRequestCommissionPolicy = "commission_applied"
	CreateCustomerRequestCommissionPolicyCommissionExempt  CreateCustomerRequestCommissionPolicy = "commission_exempt"
)

// Default priority code.
type CreateCustomerRequestDefaultPriority string

const (
	CreateCustomerRequestDefaultPriorityLow    CreateCustomerRequestDefaultPriority = "low"
	CreateCustomerRequestDefaultPriorityNormal CreateCustomerRequestDefaultPriority = "normal"
	CreateCustomerRequestDefaultPriorityHigh   CreateCustomerRequestDefaultPriority = "high"
)

// EDI status.
type CreateCustomerRequestEdiStatus string

const (
	CreateCustomerRequestEdiStatusEnabled  CreateCustomerRequestEdiStatus = "enabled"
	CreateCustomerRequestEdiStatusDisabled CreateCustomerRequestEdiStatus = "disabled"
)

// Freight policy.
type CreateCustomerRequestFreightPolicy string

const (
	CreateCustomerRequestFreightPolicyFreeFreight   CreateCustomerRequestFreightPolicy = "free_freight"
	CreateCustomerRequestFreightPolicyBilledFreight CreateCustomerRequestFreightPolicy = "billed_freight"
)

// Account status code.
type CreateCustomerRequestStatus string

const (
	CreateCustomerRequestStatusNormal       CreateCustomerRequestStatus = "normal"
	CreateCustomerRequestStatusPreferred    CreateCustomerRequestStatus = "preferred"
	CreateCustomerRequestStatusHoldShipment CreateCustomerRequestStatus = "hold_shipment"
	CreateCustomerRequestStatusHoldAll      CreateCustomerRequestStatus = "hold_all"
)

// Customer account.
type Customer struct {
	// Customer ID.
	ID string `json:"id" api:"required"`
	// Address with associated geolocation.
	BillToAddress Address `json:"bill_to_address" api:"required"`
	// List represents a paginated list of resources.
	ChildAccounts *ListCustomer `json:"child_accounts" api:"required"`
	// Commission policy.
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
	// EDI status.
	//
	// Any of "enabled", "disabled".
	EdiStatus CustomerEdiStatus `json:"edi_status" api:"required"`
	// Customer freight and carrier settings.
	FreightPreferences CustomerFreightPreferences `json:"freight_preferences" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Note.
	Note string `json:"note" api:"required"`
	// Customer notification settings.
	NotificationPreferences CustomerNotificationPreferences `json:"notification_preferences" api:"required"`
	// Customer number.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer".
	Object CustomerObject `json:"object" api:"required"`
	// Customer account.
	ParentAccount *Customer `json:"parent_account" api:"required"`
	// List represents a paginated list of resources.
	PriceGroups ListAccountGroup `json:"price_groups" api:"required"`
	// Customer relationship type.
	//
	// Any of "standalone", "parent", "child".
	RelationshipType CustomerRelationshipType `json:"relationship_type" api:"required"`
	// Address with associated geolocation.
	ShipToAddress Address `json:"ship_to_address" api:"required"`
	// Account status code.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Status CustomerStatus `json:"status" api:"required"`
	// Account group resource.
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

// Commission policy.
type CustomerCommissionPolicy string

const (
	CustomerCommissionPolicyCommissionApplied CustomerCommissionPolicy = "commission_applied"
	CustomerCommissionPolicyCommissionExempt  CustomerCommissionPolicy = "commission_exempt"
)

// EDI status.
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

// Customer relationship type.
type CustomerRelationshipType string

const (
	CustomerRelationshipTypeStandalone CustomerRelationshipType = "standalone"
	CustomerRelationshipTypeParent     CustomerRelationshipType = "parent"
	CustomerRelationshipTypeChild      CustomerRelationshipType = "child"
)

// Account status code.
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
	// Payment term resource.
	PaymentTerm PaymentTerm `json:"payment_term" api:"required"`
	// Priority level used by sales orders and picks.
	Priority Priority `json:"priority" api:"required"`
	// Account user with profile, role, and department.
	SalesRep AccountUser `json:"sales_rep" api:"required"`
	// ShippingTerm resource.
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
	// Carrier billing account number.
	BillingAccount string `json:"billing_account" api:"required"`
	// Carrier billing type.
	//
	// Any of "sender", "third_party".
	BillingType CustomerFreightPreferencesBillingType `json:"billing_type" api:"required"`
	// Carrier resource.
	Carrier Carrier `json:"carrier" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_freight_preferences".
	Object CustomerFreightPreferencesObject `json:"object" api:"required"`
	// Shipping service level for a carrier.
	ServiceLevel ServiceLevel `json:"service_level" api:"required"`
	// Freight policy.
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

// Carrier billing type.
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

// Freight policy.
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

// Department resource.
type Department struct {
	// Department ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Location resource.
	Location Location `json:"location" api:"required"`
	// List represents a paginated list of resources.
	Machines *ListMachine `json:"machines" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Notes about the department.
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

// Entity is a polymorphic reference to any resource in the system.
type Entity struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// Secondary human-readable identifier (e.g. email address, username, redacted API
	// key value).
	Handle string `json:"handle" api:"required"`
	// Human-readable display name for the entity (e.g. a user's full name, a sales
	// order number).
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "entity".
	Object EntityObject `json:"object" api:"required"`
	// The resource kind that this entity references.
	//
	// Any of "account", "actor", "entity", "record", "freight", "sales_order_totals",
	// "sales_order_related", "user", "address", "api_key", "created_api_key",
	// "refresh_token", "list", "sandbox", "registration_session", "pricing_plan",
	// "account_plan", "plan_change", "enterprise_inquiry", "request_log",
	// "audit_event", "audit_field_change", "role", "unit", "account_affiliation",
	// "agent_definition", "available_tool", "agent_definition_tool",
	// "agent_account_status", "agent_run", "agent_action", "agent_run_step",
	// "agent_token_usage", "agent_memory", "agent_alert", "tool_group",
	// "payment_term", "shipping_term", "quantity", "account_group", "account_status",
	// "geolocation", "account_user", "department", "account_integration",
	// "account_price", "product_line", "item_category", "attribute", "rate",
	// "account_group_product_line_access", "sales_target", "adjustment_type",
	// "account_branding", "account_portal", "account_logo_url", "public_account",
	// "property", "carrier", "service_level", "item", "item_inventory", "product",
	// "batch", "batch_flow_node", "scanning_consumption", "open_batch_summary",
	// "scanning_production_step_info", "scanning_station", "production_step",
	// "production_run", "machine", "child_account", "unit_group", "unit_group_unit",
	// "consumption", "customer_product_line_access", "customer",
	// "frequently_ordered_product", "priority", "delivery", "delivery_line",
	// "sales_order", "location", "location_type", "lot", "email_log",
	// "inventory_change_log", "invoice", "invoice_summary", "invoice_line",
	// "invoice_allocation", "invoice_for_payment", "shipment", "shipment_summary",
	// "shipment_line", "shipping_case", "shipping_case_label_url", "settlement",
	// "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "agent_token_detail", "account_usage_response",
	// "subscription_info", "billing_portal_session_response", "switch_plan_response",
	// "ensure_billing_customer_response", "spending_cap_response", "agent_spend_info",
	// "webhook_response", "address_suggestion", "address_components",
	// "address_details_result", "validated_address", "plan_limit",
	// "plan_change_proration", "plan_change_line_item", "setup_billing_response",
	// "confirm_payment_response", "oauth_response", "oauth_status_response",
	// "stripe_publishable_key", "stripe_status", "healthcheck",
	// "agent_definition_config", "trigger_config", "customer_contact_info",
	// "customer_freight_preferences", "customer_defaults",
	// "customer_notification_preferences", "order_discount", "sales_order_line",
	// "sales_order_type", "sales_order_status", "material", "supplier_material",
	// "part", "permission_group", "permission", "pick", "pick_line", "product_type",
	// "production", "production_flow", "map", "purchase_order", "purchase_order_line",
	// "supplier", "supplier_summary", "receivable_entry", "receiving_order",
	// "receiving_order_line", "email_contact", "allocation_entry",
	// "open_credit_entry", "volume_discount", "volume_discount_tier",
	// "analyze_deliveries_response", "analyze_manufacturing_response",
	// "analyze_manufacturing_batch_response", "analyze_quarterly_orders_response",
	// "analyze_new_customers_response", "analyze_oee_response",
	// "catalog_product_line", "catalog_category", "catalog_product",
	// "catalog_property", "catalog_attribute", "dc_location", "edi_run",
	// "inventory_item", "analyze_weeks_of_sales_response",
	// "bulk_reconcile_items_response", "sys_property", "sys_property_type",
	// "sys_property_value", "territory", "tenancy", "checkout_session",
	// "estimate_rate_result", "rate_shop_option", "rate_shop_result", "owner",
	// "message", "account_photo_upload_result", "user_photo_upload_result",
	// "user_photo_url", "batch_lot", "check_duplicate_result", "item_trend_point",
	// "pack_pick_response", "pick_shipments_response", "tenancy_pending_registration",
	// "invoice_allocation_entry", "allocation_customer",
	// "checkout_sales_order_response", "create_production_run_response".
	Type EntityType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handle      respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Entity) RawJSON() string { return r.JSON.raw }
func (r *Entity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type EntityObject string

const (
	EntityObjectEntity EntityObject = "entity"
)

// The resource kind that this entity references.
type EntityType string

const (
	EntityTypeAccount                           EntityType = "account"
	EntityTypeActor                             EntityType = "actor"
	EntityTypeEntity                            EntityType = "entity"
	EntityTypeRecord                            EntityType = "record"
	EntityTypeFreight                           EntityType = "freight"
	EntityTypeSalesOrderTotals                  EntityType = "sales_order_totals"
	EntityTypeSalesOrderRelated                 EntityType = "sales_order_related"
	EntityTypeUser                              EntityType = "user"
	EntityTypeAddress                           EntityType = "address"
	EntityTypeAPIKey                            EntityType = "api_key"
	EntityTypeCreatedAPIKey                     EntityType = "created_api_key"
	EntityTypeRefreshToken                      EntityType = "refresh_token"
	EntityTypeList                              EntityType = "list"
	EntityTypeSandbox                           EntityType = "sandbox"
	EntityTypeRegistrationSession               EntityType = "registration_session"
	EntityTypePricingPlan                       EntityType = "pricing_plan"
	EntityTypeAccountPlan                       EntityType = "account_plan"
	EntityTypePlanChange                        EntityType = "plan_change"
	EntityTypeEnterpriseInquiry                 EntityType = "enterprise_inquiry"
	EntityTypeRequestLog                        EntityType = "request_log"
	EntityTypeAuditEvent                        EntityType = "audit_event"
	EntityTypeAuditFieldChange                  EntityType = "audit_field_change"
	EntityTypeRole                              EntityType = "role"
	EntityTypeUnit                              EntityType = "unit"
	EntityTypeAccountAffiliation                EntityType = "account_affiliation"
	EntityTypeAgentDefinition                   EntityType = "agent_definition"
	EntityTypeAvailableTool                     EntityType = "available_tool"
	EntityTypeAgentDefinitionTool               EntityType = "agent_definition_tool"
	EntityTypeAgentAccountStatus                EntityType = "agent_account_status"
	EntityTypeAgentRun                          EntityType = "agent_run"
	EntityTypeAgentAction                       EntityType = "agent_action"
	EntityTypeAgentRunStep                      EntityType = "agent_run_step"
	EntityTypeAgentTokenUsage                   EntityType = "agent_token_usage"
	EntityTypeAgentMemory                       EntityType = "agent_memory"
	EntityTypeAgentAlert                        EntityType = "agent_alert"
	EntityTypeToolGroup                         EntityType = "tool_group"
	EntityTypePaymentTerm                       EntityType = "payment_term"
	EntityTypeShippingTerm                      EntityType = "shipping_term"
	EntityTypeQuantity                          EntityType = "quantity"
	EntityTypeAccountGroup                      EntityType = "account_group"
	EntityTypeAccountStatus                     EntityType = "account_status"
	EntityTypeGeolocation                       EntityType = "geolocation"
	EntityTypeAccountUser                       EntityType = "account_user"
	EntityTypeDepartment                        EntityType = "department"
	EntityTypeAccountIntegration                EntityType = "account_integration"
	EntityTypeAccountPrice                      EntityType = "account_price"
	EntityTypeProductLine                       EntityType = "product_line"
	EntityTypeItemCategory                      EntityType = "item_category"
	EntityTypeAttribute                         EntityType = "attribute"
	EntityTypeRate                              EntityType = "rate"
	EntityTypeAccountGroupProductLineAccess     EntityType = "account_group_product_line_access"
	EntityTypeSalesTarget                       EntityType = "sales_target"
	EntityTypeAdjustmentType                    EntityType = "adjustment_type"
	EntityTypeAccountBranding                   EntityType = "account_branding"
	EntityTypeAccountPortal                     EntityType = "account_portal"
	EntityTypeAccountLogoURL                    EntityType = "account_logo_url"
	EntityTypePublicAccount                     EntityType = "public_account"
	EntityTypeProperty                          EntityType = "property"
	EntityTypeCarrier                           EntityType = "carrier"
	EntityTypeServiceLevel                      EntityType = "service_level"
	EntityTypeItem                              EntityType = "item"
	EntityTypeItemInventory                     EntityType = "item_inventory"
	EntityTypeProduct                           EntityType = "product"
	EntityTypeBatch                             EntityType = "batch"
	EntityTypeBatchFlowNode                     EntityType = "batch_flow_node"
	EntityTypeScanningConsumption               EntityType = "scanning_consumption"
	EntityTypeOpenBatchSummary                  EntityType = "open_batch_summary"
	EntityTypeScanningProductionStepInfo        EntityType = "scanning_production_step_info"
	EntityTypeScanningStation                   EntityType = "scanning_station"
	EntityTypeProductionStep                    EntityType = "production_step"
	EntityTypeProductionRun                     EntityType = "production_run"
	EntityTypeMachine                           EntityType = "machine"
	EntityTypeChildAccount                      EntityType = "child_account"
	EntityTypeUnitGroup                         EntityType = "unit_group"
	EntityTypeUnitGroupUnit                     EntityType = "unit_group_unit"
	EntityTypeConsumption                       EntityType = "consumption"
	EntityTypeCustomerProductLineAccess         EntityType = "customer_product_line_access"
	EntityTypeCustomer                          EntityType = "customer"
	EntityTypeFrequentlyOrderedProduct          EntityType = "frequently_ordered_product"
	EntityTypePriority                          EntityType = "priority"
	EntityTypeDelivery                          EntityType = "delivery"
	EntityTypeDeliveryLine                      EntityType = "delivery_line"
	EntityTypeSalesOrder                        EntityType = "sales_order"
	EntityTypeLocation                          EntityType = "location"
	EntityTypeLocationType                      EntityType = "location_type"
	EntityTypeLot                               EntityType = "lot"
	EntityTypeEmailLog                          EntityType = "email_log"
	EntityTypeInventoryChangeLog                EntityType = "inventory_change_log"
	EntityTypeInvoice                           EntityType = "invoice"
	EntityTypeInvoiceSummary                    EntityType = "invoice_summary"
	EntityTypeInvoiceLine                       EntityType = "invoice_line"
	EntityTypeInvoiceAllocation                 EntityType = "invoice_allocation"
	EntityTypeInvoiceForPayment                 EntityType = "invoice_for_payment"
	EntityTypeShipment                          EntityType = "shipment"
	EntityTypeShipmentSummary                   EntityType = "shipment_summary"
	EntityTypeShipmentLine                      EntityType = "shipment_line"
	EntityTypeShippingCase                      EntityType = "shipping_case"
	EntityTypeShippingCaseLabelURL              EntityType = "shipping_case_label_url"
	EntityTypeSettlement                        EntityType = "settlement"
	EntityTypeSettlementSummary                 EntityType = "settlement_summary"
	EntityTypeRolePermission                    EntityType = "role_permission"
	EntityTypeRegistrationFlow                  EntityType = "registration_flow"
	EntityTypeRegistrationFlowOption            EntityType = "registration_flow_option"
	EntityTypeTransaction                       EntityType = "transaction"
	EntityTypeTransactionSummary                EntityType = "transaction_summary"
	EntityTypeTransactionMethod                 EntityType = "transaction_method"
	EntityTypeTransactionType                   EntityType = "transaction_type"
	EntityTypeTransactionAllocation             EntityType = "transaction_allocation"
	EntityTypeUsageItem                         EntityType = "usage_item"
	EntityTypeAgentTokenDetail                  EntityType = "agent_token_detail"
	EntityTypeAccountUsageResponse              EntityType = "account_usage_response"
	EntityTypeSubscriptionInfo                  EntityType = "subscription_info"
	EntityTypeBillingPortalSessionResponse      EntityType = "billing_portal_session_response"
	EntityTypeSwitchPlanResponse                EntityType = "switch_plan_response"
	EntityTypeEnsureBillingCustomerResponse     EntityType = "ensure_billing_customer_response"
	EntityTypeSpendingCapResponse               EntityType = "spending_cap_response"
	EntityTypeAgentSpendInfo                    EntityType = "agent_spend_info"
	EntityTypeWebhookResponse                   EntityType = "webhook_response"
	EntityTypeAddressSuggestion                 EntityType = "address_suggestion"
	EntityTypeAddressComponents                 EntityType = "address_components"
	EntityTypeAddressDetailsResult              EntityType = "address_details_result"
	EntityTypeValidatedAddress                  EntityType = "validated_address"
	EntityTypePlanLimit                         EntityType = "plan_limit"
	EntityTypePlanChangeProration               EntityType = "plan_change_proration"
	EntityTypePlanChangeLineItem                EntityType = "plan_change_line_item"
	EntityTypeSetupBillingResponse              EntityType = "setup_billing_response"
	EntityTypeConfirmPaymentResponse            EntityType = "confirm_payment_response"
	EntityTypeOAuthResponse                     EntityType = "oauth_response"
	EntityTypeOAuthStatusResponse               EntityType = "oauth_status_response"
	EntityTypeStripePublishableKey              EntityType = "stripe_publishable_key"
	EntityTypeStripeStatus                      EntityType = "stripe_status"
	EntityTypeHealthcheck                       EntityType = "healthcheck"
	EntityTypeAgentDefinitionConfig             EntityType = "agent_definition_config"
	EntityTypeTriggerConfig                     EntityType = "trigger_config"
	EntityTypeCustomerContactInfo               EntityType = "customer_contact_info"
	EntityTypeCustomerFreightPreferences        EntityType = "customer_freight_preferences"
	EntityTypeCustomerDefaults                  EntityType = "customer_defaults"
	EntityTypeCustomerNotificationPreferences   EntityType = "customer_notification_preferences"
	EntityTypeOrderDiscount                     EntityType = "order_discount"
	EntityTypeSalesOrderLine                    EntityType = "sales_order_line"
	EntityTypeSalesOrderType                    EntityType = "sales_order_type"
	EntityTypeSalesOrderStatus                  EntityType = "sales_order_status"
	EntityTypeMaterial                          EntityType = "material"
	EntityTypeSupplierMaterial                  EntityType = "supplier_material"
	EntityTypePart                              EntityType = "part"
	EntityTypePermissionGroup                   EntityType = "permission_group"
	EntityTypePermission                        EntityType = "permission"
	EntityTypePick                              EntityType = "pick"
	EntityTypePickLine                          EntityType = "pick_line"
	EntityTypeProductType                       EntityType = "product_type"
	EntityTypeProduction                        EntityType = "production"
	EntityTypeProductionFlow                    EntityType = "production_flow"
	EntityTypeMap                               EntityType = "map"
	EntityTypePurchaseOrder                     EntityType = "purchase_order"
	EntityTypePurchaseOrderLine                 EntityType = "purchase_order_line"
	EntityTypeSupplier                          EntityType = "supplier"
	EntityTypeSupplierSummary                   EntityType = "supplier_summary"
	EntityTypeReceivableEntry                   EntityType = "receivable_entry"
	EntityTypeReceivingOrder                    EntityType = "receiving_order"
	EntityTypeReceivingOrderLine                EntityType = "receiving_order_line"
	EntityTypeEmailContact                      EntityType = "email_contact"
	EntityTypeAllocationEntry                   EntityType = "allocation_entry"
	EntityTypeOpenCreditEntry                   EntityType = "open_credit_entry"
	EntityTypeVolumeDiscount                    EntityType = "volume_discount"
	EntityTypeVolumeDiscountTier                EntityType = "volume_discount_tier"
	EntityTypeAnalyzeDeliveriesResponse         EntityType = "analyze_deliveries_response"
	EntityTypeAnalyzeManufacturingResponse      EntityType = "analyze_manufacturing_response"
	EntityTypeAnalyzeManufacturingBatchResponse EntityType = "analyze_manufacturing_batch_response"
	EntityTypeAnalyzeQuarterlyOrdersResponse    EntityType = "analyze_quarterly_orders_response"
	EntityTypeAnalyzeNewCustomersResponse       EntityType = "analyze_new_customers_response"
	EntityTypeAnalyzeOeeResponse                EntityType = "analyze_oee_response"
	EntityTypeCatalogProductLine                EntityType = "catalog_product_line"
	EntityTypeCatalogCategory                   EntityType = "catalog_category"
	EntityTypeCatalogProduct                    EntityType = "catalog_product"
	EntityTypeCatalogProperty                   EntityType = "catalog_property"
	EntityTypeCatalogAttribute                  EntityType = "catalog_attribute"
	EntityTypeDcLocation                        EntityType = "dc_location"
	EntityTypeEdiRun                            EntityType = "edi_run"
	EntityTypeInventoryItem                     EntityType = "inventory_item"
	EntityTypeAnalyzeWeeksOfSalesResponse       EntityType = "analyze_weeks_of_sales_response"
	EntityTypeBulkReconcileItemsResponse        EntityType = "bulk_reconcile_items_response"
	EntityTypeSysProperty                       EntityType = "sys_property"
	EntityTypeSysPropertyType                   EntityType = "sys_property_type"
	EntityTypeSysPropertyValue                  EntityType = "sys_property_value"
	EntityTypeTerritory                         EntityType = "territory"
	EntityTypeTenancy                           EntityType = "tenancy"
	EntityTypeCheckoutSession                   EntityType = "checkout_session"
	EntityTypeEstimateRateResult                EntityType = "estimate_rate_result"
	EntityTypeRateShopOption                    EntityType = "rate_shop_option"
	EntityTypeRateShopResult                    EntityType = "rate_shop_result"
	EntityTypeOwner                             EntityType = "owner"
	EntityTypeMessage                           EntityType = "message"
	EntityTypeAccountPhotoUploadResult          EntityType = "account_photo_upload_result"
	EntityTypeUserPhotoUploadResult             EntityType = "user_photo_upload_result"
	EntityTypeUserPhotoURL                      EntityType = "user_photo_url"
	EntityTypeBatchLot                          EntityType = "batch_lot"
	EntityTypeCheckDuplicateResult              EntityType = "check_duplicate_result"
	EntityTypeItemTrendPoint                    EntityType = "item_trend_point"
	EntityTypePackPickResponse                  EntityType = "pack_pick_response"
	EntityTypePickShipmentsResponse             EntityType = "pick_shipments_response"
	EntityTypeTenancyPendingRegistration        EntityType = "tenancy_pending_registration"
	EntityTypeInvoiceAllocationEntry            EntityType = "invoice_allocation_entry"
	EntityTypeAllocationCustomer                EntityType = "allocation_customer"
	EntityTypeCheckoutSalesOrderResponse        EntityType = "checkout_sales_order_response"
	EntityTypeCreateProductionRunResponse       EntityType = "create_production_run_response"
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

// Location resource.
type Location struct {
	// Location ID.
	ID string `json:"id" api:"required"`
	// List represents a paginated list of resources.
	Children *ListLocation `json:"children" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "location".
	Object LocationObject `json:"object" api:"required"`
	// Location resource.
	Parent *Location `json:"parent" api:"required"`
	// Location type code.
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

// Machine within an account.
type Machine struct {
	// Machine ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Department resource.
	Department *Department `json:"department" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "machine".
	Object MachineObject `json:"object" api:"required"`
	// Serial number.
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

// Payment term resource.
type PaymentTerm struct {
	// Payment term ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "payment_term".
	Object PaymentTermObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Payment term status.
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

// Payment term status.
type PaymentTermStatus string

const (
	PaymentTermStatusActive   PaymentTermStatus = "active"
	PaymentTermStatusInactive PaymentTermStatus = "inactive"
)

// Production output of a production step.
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

// Production step with all nested data.
type ProductionStep struct {
	// Production step ID.
	ID string `json:"id" api:"required"`
	// Allowances as a decimal string.
	Allowances string `json:"allowances" api:"required" format:"decimal"`
	// List represents a paginated list of resources.
	Consumptions ListConsumption `json:"consumptions" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Department resource.
	Department *Department `json:"department" api:"required"`
	// List represents a paginated list of resources.
	InSteps *ListProductionStep `json:"in_steps" api:"required"`
	// Rate resource.
	LaborRate Rate `json:"labor_rate" api:"required"`
	// Rate resource.
	LaborTime Rate `json:"labor_time" api:"required"`
	// Leveling factor as a decimal string.
	LevelingFactor string `json:"leveling_factor" api:"required" format:"decimal"`
	// List represents a paginated list of resources.
	Machines *ListMachine `json:"machines" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_step".
	Object ProductionStepObject `json:"object" api:"required"`
	// List represents a paginated list of resources.
	OutSteps *ListProductionStep `json:"out_steps" api:"required"`
	// Rate resource.
	OverheadRate Rate `json:"overhead_rate" api:"required"`
	// Production output of a production step.
	Production ProductionOutput `json:"production" api:"required"`
	// Scanning station resource.
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

// QuantityInput represents a value with an associated unit for create/update
// requests.
//
// The properties UnitID, Value are required.
type QuantityInputParam struct {
	// The unit ID for the value.
	UnitID string `json:"unit_id" api:"required"`
	// The decimal value.
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

// Scanning station resource.
type ScanningStation struct {
	// Scanning station ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Department resource.
	Department *Department `json:"department" api:"required"`
	// Label size code.
	//
	// Any of "1x1", "1x3", "1x4", "2x4".
	LabelSize ScanningStationLabelSize `json:"label_size" api:"required"`
	// Label type code.
	//
	// Any of "tag", "traveler".
	LabelType ScanningStationLabelType `json:"label_type" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "scanning_station".
	Object ScanningStationObject `json:"object" api:"required"`
	// Operator requirement behavior for this station.
	//
	// Any of "none", "material_check".
	OperatorRequirement ScanningStationOperatorRequirement `json:"operator_requirement" api:"required"`
	// List represents a paginated list of resources.
	ProductionSteps *ListProductionStep `json:"production_steps" api:"required"`
	// Scanning station type.
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

// Label size code.
type ScanningStationLabelSize string

const (
	ScanningStationLabelSize1x1 ScanningStationLabelSize = "1x1"
	ScanningStationLabelSize1x3 ScanningStationLabelSize = "1x3"
	ScanningStationLabelSize1x4 ScanningStationLabelSize = "1x4"
	ScanningStationLabelSize2x4 ScanningStationLabelSize = "2x4"
)

// Label type code.
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

// Operator requirement behavior for this station.
type ScanningStationOperatorRequirement string

const (
	ScanningStationOperatorRequirementNone          ScanningStationOperatorRequirement = "none"
	ScanningStationOperatorRequirementMaterialCheck ScanningStationOperatorRequirement = "material_check"
)

// Scanning station type.
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
	// Customer portal visibility.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility ServiceLevelCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Default service level for the carrier.
	IsDefault bool `json:"is_default" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "service_level".
	Object ServiceLevelObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Service level token.
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

// Customer portal visibility.
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

// ShippingTerm resource.
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
	// Display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "shipping_term".
	Object ShippingTermObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Shipping term type.
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

// Shipping term type.
type ShippingTermType string

const (
	ShippingTermTypeFreeFreight        ShippingTermType = "free_freight"
	ShippingTermTypeFlatRateFreight    ShippingTermType = "flat_rate_freight"
	ShippingTermTypeCarrierRateFreight ShippingTermType = "carrier_rate_freight"
)

// Request to partially update a customer.
type UpdateCustomerRequestParam struct {
	// Bill-to address ID.
	BillToAddressID param.Opt[string] `json:"bill_to_address_id,omitzero"`
	// Carrier billing account number.
	CarrierBillingAccount param.Opt[string] `json:"carrier_billing_account,omitzero"`
	// The ID of the account user to assign as the default sales rep.
	DefaultSalesRepID param.Opt[string] `json:"default_sales_rep_id,omitzero"`
	// Default service level ID.
	DefaultServiceLevelID param.Opt[string] `json:"default_service_level_id,omitzero"`
	// Email address. Send null to clear.
	Email param.Opt[string] `json:"email,omitzero"`
	// Note.
	Note param.Opt[string] `json:"note,omitzero"`
	// Phone number. Send null to clear.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Ship-to address ID.
	ShipToAddressID param.Opt[string] `json:"ship_to_address_id,omitzero"`
	// Website URL. Send null to clear.
	URL param.Opt[string] `json:"url,omitzero"`
	// Customer type group ID.
	CustomerTypeGroupID param.Opt[string] `json:"customer_type_group_id,omitzero"`
	// Default carrier ID.
	DefaultCarrierID param.Opt[string] `json:"default_carrier_id,omitzero"`
	// Default payment term ID.
	DefaultPaymentTermID param.Opt[string] `json:"default_payment_term_id,omitzero"`
	// Default shipping term ID.
	DefaultShippingTermID param.Opt[string] `json:"default_shipping_term_id,omitzero"`
	// Customer name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Customer number.
	Number param.Opt[string] `json:"number,omitzero"`
	// Carrier billing type.
	//
	// Any of "sender", "third_party".
	CarrierBillingType UpdateCustomerRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// Commission policy.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy UpdateCustomerRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// QuantityInput represents a value with an associated unit for create/update
	// requests.
	CreditLimit QuantityInputParam `json:"credit_limit,omitzero"`
	// Price group IDs. Replaces all existing price groups when provided.
	CustomerPriceGroupIDs []string `json:"customer_price_group_ids,omitzero"`
	// Default priority code.
	//
	// Any of "low", "normal", "high".
	DefaultPriority UpdateCustomerRequestDefaultPriority `json:"default_priority,omitzero"`
	// EDI status.
	//
	// Any of "enabled", "disabled".
	EdiStatus UpdateCustomerRequestEdiStatus `json:"edi_status,omitzero"`
	// Freight policy.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy UpdateCustomerRequestFreightPolicy `json:"freight_policy,omitzero"`
	// Account status code.
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

// Carrier billing type.
type UpdateCustomerRequestCarrierBillingType string

const (
	UpdateCustomerRequestCarrierBillingTypeSender     UpdateCustomerRequestCarrierBillingType = "sender"
	UpdateCustomerRequestCarrierBillingTypeThirdParty UpdateCustomerRequestCarrierBillingType = "third_party"
)

// Commission policy.
type UpdateCustomerRequestCommissionPolicy string

const (
	UpdateCustomerRequestCommissionPolicyCommissionApplied UpdateCustomerRequestCommissionPolicy = "commission_applied"
	UpdateCustomerRequestCommissionPolicyCommissionExempt  UpdateCustomerRequestCommissionPolicy = "commission_exempt"
)

// Default priority code.
type UpdateCustomerRequestDefaultPriority string

const (
	UpdateCustomerRequestDefaultPriorityLow    UpdateCustomerRequestDefaultPriority = "low"
	UpdateCustomerRequestDefaultPriorityNormal UpdateCustomerRequestDefaultPriority = "normal"
	UpdateCustomerRequestDefaultPriorityHigh   UpdateCustomerRequestDefaultPriority = "high"
)

// EDI status.
type UpdateCustomerRequestEdiStatus string

const (
	UpdateCustomerRequestEdiStatusEnabled  UpdateCustomerRequestEdiStatus = "enabled"
	UpdateCustomerRequestEdiStatusDisabled UpdateCustomerRequestEdiStatus = "disabled"
)

// Freight policy.
type UpdateCustomerRequestFreightPolicy string

const (
	UpdateCustomerRequestFreightPolicyFreeFreight   UpdateCustomerRequestFreightPolicy = "free_freight"
	UpdateCustomerRequestFreightPolicyBilledFreight UpdateCustomerRequestFreightPolicy = "billed_freight"
)

// Account status code.
type UpdateCustomerRequestStatus string

const (
	UpdateCustomerRequestStatusNormal       UpdateCustomerRequestStatus = "normal"
	UpdateCustomerRequestStatusPreferred    UpdateCustomerRequestStatus = "preferred"
	UpdateCustomerRequestStatusHoldShipment UpdateCustomerRequestStatus = "hold_shipment"
	UpdateCustomerRequestStatusHoldAll      UpdateCustomerRequestStatus = "hold_all"
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
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit".
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
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit".
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
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit".
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
	// Filter by city.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Cursor token used to retrieve the next or previous page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by end date (created before).
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Maximum number of results per page (default: 100, max: 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by postal code.
	PostalCode param.Opt[string] `query:"postal_code,omitzero" json:"-"`
	// Search query used to filter results.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter by start date (created after).
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	// Filter by state.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	// Filter by carrier IDs.
	CarrierIDs []string `query:"carrier_ids,omitzero" json:"-"`
	// Filter by commission status codes.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionStatusCodes []string `query:"commission_status_codes,omitzero" json:"-"`
	// Filter by customer group IDs.
	CustomerGroupIDs []string `query:"customer_group_ids,omitzero" json:"-"`
	// Filter by freight status codes.
	//
	// Any of "free_freight", "billed_freight".
	FreightStatusCodes []string `query:"freight_status_codes,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "bill_to_address", "ship_to_address", "type", "parent_account",
	// "freight_preferences.carrier", "freight_preferences.service_level",
	// "defaults.payment_term", "defaults.shipping_term", "defaults.sales_rep",
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by whether the customer has child accounts.
	//
	// Any of "parent", "non_parent".
	ParentAccountStatus SaleCustomerListParamsParentAccountStatus `query:"parent_account_status,omitzero" json:"-"`
	// Filter by payment term IDs.
	PaymentTermIDs []string `query:"payment_term_ids,omitzero" json:"-"`
	// Filter by pricing group IDs.
	PricingGroupIDs []string `query:"pricing_group_ids,omitzero" json:"-"`
	// Filter by sales rep IDs.
	SalesRepIDs []string `query:"sales_rep_ids,omitzero" json:"-"`
	// Filter by service level IDs.
	ServiceLevelIDs []string `query:"service_level_ids,omitzero" json:"-"`
	// Filter by shipping term IDs.
	ShippingTermIDs []string `query:"shipping_term_ids,omitzero" json:"-"`
	// Filter by status codes.
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
