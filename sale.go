// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// SaleService contains methods and other services that help with interacting with
// the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleService] method instead.
type SaleService struct {
	options []option.RequestOption
	// List and manage account groups.
	AccountGroups SaleAccountGroupService
	// List and manage addresses for accounts.
	Addresses SaleAddressService
	// List and retrieve account statuses.
	AccountStatuses SaleAccountStatusService
	// List and retrieve priorities.
	Priorities SalePriorityService
	// Manage customer accounts.
	Customers   SaleCustomerService
	Contacts    SaleContactService
	SalesOrders SaleSalesOrderService
}

// NewSaleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSaleService(opts ...option.RequestOption) (r SaleService) {
	r = SaleService{}
	r.options = opts
	r.AccountGroups = NewSaleAccountGroupService(opts...)
	r.Addresses = NewSaleAddressService(opts...)
	r.AccountStatuses = NewSaleAccountStatusService(opts...)
	r.Priorities = NewSalePriorityService(opts...)
	r.Customers = NewSaleCustomerService(opts...)
	r.Contacts = NewSaleContactService(opts...)
	r.SalesOrders = NewSaleSalesOrderService(opts...)
	return
}
