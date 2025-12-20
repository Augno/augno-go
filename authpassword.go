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

// AuthPasswordService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthPasswordService] method instead.
type AuthPasswordService struct {
	Options []option.RequestOption
	Actions AuthPasswordActionService
}

// NewAuthPasswordService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthPasswordService(opts ...option.RequestOption) (r AuthPasswordService) {
	r = AuthPasswordService{}
	r.Options = opts
	r.Actions = NewAuthPasswordActionService(opts...)
	return
}

// This endpoint is utilized to update a user's password. Once completed, the user
// object is returned. Learn more about authentication and authorization in our
// [documentation](https://docs.augno.com/guides/authentication).
func (r *AuthPasswordService) UpdatePassword(ctx context.Context, body AuthPasswordUpdatePasswordParams, opts ...option.RequestOption) (res *AuthPasswordUpdatePasswordResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/passwords"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AuthPasswordUpdatePasswordResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthPasswordUpdatePasswordResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthPasswordUpdatePasswordResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthPasswordUpdatePasswordParams struct {
	// The new password to be set
	NewPassword param.Opt[string] `json:"new_password,omitzero"`
	// The user's current password
	OldPassword param.Opt[string] `json:"old_password,omitzero"`
	paramObj
}

func (r AuthPasswordUpdatePasswordParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthPasswordUpdatePasswordParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthPasswordUpdatePasswordParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
