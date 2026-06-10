# Auth

## APIKeys

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateAPIKeyRequestParam">CreateAPIKeyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountBranding">AccountBranding</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountPortal">AccountPortal</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatedAPIKey">CreatedAPIKey</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Geolocation">Geolocation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAPIKey">ListAPIKey</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Owner">Owner</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PageInfo">PageInfo</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Role">Role</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyDeleteResponse">AuthAPIKeyDeleteResponse</a>

Methods:

- <code title="post /v1/auth/api-keys">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyNewParams">AuthAPIKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatedAPIKey">CreatedAPIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auth/api-keys/{id}">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyGetParams">AuthAPIKeyGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auth/api-keys">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyListParams">AuthAPIKeyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAPIKey">ListAPIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/auth/api-keys/{id}">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyDeleteResponse">AuthAPIKeyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RotateAPIKeyRequestParam">RotateAPIKeyRequestParam</a>

Methods:

- <code title="post /v1/auth/api-keys/{id}/actions/rotate">client.Auth.APIKeys.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyActionService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuthAPIKeyActionRotateParams">AuthAPIKeyActionRotateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatedAPIKey">CreatedAPIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Core

## Sandboxes

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateSandboxRequestParam">CreateSandboxRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSandbox">ListSandbox</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Sandbox">Sandbox</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxDeleteResponse">CoreSandboxDeleteResponse</a>

Methods:

- <code title="post /v1/core/sandboxes">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxNewParams">CoreSandboxNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/sandboxes/{id}">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxGetParams">CoreSandboxGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/sandboxes">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxListParams">CoreSandboxListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSandbox">ListSandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/core/sandboxes/{id}">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreSandboxDeleteResponse">CoreSandboxDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RequestLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Actor">Actor</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListRequestLog">ListRequestLog</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RequestLog">RequestLog</a>

Methods:

- <code title="get /v1/core/request-logs/{id}">client.Core.RequestLogs.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreRequestLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreRequestLogGetParams">CoreRequestLogGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RequestLog">RequestLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/request-logs">client.Core.RequestLogs.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreRequestLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreRequestLogListParams">CoreRequestLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListRequestLog">ListRequestLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AuditEvents

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuditEvent">AuditEvent</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuditFieldChange">AuditFieldChange</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAuditEvent">ListAuditEvent</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAuditFieldChange">ListAuditFieldChange</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListObjectType">ListObjectType</a>

Methods:

- <code title="get /v1/core/audit-events/{id}">client.Core.AuditEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAuditEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAuditEventGetParams">CoreAuditEventGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AuditEvent">AuditEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/audit-events">client.Core.AuditEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAuditEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAuditEventListParams">CoreAuditEventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAuditEvent">ListAuditEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/audit-events/resource-types">client.Core.AuditEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAuditEventService.GetResourceTypes">GetResourceTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListObjectType">ListObjectType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Addresses

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AddressSuggestion">AddressSuggestion</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAddressSuggestion">ListAddressSuggestion</a>

Methods:

- <code title="get /v1/core/addresses/suggestions">client.Core.Addresses.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAddressService.GetSuggestions">GetSuggestions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAddressGetSuggestionsParams">CoreAddressGetSuggestionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAddressSuggestion">ListAddressSuggestion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ValidateAddressRequestParam">ValidateAddressRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AddressComponents">AddressComponents</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ValidatedAddress">ValidatedAddress</a>

Methods:

- <code title="put /v1/core/addresses/actions/validate">client.Core.Addresses.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAddressActionService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAddressActionValidateParams">CoreAddressActionValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ValidatedAddress">ValidatedAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EmailLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailLog">EmailLog</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEmailLog">ListEmailLog</a>

Methods:

- <code title="get /v1/core/email-logs/{id}">client.Core.EmailLogs.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreEmailLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreEmailLogGetParams">CoreEmailLogGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailLog">EmailLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/email-logs">client.Core.EmailLogs.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreEmailLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreEmailLogListParams">CoreEmailLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEmailLog">ListEmailLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Catalog

