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

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Entity">Entity</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEntity">ListEntity</a>

Methods:

- <code title="get /v1/core/search">client.Core.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreService.GetSearch">GetSearch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreGetSearchParams">CoreGetSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEntity">ListEntity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#JobExport">JobExport</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#JobResult">JobResult</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuotaInfo">QuotaInfo</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ResponseError">ResponseError</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RowError">RowError</a>

Methods:

- <code title="get /v1/core/jobs/{id}">client.Core.Jobs.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/core/jobs/{id}/cancel">client.Core.Jobs.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Analytics

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeDeliveryPerformanceRequestParam">AnalyzeDeliveryPerformanceRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeOeeRequestParam">AnalyzeOeeRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeOeeTrendRequestParam">AnalyzeOeeTrendRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeScheduleAttainmentRequestParam">AnalyzeScheduleAttainmentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OeeDepartmentPlannedTimeParam">OeeDepartmentPlannedTimeParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeDeliveryPerformanceResponse">AnalyzeDeliveryPerformanceResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeOeeResponse">AnalyzeOeeResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeOeeTrendResponse">AnalyzeOeeTrendResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeScheduleAttainmentResponse">AnalyzeScheduleAttainmentResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AttainmentBucket">AttainmentBucket</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DeliveryBacklogBucket">DeliveryBacklogBucket</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DeliveryPerformance">DeliveryPerformance</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FrozenAdherence">FrozenAdherence</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAttainmentBucket">ListAttainmentBucket</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDeliveryBacklogBucket">ListDeliveryBacklogBucket</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDeliveryPerformance">ListDeliveryPerformance</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListFrozenAdherence">ListFrozenAdherence</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListOeeDepartment">ListOeeDepartment</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListOeeDowntimeReason">ListOeeDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListOeeTrendPeriod">ListOeeTrendPeriod</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OeeDepartment">OeeDepartment</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OeeDowntimeReason">OeeDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OeeTrendPeriod">OeeTrendPeriod</a>

Methods:

- <code title="put /v1/core/analytics/delivery-performance">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsService.UpdateDeliveryPerformance">UpdateDeliveryPerformance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsUpdateDeliveryPerformanceParams">CoreAnalyticsUpdateDeliveryPerformanceParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeDeliveryPerformanceResponse">AnalyzeDeliveryPerformanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/core/analytics/oee">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsService.UpdateOee">UpdateOee</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsUpdateOeeParams">CoreAnalyticsUpdateOeeParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeOeeResponse">AnalyzeOeeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/core/analytics/oee-trend">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsService.UpdateOeeTrend">UpdateOeeTrend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsUpdateOeeTrendParams">CoreAnalyticsUpdateOeeTrendParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeOeeTrendResponse">AnalyzeOeeTrendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/core/analytics/schedule-attainment">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsService.UpdateScheduleAttainment">UpdateScheduleAttainment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CoreAnalyticsUpdateScheduleAttainmentParams">CoreAnalyticsUpdateScheduleAttainmentParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AnalyzeScheduleAttainmentResponse">AnalyzeScheduleAttainmentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertUnitsRequestParam">BulkUpsertUnitsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertUnitInputParam">UpsertUnitInputParam</a>

Methods:

- <code title="post /v1/catalog/units/actions/bulk-upsert">client.Catalog.Units.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitActionBulkUpsertParams">CatalogUnitActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertUnitGroupsRequestParam">BulkUpsertUnitGroupsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UnitIdentifierParam">UnitIdentifierParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertUnitGroupConversionInputParam">UpsertUnitGroupConversionInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertUnitGroupInputParam">UpsertUnitGroupInputParam</a>

Methods:

- <code title="post /v1/catalog/unit-groups/actions/bulk-upsert">client.Catalog.UnitGroups.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogUnitGroupActionBulkUpsertParams">CatalogUnitGroupActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertPropertiesRequestParam">BulkUpsertPropertiesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertPropertyAttributeInputParam">UpsertPropertyAttributeInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertPropertyInputParam">UpsertPropertyInputParam</a>

Methods:

- <code title="post /v1/catalog/properties/actions/bulk-upsert">client.Catalog.Properties.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPropertyActionBulkUpsertParams">CatalogPropertyActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Items

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemCategory">ItemCategory</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemInventory">ItemInventory</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemLotDefault">ItemLotDefault</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListItem">ListItem</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Quantity">Quantity</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Rate">Rate</a>

Methods:

- <code title="get /v1/catalog/items/{id}">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemGetParams">CatalogItemGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemListParams">CatalogItemListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListItem">ListItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/catalog/items/{id}/category/{category_id}">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.ChangeCategory">ChangeCategory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, categoryID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemChangeCategoryParams">CatalogItemChangeCategoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items/{id}/inventory">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.GetInventory">GetInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemGetInventoryParams">CatalogItemGetInventoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemInventory">ItemInventory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items/{id}/lot-default">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemService.GetLotDefault">GetLotDefault</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemGetLotDefaultParams">CatalogItemGetLotDefaultParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ItemLotDefault">ItemLotDefault</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertItemCategoriesRequestParam">BulkUpsertItemCategoriesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ObjectIdentifierParam">ObjectIdentifierParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertItemCategoryInputParam">UpsertItemCategoryInputParam</a>

Methods:

- <code title="post /v1/catalog/item-categories/actions/bulk-upsert">client.Catalog.ItemCategories.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogItemCategoryActionBulkUpsertParams">CatalogItemCategoryActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertMaterialsRequestParam">BulkUpsertMaterialsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertMaterialInputParam">UpsertMaterialInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertMaterialPropertyParam">UpsertMaterialPropertyParam</a>

Methods:

- <code title="post /v1/catalog/materials/actions/bulk-upsert">client.Catalog.Materials.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogMaterialActionBulkUpsertParams">CatalogMaterialActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertPartsRequestParam">BulkUpsertPartsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertPartInputParam">UpsertPartInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertPartPropertyParam">UpsertPartPropertyParam</a>

