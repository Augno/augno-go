# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthRevokeRefreshTokenResponse">AuthRevokeRefreshTokenResponse</a>

Methods:

- <code title="post /v1/auth/access-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/auth/users">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthService.RegisterUser">RegisterUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthRegisterUserParams">AuthRegisterUserParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/auth/refresh-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthService.RevokeRefreshToken">RevokeRefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthRevokeRefreshTokenResponse">AuthRevokeRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#User">User</a>

Methods:

- <code title="post /v1/auth/actions/login">client.Auth.Actions.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthActionService.LoginUser">LoginUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthActionLoginUserParams">AuthActionLoginUserParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Passwords

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordUpdatePasswordResponse">AuthPasswordUpdatePasswordResponse</a>

Methods:

- <code title="put /v1/auth/passwords">client.Auth.Passwords.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordService.UpdatePassword">UpdatePassword</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordUpdatePasswordParams">AuthPasswordUpdatePasswordParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordUpdatePasswordResponse">AuthPasswordUpdatePasswordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionRequestPasswordResetResponse">AuthPasswordActionRequestPasswordResetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionResetPasswordResponse">AuthPasswordActionResetPasswordResponse</a>

Methods:

- <code title="post /v1/auth/passwords/actions/request-reset">client.Auth.Passwords.Actions.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionService.RequestPasswordReset">RequestPasswordReset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionRequestPasswordResetParams">AuthPasswordActionRequestPasswordResetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionRequestPasswordResetResponse">AuthPasswordActionRequestPasswordResetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/auth/passwords/actions/reset">client.Auth.Passwords.Actions.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionService.ResetPassword">ResetPassword</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionResetPasswordParams">AuthPasswordActionResetPasswordParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#AuthPasswordActionResetPasswordResponse">AuthPasswordActionResetPasswordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
