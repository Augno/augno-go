// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// IdentityService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityService] method instead.
type IdentityService struct {
	options []option.RequestOption
	// List and manage account users.
	AccountUsers IdentityAccountUserService
	// Manage account details, branding, portal, logo, and favicon.
	Accounts IdentityAccountService
	// List and manage roles.
	Roles IdentityRoleService
}

// NewIdentityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityService(opts ...option.RequestOption) (r IdentityService) {
	r = IdentityService{}
	r.options = opts
	r.AccountUsers = NewIdentityAccountUserService(opts...)
	r.Accounts = NewIdentityAccountService(opts...)
	r.Roles = NewIdentityRoleService(opts...)
	return
}