Methods:

- <code title="post /v1/catalog/parts/actions/bulk-upsert">client.Catalog.Parts.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogPartActionBulkUpsertParams">CatalogPartActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProductLines

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateProductLineRequestParam">CreateProductLineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuantityInputParam">QuantityInputParam</a>
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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertProductLinesRequestParam">BulkUpsertProductLinesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertProductLineInputParam">UpsertProductLineInputParam</a>

Methods:

- <code title="post /v1/catalog/product-lines/actions/bulk-upsert">client.Catalog.ProductLines.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductLineActionBulkUpsertParams">CatalogProductLineActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertProductsRequestParam">BulkUpsertProductsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertProductInputParam">UpsertProductInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertProductPropertyParam">UpsertProductPropertyParam</a>

Methods:

- <code title="post /v1/catalog/products/actions/bulk-upsert">client.Catalog.Products.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CatalogProductActionBulkUpsertParams">CatalogProductActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AvailableTool">AvailableTool</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAvailableTool">ListAvailableTool</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListToolGroup">ListToolGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ToolGroup">ToolGroup</a>

Methods:

- <code title="get /v1/ai/tool-groups">client.AI.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIService.GetToolGroups">GetToolGroups</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIGetToolGroupsParams">AIGetToolGroupsParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListToolGroup">ListToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/tools">client.AI.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIService.GetTools">GetTools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIGetToolsParams">AIGetToolsParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAvailableTool">ListAvailableTool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ConfigInputParam">ConfigInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateAgentRequestParam">CreateAgentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ToolInputParam">ToolInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#TriggerConfigInputParam">TriggerConfigInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateAgentRequestParam">UpdateAgentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateAgentStatusRequestParam">UpdateAgentStatusRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentDefinition">AgentDefinition</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentDefinitionConfig">AgentDefinitionConfig</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentDefinitionTool">AgentDefinitionTool</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentDefinition">ListAgentDefinition</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentDefinitionTool">ListAgentDefinitionTool</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#TriggerConfig">TriggerConfig</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentDeleteResponse">AIAgentDeleteResponse</a>

Methods:

- <code title="post /v1/ai/agents">client.AI.Agents.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentNewParams">AIAgentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/agents/{id}">client.AI.Agents.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentGetParams">AIAgentGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/ai/agents/{id}">client.AI.Agents.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentUpdateParams">AIAgentUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/agents">client.AI.Agents.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentListParams">AIAgentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentDefinition">ListAgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/ai/agents/{id}">client.AI.Agents.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentDeleteResponse">AIAgentDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/ai/agents/{id}/status">client.AI.Agents.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentService.UpdateStatus">UpdateStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIAgentUpdateStatusParams">AIAgentUpdateStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#TriggerRunRequestParam">TriggerRunRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentAction">AgentAction</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentRun">AgentRun</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentRunStep">AgentRunStep</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentAction">ListAgentAction</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentRun">ListAgentRun</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentRunStep">ListAgentRunStep</a>

Methods:

- <code title="post /v1/ai/runs">client.AI.Runs.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunNewParams">AIRunNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/runs/{id}">client.AI.Runs.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunGetParams">AIRunGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/runs">client.AI.Runs.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunListParams">AIRunListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentRun">ListAgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ContinueRunRequestParam">ContinueRunRequestParam</a>

Methods:

- <code title="post /v1/ai/runs/{id}/actions/cancel">client.AI.Runs.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunActionService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunActionCancelParams">AIRunActionCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ai/runs/{id}/actions/continue">client.AI.Runs.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunActionService.Continue">Continue</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunActionContinueParams">AIRunActionContinueParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ai/runs/{id}/actions/retry">client.AI.Runs.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunActionService.Retry">Retry</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIRunActionRetryParams">AIRunActionRetryParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Memories

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateMemoryRequestParam">CreateMemoryRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateMemoryRequestParam">UpdateMemoryRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentMemory">AgentMemory</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentMemory">ListAgentMemory</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryDeleteResponse">AIMemoryDeleteResponse</a>

Methods:

- <code title="post /v1/ai/memories">client.AI.Memories.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryNewParams">AIMemoryNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentMemory">AgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/memories/{id}">client.AI.Memories.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentMemory">AgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/ai/memories/{id}">client.AI.Memories.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryUpdateParams">AIMemoryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AgentMemory">AgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/memories">client.AI.Memories.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryListParams">AIMemoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAgentMemory">ListAgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/ai/memories/{id}">client.AI.Memories.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AIMemoryDeleteResponse">AIMemoryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messaging

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListActor">ListActor</a>

Methods:

- <code title="get /v1/messaging/contacts">client.Messaging.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingService.GetContacts">GetContacts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGetContactsParams">MessagingGetContactsParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListActor">ListActor</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Notifications

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationTargetInputParam">NotificationTargetInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SendNotificationRequestParam">SendNotificationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListNotification">ListNotification</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListNotificationUnreadSummaryAccount">ListNotificationUnreadSummaryAccount</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Notification">Notification</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationSendResult">NotificationSendResult</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationUnreadCount">NotificationUnreadCount</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationUnreadSummary">NotificationUnreadSummary</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationUnreadSummaryAccount">NotificationUnreadSummaryAccount</a>

Methods:

- <code title="post /v1/messaging/notifications">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationNewParams">MessagingNotificationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationSendResult">NotificationSendResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications/{id}">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationGetParams">MessagingNotificationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationListParams">MessagingNotificationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListNotification">ListNotification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications/unread-count">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationService.GetUnreadCount">GetUnreadCount</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationUnreadCount">NotificationUnreadCount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications/unread-summary">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationService.GetUnreadSummary">GetUnreadSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationUnreadSummary">NotificationUnreadSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionMarkAllSeenResponse">MessagingNotificationActionMarkAllSeenResponse</a>

