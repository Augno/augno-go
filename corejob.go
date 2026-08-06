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
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/respjson"
)

// View the jobs that track asynchronous work. Endpoints that answer 202 Accepted
// raise one and point at it with a Location header.
//
// CoreJobService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreJobService] method instead.
type CoreJobService struct {
	options []option.RequestOption
}

// NewCoreJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCoreJobService(opts ...option.RequestOption) (r CoreJobService) {
	r = CoreJobService{}
	r.options = opts
	return
}

// Returns a job by ID — poll the job named in a `202 Accepted` response's
// `Location` to observe its outcome. A completed export carries the link to its
// file on `export.url`.
func (r *CoreJobService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/jobs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Cancels a job and returns it carrying its `cancelled` status. Work in flight is
// not interrupted but can no longer settle, and a finished job cannot be
// cancelled.
func (r *CoreJobService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/jobs/%s/cancel", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Records a piece of work the API accepted and carries out asynchronously.
// Endpoints answering `202 Accepted` point at one with a `Location` header; poll
// it for the outcome.
type Job struct {
	// Job ID.
	ID string `json:"id" api:"required"`
	// When the job was cancelled.
	CancelledAt time.Time `json:"cancelled_at" api:"required" format:"date-time"`
	// When the job finished processing, whether or not every row succeeded.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// When the job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email of the account user who requested the work.
	CreatedByEmail string `json:"created_by_email" api:"required"`
	// The ID of the account user who requested the work.
	CreatedByID string `json:"created_by_id" api:"required"`
	// Name of the account user who requested the work.
	CreatedByName string `json:"created_by_name" api:"required"`
	// Username of the account user who requested the work.
	CreatedByUsername string `json:"created_by_username" api:"required"`
	// A one-line reason the last attempt failed.
	ErrorSummary string `json:"error_summary" api:"required"`
	// One entry per failure, so a `completed` job can still carry failed rows. A
	// whole-job failure records a single entry with no `index`.
	Errors []RowError `json:"errors" api:"required"`
	// Points a completed export job at the file it produced.
	Export JobExport `json:"export" api:"required"`
	// When the most recent attempt failed. A retry that succeeds leaves this alongside
	// `completed_at`.
	FailedAt time.Time `json:"failed_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "job".
	Object JobObject `json:"object" api:"required"`
	// One entry per request row that produced a resource. A bulk create records these
	// when it accepts the request, so they stay provisional until `status` is
	// `completed`.
	Results []JobResult `json:"results" api:"required"`
	// When the job began executing.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// How far the job has got. `completed` means the work was processed, not that
	// every row succeeded — read `errors`.
	//
	// Any of "created", "started", "completed", "failed", "cancelled".
	Status JobStatus `json:"status" api:"required"`
	// The kind of work the job carries out.
	//
	// Any of "bulkcreate", "bulkupsert", "export".
	Type JobType `json:"type" api:"required"`
	// When the job was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CancelledAt       respjson.Field
		CompletedAt       respjson.Field
		CreatedAt         respjson.Field
		CreatedByEmail    respjson.Field
		CreatedByID       respjson.Field
		CreatedByName     respjson.Field
		CreatedByUsername respjson.Field
		ErrorSummary      respjson.Field
		Errors            respjson.Field
		Export            respjson.Field
		FailedAt          respjson.Field
		Object            respjson.Field
		Results           respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type JobObject string

const (
	JobObjectJob JobObject = "job"
)

// How far the job has got. `completed` means the work was processed, not that
// every row succeeded — read `errors`.
type JobStatus string

const (
	JobStatusCreated   JobStatus = "created"
	JobStatusStarted   JobStatus = "started"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
)

// The kind of work the job carries out.
type JobType string

const (
	JobTypeBulkcreate JobType = "bulkcreate"
	JobTypeBulkupsert JobType = "bulkupsert"
	JobTypeExport     JobType = "export"
)

// Points a completed export job at the file it produced.
type JobExport struct {
	// Presigned link to the file, valid for five minutes — read the job again for a
	// fresh one. It carries its own authorization, so treat it as a credential: do not
	// log it, store it, or pass it on.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExport) RawJSON() string { return r.JSON.raw }
func (r *JobExport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Accounts for one request row that produced a resource. With `errors`, also
// row-indexed, every submitted row lands in exactly one of the two once the job
// completes.
type JobResult struct {
	// ID of the resource this row produced.
	ID string `json:"id" api:"required"`
	// Whether the resource was created or updated.
	//
	// Any of "created", "updated".
	Action JobResultAction `json:"action" api:"required"`
	// Zero-based row of the request this result names.
	Index int64 `json:"index" api:"required"`
	// Resources created alongside this row's own resource — for a bulk production run
	// create, the ids of its batches. Omitted for operations that produce none.
	SubResourceIDs []string `json:"sub_resource_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Action         respjson.Field
		Index          respjson.Field
		SubResourceIDs respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResult) RawJSON() string { return r.JSON.raw }
func (r *JobResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the resource was created or updated.
type JobResultAction string

const (
	JobResultActionCreated JobResultAction = "created"
	JobResultActionUpdated JobResultAction = "updated"
)

// QuotaInfo provides machine-readable details about a plan-imposed resource limit.
// Included in limit_exceeded errors so clients can display upgrade prompts, usage
// bars, or implement programmatic retry/backoff logic.
type QuotaInfo struct {
	// Limit is the maximum number of resources allowed by the current plan.
	Limit int64 `json:"limit" api:"required"`
	// ResetAt is the time when the quota resets, if applicable. Nil for static
	// (non-metered) limits.
	ResetAt time.Time `json:"reset_at" api:"required" format:"date-time"`
	// Used is the number of resources currently consumed.
	Used int64 `json:"used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		ResetAt     respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaInfo) RawJSON() string { return r.JSON.raw }
func (r *QuotaInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseError is the JSON-serializable error body returned to API clients. It
// contains only public information. This struct is used by the OpenAPI schema
// generator to produce documentation.
type ResponseError struct {
	// A machine-readable code for the error.
	//
	// Any of "expired_token", "api_key_expired", "api_key_revoked",
	// "invalid_credentials", "insufficient_permissions", "payment_required",
	// "agent_spending_cap_reached", "validation_failed", "missing_field",
	// "invalid_format", "method_not_allowed", "resource_not_found", "resource_exists",
	// "resource_conflict", "resource_gone", "idempotency_in_progress",
	// "limit_exceeded", "registration_closed", "rate_limit_exceeded",
	// "parameter_missing", "parameter_invalid", "parameter_unknown",
	// "parameters_exclusive", "internal_error", "service_unavailable",
	// "external_service_error", "timeout", "connection_error", "request_timeout",
	// "client_closed_request", "api_version_required", "api_version_invalid",
	// "api_version_too_old".
	Code ResponseErrorCode `json:"code" api:"required"`
	// A URL to documentation about the error.
	DocURL string `json:"doc_url" api:"required"`
	// Whether this error is transient and the request can be retried.
	IsTransient bool `json:"is_transient" api:"required"`
	// A human-readable message providing more details about the error.
	Message string `json:"message" api:"required"`
	// The parameter that caused the error, if applicable.
	Param string `json:"param" api:"required"`
	// QuotaInfo provides machine-readable details about a plan-imposed resource limit.
	// Included in limit_exceeded errors so clients can display upgrade prompts, usage
	// bars, or implement programmatic retry/backoff logic.
	Quota QuotaInfo `json:"quota" api:"required"`
	// RequestLogURL is a link to the dashboard page for this request's log entry. Nil
	// when no request log is available.
	RequestLogURL string `json:"request_log_url" api:"required"`
	// The type of error.
	//
	// Any of "api_error", "idempotency_error", "invalid_request_error".
	Type ResponseErrorType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code          respjson.Field
		DocURL        respjson.Field
		IsTransient   respjson.Field
		Message       respjson.Field
		Param         respjson.Field
		Quota         respjson.Field
		RequestLogURL respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseError) RawJSON() string { return r.JSON.raw }
func (r *ResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A machine-readable code for the error.
type ResponseErrorCode string

const (
	ResponseErrorCodeExpiredToken            ResponseErrorCode = "expired_token"
	ResponseErrorCodeAPIKeyExpired           ResponseErrorCode = "api_key_expired"
	ResponseErrorCodeAPIKeyRevoked           ResponseErrorCode = "api_key_revoked"
	ResponseErrorCodeInvalidCredentials      ResponseErrorCode = "invalid_credentials"
	ResponseErrorCodeInsufficientPermissions ResponseErrorCode = "insufficient_permissions"
	ResponseErrorCodePaymentRequired         ResponseErrorCode = "payment_required"
	ResponseErrorCodeAgentSpendingCapReached ResponseErrorCode = "agent_spending_cap_reached"
	ResponseErrorCodeValidationFailed        ResponseErrorCode = "validation_failed"
	ResponseErrorCodeMissingField            ResponseErrorCode = "missing_field"
	ResponseErrorCodeInvalidFormat           ResponseErrorCode = "invalid_format"
	ResponseErrorCodeMethodNotAllowed        ResponseErrorCode = "method_not_allowed"
	ResponseErrorCodeResourceNotFound        ResponseErrorCode = "resource_not_found"
	ResponseErrorCodeResourceExists          ResponseErrorCode = "resource_exists"
	ResponseErrorCodeResourceConflict        ResponseErrorCode = "resource_conflict"
	ResponseErrorCodeResourceGone            ResponseErrorCode = "resource_gone"
	ResponseErrorCodeIdempotencyInProgress   ResponseErrorCode = "idempotency_in_progress"
	ResponseErrorCodeLimitExceeded           ResponseErrorCode = "limit_exceeded"
	ResponseErrorCodeRegistrationClosed      ResponseErrorCode = "registration_closed"
	ResponseErrorCodeRateLimitExceeded       ResponseErrorCode = "rate_limit_exceeded"
	ResponseErrorCodeParameterMissing        ResponseErrorCode = "parameter_missing"
	ResponseErrorCodeParameterInvalid        ResponseErrorCode = "parameter_invalid"
	ResponseErrorCodeParameterUnknown        ResponseErrorCode = "parameter_unknown"
	ResponseErrorCodeParametersExclusive     ResponseErrorCode = "parameters_exclusive"
	ResponseErrorCodeInternalError           ResponseErrorCode = "internal_error"
	ResponseErrorCodeServiceUnavailable      ResponseErrorCode = "service_unavailable"
	ResponseErrorCodeExternalServiceError    ResponseErrorCode = "external_service_error"
	ResponseErrorCodeTimeout                 ResponseErrorCode = "timeout"
	ResponseErrorCodeConnectionError         ResponseErrorCode = "connection_error"
	ResponseErrorCodeRequestTimeout          ResponseErrorCode = "request_timeout"
	ResponseErrorCodeClientClosedRequest     ResponseErrorCode = "client_closed_request"
	ResponseErrorCodeAPIVersionRequired      ResponseErrorCode = "api_version_required"
	ResponseErrorCodeAPIVersionInvalid       ResponseErrorCode = "api_version_invalid"
	ResponseErrorCodeAPIVersionTooOld        ResponseErrorCode = "api_version_too_old"
)

// The type of error.
type ResponseErrorType string

const (
	ResponseErrorTypeAPIError            ResponseErrorType = "api_error"
	ResponseErrorTypeIdempotencyError    ResponseErrorType = "idempotency_error"
	ResponseErrorTypeInvalidRequestError ResponseErrorType = "invalid_request_error"
)

// pairs one row of a bulk request with the failure it produced; a failure of the
// request as a whole carries no index
type RowError struct {
	// ResponseError is the JSON-serializable error body returned to API clients. It
	// contains only public information. This struct is used by the OpenAPI schema
	// generator to produce documentation.
	Error ResponseError `json:"error" api:"required"`
	// Zero-based row of the request this failure names. Absent for a failure of the
	// whole request.
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RowError) RawJSON() string { return r.JSON.raw }
func (r *RowError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
