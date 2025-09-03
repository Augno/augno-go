// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/augno-go/internal/apijson"
	"github.com/stainless-sdks/augno-go/internal/requestconfig"
	"github.com/stainless-sdks/augno-go/option"
	"github.com/stainless-sdks/augno-go/packages/param"
)

// AuthActionService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthActionService] method instead.
type AuthActionService struct {
	Options []option.RequestOption
}

// NewAuthActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthActionService(opts ...option.RequestOption) (r AuthActionService) {
	r = AuthActionService{}
	r.Options = opts
	return
}

// Login a user and get an access and refresh token.
func (r *AuthActionService) LoginUser(ctx context.Context, body AuthActionLoginUserParams, opts ...option.RequestOption) (res *CreateAccessTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/auth/actions/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthActionLoginUserParams struct {
	Password param.Opt[string] `json:"password,omitzero"`
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r AuthActionLoginUserParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthActionLoginUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthActionLoginUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