Methods:

- <code title="post /v1/messaging/notifications/{id}/actions/dismiss">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionService.Dismiss">Dismiss</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionDismissParams">MessagingNotificationActionDismissParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/notifications/actions/mark-all-seen">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionService.MarkAllSeen">MarkAllSeen</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionMarkAllSeenResponse">MessagingNotificationActionMarkAllSeenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/notifications/{id}/actions/read">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionReadParams">MessagingNotificationActionReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/notifications/{id}/actions/seen">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionService.Seen">Seen</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingNotificationActionSeenParams">MessagingNotificationActionSeenParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Announcements

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Announcement">Announcement</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAnnouncement">ListAnnouncement</a>

Methods:

- <code title="get /v1/messaging/announcements/{id}">client.Messaging.Announcements.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementGetParams">MessagingAnnouncementGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/announcements">client.Messaging.Announcements.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementListParams">MessagingAnnouncementListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAnnouncement">ListAnnouncement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="post /v1/messaging/announcements/{id}/actions/dismiss">client.Messaging.Announcements.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementActionService.Dismiss">Dismiss</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementActionDismissParams">MessagingAnnouncementActionDismissParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/announcements/{id}/actions/read">client.Messaging.Announcements.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementActionService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementActionReadParams">MessagingAnnouncementActionReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/announcements/{id}/actions/seen">client.Messaging.Announcements.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementActionService.Seen">Seen</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingAnnouncementActionSeenParams">MessagingAnnouncementActionSeenParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Conversations

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateConversationRequestParam">CreateConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateConversationRequestParam">UpdateConversationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ConversationParticipant">ConversationParticipant</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListConversation">ListConversation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListConversationParticipant">ListConversationParticipant</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessageAttachment">ListMessageAttachment</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessagingGroupMember">ListMessagingGroupMember</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessageAttachment">MessageAttachment</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroup">MessagingGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupMember">MessagingGroupMember</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReadCursor">ReadCursor</a>

Methods:

- <code title="post /v1/messaging/conversations">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationNewParams">MessagingConversationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations/{id}">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationGetParams">MessagingConversationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/messaging/conversations/{id}">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationUpdateParams">MessagingConversationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationListParams">MessagingConversationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListConversation">ListConversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AssignConversationRequestParam">AssignConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MarkConversationReadRequestParam">MarkConversationReadRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MuteConversationRequestParam">MuteConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReportConversationRequestParam">ReportConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SetLegalHoldRequestParam">SetLegalHoldRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SetWorkflowStatusRequestParam">SetWorkflowStatusRequestParam</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/actions/archive">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionArchiveParams">MessagingConversationActionArchiveParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/assign">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Assign">Assign</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionAssignParams">MessagingConversationActionAssignParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/hide">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Hide">Hide</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionHideParams">MessagingConversationActionHideParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/leave">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Leave">Leave</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionLeaveParams">MessagingConversationActionLeaveParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/mute">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Mute">Mute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionMuteParams">MessagingConversationActionMuteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/read">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionReadParams">MessagingConversationActionReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/redact">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Redact">Redact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionRedactParams">MessagingConversationActionRedactParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/report">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Report">Report</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionReportParams">MessagingConversationActionReportParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/set-legal-hold">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.SetLegalHold">SetLegalHold</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionSetLegalHoldParams">MessagingConversationActionSetLegalHoldParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/set-status">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.SetStatus">SetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionSetStatusParams">MessagingConversationActionSetStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/unarchive">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Unarchive">Unarchive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionUnarchiveParams">MessagingConversationActionUnarchiveParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/unhide">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Unhide">Unhide</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionUnhideParams">MessagingConversationActionUnhideParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/unmute">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionService.Unmute">Unmute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationActionUnmuteParams">MessagingConversationActionUnmuteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Links

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AddConversationLinkRequestParam">AddConversationLinkRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ConversationLink">ConversationLink</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListConversationLink">ListConversationLink</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkDeleteResponse">MessagingConversationLinkDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/links">client.Messaging.Conversations.Links.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkNewParams">MessagingConversationLinkNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ConversationLink">ConversationLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations/{id}/links">client.Messaging.Conversations.Links.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkListParams">MessagingConversationLinkListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListConversationLink">ListConversationLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/conversations/{id}/links/{link_id}">client.Messaging.Conversations.Links.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, linkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkDeleteParams">MessagingConversationLinkDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationLinkDeleteResponse">MessagingConversationLinkDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessageAttachmentInputParam">MessageAttachmentInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SendMessageRequestParam">SendMessageRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessage">ListMessage</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/messages">client.Messaging.Conversations.Messages.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationMessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationMessageNewParams">MessagingConversationMessageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations/{id}/messages">client.Messaging.Conversations.Messages.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationMessageListParams">MessagingConversationMessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessage">ListMessage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Participants

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AddParticipantRequestParam">AddParticipantRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantDeleteResponse">MessagingConversationParticipantDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/participants">client.Messaging.Conversations.Participants.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantNewParams">MessagingConversationParticipantNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/conversations/{id}/participants/{pid}">client.Messaging.Conversations.Participants.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantDeleteParams">MessagingConversationParticipantDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantDeleteResponse">MessagingConversationParticipantDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateParticipantRoleRequestParam">UpdateParticipantRoleRequestParam</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/participants/{pid}/actions/set-role">client.Messaging.Conversations.Participants.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantActionService.SetRole">SetRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pid <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationParticipantActionSetRoleParams">MessagingConversationParticipantActionSetRoleParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Attachments

#### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateAttachmentUploadURLRequestParam">CreateAttachmentUploadURLRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AttachmentUploadTarget">AttachmentUploadTarget</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/attachments/actions/upload-url">client.Messaging.Conversations.Attachments.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationAttachmentActionService.UploadURL">UploadURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingConversationAttachmentActionUploadURLParams">MessagingConversationAttachmentActionUploadURLParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AttachmentUploadTarget">AttachmentUploadTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateDraftRequestParam">UpdateDraftRequestParam</a>

Methods:

- <code title="patch /v1/messaging/messages/{id}">client.Messaging.Messages.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageUpdateParams">MessagingMessageUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ApproveSendDraftRequestParam">ApproveSendDraftRequestParam</a>

Methods:

- <code title="post /v1/messaging/messages/{id}/actions/approve-send">client.Messaging.Messages.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageActionService.ApproveSend">ApproveSend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageActionApproveSendParams">MessagingMessageActionApproveSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/messages/{id}/actions/cancel">client.Messaging.Messages.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageActionService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageActionCancelParams">MessagingMessageActionCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/messages/{id}/actions/reject">client.Messaging.Messages.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageActionService.Reject">Reject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingMessageActionRejectParams">MessagingMessageActionRejectParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Groups

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateMessagingGroupRequestParam">CreateMessagingGroupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateMessagingGroupRequestParam">UpdateMessagingGroupRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessagingGroup">ListMessagingGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupDeleteResponse">MessagingGroupDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/groups">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupNewParams">MessagingGroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/groups/{id}">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/messaging/groups/{id}">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupUpdateParams">MessagingGroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/groups">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessagingGroup">ListMessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/groups/{id}">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupDeleteResponse">MessagingGroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Members

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AddMessagingGroupMemberRequestParam">AddMessagingGroupMemberRequestParam</a>

Methods:

- <code title="post /v1/messaging/groups/{id}/members">client.Messaging.Groups.Members.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupMemberService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupMemberNewParams">MessagingGroupMemberNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/groups/{id}/members/{member_id}">client.Messaging.Groups.Members.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupMemberService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memberID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroupMemberDeleteParams">MessagingGroupMemberDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Blocks

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BlockRequestParam">BlockRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#LocationTypeCode">LocationTypeCode</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountUser">AccountUser</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Consumption">Consumption</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Department">Department</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListConsumption">ListConsumption</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListLocation">ListLocation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachine">ListMachine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessagingBlock">ListMessagingBlock</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionStep">ListProductionStep</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScanningStation">ListScanningStation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#LocationTypeCode">LocationTypeCode</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Machine">Machine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlock">MessagingBlock</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionOutput">ProductionOutput</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionStep">ProductionStep</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScanningStation">ScanningStation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#User">User</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlockDeleteResponse">MessagingBlockDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/blocks">client.Messaging.Blocks.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlockService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlockNewParams">MessagingBlockNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlock">MessagingBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/blocks">client.Messaging.Blocks.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlockService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlockListParams">MessagingBlockListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMessagingBlock">ListMessagingBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/blocks/{id}">client.Messaging.Blocks.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlockService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingBlockDeleteResponse">MessagingBlockDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Preferences

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertNotificationPreferenceRequestParam">UpsertNotificationPreferenceRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListNotificationPreference">ListNotificationPreference</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationPreference">NotificationPreference</a>

Methods:

- <code title="put /v1/messaging/preferences">client.Messaging.Preferences.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingPreferenceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingPreferenceUpdateParams">MessagingPreferenceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#NotificationPreference">NotificationPreference</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/preferences">client.Messaging.Preferences.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingPreferenceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListNotificationPreference">ListNotificationPreference</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EmailDomains

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateEmailDomainRequestParam">CreateEmailDomainRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailDomain">EmailDomain</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEmailDomain">ListEmailDomain</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainDeleteResponse">MessagingEmailDomainDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/email-domains">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainNewParams">MessagingEmailDomainNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailDomain">EmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-domains/{id}">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailDomain">EmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-domains">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEmailDomain">ListEmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/email-domains/{id}">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainDeleteResponse">MessagingEmailDomainDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="post /v1/messaging/email-domains/{id}/actions/verify">client.Messaging.EmailDomains.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailDomainActionService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailDomain">EmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EmailInboxes

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateEmailInboxRequestParam">CreateEmailInboxRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateEmailInboxRequestParam">UpdateEmailInboxRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailInbox">EmailInbox</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEmailInbox">ListEmailInbox</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxDeleteResponse">MessagingEmailInboxDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/email-inboxes">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxNewParams">MessagingEmailInboxNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailInbox">EmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-inboxes/{id}">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxGetParams">MessagingEmailInboxGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailInbox">EmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/messaging/email-inboxes/{id}">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxUpdateParams">MessagingEmailInboxUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#EmailInbox">EmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-inboxes">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxListParams">MessagingEmailInboxListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListEmailInbox">ListEmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/email-inboxes/{id}">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MessagingEmailInboxDeleteResponse">MessagingEmailInboxDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

## AccountUsers

### SalesTargets

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateSalesTargetRequestParam">CreateSalesTargetRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertSalesTargetRequestParam">UpsertSalesTargetRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesTarget">ListSalesTarget</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesTarget">SalesTarget</a>

Methods:

