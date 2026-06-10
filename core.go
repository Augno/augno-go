// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// CoreService contains methods and other services that help with interacting with
// the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreService] method instead.
type CoreService struct {
	options []option.RequestOption
	// List and manage sandbox environments.
	Sandboxes CoreSandboxService
	// List and retrieve request logs.
	RequestLogs CoreRequestLogService
	// List and retrieve audit events.
	AuditEvents CoreAuditEventService
	// Autocomplete, look up details, and validate addresses.
	Addresses CoreAddressService
	// View email logs for accounts.
	EmailLogs CoreEmailLogService
}

// NewCoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCoreService(opts ...option.RequestOption) (r CoreService) {
	r = CoreService{}
	r.options = opts
	r.Sandboxes = NewCoreSandboxService(opts...)
	r.RequestLogs = NewCoreRequestLogService(opts...)
	r.AuditEvents = NewCoreAuditEventService(opts...)
	r.Addresses = NewCoreAddressService(opts...)
	r.EmailLogs = NewCoreEmailLogService(opts...)
	return
}
