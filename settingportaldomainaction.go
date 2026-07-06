// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
)

// Connect a custom domain to the account's customer portal, verify its DNS, and
// resolve custom hosts to portal accounts.
//
// SettingPortalDomainActionService contains methods and other services that help
// with interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingPortalDomainActionService] method instead.
type SettingPortalDomainActionService struct {
	options []option.RequestOption
}

// NewSettingPortalDomainActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingPortalDomainActionService(opts ...option.RequestOption) (r SettingPortalDomainActionService) {
	r = SettingPortalDomainActionService{}
	r.options = opts
	return
}

// Re-checks the domain's DNS configuration and flips it to `verified` once the
// published records are confirmed.
//
// Returns the updated domain (still `pending` if DNS has not propagated yet) along
// with the currently required DNS records.
//
// This endpoint requires the permission: `self:update`.
func (r *SettingPortalDomainActionService) Verify(ctx context.Context, id string, opts ...option.RequestOption) (res *PortalDomain, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/settings/portal-domains/%s/actions/verify", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}