- <code title="post /v1/sales/account-users/{id}/sales-targets">client.Sales.AccountUsers.SalesTargets.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountUserSalesTargetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountUserSalesTargetNewParams">SaleAccountUserSalesTargetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesTarget">SalesTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/account-users/{id}/sales-targets/{target_id}">client.Sales.AccountUsers.SalesTargets.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountUserSalesTargetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, targetID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountUserSalesTargetUpdateParams">SaleAccountUserSalesTargetUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesTarget">SalesTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-users/{id}/sales-targets">client.Sales.AccountUsers.SalesTargets.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountUserSalesTargetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleAccountUserSalesTargetListParams">SaleAccountUserSalesTargetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesTarget">ListSalesTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateCustomerRequestParam">UpdateCustomerRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Carrier">Carrier</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerContactInfo">CustomerContactInfo</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerDefaults">CustomerDefaults</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerFreightPreferences">CustomerFreightPreferences</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerLeadTime">CustomerLeadTime</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerNotificationPreferences">CustomerNotificationPreferences</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListCustomer">ListCustomer</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListServiceLevel">ListServiceLevel</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PaymentTerm">PaymentTerm</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ServiceLevel">ServiceLevel</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ShippingTerm">ShippingTerm</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerDeleteResponse">SaleCustomerDeleteResponse</a>

Methods:

- <code title="post /v1/sales/customers">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerNewParams">SaleCustomerNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerGetParams">SaleCustomerGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerUpdateParams">SaleCustomerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerListParams">SaleCustomerListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListCustomer">ListCustomer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerDeleteResponse">SaleCustomerDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers/{id}/lead-time">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerService.GetLeadTime">GetLeadTime</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CustomerLeadTime">CustomerLeadTime</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MergeCustomersRequestParam">MergeCustomersRequestParam</a>

Methods:

- <code title="post /v1/sales/customers/{id}/actions/merge">client.Sales.Customers.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerActionService.Merge">Merge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleCustomerActionMergeParams">SaleCustomerActionMergeParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Contacts

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FindContactByEmailRequestParam">FindContactByEmailRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ContactMatch">ContactMatch</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListContactMatch">ListContactMatch</a>

Methods:

- <code title="post /v1/sales/contacts/actions/find-by-email">client.Sales.Contacts.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleContactActionService.FindByEmail">FindByEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleContactActionFindByEmailParams">SaleContactActionFindByEmailParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListContactMatch">ListContactMatch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SalesOrders

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CheckoutSalesOrderRequestParam">CheckoutSalesOrderRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateSalesOrderLineInputParam">CreateSalesOrderLineInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateSalesOrderRequestParam">CreateSalesOrderRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuoteSalesOrderLineInputParam">QuoteSalesOrderLineInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuoteSalesOrderPricesRequestParam">QuoteSalesOrderPricesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderEmailContactInputParam">SalesOrderEmailContactInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateSalesOrderRequestParam">UpdateSalesOrderRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CheckoutSalesOrderResponse">CheckoutSalesOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ComputedRate">ComputedRate</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatedBy">CreatedBy</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Freight">Freight</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListQuotedSalesOrderLine">ListQuotedSalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListRecord">ListRecord</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesOrder">ListSalesOrder</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesOrderLine">ListSalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesOrderStatus">ListSalesOrderStatus</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OrderContact">OrderContact</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OrderDiscount">OrderDiscount</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuoteSalesOrderPricesResponse">QuoteSalesOrderPricesResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuotedSalesOrderLine">QuotedSalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Record">Record</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderLine">SalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderRelated">SalesOrderRelated</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderStageTotal">SalesOrderStageTotal</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderStatus">SalesOrderStatus</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderTotals">SalesOrderTotals</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderDeleteResponse">SaleSalesOrderDeleteResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderNewParams">SaleSalesOrderNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/sales-orders/{id}">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderGetParams">SaleSalesOrderGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/sales-orders/{id}">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderUpdateParams">SaleSalesOrderUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/sales-orders">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderListParams">SaleSalesOrderListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesOrder">ListSalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/sales-orders/{id}">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderDeleteResponse">SaleSalesOrderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/{id}/checkout">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.Checkout">Checkout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderCheckoutParams">SaleSalesOrderCheckoutParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CheckoutSalesOrderResponse">CheckoutSalesOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/price-quote">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.PriceQuote">PriceQuote</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderPriceQuoteParams">SaleSalesOrderPriceQuoteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuoteSalesOrderPricesResponse">QuoteSalesOrderPricesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/sales-orders/statuses">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderService.GetStatuses">GetStatuses</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderGetStatusesParams">SaleSalesOrderGetStatusesParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSalesOrderStatus">ListSalesOrderStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkDeleteSalesOrdersRequestParam">BulkDeleteSalesOrdersRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IssueSalesOrderRequestParam">IssueSalesOrderRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionRun">ProductionRun</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuoteSalesOrderFreightResponse">QuoteSalesOrderFreightResponse</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionBulkDeleteResponse">SaleSalesOrderActionBulkDeleteResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders/actions/bulk-delete">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionService.BulkDelete">BulkDelete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionBulkDeleteParams">SaleSalesOrderActionBulkDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionBulkDeleteResponse">SaleSalesOrderActionBulkDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/close">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/{id}/actions/create-production-run">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionService.NewProductionRun">NewProductionRun</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionNewProductionRunParams">SaleSalesOrderActionNewProductionRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionRun">ProductionRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/issue">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionService.Issue">Issue</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionIssueParams">SaleSalesOrderActionIssueParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/open">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionService.Open">Open</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/{id}/actions/quote-freight">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionService.QuoteFreight">QuoteFreight</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuoteSalesOrderFreightResponse">QuoteSalesOrderFreightResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/unissue">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderActionService.Unissue">Unissue</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Lines

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateSalesOrderLineRequestParam">CreateSalesOrderLineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateSalesOrderLineRequestParam">UpdateSalesOrderLineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineDeleteResponse">SaleSalesOrderLineDeleteResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders/{id}/lines">client.Sales.SalesOrders.Lines.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineNewParams">SaleSalesOrderLineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderLine">SalesOrderLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/sales-orders/{id}/lines/{line_id}">client.Sales.SalesOrders.Lines.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineUpdateParams">SaleSalesOrderLineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SalesOrderLine">SalesOrderLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/sales-orders/{id}/lines/{line_id}">client.Sales.SalesOrders.Lines.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineDeleteParams">SaleSalesOrderLineDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineDeleteResponse">SaleSalesOrderLineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReorderSalesOrderLinesRequestParam">ReorderSalesOrderLinesRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineActionReorderResponse">SaleSalesOrderLineActionReorderResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders/{id}/lines/actions/reorder">client.Sales.SalesOrders.Lines.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineActionService.Reorder">Reorder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineActionReorderParams">SaleSalesOrderLineActionReorderParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SaleSalesOrderLineActionReorderResponse">SaleSalesOrderLineActionReorderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DemandOverrideType">DemandOverrideType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDemandOverrideType">ListDemandOverrideType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachineDowntimeReason">ListMachineDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachineStatus">ListMachineStatus</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleDeviationType">ListScheduleDeviationType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineCampaign">MachineCampaign</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineDowntimeReason">MachineDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineDowntimeReasonSummary">MachineDowntimeReasonSummary</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineDowntimeSummary">MachineDowntimeSummary</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineStatus">MachineStatus</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleDeviationType">ScheduleDeviationType</a>