## Units

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateUnitRequestParam">CreateUnitRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateUnitRequestParam">UpdateUnitRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListUnit">ListUnit</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Unit">Unit</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitDeleteResponse">CatalogUnitDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/units">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitNewParams">CatalogUnitNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Unit">Unit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/units/{id}">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGetParams">CatalogUnitGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Unit">Unit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/units/{id}">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitUpdateParams">CatalogUnitUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Unit">Unit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/units">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitListParams">CatalogUnitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListUnit">ListUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/units/{id}">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitDeleteResponse">CatalogUnitDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## UnitGroups

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateUnitGroupRequestParam">CreateUnitGroupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateUnitGroupUnitParam">CreateUnitGroupUnitParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateUnitGroupRequestParam">UpdateUnitGroupRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListUnitGroup">ListUnitGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListUnitGroupUnit">ListUnitGroupUnit</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroup">UnitGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroupUnit">UnitGroupUnit</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupDeleteResponse">CatalogUnitGroupDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/unit-groups">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupNewParams">CatalogUnitGroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroup">UnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups/{id}">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupGetParams">CatalogUnitGroupGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroup">UnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/unit-groups/{id}">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUpdateParams">CatalogUnitGroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroup">UnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupListParams">CatalogUnitGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListUnitGroup">ListUnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/unit-groups/{id}">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupDeleteResponse">CatalogUnitGroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Units

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateUnitGroupUnitRequestParam">CreateUnitGroupUnitRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateUnitGroupUnitRequestParam">UpdateUnitGroupUnitRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitDeleteResponse">CatalogUnitGroupUnitDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/unit-groups/{unit_group_id}/units">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, unitGroupID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitNewParams">CatalogUnitGroupUnitNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroupUnit">UnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups/{unit_group_id}/units/{id}">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitGetParams">CatalogUnitGroupUnitGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroupUnit">UnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/unit-groups/{unit_group_id}/units/{id}">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitUpdateParams">CatalogUnitGroupUnitUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitGroupUnit">UnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups/{unit_group_id}/units">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, unitGroupID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitListParams">CatalogUnitGroupUnitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListUnitGroupUnit">ListUnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/unit-groups/{unit_group_id}/units/{id}">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitDeleteParams">CatalogUnitGroupUnitDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupUnitDeleteResponse">CatalogUnitGroupUnitDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Properties

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatePropertyRequestParam">CreatePropertyRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdatePropertyRequestParam">UpdatePropertyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Attribute">Attribute</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAttribute">ListAttribute</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProperty">ListProperty</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Property">Property</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyDeleteResponse">CatalogPropertyDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/properties">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyNewParams">CatalogPropertyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Property">Property</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties/{id}">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyGetParams">CatalogPropertyGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Property">Property</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/properties/{id}">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyUpdateParams">CatalogPropertyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Property">Property</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyListParams">CatalogPropertyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProperty">ListProperty</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/properties/{id}">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyDeleteResponse">CatalogPropertyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Attributes

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateAttributeRequestParam">CreateAttributeRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateAttributeRequestParam">UpdateAttributeRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeDeleteResponse">CatalogPropertyAttributeDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/properties/{property_id}/attributes">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeNewParams">CatalogPropertyAttributeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Attribute">Attribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties/{property_id}/attributes/{id}">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeGetParams">CatalogPropertyAttributeGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Attribute">Attribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/properties/{property_id}/attributes/{id}">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeUpdateParams">CatalogPropertyAttributeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Attribute">Attribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties/{property_id}/attributes">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeListParams">CatalogPropertyAttributeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAttribute">ListAttribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/properties/{property_id}/attributes/{id}">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeDeleteParams">CatalogPropertyAttributeDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyAttributeDeleteResponse">CatalogPropertyAttributeDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Items

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemCategory">ItemCategory</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemInventory">ItemInventory</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListItem">ListItem</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Quantity">Quantity</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Rate">Rate</a>

Methods:

- <code title="get /v1/catalog/items/{id}">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemGetParams">CatalogItemGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemListParams">CatalogItemListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListItem">ListItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/catalog/items/{id}/category/{category_id}">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.ChangeCategory">ChangeCategory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, categoryID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemChangeCategoryParams">CatalogItemChangeCategoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items/{id}/inventory">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.GetInventory">GetInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemGetInventoryParams">CatalogItemGetInventoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemInventory">ItemInventory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Attributes

