<p align="center">
	<img src="./docs/jagw-logo.png">
</p>
<h1 align="center">JAGW-Go</h1>
<p align="center">
	<a href="https://github.com/jalapeno-api-gateway/jagw-go/releases/latest"><img src="https://img.shields.io/github/v/tag/jalapeno-api-gateway/jagw-go.svg?label=release&logo=github&style=flat-square" alt="Latest"></a>
	<a href="https://pkg.go.dev/github.com/jalapeno-api-gateway/jagw-go"><img src="https://pkg.go.dev/badge/github.com/jalapeno-api-gateway/jagw-go" alt="PkgGoDev"></a>
</p>

<p align="center">
The JAGW-Go repository contains a Go client library for interacting with the Jalapeno API Gateway
</p>

---

## Documentation
Go reference is available at https://pkg.go.dev/github.com/jalapeno-api-gateway/jagw-go. More documentation can be found under https://jalapeno-api-gateway.github.io/jagw.
The mock files are generated with https://github.com/uber-go/mock.

## Example
Here is a list of code examples with short description of demonstrated JAGW-Go functionality:

- [request-all-nodes](examples/request_all_nodes.go) - request all lsnodes from the gateway

All code examples can be found under [examples](examples) directory.

## Quick Start
Below are short code samples showing a JAGW-Go client interacting with the Jalapeno API Gateway.

```go
// Define the endpoint request/subscription
endpoint := jagw.JagwEndpoint{
    EndpointAddress: "localhost",
    EndpointPort: 9902,
}
// Get a connection
connection, err := jagw.NewJagwConnection(endpoint)
if err != nil {
    panic(err)
}
defer connection.Close()

// Create a new request client
client := jagw.NewRequestServiceClient(connection)
```
Complete example in [request_all_nodes.go](examples/request_all_nodes.go)

## How to contribute?

- Contribute code by submitting a [Pull Request](https://github.com/jalapeno-api-gateway/jagw-go/pulls).
- Report bugs by opening an [Issue](https://github.com/jalapeno-api-gateway/jagw-go/issues).
- Ask questions & open discussions by starting a [Discussion](https://github.com/jalapeno-api-gateway/jagw-go/discussions).