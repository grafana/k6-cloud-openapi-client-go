# \TestRunsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadTestsTestRunsRetrieve**](TestRunsAPI.md#LoadTestsTestRunsRetrieve) | **Get** /cloud/v6/load_tests/{id}/test_runs | List all runs of a load test.
[**TestRunsAbort**](TestRunsAPI.md#TestRunsAbort) | **Post** /cloud/v6/test_runs/{id}/abort | Abort a running test.
[**TestRunsDestroy**](TestRunsAPI.md#TestRunsDestroy) | **Delete** /cloud/v6/test_runs/{id} | Delete a test run.
[**TestRunsDistributionRetrieve**](TestRunsAPI.md#TestRunsDistributionRetrieve) | **Get** /cloud/v6/test_runs/{id}/distribution | Get test run distribution.
[**TestRunsList**](TestRunsAPI.md#TestRunsList) | **Get** /cloud/v6/test_runs | List all test runs.
[**TestRunsPartialUpdate**](TestRunsAPI.md#TestRunsPartialUpdate) | **Patch** /cloud/v6/test_runs/{id} | Update a test run.
[**TestRunsRetrieve**](TestRunsAPI.md#TestRunsRetrieve) | **Get** /cloud/v6/test_runs/{id} | Get a test run by ID.
[**TestRunsSave**](TestRunsAPI.md#TestRunsSave) | **Post** /cloud/v6/test_runs/{id}/save | Save test run results.
[**TestRunsScriptRetrieve**](TestRunsAPI.md#TestRunsScriptRetrieve) | **Get** /cloud/v6/test_runs/{id}/script | Download the test run script.
[**TestRunsUnsave**](TestRunsAPI.md#TestRunsUnsave) | **Post** /cloud/v6/test_runs/{id}/unsave | Unsave test run results.



## LoadTestsTestRunsRetrieve

> TestRunListResponse LoadTestsTestRunsRetrieve(ctx, id).XStackId(xStackId).Count(count).Skip(skip).Top(top).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

List all runs of a load test.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test.
	count := true // bool | Include collection length in the response object as `@count`. (optional)
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)
	createdAfter := time.Now() // time.Time | Filter test runs created on or after this date and time (inclusive). (optional)
	createdBefore := time.Now() // time.Time | Filter test runs created before this date and time (non-inclusive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.LoadTestsTestRunsRetrieve(context.Background(), id).XStackId(xStackId).Count(count).Skip(skip).Top(top).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.LoadTestsTestRunsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsTestRunsRetrieve`: TestRunListResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.LoadTestsTestRunsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsTestRunsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **count** | **bool** | Include collection length in the response object as &#x60;@count&#x60;. | 
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]
 **createdAfter** | **time.Time** | Filter test runs created on or after this date and time (inclusive). | 
 **createdBefore** | **time.Time** | Filter test runs created before this date and time (non-inclusive). | 

### Return type

[**TestRunListResponse**](TestRunListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsAbort

> TestRunsAbort(ctx, id).XStackId(xStackId).Execute()

Abort a running test.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.TestRunsAbort(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsDestroy

> TestRunsDestroy(ctx, id).XStackId(xStackId).Execute()

Delete a test run.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.TestRunsDestroy(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsDistributionRetrieve

> TestRunDistributionApiModel TestRunsDistributionRetrieve(ctx, id).XStackId(xStackId).Execute()

Get test run distribution.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.TestRunsDistributionRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsDistributionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsDistributionRetrieve`: TestRunDistributionApiModel
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.TestRunsDistributionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsDistributionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**TestRunDistributionApiModel**](TestRunDistributionApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsList

> TestRunListResponse TestRunsList(ctx).XStackId(xStackId).Count(count).Skip(skip).Top(top).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

List all test runs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	count := true // bool | Include collection length in the response object as `@count`. (optional)
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)
	createdAfter := time.Now() // time.Time | Filter test runs created on or after this date and time (inclusive). (optional)
	createdBefore := time.Now() // time.Time | Filter test runs created before this date and time (non-inclusive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.TestRunsList(context.Background()).XStackId(xStackId).Count(count).Skip(skip).Top(top).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsList`: TestRunListResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.TestRunsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
 **count** | **bool** | Include collection length in the response object as &#x60;@count&#x60;. | 
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]
 **createdAfter** | **time.Time** | Filter test runs created on or after this date and time (inclusive). | 
 **createdBefore** | **time.Time** | Filter test runs created before this date and time (non-inclusive). | 

### Return type

[**TestRunListResponse**](TestRunListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsPartialUpdate

> TestRunsPartialUpdate(ctx, id).XStackId(xStackId).PatchTestRunApiModel(patchTestRunApiModel).Execute()

Update a test run.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.
	patchTestRunApiModel := *openapiclient.NewPatchTestRunApiModel("Note_example") // PatchTestRunApiModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.TestRunsPartialUpdate(context.Background(), id).XStackId(xStackId).PatchTestRunApiModel(patchTestRunApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **patchTestRunApiModel** | [**PatchTestRunApiModel**](PatchTestRunApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsRetrieve

> TestRunApiModel TestRunsRetrieve(ctx, id).XStackId(xStackId).Execute()

Get a test run by ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.TestRunsRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsRetrieve`: TestRunApiModel
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.TestRunsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**TestRunApiModel**](TestRunApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsSave

> TestRunsSave(ctx, id).XStackId(xStackId).Execute()

Save test run results.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.TestRunsSave(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsSave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsScriptRetrieve

> string TestRunsScriptRetrieve(ctx, id).XStackId(xStackId).Execute()

Download the test run script.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.TestRunsScriptRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsScriptRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsScriptRetrieve`: string
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.TestRunsScriptRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsScriptRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

**string**

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/javascript, application/x-tar, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsUnsave

> TestRunsUnsave(ctx, id).XStackId(xStackId).Execute()

Unsave test run results.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/k6-cloud-openapi-client/k6"
)

func main() {
	xStackId := int32(56) // int32 | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
	id := int32(56) // int32 | ID of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.TestRunsUnsave(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsUnsave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsUnsaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

