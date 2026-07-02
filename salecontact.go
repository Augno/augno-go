// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// SaleContactService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleContactService] method instead.
type SaleContactService struct {
	options []option.RequestOption
	// Look up the people you do business with.
	Actions SaleContactActionService
}

// NewSaleContactService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleContactService(opts ...option.RequestOption) (r SaleContactService) {
	r = SaleContactService{}
	r.options = opts
	r.Actions = NewSaleContactActionService(opts...)
	return
}
