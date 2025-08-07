# Customers

## Addresses

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#GetCustomerAddress">GetCustomerAddress</a>

Methods:

- <code title="post /customers/{customer_id}/addresses">client.Customers.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressNewParams">CustomerAddressNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{customer_id}/addresses/{address_id}">client.Customers.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, addressID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressGetParams">CustomerAddressGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#GetCustomerAddress">GetCustomerAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customers/{customer_id}/addresses/{address_id}">client.Customers.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, addressID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressUpdateParams">CustomerAddressUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{customer_id}/addresses">client.Customers.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#GetCustomerAddress">GetCustomerAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /customers/{customer_id}/addresses/{address_id}">client.Customers.Addresses.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, addressID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#CustomerAddressDeleteParams">CustomerAddressDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go">augno</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/augno-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
