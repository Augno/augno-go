// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	shimjson "github.com/augno/augno-go/internal/encoding/json"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
)

// Create and manage API keys for programmatic access.
//
// AuthAPIKeyActionService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthAPIKeyActionService] method instead.
type AuthAPIKeyActionService struct {
	options []option.RequestOption
}

// NewAuthAPIKeyActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthAPIKeyActionService(opts ...option.RequestOption) (r AuthAPIKeyActionService) {
	r = AuthAPIKeyActionService{}
	r.options = opts
	return
}

// Rotates an [API key](https://docs.augno.com/api/api-keys) by revoking the
// existing key and issuing a replacement with the same name, role, and expiration
// (unless overridden).
//
// The secret key is returned once and cannot be retrieved later, so you should
// store it securely. We provide some
// [recommendations](https://docs.augno.com/api/managing-api-keys) on how you can
// manage your API keys.
func (r *AuthAPIKeyActionService) Rotate(ctx context.Context, id string, params AuthAPIKeyActionRotateParams, opts ...option.RequestOption) (res *CreatedAPIKey, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/auth/api-keys/%s/actions/rotate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to rotate an API key.
type RotateAPIKeyRequestParam struct {
	// Expiration timestamp override for the new key.
	//
	// If omitted, the previous key's expiration is used.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// When to revoke the old key.
	//
	// If omitted (or in the past), the old key is revoked immediately. A future
	// timestamp schedules revocation (keeping the old key valid until then) up to a
	// maximum of 30 days out.
	RevokeAt param.Opt[time.Time] `json:"revoke_at,omitzero" format:"date-time"`
	paramObj
}

func (r RotateAPIKeyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RotateAPIKeyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RotateAPIKeyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthAPIKeyActionRotateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to rotate an API key.
	RotateAPIKeyRequest RotateAPIKeyRequestParam
	paramObj
}

func (r AuthAPIKeyActionRotateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RotateAPIKeyRequest)
}
func (r *AuthAPIKeyActionRotateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AuthAPIKeyActionRotateParams]'s query parameters as
// `url.Values`.
func (r AuthAPIKeyActionRotateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
