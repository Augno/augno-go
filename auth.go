// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/augno-go/internal/apijson"
	"github.com/stainless-sdks/augno-go/internal/requestconfig"
	"github.com/stainless-sdks/augno-go/option"
	"github.com/stainless-sdks/augno-go/packages/param"
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
func (r *AuthService) RefreshToken(ctx context.Context, body AuthRefreshTokenParams, opts ...option.RequestOption) (res *AuthRefreshTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/auth/access-tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Represents a RefreshToken resource
type RefreshToken struct {
	// The refresh token
	Token string `json:"token,required"`
	// The refresh token expires at
	ExpiresAt string `json:"expires_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RefreshToken) RawJSON() string { return r.JSON.raw }
func (r *RefreshToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response schema for CreateAccessTokenResponse
type AuthRefreshTokenResponse struct {
	// The new access token
	AccessToken string `json:"access_token,required"`
	// A new refresh token
	RefreshToken RefreshToken `json:"refresh_token,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken  respjson.Field
		RefreshToken respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthRefreshTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthRefreshTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRefreshTokenParams struct {
	RefreshToken param.Opt[string] `json:"RefreshToken,omitzero"`
	paramObj
}

func (r AuthRefreshTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthRefreshTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthRefreshTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