Methods:

- <code title="get /v1/operations/demand-override-types">client.Operations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationService.GetDemandOverrideTypes">GetDemandOverrideTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDemandOverrideType">ListDemandOverrideType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-downtime-reasons">client.Operations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationService.GetMachineDowntimeReasons">GetMachineDowntimeReasons</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachineDowntimeReason">ListMachineDowntimeReason</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-status">client.Operations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationService.GetMachineStatus">GetMachineStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationGetMachineStatusParams">OperationGetMachineStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachineStatus">ListMachineStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/schedule-deviation-types">client.Operations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationService.GetScheduleDeviationTypes">GetScheduleDeviationTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleDeviationType">ListScheduleDeviationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

## Departments

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateDepartmentRequestParam">CreateDepartmentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DepartmentRateInputParam">DepartmentRateInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateDepartmentRequestParam">UpdateDepartmentRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDepartment">ListDepartment</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentDeleteResponse">OperationDepartmentDeleteResponse</a>

Methods:

- <code title="post /v1/operations/departments">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentNewParams">OperationDepartmentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Department">Department</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/departments/{id}">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentGetParams">OperationDepartmentGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Department">Department</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/departments/{id}">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentUpdateParams">OperationDepartmentUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Department">Department</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/departments">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentListParams">OperationDepartmentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDepartment">ListDepartment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/departments/{id}">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDepartmentDeleteResponse">OperationDepartmentDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Machines

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateMachineRequestParam">CreateMachineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateMachineRequestParam">UpdateMachineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDeleteResponse">OperationMachineDeleteResponse</a>

Methods:

- <code title="post /v1/operations/machines">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineNewParams">OperationMachineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Machine">Machine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machines/{id}">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineGetParams">OperationMachineGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Machine">Machine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/machines/{id}">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineUpdateParams">OperationMachineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Machine">Machine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machines">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineListParams">OperationMachineListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachine">ListMachine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/machines/{id}">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDeleteResponse">OperationMachineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## MachineDowntimeEvents

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateMachineDowntimeEventRequestParam">CreateMachineDowntimeEventRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateMachineDowntimeEventRequestParam">UpdateMachineDowntimeEventRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachineDowntimeEvent">ListMachineDowntimeEvent</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineDowntimeEvent">MachineDowntimeEvent</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventDeleteResponse">OperationMachineDowntimeEventDeleteResponse</a>

Methods:

- <code title="post /v1/operations/machine-downtime-events">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventNewParams">OperationMachineDowntimeEventNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineDowntimeEvent">MachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-downtime-events/{id}">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventGetParams">OperationMachineDowntimeEventGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineDowntimeEvent">MachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/machine-downtime-events/{id}">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventUpdateParams">OperationMachineDowntimeEventUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#MachineDowntimeEvent">MachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-downtime-events">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventListParams">OperationMachineDowntimeEventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListMachineDowntimeEvent">ListMachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/machine-downtime-events/{id}">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationMachineDowntimeEventDeleteResponse">OperationMachineDowntimeEventDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DemandOverrides

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateDemandOverrideRequestParam">CreateDemandOverrideRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateDemandOverrideRequestParam">UpdateDemandOverrideRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DemandOverride">DemandOverride</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDemandOverride">ListDemandOverride</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideDeleteResponse">OperationDemandOverrideDeleteResponse</a>

Methods:

- <code title="post /v1/operations/demand-overrides">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideNewParams">OperationDemandOverrideNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DemandOverride">DemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/demand-overrides/{id}">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideGetParams">OperationDemandOverrideGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DemandOverride">DemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/demand-overrides/{id}">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideUpdateParams">OperationDemandOverrideUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DemandOverride">DemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/demand-overrides">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideListParams">OperationDemandOverrideListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDemandOverride">ListDemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/demand-overrides/{id}">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationDemandOverrideDeleteResponse">OperationDemandOverrideDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProductionSchedules

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#GenerateProductionScheduleRequestParam">GenerateProductionScheduleRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionSchedule">ListProductionSchedule</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleDerivedLine">ListProductionScheduleDerivedLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleDeviation">ListProductionScheduleDeviation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleFinishedPolicy">ListProductionScheduleFinishedPolicy</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleItemPolicy">ListProductionScheduleItemPolicy</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListReleaseScheduleBatch">ListReleaseScheduleBatch</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListReleasedScheduleLine">ListReleasedScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleAppliedOverride">ListScheduleAppliedOverride</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleAtRiskOrder">ListScheduleAtRiskOrder</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleOrderCoverage">ListScheduleOrderCoverage</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleOrderCoverageLine">ListScheduleOrderCoverageLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedule">ProductionSchedule</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleDerivedLine">ProductionScheduleDerivedLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleDeviation">ProductionScheduleDeviation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleFinishedPolicy">ProductionScheduleFinishedPolicy</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleItemPolicy">ProductionScheduleItemPolicy</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReleaseScheduleBatch">ReleaseScheduleBatch</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReleaseScheduleWeekPreview">ReleaseScheduleWeekPreview</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReleasedScheduleLine">ReleasedScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleAppliedOverride">ScheduleAppliedOverride</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleAtRiskOrder">ScheduleAtRiskOrder</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleDiagnostics">ScheduleDiagnostics</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleOrderCoverage">ScheduleOrderCoverage</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleOrderCoverageLine">ScheduleOrderCoverageLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleDeleteResponse">OperationProductionScheduleDeleteResponse</a>

