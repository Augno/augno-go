// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/augno-go/internal/apijson"
	"github.com/stainless-sdks/augno-go/internal/requestconfig"
	"github.com/stainless-sdks/augno-go/option"
	"github.com/stainless-sdks/augno-go/packages/respjson"
)

// AuthService contains methods and other services that help with interacting with
// the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
	Actions AuthActionService
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	r.Actions = NewAuthActionService(opts...)
	return
}

// Refresh an access token using a refresh token.
func (r *AuthService) RefreshToken(ctx context.Context, opts ...option.RequestOption) (res *EmptyResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/auth/access-tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Revoke a refresh token.
func (r *AuthService) RevokeRefreshToken(ctx context.Context, opts ...option.RequestOption) (res *EmptyResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/auth/refresh-tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Request schema for EmptyResource
type EmptyResource struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmptyResource) RawJSON() string { return r.JSON.raw }
func (r *EmptyResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
