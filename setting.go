// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// SettingService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	options []option.RequestOption
	// Connect a custom domain to the account's customer portal, verify its DNS, and
	// resolve custom hosts to portal accounts.
	PortalDomains SettingPortalDomainService
	// List and manage third-party account integrations.
	Integrations SettingIntegrationService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r SettingService) {
	r = SettingService{}
	r.options = opts
	r.PortalDomains = NewSettingPortalDomainService(opts...)
	r.Integrations = NewSettingIntegrationService(opts...)
	return
}
