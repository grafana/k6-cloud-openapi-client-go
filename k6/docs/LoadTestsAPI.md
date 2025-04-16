# \LoadTestsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadTestsDestroy**](LoadTestsAPI.md#LoadTestsDestroy) | **Delete** /cloud/v6/load_tests/{id} | Delete a load test.
[**LoadTestsList**](LoadTestsAPI.md#LoadTestsList) | **Get** /cloud/v6/load_tests | List all tests.
[**LoadTestsMove**](LoadTestsAPI.md#LoadTestsMove) | **Put** /cloud/v6/load_tests/{id}/move | Move a test to another project.
[**LoadTestsPartialUpdate**](LoadTestsAPI.md#LoadTestsPartialUpdate) | **Patch** /cloud/v6/load_tests/{id} | Update a load test.
[**LoadTestsRetrieve**](LoadTestsAPI.md#LoadTestsRetrieve) | **Get** /cloud/v6/load_tests/{id} | Get a load test by ID.
[**LoadTestsScriptRetrieve**](LoadTestsAPI.md#LoadTestsScriptRetrieve) | **Get** /cloud/v6/load_tests/{id}/script | Download the test script.
[**LoadTestsScriptUpdate**](LoadTestsAPI.md#LoadTestsScriptUpdate) | **Put** /cloud/v6/load_tests/{id}/script | Upload the script for a test.
[**LoadTestsStart**](LoadTestsAPI.md#LoadTestsStart) | **Post** /cloud/v6/load_tests/{id}/start | Start a test in Grafana Cloud.
[**ProjectsLoadTestsCreate**](LoadTestsAPI.md#ProjectsLoadTestsCreate) | **Post** /cloud/v6/projects/{id}/load_tests | Create a new test.
[**ProjectsLoadTestsRetrieve**](LoadTestsAPI.md#ProjectsLoadTestsRetrieve) | **Get** /cloud/v6/projects/{id}/load_tests | List load tests in a project.
[**ValidateOptions**](LoadTestsAPI.md#ValidateOptions) | **Post** /cloud/v6/validate_options | Validate k6 script options.



## LoadTestsDestroy

> LoadTestsDestroy(ctx, id).XStackId(xStackId).Execute()

Delete a load test.



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
	id := int32(56) // int32 | ID of the load test.

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
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsDestroyRequest struct via the builder pattern


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


## LoadTestsList

> LoadTestListResponse LoadTestsList(ctx).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Name(name).Execute()

List all tests.



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
	count := true // bool | Include collection length in the response object as `@count`. (optional)
	orderby := "id desc,project_id" // string | Comma-separated list of fields to use when ordering the results. Available fields: - id - project_id - name - created - updated  The default ascending order can be reversed by appending the `desc` specifier. (optional)
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)
	name := "name_example" // string | Filter results by load test name (exact match). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsList(context.Background()).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsList`: LoadTestListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
 **count** | **bool** | Include collection length in the response object as &#x60;@count&#x60;. | 
 **orderby** | **string** | Comma-separated list of fields to use when ordering the results. Available fields: - id - project_id - name - created - updated  The default ascending order can be reversed by appending the &#x60;desc&#x60; specifier. | 
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]
 **name** | **string** | Filter results by load test name (exact match). | 

### Return type

[**LoadTestListResponse**](LoadTestListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadTestsMove

> LoadTestsMove(ctx, id).XStackId(xStackId).MoveLoadTestApiModel(moveLoadTestApiModel).Execute()

Move a test to another project.



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
	id := int32(56) // int32 | ID of the load test.
	moveLoadTestApiModel := *openapiclient.NewMoveLoadTestApiModel(int32(123)) // MoveLoadTestApiModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadTestsAPI.LoadTestsMove(context.Background(), id).XStackId(xStackId).MoveLoadTestApiModel(moveLoadTestApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **moveLoadTestApiModel** | [**MoveLoadTestApiModel**](MoveLoadTestApiModel.md) |  | 

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


## LoadTestsPartialUpdate

> LoadTestsPartialUpdate(ctx, id).XStackId(xStackId).PatchLoadTestApiModel(patchLoadTestApiModel).Execute()

Update a load test.



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
	id := int32(56) // int32 | ID of the load test.
	patchLoadTestApiModel := *openapiclient.NewPatchLoadTestApiModel() // PatchLoadTestApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadTestsAPI.LoadTestsPartialUpdate(context.Background(), id).XStackId(xStackId).PatchLoadTestApiModel(patchLoadTestApiModel).Execute()
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
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **patchLoadTestApiModel** | [**PatchLoadTestApiModel**](PatchLoadTestApiModel.md) |  | 

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

> LoadTestApiModel LoadTestsRetrieve(ctx, id).XStackId(xStackId).Execute()

Get a load test by ID.



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
	id := int32(56) // int32 | ID of the load test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsRetrieve`: LoadTestApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**LoadTestApiModel**](LoadTestApiModel.md)

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

Download the test script.



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
	id := int32(56) // int32 | ID of the load test.

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
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsScriptRetrieveRequest struct via the builder pattern


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


## LoadTestsScriptUpdate

> LoadTestsScriptUpdate(ctx, id).XStackId(xStackId).Body(body).Execute()

Upload the script for a test.



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
	id := int32(56) // int32 | ID of the load test.
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
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsScriptUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

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

> TestRunApiModel LoadTestsStart(ctx, id).XStackId(xStackId).Execute()

Start a test in Grafana Cloud.



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
	id := int32(56) // int32 | ID of the load test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.LoadTestsStart(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.LoadTestsStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsStart`: TestRunApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.LoadTestsStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsStartRequest struct via the builder pattern


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


## ProjectsLoadTestsCreate

> LoadTestApiModel ProjectsLoadTestsCreate(ctx, id).XStackId(xStackId).Name(name).Script(script).Execute()

Create a new test.



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
	id := int32(56) // int32 | ID of the project.
	name := "name_example" // string | Unique name of the test within the project.
	script := os.NewFile(1234, "some_file") // *os.File | Test script in the form of a UTF-8 encoded text or a k6 .tar archive.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.ProjectsLoadTestsCreate(context.Background(), id).XStackId(xStackId).Name(name).Script(script).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.ProjectsLoadTestsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLoadTestsCreate`: LoadTestApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.ProjectsLoadTestsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLoadTestsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **name** | **string** | Unique name of the test within the project. | 
 **script** | ***os.File** | Test script in the form of a UTF-8 encoded text or a k6 .tar archive. | 

### Return type

[**LoadTestApiModel**](LoadTestApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLoadTestsRetrieve

> LoadTestListResponse ProjectsLoadTestsRetrieve(ctx, id).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()

List load tests in a project.



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
	id := int32(56) // int32 | ID of the project.
	count := true // bool | Include collection length in the response object as `@count`. (optional)
	orderby := "id desc,project_id" // string | Comma-separated list of fields to use when ordering the results. Available fields: - id - project_id - name - created - updated  The default ascending order can be reversed by appending the `desc` specifier. (optional)
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadTestsAPI.ProjectsLoadTestsRetrieve(context.Background(), id).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadTestsAPI.ProjectsLoadTestsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLoadTestsRetrieve`: LoadTestListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadTestsAPI.ProjectsLoadTestsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLoadTestsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **count** | **bool** | Include collection length in the response object as &#x60;@count&#x60;. | 
 **orderby** | **string** | Comma-separated list of fields to use when ordering the results. Available fields: - id - project_id - name - created - updated  The default ascending order can be reversed by appending the &#x60;desc&#x60; specifier. | 
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]

### Return type

[**LoadTestListResponse**](LoadTestListResponse.md)

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

Validate k6 script options.



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
	validateOptionsRequest := *openapiclient.NewValidateOptionsRequest(*openapiclient.NewOptions()) // ValidateOptionsRequest | 

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
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
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

