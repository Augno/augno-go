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

// AuthPasswordActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthPasswordActionService] method instead.
type AuthPasswordActionService struct {
	Options []option.RequestOption
}

// NewAuthPasswordActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthPasswordActionService(opts ...option.RequestOption) (r AuthPasswordActionService) {
	r = AuthPasswordActionService{}
	r.Options = opts
	return
}

// This endpoint is utilized to request a password reset for a user. An email will
// be sent to the user with a link to reset their password. Learn more about
// authentication and authorization in our
// [documentation](https://docs.augno.com/guides/authentication).
func (r *AuthPasswordActionService) RequestPasswordReset(ctx context.Context, body AuthPasswordActionRequestPasswordResetParams, opts ...option.RequestOption) (res *AuthPasswordActionRequestPasswordResetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/passwords/actions/request-reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is utilized to reset a user's password using a password reset
// token. Once completed, new access and refresh tokens are set in cookies. Learn
// more about authentication and authorization in our
// [documentation](https://docs.augno.com/guides/authentication).
func (r *AuthPasswordActionService) ResetPassword(ctx context.Context, body AuthPasswordActionResetPasswordParams, opts ...option.RequestOption) (res *AuthPasswordActionResetPasswordResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/passwords/actions/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthPasswordActionRequestPasswordResetResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthPasswordActionRequestPasswordResetResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthPasswordActionRequestPasswordResetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthPasswordActionResetPasswordResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthPasswordActionResetPasswordResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthPasswordActionResetPasswordResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthPasswordActionRequestPasswordResetParams struct {
	// The account slug (optional)
	AccountSlug param.Opt[string] `json:"account_slug,omitzero"`
	// The username or email of the account to reset
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	paramObj
}

func (r AuthPasswordActionRequestPasswordResetParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthPasswordActionRequestPasswordResetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthPasswordActionRequestPasswordResetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthPasswordActionResetPasswordParams struct {
	// The password reset token
	Token param.Opt[string] `json:"token,omitzero"`
	// The new password of the user
	Password param.Opt[string] `json:"password,omitzero"`
	paramObj
}

func (r AuthPasswordActionResetPasswordParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthPasswordActionResetPasswordParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthPasswordActionResetPasswordParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
