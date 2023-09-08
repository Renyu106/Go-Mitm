# Go-Mitm

A Go application that demonstrates how to use `go-mitmproxy` to intercept, modify, and display HTTP requests and responses.

## Features

1. **URL Rewriting**: Redirects requests intended for `naver.com` to a local web backend running on `localhost:51231`.
2. **Local Web Backend**: A simple HTTP server that listens on `localhost:51231`, prints the details of incoming requests, and sends a response with those details.
3. **Proxy Server**: Uses `go-mitmproxy` to start a proxy server on port `:9080`.

## How it works

1. The `RewriteHost` struct implements the `Requestheaders` method of the `go-mitmproxy` Addon interface. This method checks if the incoming request's host is `naver.com`. If it is, the request's URL and scheme are modified to target the local web backend.
2. The `runWebBackend` function initializes an HTTP server on `localhost:51231`. The server uses the `handleRequest` function to handle incoming requests.
3. The `printRequestDetails` function is used to print the details (headers and body) of the incoming request to the console.
4. The `writeResponseDetails` function sends a response to the requester with the request's details.
5. The `main` function sets up and starts the proxy server. The proxy server uses two addons: `RewriteHost` (for URL rewriting) and `proxy.LogAddon` (for logging).

## Usage

To run the application:

1. Ensure you have the `go-mitmproxy` package installed.
2. Run the main application:
```bash
go run main.go
```
3. Configure your browser or HTTP client to use a proxy server with the address `localhost:9080`.
4. Make a request to `naver.com`. You should see the request details printed in the console, and the response will contain the same details.
