# Healthz

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#HealthzCheckResponse">HealthzCheckResponse</a>

Methods:

- <code title="get /healthz">client.Healthz.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#HealthzService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#HealthzCheckResponse">HealthzCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#EmptyResource">EmptyResource</a>

Methods:

- <code title="post /v2/auth/access-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#EmptyResource">EmptyResource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/auth/refresh-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthService.RevokeRefreshToken">RevokeRefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#EmptyResource">EmptyResource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthActionLoginUserResponse">AuthActionLoginUserResponse</a>

Methods:

- <code title="post /v2/auth/actions/login">client.Auth.Actions.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthActionService.LoginUser">LoginUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthActionLoginUserParams">AuthActionLoginUserParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthActionLoginUserResponse">AuthActionLoginUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
