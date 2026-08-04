// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/augno/augno-go/internal/apijson"
	"github.com/augno/augno-go/internal/apiquery"
	"github.com/augno/augno-go/internal/requestconfig"
	"github.com/augno/augno-go/option"
	"github.com/augno/augno-go/packages/param"
	"github.com/augno/augno-go/packages/respjson"
)

// Autocomplete, look up details, and validate addresses.
//
// CoreAddressService contains methods and other services that help with
// interacting with the augno API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreAddressService] method instead.
type CoreAddressService struct {
	options []option.RequestOption
	// Autocomplete, look up details, and validate addresses.
	Actions CoreAddressActionService
}

// NewCoreAddressService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreAddressService(opts ...option.RequestOption) (r CoreAddressService) {
	r = CoreAddressService{}
	r.options = opts
	r.Actions = NewCoreAddressActionService(opts...)
	return
}

// Returns address suggestions for partial address text, for use in type-ahead
// address entry.
//
// Only street addresses are suggested; cities, regions, and business listings are
// not returned. Suggestions are lookup results, not saved addresses in your
// account. Pass a suggestion's `id` to the address details endpoint to get the
// full parsed address.
func (r *CoreAddressService) GetSuggestions(ctx context.Context, query CoreAddressGetSuggestionsParams, opts ...option.RequestOption) (res *ListAddressSuggestion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/addresses/suggestions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A candidate address returned by address autocomplete.
//
// A suggestion is a lookup result from the address provider, not a saved address
// in your account. Creating an address from one is a separate step.
type AddressSuggestion struct {
	// Identifier of the suggested place.
	//
	// Pass this value as the `id` path parameter of the address details endpoint to
	// retrieve the full parsed address. It is issued by the underlying address
	// provider rather than by Augno, so it is not a durable Augno resource ID.
	ID string `json:"id" api:"required"`
	// Full description of the address.
	Description string `json:"description" api:"required"`
	// Main text (typically the street address).
	MainText string `json:"main_text" api:"required"`
	// Resource type identifier.
	//
	// Any of "address_suggestion".
	Object AddressSuggestionObject `json:"object" api:"required"`
	// Secondary text (typically city, state, country).
	SecondaryText string `json:"secondary_text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Description   respjson.Field
		MainText      respjson.Field
		Object        respjson.Field
		SecondaryText respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressSuggestion) RawJSON() string { return r.JSON.raw }
func (r *AddressSuggestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AddressSuggestionObject string

const (
	AddressSuggestionObjectAddressSuggestion AddressSuggestionObject = "address_suggestion"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAddressSuggestion struct {
	// Resources in this page.
	Data []AddressSuggestion `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAddressSuggestionObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListAddressSuggestion) RawJSON() string { return r.JSON.raw }
func (r *ListAddressSuggestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAddressSuggestionObject string

const (
	ListAddressSuggestionObjectList ListAddressSuggestionObject = "list"
)

type CoreAddressGetSuggestionsParams struct {
	// Partial address text to generate suggestions for.
	Input string `query:"input" api:"required" json:"-"`
	// Opaque token that groups a series of related autocomplete requests into a single
	// session.
	//
	// Reuse the same token for each keystroke of one address entry, and again when you
	// retrieve the details of the suggestion the user picks, so the whole entry is
	// treated as one lookup.
	SessionToken param.Opt[string] `query:"session_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreAddressGetSuggestionsParams]'s query parameters as
// `url.Values`.
func (r CoreAddressGetSuggestionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
