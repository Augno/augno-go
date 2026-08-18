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

// List, create, update, and delete agent memories.
//
// AIMemoryService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMemoryService] method instead.
type AIMemoryService struct {
	options []option.RequestOption
}

// NewAIMemoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIMemoryService(opts ...option.RequestOption) (r AIMemoryService) {
	r = AIMemoryService{}
	r.options = opts
	return
}

// Saves a piece of information for agents to recall on future runs.
//
// This endpoint requires the permission: `agent_memories:create`.
func (r *AIMemoryService) New(ctx context.Context, body AIMemoryNewParams, opts ...option.RequestOption) (res *AgentMemory, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/memories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns an agent memory by ID.
//
// An expired memory is still returned here, even though it is excluded from list
// results and no longer recalled by agents.
//
// This endpoint requires the permission: `agent_memories:read`.
func (r *AIMemoryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentMemory, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/memories/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an agent memory.
//
// Only the fields included in the request are changed; everything else keeps its
// current value.
//
// This endpoint requires the permission: `agent_memories:update`.
func (r *AIMemoryService) Update(ctx context.Context, id string, body AIMemoryUpdateParams, opts ...option.RequestOption) (res *AgentMemory, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/memories/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of agent memories for the current account, newest
// first.
//
// Memories whose `expires_at` has passed are excluded. The `q` search term matches
// against a memory's ID, category, content, and the ID of the record it is scoped
// to.
//
// This endpoint requires the permission: `agent_memories:read`.
func (r *AIMemoryService) List(ctx context.Context, query AIMemoryListParams, opts ...option.RequestOption) (res *ListAgentMemory, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/memories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes an agent memory so it is no longer recalled.
//
// Deleting a memory that has already been deleted succeeds rather than returning
// an error.
//
// This endpoint requires the permission: `agent_memories:delete`.
func (r *AIMemoryService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AIMemoryDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/memories/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A piece of information an agent has saved for recall in future runs.
type AgentMemory struct {
	// Memory ID.
	ID string `json:"id" api:"required"`
	// The kind of information this memory holds, used to group related memories.
	//
	//   - `preference`: how someone likes things done, such as a customer who always
	//     wants express shipping.
	//   - `fact`: a durable detail worth remembering about the account or one of its
	//     records, such as a customer's typical order size.
	//   - `instruction`: standing guidance for agents to follow, such as always
	//     confirming freight before issuing an order.
	//
	// Any of "preference", "fact", "instruction".
	Category AgentMemoryCategory `json:"category" api:"required"`
	// The information itself, written as plain text for an agent to read.
	Content string `json:"content" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Entity Entity `json:"entity" api:"required"`
	// When this memory stops being used.
	//
	// Past this time the memory is no longer recalled by agents and is omitted from
	// list results, but it is not deleted and can still be retrieved by ID. A memory
	// with no expiration is used indefinitely.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Relative importance from `0` to `1`, used to prioritize which memories the agent
	// recalls.
	//
	// An agent takes in only a limited number of memories per run, and the
	// highest-importance ones are recalled first.
	Importance float64 `json:"importance" api:"required"`
	// Arbitrary metadata as JSON. Encoded as a JSON value (object, array, string,
	// number, boolean, or null), not a JSON-encoded string.
	Metadata any `json:"metadata" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_memory".
	Object AgentMemoryObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Category    respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		Entity      respjson.Field
		ExpiresAt   respjson.Field
		Importance  respjson.Field
		Metadata    respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentMemory) RawJSON() string { return r.JSON.raw }
func (r *AgentMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of information this memory holds, used to group related memories.
//
//   - `preference`: how someone likes things done, such as a customer who always
//     wants express shipping.
//   - `fact`: a durable detail worth remembering about the account or one of its
//     records, such as a customer's typical order size.
//   - `instruction`: standing guidance for agents to follow, such as always
//     confirming freight before issuing an order.
type AgentMemoryCategory string

const (
	AgentMemoryCategoryPreference  AgentMemoryCategory = "preference"
	AgentMemoryCategoryFact        AgentMemoryCategory = "fact"
	AgentMemoryCategoryInstruction AgentMemoryCategory = "instruction"
)

// Resource type identifier.
type AgentMemoryObject string

const (
	AgentMemoryObjectAgentMemory AgentMemoryObject = "agent_memory"
)

// Request to create an agent memory.
//
// The properties Category, Content are required.
type CreateMemoryRequestParam struct {
	// The kind of information this memory holds, used to group related memories.
	//
	//   - `preference`: how someone likes things done, such as a customer who always
	//     wants express shipping.
	//   - `fact`: a durable detail worth remembering about the account or one of its
	//     records, such as a customer's typical order size.
	//   - `instruction`: standing guidance for agents to follow, such as always
	//     confirming freight before issuing an order.
	//
	// Any of "preference", "fact", "instruction".
	Category CreateMemoryRequestCategory `json:"category,omitzero" api:"required"`
	// The information to remember, written as plain text for an agent to read.
	Content string `json:"content" api:"required"`
	// ID of the platform record this memory is scoped to.
	//
	// Provide together with `entity_type`.
	EntityID param.Opt[string] `json:"entity_id,omitzero"`
	// Type of platform record this memory is scoped to (e.g. `customer`, `product`).
	//
	// Provide together with `entity_id` to scope the memory to a specific record; omit
	// both for a memory that is not tied to any particular record.
	EntityType param.Opt[string] `json:"entity_type,omitzero"`
	// When this memory should stop being used, as an ISO 8601 timestamp (e.g.
	// `2026-01-02T15:04:05Z`).
	//
	// Past this time the memory is no longer recalled by agents and is omitted from
	// list results, but it is not deleted. Omit it for a memory that should be used
	// indefinitely.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	// Relative importance from `0` to `1` in increments of `0.1`, used to prioritize
	// which memories the agent recalls.
	//
	// An agent takes in only a limited number of memories per run and recalls the
	// highest-importance ones first, so a memory created without an importance is
	// stored at `0` and is the first to be left out.
	Importance param.Opt[float64] `json:"importance,omitzero"`
	// Arbitrary metadata as JSON. Encoded as a JSON value (object, array, string,
	// number, boolean, or null), not a JSON-encoded string.
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r CreateMemoryRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMemoryRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMemoryRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of information this memory holds, used to group related memories.
//
//   - `preference`: how someone likes things done, such as a customer who always
//     wants express shipping.
//   - `fact`: a durable detail worth remembering about the account or one of its
//     records, such as a customer's typical order size.
//   - `instruction`: standing guidance for agents to follow, such as always
//     confirming freight before issuing an order.
type CreateMemoryRequestCategory string

const (
	CreateMemoryRequestCategoryPreference  CreateMemoryRequestCategory = "preference"
	CreateMemoryRequestCategoryFact        CreateMemoryRequestCategory = "fact"
	CreateMemoryRequestCategoryInstruction CreateMemoryRequestCategory = "instruction"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAgentMemory struct {
	// Resources in this page.
	Data []AgentMemory `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentMemoryObject `json:"object" api:"required"`
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
func (r ListAgentMemory) RawJSON() string { return r.JSON.raw }
func (r *ListAgentMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentMemoryObject string

const (
	ListAgentMemoryObjectList ListAgentMemoryObject = "list"
)

// Request to update an agent memory.
type UpdateMemoryRequestParam struct {
	// ID of the platform record this memory is scoped to.
	//
	// Provide together with `entity_type`; send `null` to unscope the memory.
	EntityID param.Opt[string] `json:"entity_id,omitzero"`
	// Type of platform record this memory is scoped to (e.g. `customer`, `product`).
	//
	// Provide together with `entity_id` to scope the memory to a specific record; send
	// `null` (on either entity field) to unscope the memory.
	EntityType param.Opt[string] `json:"entity_type,omitzero"`
	// When this memory should stop being used, as an ISO 8601 timestamp (e.g.
	// `2026-01-02T15:04:05Z`).
	//
	// Past this time the memory is no longer recalled by agents and is omitted from
	// list results, but it is not deleted. Send `null` so the memory is used
	// indefinitely.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	// The information to remember, written as plain text for an agent to read.
	Content param.Opt[string] `json:"content,omitzero"`
	// Relative importance from `0` to `1` in increments of `0.1`, used to prioritize
	// which memories the agent recalls.
	//
	// An agent takes in only a limited number of memories per run and recalls the
	// highest-importance ones first.
	Importance param.Opt[float64] `json:"importance,omitzero"`
	// Arbitrary metadata as JSON.
	//
	// Replaces the stored metadata outright rather than merging into it. Encoded as a
	// JSON value (object, array, string, number, boolean, or null), not a JSON-encoded
	// string.
	Metadata any `json:"metadata,omitzero"`
	// The kind of information this memory holds, used to group related memories.
	//
	//   - `preference`: how someone likes things done, such as a customer who always
	//     wants express shipping.
	//   - `fact`: a durable detail worth remembering about the account or one of its
	//     records, such as a customer's typical order size.
	//   - `instruction`: standing guidance for agents to follow, such as always
	//     confirming freight before issuing an order.
	//
	// Any of "preference", "fact", "instruction".
	Category UpdateMemoryRequestCategory `json:"category,omitzero"`
	paramObj
}

func (r UpdateMemoryRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMemoryRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMemoryRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of information this memory holds, used to group related memories.
//
//   - `preference`: how someone likes things done, such as a customer who always
//     wants express shipping.
//   - `fact`: a durable detail worth remembering about the account or one of its
//     records, such as a customer's typical order size.
//   - `instruction`: standing guidance for agents to follow, such as always
//     confirming freight before issuing an order.
type UpdateMemoryRequestCategory string

const (
	UpdateMemoryRequestCategoryPreference  UpdateMemoryRequestCategory = "preference"
	UpdateMemoryRequestCategoryFact        UpdateMemoryRequestCategory = "fact"
	UpdateMemoryRequestCategoryInstruction UpdateMemoryRequestCategory = "instruction"
)

type AIMemoryDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNewParams struct {
	// Request to create an agent memory.
	CreateMemoryRequest CreateMemoryRequestParam
	paramObj
}

func (r AIMemoryNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMemoryRequest)
}
func (r *AIMemoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryUpdateParams struct {
	// Request to update an agent memory.
	UpdateMemoryRequest UpdateMemoryRequestParam
	paramObj
}

func (r AIMemoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateMemoryRequest)
}
func (r *AIMemoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter to memories scoped to this entity type (e.g. `customer`, `product`).
	EntityType param.Opt[string] `query:"entity_type,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter to memories with this exact category (e.g. `preference`, `fact`).
	//
	// Any of "preference", "fact", "instruction".
	Category AIMemoryListParamsCategory `query:"category,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMemoryListParams]'s query parameters as `url.Values`.
func (r AIMemoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter to memories with this exact category (e.g. `preference`, `fact`).
type AIMemoryListParamsCategory string

const (
	AIMemoryListParamsCategoryPreference  AIMemoryListParamsCategory = "preference"
	AIMemoryListParamsCategoryFact        AIMemoryListParamsCategory = "fact"
	AIMemoryListParamsCategoryInstruction AIMemoryListParamsCategory = "instruction"
)
