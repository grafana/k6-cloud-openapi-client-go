# \LoadTestsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadTestsDestroy**](LoadTestsAPI.md#LoadTestsDestroy) | **Delete** /load_tests/{id} | 
[**LoadTestsPartialUpdate**](LoadTestsAPI.md#LoadTestsPartialUpdate) | **Patch** /load_tests/{id} | 
[**LoadTestsRetrieve**](LoadTestsAPI.md#LoadTestsRetrieve) | **Get** /load_tests/{id} | 
[**LoadTestsSchedulesRetrieve**](LoadTestsAPI.md#LoadTestsSchedulesRetrieve) | **Get** /load_tests/{id}/schedule | 
[**LoadTestsScriptRetrieve**](LoadTestsAPI.md#LoadTestsScriptRetrieve) | **Get** /load_tests/{id}/script | 
[**LoadTestsScriptUpdate**](LoadTestsAPI.md#LoadTestsScriptUpdate) | **Put** /load_tests/{id}/script | 
[**LoadTestsStart**](LoadTestsAPI.md#LoadTestsStart) | **Post** /load_tests/{id}/start | 
[**LoadTestsTestRunsList**](LoadTestsAPI.md#LoadTestsTestRunsList) | **Get** /load_tests/{id}/test_runs | 
[**ProjectsLoadTestsList**](LoadTestsAPI.md#ProjectsLoadTestsList) | **Get** /load_tests | 
[**ValidateOptions**](LoadTestsAPI.md#ValidateOptions) | **Post** /validate_options | 



## LoadTestsDestroy

> LoadTestsDestroy(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id of the load test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadTestsAPI.LoadTestsDestroy(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsDestroyRequest struct via the builder pattern


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


## LoadTestsPartialUpdate

> LoadTestsPartialUpdate(ctx, id).XStackId(xStackId).PatchedLoadTestRequest(patchedLoadTestRequest).Execute()





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
	id := int32(56) // int32 | Id of the load test.
	patchedLoadTestRequest := *openapiclient.NewPatchedLoadTestRequest() // PatchedLoadTestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadTestsAPI.LoadTestsPartialUpdate(context.Background(), id).XStackId(xStackId).PatchedLoadTestRequest(patchedLoadTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 

 **patchedLoadTestRequest** | [**PatchedLoadTestRequest**](PatchedLoadTestRequest.md) |  | 

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


## LoadTestsRetrieve

> LoadTest LoadTestsRetrieve(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id of the load test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsRetrieve`: LoadTest
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


### Return type

[**LoadTest**](LoadTest.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadTestsSchedulesRetrieve

> Schedule LoadTestsSchedulesRetrieve(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id of the load test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsSchedulesRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsSchedulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsSchedulesRetrieve`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsSchedulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsSchedulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


### Return type

[**Schedule**](Schedule.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadTestsScriptRetrieve

> string LoadTestsScriptRetrieve(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id of the load test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsScriptRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsScriptRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsScriptRetrieve`: string
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsScriptRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsScriptRetrieveRequest struct via the builder pattern


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


## LoadTestsScriptUpdate

> LoadTestsScriptUpdate(ctx, id).XStackId(xStackId).Body(body).Execute()





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
	id := int32(56) // int32 | Id of the load test.
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadTestsAPI.LoadTestsScriptUpdate(context.Background(), id).XStackId(xStackId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsScriptUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsScriptUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 

 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadTestsStart

> TestRun LoadTestsStart(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id of the load test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsStart(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsStart`: TestRun
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsStartRequest struct via the builder pattern


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


## LoadTestsTestRunsList

> PaginatedTestRunList LoadTestsTestRunsList(ctx, id).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()





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
	id := int32(56) // int32 | Id of the load test.
	count := true // bool | Include collection length in the response object as '@count'. (optional)
	orderby := "created desc,ended" // string | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending `desc` specifier. Available fields: - created - ended - retention_expiry - status - cost/total_vuh (optional) (default to "created desc")
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsTestRunsList(context.Background(), id).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsTestRunsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsTestRunsList`: PaginatedTestRunList
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsTestRunsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsTestRunsListRequest struct via the builder pattern


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


## ProjectsLoadTestsList

> PaginatedLoadTestList ProjectsLoadTestsList(ctx).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()





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
	orderby := "name desc,created" // string | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending `desc` specifier. Available fields: - name - created - updated (optional) (default to "created desc")
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.ProjectsLoadTestsList(context.Background()).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.ProjectsLoadTestsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLoadTestsList`: PaginatedLoadTestList
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.ProjectsLoadTestsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLoadTestsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **count** | **bool** | Include collection length in the response object as &#39;@count&#39;. | 
 **orderby** | **string** | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending &#x60;desc&#x60; specifier. Available fields: - name - created - updated | [default to &quot;created desc&quot;]
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]

### Return type

[**PaginatedLoadTestList**](PaginatedLoadTestList.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateOptions

> ValidateOptionsResponse ValidateOptions(ctx).XStackId(xStackId).ValidateOptionsRequest(validateOptionsRequest).Execute()





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
	validateOptionsRequest := *openapiclient.NewValidateOptionsRequest(map[string]interface{}{"key": interface{}(123)}) // ValidateOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.ValidateOptions(context.Background()).XStackId(xStackId).ValidateOptionsRequest(validateOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.ValidateOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateOptions`: ValidateOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.ValidateOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **validateOptionsRequest** | [**ValidateOptionsRequest**](ValidateOptionsRequest.md) |  | 

### Return type

[**ValidateOptionsResponse**](ValidateOptionsResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

