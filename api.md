# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthRevokeRefreshTokenResponse">AuthRevokeRefreshTokenResponse</a>

Methods:

- <code title="post /v1/auth/access-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthService.RefreshToken">RefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthRefreshTokenResponse">AuthRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/auth/users">client.Auth.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthService.RegisterUser">RegisterUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthRegisterUserParams">AuthRegisterUserParams</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/auth/refresh-tokens">client.Auth.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthService.RevokeRefreshToken">RevokeRefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthRevokeRefreshTokenResponse">AuthRevokeRefreshTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#User">User</a>

Methods:

- <code title="post /v1/auth/actions/login">client.Auth.Actions.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthActionService.LoginUser">LoginUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthActionLoginUserParams">AuthActionLoginUserParams</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Passwords

Response Types:

- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordUpdatePasswordResponse">AuthPasswordUpdatePasswordResponse</a>

Methods:

- <code title="put /v1/auth/passwords">client.Auth.Passwords.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordService.UpdatePassword">UpdatePassword</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordUpdatePasswordParams">AuthPasswordUpdatePasswordParams</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordUpdatePasswordResponse">AuthPasswordUpdatePasswordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionRequestPasswordResetResponse">AuthPasswordActionRequestPasswordResetResponse</a>
- <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionResetPasswordResponse">AuthPasswordActionResetPasswordResponse</a>

Methods:

- <code title="post /v1/auth/passwords/actions/request-reset">client.Auth.Passwords.Actions.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionService.RequestPasswordReset">RequestPasswordReset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionRequestPasswordResetParams">AuthPasswordActionRequestPasswordResetParams</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionRequestPasswordResetResponse">AuthPasswordActionRequestPasswordResetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/auth/passwords/actions/reset">client.Auth.Passwords.Actions.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionService.ResetPassword">ResetPassword</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionResetPasswordParams">AuthPasswordActionResetPasswordParams</a>) (<a href="https://pkg.go.dev/github.com/Augno/go-sdk">augno</a>.<a href="https://pkg.go.dev/github.com/Augno/go-sdk#AuthPasswordActionResetPasswordResponse">AuthPasswordActionResetPasswordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
