// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// Unified free-text search across resource types, returning lightweight entity
// references.
//
// CoreService contains methods and other services that help with interacting with
// the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreService] method instead.
type CoreService struct {
	options []option.RequestOption
	// List and manage sandbox environments.
	Sandboxes CoreSandboxService
	// List and retrieve request logs.
	RequestLogs CoreRequestLogService
	// List and retrieve audit events.
	AuditEvents CoreAuditEventService
	// Autocomplete, look up details, and validate addresses.
	Addresses CoreAddressService
	// View email logs for accounts.
	EmailLogs CoreEmailLogService
}

// NewCoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCoreService(opts ...option.RequestOption) (r CoreService) {
	r = CoreService{}
	r.options = opts
	r.Sandboxes = NewCoreSandboxService(opts...)
	r.RequestLogs = NewCoreRequestLogService(opts...)
	r.AuditEvents = NewCoreAuditEventService(opts...)
	r.Addresses = NewCoreAddressService(opts...)
	r.EmailLogs = NewCoreEmailLogService(opts...)
	return
}

// Search returns lightweight `entity` references matching the query across the
// resource types the caller can read.
//
// This endpoint requires the permissions: `sales_orders:read`,
// `purchase_orders:read`, `invoices:read`, `customers:read`, `items:read`,
// `shipments:read`, `messaging:read`, `agents:read`.
func (r *CoreService) GetSearch(ctx context.Context, query CoreGetSearchParams, opts ...option.RequestOption) (res *ListEntity, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

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
	// The resource kind that this entity references, as an object-type value (e.g.
	// `user`, `account`).
	//
	// Unlike `object` — which is always `entity` — this names the underlying resource
	// the `id` points to.
	//
	// Any of "account", "actor", "entity", "record", "freight", "sales_order_totals",
	// "sales_order_related", "order_contact", "user", "address", "api_key",
	// "created_api_key", "refresh_token", "list", "sandbox", "registration_session",
	// "pricing_plan", "account_plan", "plan_change", "enterprise_inquiry",
	// "request_log", "audit_event", "audit_field_change", "role", "unit",
	// "account_affiliation", "agent_definition", "available_tool",
	// "agent_definition_tool", "agent_account_status", "agent_run", "agent_action",
	// "agent_run_step", "agent_token_usage", "agent_memory", "notification",
	// "notification_unread_count", "notification_send_result",
	// "notification_unread_summary", "announcement", "conversation",
	// "conversation_participant", "read_cursor", "chat_message",
	// "notification_unread_summary_account", "messaging_block",
	// "notification_preference", "message_attachment", "attachment_upload_target",
	// "scheduled_message", "messaging_contact", "message_report", "tool_group",
	// "model", "payment_term", "shipping_term", "quantity", "account_group",
	// "support_route", "support_availability", "account_status", "geolocation",
	// "account_user", "department", "account_integration", "account_price",
	// "product_line", "item_category", "attribute", "rate",
	// "account_group_product_line_access", "sales_target", "adjustment_type",
	// "account_branding", "account_portal", "account_logo_url", "public_account",
	// "property", "carrier", "service_level", "item", "item_inventory", "product",
	// "batch", "batch_flow_node", "scanning_consumption", "open_batch_summary",
	// "scanning_production_step_info", "scanning_station", "production_step",
	// "production_run", "machine", "child_account", "unit_group", "unit_group_unit",
	// "consumption", "customer_product_line_access", "customer",
	// "frequently_ordered_product", "priority", "delivery", "delivery_line",
	// "sales_order", "location", "location_type", "lot", "email_log", "email_domain",
	// "email_inbox", "portal_domain", "dns_record", "inventory_change_log", "invoice",
	// "invoice_summary", "invoice_line", "invoice_allocation", "invoice_for_payment",
	// "shipment", "shipment_summary", "shipment_line", "shipping_case",
	// "shipping_case_label_url", "settlement", "settlement_summary",
	// "role_permission", "registration_flow", "registration_flow_option",
	// "transaction", "transaction_summary", "transaction_method", "transaction_type",
	// "transaction_allocation", "usage_item", "account_usage_response",
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
	// "created_by", "message", "account_photo_upload_result",
	// "user_photo_upload_result", "user_photo_url", "batch_lot",
	// "check_duplicate_result", "item_trend_point", "pack_pick_response",
	// "pick_shipments_response", "tenancy_pending_registration",
	// "invoice_allocation_entry", "allocation_customer",
	// "checkout_sales_order_response", "create_production_run_response",
	// "sales_order_price_quote", "hubspot_sync_job", "hubspot_sync_report",
	// "hubspot_company_review", "hubspot_company_candidate", "contact_match",
	// "reply_draft", "conversation_link", "messaging_group", "messaging_group_member".
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

// The resource kind that this entity references, as an object-type value (e.g.
// `user`, `account`).
//
// Unlike `object` — which is always `entity` — this names the underlying resource
// the `id` points to.
type EntityType string

const (
	EntityTypeAccount                           EntityType = "account"
	EntityTypeActor                             EntityType = "actor"
	EntityTypeEntity                            EntityType = "entity"
	EntityTypeRecord                            EntityType = "record"
	EntityTypeFreight                           EntityType = "freight"
	EntityTypeSalesOrderTotals                  EntityType = "sales_order_totals"
	EntityTypeSalesOrderRelated                 EntityType = "sales_order_related"
	EntityTypeOrderContact                      EntityType = "order_contact"
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
	EntityTypeNotification                      EntityType = "notification"
	EntityTypeNotificationUnreadCount           EntityType = "notification_unread_count"
	EntityTypeNotificationSendResult            EntityType = "notification_send_result"
	EntityTypeNotificationUnreadSummary         EntityType = "notification_unread_summary"
	EntityTypeAnnouncement                      EntityType = "announcement"
	EntityTypeConversation                      EntityType = "conversation"
	EntityTypeConversationParticipant           EntityType = "conversation_participant"
	EntityTypeReadCursor                        EntityType = "read_cursor"
	EntityTypeChatMessage                       EntityType = "chat_message"
	EntityTypeNotificationUnreadSummaryAccount  EntityType = "notification_unread_summary_account"
	EntityTypeMessagingBlock                    EntityType = "messaging_block"
	EntityTypeNotificationPreference            EntityType = "notification_preference"
	EntityTypeMessageAttachment                 EntityType = "message_attachment"
	EntityTypeAttachmentUploadTarget            EntityType = "attachment_upload_target"
	EntityTypeScheduledMessage                  EntityType = "scheduled_message"
	EntityTypeMessagingContact                  EntityType = "messaging_contact"
	EntityTypeMessageReport                     EntityType = "message_report"
	EntityTypeToolGroup                         EntityType = "tool_group"
	EntityTypeModel                             EntityType = "model"
	EntityTypePaymentTerm                       EntityType = "payment_term"
	EntityTypeShippingTerm                      EntityType = "shipping_term"
	EntityTypeQuantity                          EntityType = "quantity"
	EntityTypeAccountGroup                      EntityType = "account_group"
	EntityTypeSupportRoute                      EntityType = "support_route"
	EntityTypeSupportAvailability               EntityType = "support_availability"
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
	EntityTypeEmailDomain                       EntityType = "email_domain"
	EntityTypeEmailInbox                        EntityType = "email_inbox"
	EntityTypePortalDomain                      EntityType = "portal_domain"
	EntityTypeDNSRecord                         EntityType = "dns_record"
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
	EntityTypeCreatedBy                         EntityType = "created_by"
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
	EntityTypeSalesOrderPriceQuote              EntityType = "sales_order_price_quote"
	EntityTypeHubspotSyncJob                    EntityType = "hubspot_sync_job"
	EntityTypeHubspotSyncReport                 EntityType = "hubspot_sync_report"
	EntityTypeHubspotCompanyReview              EntityType = "hubspot_company_review"
	EntityTypeHubspotCompanyCandidate           EntityType = "hubspot_company_candidate"
	EntityTypeContactMatch                      EntityType = "contact_match"
	EntityTypeReplyDraft                        EntityType = "reply_draft"
	EntityTypeConversationLink                  EntityType = "conversation_link"
	EntityTypeMessagingGroup                    EntityType = "messaging_group"
	EntityTypeMessagingGroupMember              EntityType = "messaging_group_member"
)

// List represents a paginated list of resources.
type ListEntity struct {
	// Resources in this page.
	Data []Entity `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListEntityObject `json:"object" api:"required"`
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
func (r ListEntity) RawJSON() string { return r.JSON.raw }
func (r *ListEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListEntityObject string

const (
	ListEntityObjectList ListEntityObject = "list"
)

type CoreGetSearchParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Restrict the search to a single customer by their account ID.
	//
	// When set, only resource types that are safe to expose to a customer are searched
	// (their sales orders, invoices, and shipments), and results are limited to
	// records belonging to that customer.
	Customer param.Opt[string] `query:"customer,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter the search to specific resource types.
	//
	// Omit to search every supported type the caller can read.
	//
	// Any of "account", "actor", "entity", "record", "freight", "sales_order_totals",
	// "sales_order_related", "order_contact", "user", "address", "api_key",
	// "created_api_key", "refresh_token", "list", "sandbox", "registration_session",
	// "pricing_plan", "account_plan", "plan_change", "enterprise_inquiry",
	// "request_log", "audit_event", "audit_field_change", "role", "unit",
	// "account_affiliation", "agent_definition", "available_tool",
	// "agent_definition_tool", "agent_account_status", "agent_run", "agent_action",
	// "agent_run_step", "agent_token_usage", "agent_memory", "notification",
	// "notification_unread_count", "notification_send_result",
	// "notification_unread_summary", "announcement", "conversation",
	// "conversation_participant", "read_cursor", "chat_message",
	// "notification_unread_summary_account", "messaging_block",
	// "notification_preference", "message_attachment", "attachment_upload_target",
	// "scheduled_message", "messaging_contact", "message_report", "tool_group",
	// "model", "payment_term", "shipping_term", "quantity", "account_group",
	// "support_route", "support_availability", "account_status", "geolocation",
	// "account_user", "department", "account_integration", "account_price",
	// "product_line", "item_category", "attribute", "rate",
	// "account_group_product_line_access", "sales_target", "adjustment_type",
	// "account_branding", "account_portal", "account_logo_url", "public_account",
	// "property", "carrier", "service_level", "item", "item_inventory", "product",
	// "batch", "batch_flow_node", "scanning_consumption", "open_batch_summary",
	// "scanning_production_step_info", "scanning_station", "production_step",
	// "production_run", "machine", "child_account", "unit_group", "unit_group_unit",
	// "consumption", "customer_product_line_access", "customer",
	// "frequently_ordered_product", "priority", "delivery", "delivery_line",
	// "sales_order", "location", "location_type", "lot", "email_log", "email_domain",
	// "email_inbox", "portal_domain", "dns_record", "inventory_change_log", "invoice",
	// "invoice_summary", "invoice_line", "invoice_allocation", "invoice_for_payment",
	// "shipment", "shipment_summary", "shipment_line", "shipping_case",
	// "shipping_case_label_url", "settlement", "settlement_summary",
	// "role_permission", "registration_flow", "registration_flow_option",
	// "transaction", "transaction_summary", "transaction_method", "transaction_type",
	// "transaction_allocation", "usage_item", "account_usage_response",
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
	// "created_by", "message", "account_photo_upload_result",
	// "user_photo_upload_result", "user_photo_url", "batch_lot",
	// "check_duplicate_result", "item_trend_point", "pack_pick_response",
	// "pick_shipments_response", "tenancy_pending_registration",
	// "invoice_allocation_entry", "allocation_customer",
	// "checkout_sales_order_response", "create_production_run_response",
	// "sales_order_price_quote", "hubspot_sync_job", "hubspot_sync_report",
	// "hubspot_company_review", "hubspot_company_candidate", "contact_match",
	// "reply_draft", "conversation_link", "messaging_group", "messaging_group_member".
	Types []string `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreGetSearchParams]'s query parameters as `url.Values`.
func (r CoreGetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
