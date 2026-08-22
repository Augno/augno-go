// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
)

// Connect a custom domain to the account's customer portal, verify its DNS, and
// resolve custom hosts to portal accounts.
//
// SettingPortalDomainActionService contains methods and other services that help
// with interacting with the openmrp API.
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

// Re-checks a portal domain against the serving provider and advances its status.
//
// Run this after publishing the DNS records, and keep polling it: the domain stays
// `pending` while its records are missing or misconfigured, moves to `securing`
// once they are correct and its TLS certificate is being issued, and reaches
// `verified` only once that certificate is live and the portal answers on the
// domain. The response carries the updated domain along with the records still
// required. Verifying an already-verified domain returns it unchanged.
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
