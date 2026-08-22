// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
)

// List, retrieve, trigger, cancel, and continue agent runs.
//
// AIRunActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIRunActionService] method instead.
type AIRunActionService struct {
	options []option.RequestOption
}

// NewAIRunActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIRunActionService(opts ...option.RequestOption) (r AIRunActionService) {
	r = AIRunActionService{}
	r.options = opts
	return
}

// Cancels an in-progress agent run.
//
// A run can be cancelled while it is working or paused waiting on the user —
// `pending`, `running`, `awaiting_input`, or `awaiting_approval`. Cancelling a run
// in a terminal status (`completed`, `failed`, `cancelled`) returns a validation
// error.
//
// Cancelling a run that is `awaiting_approval` counts as denying the review: every
// action still pending review is recorded as rejected, attributed to the caller.
// Work the agent already completed is not undone.
//
// This endpoint requires the permission: `agent_runs:update`.
func (r *AIRunActionService) Cancel(ctx context.Context, id string, body AIRunActionCancelParams, opts ...option.RequestOption) (res *AgentRun, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/runs/%s/actions/cancel", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Resumes a paused agent run with a user message and any tool review decisions.
//
// The run must be `awaiting_input` or `awaiting_approval`; resuming it from any
// other status returns a validation error. It moves back to `running` and
// continues asynchronously, so poll Retrieve Agent Run to follow it. Each approval
// and denial is recorded on the matching action and attributed to the caller.
//
// This endpoint requires the permission: `agent_runs:update`.
func (r *AIRunActionService) Continue(ctx context.Context, id string, params AIRunActionContinueParams, opts ...option.RequestOption) (res *AgentRun, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/runs/%s/actions/continue", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retries a failed agent run by resuming its existing transcript.
//
// Only runs in the `failed` status can be retried; retrying a run in any other
// status returns a validation error. The run is re-attempted from where it left
// off — its prior reasoning and tool results are replayed, so the agent continues
// with full knowledge of what it already did rather than starting over, which
// minimizes the chance of it repeating side effects it has already caused.
//
// A run can be retried at most five times in total, and any automatic retries the
// platform already performed for transient failures count against that budget.
//
// This endpoint requires the permission: `agent_runs:update`.
func (r *AIRunActionService) Retry(ctx context.Context, id string, body AIRunActionRetryParams, opts ...option.RequestOption) (res *AgentRun, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/runs/%s/actions/retry", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to resume a paused agent run.
//
// The property Message is required.
type ContinueRunRequestParam struct {
	// Message to send to the agent as the next turn of the run.
	//
	// It accompanies any approval or denial in the same request, so use it to tell the
	// agent how to proceed with what you just allowed or blocked.
	Message string `json:"message" api:"required"`
	// Tool-call IDs (the `tool_use_id` of individual blocked calls) to approve.
	//
	// Use this instead of `approved_tool_slugs` to approve ONE specific call when
	// several pending calls share the same tool slug — approving by slug would approve
	// all of them. Approvals are one-time.
	ApprovedToolCallIDs []string `json:"approved_tool_call_ids,omitzero"`
	// Slugs of tools whose pending calls should be approved.
	//
	// Approves every call currently pending review for each named tool. Approval is
	// one-time — the next call to the same tool pauses for review again. Tools you do
	// not name are left pending, and the run resumes without them.
	ApprovedToolSlugs []string `json:"approved_tool_slugs,omitzero"`
	// Tool-call IDs (the `tool_use_id` of individual blocked calls) to deny.
	//
	// Per-call counterpart of `rejected_tool_slugs`, letting you deny one specific
	// call among several that share a slug. Each denied call is answered with a
	// "denied by user" result and the run continues.
	RejectedToolCallIDs []string `json:"rejected_tool_call_ids,omitzero"`
	// Slugs of tools whose pending calls should be denied.
	//
	// The run keeps going: each denied call is answered with a "denied by user" result
	// so the agent proceeds without it, instead of cancelling the run. A single resume
	// may both approve and reject different tools.
	RejectedToolSlugs []string `json:"rejected_tool_slugs,omitzero"`
	paramObj
}

func (r ContinueRunRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ContinueRunRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContinueRunRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIRunActionCancelParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "actions", "definition", "definition.config", "definition.tools",
	// "definition.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIRunActionCancelParams]'s query parameters as
// `url.Values`.
func (r AIRunActionCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIRunActionContinueParams struct {
	// Request to resume a paused agent run.
	ContinueRunRequest ContinueRunRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "actions", "definition", "definition.config", "definition.tools",
	// "definition.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r AIRunActionContinueParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContinueRunRequest)
}
func (r *AIRunActionContinueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AIRunActionContinueParams]'s query parameters as
// `url.Values`.
func (r AIRunActionContinueParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIRunActionRetryParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "actions", "definition", "definition.config", "definition.tools",
	// "definition.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIRunActionRetryParams]'s query parameters as `url.Values`.
func (r AIRunActionRetryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