Methods:

- <code title="put /v1/catalog/items/{id}/attributes/{attribute_id}">client.Catalog.Items.Attributes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemAttributeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attributeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemAttributeUpdateParams">CatalogItemAttributeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/items/{id}/attributes/{attribute_id}">client.Catalog.Items.Attributes.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemAttributeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attributeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemAttributeDeleteParams">CatalogItemAttributeDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ItemCategories

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateItemCategoryRequestParam">CreateItemCategoryRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateItemCategoryRequestParam">UpdateItemCategoryRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListItemCategory">ListItemCategory</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryDeleteResponse">CatalogItemCategoryDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryChangeUnitGroupResponse">CatalogItemCategoryChangeUnitGroupResponse</a>

Methods:

- <code title="post /v1/catalog/item-categories">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryNewParams">CatalogItemCategoryNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemCategory">ItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/item-categories/{id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryGetParams">CatalogItemCategoryGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemCategory">ItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/item-categories/{id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryUpdateParams">CatalogItemCategoryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemCategory">ItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/item-categories">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryListParams">CatalogItemCategoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListItemCategory">ListItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/item-categories/{id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryDeleteResponse">CatalogItemCategoryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/catalog/item-categories/{id}/unit-groups/{unit_group_id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryService.ChangeUnitGroup">ChangeUnitGroup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, unitGroupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryChangeUnitGroupParams">CatalogItemCategoryChangeUnitGroupParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryChangeUnitGroupResponse">CatalogItemCategoryChangeUnitGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Properties

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyUpdateResponse">CatalogItemCategoryPropertyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyDeleteResponse">CatalogItemCategoryPropertyDeleteResponse</a>

Methods:

- <code title="put /v1/catalog/item-categories/{id}/properties/{property_id}">client.Catalog.ItemCategories.Properties.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyUpdateParams">CatalogItemCategoryPropertyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyUpdateResponse">CatalogItemCategoryPropertyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/item-categories/{id}/properties/{property_id}">client.Catalog.ItemCategories.Properties.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyDeleteParams">CatalogItemCategoryPropertyDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryPropertyDeleteResponse">CatalogItemCategoryPropertyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Materials

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateMaterialRequestParam">CreateMaterialRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuantityInputRequestParam">QuantityInputRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RateInputParam">RateInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateMaterialRequestParam">UpdateMaterialRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMaterial">ListMaterial</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Material">Material</a>

Methods:

- <code title="post /v1/catalog/materials">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialNewParams">CatalogMaterialNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/materials/{id}">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialGetParams">CatalogMaterialGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/materials/{id}">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialUpdateParams">CatalogMaterialUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/materials">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialListParams">CatalogMaterialListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMaterial">ListMaterial</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/materials/{id}">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Parts

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatePartRequestParam">CreatePartRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdatePartRequestParam">UpdatePartRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPart">ListPart</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Part">Part</a>

Methods:

- <code title="post /v1/catalog/parts">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartNewParams">CatalogPartNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/parts/{id}">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartGetParams">CatalogPartGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/parts/{id}">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartUpdateParams">CatalogPartUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/parts">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartListParams">CatalogPartListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPart">ListPart</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/parts/{id}">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProductLines

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateProductLineRequestParam">CreateProductLineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateProductLineRequestParam">UpdateProductLineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductLine">ListProductLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductLine">ProductLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineDeleteResponse">CatalogProductLineDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/product-lines">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineNewParams">CatalogProductLineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductLine">ProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/product-lines/{id}">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineGetParams">CatalogProductLineGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductLine">ProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/product-lines/{id}">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineUpdateParams">CatalogProductLineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductLine">ProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/product-lines">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineListParams">CatalogProductLineListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductLine">ListProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/product-lines/{id}">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineDeleteResponse">CatalogProductLineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Products

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateProductRequestParam">CreateProductRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateProductRequestParam">UpdateProductRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProduct">ListProduct</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Product">Product</a>

Methods:

- <code title="post /v1/catalog/products">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductNewParams">CatalogProductNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/products/{id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductGetParams">CatalogProductGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/products/{id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductUpdateParams">CatalogProductUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/products">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductListParams">CatalogProductListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProduct">ListProduct</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/products/{id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductDeleteParams">CatalogProductDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/catalog/products/{id}/product-line/{product_line_id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductService.ChangeProductLine">ChangeProductLine</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, productLineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductChangeProductLineParams">CatalogProductChangeProductLineParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sales

## AccountGroups

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateAccountGroupRequestParam">CreateAccountGroupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateAccountGroupRequestParam">UpdateAccountGroupRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountGroup">AccountGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountGroup">ListAccountGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupDeleteResponse">SaleAccountGroupDeleteResponse</a>

Methods:

- <code title="post /v1/sales/account-groups">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupNewParams">SaleAccountGroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountGroup">AccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-groups/{id}">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountGroup">AccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/account-groups/{id}">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupUpdateParams">SaleAccountGroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountGroup">AccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-groups">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupListParams">SaleAccountGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountGroup">ListAccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/account-groups/{id}">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountGroupDeleteResponse">SaleAccountGroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Addresses

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AddressInputParam">AddressInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateAddressRequestParam">UpdateAddressRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAddress">ListAddress</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressDeleteResponse">SaleAddressDeleteResponse</a>

Methods:

- <code title="post /v1/sales/addresses">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressNewParams">SaleAddressNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/addresses/{id}">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/addresses/{id}">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressUpdateParams">SaleAddressUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/addresses">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressListParams">SaleAddressListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAddress">ListAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/addresses/{id}">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAddressDeleteResponse">SaleAddressDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountStatuses

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountStatus">AccountStatus</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountStatus">ListAccountStatus</a>

Methods:

- <code title="get /v1/sales/account-statuses/{id}">client.Sales.AccountStatuses.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountStatusGetParams">SaleAccountStatusGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountStatus">AccountStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-statuses">client.Sales.AccountStatuses.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountStatusService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountStatusListParams">SaleAccountStatusListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountStatus">ListAccountStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Priorities

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPriority">ListPriority</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Priority">Priority</a>

Methods:

- <code title="get /v1/sales/priorities/{id}">client.Sales.Priorities.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalePriorityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalePriorityGetParams">SalePriorityGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Priority">Priority</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/priorities">client.Sales.Priorities.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalePriorityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalePriorityListParams">SalePriorityListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPriority">ListPriority</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Customers

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateCustomerRequestParam">CreateCustomerRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuantityInputParam">QuantityInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateCustomerRequestParam">UpdateCustomerRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountUser">AccountUser</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Carrier">Carrier</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Consumption">Consumption</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerContactInfo">CustomerContactInfo</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerDefaults">CustomerDefaults</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerFreightPreferences">CustomerFreightPreferences</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerNotificationPreferences">CustomerNotificationPreferences</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Department">Department</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListConsumption">ListConsumption</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListCustomer">ListCustomer</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListLocation">ListLocation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachine">ListMachine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionStep">ListProductionStep</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScanningStation">ListScanningStation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListServiceLevel">ListServiceLevel</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Machine">Machine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PaymentTerm">PaymentTerm</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionOutput">ProductionOutput</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionStep">ProductionStep</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScanningStation">ScanningStation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ServiceLevel">ServiceLevel</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ShippingTerm">ShippingTerm</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerDeleteResponse">SaleCustomerDeleteResponse</a>

Methods:

- <code title="post /v1/sales/customers">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerNewParams">SaleCustomerNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerGetParams">SaleCustomerGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerUpdateParams">SaleCustomerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerListParams">SaleCustomerListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListCustomer">ListCustomer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerDeleteResponse">SaleCustomerDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MergeCustomersRequestParam">MergeCustomersRequestParam</a>

Methods:

- <code title="post /v1/sales/customers/{id}/actions/merge">client.Sales.Customers.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerActionService.Merge">Merge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerActionMergeParams">SaleCustomerActionMergeParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SalesOrders

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesOrderStatus">ListSalesOrderStatus</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderStatus">SalesOrderStatus</a>

Methods:

- <code title="get /v1/sales/sales-orders/statuses">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.GetStatuses">GetStatuses</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderGetStatusesParams">SaleSalesOrderGetStatusesParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesOrderStatus">ListSalesOrderStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Finance

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AdjustmentType">AdjustmentType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAdjustmentType">ListAdjustmentType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListTransactionMethod">ListTransactionMethod</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListTransactionType">ListTransactionType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#TransactionMethod">TransactionMethod</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#TransactionType">TransactionType</a>

Methods:

- <code title="get /v1/finance/adjustment-types">client.Finance.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinanceService.GetAdjustmentTypes">GetAdjustmentTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinanceGetAdjustmentTypesParams">FinanceGetAdjustmentTypesParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAdjustmentType">ListAdjustmentType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/transaction-methods">client.Finance.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinanceService.GetTransactionMethods">GetTransactionMethods</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinanceGetTransactionMethodsParams">FinanceGetTransactionMethodsParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListTransactionMethod">ListTransactionMethod</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/transaction-types">client.Finance.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinanceService.GetTransactionTypes">GetTransactionTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinanceGetTransactionTypesParams">FinanceGetTransactionTypesParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListTransactionType">ListTransactionType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PaymentTerms

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatePaymentTermRequestParam">CreatePaymentTermRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdatePaymentTermRequestParam">UpdatePaymentTermRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPaymentTerm">ListPaymentTerm</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermDeleteResponse">FinancePaymentTermDeleteResponse</a>

Methods:

- <code title="post /v1/finance/payment-terms">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermNewParams">FinancePaymentTermNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PaymentTerm">PaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/payment-terms/{id}">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermGetParams">FinancePaymentTermGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PaymentTerm">PaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/finance/payment-terms/{id}">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermUpdateParams">FinancePaymentTermUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PaymentTerm">PaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/payment-terms">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermListParams">FinancePaymentTermListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPaymentTerm">ListPaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/finance/payment-terms/{id}">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FinancePaymentTermDeleteResponse">FinancePaymentTermDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Operations

## ShippingTerms

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateShippingTermRequestParam">CreateShippingTermRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateShippingTermRequestParam">UpdateShippingTermRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListShippingTerm">ListShippingTerm</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermDeleteResponse">OperationShippingTermDeleteResponse</a>

Methods:

- <code title="post /v1/operations/shipping-terms">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermNewParams">OperationShippingTermNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ShippingTerm">ShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/shipping-terms/{id}">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermGetParams">OperationShippingTermGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ShippingTerm">ShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/shipping-terms/{id}">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermUpdateParams">OperationShippingTermUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ShippingTerm">ShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/shipping-terms">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermListParams">OperationShippingTermListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListShippingTerm">ListShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/shipping-terms/{id}">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShippingTermDeleteResponse">OperationShippingTermDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Carriers

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateCarrierRequestParam">CreateCarrierRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateCarrierRequestParam">UpdateCarrierRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListCarrier">ListCarrier</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierDeleteResponse">OperationCarrierDeleteResponse</a>

Methods:

- <code title="post /v1/operations/carriers">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierNewParams">OperationCarrierNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Carrier">Carrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers/{id}">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierGetParams">OperationCarrierGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Carrier">Carrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/carriers/{id}">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierUpdateParams">OperationCarrierUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Carrier">Carrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierListParams">OperationCarrierListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListCarrier">ListCarrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/carriers/{id}">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierDeleteResponse">OperationCarrierDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ServiceLevels

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateServiceLevelRequestParam">CreateServiceLevelRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateServiceLevelRequestParam">UpdateServiceLevelRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelDeleteResponse">OperationCarrierServiceLevelDeleteResponse</a>

Methods:

- <code title="post /v1/operations/carriers/{carrier_id}/service-levels">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, carrierID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelNewParams">OperationCarrierServiceLevelNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ServiceLevel">ServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers/{carrier_id}/service-levels/{id}">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelGetParams">OperationCarrierServiceLevelGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ServiceLevel">ServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/carriers/{carrier_id}/service-levels/{id}">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelUpdateParams">OperationCarrierServiceLevelUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ServiceLevel">ServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers/{carrier_id}/service-levels">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, carrierID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelListParams">OperationCarrierServiceLevelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListServiceLevel">ListServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/carriers/{carrier_id}/service-levels/{id}">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelDeleteParams">OperationCarrierServiceLevelDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationCarrierServiceLevelDeleteResponse">OperationCarrierServiceLevelDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Locations

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateLocationRequestParam">CreateLocationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateLocationRequestParam">UpdateLocationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationDeleteResponse">OperationLocationDeleteResponse</a>

Methods:

- <code title="post /v1/operations/locations">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationNewParams">OperationLocationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Location">Location</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/locations/{id}">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationGetParams">OperationLocationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Location">Location</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/locations/{id}">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationUpdateParams">OperationLocationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Location">Location</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/locations">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationListParams">OperationLocationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListLocation">ListLocation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/locations/{id}">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationDeleteResponse">OperationLocationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LocationTypes

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListLocationType">ListLocationType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#LocationType">LocationType</a>

Methods:

- <code title="get /v1/operations/location-types/{id}">client.Operations.LocationTypes.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#LocationType">LocationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/location-types">client.Operations.LocationTypes.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationTypeListParams">OperationLocationTypeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListLocationType">ListLocationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ScanningStations

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateScanningStationRequestParam">CreateScanningStationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateScanningStationRequestParam">UpdateScanningStationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationDeleteResponse">OperationScanningStationDeleteResponse</a>

Methods:

- <code title="post /v1/operations/scanning-stations">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationNewParams">OperationScanningStationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScanningStation">ScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/scanning-stations/{id}">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationGetParams">OperationScanningStationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScanningStation">ScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/scanning-stations/{id}">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationUpdateParams">OperationScanningStationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScanningStation">ScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/scanning-stations">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationListParams">OperationScanningStationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScanningStation">ListScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/scanning-stations/{id}">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationScanningStationDeleteResponse">OperationScanningStationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Identity

## AccountUsers

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateAccountUserRequestParam">CreateAccountUserRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationPreferenceItemParam">NotificationPreferenceItemParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateAccountUserRequestParam">UpdateAccountUserRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountUser">ListAccountUser</a>

Methods:

- <code title="post /v1/identity/account-users">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserNewParams">IdentityAccountUserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountUser">AccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/account-users/{id}">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserGetParams">IdentityAccountUserGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountUser">AccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/identity/account-users/{id}">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserUpdateParams">IdentityAccountUserUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountUser">AccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/account-users">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserListParams">IdentityAccountUserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountUser">ListAccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionActivateResponse">IdentityAccountUserActionActivateResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionDisableResponse">IdentityAccountUserActionDisableResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionRemoveResponse">IdentityAccountUserActionRemoveResponse</a>

Methods:

- <code title="put /v1/identity/account-users/{id}/actions/activate">client.Identity.AccountUsers.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionActivateResponse">IdentityAccountUserActionActivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/identity/account-users/{id}/actions/disable">client.Identity.AccountUsers.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionService.Disable">Disable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionDisableResponse">IdentityAccountUserActionDisableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/identity/account-users/{id}/actions/remove">client.Identity.AccountUsers.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUserActionRemoveResponse">IdentityAccountUserActionRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Roles

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateRoleRequestParam">CreateRoleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateRoleRequestParam">UpdateRoleRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListRole">ListRole</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleDeleteResponse">IdentityRoleDeleteResponse</a>

Methods:

- <code title="post /v1/identity/roles">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleNewParams">IdentityRoleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Role">Role</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/roles/{id}">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleGetParams">IdentityRoleGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Role">Role</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/identity/roles/{id}">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleUpdateParams">IdentityRoleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Role">Role</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/roles">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleListParams">IdentityRoleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListRole">ListRole</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/identity/roles/{id}">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityRoleDeleteResponse">IdentityRoleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
