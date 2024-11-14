# \VariablesAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VariablesCreate**](VariablesAPI.md#VariablesCreate) | **Post** /variables | 
[**VariablesDestroy**](VariablesAPI.md#VariablesDestroy) | **Delete** /variables/{id} | 
[**VariablesList**](VariablesAPI.md#VariablesList) | **Get** /variables | 
[**VariablesPartialUpdate**](VariablesAPI.md#VariablesPartialUpdate) | **Patch** /variables/{id} | 



## VariablesCreate

> StackEnvVar VariablesCreate(ctx).CreateStackEnvVarRequestRequest(createStackEnvVarRequestRequest).Execute()





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
	createStackEnvVarRequestRequest := *openapiclient.NewCreateStackEnvVarRequestRequest("Name_example", "Value_example") // CreateStackEnvVarRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesAPI.VariablesCreate(context.Background()).CreateStackEnvVarRequestRequest(createStackEnvVarRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.VariablesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VariablesCreate`: StackEnvVar
	fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.VariablesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariablesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStackEnvVarRequestRequest** | [**CreateStackEnvVarRequestRequest**](CreateStackEnvVarRequestRequest.md) |  | 

### Return type

[**StackEnvVar**](StackEnvVar.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariablesDestroy

> VariablesDestroy(ctx, id).XStackId(xStackId).Execute()





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
	id := int32(56) // int32 | Id or name of the variable.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VariablesAPI.VariablesDestroy(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.VariablesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id or name of the variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVariablesDestroyRequest struct via the builder pattern


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


## VariablesList

> PaginatedStackEnvVarList VariablesList(ctx).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()





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
	orderby := "name desc,created" // string | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending `desc` specifier. Available fields: - name - created - updated (optional) (default to "name")
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariablesAPI.VariablesList(context.Background()).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.VariablesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VariablesList`: PaginatedStackEnvVarList
	fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.VariablesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariablesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **count** | **bool** | Include collection length in the response object as &#39;@count&#39;. | 
 **orderby** | **string** | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending &#x60;desc&#x60; specifier. Available fields: - name - created - updated | [default to &quot;name&quot;]
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]

### Return type

[**PaginatedStackEnvVarList**](PaginatedStackEnvVarList.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariablesPartialUpdate

> VariablesPartialUpdate(ctx, id).XStackId(xStackId).PatchedStackEnvVarRequest(patchedStackEnvVarRequest).Execute()





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
	id := int32(56) // int32 | Id or name of the variable.
	patchedStackEnvVarRequest := *openapiclient.NewPatchedStackEnvVarRequest() // PatchedStackEnvVarRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VariablesAPI.VariablesPartialUpdate(context.Background(), id).XStackId(xStackId).PatchedStackEnvVarRequest(patchedStackEnvVarRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.VariablesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id or name of the variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVariablesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 

 **patchedStackEnvVarRequest** | [**PatchedStackEnvVarRequest**](PatchedStackEnvVarRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

