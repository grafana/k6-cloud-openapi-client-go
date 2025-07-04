# \ProjectsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectLimitsRetrieve**](ProjectsAPI.md#ProjectLimitsRetrieve) | **Get** /cloud/v6/project-limits | List limits for all projects.
[**ProjectsCreate**](ProjectsAPI.md#ProjectsCreate) | **Post** /cloud/v6/projects | Create a project.
[**ProjectsDestroy**](ProjectsAPI.md#ProjectsDestroy) | **Delete** /cloud/v6/projects/{id} | Delete a project.
[**ProjectsLimitsPartialUpdate**](ProjectsAPI.md#ProjectsLimitsPartialUpdate) | **Patch** /cloud/v6/projects/{id}/limits | Update project limits.
[**ProjectsLimitsRetrieve**](ProjectsAPI.md#ProjectsLimitsRetrieve) | **Get** /cloud/v6/projects/{id}/limits | Get project limits.
[**ProjectsList**](ProjectsAPI.md#ProjectsList) | **Get** /cloud/v6/projects | List all projects.
[**ProjectsPartialUpdate**](ProjectsAPI.md#ProjectsPartialUpdate) | **Patch** /cloud/v6/projects/{id} | Update a project.
[**ProjectsRetrieve**](ProjectsAPI.md#ProjectsRetrieve) | **Get** /cloud/v6/projects/{id} | Get a project by ID.



## ProjectLimitsRetrieve

> ProjectLimitsListResponse ProjectLimitsRetrieve(ctx).XStackId(xStackId).Count(count).Skip(skip).Top(top).Execute()

List limits for all projects.



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
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectLimitsRetrieve(context.Background()).XStackId(xStackId).Count(count).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectLimitsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectLimitsRetrieve`: ProjectLimitsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectLimitsRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectLimitsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
 **count** | **bool** | Include collection length in the response object as &#x60;@count&#x60;. | 
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]

### Return type

[**ProjectLimitsListResponse**](ProjectLimitsListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreate

> ProjectApiModel ProjectsCreate(ctx).XStackId(xStackId).CreateProjectApiModel(createProjectApiModel).Execute()

Create a project.



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
	createProjectApiModel := *openapiclient.NewCreateProjectApiModel("Name_example") // CreateProjectApiModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsCreate(context.Background()).XStackId(xStackId).CreateProjectApiModel(createProjectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCreate`: ProjectApiModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
 **createProjectApiModel** | [**CreateProjectApiModel**](CreateProjectApiModel.md) |  | 

### Return type

[**ProjectApiModel**](ProjectApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDestroy

> ProjectsDestroy(ctx, id).XStackId(xStackId).Execute()

Delete a project.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsDestroy(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDestroyRequest struct via the builder pattern


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


## ProjectsLimitsPartialUpdate

> ProjectsLimitsPartialUpdate(ctx, id).XStackId(xStackId).PatchProjectLimitsRequest(patchProjectLimitsRequest).Execute()

Update project limits.



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
	patchProjectLimitsRequest := *openapiclient.NewPatchProjectLimitsRequest() // PatchProjectLimitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsLimitsPartialUpdate(context.Background(), id).XStackId(xStackId).PatchProjectLimitsRequest(patchProjectLimitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsLimitsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLimitsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **patchProjectLimitsRequest** | [**PatchProjectLimitsRequest**](PatchProjectLimitsRequest.md) |  | 

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


## ProjectsLimitsRetrieve

> ProjectLimitsApiModel ProjectsLimitsRetrieve(ctx, id).XStackId(xStackId).Execute()

Get project limits.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsLimitsRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsLimitsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLimitsRetrieve`: ProjectLimitsApiModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsLimitsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLimitsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**ProjectLimitsApiModel**](ProjectLimitsApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> ProjectListResponse ProjectsList(ctx).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Name(name).Execute()

List all projects.



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
	orderby := "created desc" // string | Comma-separated list of fields to use when ordering the results. Available fields: - created  The default ascending order can be reversed by appending the `desc` specifier. (optional)
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)
	name := "name_example" // string | Filter results by project name (exact match). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsList(context.Background()).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsList`: ProjectListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
 **count** | **bool** | Include collection length in the response object as &#x60;@count&#x60;. | 
 **orderby** | **string** | Comma-separated list of fields to use when ordering the results. Available fields: - created  The default ascending order can be reversed by appending the &#x60;desc&#x60; specifier. | 
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]
 **name** | **string** | Filter results by project name (exact match). | 

### Return type

[**ProjectListResponse**](ProjectListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPartialUpdate

> ProjectsPartialUpdate(ctx, id).XStackId(xStackId).PatchProjectApiModel(patchProjectApiModel).Execute()

Update a project.



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
	patchProjectApiModel := *openapiclient.NewPatchProjectApiModel("Name_example") // PatchProjectApiModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsPartialUpdate(context.Background(), id).XStackId(xStackId).PatchProjectApiModel(patchProjectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **patchProjectApiModel** | [**PatchProjectApiModel**](PatchProjectApiModel.md) |  | 

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


## ProjectsRetrieve

> ProjectApiModel ProjectsRetrieve(ctx, id).XStackId(xStackId).Execute()

Get a project by ID.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsRetrieve`: ProjectApiModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**ProjectApiModel**](ProjectApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

