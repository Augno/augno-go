// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// OperationShipmentService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationShipmentService] method instead.
type OperationShipmentService struct {
	options []option.RequestOption
	// List and manage shipments, shipment lines, and shipping operations.
	Actions OperationShipmentActionService
}

// NewOperationShipmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationShipmentService(opts ...option.RequestOption) (r OperationShipmentService) {
	r = OperationShipmentService{}
	r.options = opts
	r.Actions = NewOperationShipmentActionService(opts...)
	return
}
