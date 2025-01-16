# k6-cloud-openapi-client-go

Go client for the
[Grafana Cloud k6 REST API v6](https://grafana.com/docs/grafana-cloud/testing/k6/reference/cloud-rest-api/v6/).

_Code generated with the [openapi-generator](https://openapi-generator.tech/) and the OpenAPI specification that can be
found at https://api.k6.io/cloud/v6/openapi._

## Getting started

### Installation

Install the following dependency:

```sh
go get github.com/grafana/k6-cloud-openapi-client-go
```

### Usage

To use it, you need to import the main `k6` package:

```go
import k6 "github.com/grafana/k6-cloud-openapi-client-go/k6"
```

#### Configuration

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

Alternatively, you can define your own server list with:

```go
cfg := k6.NewConfiguration()
cfg.Servers = []k6.ServerConfiguration{
    {URL: "<your-server-url>"},
}
```

#### Initialization

Once you have a configuration, you can initialize the client:

```go
client := k6.NewAPIClient(cfg)
```

#### Authentication

The authentication schema used by the API is HTTP Bearer token authentication and the way to provide the token
in the client is by using the `k6.ContextAccessToken` context key, as follows:


```go
// You may want to use `context.Background()` if your application doesn't have a context yet.
ctx = context.WithValue(ctx, k6.ContextAccessToken, "<your-token>")
```

#### Using the client's APIs

Below, you can see a **full example** of how to use the client to get the list of all the projects:

```go
package main

import (
	"context"
	"fmt"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

func main() {
	// First, we initialize the configuration and the client:
	cfg := k6.NewConfiguration()
	client := k6.NewAPIClient(cfg)

	// Then, we prepare the context with the Bearer token:
	ctx := context.WithValue(context.Background(), k6.ContextAccessToken, "<your-token>")

	// Now, we can build the base request:
	req := client.ProjectsAPI.ProjectsList(ctx).
		// Specify what stack id we want to make the request for:
		XStackId(123 /*<your-stack-id>*/).
		// Specify some optional parameters, like sorting by creation date:
		Orderby("created").
		// Request the total amount of projects to be present in the response:
		Count(true)

	// And finally, execute the request:
	projectsRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsList request failed: %s", err.Error())
	}

	// We print the status code and the total amount of projects, if present,
	// as well as the name of each project:
	log.Printf("Status code: %d", httpRes.StatusCode)
	if projectsRes.HasCount() {
		log.Printf("Total amount of projects: %d", *projectsRes.Count)
	}

	log.Println("The list of available projects is:")
	for _, p := range projectsRes.Value {
		log.Printf("%s\n", p.GetName())
	}
}
```

#### Examples

If you want to see more working examples, navigate to the [examples](./examples) directory.

#### Docs

If you want to see the docs of the APIs, navigate to the
[./k6/README.md](k6/README.md#documentation-for-api-endpoints).

## Other

### Versioning

This project follows the [Semantic Versioning](https://semver.org/) guidelines, 
using [pre-release syntax](https://semver.org/#spec-item-9) to indicate the client's version. 

So, for instance, we can fix a bug in the client if needed, while keeping it targeting the same API version.

For instance, given `1.0.0-0.1.0`:
- The API version (from spec) would be: `1.0.0`
- The client version would be `0.1.0`.

### Contributing

In you case you're willing to contribute to this project, navigate to [DEVELOPMENT](./DEVELOPMENT.md).