Methods:

- <code title="post /v1/operations/production-schedules">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleNewParams">OperationProductionScheduleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleListParams">OperationProductionScheduleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionSchedule">ListProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedules/{id}">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleDeleteResponse">OperationProductionScheduleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/at-risk-orders">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.GetAtRiskOrders">GetAtRiskOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleOrderCoverage">ListScheduleOrderCoverage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/current">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.GetCurrent">GetCurrent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/derived-lines">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.GetDerivedLines">GetDerivedLines</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleGetDerivedLinesParams">OperationProductionScheduleGetDerivedLinesParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleDerivedLine">ListProductionScheduleDerivedLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/deviations">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.GetDeviations">GetDeviations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleGetDeviationsParams">OperationProductionScheduleGetDeviationsParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleDeviation">ListProductionScheduleDeviation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/finished-policies">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.GetFinishedPolicies">GetFinishedPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleFinishedPolicy">ListProductionScheduleFinishedPolicy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/item-policies">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.GetItemPolicies">GetItemPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleItemPolicy">ListProductionScheduleItemPolicy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/week-release-preview">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleService.GetWeekReleasePreview">GetWeekReleasePreview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleGetWeekReleasePreviewParams">OperationProductionScheduleGetWeekReleasePreviewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReleaseScheduleWeekPreview">ReleaseScheduleWeekPreview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Lines

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateProductionScheduleLineRequestParam">CreateProductionScheduleLineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateProductionScheduleLineRequestParam">UpdateProductionScheduleLineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleLine">ListProductionScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleLine">ProductionScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineDeleteResponse">OperationProductionScheduleLineDeleteResponse</a>

Methods:

- <code title="post /v1/operations/production-schedules/{id}/lines">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineNewParams">OperationProductionScheduleLineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleLine">ProductionScheduleLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/production-schedules/{id}/lines/{line_id}">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineUpdateParams">OperationProductionScheduleLineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleLine">ProductionScheduleLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/lines">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineListParams">OperationProductionScheduleLineListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleLine">ListProductionScheduleLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedules/{id}/lines/{line_id}">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineDeleteParams">OperationProductionScheduleLineDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleLineDeleteResponse">OperationProductionScheduleLineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PreviewProductionScheduleRequestParam">PreviewProductionScheduleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PreviewRegenerateProductionScheduleRequestParam">PreviewRegenerateProductionScheduleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#QuotePromiseDateRequestParam">QuotePromiseDateRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RegenerateProductionScheduleRequestParam">RegenerateProductionScheduleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReleaseProductionScheduleWeekRequestParam">ReleaseProductionScheduleWeekRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleCampaign">ListScheduleCampaign</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleDiffLine">ListScheduleDiffLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListSchedulePolicy">ListSchedulePolicy</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListScheduleProjection">ListScheduleProjection</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedulePreview">ProductionSchedulePreview</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleRegeneratePreview">ProductionScheduleRegeneratePreview</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PromiseDateQuote">PromiseDateQuote</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReleaseScheduleWeekResult">ReleaseScheduleWeekResult</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleCampaign">ScheduleCampaign</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleDiffLine">ScheduleDiffLine</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SchedulePolicy">SchedulePolicy</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ScheduleProjection">ScheduleProjection</a>

Methods:

- <code title="put /v1/operations/production-schedules/{id}/actions/archive">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/actions/preview">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionService.Preview">Preview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionPreviewParams">OperationProductionScheduleActionPreviewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedulePreview">ProductionSchedulePreview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/{id}/actions/preview-regenerate">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionService.PreviewRegenerate">PreviewRegenerate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionPreviewRegenerateParams">OperationProductionScheduleActionPreviewRegenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleRegeneratePreview">ProductionScheduleRegeneratePreview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/{id}/actions/publish">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionService.Publish">Publish</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/operations/production-schedules/actions/quote-promise-date">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionService.QuotePromiseDate">QuotePromiseDate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionQuotePromiseDateParams">OperationProductionScheduleActionQuotePromiseDateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PromiseDateQuote">PromiseDateQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/{id}/actions/regenerate">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionService.Regenerate">Regenerate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionRegenerateParams">OperationProductionScheduleActionRegenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/operations/production-schedules/{id}/actions/release-week">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionService.ReleaseWeek">ReleaseWeek</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleActionReleaseWeekParams">OperationProductionScheduleActionReleaseWeekParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ReleaseScheduleWeekResult">ReleaseScheduleWeekResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProductionScheduleSettings

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateProductionScheduleSettingsRequestParam">UpdateProductionScheduleSettingsRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleSettings">ProductionScheduleSettings</a>

Methods:

- <code title="put /v1/operations/production-schedule-settings">client.Operations.ProductionScheduleSettings.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingUpdateParams">OperationProductionScheduleSettingUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleSettings">ProductionScheduleSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedule-settings">client.Operations.ProductionScheduleSettings.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleSettings">ProductionScheduleSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Resources

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertResourceSettingRequestParam">UpsertResourceSettingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleResourceSetting">ListProductionScheduleResourceSetting</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleResourceSetting">ProductionScheduleResourceSetting</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingResourceDeleteResponse">OperationProductionScheduleSettingResourceDeleteResponse</a>

