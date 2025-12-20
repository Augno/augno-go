// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"

	"github.com/Augno/go-sdk/internal/apijson"
	"github.com/Augno/go-sdk/internal/requestconfig"
	"github.com/Augno/go-sdk/option"
	"github.com/Augno/go-sdk/packages/respjson"
)

// HealthzService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthzService] method instead.
type HealthzService struct {
	Options []option.RequestOption
}

// NewHealthzService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthzService(opts ...option.RequestOption) (r HealthzService) {
	r = HealthzService{}
	r.Options = opts
	return
}

// Returns the current health status, environment, and version.
func (r *HealthzService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthzCheckResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "healthz"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Represents a Healthcheck resource
type HealthzCheckResponse struct {
	// Deployment environment (development, production)
	Environment string `json:"environment,required"`
	// Current operational status of the API service
	Status string `json:"status,required"`
	// Application version number
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Environment respjson.Field
		Status      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthzCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthzCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
