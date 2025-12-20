// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Augno/go-sdk/internal/apijson"
	"github.com/Augno/go-sdk/internal/requestconfig"
	"github.com/Augno/go-sdk/option"
	"github.com/Augno/go-sdk/packages/param"
	"github.com/Augno/go-sdk/packages/respjson"
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

// This endpoint is utilized to login a user. Once completed, the user object is
// returned. An access and refresh token are set in cookies. Learn more about
// authentication and authorization in our
// [documentation](https://docs.augno.com/guides/authentication).
func (r *AuthActionService) LoginUser(ctx context.Context, body AuthActionLoginUserParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/actions/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A user in the Augno system
type User struct {
	// The ID of the user
	ID string `json:"id"`
	// The created at timestamp of the user
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The email of the user
	Email string `json:"email"`
	// The email verified status of the user
	EmailVerified time.Time `json:"email_verified" format:"date-time"`
	// The image URL of the user
	ImageURL string `json:"image_url"`
	// The name of the user
	Name string `json:"name"`
	// The object type, always "user"
	Object string `json:"object"`
	// The updated at timestamp of the user
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The username of the user
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Email         respjson.Field
		EmailVerified respjson.Field
		ImageURL      respjson.Field
		Name          respjson.Field
		Object        respjson.Field
		UpdatedAt     respjson.Field
		Username      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthActionLoginUserParams struct {
	// The username or email for authentication
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	// The password for authentication
	Password param.Opt[string] `json:"password,omitzero"`
	paramObj
}

func (r AuthActionLoginUserParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthActionLoginUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthActionLoginUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
