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

// List and retrieve audit events.
//
// CoreAuditEventService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreAuditEventService] method instead.
type CoreAuditEventService struct {
	options []option.RequestOption
}

// NewCoreAuditEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreAuditEventService(opts ...option.RequestOption) (r CoreAuditEventService) {
	r = CoreAuditEventService{}
	r.options = opts
	return
}

// Returns an audit event by ID.
//
// This endpoint requires the permission: `audit_events:read`.
func (r *CoreAuditEventService) Get(ctx context.Context, id string, query CoreAuditEventGetParams, opts ...option.RequestOption) (res *AuditEvent, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/audit-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of audit events for the current account.
//
// This endpoint requires the permission: `audit_events:read`.
func (r *CoreAuditEventService) List(ctx context.Context, query CoreAuditEventListParams, opts ...option.RequestOption) (res *ListAuditEvent, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/audit-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the full set of resource types that may appear on audit events.
//
// Values are plain strings, suitable for the `resource_types` filter when listing
// audit events.
//
// This endpoint requires the permission: `audit_events:read`.
func (r *CoreAuditEventService) GetResourceTypes(ctx context.Context, opts ...option.RequestOption) (res *ListObjectType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/audit-events/resource-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Immutable audit event record.
//
// Captures the actor, changed resource, and timestamp.
type AuditEvent struct {
	// Audit event ID.
	ID string `json:"id" api:"required"`
	// A customer account, including its branding and customer portal sub-resources.
	Account Account `json:"account" api:"required"`
	// The type of action this event records.
	//
	//   - `create`: the resource was created.
	//   - `update`: one or more fields were changed.
	//   - `delete`: the resource was deleted.
	//   - `restore`: a previously deleted resource was restored.
	//   - `archive`: the resource was archived.
	//   - `approve`: a human approved a gated action, such as allowing a review-gated
	//     agent tool to run.
	//   - `deny`: a human denied a gated action, such as rejecting a review-gated agent
	//     tool.
	//
	// Any of "create", "update", "delete", "restore", "archive", "approve", "deny".
	Action AuditEventAction `json:"action" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// List represents a paginated list of resources.
	Changes ListAuditFieldChange `json:"changes" api:"required"`
	// When the audit event record was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Idempotency key of the originating request.
	IdempotencyKey string `json:"idempotency_key" api:"required"`
	// Arbitrary JSON metadata for the mutation (e.g. reason, source, tags). Encoded as
	// a JSON value (object, array, string, number, boolean, or null), not a
	// JSON-encoded string.
	Metadata any `json:"metadata" api:"required"`
	// Resource type identifier.
	//
	// Any of "audit_event".
	Object AuditEventObject `json:"object" api:"required"`
	// When the audited mutation occurred.
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// A log of a single API request, capturing its route, outcome, latency, and actor.
	Request RequestLog `json:"request" api:"required"`
	// Audited resource ID.
	ResourceID string `json:"resource_id" api:"required"`
	// Resource type of the audited entity.
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
	// "conversation_participant", "chat_message",
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
	// "email_inbox", "inventory_change_log", "invoice", "invoice_summary",
	// "invoice_line", "invoice_allocation", "invoice_for_payment", "shipment",
	// "shipment_summary", "shipment_line", "shipping_case", "shipping_case_label_url",
	// "settlement", "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "account_usage_response", "subscription_info",
	// "billing_portal_session_response", "switch_plan_response",
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
	ResourceType AuditEventResourceType `json:"resource_type" api:"required"`
	// Originating client IP address.
	SourceIP string `json:"source_ip" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Account        respjson.Field
		Action         respjson.Field
		Actor          respjson.Field
		Changes        respjson.Field
		CreatedAt      respjson.Field
		IdempotencyKey respjson.Field
		Metadata       respjson.Field
		Object         respjson.Field
		OccurredAt     respjson.Field
		Request        respjson.Field
		ResourceID     respjson.Field
		ResourceType   respjson.Field
		SourceIP       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEvent) RawJSON() string { return r.JSON.raw }
func (r *AuditEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of action this event records.
//
//   - `create`: the resource was created.
//   - `update`: one or more fields were changed.
//   - `delete`: the resource was deleted.
//   - `restore`: a previously deleted resource was restored.
//   - `archive`: the resource was archived.
//   - `approve`: a human approved a gated action, such as allowing a review-gated
//     agent tool to run.
//   - `deny`: a human denied a gated action, such as rejecting a review-gated agent
//     tool.
type AuditEventAction string

const (
	AuditEventActionCreate  AuditEventAction = "create"
	AuditEventActionUpdate  AuditEventAction = "update"
	AuditEventActionDelete  AuditEventAction = "delete"
	AuditEventActionRestore AuditEventAction = "restore"
	AuditEventActionArchive AuditEventAction = "archive"
	AuditEventActionApprove AuditEventAction = "approve"
	AuditEventActionDeny    AuditEventAction = "deny"
)

// Resource type identifier.
type AuditEventObject string

const (
	AuditEventObjectAuditEvent AuditEventObject = "audit_event"
)

// Resource type of the audited entity.
type AuditEventResourceType string

const (
	AuditEventResourceTypeAccount                           AuditEventResourceType = "account"
	AuditEventResourceTypeActor                             AuditEventResourceType = "actor"
	AuditEventResourceTypeEntity                            AuditEventResourceType = "entity"
	AuditEventResourceTypeRecord                            AuditEventResourceType = "record"
	AuditEventResourceTypeFreight                           AuditEventResourceType = "freight"
	AuditEventResourceTypeSalesOrderTotals                  AuditEventResourceType = "sales_order_totals"
	AuditEventResourceTypeSalesOrderRelated                 AuditEventResourceType = "sales_order_related"
	AuditEventResourceTypeOrderContact                      AuditEventResourceType = "order_contact"
	AuditEventResourceTypeUser                              AuditEventResourceType = "user"
	AuditEventResourceTypeAddress                           AuditEventResourceType = "address"
	AuditEventResourceTypeAPIKey                            AuditEventResourceType = "api_key"
	AuditEventResourceTypeCreatedAPIKey                     AuditEventResourceType = "created_api_key"
	AuditEventResourceTypeRefreshToken                      AuditEventResourceType = "refresh_token"
	AuditEventResourceTypeList                              AuditEventResourceType = "list"
	AuditEventResourceTypeSandbox                           AuditEventResourceType = "sandbox"
	AuditEventResourceTypeRegistrationSession               AuditEventResourceType = "registration_session"
	AuditEventResourceTypePricingPlan                       AuditEventResourceType = "pricing_plan"
	AuditEventResourceTypeAccountPlan                       AuditEventResourceType = "account_plan"
	AuditEventResourceTypePlanChange                        AuditEventResourceType = "plan_change"
	AuditEventResourceTypeEnterpriseInquiry                 AuditEventResourceType = "enterprise_inquiry"
	AuditEventResourceTypeRequestLog                        AuditEventResourceType = "request_log"
	AuditEventResourceTypeAuditEvent                        AuditEventResourceType = "audit_event"
	AuditEventResourceTypeAuditFieldChange                  AuditEventResourceType = "audit_field_change"
	AuditEventResourceTypeRole                              AuditEventResourceType = "role"
	AuditEventResourceTypeUnit                              AuditEventResourceType = "unit"
	AuditEventResourceTypeAccountAffiliation                AuditEventResourceType = "account_affiliation"
	AuditEventResourceTypeAgentDefinition                   AuditEventResourceType = "agent_definition"
	AuditEventResourceTypeAvailableTool                     AuditEventResourceType = "available_tool"
	AuditEventResourceTypeAgentDefinitionTool               AuditEventResourceType = "agent_definition_tool"
	AuditEventResourceTypeAgentAccountStatus                AuditEventResourceType = "agent_account_status"
	AuditEventResourceTypeAgentRun                          AuditEventResourceType = "agent_run"
	AuditEventResourceTypeAgentAction                       AuditEventResourceType = "agent_action"
	AuditEventResourceTypeAgentRunStep                      AuditEventResourceType = "agent_run_step"
	AuditEventResourceTypeAgentTokenUsage                   AuditEventResourceType = "agent_token_usage"
	AuditEventResourceTypeAgentMemory                       AuditEventResourceType = "agent_memory"
	AuditEventResourceTypeNotification                      AuditEventResourceType = "notification"
	AuditEventResourceTypeNotificationUnreadCount           AuditEventResourceType = "notification_unread_count"
	AuditEventResourceTypeNotificationSendResult            AuditEventResourceType = "notification_send_result"
	AuditEventResourceTypeNotificationUnreadSummary         AuditEventResourceType = "notification_unread_summary"
	AuditEventResourceTypeAnnouncement                      AuditEventResourceType = "announcement"
	AuditEventResourceTypeConversation                      AuditEventResourceType = "conversation"
	AuditEventResourceTypeConversationParticipant           AuditEventResourceType = "conversation_participant"
	AuditEventResourceTypeChatMessage                       AuditEventResourceType = "chat_message"
	AuditEventResourceTypeNotificationUnreadSummaryAccount  AuditEventResourceType = "notification_unread_summary_account"
	AuditEventResourceTypeMessagingBlock                    AuditEventResourceType = "messaging_block"
	AuditEventResourceTypeNotificationPreference            AuditEventResourceType = "notification_preference"
	AuditEventResourceTypeMessageAttachment                 AuditEventResourceType = "message_attachment"
	AuditEventResourceTypeAttachmentUploadTarget            AuditEventResourceType = "attachment_upload_target"
	AuditEventResourceTypeScheduledMessage                  AuditEventResourceType = "scheduled_message"
	AuditEventResourceTypeMessagingContact                  AuditEventResourceType = "messaging_contact"
	AuditEventResourceTypeMessageReport                     AuditEventResourceType = "message_report"
	AuditEventResourceTypeToolGroup                         AuditEventResourceType = "tool_group"
	AuditEventResourceTypeModel                             AuditEventResourceType = "model"
	AuditEventResourceTypePaymentTerm                       AuditEventResourceType = "payment_term"
	AuditEventResourceTypeShippingTerm                      AuditEventResourceType = "shipping_term"
	AuditEventResourceTypeQuantity                          AuditEventResourceType = "quantity"
	AuditEventResourceTypeAccountGroup                      AuditEventResourceType = "account_group"
	AuditEventResourceTypeSupportRoute                      AuditEventResourceType = "support_route"
	AuditEventResourceTypeSupportAvailability               AuditEventResourceType = "support_availability"
	AuditEventResourceTypeAccountStatus                     AuditEventResourceType = "account_status"
	AuditEventResourceTypeGeolocation                       AuditEventResourceType = "geolocation"
	AuditEventResourceTypeAccountUser                       AuditEventResourceType = "account_user"
	AuditEventResourceTypeDepartment                        AuditEventResourceType = "department"
	AuditEventResourceTypeAccountIntegration                AuditEventResourceType = "account_integration"
	AuditEventResourceTypeAccountPrice                      AuditEventResourceType = "account_price"
	AuditEventResourceTypeProductLine                       AuditEventResourceType = "product_line"
	AuditEventResourceTypeItemCategory                      AuditEventResourceType = "item_category"
	AuditEventResourceTypeAttribute                         AuditEventResourceType = "attribute"
	AuditEventResourceTypeRate                              AuditEventResourceType = "rate"
	AuditEventResourceTypeAccountGroupProductLineAccess     AuditEventResourceType = "account_group_product_line_access"
	AuditEventResourceTypeSalesTarget                       AuditEventResourceType = "sales_target"
	AuditEventResourceTypeAdjustmentType                    AuditEventResourceType = "adjustment_type"
	AuditEventResourceTypeAccountBranding                   AuditEventResourceType = "account_branding"
	AuditEventResourceTypeAccountPortal                     AuditEventResourceType = "account_portal"
	AuditEventResourceTypeAccountLogoURL                    AuditEventResourceType = "account_logo_url"
	AuditEventResourceTypePublicAccount                     AuditEventResourceType = "public_account"
	AuditEventResourceTypeProperty                          AuditEventResourceType = "property"
	AuditEventResourceTypeCarrier                           AuditEventResourceType = "carrier"
	AuditEventResourceTypeServiceLevel                      AuditEventResourceType = "service_level"
	AuditEventResourceTypeItem                              AuditEventResourceType = "item"
	AuditEventResourceTypeItemInventory                     AuditEventResourceType = "item_inventory"
	AuditEventResourceTypeProduct                           AuditEventResourceType = "product"
	AuditEventResourceTypeBatch                             AuditEventResourceType = "batch"
	AuditEventResourceTypeBatchFlowNode                     AuditEventResourceType = "batch_flow_node"
	AuditEventResourceTypeScanningConsumption               AuditEventResourceType = "scanning_consumption"
	AuditEventResourceTypeOpenBatchSummary                  AuditEventResourceType = "open_batch_summary"
	AuditEventResourceTypeScanningProductionStepInfo        AuditEventResourceType = "scanning_production_step_info"
	AuditEventResourceTypeScanningStation                   AuditEventResourceType = "scanning_station"
	AuditEventResourceTypeProductionStep                    AuditEventResourceType = "production_step"
	AuditEventResourceTypeProductionRun                     AuditEventResourceType = "production_run"
	AuditEventResourceTypeMachine                           AuditEventResourceType = "machine"
	AuditEventResourceTypeChildAccount                      AuditEventResourceType = "child_account"
	AuditEventResourceTypeUnitGroup                         AuditEventResourceType = "unit_group"
	AuditEventResourceTypeUnitGroupUnit                     AuditEventResourceType = "unit_group_unit"
	AuditEventResourceTypeConsumption                       AuditEventResourceType = "consumption"
	AuditEventResourceTypeCustomerProductLineAccess         AuditEventResourceType = "customer_product_line_access"
	AuditEventResourceTypeCustomer                          AuditEventResourceType = "customer"
	AuditEventResourceTypeFrequentlyOrderedProduct          AuditEventResourceType = "frequently_ordered_product"
	AuditEventResourceTypePriority                          AuditEventResourceType = "priority"
	AuditEventResourceTypeDelivery                          AuditEventResourceType = "delivery"
	AuditEventResourceTypeDeliveryLine                      AuditEventResourceType = "delivery_line"
	AuditEventResourceTypeSalesOrder                        AuditEventResourceType = "sales_order"
	AuditEventResourceTypeLocation                          AuditEventResourceType = "location"
	AuditEventResourceTypeLocationType                      AuditEventResourceType = "location_type"
	AuditEventResourceTypeLot                               AuditEventResourceType = "lot"
	AuditEventResourceTypeEmailLog                          AuditEventResourceType = "email_log"
	AuditEventResourceTypeEmailDomain                       AuditEventResourceType = "email_domain"
	AuditEventResourceTypeEmailInbox                        AuditEventResourceType = "email_inbox"
	AuditEventResourceTypeInventoryChangeLog                AuditEventResourceType = "inventory_change_log"
	AuditEventResourceTypeInvoice                           AuditEventResourceType = "invoice"
	AuditEventResourceTypeInvoiceSummary                    AuditEventResourceType = "invoice_summary"
	AuditEventResourceTypeInvoiceLine                       AuditEventResourceType = "invoice_line"
	AuditEventResourceTypeInvoiceAllocation                 AuditEventResourceType = "invoice_allocation"
	AuditEventResourceTypeInvoiceForPayment                 AuditEventResourceType = "invoice_for_payment"
	AuditEventResourceTypeShipment                          AuditEventResourceType = "shipment"
	AuditEventResourceTypeShipmentSummary                   AuditEventResourceType = "shipment_summary"
	AuditEventResourceTypeShipmentLine                      AuditEventResourceType = "shipment_line"
	AuditEventResourceTypeShippingCase                      AuditEventResourceType = "shipping_case"
	AuditEventResourceTypeShippingCaseLabelURL              AuditEventResourceType = "shipping_case_label_url"
	AuditEventResourceTypeSettlement                        AuditEventResourceType = "settlement"
	AuditEventResourceTypeSettlementSummary                 AuditEventResourceType = "settlement_summary"
	AuditEventResourceTypeRolePermission                    AuditEventResourceType = "role_permission"
	AuditEventResourceTypeRegistrationFlow                  AuditEventResourceType = "registration_flow"
	AuditEventResourceTypeRegistrationFlowOption            AuditEventResourceType = "registration_flow_option"
	AuditEventResourceTypeTransaction                       AuditEventResourceType = "transaction"
	AuditEventResourceTypeTransactionSummary                AuditEventResourceType = "transaction_summary"
	AuditEventResourceTypeTransactionMethod                 AuditEventResourceType = "transaction_method"
	AuditEventResourceTypeTransactionType                   AuditEventResourceType = "transaction_type"
	AuditEventResourceTypeTransactionAllocation             AuditEventResourceType = "transaction_allocation"
	AuditEventResourceTypeUsageItem                         AuditEventResourceType = "usage_item"
	AuditEventResourceTypeAccountUsageResponse              AuditEventResourceType = "account_usage_response"
	AuditEventResourceTypeSubscriptionInfo                  AuditEventResourceType = "subscription_info"
	AuditEventResourceTypeBillingPortalSessionResponse      AuditEventResourceType = "billing_portal_session_response"
	AuditEventResourceTypeSwitchPlanResponse                AuditEventResourceType = "switch_plan_response"
	AuditEventResourceTypeEnsureBillingCustomerResponse     AuditEventResourceType = "ensure_billing_customer_response"
	AuditEventResourceTypeSpendingCapResponse               AuditEventResourceType = "spending_cap_response"
	AuditEventResourceTypeAgentSpendInfo                    AuditEventResourceType = "agent_spend_info"
	AuditEventResourceTypeWebhookResponse                   AuditEventResourceType = "webhook_response"
	AuditEventResourceTypeAddressSuggestion                 AuditEventResourceType = "address_suggestion"
	AuditEventResourceTypeAddressComponents                 AuditEventResourceType = "address_components"
	AuditEventResourceTypeAddressDetailsResult              AuditEventResourceType = "address_details_result"
	AuditEventResourceTypeValidatedAddress                  AuditEventResourceType = "validated_address"
	AuditEventResourceTypePlanLimit                         AuditEventResourceType = "plan_limit"
	AuditEventResourceTypePlanChangeProration               AuditEventResourceType = "plan_change_proration"
	AuditEventResourceTypePlanChangeLineItem                AuditEventResourceType = "plan_change_line_item"
	AuditEventResourceTypeSetupBillingResponse              AuditEventResourceType = "setup_billing_response"
	AuditEventResourceTypeConfirmPaymentResponse            AuditEventResourceType = "confirm_payment_response"
	AuditEventResourceTypeOAuthResponse                     AuditEventResourceType = "oauth_response"
	AuditEventResourceTypeOAuthStatusResponse               AuditEventResourceType = "oauth_status_response"
	AuditEventResourceTypeStripePublishableKey              AuditEventResourceType = "stripe_publishable_key"
	AuditEventResourceTypeStripeStatus                      AuditEventResourceType = "stripe_status"
	AuditEventResourceTypeHealthcheck                       AuditEventResourceType = "healthcheck"
	AuditEventResourceTypeAgentDefinitionConfig             AuditEventResourceType = "agent_definition_config"
	AuditEventResourceTypeTriggerConfig                     AuditEventResourceType = "trigger_config"
	AuditEventResourceTypeCustomerContactInfo               AuditEventResourceType = "customer_contact_info"
	AuditEventResourceTypeCustomerFreightPreferences        AuditEventResourceType = "customer_freight_preferences"
	AuditEventResourceTypeCustomerDefaults                  AuditEventResourceType = "customer_defaults"
	AuditEventResourceTypeCustomerNotificationPreferences   AuditEventResourceType = "customer_notification_preferences"
	AuditEventResourceTypeOrderDiscount                     AuditEventResourceType = "order_discount"
	AuditEventResourceTypeSalesOrderLine                    AuditEventResourceType = "sales_order_line"
	AuditEventResourceTypeSalesOrderType                    AuditEventResourceType = "sales_order_type"
	AuditEventResourceTypeSalesOrderStatus                  AuditEventResourceType = "sales_order_status"
	AuditEventResourceTypeMaterial                          AuditEventResourceType = "material"
	AuditEventResourceTypeSupplierMaterial                  AuditEventResourceType = "supplier_material"
	AuditEventResourceTypePart                              AuditEventResourceType = "part"
	AuditEventResourceTypePermissionGroup                   AuditEventResourceType = "permission_group"
	AuditEventResourceTypePermission                        AuditEventResourceType = "permission"
	AuditEventResourceTypePick                              AuditEventResourceType = "pick"
	AuditEventResourceTypePickLine                          AuditEventResourceType = "pick_line"
	AuditEventResourceTypeProductType                       AuditEventResourceType = "product_type"
	AuditEventResourceTypeProduction                        AuditEventResourceType = "production"
	AuditEventResourceTypeProductionFlow                    AuditEventResourceType = "production_flow"
	AuditEventResourceTypeMap                               AuditEventResourceType = "map"
	AuditEventResourceTypePurchaseOrder                     AuditEventResourceType = "purchase_order"
	AuditEventResourceTypePurchaseOrderLine                 AuditEventResourceType = "purchase_order_line"
	AuditEventResourceTypeSupplier                          AuditEventResourceType = "supplier"
	AuditEventResourceTypeSupplierSummary                   AuditEventResourceType = "supplier_summary"
	AuditEventResourceTypeReceivableEntry                   AuditEventResourceType = "receivable_entry"
	AuditEventResourceTypeReceivingOrder                    AuditEventResourceType = "receiving_order"
	AuditEventResourceTypeReceivingOrderLine                AuditEventResourceType = "receiving_order_line"
	AuditEventResourceTypeEmailContact                      AuditEventResourceType = "email_contact"
	AuditEventResourceTypeAllocationEntry                   AuditEventResourceType = "allocation_entry"
	AuditEventResourceTypeOpenCreditEntry                   AuditEventResourceType = "open_credit_entry"
	AuditEventResourceTypeVolumeDiscount                    AuditEventResourceType = "volume_discount"
	AuditEventResourceTypeVolumeDiscountTier                AuditEventResourceType = "volume_discount_tier"
	AuditEventResourceTypeAnalyzeDeliveriesResponse         AuditEventResourceType = "analyze_deliveries_response"
	AuditEventResourceTypeAnalyzeManufacturingResponse      AuditEventResourceType = "analyze_manufacturing_response"
	AuditEventResourceTypeAnalyzeManufacturingBatchResponse AuditEventResourceType = "analyze_manufacturing_batch_response"
	AuditEventResourceTypeAnalyzeQuarterlyOrdersResponse    AuditEventResourceType = "analyze_quarterly_orders_response"
	AuditEventResourceTypeAnalyzeNewCustomersResponse       AuditEventResourceType = "analyze_new_customers_response"
	AuditEventResourceTypeAnalyzeOeeResponse                AuditEventResourceType = "analyze_oee_response"
	AuditEventResourceTypeCatalogProductLine                AuditEventResourceType = "catalog_product_line"
	AuditEventResourceTypeCatalogCategory                   AuditEventResourceType = "catalog_category"
	AuditEventResourceTypeCatalogProduct                    AuditEventResourceType = "catalog_product"
	AuditEventResourceTypeCatalogProperty                   AuditEventResourceType = "catalog_property"
	AuditEventResourceTypeCatalogAttribute                  AuditEventResourceType = "catalog_attribute"
	AuditEventResourceTypeDcLocation                        AuditEventResourceType = "dc_location"
	AuditEventResourceTypeEdiRun                            AuditEventResourceType = "edi_run"
	AuditEventResourceTypeInventoryItem                     AuditEventResourceType = "inventory_item"
	AuditEventResourceTypeAnalyzeWeeksOfSalesResponse       AuditEventResourceType = "analyze_weeks_of_sales_response"
	AuditEventResourceTypeBulkReconcileItemsResponse        AuditEventResourceType = "bulk_reconcile_items_response"
	AuditEventResourceTypeSysProperty                       AuditEventResourceType = "sys_property"
	AuditEventResourceTypeSysPropertyType                   AuditEventResourceType = "sys_property_type"
	AuditEventResourceTypeSysPropertyValue                  AuditEventResourceType = "sys_property_value"
	AuditEventResourceTypeTerritory                         AuditEventResourceType = "territory"
	AuditEventResourceTypeTenancy                           AuditEventResourceType = "tenancy"
	AuditEventResourceTypeCheckoutSession                   AuditEventResourceType = "checkout_session"
	AuditEventResourceTypeEstimateRateResult                AuditEventResourceType = "estimate_rate_result"
	AuditEventResourceTypeRateShopOption                    AuditEventResourceType = "rate_shop_option"
	AuditEventResourceTypeRateShopResult                    AuditEventResourceType = "rate_shop_result"
	AuditEventResourceTypeOwner                             AuditEventResourceType = "owner"
	AuditEventResourceTypeCreatedBy                         AuditEventResourceType = "created_by"
	AuditEventResourceTypeMessage                           AuditEventResourceType = "message"
	AuditEventResourceTypeAccountPhotoUploadResult          AuditEventResourceType = "account_photo_upload_result"
	AuditEventResourceTypeUserPhotoUploadResult             AuditEventResourceType = "user_photo_upload_result"
	AuditEventResourceTypeUserPhotoURL                      AuditEventResourceType = "user_photo_url"
	AuditEventResourceTypeBatchLot                          AuditEventResourceType = "batch_lot"
	AuditEventResourceTypeCheckDuplicateResult              AuditEventResourceType = "check_duplicate_result"
	AuditEventResourceTypeItemTrendPoint                    AuditEventResourceType = "item_trend_point"
	AuditEventResourceTypePackPickResponse                  AuditEventResourceType = "pack_pick_response"
	AuditEventResourceTypePickShipmentsResponse             AuditEventResourceType = "pick_shipments_response"
	AuditEventResourceTypeTenancyPendingRegistration        AuditEventResourceType = "tenancy_pending_registration"
	AuditEventResourceTypeInvoiceAllocationEntry            AuditEventResourceType = "invoice_allocation_entry"
	AuditEventResourceTypeAllocationCustomer                AuditEventResourceType = "allocation_customer"
	AuditEventResourceTypeCheckoutSalesOrderResponse        AuditEventResourceType = "checkout_sales_order_response"
	AuditEventResourceTypeCreateProductionRunResponse       AuditEventResourceType = "create_production_run_response"
	AuditEventResourceTypeSalesOrderPriceQuote              AuditEventResourceType = "sales_order_price_quote"
	AuditEventResourceTypeHubspotSyncJob                    AuditEventResourceType = "hubspot_sync_job"
	AuditEventResourceTypeHubspotSyncReport                 AuditEventResourceType = "hubspot_sync_report"
	AuditEventResourceTypeHubspotCompanyReview              AuditEventResourceType = "hubspot_company_review"
	AuditEventResourceTypeHubspotCompanyCandidate           AuditEventResourceType = "hubspot_company_candidate"
	AuditEventResourceTypeContactMatch                      AuditEventResourceType = "contact_match"
	AuditEventResourceTypeReplyDraft                        AuditEventResourceType = "reply_draft"
	AuditEventResourceTypeConversationLink                  AuditEventResourceType = "conversation_link"
	AuditEventResourceTypeMessagingGroup                    AuditEventResourceType = "messaging_group"
	AuditEventResourceTypeMessagingGroupMember              AuditEventResourceType = "messaging_group_member"
)

// Field-level before/after transition recorded during a mutation.
type AuditFieldChange struct {
	// Name of the changed field.
	Field string `json:"field" api:"required"`
	// New value as a JSON fragment.
	//
	// `null` for deletion events. Encoded as a JSON value (object, array, string,
	// number, boolean, or null), not a JSON-encoded string.
	NewValue any `json:"new_value" api:"required"`
	// Resource type identifier.
	//
	// Any of "audit_field_change".
	Object AuditFieldChangeObject `json:"object" api:"required"`
	// Previous value as a JSON fragment.
	//
	// `null` for creation events. Encoded as a JSON value (object, array, string,
	// number, boolean, or null), not a JSON-encoded string.
	OldValue any `json:"old_value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		NewValue    respjson.Field
		Object      respjson.Field
		OldValue    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditFieldChange) RawJSON() string { return r.JSON.raw }
func (r *AuditFieldChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AuditFieldChangeObject string

const (
	AuditFieldChangeObjectAuditFieldChange AuditFieldChangeObject = "audit_field_change"
)

// List represents a paginated list of resources.
type ListAuditEvent struct {
	// Resources in this page.
	Data []AuditEvent `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAuditEventObject `json:"object" api:"required"`
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
func (r ListAuditEvent) RawJSON() string { return r.JSON.raw }
func (r *ListAuditEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAuditEventObject string

const (
	ListAuditEventObjectList ListAuditEventObject = "list"
)

// List represents a paginated list of resources.
type ListAuditFieldChange struct {
	// Resources in this page.
	Data []AuditFieldChange `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAuditFieldChangeObject `json:"object" api:"required"`
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
func (r ListAuditFieldChange) RawJSON() string { return r.JSON.raw }
func (r *ListAuditFieldChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAuditFieldChangeObject string

const (
	ListAuditFieldChangeObjectList ListAuditFieldChangeObject = "list"
)

// List represents a paginated list of resources.
type ListObjectType struct {
	// Resources in this page.
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
	// "conversation_participant", "chat_message",
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
	// "email_inbox", "inventory_change_log", "invoice", "invoice_summary",
	// "invoice_line", "invoice_allocation", "invoice_for_payment", "shipment",
	// "shipment_summary", "shipment_line", "shipping_case", "shipping_case_label_url",
	// "settlement", "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "account_usage_response", "subscription_info",
	// "billing_portal_session_response", "switch_plan_response",
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
	Data []string `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListObjectTypeObject `json:"object" api:"required"`
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
func (r ListObjectType) RawJSON() string { return r.JSON.raw }
func (r *ListObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListObjectTypeObject string

const (
	ListObjectTypeObjectList ListObjectTypeObject = "list"
)

type CoreAuditEventGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account", "actor", "changes", "metadata", "request".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreAuditEventGetParams]'s query parameters as
// `url.Values`.
func (r CoreAuditEventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CoreAuditEventListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Restricts results to audit events on or before this timestamp.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Restricts results to audit events on or after this timestamp.
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	// Filter by the mutation type recorded on the event.
	//
	// Any of "create", "update", "delete", "restore", "archive", "approve", "deny".
	Actions []string `query:"actions,omitzero" json:"-"`
	// Filter by the _acting_ account: the account that performed the mutation.
	//
	// Results are always scoped to events where your account is either the acting
	// account or the target account; this narrows that set to specific acting accounts
	// — for example a specific customer's account that mutated a resource on your
	// account.
	ActorAccountIDs []string `query:"actor_account_ids,omitzero" json:"-"`
	// Filter by the actor identifier.
	//
	// Matches the event's `actor.id`: a user ID for `user` actors or an API key ID for
	// `api_key` actors.
	ActorIDs []string `query:"actor_ids,omitzero" json:"-"`
	// Filter by the actor type.
	//
	// Any of "user", "api_key", "agent", "group".
	ActorTypes []string `query:"actor_types,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account", "actor", "changes", "metadata", "request".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by the audited resource IDs.
	ResourceIDs []string `query:"resource_ids,omitzero" json:"-"`
	// Filter by the resource type of the audited entity.
	//
	// The full set of valid values is available from the List Audit Event Resource
	// Types endpoint.
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
	// "conversation_participant", "chat_message",
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
	// "email_inbox", "inventory_change_log", "invoice", "invoice_summary",
	// "invoice_line", "invoice_allocation", "invoice_for_payment", "shipment",
	// "shipment_summary", "shipment_line", "shipping_case", "shipping_case_label_url",
	// "settlement", "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "account_usage_response", "subscription_info",
	// "billing_portal_session_response", "switch_plan_response",
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
	ResourceTypes []string `query:"resource_types,omitzero" json:"-"`
	// Filter by the _target_ account the mutation was performed against (the event's
	// `account`).
	//
	// Results are always scoped to events where your account is either the acting
	// account or the target account; this narrows that set to specific target accounts
	// — for example a specific customer's or supplier's account.
	TargetAccountIDs []string `query:"target_account_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreAuditEventListParams]'s query parameters as
// `url.Values`.
func (r CoreAuditEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
