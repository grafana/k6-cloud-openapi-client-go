# k6-cloud-openapi-client examples

Here you can find a working example of how to use the `k6-cloud-openapi-client`.

## Instructions

The [`main.go`](./main.go) file contains a simple `main` function where the client is initialized and used to list all
the load tests (*`GET /load_tests`*), the results are printed through the standard output.

### How to run the example

> [!IMPORTANT]  
> This example assumes you have the service forwarded to a local port.
> You can do that with `kubectl --context=dev-us-east-0 --namespace k6-cloud port-forward service/general-api 8080:80`.
> Otherwise, you'll need to adjust the `URL` specified in the `main.go` file.

This example tries to read the _token_ and the _stack id_, which are required, from environment variables.

So, you can run it with:

```bash
K6_CLOUD_TOKEN=<your-token> K6_CLOUD_STACK_ID=<stack-or-instance-id> go run main.go
```