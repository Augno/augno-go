# Healthz

Response Types:

- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#HealthzCheckResponse">HealthzCheckResponse</a>

Methods:

- <code title="get /healthz">client.Healthz.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#HealthzService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#HealthzCheckResponse">HealthzCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#EmptyResource">EmptyResource</a>

Methods:

- <code title="post /v2/auth/access-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#EmptyResource">EmptyResource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/auth/refresh-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthService.RevokeRefreshToken">RevokeRefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#EmptyResource">EmptyResource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthActionLoginUserResponse">AuthActionLoginUserResponse</a>

Methods:

- <code title="post /v2/auth/actions/login">client.Auth.Actions.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthActionService.LoginUser">LoginUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthActionLoginUserParams">AuthActionLoginUserParams</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthActionLoginUserResponse">AuthActionLoginUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
