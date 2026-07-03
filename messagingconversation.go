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

// Create conversations, send and read messages (1:1 direct messages).
//
// MessagingConversationService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationService] method instead.
type MessagingConversationService struct {
	options []option.RequestOption
	// Create conversations, send and read messages (1:1 direct messages).
	Actions MessagingConversationActionService
	// Create conversations, send and read messages (1:1 direct messages).
	Links MessagingConversationLinkService
	// Send, list, edit, and delete chat messages.
	Messages MessagingConversationMessageService
	// Add, remove, and manage participants (including agents) in a conversation.
	Participants MessagingConversationParticipantService
	Attachments  MessagingConversationAttachmentService
}

// NewMessagingConversationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingConversationService(opts ...option.RequestOption) (r MessagingConversationService) {
	r = MessagingConversationService{}
	r.options = opts
	r.Actions = NewMessagingConversationActionService(opts...)
	r.Links = NewMessagingConversationLinkService(opts...)
	r.Messages = NewMessagingConversationMessageService(opts...)
	r.Participants = NewMessagingConversationParticipantService(opts...)
	r.Attachments = NewMessagingConversationAttachmentService(opts...)
	return
}

// Starts a conversation between participants.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingConversationService) New(ctx context.Context, params MessagingConversationNewParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns one conversation (with participants) the caller belongs to.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingConversationService) Get(ctx context.Context, id string, query MessagingConversationGetParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Renames a conversation.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationService) Update(ctx context.Context, id string, params MessagingConversationUpdateParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns the caller's conversations, most-recently-active first.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingConversationService) List(ctx context.Context, query MessagingConversationListParams, opts ...option.RequestOption) (res *ListConversation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single tool invocation performed by an agent during a run.
//
// Each action records the tool that was called, its input and output, and any
// human review decision.
type AgentAction struct {
	// Agent action ID.
	ID string `json:"id" api:"required"`
	// When this action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Longer description of what the action does.
	Description string `json:"description" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Entity Entity `json:"entity" api:"required"`
	// Error message if the action failed.
	ErrorMessage string `json:"error_message" api:"required"`
	// When the action was executed.
	ExecutedAt time.Time `json:"executed_at" api:"required" format:"date-time"`
	// Arguments passed to the tool, as JSON.
	//
	// Shape depends on `tool`. Encoded as a JSON value (object, array, string, number,
	// boolean, or null), not a JSON-encoded string.
	Input any `json:"input" api:"required"`
	// Short human-readable label summarizing the action.
	Label string `json:"label" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_action".
	Object AgentActionObject `json:"object" api:"required"`
	// Result returned by the tool, as JSON.
	//
	// Recorded when the tool runs, so it is present even while the action is still
	// `pending_review` or `auto_approved`; the shape depends on `tool`, and it is `{}`
	// when the tool returned no output. Encoded as a JSON value (object, array,
	// string, number, boolean, or null), not a JSON-encoded string.
	Output any `json:"output" api:"required"`
	// Whether this action must be reviewed by a human before it can execute.
	//
	// Any of "not_required", "required".
	ReviewRequirement AgentActionReviewRequirement `json:"review_requirement" api:"required"`
	// When a human review decision was recorded for the action.
	ReviewedAt time.Time `json:"reviewed_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	ReviewedBy Actor `json:"reviewed_by" api:"required"`
	// A single execution of an agent, from trigger through completion.
	Run *AgentRun `json:"run" api:"required"`
	// Current action status.
	//
	// - `pending_review`: awaiting human review before it can execute.
	// - `auto_approved`: automatically approved by policy.
	// - `approved`: manually approved by a user.
	// - `rejected`: rejected by a user; will not execute.
	// - `executed`: successfully executed.
	// - `failed`: errored during execution; see `error_message`.
	//
	// Any of "pending_review", "auto_approved", "approved", "rejected", "executed",
	// "failed".
	Status AgentActionStatus `json:"status" api:"required"`
	// The tool the agent invoked for this action.
	//
	//   - `create_artifact`: create an artifact such as a report, document, or data
	//     export.
	//   - `read_doc`: read Augno documentation pages.
	//   - `fetch_url`: fetch content from a public URL.
	//   - `draft_reply`: propose a reply to the case's external party as a draft held
	//     for human approval (not sent).
	//   - `send_email`: send an email reply through the conversation's bound inbox.
	//
	// Any of "create_artifact", "read_doc", "fetch_url", "send_email", "draft_reply".
	Tool AgentActionTool `json:"tool" api:"required"`
	// When this action was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		Entity            respjson.Field
		ErrorMessage      respjson.Field
		ExecutedAt        respjson.Field
		Input             respjson.Field
		Label             respjson.Field
		Object            respjson.Field
		Output            respjson.Field
		ReviewRequirement respjson.Field
		ReviewedAt        respjson.Field
		ReviewedBy        respjson.Field
		Run               respjson.Field
		Status            respjson.Field
		Tool              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAction) RawJSON() string { return r.JSON.raw }
func (r *AgentAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentActionObject string

const (
	AgentActionObjectAgentAction AgentActionObject = "agent_action"
)

// Whether this action must be reviewed by a human before it can execute.
type AgentActionReviewRequirement string

const (
	AgentActionReviewRequirementNotRequired AgentActionReviewRequirement = "not_required"
	AgentActionReviewRequirementRequired    AgentActionReviewRequirement = "required"
)

// Current action status.
//
// - `pending_review`: awaiting human review before it can execute.
// - `auto_approved`: automatically approved by policy.
// - `approved`: manually approved by a user.
// - `rejected`: rejected by a user; will not execute.
// - `executed`: successfully executed.
// - `failed`: errored during execution; see `error_message`.
type AgentActionStatus string

const (
	AgentActionStatusPendingReview AgentActionStatus = "pending_review"
	AgentActionStatusAutoApproved  AgentActionStatus = "auto_approved"
	AgentActionStatusApproved      AgentActionStatus = "approved"
	AgentActionStatusRejected      AgentActionStatus = "rejected"
	AgentActionStatusExecuted      AgentActionStatus = "executed"
	AgentActionStatusFailed        AgentActionStatus = "failed"
)

// The tool the agent invoked for this action.
//
//   - `create_artifact`: create an artifact such as a report, document, or data
//     export.
//   - `read_doc`: read Augno documentation pages.
//   - `fetch_url`: fetch content from a public URL.
//   - `draft_reply`: propose a reply to the case's external party as a draft held
//     for human approval (not sent).
//   - `send_email`: send an email reply through the conversation's bound inbox.
type AgentActionTool string

const (
	AgentActionToolCreateArtifact AgentActionTool = "create_artifact"
	AgentActionToolReadDoc        AgentActionTool = "read_doc"
	AgentActionToolFetchURL       AgentActionTool = "fetch_url"
	AgentActionToolSendEmail      AgentActionTool = "send_email"
	AgentActionToolDraftReply     AgentActionTool = "draft_reply"
)

// An AI agent available to the account.
//
// The definition describes what the agent does, how its runs are triggered, the
// tools it can use, and whether it is currently enabled for the account.
type AgentDefinition struct {
	// Agent definition ID.
	ID string `json:"id" api:"required"`
	// Category grouping for the agent (e.g. `order_processing`), used to organize
	// agents in the UI.
	CategoryCode string `json:"category_code" api:"required"`
	// Agent-level configuration controlling LLM behavior and trigger settings.
	//
	// Distinct from per-tool configuration (`tools[].config`), which configures
	// individual tools attached to the agent.
	Config AgentDefinitionConfig `json:"config" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the agent is provided by Augno or created in this account.
	//
	// - `system`: provided by Augno; cannot be edited or deleted.
	// - `custom`: created by a user in this account.
	//
	// Any of "system", "custom".
	DefinitionType AgentDefinitionDefinitionType `json:"definition_type" api:"required"`
	// Description of what the agent does.
	Description string `json:"description" api:"required"`
	// Whether the current user can edit this agent definition.
	//
	// Always `read_only` for `system` definitions.
	//
	// Any of "editable", "read_only".
	Editability AgentDefinitionEditability `json:"editability" api:"required"`
	// Human-readable name of the agent.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_definition".
	Object AgentDefinitionObject `json:"object" api:"required"`
	// A named set of permissions that can be assigned to users to control what they
	// can access.
	Role Role `json:"role" api:"required"`
	// URL-friendly identifier for the agent.
	Slug string `json:"slug" api:"required"`
	// Whether this agent is enabled for the current account.
	//
	// Activation is per-account: a `system` agent shared across accounts can be
	// `active` for one account and `inactive` for another. An `inactive` agent does
	// not run.
	//
	// Any of "active", "inactive".
	Status AgentDefinitionStatus `json:"status" api:"required"`
	// List represents a paginated list of resources.
	Tools ListAgentDefinitionTool `json:"tools" api:"required"`
	// How runs of this agent are initiated.
	//
	//   - `scheduled`: runs on a cron schedule (see
	//     `config.trigger_config.cron_schedule`).
	//   - `event`: runs in response to platform events (see
	//     `config.trigger_config.event_filters`).
	//   - `manual`: runs only when explicitly invoked.
	//   - `chat`: runs in response to a chat message; the run is linked to a
	//     conversation and posts its reply back into it.
	//
	// Any of "scheduled", "manual", "event", "chat".
	TriggerType AgentDefinitionTriggerType `json:"trigger_type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CategoryCode   respjson.Field
		Config         respjson.Field
		CreatedAt      respjson.Field
		DefinitionType respjson.Field
		Description    respjson.Field
		Editability    respjson.Field
		Name           respjson.Field
		Object         respjson.Field
		Role           respjson.Field
		Slug           respjson.Field
		Status         respjson.Field
		Tools          respjson.Field
		TriggerType    respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentDefinition) RawJSON() string { return r.JSON.raw }
func (r *AgentDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the agent is provided by Augno or created in this account.
//
// - `system`: provided by Augno; cannot be edited or deleted.
// - `custom`: created by a user in this account.
type AgentDefinitionDefinitionType string

const (
	AgentDefinitionDefinitionTypeSystem AgentDefinitionDefinitionType = "system"
	AgentDefinitionDefinitionTypeCustom AgentDefinitionDefinitionType = "custom"
)

// Whether the current user can edit this agent definition.
//
// Always `read_only` for `system` definitions.
type AgentDefinitionEditability string

const (
	AgentDefinitionEditabilityEditable AgentDefinitionEditability = "editable"
	AgentDefinitionEditabilityReadOnly AgentDefinitionEditability = "read_only"
)

// Resource type identifier.
type AgentDefinitionObject string

const (
	AgentDefinitionObjectAgentDefinition AgentDefinitionObject = "agent_definition"
)

// Whether this agent is enabled for the current account.
//
// Activation is per-account: a `system` agent shared across accounts can be
// `active` for one account and `inactive` for another. An `inactive` agent does
// not run.
type AgentDefinitionStatus string

const (
	AgentDefinitionStatusActive   AgentDefinitionStatus = "active"
	AgentDefinitionStatusInactive AgentDefinitionStatus = "inactive"
)

// How runs of this agent are initiated.
//
//   - `scheduled`: runs on a cron schedule (see
//     `config.trigger_config.cron_schedule`).
//   - `event`: runs in response to platform events (see
//     `config.trigger_config.event_filters`).
//   - `manual`: runs only when explicitly invoked.
//   - `chat`: runs in response to a chat message; the run is linked to a
//     conversation and posts its reply back into it.
type AgentDefinitionTriggerType string

const (
	AgentDefinitionTriggerTypeScheduled AgentDefinitionTriggerType = "scheduled"
	AgentDefinitionTriggerTypeManual    AgentDefinitionTriggerType = "manual"
	AgentDefinitionTriggerTypeEvent     AgentDefinitionTriggerType = "event"
	AgentDefinitionTriggerTypeChat      AgentDefinitionTriggerType = "chat"
)

// Agent-level configuration controlling LLM behavior and trigger settings.
//
// Distinct from per-tool configuration (`tools[].config`), which configures
// individual tools attached to the agent.
type AgentDefinitionConfig struct {
	// Per-endpoint-tool human-review overrides, keyed by tool slug.
	//
	// When an entry is `true`, the run pauses in `awaiting_approval` each time the
	// agent calls that endpoint-tool until it is approved via the Continue Agent Run
	// endpoint. Slugs absent from the map do not require review.
	EndpointToolReview map[string]bool `json:"endpoint_tool_review" api:"required"`
	// API-endpoint tools the agent may discover and use, by slug (e.g.
	// `create_account_group`).
	//
	// These correspond to tools listed by the List Tools endpoint with category
	// `api_endpoint`. A single entry `*` grants the entire endpoint-tool catalog.
	EndpointToolSlugs []string `json:"endpoint_tool_slugs" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_definition_config".
	Object AgentDefinitionConfigObject `json:"object" api:"required"`
	// System prompt / instructions for the agent.
	SystemPrompt string `json:"system_prompt" api:"required"`
	// LLM sampling temperature between 0 and 1.
	Temperature float64 `json:"temperature" api:"required"`
	// Intelligence and cost tier for the agent's reasoning.
	//
	// Selects how capable and expensive a model the agent uses without pinning a
	// specific model; higher tiers reason better but cost more. Leaving it unset uses
	// the default tier.
	//
	//   - `frontier`: the most capable tier, for multi-step planning, ambiguous agent
	//     work, and hard coding or architecture tasks.
	//   - `high`: the default tier, for normal planning, code edits, synthesis, and
	//     customer-facing reasoning.
	//   - `balanced`: for research, summarization, classification, structured
	//     extraction, and light tool use.
	//   - `cheap`: for simple transforms, validation, formatting, and routing.
	//
	// Any of "frontier", "high", "balanced", "cheap", "legacy".
	Tier AgentDefinitionConfigTier `json:"tier" api:"required"`
	// Trigger-type-specific configuration.
	//
	// Which fields are populated depends on the agent's `trigger_type`:
	//
	// - `scheduled`: `cron_schedule` (and optionally `timezone`) is set.
	// - `event`: `event_filters` is set.
	// - `manual`: all fields are empty.
	TriggerConfig TriggerConfig `json:"trigger_config" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointToolReview respjson.Field
		EndpointToolSlugs  respjson.Field
		Object             respjson.Field
		SystemPrompt       respjson.Field
		Temperature        respjson.Field
		Tier               respjson.Field
		TriggerConfig      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentDefinitionConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentDefinitionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentDefinitionConfigObject string

const (
	AgentDefinitionConfigObjectAgentDefinitionConfig AgentDefinitionConfigObject = "agent_definition_config"
)

// Intelligence and cost tier for the agent's reasoning.
//
// Selects how capable and expensive a model the agent uses without pinning a
// specific model; higher tiers reason better but cost more. Leaving it unset uses
// the default tier.
//
//   - `frontier`: the most capable tier, for multi-step planning, ambiguous agent
//     work, and hard coding or architecture tasks.
//   - `high`: the default tier, for normal planning, code edits, synthesis, and
//     customer-facing reasoning.
//   - `balanced`: for research, summarization, classification, structured
//     extraction, and light tool use.
//   - `cheap`: for simple transforms, validation, formatting, and routing.
type AgentDefinitionConfigTier string

const (
	AgentDefinitionConfigTierFrontier AgentDefinitionConfigTier = "frontier"
	AgentDefinitionConfigTierHigh     AgentDefinitionConfigTier = "high"
	AgentDefinitionConfigTierBalanced AgentDefinitionConfigTier = "balanced"
	AgentDefinitionConfigTierCheap    AgentDefinitionConfigTier = "cheap"
	AgentDefinitionConfigTierLegacy   AgentDefinitionConfigTier = "legacy"
)

// Tool attached to an agent definition.
//
// Pairs an AvailableTool with agent-specific config values.
type AgentDefinitionTool struct {
	// Agent definition tool ID.
	ID string `json:"id" api:"required"`
	// Instance-specific configuration for this tool.
	//
	// Must conform to the tool's `config_schema`. Encoded as a JSON value (object,
	// array, string, number, boolean, or null), not a JSON-encoded string.
	Config any `json:"config" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_definition_tool".
	Object AgentDefinitionToolObject `json:"object" api:"required"`
	// Whether calls to this tool must be approved by a user before they execute.
	//
	// When `required`, the run pauses in the `awaiting_approval` status each time the
	// agent invokes this tool; approve or allow the tool via the Continue Agent Run
	// endpoint to proceed.
	//
	// Any of "not_required", "required".
	ReviewRequirement AgentDefinitionToolReviewRequirement `json:"review_requirement" api:"required"`
	// Sort order within the agent.
	SortOrder int64 `json:"sort_order" api:"required"`
	// Platform tool that can be attached to agents.
	Tool AvailableTool `json:"tool" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Config            respjson.Field
		Object            respjson.Field
		ReviewRequirement respjson.Field
		SortOrder         respjson.Field
		Tool              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentDefinitionTool) RawJSON() string { return r.JSON.raw }
func (r *AgentDefinitionTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentDefinitionToolObject string

const (
	AgentDefinitionToolObjectAgentDefinitionTool AgentDefinitionToolObject = "agent_definition_tool"
)

// Whether calls to this tool must be approved by a user before they execute.
//
// When `required`, the run pauses in the `awaiting_approval` status each time the
// agent invokes this tool; approve or allow the tool via the Continue Agent Run
// endpoint to proceed.
type AgentDefinitionToolReviewRequirement string

const (
	AgentDefinitionToolReviewRequirementNotRequired AgentDefinitionToolReviewRequirement = "not_required"
	AgentDefinitionToolReviewRequirementRequired    AgentDefinitionToolReviewRequirement = "required"
)

// A single execution of an agent, from trigger through completion.
type AgentRun struct {
	// Agent run ID.
	ID string `json:"id" api:"required"`
	// List represents a paginated list of resources.
	Actions *ListAgentAction `json:"actions" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// When this run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An AI agent available to the account.
	//
	// The definition describes what the agent does, how its runs are triggered, the
	// tools it can use, and whether it is currently enabled for the account.
	Definition AgentDefinition `json:"definition" api:"required"`
	// Duration in milliseconds.
	DurationMs int64 `json:"duration_ms" api:"required"`
	// Error message if the run failed.
	ErrorMessage string `json:"error_message" api:"required"`
	// Input provided to the agent at the start of the run, as JSON. Encoded as a JSON
	// value (object, array, string, number, boolean, or null), not a JSON-encoded
	// string.
	Input any `json:"input" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_run".
	Object AgentRunObject `json:"object" api:"required"`
	// Final output produced by the agent, as JSON.
	//
	// Populated only once the run has completed successfully. Encoded as a JSON value
	// (object, array, string, number, boolean, or null), not a JSON-encoded string.
	Output any `json:"output" api:"required"`
	// When the run started executing.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// Current run status.
	//
	// - `pending`: queued but not yet started.
	// - `running`: currently executing.
	// - `awaiting_input`: paused, waiting for user input before continuing.
	// - `awaiting_approval`: paused, waiting for a pending action to be approved.
	// - `completed`: finished successfully.
	// - `failed`: stopped after an error; see `error_message`.
	// - `cancelled`: stopped before completion by a user.
	//
	// Any of "pending", "running", "completed", "failed", "cancelled",
	// "awaiting_input", "awaiting_approval".
	Status AgentRunStatus `json:"status" api:"required"`
	// List represents a paginated list of resources.
	Steps ListAgentRunStep `json:"steps" api:"required"`
	// How this run was initiated.
	//
	//   - `scheduled`: started by the agent's cron schedule.
	//   - `event`: started in response to a platform event.
	//   - `manual`: started by an explicit request; see `triggered_by`.
	//   - `chat`: started by a message in a conversation, with the agent's reply posted
	//     back into that conversation.
	//
	// Any of "scheduled", "manual", "event", "chat".
	TriggerType AgentRunTriggerType `json:"trigger_type" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	TriggeredBy Actor `json:"triggered_by" api:"required"`
	// When this run was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Actions      respjson.Field
		CompletedAt  respjson.Field
		CreatedAt    respjson.Field
		Definition   respjson.Field
		DurationMs   respjson.Field
		ErrorMessage respjson.Field
		Input        respjson.Field
		Object       respjson.Field
		Output       respjson.Field
		StartedAt    respjson.Field
		Status       respjson.Field
		Steps        respjson.Field
		TriggerType  respjson.Field
		TriggeredBy  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRun) RawJSON() string { return r.JSON.raw }
