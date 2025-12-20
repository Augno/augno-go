// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"

	"github.com/Augno/go-sdk/internal/apijson"
	"github.com/Augno/go-sdk/internal/requestconfig"
	"github.com/Augno/go-sdk/option"
	"github.com/Augno/go-sdk/packages/param"
	"github.com/Augno/go-sdk/packages/respjson"
)

// AuthService contains methods and other services that help with interacting with
// the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options   []option.RequestOption
	Actions   AuthActionService
	Passwords AuthPasswordService
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	r.Actions = NewAuthActionService(opts...)
	r.Passwords = NewAuthPasswordService(opts...)
	return
}

// This endpoint is utilized to refresh an access token using a refresh token. Once
// completed, a new access token is set in a cookie. Learn more about
// authentication and authorization in our
// [documentation](https://docs.augno.com/guides/authentication).
func (r *AuthService) RefreshToken(ctx context.Context, opts ...option.RequestOption) (res *AuthRefreshTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/access-tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint is utilized to register a new user. Once completed, the user
// object is returned. An access and refresh token are set in cookies. Learn more
// about authentication and authorization in our
// [documentation](https://docs.augno.com/guides/authentication).
func (r *AuthService) RegisterUser(ctx context.Context, body AuthRegisterUserParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is utilized to revoke a refresh token. Once completed, the refresh
// token is revoked and is no longer valid for refreshing an access token. Learn
// more about authentication and authorization in our
// [documentation](https://docs.augno.com/guides/authentication).
func (r *AuthService) RevokeRefreshToken(ctx context.Context, opts ...option.RequestOption) (res *AuthRevokeRefreshTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/refresh-tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AuthRefreshTokenResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthRefreshTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthRefreshTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRevokeRefreshTokenResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthRevokeRefreshTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthRevokeRefreshTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRegisterUserParams struct {
	// The email address for the new user
	Email param.Opt[string] `json:"email,omitzero"`
	// The full name of the new user
	Name param.Opt[string] `json:"name,omitzero"`
	// The password for the new user
	Password param.Opt[string] `json:"password,omitzero"`
	paramObj
}

func (r AuthRegisterUserParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthRegisterUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthRegisterUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
