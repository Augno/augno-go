// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// OperationService contains methods and other services that help with interacting
// with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationService] method instead.
type OperationService struct {
	options []option.RequestOption
	// List and manage shipping terms.
	ShippingTerms OperationShippingTermService
	// List and manage carriers and their Shippo integrations.
	Carriers OperationCarrierService
	// List and manage locations.
	Locations OperationLocationService
	// List and manage locations.
	LocationTypes OperationLocationTypeService
	// List and manage scanning stations.
	ScanningStations OperationScanningStationService
}

// NewOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationService(opts ...option.RequestOption) (r OperationService) {
	r = OperationService{}
	r.options = opts
	r.ShippingTerms = NewOperationShippingTermService(opts...)
	r.Carriers = NewOperationCarrierService(opts...)
	r.Locations = NewOperationLocationService(opts...)
	r.LocationTypes = NewOperationLocationTypeService(opts...)
	r.ScanningStations = NewOperationScanningStationService(opts...)
	return
}