func (r *AgentRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentRunObject string

const (
	AgentRunObjectAgentRun AgentRunObject = "agent_run"
)

// Current run status.
//
// - `pending`: queued but not yet started.
// - `running`: currently executing.
// - `awaiting_input`: paused, waiting for user input before continuing.
// - `awaiting_approval`: paused, waiting for a pending action to be approved.
// - `completed`: finished successfully.
// - `failed`: stopped after an error; see `error_message`.
// - `cancelled`: stopped before completion by a user.
type AgentRunStatus string

const (
	AgentRunStatusPending          AgentRunStatus = "pending"
	AgentRunStatusRunning          AgentRunStatus = "running"
	AgentRunStatusCompleted        AgentRunStatus = "completed"
	AgentRunStatusFailed           AgentRunStatus = "failed"
	AgentRunStatusCancelled        AgentRunStatus = "cancelled"
	AgentRunStatusAwaitingInput    AgentRunStatus = "awaiting_input"
	AgentRunStatusAwaitingApproval AgentRunStatus = "awaiting_approval"
)

// How this run was initiated.
//
//   - `scheduled`: started by the agent's cron schedule.
//   - `event`: started in response to a platform event.
//   - `manual`: started by an explicit request; see `triggered_by`.
//   - `chat`: started by a message in a conversation, with the agent's reply posted
//     back into that conversation.
type AgentRunTriggerType string

const (
	AgentRunTriggerTypeScheduled AgentRunTriggerType = "scheduled"
	AgentRunTriggerTypeManual    AgentRunTriggerType = "manual"
	AgentRunTriggerTypeEvent     AgentRunTriggerType = "event"
	AgentRunTriggerTypeChat      AgentRunTriggerType = "chat"
)

// A single event in an agent run's execution timeline.
type AgentRunStep struct {
	// Agent run step ID.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// Text payload for the step, such as a message body or a tool result.
	Content string `json:"content" api:"required"`
	// When this step was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Duration in milliseconds.
	DurationMs int64 `json:"duration_ms" api:"required"`
	// Additional structured data for the step, as JSON. Encoded as a JSON value
	// (object, array, string, number, boolean, or null), not a JSON-encoded string.
	Metadata any `json:"metadata" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_run_step".
	Object AgentRunStepObject `json:"object" api:"required"`
	// Zero-based position of this step within the run's timeline.
	Sequence int64 `json:"sequence" api:"required"`
	// The kind of timeline event (e.g. `trigger_received`, `user_message`,
	// `assistant_message`, `tool_call`, `tool_result`, `awaiting_approval`,
	// `completion`, `error`).
	StepType string `json:"step_type" api:"required"`
	// Short title for the step.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actor       respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		DurationMs  respjson.Field
		Metadata    respjson.Field
		Object      respjson.Field
		Sequence    respjson.Field
		StepType    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunStep) RawJSON() string { return r.JSON.raw }
func (r *AgentRunStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentRunStepObject string

const (
	AgentRunStepObjectAgentRunStep AgentRunStepObject = "agent_run_step"
)

// Platform tool that can be attached to agents.
type AvailableTool struct {
	// Category grouping for the tool (e.g. `built_in`).
	Category string `json:"category" api:"required"`
	// JSON schema describing the configuration options this tool accepts.
	//
	// Defines the shape of the `config` field on AgentDefinitionTool.
	//
	// For example:
	//
	// ````json
	//
	//	{
	//	  "type": "object",
	//	  "properties": {
	//	    "max_results": {
	//	      "type": "integer",
	//	      "default": 10
	//	    }
	//	  }
	//	}
	//
	// ``` Encoded as a JSON value (object, array, string, number, boolean, or null), not a JSON-encoded string.
	// ````
	ConfigSchema any `json:"config_schema" api:"required"`
	// Tool description.
	Description string `json:"description" api:"required"`
	// Whether invoking this tool changes server state.
	//
	// True for any `api_endpoint` tool whose underlying operation is not a read
	// (non-GET); always false for `built_in` tools. The agent-configuration UI uses
	// this to default such tools to requiring human review.
	Mutating bool `json:"mutating" api:"required"`
	// Tool name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "available_tool".
	Object AvailableToolObject `json:"object" api:"required"`
	// Permission scopes the agent's role must hold for this tool to be usable (e.g.
	// `products:read`).
	RequiredPermissions []string `json:"required_permissions" api:"required"`
	// Role type the caller must have for this tool, when the operation is gated by
	// role rather than a permission (e.g. `admin`).
	RequiredRoleType string `json:"required_role_type" api:"required"`
	// A stable identifier used when attaching the tool to an agent.
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category            respjson.Field
		ConfigSchema        respjson.Field
		Description         respjson.Field
		Mutating            respjson.Field
		Name                respjson.Field
		Object              respjson.Field
		RequiredPermissions respjson.Field
		RequiredRoleType    respjson.Field
		Slug                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailableTool) RawJSON() string { return r.JSON.raw }
func (r *AvailableTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AvailableToolObject string

const (
	AvailableToolObjectAvailableTool AvailableToolObject = "available_tool"
)

// A conversation thread the caller participates in.
type Conversation struct {
	// Conversation ID.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Assignee Actor `json:"assignee" api:"required"`
	// Whether this is a team-only conversation (`internal`) or a customer-facing case
	// (`customer`).
	//
	// Any of "internal", "customer".
	Audience ConversationAudience `json:"audience" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A reusable roster: a named set of members (users and/or agents) that seeds new
	// conversations.
	//
	// Starting a conversation from a group snapshots its current members into that
	// conversation, so the same group can back many conversations (each with its own
	// title); later edits to the group never change conversations already created from
	// it.
	Group MessagingGroup `json:"group" api:"required"`
	// A chat message within a conversation.
	LastMessage *Message `json:"last_message" api:"required"`
	// When the most recent message was sent.
	//
	// `null` when the conversation has no messages yet.
	LastMessageAt time.Time `json:"last_message_at" api:"required" format:"date-time"`
	// Whether the conversation is under legal hold.
	//
	// Exempts the conversation from retention purging and redaction.
	//
	// Any of "released", "held".
	LegalHold ConversationLegalHold `json:"legal_hold" api:"required"`
	// Resource type identifier.
	//
	// Any of "conversation".
	Object ConversationObject `json:"object" api:"required"`
	// List represents a paginated list of resources.
	Participants ListConversationParticipant `json:"participants" api:"required"`
	// The caller's effective status.
	//
	// - `hidden` when the caller has hidden the conversation
	// - otherwise the account-level lifecycle state
	//
	// Any of "active", "archived", "hidden".
	Status ConversationStatus `json:"status" api:"required"`
	// The display title of a group conversation.
	//
	// `null` for direct messages, where the client derives a title from the
	// participants.
	Title string `json:"title" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Topic Entity `json:"topic" api:"required"`
	// What kind of conversation this is.
	//
	//   - `direct_message`: a 1:1 thread between two users.
	//   - `group`: a named thread with multiple user or agent members (including
	//     customer-facing support cases).
	//   - `system`: a system channel that delivers automated account alerts.
	//
	// Any of "direct_message", "group", "system".
	Type ConversationType `json:"type" api:"required"`
	// Number of messages the caller has not yet read.
	Unread int64 `json:"unread" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The triage lane of a customer-facing case.
	//
	// Only set for customer-audience conversations.
	//
	// - `new`: opened but not yet triaged.
	// - `open`: actively being worked.
	// - `waiting_internal`: blocked on the internal team.
	// - `waiting_external`: blocked on an external reply.
	// - `needs_approval`: a drafted reply is awaiting human approval.
	// - `resolved`: closed out.
	//
	// Any of "new", "open", "waiting_internal", "waiting_external", "needs_approval",
	// "resolved".
	WorkflowStatus ConversationWorkflowStatus `json:"workflow_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Assignee       respjson.Field
		Audience       respjson.Field
		CreatedAt      respjson.Field
		Group          respjson.Field
		LastMessage    respjson.Field
		LastMessageAt  respjson.Field
		LegalHold      respjson.Field
		Object         respjson.Field
		Participants   respjson.Field
		Status         respjson.Field
		Title          respjson.Field
		Topic          respjson.Field
		Type           respjson.Field
		Unread         respjson.Field
		UpdatedAt      respjson.Field
		WorkflowStatus respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Conversation) RawJSON() string { return r.JSON.raw }
func (r *Conversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this is a team-only conversation (`internal`) or a customer-facing case
// (`customer`).
type ConversationAudience string

const (
	ConversationAudienceInternal ConversationAudience = "internal"
	ConversationAudienceCustomer ConversationAudience = "customer"
)

// Whether the conversation is under legal hold.
//
// Exempts the conversation from retention purging and redaction.
type ConversationLegalHold string

const (
	ConversationLegalHoldReleased ConversationLegalHold = "released"
	ConversationLegalHoldHeld     ConversationLegalHold = "held"
)

// Resource type identifier.
type ConversationObject string

const (
	ConversationObjectConversation ConversationObject = "conversation"
)

// The caller's effective status.
//
// - `hidden` when the caller has hidden the conversation
// - otherwise the account-level lifecycle state
type ConversationStatus string

const (
	ConversationStatusActive   ConversationStatus = "active"
	ConversationStatusArchived ConversationStatus = "archived"
	ConversationStatusHidden   ConversationStatus = "hidden"
)

// What kind of conversation this is.
//
//   - `direct_message`: a 1:1 thread between two users.
//   - `group`: a named thread with multiple user or agent members (including
//     customer-facing support cases).
//   - `system`: a system channel that delivers automated account alerts.
type ConversationType string

const (
	ConversationTypeDirectMessage ConversationType = "direct_message"
	ConversationTypeGroup         ConversationType = "group"
	ConversationTypeSystem        ConversationType = "system"
)

// The triage lane of a customer-facing case.
//
// Only set for customer-audience conversations.
//
// - `new`: opened but not yet triaged.
// - `open`: actively being worked.
// - `waiting_internal`: blocked on the internal team.
// - `waiting_external`: blocked on an external reply.
// - `needs_approval`: a drafted reply is awaiting human approval.
// - `resolved`: closed out.
type ConversationWorkflowStatus string

const (
	ConversationWorkflowStatusNew             ConversationWorkflowStatus = "new"
	ConversationWorkflowStatusOpen            ConversationWorkflowStatus = "open"
	ConversationWorkflowStatusWaitingInternal ConversationWorkflowStatus = "waiting_internal"
	ConversationWorkflowStatusWaitingExternal ConversationWorkflowStatus = "waiting_external"
	ConversationWorkflowStatusNeedsApproval   ConversationWorkflowStatus = "needs_approval"
	ConversationWorkflowStatusResolved        ConversationWorkflowStatus = "resolved"
)

// A participant (membership) in a conversation.
type ConversationParticipant struct {
	// Participant ID.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// For agent participants with a keyword/mention policy, the keywords that trigger
	// it.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords" api:"required"`
	// For agent participants, when the agent is invoked in response to messages.
	//
	// `null` for non-agent participants.
	//
	// - `mention`: only when the agent is @mentioned.
	// - `keyword`: when a message contains one of the agent's trigger keywords.
	// - `always`: on every human message in the conversation.
	//
	// Any of "mention", "keyword", "always".
	AgentTriggerPolicy ConversationParticipantAgentTriggerPolicy `json:"agent_trigger_policy" api:"required"`
	// The participant's membership in the conversation.
	//
	// - `active`: currently a member.
	// - `left`: voluntarily left the conversation.
	// - `removed`: removed by an admin.
	// - `hidden`: still a member but has hidden the conversation from their own list.
	//
	// Any of "active", "left", "removed", "hidden".
	Membership ConversationParticipantMembership `json:"membership" api:"required"`
	// The participant's notification preference for the conversation.
	//
	// - `unmuted`: receives normal notifications.
	// - `muted`: notifications are suppressed (mentions may still pierce the mute).
	//
	// Any of "unmuted", "muted".
	Notifications ConversationParticipantNotifications `json:"notifications" api:"required"`
	// Resource type identifier.
	//
	// Any of "conversation_participant".
	Object ConversationParticipantObject `json:"object" api:"required"`
	// A participant's read position in a conversation — the basis for read receipts
	// ("who has seen this").
	ReadCursor ReadCursor `json:"read_cursor" api:"required"`
	// The participant's permission level in the conversation.
	//
	//   - `owner`: can rename or delete the conversation and manage its members and
	//     their roles.
	//   - `admin`: can add or remove members and rename the conversation.
	//   - `member`: can post, react, mute, and leave.
	//   - `viewer`: read-only access.
	//
	// Any of "owner", "admin", "member", "viewer".
	Role ConversationParticipantRole `json:"role" api:"required"`
	// The kind of participant.
	//
	// - `user`: an account user (a teammate).
	// - `agent`: an AI agent.
	// - `system`: the system itself, which posts automated messages.
	// - `customer`: an external customer in a support case.
	//
	// Any of "user", "agent", "system", "customer".
	Type ConversationParticipantType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Actor                respjson.Field
		AgentTriggerKeywords respjson.Field
		AgentTriggerPolicy   respjson.Field
		Membership           respjson.Field
		Notifications        respjson.Field
		Object               respjson.Field
		ReadCursor           respjson.Field
		Role                 respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationParticipant) RawJSON() string { return r.JSON.raw }
func (r *ConversationParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// For agent participants, when the agent is invoked in response to messages.
//
// `null` for non-agent participants.
//
// - `mention`: only when the agent is @mentioned.
// - `keyword`: when a message contains one of the agent's trigger keywords.
// - `always`: on every human message in the conversation.
type ConversationParticipantAgentTriggerPolicy string

const (
	ConversationParticipantAgentTriggerPolicyMention ConversationParticipantAgentTriggerPolicy = "mention"
	ConversationParticipantAgentTriggerPolicyKeyword ConversationParticipantAgentTriggerPolicy = "keyword"
	ConversationParticipantAgentTriggerPolicyAlways  ConversationParticipantAgentTriggerPolicy = "always"
)

// The participant's membership in the conversation.
//
// - `active`: currently a member.
// - `left`: voluntarily left the conversation.
// - `removed`: removed by an admin.
// - `hidden`: still a member but has hidden the conversation from their own list.
type ConversationParticipantMembership string

const (
	ConversationParticipantMembershipActive  ConversationParticipantMembership = "active"
	ConversationParticipantMembershipLeft    ConversationParticipantMembership = "left"
	ConversationParticipantMembershipRemoved ConversationParticipantMembership = "removed"
	ConversationParticipantMembershipHidden  ConversationParticipantMembership = "hidden"
)

// The participant's notification preference for the conversation.
//
// - `unmuted`: receives normal notifications.
// - `muted`: notifications are suppressed (mentions may still pierce the mute).
type ConversationParticipantNotifications string

const (
	ConversationParticipantNotificationsUnmuted ConversationParticipantNotifications = "unmuted"
	ConversationParticipantNotificationsMuted   ConversationParticipantNotifications = "muted"
)

// Resource type identifier.
type ConversationParticipantObject string

const (
	ConversationParticipantObjectConversationParticipant ConversationParticipantObject = "conversation_participant"
)

// The participant's permission level in the conversation.
//
//   - `owner`: can rename or delete the conversation and manage its members and
//     their roles.
//   - `admin`: can add or remove members and rename the conversation.
//   - `member`: can post, react, mute, and leave.
//   - `viewer`: read-only access.
type ConversationParticipantRole string

const (
	ConversationParticipantRoleOwner  ConversationParticipantRole = "owner"
	ConversationParticipantRoleAdmin  ConversationParticipantRole = "admin"
	ConversationParticipantRoleMember ConversationParticipantRole = "member"
	ConversationParticipantRoleViewer ConversationParticipantRole = "viewer"
)

// The kind of participant.
//
// - `user`: an account user (a teammate).
// - `agent`: an AI agent.
// - `system`: the system itself, which posts automated messages.
// - `customer`: an external customer in a support case.
type ConversationParticipantType string

const (
	ConversationParticipantTypeUser     ConversationParticipantType = "user"
	ConversationParticipantTypeAgent    ConversationParticipantType = "agent"
	ConversationParticipantTypeSystem   ConversationParticipantType = "system"
	ConversationParticipantTypeCustomer ConversationParticipantType = "customer"
)

// Request to create a conversation.
//
// The property Type is required.
type CreateConversationRequestParam struct {
	// The kind of conversation to create.
	//
	// Any of "direct_message", "group", "system".
	Type CreateConversationRequestType `json:"type,omitzero" api:"required"`
	// Seed a group conversation from a reusable roster.
	//
	// The roster's current members are copied into this conversation (in addition to
	// any `participant_account_user_ids`); the conversation is independent afterward.
	// Ignored for direct messages.
	GroupID param.Opt[string] `json:"group_id,omitzero"`
	// Title for a group conversation.
	//
	// Ignored for direct messages.
	Title param.Opt[string] `json:"title,omitzero"`
	// The id of the business record to anchor this conversation to.
	TopicResourceID param.Opt[string] `json:"topic_resource_id,omitzero"`
	// The other participant(s).
	//
	// For a direct message, exactly one account_user ID. For a group, the members to
	// add — optional when `group_id` seeds the roster.
	ParticipantAccountUserIDs []string `json:"participant_account_user_ids,omitzero"`
	// The type of business record to anchor this conversation to.
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
	TopicResourceType CreateConversationRequestTopicResourceType `json:"topic_resource_type,omitzero"`
	paramObj
}

func (r CreateConversationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateConversationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateConversationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of conversation to create.
type CreateConversationRequestType string

const (
	CreateConversationRequestTypeDirectMessage CreateConversationRequestType = "direct_message"
	CreateConversationRequestTypeGroup         CreateConversationRequestType = "group"
	CreateConversationRequestTypeSystem        CreateConversationRequestType = "system"
)

// The type of business record to anchor this conversation to.
type CreateConversationRequestTopicResourceType string

const (
	CreateConversationRequestTopicResourceTypeAccount                           CreateConversationRequestTopicResourceType = "account"
	CreateConversationRequestTopicResourceTypeActor                             CreateConversationRequestTopicResourceType = "actor"
	CreateConversationRequestTopicResourceTypeEntity                            CreateConversationRequestTopicResourceType = "entity"
	CreateConversationRequestTopicResourceTypeRecord                            CreateConversationRequestTopicResourceType = "record"
	CreateConversationRequestTopicResourceTypeFreight                           CreateConversationRequestTopicResourceType = "freight"
	CreateConversationRequestTopicResourceTypeSalesOrderTotals                  CreateConversationRequestTopicResourceType = "sales_order_totals"
	CreateConversationRequestTopicResourceTypeSalesOrderRelated                 CreateConversationRequestTopicResourceType = "sales_order_related"
	CreateConversationRequestTopicResourceTypeOrderContact                      CreateConversationRequestTopicResourceType = "order_contact"
	CreateConversationRequestTopicResourceTypeUser                              CreateConversationRequestTopicResourceType = "user"
	CreateConversationRequestTopicResourceTypeAddress                           CreateConversationRequestTopicResourceType = "address"
	CreateConversationRequestTopicResourceTypeAPIKey                            CreateConversationRequestTopicResourceType = "api_key"
	CreateConversationRequestTopicResourceTypeCreatedAPIKey                     CreateConversationRequestTopicResourceType = "created_api_key"
	CreateConversationRequestTopicResourceTypeRefreshToken                      CreateConversationRequestTopicResourceType = "refresh_token"
	CreateConversationRequestTopicResourceTypeList                              CreateConversationRequestTopicResourceType = "list"
	CreateConversationRequestTopicResourceTypeSandbox                           CreateConversationRequestTopicResourceType = "sandbox"
	CreateConversationRequestTopicResourceTypeRegistrationSession               CreateConversationRequestTopicResourceType = "registration_session"
	CreateConversationRequestTopicResourceTypePricingPlan                       CreateConversationRequestTopicResourceType = "pricing_plan"
	CreateConversationRequestTopicResourceTypeAccountPlan                       CreateConversationRequestTopicResourceType = "account_plan"
	CreateConversationRequestTopicResourceTypePlanChange                        CreateConversationRequestTopicResourceType = "plan_change"
	CreateConversationRequestTopicResourceTypeEnterpriseInquiry                 CreateConversationRequestTopicResourceType = "enterprise_inquiry"
	CreateConversationRequestTopicResourceTypeRequestLog                        CreateConversationRequestTopicResourceType = "request_log"
	CreateConversationRequestTopicResourceTypeAuditEvent                        CreateConversationRequestTopicResourceType = "audit_event"
	CreateConversationRequestTopicResourceTypeAuditFieldChange                  CreateConversationRequestTopicResourceType = "audit_field_change"
	CreateConversationRequestTopicResourceTypeRole                              CreateConversationRequestTopicResourceType = "role"
	CreateConversationRequestTopicResourceTypeUnit                              CreateConversationRequestTopicResourceType = "unit"
	CreateConversationRequestTopicResourceTypeAccountAffiliation                CreateConversationRequestTopicResourceType = "account_affiliation"
	CreateConversationRequestTopicResourceTypeAgentDefinition                   CreateConversationRequestTopicResourceType = "agent_definition"
	CreateConversationRequestTopicResourceTypeAvailableTool                     CreateConversationRequestTopicResourceType = "available_tool"
	CreateConversationRequestTopicResourceTypeAgentDefinitionTool               CreateConversationRequestTopicResourceType = "agent_definition_tool"
	CreateConversationRequestTopicResourceTypeAgentAccountStatus                CreateConversationRequestTopicResourceType = "agent_account_status"
	CreateConversationRequestTopicResourceTypeAgentRun                          CreateConversationRequestTopicResourceType = "agent_run"
	CreateConversationRequestTopicResourceTypeAgentAction                       CreateConversationRequestTopicResourceType = "agent_action"
	CreateConversationRequestTopicResourceTypeAgentRunStep                      CreateConversationRequestTopicResourceType = "agent_run_step"
	CreateConversationRequestTopicResourceTypeAgentTokenUsage                   CreateConversationRequestTopicResourceType = "agent_token_usage"
	CreateConversationRequestTopicResourceTypeAgentMemory                       CreateConversationRequestTopicResourceType = "agent_memory"
	CreateConversationRequestTopicResourceTypeNotification                      CreateConversationRequestTopicResourceType = "notification"
	CreateConversationRequestTopicResourceTypeNotificationUnreadCount           CreateConversationRequestTopicResourceType = "notification_unread_count"
	CreateConversationRequestTopicResourceTypeNotificationSendResult            CreateConversationRequestTopicResourceType = "notification_send_result"
	CreateConversationRequestTopicResourceTypeNotificationUnreadSummary         CreateConversationRequestTopicResourceType = "notification_unread_summary"
	CreateConversationRequestTopicResourceTypeAnnouncement                      CreateConversationRequestTopicResourceType = "announcement"
	CreateConversationRequestTopicResourceTypeConversation                      CreateConversationRequestTopicResourceType = "conversation"
	CreateConversationRequestTopicResourceTypeConversationParticipant           CreateConversationRequestTopicResourceType = "conversation_participant"
	CreateConversationRequestTopicResourceTypeReadCursor                        CreateConversationRequestTopicResourceType = "read_cursor"
	CreateConversationRequestTopicResourceTypeChatMessage                       CreateConversationRequestTopicResourceType = "chat_message"
	CreateConversationRequestTopicResourceTypeNotificationUnreadSummaryAccount  CreateConversationRequestTopicResourceType = "notification_unread_summary_account"
	CreateConversationRequestTopicResourceTypeMessagingBlock                    CreateConversationRequestTopicResourceType = "messaging_block"
	CreateConversationRequestTopicResourceTypeNotificationPreference            CreateConversationRequestTopicResourceType = "notification_preference"
	CreateConversationRequestTopicResourceTypeMessageAttachment                 CreateConversationRequestTopicResourceType = "message_attachment"
	CreateConversationRequestTopicResourceTypeAttachmentUploadTarget            CreateConversationRequestTopicResourceType = "attachment_upload_target"
	CreateConversationRequestTopicResourceTypeScheduledMessage                  CreateConversationRequestTopicResourceType = "scheduled_message"
	CreateConversationRequestTopicResourceTypeMessagingContact                  CreateConversationRequestTopicResourceType = "messaging_contact"
	CreateConversationRequestTopicResourceTypeMessageReport                     CreateConversationRequestTopicResourceType = "message_report"
	CreateConversationRequestTopicResourceTypeToolGroup                         CreateConversationRequestTopicResourceType = "tool_group"
	CreateConversationRequestTopicResourceTypeModel                             CreateConversationRequestTopicResourceType = "model"
	CreateConversationRequestTopicResourceTypePaymentTerm                       CreateConversationRequestTopicResourceType = "payment_term"
	CreateConversationRequestTopicResourceTypeShippingTerm                      CreateConversationRequestTopicResourceType = "shipping_term"
	CreateConversationRequestTopicResourceTypeQuantity                          CreateConversationRequestTopicResourceType = "quantity"
	CreateConversationRequestTopicResourceTypeAccountGroup                      CreateConversationRequestTopicResourceType = "account_group"
	CreateConversationRequestTopicResourceTypeSupportRoute                      CreateConversationRequestTopicResourceType = "support_route"
	CreateConversationRequestTopicResourceTypeSupportAvailability               CreateConversationRequestTopicResourceType = "support_availability"
	CreateConversationRequestTopicResourceTypeAccountStatus                     CreateConversationRequestTopicResourceType = "account_status"
	CreateConversationRequestTopicResourceTypeGeolocation                       CreateConversationRequestTopicResourceType = "geolocation"
	CreateConversationRequestTopicResourceTypeAccountUser                       CreateConversationRequestTopicResourceType = "account_user"
	CreateConversationRequestTopicResourceTypeDepartment                        CreateConversationRequestTopicResourceType = "department"
	CreateConversationRequestTopicResourceTypeAccountIntegration                CreateConversationRequestTopicResourceType = "account_integration"
	CreateConversationRequestTopicResourceTypeAccountPrice                      CreateConversationRequestTopicResourceType = "account_price"
	CreateConversationRequestTopicResourceTypeProductLine                       CreateConversationRequestTopicResourceType = "product_line"
	CreateConversationRequestTopicResourceTypeItemCategory                      CreateConversationRequestTopicResourceType = "item_category"
	CreateConversationRequestTopicResourceTypeAttribute                         CreateConversationRequestTopicResourceType = "attribute"
	CreateConversationRequestTopicResourceTypeRate                              CreateConversationRequestTopicResourceType = "rate"
	CreateConversationRequestTopicResourceTypeAccountGroupProductLineAccess     CreateConversationRequestTopicResourceType = "account_group_product_line_access"
	CreateConversationRequestTopicResourceTypeSalesTarget                       CreateConversationRequestTopicResourceType = "sales_target"
	CreateConversationRequestTopicResourceTypeAdjustmentType                    CreateConversationRequestTopicResourceType = "adjustment_type"
	CreateConversationRequestTopicResourceTypeAccountBranding                   CreateConversationRequestTopicResourceType = "account_branding"
	CreateConversationRequestTopicResourceTypeAccountPortal                     CreateConversationRequestTopicResourceType = "account_portal"
	CreateConversationRequestTopicResourceTypeAccountLogoURL                    CreateConversationRequestTopicResourceType = "account_logo_url"
	CreateConversationRequestTopicResourceTypePublicAccount                     CreateConversationRequestTopicResourceType = "public_account"
	CreateConversationRequestTopicResourceTypeProperty                          CreateConversationRequestTopicResourceType = "property"
	CreateConversationRequestTopicResourceTypeCarrier                           CreateConversationRequestTopicResourceType = "carrier"
	CreateConversationRequestTopicResourceTypeServiceLevel                      CreateConversationRequestTopicResourceType = "service_level"
	CreateConversationRequestTopicResourceTypeItem                              CreateConversationRequestTopicResourceType = "item"
	CreateConversationRequestTopicResourceTypeItemInventory                     CreateConversationRequestTopicResourceType = "item_inventory"
	CreateConversationRequestTopicResourceTypeProduct                           CreateConversationRequestTopicResourceType = "product"
	CreateConversationRequestTopicResourceTypeBatch                             CreateConversationRequestTopicResourceType = "batch"
	CreateConversationRequestTopicResourceTypeBatchFlowNode                     CreateConversationRequestTopicResourceType = "batch_flow_node"
	CreateConversationRequestTopicResourceTypeScanningConsumption               CreateConversationRequestTopicResourceType = "scanning_consumption"
	CreateConversationRequestTopicResourceTypeOpenBatchSummary                  CreateConversationRequestTopicResourceType = "open_batch_summary"
	CreateConversationRequestTopicResourceTypeScanningProductionStepInfo        CreateConversationRequestTopicResourceType = "scanning_production_step_info"
	CreateConversationRequestTopicResourceTypeScanningStation                   CreateConversationRequestTopicResourceType = "scanning_station"
	CreateConversationRequestTopicResourceTypeProductionStep                    CreateConversationRequestTopicResourceType = "production_step"
	CreateConversationRequestTopicResourceTypeProductionRun                     CreateConversationRequestTopicResourceType = "production_run"
	CreateConversationRequestTopicResourceTypeMachine                           CreateConversationRequestTopicResourceType = "machine"
	CreateConversationRequestTopicResourceTypeChildAccount                      CreateConversationRequestTopicResourceType = "child_account"
	CreateConversationRequestTopicResourceTypeUnitGroup                         CreateConversationRequestTopicResourceType = "unit_group"
	CreateConversationRequestTopicResourceTypeUnitGroupUnit                     CreateConversationRequestTopicResourceType = "unit_group_unit"
	CreateConversationRequestTopicResourceTypeConsumption                       CreateConversationRequestTopicResourceType = "consumption"
	CreateConversationRequestTopicResourceTypeCustomerProductLineAccess         CreateConversationRequestTopicResourceType = "customer_product_line_access"
	CreateConversationRequestTopicResourceTypeCustomer                          CreateConversationRequestTopicResourceType = "customer"
	CreateConversationRequestTopicResourceTypeFrequentlyOrderedProduct          CreateConversationRequestTopicResourceType = "frequently_ordered_product"
	CreateConversationRequestTopicResourceTypePriority                          CreateConversationRequestTopicResourceType = "priority"
	CreateConversationRequestTopicResourceTypeDelivery                          CreateConversationRequestTopicResourceType = "delivery"
	CreateConversationRequestTopicResourceTypeDeliveryLine                      CreateConversationRequestTopicResourceType = "delivery_line"
	CreateConversationRequestTopicResourceTypeSalesOrder                        CreateConversationRequestTopicResourceType = "sales_order"
	CreateConversationRequestTopicResourceTypeLocation                          CreateConversationRequestTopicResourceType = "location"
	CreateConversationRequestTopicResourceTypeLocationType                      CreateConversationRequestTopicResourceType = "location_type"
	CreateConversationRequestTopicResourceTypeLot                               CreateConversationRequestTopicResourceType = "lot"
	CreateConversationRequestTopicResourceTypeEmailLog                          CreateConversationRequestTopicResourceType = "email_log"
	CreateConversationRequestTopicResourceTypeEmailDomain                       CreateConversationRequestTopicResourceType = "email_domain"
	CreateConversationRequestTopicResourceTypeEmailInbox                        CreateConversationRequestTopicResourceType = "email_inbox"
	CreateConversationRequestTopicResourceTypeInventoryChangeLog                CreateConversationRequestTopicResourceType = "inventory_change_log"
	CreateConversationRequestTopicResourceTypeInvoice                           CreateConversationRequestTopicResourceType = "invoice"
	CreateConversationRequestTopicResourceTypeInvoiceSummary                    CreateConversationRequestTopicResourceType = "invoice_summary"
	CreateConversationRequestTopicResourceTypeInvoiceLine                       CreateConversationRequestTopicResourceType = "invoice_line"
	CreateConversationRequestTopicResourceTypeInvoiceAllocation                 CreateConversationRequestTopicResourceType = "invoice_allocation"
	CreateConversationRequestTopicResourceTypeInvoiceForPayment                 CreateConversationRequestTopicResourceType = "invoice_for_payment"
	CreateConversationRequestTopicResourceTypeShipment                          CreateConversationRequestTopicResourceType = "shipment"
	CreateConversationRequestTopicResourceTypeShipmentSummary                   CreateConversationRequestTopicResourceType = "shipment_summary"
	CreateConversationRequestTopicResourceTypeShipmentLine                      CreateConversationRequestTopicResourceType = "shipment_line"
	CreateConversationRequestTopicResourceTypeShippingCase                      CreateConversationRequestTopicResourceType = "shipping_case"
	CreateConversationRequestTopicResourceTypeShippingCaseLabelURL              CreateConversationRequestTopicResourceType = "shipping_case_label_url"
	CreateConversationRequestTopicResourceTypeSettlement                        CreateConversationRequestTopicResourceType = "settlement"
	CreateConversationRequestTopicResourceTypeSettlementSummary                 CreateConversationRequestTopicResourceType = "settlement_summary"
	CreateConversationRequestTopicResourceTypeRolePermission                    CreateConversationRequestTopicResourceType = "role_permission"
	CreateConversationRequestTopicResourceTypeRegistrationFlow                  CreateConversationRequestTopicResourceType = "registration_flow"
	CreateConversationRequestTopicResourceTypeRegistrationFlowOption            CreateConversationRequestTopicResourceType = "registration_flow_option"
	CreateConversationRequestTopicResourceTypeTransaction                       CreateConversationRequestTopicResourceType = "transaction"
	CreateConversationRequestTopicResourceTypeTransactionSummary                CreateConversationRequestTopicResourceType = "transaction_summary"
	CreateConversationRequestTopicResourceTypeTransactionMethod                 CreateConversationRequestTopicResourceType = "transaction_method"
	CreateConversationRequestTopicResourceTypeTransactionType                   CreateConversationRequestTopicResourceType = "transaction_type"
	CreateConversationRequestTopicResourceTypeTransactionAllocation             CreateConversationRequestTopicResourceType = "transaction_allocation"
	CreateConversationRequestTopicResourceTypeUsageItem                         CreateConversationRequestTopicResourceType = "usage_item"
	CreateConversationRequestTopicResourceTypeAccountUsageResponse              CreateConversationRequestTopicResourceType = "account_usage_response"
	CreateConversationRequestTopicResourceTypeSubscriptionInfo                  CreateConversationRequestTopicResourceType = "subscription_info"
	CreateConversationRequestTopicResourceTypeBillingPortalSessionResponse      CreateConversationRequestTopicResourceType = "billing_portal_session_response"
	CreateConversationRequestTopicResourceTypeSwitchPlanResponse                CreateConversationRequestTopicResourceType = "switch_plan_response"
	CreateConversationRequestTopicResourceTypeEnsureBillingCustomerResponse     CreateConversationRequestTopicResourceType = "ensure_billing_customer_response"
	CreateConversationRequestTopicResourceTypeSpendingCapResponse               CreateConversationRequestTopicResourceType = "spending_cap_response"
	CreateConversationRequestTopicResourceTypeAgentSpendInfo                    CreateConversationRequestTopicResourceType = "agent_spend_info"
	CreateConversationRequestTopicResourceTypeWebhookResponse                   CreateConversationRequestTopicResourceType = "webhook_response"
	CreateConversationRequestTopicResourceTypeAddressSuggestion                 CreateConversationRequestTopicResourceType = "address_suggestion"
	CreateConversationRequestTopicResourceTypeAddressComponents                 CreateConversationRequestTopicResourceType = "address_components"
	CreateConversationRequestTopicResourceTypeAddressDetailsResult              CreateConversationRequestTopicResourceType = "address_details_result"
	CreateConversationRequestTopicResourceTypeValidatedAddress                  CreateConversationRequestTopicResourceType = "validated_address"
	CreateConversationRequestTopicResourceTypePlanLimit                         CreateConversationRequestTopicResourceType = "plan_limit"
	CreateConversationRequestTopicResourceTypePlanChangeProration               CreateConversationRequestTopicResourceType = "plan_change_proration"
	CreateConversationRequestTopicResourceTypePlanChangeLineItem                CreateConversationRequestTopicResourceType = "plan_change_line_item"
	CreateConversationRequestTopicResourceTypeSetupBillingResponse              CreateConversationRequestTopicResourceType = "setup_billing_response"
	CreateConversationRequestTopicResourceTypeConfirmPaymentResponse            CreateConversationRequestTopicResourceType = "confirm_payment_response"
	CreateConversationRequestTopicResourceTypeOAuthResponse                     CreateConversationRequestTopicResourceType = "oauth_response"
	CreateConversationRequestTopicResourceTypeOAuthStatusResponse               CreateConversationRequestTopicResourceType = "oauth_status_response"
	CreateConversationRequestTopicResourceTypeStripePublishableKey              CreateConversationRequestTopicResourceType = "stripe_publishable_key"
	CreateConversationRequestTopicResourceTypeStripeStatus                      CreateConversationRequestTopicResourceType = "stripe_status"
	CreateConversationRequestTopicResourceTypeHealthcheck                       CreateConversationRequestTopicResourceType = "healthcheck"
	CreateConversationRequestTopicResourceTypeAgentDefinitionConfig             CreateConversationRequestTopicResourceType = "agent_definition_config"
	CreateConversationRequestTopicResourceTypeTriggerConfig                     CreateConversationRequestTopicResourceType = "trigger_config"
	CreateConversationRequestTopicResourceTypeCustomerContactInfo               CreateConversationRequestTopicResourceType = "customer_contact_info"
	CreateConversationRequestTopicResourceTypeCustomerFreightPreferences        CreateConversationRequestTopicResourceType = "customer_freight_preferences"
	CreateConversationRequestTopicResourceTypeCustomerDefaults                  CreateConversationRequestTopicResourceType = "customer_defaults"
	CreateConversationRequestTopicResourceTypeCustomerNotificationPreferences   CreateConversationRequestTopicResourceType = "customer_notification_preferences"
	CreateConversationRequestTopicResourceTypeOrderDiscount                     CreateConversationRequestTopicResourceType = "order_discount"
	CreateConversationRequestTopicResourceTypeSalesOrderLine                    CreateConversationRequestTopicResourceType = "sales_order_line"
	CreateConversationRequestTopicResourceTypeSalesOrderType                    CreateConversationRequestTopicResourceType = "sales_order_type"
	CreateConversationRequestTopicResourceTypeSalesOrderStatus                  CreateConversationRequestTopicResourceType = "sales_order_status"
	CreateConversationRequestTopicResourceTypeMaterial                          CreateConversationRequestTopicResourceType = "material"
	CreateConversationRequestTopicResourceTypeSupplierMaterial                  CreateConversationRequestTopicResourceType = "supplier_material"
	CreateConversationRequestTopicResourceTypePart                              CreateConversationRequestTopicResourceType = "part"
	CreateConversationRequestTopicResourceTypePermissionGroup                   CreateConversationRequestTopicResourceType = "permission_group"
	CreateConversationRequestTopicResourceTypePermission                        CreateConversationRequestTopicResourceType = "permission"
	CreateConversationRequestTopicResourceTypePick                              CreateConversationRequestTopicResourceType = "pick"
	CreateConversationRequestTopicResourceTypePickLine                          CreateConversationRequestTopicResourceType = "pick_line"
	CreateConversationRequestTopicResourceTypeProductType                       CreateConversationRequestTopicResourceType = "product_type"
	CreateConversationRequestTopicResourceTypeProduction                        CreateConversationRequestTopicResourceType = "production"
	CreateConversationRequestTopicResourceTypeProductionFlow                    CreateConversationRequestTopicResourceType = "production_flow"
	CreateConversationRequestTopicResourceTypeMap                               CreateConversationRequestTopicResourceType = "map"
	CreateConversationRequestTopicResourceTypePurchaseOrder                     CreateConversationRequestTopicResourceType = "purchase_order"
	CreateConversationRequestTopicResourceTypePurchaseOrderLine                 CreateConversationRequestTopicResourceType = "purchase_order_line"
	CreateConversationRequestTopicResourceTypeSupplier                          CreateConversationRequestTopicResourceType = "supplier"
	CreateConversationRequestTopicResourceTypeSupplierSummary                   CreateConversationRequestTopicResourceType = "supplier_summary"
	CreateConversationRequestTopicResourceTypeReceivableEntry                   CreateConversationRequestTopicResourceType = "receivable_entry"
	CreateConversationRequestTopicResourceTypeReceivingOrder                    CreateConversationRequestTopicResourceType = "receiving_order"
	CreateConversationRequestTopicResourceTypeReceivingOrderLine                CreateConversationRequestTopicResourceType = "receiving_order_line"
	CreateConversationRequestTopicResourceTypeEmailContact                      CreateConversationRequestTopicResourceType = "email_contact"
	CreateConversationRequestTopicResourceTypeAllocationEntry                   CreateConversationRequestTopicResourceType = "allocation_entry"
	CreateConversationRequestTopicResourceTypeOpenCreditEntry                   CreateConversationRequestTopicResourceType = "open_credit_entry"
	CreateConversationRequestTopicResourceTypeVolumeDiscount                    CreateConversationRequestTopicResourceType = "volume_discount"
	CreateConversationRequestTopicResourceTypeVolumeDiscountTier                CreateConversationRequestTopicResourceType = "volume_discount_tier"
	CreateConversationRequestTopicResourceTypeAnalyzeDeliveriesResponse         CreateConversationRequestTopicResourceType = "analyze_deliveries_response"
	CreateConversationRequestTopicResourceTypeAnalyzeManufacturingResponse      CreateConversationRequestTopicResourceType = "analyze_manufacturing_response"
	CreateConversationRequestTopicResourceTypeAnalyzeManufacturingBatchResponse CreateConversationRequestTopicResourceType = "analyze_manufacturing_batch_response"
	CreateConversationRequestTopicResourceTypeAnalyzeQuarterlyOrdersResponse    CreateConversationRequestTopicResourceType = "analyze_quarterly_orders_response"
	CreateConversationRequestTopicResourceTypeAnalyzeNewCustomersResponse       CreateConversationRequestTopicResourceType = "analyze_new_customers_response"
	CreateConversationRequestTopicResourceTypeAnalyzeOeeResponse                CreateConversationRequestTopicResourceType = "analyze_oee_response"
	CreateConversationRequestTopicResourceTypeCatalogProductLine                CreateConversationRequestTopicResourceType = "catalog_product_line"
	CreateConversationRequestTopicResourceTypeCatalogCategory                   CreateConversationRequestTopicResourceType = "catalog_category"
	CreateConversationRequestTopicResourceTypeCatalogProduct                    CreateConversationRequestTopicResourceType = "catalog_product"
	CreateConversationRequestTopicResourceTypeCatalogProperty                   CreateConversationRequestTopicResourceType = "catalog_property"
	CreateConversationRequestTopicResourceTypeCatalogAttribute                  CreateConversationRequestTopicResourceType = "catalog_attribute"
	CreateConversationRequestTopicResourceTypeDcLocation                        CreateConversationRequestTopicResourceType = "dc_location"
	CreateConversationRequestTopicResourceTypeEdiRun                            CreateConversationRequestTopicResourceType = "edi_run"
	CreateConversationRequestTopicResourceTypeInventoryItem                     CreateConversationRequestTopicResourceType = "inventory_item"
	CreateConversationRequestTopicResourceTypeAnalyzeWeeksOfSalesResponse       CreateConversationRequestTopicResourceType = "analyze_weeks_of_sales_response"
	CreateConversationRequestTopicResourceTypeBulkReconcileItemsResponse        CreateConversationRequestTopicResourceType = "bulk_reconcile_items_response"
	CreateConversationRequestTopicResourceTypeSysProperty                       CreateConversationRequestTopicResourceType = "sys_property"
	CreateConversationRequestTopicResourceTypeSysPropertyType                   CreateConversationRequestTopicResourceType = "sys_property_type"
	CreateConversationRequestTopicResourceTypeSysPropertyValue                  CreateConversationRequestTopicResourceType = "sys_property_value"
	CreateConversationRequestTopicResourceTypeTerritory                         CreateConversationRequestTopicResourceType = "territory"
	CreateConversationRequestTopicResourceTypeTenancy                           CreateConversationRequestTopicResourceType = "tenancy"
	CreateConversationRequestTopicResourceTypeCheckoutSession                   CreateConversationRequestTopicResourceType = "checkout_session"
	CreateConversationRequestTopicResourceTypeEstimateRateResult                CreateConversationRequestTopicResourceType = "estimate_rate_result"
	CreateConversationRequestTopicResourceTypeRateShopOption                    CreateConversationRequestTopicResourceType = "rate_shop_option"
	CreateConversationRequestTopicResourceTypeRateShopResult                    CreateConversationRequestTopicResourceType = "rate_shop_result"
	CreateConversationRequestTopicResourceTypeOwner                             CreateConversationRequestTopicResourceType = "owner"
	CreateConversationRequestTopicResourceTypeCreatedBy                         CreateConversationRequestTopicResourceType = "created_by"
	CreateConversationRequestTopicResourceTypeMessage                           CreateConversationRequestTopicResourceType = "message"
	CreateConversationRequestTopicResourceTypeAccountPhotoUploadResult          CreateConversationRequestTopicResourceType = "account_photo_upload_result"
	CreateConversationRequestTopicResourceTypeUserPhotoUploadResult             CreateConversationRequestTopicResourceType = "user_photo_upload_result"
	CreateConversationRequestTopicResourceTypeUserPhotoURL                      CreateConversationRequestTopicResourceType = "user_photo_url"
	CreateConversationRequestTopicResourceTypeBatchLot                          CreateConversationRequestTopicResourceType = "batch_lot"
	CreateConversationRequestTopicResourceTypeCheckDuplicateResult              CreateConversationRequestTopicResourceType = "check_duplicate_result"
	CreateConversationRequestTopicResourceTypeItemTrendPoint                    CreateConversationRequestTopicResourceType = "item_trend_point"
	CreateConversationRequestTopicResourceTypePackPickResponse                  CreateConversationRequestTopicResourceType = "pack_pick_response"
	CreateConversationRequestTopicResourceTypePickShipmentsResponse             CreateConversationRequestTopicResourceType = "pick_shipments_response"
	CreateConversationRequestTopicResourceTypeTenancyPendingRegistration        CreateConversationRequestTopicResourceType = "tenancy_pending_registration"
	CreateConversationRequestTopicResourceTypeInvoiceAllocationEntry            CreateConversationRequestTopicResourceType = "invoice_allocation_entry"
	CreateConversationRequestTopicResourceTypeAllocationCustomer                CreateConversationRequestTopicResourceType = "allocation_customer"
	CreateConversationRequestTopicResourceTypeCheckoutSalesOrderResponse        CreateConversationRequestTopicResourceType = "checkout_sales_order_response"
	CreateConversationRequestTopicResourceTypeCreateProductionRunResponse       CreateConversationRequestTopicResourceType = "create_production_run_response"
	CreateConversationRequestTopicResourceTypeSalesOrderPriceQuote              CreateConversationRequestTopicResourceType = "sales_order_price_quote"
	CreateConversationRequestTopicResourceTypeHubspotSyncJob                    CreateConversationRequestTopicResourceType = "hubspot_sync_job"
	CreateConversationRequestTopicResourceTypeHubspotSyncReport                 CreateConversationRequestTopicResourceType = "hubspot_sync_report"
	CreateConversationRequestTopicResourceTypeHubspotCompanyReview              CreateConversationRequestTopicResourceType = "hubspot_company_review"
	CreateConversationRequestTopicResourceTypeHubspotCompanyCandidate           CreateConversationRequestTopicResourceType = "hubspot_company_candidate"
	CreateConversationRequestTopicResourceTypeContactMatch                      CreateConversationRequestTopicResourceType = "contact_match"
	CreateConversationRequestTopicResourceTypeReplyDraft                        CreateConversationRequestTopicResourceType = "reply_draft"
	CreateConversationRequestTopicResourceTypeConversationLink                  CreateConversationRequestTopicResourceType = "conversation_link"
	CreateConversationRequestTopicResourceTypeMessagingGroup                    CreateConversationRequestTopicResourceType = "messaging_group"
	CreateConversationRequestTopicResourceTypeMessagingGroupMember              CreateConversationRequestTopicResourceType = "messaging_group_member"
)

// List represents a paginated list of resources.
type ListAgentAction struct {
	// Resources in this page.
	Data []AgentAction `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentActionObject `json:"object" api:"required"`
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
func (r ListAgentAction) RawJSON() string { return r.JSON.raw }
func (r *ListAgentAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentActionObject string

const (
	ListAgentActionObjectList ListAgentActionObject = "list"
)

// List represents a paginated list of resources.
type ListAgentDefinitionTool struct {
	// Resources in this page.
	Data []AgentDefinitionTool `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentDefinitionToolObject `json:"object" api:"required"`
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
func (r ListAgentDefinitionTool) RawJSON() string { return r.JSON.raw }
func (r *ListAgentDefinitionTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentDefinitionToolObject string

const (
	ListAgentDefinitionToolObjectList ListAgentDefinitionToolObject = "list"
)

// List represents a paginated list of resources.
type ListAgentRunStep struct {
	// Resources in this page.
	Data []AgentRunStep `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentRunStepObject `json:"object" api:"required"`
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
func (r ListAgentRunStep) RawJSON() string { return r.JSON.raw }
func (r *ListAgentRunStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentRunStepObject string

const (
	ListAgentRunStepObjectList ListAgentRunStepObject = "list"
)

// List represents a paginated list of resources.
type ListConversation struct {
	// Resources in this page.
	Data []Conversation `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListConversationObject `json:"object" api:"required"`
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
func (r ListConversation) RawJSON() string { return r.JSON.raw }
func (r *ListConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListConversationObject string

const (
	ListConversationObjectList ListConversationObject = "list"
)

// List represents a paginated list of resources.
type ListConversationParticipant struct {
	// Resources in this page.
	Data []ConversationParticipant `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListConversationParticipantObject `json:"object" api:"required"`
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
func (r ListConversationParticipant) RawJSON() string { return r.JSON.raw }
func (r *ListConversationParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListConversationParticipantObject string

const (
	ListConversationParticipantObjectList ListConversationParticipantObject = "list"
)

// List represents a paginated list of resources.
type ListMessageAttachment struct {
	// Resources in this page.
	Data []MessageAttachment `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessageAttachmentObject `json:"object" api:"required"`
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
func (r ListMessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *ListMessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessageAttachmentObject string

const (
	ListMessageAttachmentObjectList ListMessageAttachmentObject = "list"
)

// List represents a paginated list of resources.
type ListMessagingGroupMember struct {
	// Resources in this page.
	Data []MessagingGroupMember `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessagingGroupMemberObject `json:"object" api:"required"`
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
func (r ListMessagingGroupMember) RawJSON() string { return r.JSON.raw }
func (r *ListMessagingGroupMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessagingGroupMemberObject string

const (
	ListMessagingGroupMemberObjectList ListMessagingGroupMemberObject = "list"
)

// A chat message within a conversation.
type Message struct {
	// Message ID.
	ID string `json:"id" api:"required"`
	// A single execution of an agent, from trigger through completion.
	AgentRun AgentRun `json:"agent_run" api:"required"`
	// List represents a paginated list of resources.
	Attachments ListMessageAttachment `json:"attachments" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Author Actor `json:"author" api:"required"`
	// Message body.
	//
	// `null` for templated or deleted messages.
	Body string `json:"body" api:"required"`
	// How the message was delivered (or, for a draft, how it will be on approve).
	//
	// - `message`: delivered as an in-conversation chat message.
	// - `email`: delivered as email through the conversation's bridged inbox.
	//
	// Any of "message", "email".
	Channel MessageChannel `json:"channel" api:"required"`
	// The client-supplied dedupe key echoed back for optimistic-UI reconciliation.
	//
	// `null` for server-generated messages.
	ClientMessageID string `json:"client_message_id" api:"required"`
	// A conversation thread the caller participates in.
	Conversation *Conversation `json:"conversation" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the message was deleted (tombstone).
	DeletedAt time.Time `json:"deleted_at" api:"required" format:"date-time"`
	// When the message was last edited.
	EditedAt time.Time `json:"edited_at" api:"required" format:"date-time"`
	// The kind of message.
	//
	//   - `chat`: a user-authored chat message.
	//   - `system_event`: a system-generated event message.
	//   - `agent`: a message authored by an AI agent participant.
	//   - `scheduled`: a message materialized from a scheduled send.
	//   - `alert`: a system or producer alert rendered as a message.
	//   - `email`: an inbound email materialized into the conversation by the email
	//     bridge.
	//
	// Any of "chat", "system_event", "agent", "scheduled", "alert", "email".
	Kind MessageKind `json:"kind" api:"required"`
	// Resource type identifier.
	//
	// Any of "chat_message".
	Object MessageObject `json:"object" api:"required"`
	// A chat message within a conversation.
	ReplyTo *Message `json:"reply_to" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// When a `scheduled` message is due to be delivered.
	ScheduledAt time.Time `json:"scheduled_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Sender Actor `json:"sender" api:"required"`
	// Monotonic per-conversation ordering sequence.
	Sequence int64 `json:"sequence" api:"required"`
	// The lifecycle state of the message.
	//
	//   - `draft`: an editable customer-reply draft awaiting approval; not in the
	//     timeline.
	//   - `scheduled`: queued for delivery at a future time; not yet in the timeline.
	//   - `sent`: a delivered timeline message; only `sent` messages carry a `sequence`.
	//   - `canceled`: a scheduled message canceled before delivery.
	//   - `rejected`: a draft discarded without sending.
	//   - `failed`: a scheduled message that exhausted delivery attempts.
	//   - `superseded`: a draft replaced by a newer one for the same source thread.
	//
	// Any of "draft", "scheduled", "sent", "canceled", "rejected", "failed",
	// "superseded".
	Status MessageStatus `json:"status" api:"required"`
	// The streaming state of a reply.
	//
	// `streaming` while the body is still being generated (it fills in via realtime
	// updates); `complete` once finalized.
	//
	// `null` for ordinary messages.
	StreamingState string `json:"streaming_state" api:"required"`
	// The email subject of a customer-reply `draft` on an email-bridged case.
	Subject string `json:"subject" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Who can see this message.
	//
	//   - `internal`: a team-only note.
	//   - `external`: sent to or received from an external party (e.g. the customer on a
	//     support case).
	//   - `system`: an event shown to both the team and the customer.
	//
	// On a customer-facing conversation, customer payloads only ever carry `external`
	// and `system` messages.
	//
	// Any of "internal", "external", "system".
	Visibility MessageVisibility `json:"visibility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AgentRun        respjson.Field
		Attachments     respjson.Field
		Author          respjson.Field
		Body            respjson.Field
		Channel         respjson.Field
		ClientMessageID respjson.Field
		Conversation    respjson.Field
		CreatedAt       respjson.Field
		DeletedAt       respjson.Field
		EditedAt        respjson.Field
		Kind            respjson.Field
		Object          respjson.Field
		ReplyTo         respjson.Field
		Resource        respjson.Field
		ScheduledAt     respjson.Field
		Sender          respjson.Field
		Sequence        respjson.Field
		Status          respjson.Field
		StreamingState  respjson.Field
		Subject         respjson.Field
		UpdatedAt       respjson.Field
		Visibility      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the message was delivered (or, for a draft, how it will be on approve).
//
// - `message`: delivered as an in-conversation chat message.
// - `email`: delivered as email through the conversation's bridged inbox.
type MessageChannel string

const (
	MessageChannelMessage MessageChannel = "message"
	MessageChannelEmail   MessageChannel = "email"
)

// The kind of message.
//
//   - `chat`: a user-authored chat message.
//   - `system_event`: a system-generated event message.
//   - `agent`: a message authored by an AI agent participant.
//   - `scheduled`: a message materialized from a scheduled send.
//   - `alert`: a system or producer alert rendered as a message.
//   - `email`: an inbound email materialized into the conversation by the email
//     bridge.
type MessageKind string

const (
	MessageKindChat        MessageKind = "chat"
	MessageKindSystemEvent MessageKind = "system_event"
	MessageKindAgent       MessageKind = "agent"
	MessageKindScheduled   MessageKind = "scheduled"
	MessageKindAlert       MessageKind = "alert"
	MessageKindEmail       MessageKind = "email"
)

// Resource type identifier.
type MessageObject string

const (
	MessageObjectChatMessage MessageObject = "chat_message"
)

// The lifecycle state of the message.
//
//   - `draft`: an editable customer-reply draft awaiting approval; not in the
//     timeline.
//   - `scheduled`: queued for delivery at a future time; not yet in the timeline.
//   - `sent`: a delivered timeline message; only `sent` messages carry a `sequence`.
//   - `canceled`: a scheduled message canceled before delivery.
//   - `rejected`: a draft discarded without sending.
//   - `failed`: a scheduled message that exhausted delivery attempts.
//   - `superseded`: a draft replaced by a newer one for the same source thread.
type MessageStatus string

const (
	MessageStatusDraft      MessageStatus = "draft"
	MessageStatusScheduled  MessageStatus = "scheduled"
	MessageStatusSent       MessageStatus = "sent"
	MessageStatusCanceled   MessageStatus = "canceled"
	MessageStatusRejected   MessageStatus = "rejected"
	MessageStatusFailed     MessageStatus = "failed"
	MessageStatusSuperseded MessageStatus = "superseded"
)

// Who can see this message.
//
//   - `internal`: a team-only note.
//   - `external`: sent to or received from an external party (e.g. the customer on a
//     support case).
//   - `system`: an event shown to both the team and the customer.
//
// On a customer-facing conversation, customer payloads only ever carry `external`
// and `system` messages.
type MessageVisibility string

const (
	MessageVisibilityInternal MessageVisibility = "internal"
	MessageVisibilityExternal MessageVisibility = "external"
	MessageVisibilitySystem   MessageVisibility = "system"
)

// A file, image, link, or resource attached to a message.
type MessageAttachment struct {
	// Attachment ID.
	ID string `json:"id" api:"required"`
	// The MIME content type for uploaded attachments.
	//
	// `null` for link/resource attachments.
	ContentType string `json:"content_type" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The original filename for uploaded attachments.
	//
	// `null` for link/resource attachments.
	Filename string `json:"filename" api:"required"`
	// The kind of attachment, which determines how it is stored and which of the
	// fields below are populated.
	//
	// - `file`: an uploaded non-image file.
	// - `image`: an uploaded image.
	// - `link`: an external URL reference, with no stored file.
	// - `resource`: a reference to an in-app resource, such as an order.
	//
	// Any of "file", "image", "link", "resource".
	Kind MessageAttachmentKind `json:"kind" api:"required"`
	// Resource type identifier.
	//
	// Any of "message_attachment".
	Object MessageAttachmentObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// The size in bytes for uploaded attachments.
	//
	// `null` when unknown or for link/resource attachments.
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// A time-limited download URL for uploaded (file/image) attachments, or the link
	// URL.
	//
	// `null` for resource attachments.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		CreatedAt   respjson.Field
		Filename    respjson.Field
		Kind        respjson.Field
		Object      respjson.Field
		Resource    respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *MessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of attachment, which determines how it is stored and which of the
// fields below are populated.
//
// - `file`: an uploaded non-image file.
// - `image`: an uploaded image.
// - `link`: an external URL reference, with no stored file.
// - `resource`: a reference to an in-app resource, such as an order.
type MessageAttachmentKind string

const (
	MessageAttachmentKindFile     MessageAttachmentKind = "file"
	MessageAttachmentKindImage    MessageAttachmentKind = "image"
	MessageAttachmentKindLink     MessageAttachmentKind = "link"
	MessageAttachmentKindResource MessageAttachmentKind = "resource"
)

// Resource type identifier.
type MessageAttachmentObject string

const (
	MessageAttachmentObjectMessageAttachment MessageAttachmentObject = "message_attachment"
)

// A reusable roster: a named set of members (users and/or agents) that seeds new
// conversations.
//
// Starting a conversation from a group snapshots its current members into that
// conversation, so the same group can back many conversations (each with its own
// title); later edits to the group never change conversations already created from
// it.
type MessagingGroup struct {
	// Messaging group ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// List represents a paginated list of resources.
	Members ListMessagingGroupMember `json:"members" api:"required"`
	// The roster's display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "messaging_group".
	Object MessagingGroupObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Members     respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingGroup) RawJSON() string { return r.JSON.raw }
func (r *MessagingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MessagingGroupObject string

const (
	MessagingGroupObjectMessagingGroup MessagingGroupObject = "messaging_group"
)

// A member of a reusable roster: either a user or an agent, represented by its
// actor.
type MessagingGroupMember struct {
	// Membership ID (used to remove the member from the roster).
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// Resource type identifier.
	//
	// Any of "messaging_group_member".
	Object MessagingGroupMemberObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actor       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingGroupMember) RawJSON() string { return r.JSON.raw }
func (r *MessagingGroupMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MessagingGroupMemberObject string

const (
	MessagingGroupMemberObjectMessagingGroupMember MessagingGroupMemberObject = "messaging_group_member"
)

// A participant's read position in a conversation — the basis for read receipts
// ("who has seen this").
type ReadCursor struct {
	// The id of the last message the participant has read.
	//
	// `null` if they have not read any message yet.
	MessageID string `json:"message_id" api:"required"`
	// Resource type identifier.
	//
	// Any of "read_cursor".
	Object ReadCursorObject `json:"object" api:"required"`
	// When the participant last advanced their read cursor.
	//
	// `null` if they have not read any message yet.
	ReadAt time.Time `json:"read_at" api:"required" format:"date-time"`
	// The sequence number of the last message the participant has read in the
	// conversation.
	//
	// A message is "seen" by this participant when its `sequence` is `<=` this value.
	// `0` means they have not read any message in the conversation yet.
	Sequence int64 `json:"sequence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		Object      respjson.Field
		ReadAt      respjson.Field
		Sequence    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReadCursor) RawJSON() string { return r.JSON.raw }
func (r *ReadCursor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ReadCursorObject string

const (
	ReadCursorObjectReadCursor ReadCursorObject = "read_cursor"
)

// Trigger-type-specific configuration.
//
// Which fields are populated depends on the agent's `trigger_type`:
//
// - `scheduled`: `cron_schedule` (and optionally `timezone`) is set.
// - `event`: `event_filters` is set.
// - `manual`: all fields are empty.
type TriggerConfig struct {
	// Cron expression for scheduled triggers (e.g. `0 9 * * *`).
	CronSchedule string `json:"cron_schedule" api:"required"`
	// Event types that trigger this agent (e.g.
	// `["email.received", "order.created"]`).
	EventFilters []string `json:"event_filters" api:"required"`
	// Resource type identifier.
	//
	// Any of "trigger_config".
	Object TriggerConfigObject `json:"object" api:"required"`
	// IANA timezone for the cron schedule (e.g. `America/New_York`).
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CronSchedule respjson.Field
		EventFilters respjson.Field
		Object       respjson.Field
		Timezone     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TriggerConfig) RawJSON() string { return r.JSON.raw }
func (r *TriggerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type TriggerConfigObject string

const (
	TriggerConfigObjectTriggerConfig TriggerConfigObject = "trigger_config"
)

// Request to rename a conversation (owner/admin; groups only).
type UpdateConversationRequestParam struct {
	// New group title.
	//
	// Send `null` to clear the title; omit to leave it unchanged.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r UpdateConversationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateConversationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateConversationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingConversationNewParams struct {
	// Request to create a conversation.
	CreateConversationRequest CreateConversationRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateConversationRequest)
}
func (r *MessagingConversationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationNewParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationGetParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to rename a conversation (owner/admin; groups only).
	UpdateConversationRequest UpdateConversationRequestParam
	paramObj
}

func (r MessagingConversationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateConversationRequest)
}
func (r *MessagingConversationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationUpdateParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationListParams struct {
	// Support inbox: filter to cases owned by this assignee (a user or a team),
	// matched by id.
	AssigneeResourceID param.Opt[string] `query:"assignee_resource_id,omitzero" json:"-"`
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Support inbox: include archived (resolved-and-closed) cases.
	IncludeArchived param.Opt[bool] `query:"include_archived,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// The id of the anchoring business record (with `topic_resource_type`).
	TopicResourceID param.Opt[string] `query:"topic_resource_id,omitzero" json:"-"`
	// Support inbox: restrict to cases with no assignee.
	Unassigned param.Opt[bool] `query:"unassigned,omitzero" json:"-"`
	// Filter by conversation audience direction.
	//
	// Any of "internal", "customer".
	Audience MessagingConversationListParamsAudience `query:"audience,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by conversation visibility.
	//
	// Any of "active", "hidden".
	Status MessagingConversationListParamsStatus `query:"status,omitzero" json:"-"`
	// Restrict to conversations anchored to a business record of this type (with
	// `topic_resource_id`).
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
	TopicResourceType MessagingConversationListParamsTopicResourceType `query:"topic_resource_type,omitzero" json:"-"`
	// Filter by conversation type.
	//
	// Any of "direct_message", "group", "system".
	Type MessagingConversationListParamsType `query:"type,omitzero" json:"-"`
	// Support inbox: filter external cases to a single triage lane.
	//
	// Any of "new", "open", "waiting_internal", "waiting_external", "needs_approval",
	// "resolved".
	WorkflowStatus MessagingConversationListParamsWorkflowStatus `query:"workflow_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationListParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by conversation audience direction.
type MessagingConversationListParamsAudience string

const (
	MessagingConversationListParamsAudienceInternal MessagingConversationListParamsAudience = "internal"
	MessagingConversationListParamsAudienceCustomer MessagingConversationListParamsAudience = "customer"
)

// Filter by conversation visibility.
type MessagingConversationListParamsStatus string

const (
	MessagingConversationListParamsStatusActive MessagingConversationListParamsStatus = "active"
	MessagingConversationListParamsStatusHidden MessagingConversationListParamsStatus = "hidden"
)

// Restrict to conversations anchored to a business record of this type (with
// `topic_resource_id`).
type MessagingConversationListParamsTopicResourceType string

const (
	MessagingConversationListParamsTopicResourceTypeAccount                           MessagingConversationListParamsTopicResourceType = "account"
	MessagingConversationListParamsTopicResourceTypeActor                             MessagingConversationListParamsTopicResourceType = "actor"
	MessagingConversationListParamsTopicResourceTypeEntity                            MessagingConversationListParamsTopicResourceType = "entity"
	MessagingConversationListParamsTopicResourceTypeRecord                            MessagingConversationListParamsTopicResourceType = "record"
	MessagingConversationListParamsTopicResourceTypeFreight                           MessagingConversationListParamsTopicResourceType = "freight"
	MessagingConversationListParamsTopicResourceTypeSalesOrderTotals                  MessagingConversationListParamsTopicResourceType = "sales_order_totals"
	MessagingConversationListParamsTopicResourceTypeSalesOrderRelated                 MessagingConversationListParamsTopicResourceType = "sales_order_related"
	MessagingConversationListParamsTopicResourceTypeOrderContact                      MessagingConversationListParamsTopicResourceType = "order_contact"
	MessagingConversationListParamsTopicResourceTypeUser                              MessagingConversationListParamsTopicResourceType = "user"
	MessagingConversationListParamsTopicResourceTypeAddress                           MessagingConversationListParamsTopicResourceType = "address"
	MessagingConversationListParamsTopicResourceTypeAPIKey                            MessagingConversationListParamsTopicResourceType = "api_key"
	MessagingConversationListParamsTopicResourceTypeCreatedAPIKey                     MessagingConversationListParamsTopicResourceType = "created_api_key"
	MessagingConversationListParamsTopicResourceTypeRefreshToken                      MessagingConversationListParamsTopicResourceType = "refresh_token"
	MessagingConversationListParamsTopicResourceTypeList                              MessagingConversationListParamsTopicResourceType = "list"
	MessagingConversationListParamsTopicResourceTypeSandbox                           MessagingConversationListParamsTopicResourceType = "sandbox"
	MessagingConversationListParamsTopicResourceTypeRegistrationSession               MessagingConversationListParamsTopicResourceType = "registration_session"
	MessagingConversationListParamsTopicResourceTypePricingPlan                       MessagingConversationListParamsTopicResourceType = "pricing_plan"
	MessagingConversationListParamsTopicResourceTypeAccountPlan                       MessagingConversationListParamsTopicResourceType = "account_plan"
	MessagingConversationListParamsTopicResourceTypePlanChange                        MessagingConversationListParamsTopicResourceType = "plan_change"
	MessagingConversationListParamsTopicResourceTypeEnterpriseInquiry                 MessagingConversationListParamsTopicResourceType = "enterprise_inquiry"
	MessagingConversationListParamsTopicResourceTypeRequestLog                        MessagingConversationListParamsTopicResourceType = "request_log"
	MessagingConversationListParamsTopicResourceTypeAuditEvent                        MessagingConversationListParamsTopicResourceType = "audit_event"
	MessagingConversationListParamsTopicResourceTypeAuditFieldChange                  MessagingConversationListParamsTopicResourceType = "audit_field_change"
	MessagingConversationListParamsTopicResourceTypeRole                              MessagingConversationListParamsTopicResourceType = "role"
	MessagingConversationListParamsTopicResourceTypeUnit                              MessagingConversationListParamsTopicResourceType = "unit"
	MessagingConversationListParamsTopicResourceTypeAccountAffiliation                MessagingConversationListParamsTopicResourceType = "account_affiliation"
	MessagingConversationListParamsTopicResourceTypeAgentDefinition                   MessagingConversationListParamsTopicResourceType = "agent_definition"
	MessagingConversationListParamsTopicResourceTypeAvailableTool                     MessagingConversationListParamsTopicResourceType = "available_tool"
	MessagingConversationListParamsTopicResourceTypeAgentDefinitionTool               MessagingConversationListParamsTopicResourceType = "agent_definition_tool"
	MessagingConversationListParamsTopicResourceTypeAgentAccountStatus                MessagingConversationListParamsTopicResourceType = "agent_account_status"
	MessagingConversationListParamsTopicResourceTypeAgentRun                          MessagingConversationListParamsTopicResourceType = "agent_run"
	MessagingConversationListParamsTopicResourceTypeAgentAction                       MessagingConversationListParamsTopicResourceType = "agent_action"
	MessagingConversationListParamsTopicResourceTypeAgentRunStep                      MessagingConversationListParamsTopicResourceType = "agent_run_step"
	MessagingConversationListParamsTopicResourceTypeAgentTokenUsage                   MessagingConversationListParamsTopicResourceType = "agent_token_usage"
	MessagingConversationListParamsTopicResourceTypeAgentMemory                       MessagingConversationListParamsTopicResourceType = "agent_memory"
	MessagingConversationListParamsTopicResourceTypeNotification                      MessagingConversationListParamsTopicResourceType = "notification"
	MessagingConversationListParamsTopicResourceTypeNotificationUnreadCount           MessagingConversationListParamsTopicResourceType = "notification_unread_count"
	MessagingConversationListParamsTopicResourceTypeNotificationSendResult            MessagingConversationListParamsTopicResourceType = "notification_send_result"
	MessagingConversationListParamsTopicResourceTypeNotificationUnreadSummary         MessagingConversationListParamsTopicResourceType = "notification_unread_summary"
	MessagingConversationListParamsTopicResourceTypeAnnouncement                      MessagingConversationListParamsTopicResourceType = "announcement"
	MessagingConversationListParamsTopicResourceTypeConversation                      MessagingConversationListParamsTopicResourceType = "conversation"
	MessagingConversationListParamsTopicResourceTypeConversationParticipant           MessagingConversationListParamsTopicResourceType = "conversation_participant"
	MessagingConversationListParamsTopicResourceTypeReadCursor                        MessagingConversationListParamsTopicResourceType = "read_cursor"
	MessagingConversationListParamsTopicResourceTypeChatMessage                       MessagingConversationListParamsTopicResourceType = "chat_message"
	MessagingConversationListParamsTopicResourceTypeNotificationUnreadSummaryAccount  MessagingConversationListParamsTopicResourceType = "notification_unread_summary_account"
	MessagingConversationListParamsTopicResourceTypeMessagingBlock                    MessagingConversationListParamsTopicResourceType = "messaging_block"
	MessagingConversationListParamsTopicResourceTypeNotificationPreference            MessagingConversationListParamsTopicResourceType = "notification_preference"
	MessagingConversationListParamsTopicResourceTypeMessageAttachment                 MessagingConversationListParamsTopicResourceType = "message_attachment"
	MessagingConversationListParamsTopicResourceTypeAttachmentUploadTarget            MessagingConversationListParamsTopicResourceType = "attachment_upload_target"
	MessagingConversationListParamsTopicResourceTypeScheduledMessage                  MessagingConversationListParamsTopicResourceType = "scheduled_message"
	MessagingConversationListParamsTopicResourceTypeMessagingContact                  MessagingConversationListParamsTopicResourceType = "messaging_contact"
	MessagingConversationListParamsTopicResourceTypeMessageReport                     MessagingConversationListParamsTopicResourceType = "message_report"
	MessagingConversationListParamsTopicResourceTypeToolGroup                         MessagingConversationListParamsTopicResourceType = "tool_group"
	MessagingConversationListParamsTopicResourceTypeModel                             MessagingConversationListParamsTopicResourceType = "model"
	MessagingConversationListParamsTopicResourceTypePaymentTerm                       MessagingConversationListParamsTopicResourceType = "payment_term"
	MessagingConversationListParamsTopicResourceTypeShippingTerm                      MessagingConversationListParamsTopicResourceType = "shipping_term"
	MessagingConversationListParamsTopicResourceTypeQuantity                          MessagingConversationListParamsTopicResourceType = "quantity"
	MessagingConversationListParamsTopicResourceTypeAccountGroup                      MessagingConversationListParamsTopicResourceType = "account_group"
	MessagingConversationListParamsTopicResourceTypeSupportRoute                      MessagingConversationListParamsTopicResourceType = "support_route"
	MessagingConversationListParamsTopicResourceTypeSupportAvailability               MessagingConversationListParamsTopicResourceType = "support_availability"
	MessagingConversationListParamsTopicResourceTypeAccountStatus                     MessagingConversationListParamsTopicResourceType = "account_status"
	MessagingConversationListParamsTopicResourceTypeGeolocation                       MessagingConversationListParamsTopicResourceType = "geolocation"
	MessagingConversationListParamsTopicResourceTypeAccountUser                       MessagingConversationListParamsTopicResourceType = "account_user"
	MessagingConversationListParamsTopicResourceTypeDepartment                        MessagingConversationListParamsTopicResourceType = "department"
	MessagingConversationListParamsTopicResourceTypeAccountIntegration                MessagingConversationListParamsTopicResourceType = "account_integration"
	MessagingConversationListParamsTopicResourceTypeAccountPrice                      MessagingConversationListParamsTopicResourceType = "account_price"
	MessagingConversationListParamsTopicResourceTypeProductLine                       MessagingConversationListParamsTopicResourceType = "product_line"
	MessagingConversationListParamsTopicResourceTypeItemCategory                      MessagingConversationListParamsTopicResourceType = "item_category"
	MessagingConversationListParamsTopicResourceTypeAttribute                         MessagingConversationListParamsTopicResourceType = "attribute"
	MessagingConversationListParamsTopicResourceTypeRate                              MessagingConversationListParamsTopicResourceType = "rate"
	MessagingConversationListParamsTopicResourceTypeAccountGroupProductLineAccess     MessagingConversationListParamsTopicResourceType = "account_group_product_line_access"
	MessagingConversationListParamsTopicResourceTypeSalesTarget                       MessagingConversationListParamsTopicResourceType = "sales_target"
	MessagingConversationListParamsTopicResourceTypeAdjustmentType                    MessagingConversationListParamsTopicResourceType = "adjustment_type"
	MessagingConversationListParamsTopicResourceTypeAccountBranding                   MessagingConversationListParamsTopicResourceType = "account_branding"
	MessagingConversationListParamsTopicResourceTypeAccountPortal                     MessagingConversationListParamsTopicResourceType = "account_portal"
	MessagingConversationListParamsTopicResourceTypeAccountLogoURL                    MessagingConversationListParamsTopicResourceType = "account_logo_url"
	MessagingConversationListParamsTopicResourceTypePublicAccount                     MessagingConversationListParamsTopicResourceType = "public_account"
	MessagingConversationListParamsTopicResourceTypeProperty                          MessagingConversationListParamsTopicResourceType = "property"
	MessagingConversationListParamsTopicResourceTypeCarrier                           MessagingConversationListParamsTopicResourceType = "carrier"
	MessagingConversationListParamsTopicResourceTypeServiceLevel                      MessagingConversationListParamsTopicResourceType = "service_level"
	MessagingConversationListParamsTopicResourceTypeItem                              MessagingConversationListParamsTopicResourceType = "item"
	MessagingConversationListParamsTopicResourceTypeItemInventory                     MessagingConversationListParamsTopicResourceType = "item_inventory"
	MessagingConversationListParamsTopicResourceTypeProduct                           MessagingConversationListParamsTopicResourceType = "product"
	MessagingConversationListParamsTopicResourceTypeBatch                             MessagingConversationListParamsTopicResourceType = "batch"
	MessagingConversationListParamsTopicResourceTypeBatchFlowNode                     MessagingConversationListParamsTopicResourceType = "batch_flow_node"
	MessagingConversationListParamsTopicResourceTypeScanningConsumption               MessagingConversationListParamsTopicResourceType = "scanning_consumption"
	MessagingConversationListParamsTopicResourceTypeOpenBatchSummary                  MessagingConversationListParamsTopicResourceType = "open_batch_summary"
	MessagingConversationListParamsTopicResourceTypeScanningProductionStepInfo        MessagingConversationListParamsTopicResourceType = "scanning_production_step_info"
	MessagingConversationListParamsTopicResourceTypeScanningStation                   MessagingConversationListParamsTopicResourceType = "scanning_station"
	MessagingConversationListParamsTopicResourceTypeProductionStep                    MessagingConversationListParamsTopicResourceType = "production_step"
	MessagingConversationListParamsTopicResourceTypeProductionRun                     MessagingConversationListParamsTopicResourceType = "production_run"
	MessagingConversationListParamsTopicResourceTypeMachine                           MessagingConversationListParamsTopicResourceType = "machine"
	MessagingConversationListParamsTopicResourceTypeChildAccount                      MessagingConversationListParamsTopicResourceType = "child_account"
	MessagingConversationListParamsTopicResourceTypeUnitGroup                         MessagingConversationListParamsTopicResourceType = "unit_group"
	MessagingConversationListParamsTopicResourceTypeUnitGroupUnit                     MessagingConversationListParamsTopicResourceType = "unit_group_unit"
	MessagingConversationListParamsTopicResourceTypeConsumption                       MessagingConversationListParamsTopicResourceType = "consumption"
	MessagingConversationListParamsTopicResourceTypeCustomerProductLineAccess         MessagingConversationListParamsTopicResourceType = "customer_product_line_access"
	MessagingConversationListParamsTopicResourceTypeCustomer                          MessagingConversationListParamsTopicResourceType = "customer"
	MessagingConversationListParamsTopicResourceTypeFrequentlyOrderedProduct          MessagingConversationListParamsTopicResourceType = "frequently_ordered_product"
	MessagingConversationListParamsTopicResourceTypePriority                          MessagingConversationListParamsTopicResourceType = "priority"
	MessagingConversationListParamsTopicResourceTypeDelivery                          MessagingConversationListParamsTopicResourceType = "delivery"
	MessagingConversationListParamsTopicResourceTypeDeliveryLine                      MessagingConversationListParamsTopicResourceType = "delivery_line"
	MessagingConversationListParamsTopicResourceTypeSalesOrder                        MessagingConversationListParamsTopicResourceType = "sales_order"
	MessagingConversationListParamsTopicResourceTypeLocation                          MessagingConversationListParamsTopicResourceType = "location"
	MessagingConversationListParamsTopicResourceTypeLocationType                      MessagingConversationListParamsTopicResourceType = "location_type"
	MessagingConversationListParamsTopicResourceTypeLot                               MessagingConversationListParamsTopicResourceType = "lot"
	MessagingConversationListParamsTopicResourceTypeEmailLog                          MessagingConversationListParamsTopicResourceType = "email_log"
	MessagingConversationListParamsTopicResourceTypeEmailDomain                       MessagingConversationListParamsTopicResourceType = "email_domain"
	MessagingConversationListParamsTopicResourceTypeEmailInbox                        MessagingConversationListParamsTopicResourceType = "email_inbox"
	MessagingConversationListParamsTopicResourceTypeInventoryChangeLog                MessagingConversationListParamsTopicResourceType = "inventory_change_log"
	MessagingConversationListParamsTopicResourceTypeInvoice                           MessagingConversationListParamsTopicResourceType = "invoice"
	MessagingConversationListParamsTopicResourceTypeInvoiceSummary                    MessagingConversationListParamsTopicResourceType = "invoice_summary"
	MessagingConversationListParamsTopicResourceTypeInvoiceLine                       MessagingConversationListParamsTopicResourceType = "invoice_line"
	MessagingConversationListParamsTopicResourceTypeInvoiceAllocation                 MessagingConversationListParamsTopicResourceType = "invoice_allocation"
	MessagingConversationListParamsTopicResourceTypeInvoiceForPayment                 MessagingConversationListParamsTopicResourceType = "invoice_for_payment"
	MessagingConversationListParamsTopicResourceTypeShipment                          MessagingConversationListParamsTopicResourceType = "shipment"
	MessagingConversationListParamsTopicResourceTypeShipmentSummary                   MessagingConversationListParamsTopicResourceType = "shipment_summary"
	MessagingConversationListParamsTopicResourceTypeShipmentLine                      MessagingConversationListParamsTopicResourceType = "shipment_line"
	MessagingConversationListParamsTopicResourceTypeShippingCase                      MessagingConversationListParamsTopicResourceType = "shipping_case"
	MessagingConversationListParamsTopicResourceTypeShippingCaseLabelURL              MessagingConversationListParamsTopicResourceType = "shipping_case_label_url"
	MessagingConversationListParamsTopicResourceTypeSettlement                        MessagingConversationListParamsTopicResourceType = "settlement"
	MessagingConversationListParamsTopicResourceTypeSettlementSummary                 MessagingConversationListParamsTopicResourceType = "settlement_summary"
	MessagingConversationListParamsTopicResourceTypeRolePermission                    MessagingConversationListParamsTopicResourceType = "role_permission"
	MessagingConversationListParamsTopicResourceTypeRegistrationFlow                  MessagingConversationListParamsTopicResourceType = "registration_flow"
	MessagingConversationListParamsTopicResourceTypeRegistrationFlowOption            MessagingConversationListParamsTopicResourceType = "registration_flow_option"
	MessagingConversationListParamsTopicResourceTypeTransaction                       MessagingConversationListParamsTopicResourceType = "transaction"
	MessagingConversationListParamsTopicResourceTypeTransactionSummary                MessagingConversationListParamsTopicResourceType = "transaction_summary"
	MessagingConversationListParamsTopicResourceTypeTransactionMethod                 MessagingConversationListParamsTopicResourceType = "transaction_method"
	MessagingConversationListParamsTopicResourceTypeTransactionType                   MessagingConversationListParamsTopicResourceType = "transaction_type"
	MessagingConversationListParamsTopicResourceTypeTransactionAllocation             MessagingConversationListParamsTopicResourceType = "transaction_allocation"
	MessagingConversationListParamsTopicResourceTypeUsageItem                         MessagingConversationListParamsTopicResourceType = "usage_item"
	MessagingConversationListParamsTopicResourceTypeAccountUsageResponse              MessagingConversationListParamsTopicResourceType = "account_usage_response"
	MessagingConversationListParamsTopicResourceTypeSubscriptionInfo                  MessagingConversationListParamsTopicResourceType = "subscription_info"
	MessagingConversationListParamsTopicResourceTypeBillingPortalSessionResponse      MessagingConversationListParamsTopicResourceType = "billing_portal_session_response"
	MessagingConversationListParamsTopicResourceTypeSwitchPlanResponse                MessagingConversationListParamsTopicResourceType = "switch_plan_response"
	MessagingConversationListParamsTopicResourceTypeEnsureBillingCustomerResponse     MessagingConversationListParamsTopicResourceType = "ensure_billing_customer_response"
	MessagingConversationListParamsTopicResourceTypeSpendingCapResponse               MessagingConversationListParamsTopicResourceType = "spending_cap_response"
	MessagingConversationListParamsTopicResourceTypeAgentSpendInfo                    MessagingConversationListParamsTopicResourceType = "agent_spend_info"
	MessagingConversationListParamsTopicResourceTypeWebhookResponse                   MessagingConversationListParamsTopicResourceType = "webhook_response"
	MessagingConversationListParamsTopicResourceTypeAddressSuggestion                 MessagingConversationListParamsTopicResourceType = "address_suggestion"
	MessagingConversationListParamsTopicResourceTypeAddressComponents                 MessagingConversationListParamsTopicResourceType = "address_components"
	MessagingConversationListParamsTopicResourceTypeAddressDetailsResult              MessagingConversationListParamsTopicResourceType = "address_details_result"
	MessagingConversationListParamsTopicResourceTypeValidatedAddress                  MessagingConversationListParamsTopicResourceType = "validated_address"
	MessagingConversationListParamsTopicResourceTypePlanLimit                         MessagingConversationListParamsTopicResourceType = "plan_limit"
	MessagingConversationListParamsTopicResourceTypePlanChangeProration               MessagingConversationListParamsTopicResourceType = "plan_change_proration"
	MessagingConversationListParamsTopicResourceTypePlanChangeLineItem                MessagingConversationListParamsTopicResourceType = "plan_change_line_item"
	MessagingConversationListParamsTopicResourceTypeSetupBillingResponse              MessagingConversationListParamsTopicResourceType = "setup_billing_response"
	MessagingConversationListParamsTopicResourceTypeConfirmPaymentResponse            MessagingConversationListParamsTopicResourceType = "confirm_payment_response"
	MessagingConversationListParamsTopicResourceTypeOAuthResponse                     MessagingConversationListParamsTopicResourceType = "oauth_response"
	MessagingConversationListParamsTopicResourceTypeOAuthStatusResponse               MessagingConversationListParamsTopicResourceType = "oauth_status_response"
	MessagingConversationListParamsTopicResourceTypeStripePublishableKey              MessagingConversationListParamsTopicResourceType = "stripe_publishable_key"
	MessagingConversationListParamsTopicResourceTypeStripeStatus                      MessagingConversationListParamsTopicResourceType = "stripe_status"
	MessagingConversationListParamsTopicResourceTypeHealthcheck                       MessagingConversationListParamsTopicResourceType = "healthcheck"
	MessagingConversationListParamsTopicResourceTypeAgentDefinitionConfig             MessagingConversationListParamsTopicResourceType = "agent_definition_config"
	MessagingConversationListParamsTopicResourceTypeTriggerConfig                     MessagingConversationListParamsTopicResourceType = "trigger_config"
	MessagingConversationListParamsTopicResourceTypeCustomerContactInfo               MessagingConversationListParamsTopicResourceType = "customer_contact_info"
	MessagingConversationListParamsTopicResourceTypeCustomerFreightPreferences        MessagingConversationListParamsTopicResourceType = "customer_freight_preferences"
	MessagingConversationListParamsTopicResourceTypeCustomerDefaults                  MessagingConversationListParamsTopicResourceType = "customer_defaults"
	MessagingConversationListParamsTopicResourceTypeCustomerNotificationPreferences   MessagingConversationListParamsTopicResourceType = "customer_notification_preferences"
	MessagingConversationListParamsTopicResourceTypeOrderDiscount                     MessagingConversationListParamsTopicResourceType = "order_discount"
	MessagingConversationListParamsTopicResourceTypeSalesOrderLine                    MessagingConversationListParamsTopicResourceType = "sales_order_line"
	MessagingConversationListParamsTopicResourceTypeSalesOrderType                    MessagingConversationListParamsTopicResourceType = "sales_order_type"
	MessagingConversationListParamsTopicResourceTypeSalesOrderStatus                  MessagingConversationListParamsTopicResourceType = "sales_order_status"
	MessagingConversationListParamsTopicResourceTypeMaterial                          MessagingConversationListParamsTopicResourceType = "material"
	MessagingConversationListParamsTopicResourceTypeSupplierMaterial                  MessagingConversationListParamsTopicResourceType = "supplier_material"
	MessagingConversationListParamsTopicResourceTypePart                              MessagingConversationListParamsTopicResourceType = "part"
	MessagingConversationListParamsTopicResourceTypePermissionGroup                   MessagingConversationListParamsTopicResourceType = "permission_group"
	MessagingConversationListParamsTopicResourceTypePermission                        MessagingConversationListParamsTopicResourceType = "permission"
	MessagingConversationListParamsTopicResourceTypePick                              MessagingConversationListParamsTopicResourceType = "pick"
	MessagingConversationListParamsTopicResourceTypePickLine                          MessagingConversationListParamsTopicResourceType = "pick_line"
	MessagingConversationListParamsTopicResourceTypeProductType                       MessagingConversationListParamsTopicResourceType = "product_type"
	MessagingConversationListParamsTopicResourceTypeProduction                        MessagingConversationListParamsTopicResourceType = "production"
	MessagingConversationListParamsTopicResourceTypeProductionFlow                    MessagingConversationListParamsTopicResourceType = "production_flow"
	MessagingConversationListParamsTopicResourceTypeMap                               MessagingConversationListParamsTopicResourceType = "map"
	MessagingConversationListParamsTopicResourceTypePurchaseOrder                     MessagingConversationListParamsTopicResourceType = "purchase_order"
	MessagingConversationListParamsTopicResourceTypePurchaseOrderLine                 MessagingConversationListParamsTopicResourceType = "purchase_order_line"
	MessagingConversationListParamsTopicResourceTypeSupplier                          MessagingConversationListParamsTopicResourceType = "supplier"
	MessagingConversationListParamsTopicResourceTypeSupplierSummary                   MessagingConversationListParamsTopicResourceType = "supplier_summary"
	MessagingConversationListParamsTopicResourceTypeReceivableEntry                   MessagingConversationListParamsTopicResourceType = "receivable_entry"
	MessagingConversationListParamsTopicResourceTypeReceivingOrder                    MessagingConversationListParamsTopicResourceType = "receiving_order"
	MessagingConversationListParamsTopicResourceTypeReceivingOrderLine                MessagingConversationListParamsTopicResourceType = "receiving_order_line"
	MessagingConversationListParamsTopicResourceTypeEmailContact                      MessagingConversationListParamsTopicResourceType = "email_contact"
	MessagingConversationListParamsTopicResourceTypeAllocationEntry                   MessagingConversationListParamsTopicResourceType = "allocation_entry"
	MessagingConversationListParamsTopicResourceTypeOpenCreditEntry                   MessagingConversationListParamsTopicResourceType = "open_credit_entry"
	MessagingConversationListParamsTopicResourceTypeVolumeDiscount                    MessagingConversationListParamsTopicResourceType = "volume_discount"
	MessagingConversationListParamsTopicResourceTypeVolumeDiscountTier                MessagingConversationListParamsTopicResourceType = "volume_discount_tier"
	MessagingConversationListParamsTopicResourceTypeAnalyzeDeliveriesResponse         MessagingConversationListParamsTopicResourceType = "analyze_deliveries_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeManufacturingResponse      MessagingConversationListParamsTopicResourceType = "analyze_manufacturing_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeManufacturingBatchResponse MessagingConversationListParamsTopicResourceType = "analyze_manufacturing_batch_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeQuarterlyOrdersResponse    MessagingConversationListParamsTopicResourceType = "analyze_quarterly_orders_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeNewCustomersResponse       MessagingConversationListParamsTopicResourceType = "analyze_new_customers_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeOeeResponse                MessagingConversationListParamsTopicResourceType = "analyze_oee_response"
	MessagingConversationListParamsTopicResourceTypeCatalogProductLine                MessagingConversationListParamsTopicResourceType = "catalog_product_line"
	MessagingConversationListParamsTopicResourceTypeCatalogCategory                   MessagingConversationListParamsTopicResourceType = "catalog_category"
	MessagingConversationListParamsTopicResourceTypeCatalogProduct                    MessagingConversationListParamsTopicResourceType = "catalog_product"
	MessagingConversationListParamsTopicResourceTypeCatalogProperty                   MessagingConversationListParamsTopicResourceType = "catalog_property"
	MessagingConversationListParamsTopicResourceTypeCatalogAttribute                  MessagingConversationListParamsTopicResourceType = "catalog_attribute"
	MessagingConversationListParamsTopicResourceTypeDcLocation                        MessagingConversationListParamsTopicResourceType = "dc_location"
	MessagingConversationListParamsTopicResourceTypeEdiRun                            MessagingConversationListParamsTopicResourceType = "edi_run"
	MessagingConversationListParamsTopicResourceTypeInventoryItem                     MessagingConversationListParamsTopicResourceType = "inventory_item"
	MessagingConversationListParamsTopicResourceTypeAnalyzeWeeksOfSalesResponse       MessagingConversationListParamsTopicResourceType = "analyze_weeks_of_sales_response"
	MessagingConversationListParamsTopicResourceTypeBulkReconcileItemsResponse        MessagingConversationListParamsTopicResourceType = "bulk_reconcile_items_response"
	MessagingConversationListParamsTopicResourceTypeSysProperty                       MessagingConversationListParamsTopicResourceType = "sys_property"
	MessagingConversationListParamsTopicResourceTypeSysPropertyType                   MessagingConversationListParamsTopicResourceType = "sys_property_type"
	MessagingConversationListParamsTopicResourceTypeSysPropertyValue                  MessagingConversationListParamsTopicResourceType = "sys_property_value"
	MessagingConversationListParamsTopicResourceTypeTerritory                         MessagingConversationListParamsTopicResourceType = "territory"
	MessagingConversationListParamsTopicResourceTypeTenancy                           MessagingConversationListParamsTopicResourceType = "tenancy"
	MessagingConversationListParamsTopicResourceTypeCheckoutSession                   MessagingConversationListParamsTopicResourceType = "checkout_session"
	MessagingConversationListParamsTopicResourceTypeEstimateRateResult                MessagingConversationListParamsTopicResourceType = "estimate_rate_result"
	MessagingConversationListParamsTopicResourceTypeRateShopOption                    MessagingConversationListParamsTopicResourceType = "rate_shop_option"
	MessagingConversationListParamsTopicResourceTypeRateShopResult                    MessagingConversationListParamsTopicResourceType = "rate_shop_result"
	MessagingConversationListParamsTopicResourceTypeOwner                             MessagingConversationListParamsTopicResourceType = "owner"
	MessagingConversationListParamsTopicResourceTypeCreatedBy                         MessagingConversationListParamsTopicResourceType = "created_by"
	MessagingConversationListParamsTopicResourceTypeMessage                           MessagingConversationListParamsTopicResourceType = "message"
	MessagingConversationListParamsTopicResourceTypeAccountPhotoUploadResult          MessagingConversationListParamsTopicResourceType = "account_photo_upload_result"
	MessagingConversationListParamsTopicResourceTypeUserPhotoUploadResult             MessagingConversationListParamsTopicResourceType = "user_photo_upload_result"
	MessagingConversationListParamsTopicResourceTypeUserPhotoURL                      MessagingConversationListParamsTopicResourceType = "user_photo_url"
	MessagingConversationListParamsTopicResourceTypeBatchLot                          MessagingConversationListParamsTopicResourceType = "batch_lot"
	MessagingConversationListParamsTopicResourceTypeCheckDuplicateResult              MessagingConversationListParamsTopicResourceType = "check_duplicate_result"
	MessagingConversationListParamsTopicResourceTypeItemTrendPoint                    MessagingConversationListParamsTopicResourceType = "item_trend_point"
	MessagingConversationListParamsTopicResourceTypePackPickResponse                  MessagingConversationListParamsTopicResourceType = "pack_pick_response"
	MessagingConversationListParamsTopicResourceTypePickShipmentsResponse             MessagingConversationListParamsTopicResourceType = "pick_shipments_response"
	MessagingConversationListParamsTopicResourceTypeTenancyPendingRegistration        MessagingConversationListParamsTopicResourceType = "tenancy_pending_registration"
	MessagingConversationListParamsTopicResourceTypeInvoiceAllocationEntry            MessagingConversationListParamsTopicResourceType = "invoice_allocation_entry"
	MessagingConversationListParamsTopicResourceTypeAllocationCustomer                MessagingConversationListParamsTopicResourceType = "allocation_customer"
	MessagingConversationListParamsTopicResourceTypeCheckoutSalesOrderResponse        MessagingConversationListParamsTopicResourceType = "checkout_sales_order_response"
	MessagingConversationListParamsTopicResourceTypeCreateProductionRunResponse       MessagingConversationListParamsTopicResourceType = "create_production_run_response"
	MessagingConversationListParamsTopicResourceTypeSalesOrderPriceQuote              MessagingConversationListParamsTopicResourceType = "sales_order_price_quote"
	MessagingConversationListParamsTopicResourceTypeHubspotSyncJob                    MessagingConversationListParamsTopicResourceType = "hubspot_sync_job"
	MessagingConversationListParamsTopicResourceTypeHubspotSyncReport                 MessagingConversationListParamsTopicResourceType = "hubspot_sync_report"
	MessagingConversationListParamsTopicResourceTypeHubspotCompanyReview              MessagingConversationListParamsTopicResourceType = "hubspot_company_review"
	MessagingConversationListParamsTopicResourceTypeHubspotCompanyCandidate           MessagingConversationListParamsTopicResourceType = "hubspot_company_candidate"
	MessagingConversationListParamsTopicResourceTypeContactMatch                      MessagingConversationListParamsTopicResourceType = "contact_match"
	MessagingConversationListParamsTopicResourceTypeReplyDraft                        MessagingConversationListParamsTopicResourceType = "reply_draft"
	MessagingConversationListParamsTopicResourceTypeConversationLink                  MessagingConversationListParamsTopicResourceType = "conversation_link"
	MessagingConversationListParamsTopicResourceTypeMessagingGroup                    MessagingConversationListParamsTopicResourceType = "messaging_group"
	MessagingConversationListParamsTopicResourceTypeMessagingGroupMember              MessagingConversationListParamsTopicResourceType = "messaging_group_member"
)

// Filter by conversation type.
type MessagingConversationListParamsType string

const (
	MessagingConversationListParamsTypeDirectMessage MessagingConversationListParamsType = "direct_message"
	MessagingConversationListParamsTypeGroup         MessagingConversationListParamsType = "group"
	MessagingConversationListParamsTypeSystem        MessagingConversationListParamsType = "system"
)

// Support inbox: filter external cases to a single triage lane.
type MessagingConversationListParamsWorkflowStatus string

const (
	MessagingConversationListParamsWorkflowStatusNew             MessagingConversationListParamsWorkflowStatus = "new"
	MessagingConversationListParamsWorkflowStatusOpen            MessagingConversationListParamsWorkflowStatus = "open"
	MessagingConversationListParamsWorkflowStatusWaitingInternal MessagingConversationListParamsWorkflowStatus = "waiting_internal"
	MessagingConversationListParamsWorkflowStatusWaitingExternal MessagingConversationListParamsWorkflowStatus = "waiting_external"
	MessagingConversationListParamsWorkflowStatusNeedsApproval   MessagingConversationListParamsWorkflowStatus = "needs_approval"
	MessagingConversationListParamsWorkflowStatusResolved        MessagingConversationListParamsWorkflowStatus = "resolved"
)
