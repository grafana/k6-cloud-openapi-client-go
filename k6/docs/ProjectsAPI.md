# \ProjectsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadTestsList**](ProjectsAPI.md#LoadTestsList) | **Get** /projects/{id}/load_tests | 
[**ProjectDelete**](ProjectsAPI.md#ProjectDelete) | **Delete** /projects/{id} | 
[**ProjectRetrieve**](ProjectsAPI.md#ProjectRetrieve) | **Get** /projects/{id} | 
[**ProjectSetDefault**](ProjectsAPI.md#ProjectSetDefault) | **Post** /projects/{id}/set_default | 
[**ProjectUpdate**](ProjectsAPI.md#ProjectUpdate) | **Patch** /projects/{id} | 
[**ProjectsCreate**](ProjectsAPI.md#ProjectsCreate) | **Post** /projects | 
[**ProjectsList**](ProjectsAPI.md#ProjectsList) | **Get** /projects | 
[**ProjectsLoadTestsCreate**](ProjectsAPI.md#ProjectsLoadTestsCreate) | **Post** /projects/{id}/load_tests | 



## LoadTestsList

> PaginatedLoadTestList LoadTestsList(ctx, id).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()





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
	id := int32(56) // int32 | Id of the project
	count := true // bool | Include collection length in the response object as '@count'. (optional)
	orderby := "name desc,created" // string | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending `desc` specifier. Available fields: - name - created - updated (optional) (default to "created desc")
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.LoadTestsList(context.Background(), id).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.LoadTestsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTestsList`: PaginatedLoadTestList
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.LoadTestsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTestsListRequest struct via the builder pattern


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


## ProjectDelete

> ProjectDelete(ctx, id).XStackId(xStackId).Execute()





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
	id := int64(789) // int64 | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectDelete(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeleteRequest struct via the builder pattern


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


## ProjectRetrieve

> Project ProjectRetrieve(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectRetrieve`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


### Return type

[**Project**](Project.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectSetDefault

> ProjectSetDefault(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectSetDefault(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectSetDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectSetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 


### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectUpdate

> ProjectUpdate(ctx, id).XStackId(xStackId).ProjectUpdateRequest(projectUpdateRequest).Execute()





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
	id := int32(56) // int32 | Id of the project
	projectUpdateRequest := *openapiclient.NewProjectUpdateRequest() // ProjectUpdateRequest | New values for an existing `Project` item. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectUpdate(context.Background(), id).XStackId(xStackId).ProjectUpdateRequest(projectUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 

 **projectUpdateRequest** | [**ProjectUpdateRequest**](ProjectUpdateRequest.md) | New values for an existing &#x60;Project&#x60; item. | 

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


## ProjectsCreate

> Project ProjectsCreate(ctx).CreateProjectRequest(createProjectRequest).Execute()





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
	createProjectRequest := *openapiclient.NewCreateProjectRequest("Name_example") // CreateProjectRequest | New `Project` item to be created (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsCreate(context.Background()).CreateProjectRequest(createProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCreate`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectRequest** | [**CreateProjectRequest**](CreateProjectRequest.md) | New &#x60;Project&#x60; item to be created | 

### Return type

[**Project**](Project.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> ProjectsList200Response ProjectsList(ctx).XStackId(xStackId).Orderby(orderby).Filter(filter).Top(top).Skip(skip).Count(count).Execute()





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
	orderby := []string{"Orderby_example"} // []string | Comma separated list of fields to use when ordering the results. Use `field_name desc` for descending ordering. (optional)
	filter := "filter_example" // string | Filter items by property values. (optional)
	top := int32(56) // int32 | Limit number of returned records. (optional)
	skip := int32(56) // int32 | Skip (offset) specified number of records before returning result. (optional)
	count := true // bool | Return count of items after applying any filtering. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsList(context.Background()).XStackId(xStackId).Orderby(orderby).Filter(filter).Top(top).Skip(skip).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsList`: ProjectsList200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **orderby** | **[]string** | Comma separated list of fields to use when ordering the results. Use &#x60;field_name desc&#x60; for descending ordering. | 
 **filter** | **string** | Filter items by property values. | 
 **top** | **int32** | Limit number of returned records. | 
 **skip** | **int32** | Skip (offset) specified number of records before returning result. | 
 **count** | **bool** | Return count of items after applying any filtering. | 

### Return type

[**ProjectsList200Response**](ProjectsList200Response.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLoadTestsCreate

> LoadTest ProjectsLoadTestsCreate(ctx, id).XStackId(xStackId).Name(name).Script(script).Execute()





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
	id := int32(56) // int32 | Id of the project
	name := "name_example" // string | Unique name of the test within the project
	script := os.NewFile(1234, "some_file") // *os.File | k6 script file that is executed when the test is started

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsLoadTestsCreate(context.Background(), id).XStackId(xStackId).Name(name).Script(script).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsLoadTestsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLoadTestsCreate`: LoadTest
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsLoadTestsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLoadTestsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 

 **name** | **string** | Unique name of the test within the project | 
 **script** | ***os.File** | k6 script file that is executed when the test is started | 

### Return type

[**LoadTest**](LoadTest.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

