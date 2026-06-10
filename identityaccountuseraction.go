// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/respjson"
)

// List and manage account users.
//
// IdentityAccountUserActionService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityAccountUserActionService] method instead.
type IdentityAccountUserActionService struct {
	options []option.RequestOption
}

// NewIdentityAccountUserActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIdentityAccountUserActionService(opts ...option.RequestOption) (r IdentityAccountUserActionService) {
	r = IdentityAccountUserActionService{}
	r.options = opts
	return
}

// Activates a disabled or removed account user.
func (r *IdentityAccountUserActionService) Activate(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityAccountUserActionActivateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s/actions/activate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Disables an account user. Disabled users will not be able to access the target
// account.
func (r *IdentityAccountUserActionService) Disable(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityAccountUserActionDisableResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s/actions/disable", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Removes a user from the target account.
func (r *IdentityAccountUserActionService) Remove(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityAccountUserActionRemoveResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s/actions/remove", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type IdentityAccountUserActionActivateResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityAccountUserActionActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityAccountUserActionActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityAccountUserActionDisableResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityAccountUserActionDisableResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityAccountUserActionDisableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityAccountUserActionRemoveResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityAccountUserActionRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityAccountUserActionRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
