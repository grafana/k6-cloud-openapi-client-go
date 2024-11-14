# \TestRunsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestRunsAbort**](TestRunsAPI.md#TestRunsAbort) | **Post** /test_runs/{id}/abort | 
[**TestRunsDestroy**](TestRunsAPI.md#TestRunsDestroy) | **Delete** /test_runs/{id} | 
[**TestRunsList**](TestRunsAPI.md#TestRunsList) | **Get** /test_runs | 
[**TestRunsNoteUpdate**](TestRunsAPI.md#TestRunsNoteUpdate) | **Put** /test_runs/{id}/note | 
[**TestRunsRetrieve**](TestRunsAPI.md#TestRunsRetrieve) | **Get** /test_runs/{id} | 
[**TestRunsSave**](TestRunsAPI.md#TestRunsSave) | **Post** /test_runs/{id}/save | 
[**TestRunsScriptRetrieve**](TestRunsAPI.md#TestRunsScriptRetrieve) | **Get** /test_runs/{id}/script | 
[**TestRunsUnsave**](TestRunsAPI.md#TestRunsUnsave) | **Post** /test_runs/{id}/unsave | 



## TestRunsAbort

> TestRunsAbort(ctx, id).XStackId(xStackId).Execute()





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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	id := int32(56) // int32 | Id of the load test run.

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
**id** | **int32** | Id of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	id := int32(56) // int32 | Id of the load test run.

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
**id** | **int32** | Id of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


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


## TestRunsList

> PaginatedTestRunList TestRunsList(ctx).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()





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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	count := true // bool | Include collection length in the response object as '@count'. (optional)
	orderby := "created desc,ended" // string | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending `desc` specifier. Available fields: - created - ended - retention_expiry - status - cost/total_vuh (optional) (default to "created desc")
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.TestRunsList(context.Background()).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsList`: PaginatedTestRunList
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.TestRunsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **count** | **bool** | Include collection length in the response object as &#39;@count&#39;. | 
 **orderby** | **string** | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending &#x60;desc&#x60; specifier. Available fields: - created - ended - retention_expiry - status - cost/total_vuh | [default to &quot;created desc&quot;]
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]

### Return type

[**PaginatedTestRunList**](PaginatedTestRunList.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsNoteUpdate

> TestRunsNoteUpdate(ctx, id).XStackId(xStackId).TestRunUpdateNoteRequest(testRunUpdateNoteRequest).Execute()





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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	id := int32(56) // int32 | Id of the load test run.
	testRunUpdateNoteRequest := *openapiclient.NewTestRunUpdateNoteRequest("Value_example") // TestRunUpdateNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.TestRunsNoteUpdate(context.Background(), id).XStackId(xStackId).TestRunUpdateNoteRequest(testRunUpdateNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsNoteUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsNoteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 

 **testRunUpdateNoteRequest** | [**TestRunUpdateNoteRequest**](TestRunUpdateNoteRequest.md) |  | 

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

> TestRun TestRunsRetrieve(ctx, id).XStackId(xStackId).Execute()





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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	id := int32(56) // int32 | Id of the load test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.TestRunsRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.TestRunsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsRetrieve`: TestRun
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.TestRunsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


### Return type

[**TestRun**](TestRun.md)

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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	id := int32(56) // int32 | Id of the load test run.

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
**id** | **int32** | Id of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	id := int32(56) // int32 | Id of the load test run.

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
**id** | **int32** | Id of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsScriptRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


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
	xStackId := "xStackId_example" // string | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. 
	id := int32(56) // int32 | Id of the load test run.

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
**id** | **int32** | Id of the load test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsUnsaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


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

