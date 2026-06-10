// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"github.com/augno/augno-go/option"
)

// AuthService contains methods and other services that help with interacting with
// the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	options []option.RequestOption
	// Create and manage API keys for programmatic access.
	APIKeys AuthAPIKeyService
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.options = opts
	r.APIKeys = NewAuthAPIKeyService(opts...)
	return
}
