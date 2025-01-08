# k6-cloud-openapi-client examples

Here there are some working examples that show how to use the `k6-cloud-openapi-client` in the form of
[Go Testable Examples](https://go.dev/blog/examples).

## Instructions

### Requirements

To run these examples, the following environment variables must be defined:

- `K6_CLOUD_TOKEN`: Personal API token or a Grafana Stack API token, to authenticate with the GCk6 API.
- `K6_CLOUD_STACK_ID`: The Grafana stack instance ID.
- _(Optional)_ `K6_CLOUD_URL`: The root URL for the Grafana Cloud k6 REST API. *Default is `https://api.k6.io`.*

_(Further details can be found
[here](https://grafana.com/docs/grafana-cloud/testing/k6/reference/cloud-rest-api/v6/#authentication-and-authorization))_

### How to run the examples

> [!IMPORTANT]
> Some examples provided here are not generic operations over collections, such as listing all projects for a given
> set of credentials and stack. Instead, they are operations on specific resources, like destroying (deleting) a
> project.
> These operations require you to specify the resource id in the request.
>
> The examples include placeholder values for demonstration purposes, but you must replace these with the actual
> resource ids from your account. If you do not, the examples will fail, returning errors like: `404 Not Found`.

You can run one of the examples, given its name (e.g. `ExampleProjectsAPI_ProjectsList`), by executing the
`go test` command as follows:

```sh
K6_CLOUD_TOKEN=<your-token> K6_CLOUD_STACK_ID=<stack-or-instance-id> go test ./examples/... -run ExampleProjectsAPI_ProjectsList -v
```

_Note the use of `-v` to display the output of the example, which contains some information retrieved from the API_

