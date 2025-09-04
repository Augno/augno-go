// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/augno-go/internal/apijson"
	"github.com/stainless-sdks/augno-go/internal/requestconfig"
	"github.com/stainless-sdks/augno-go/option"
	"github.com/stainless-sdks/augno-go/packages/param"
	"github.com/stainless-sdks/augno-go/packages/respjson"
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
func (r *AuthActionService) LoginUser(ctx context.Context, body AuthActionLoginUserParams, opts ...option.RequestOption) (res *AuthActionLoginUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/auth/actions/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response schema for LoginResponse
type AuthActionLoginUserResponse struct {
	// The access token for the user
	AccessToken string `json:"access_token,required"`
	// The account affiliations
	AccountAffiliations []AuthActionLoginUserResponseAccountAffiliation `json:"account_affiliations,required"`
	// The current account in use
	CurrentAccount AuthActionLoginUserResponseCurrentAccount `json:"current_account,required"`
	// The refresh token for the user
	RefreshToken RefreshToken `json:"refresh_token,required"`
	// The user that was logged in
	User AuthActionLoginUserResponseUser `json:"user,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken         respjson.Field
		AccountAffiliations respjson.Field
		CurrentAccount      respjson.Field
		RefreshToken        respjson.Field
		User                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthActionLoginUserResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthActionLoginUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a AccountAffiliation resource
type AuthActionLoginUserResponseAccountAffiliation struct {
	// The ID of the account affiliation
	ID string `json:"id,required"`
	// The name of the account affiliation
	Name string `json:"name,required"`
	// Represents a AccountAffiliationRole resource
	Role AuthActionLoginUserResponseAccountAffiliationRole `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthActionLoginUserResponseAccountAffiliation) RawJSON() string { return r.JSON.raw }
func (r *AuthActionLoginUserResponseAccountAffiliation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a AccountAffiliationRole resource
type AuthActionLoginUserResponseAccountAffiliationRole struct {
	// The ID of the role
	ID string `json:"id,required"`
	// The name of the role
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthActionLoginUserResponseAccountAffiliationRole) RawJSON() string { return r.JSON.raw }
func (r *AuthActionLoginUserResponseAccountAffiliationRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current account in use
type AuthActionLoginUserResponseCurrentAccount struct {
	// The ID of the current account
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthActionLoginUserResponseCurrentAccount) RawJSON() string { return r.JSON.raw }
func (r *AuthActionLoginUserResponseCurrentAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user that was logged in
type AuthActionLoginUserResponseUser struct {
	// The ID of the user
	ID string `json:"id,required"`
	// The created at timestamp of the user
	CreatedAt string `json:"created_at,required"`
	// The updated at timestamp of the user
	UpdatedAt string `json:"updated_at,required"`
	// The email of the user
	Email string `json:"email"`
	// The email verified status of the user
	EmailVerified string `json:"email_verified"`
	// The image URL of the user
	ImageURL string `json:"image_url,nullable"`
	// The name of the user
	Name string `json:"name"`
	// The username of the user
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		UpdatedAt     respjson.Field
		Email         respjson.Field
		EmailVerified respjson.Field
		ImageURL      respjson.Field
		Name          respjson.Field
		Username      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthActionLoginUserResponseUser) RawJSON() string { return r.JSON.raw }
func (r *AuthActionLoginUserResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthActionLoginUserParams struct {
	Password string `json:"password,required"`
	Username string `json:"username,required"`
	paramObj
}

func (r AuthActionLoginUserParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthActionLoginUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthActionLoginUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
