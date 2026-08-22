// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"github.com/open-mrp/openmrp-go/option"
)

// SaleAccountUserService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleAccountUserService] method instead.
type SaleAccountUserService struct {
	options []option.RequestOption
	// List and manage sales targets for account users.
	SalesTargets SaleAccountUserSalesTargetService
}

// NewSaleAccountUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleAccountUserService(opts ...option.RequestOption) (r SaleAccountUserService) {
	r = SaleAccountUserService{}
	r.options = opts
	r.SalesTargets = NewSaleAccountUserSalesTargetService(opts...)
	return
}