Methods:

- <code title="put /v1/operations/production-schedule-settings/resources">client.Operations.ProductionScheduleSettings.Resources.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingResourceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingResourceUpdateParams">OperationProductionScheduleSettingResourceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleResourceSetting">ProductionScheduleResourceSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedule-settings/resources">client.Operations.ProductionScheduleSettings.Resources.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingResourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleResourceSetting">ListProductionScheduleResourceSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedule-settings/resources/{id}">client.Operations.ProductionScheduleSettings.Resources.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingResourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingResourceDeleteResponse">OperationProductionScheduleSettingResourceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Items

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertItemSettingRequestParam">UpsertItemSettingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleItemSetting">ListProductionScheduleItemSetting</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleItemSetting">ProductionScheduleItemSetting</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingItemDeleteResponse">OperationProductionScheduleSettingItemDeleteResponse</a>

Methods:

- <code title="get /v1/operations/production-schedule-settings/items/{item_id}">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleItemSetting">ProductionScheduleItemSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedule-settings/items/{item_id}">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingItemService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingItemUpdateParams">OperationProductionScheduleSettingItemUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ProductionScheduleItemSetting">ProductionScheduleItemSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedule-settings/items">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListProductionScheduleItemSetting">ListProductionScheduleItemSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedule-settings/items/{item_id}">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingItemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationProductionScheduleSettingItemDeleteResponse">OperationProductionScheduleSettingItemDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FulfillmentRecommendations

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#FulfillmentRecommendation">FulfillmentRecommendation</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListFulfillmentRecommendation">ListFulfillmentRecommendation</a>

Methods:

- <code title="get /v1/operations/fulfillment-recommendations">client.Operations.FulfillmentRecommendations.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationFulfillmentRecommendationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListFulfillmentRecommendation">ListFulfillmentRecommendation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ApplyFulfillmentRecommendationsRequestParam">ApplyFulfillmentRecommendationsRequestParam</a>

Methods:

- <code title="post /v1/operations/fulfillment-recommendations/actions/apply">client.Operations.FulfillmentRecommendations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationFulfillmentRecommendationActionService.Apply">Apply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationFulfillmentRecommendationActionApplyParams">OperationFulfillmentRecommendationActionApplyParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListFulfillmentRecommendation">ListFulfillmentRecommendation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#BulkUpsertLocationsRequestParam">BulkUpsertLocationsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpsertLocationInputParam">UpsertLocationInputParam</a>

Methods:

- <code title="post /v1/operations/locations/actions/bulk-upsert">client.Operations.Locations.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationActionBulkUpsertParams">OperationLocationActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LocationTypes

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListLocationType">ListLocationType</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#LocationType">LocationType</a>

Methods:

- <code title="get /v1/operations/location-types/{id}">client.Operations.LocationTypes.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#LocationType">LocationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/location-types">client.Operations.LocationTypes.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationLocationTypeListParams">OperationLocationTypeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListLocationType">ListLocationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Shipments

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ParcelInputParam">ParcelInputParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RateShopRequestParam">RateShopRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListRateShopOption">ListRateShopOption</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RateShopOption">RateShopOption</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RateShopResult">RateShopResult</a>

Methods:

- <code title="post /v1/operations/shipments/actions/rate-shop">client.Operations.Shipments.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShipmentActionService.RateShop">RateShop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#OperationShipmentActionRateShopParams">OperationShipmentActionRateShopParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#RateShopResult">RateShopResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPermission">ListPermission</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPermissionGroup">ListPermissionGroup</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#Permission">Permission</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PermissionGroup">PermissionGroup</a>

Methods:

- <code title="get /v1/identity/permission-groups">client.Identity.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityService.GetPermissionGroups">GetPermissionGroups</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityGetPermissionGroupsParams">IdentityGetPermissionGroupsParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPermissionGroup">ListPermissionGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

## Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUpdateFaviconResponse">IdentityAccountUpdateFaviconResponse</a>

Methods:

- <code title="put /v1/identity/accounts/{id}/favicon">client.Identity.Accounts.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountService.UpdateFavicon">UpdateFavicon</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#IdentityAccountUpdateFaviconResponse">IdentityAccountUpdateFaviconResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

# Settings

## PortalDomains

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreatePortalDomainRequestParam">CreatePortalDomainRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#DNSRecord">DNSRecord</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListDNSRecord">ListDNSRecord</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPortalDomain">ListPortalDomain</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PortalDomain">PortalDomain</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainDeleteResponse">SettingPortalDomainDeleteResponse</a>

Methods:

- <code title="post /v1/settings/portal-domains">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainNewParams">SettingPortalDomainNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PortalDomain">PortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/settings/portal-domains/{id}">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PortalDomain">PortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/settings/portal-domains">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListPortalDomain">ListPortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/settings/portal-domains/{id}">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainDeleteResponse">SettingPortalDomainDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="post /v1/settings/portal-domains/{id}/actions/verify">client.Settings.PortalDomains.Actions.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingPortalDomainActionService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#PortalDomain">PortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Integrations

Params Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#CreateAccountIntegrationRequestParam">CreateAccountIntegrationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#UpdateAccountIntegrationRequestParam">UpdateAccountIntegrationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountIntegration">AccountIntegration</a>
- <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountIntegration">ListAccountIntegration</a>

Methods:

- <code title="post /v1/settings/integrations">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingIntegrationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingIntegrationNewParams">SettingIntegrationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountIntegration">AccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/settings/integrations/{id}">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingIntegrationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingIntegrationUpdateParams">SettingIntegrationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountIntegration">AccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/settings/integrations">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingIntegrationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingIntegrationListParams">SettingIntegrationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#ListAccountIntegration">ListAccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/settings/integrations/{id}">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/augno/augno-go#SettingIntegrationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/augno/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/augno/augno-go#AccountIntegration">AccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
