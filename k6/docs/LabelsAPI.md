# \LabelsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LabelsCreate**](LabelsAPI.md#LabelsCreate) | **Post** /cloud/v6/labels | Create label keys.
[**LabelsDestroy**](LabelsAPI.md#LabelsDestroy) | **Delete** /cloud/v6/labels/{id} | Delete a label key.
[**LabelsList**](LabelsAPI.md#LabelsList) | **Get** /cloud/v6/labels | List label keys.
[**LabelsPartialUpdate**](LabelsAPI.md#LabelsPartialUpdate) | **Patch** /cloud/v6/labels/{id} | Update a label key.
[**ProjectsLabelsDestroy**](LabelsAPI.md#ProjectsLabelsDestroy) | **Delete** /cloud/v6/projects/{project_id}/labels/{key} | Remove a label from project.
[**ProjectsLabelsList**](LabelsAPI.md#ProjectsLabelsList) | **Get** /cloud/v6/projects/{project_id}/labels | List project labels.
[**ProjectsLabelsUpdate**](LabelsAPI.md#ProjectsLabelsUpdate) | **Put** /cloud/v6/projects/{project_id}/labels | Replace project labels.



## LabelsCreate

> LabelKeyListResponse LabelsCreate(ctx).XStackId(xStackId).LabelKeyCreateRequest(labelKeyCreateRequest).Execute()

Create label keys.



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
	labelKeyCreateRequest := *openapiclient.NewLabelKeyCreateRequest([]openapiclient.LabelKeyCreateItem{*openapiclient.NewLabelKeyCreateItem("Key_example")}) // LabelKeyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsCreate(context.Background()).XStackId(xStackId).LabelKeyCreateRequest(labelKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsCreate`: LabelKeyListResponse
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
 **labelKeyCreateRequest** | [**LabelKeyCreateRequest**](LabelKeyCreateRequest.md) |  | 

### Return type

[**LabelKeyListResponse**](LabelKeyListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsDestroy

> LabelsDestroy(ctx, id).XStackId(xStackId).Execute()

Delete a label key.



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
	id := int32(56) // int32 | ID of the label key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LabelsAPI.LabelsDestroy(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the label key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsDestroyRequest struct via the builder pattern


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


## LabelsList

> []LabelKeyListResponse LabelsList(ctx).XStackId(xStackId).Execute()

List label keys.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsList(context.Background()).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsList`: []LabelKeyListResponse
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

### Return type

[**[]LabelKeyListResponse**](LabelKeyListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsPartialUpdate

> LabelKeyResponse LabelsPartialUpdate(ctx, id).XStackId(xStackId).LabelKeyPatchRequest(labelKeyPatchRequest).Execute()

Update a label key.



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
	id := int32(56) // int32 | ID of the label key.
	labelKeyPatchRequest := *openapiclient.NewLabelKeyPatchRequest() // LabelKeyPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.LabelsPartialUpdate(context.Background(), id).XStackId(xStackId).LabelKeyPatchRequest(labelKeyPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.LabelsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelsPartialUpdate`: LabelKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.LabelsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the label key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **labelKeyPatchRequest** | [**LabelKeyPatchRequest**](LabelKeyPatchRequest.md) |  | 

### Return type

[**LabelKeyResponse**](LabelKeyResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLabelsDestroy

> ProjectsLabelsDestroy(ctx, key, projectId).XStackId(xStackId).Execute()

Remove a label from project.



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
	key := "key_example" // string | Label key name.
	projectId := int32(56) // int32 | ID of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LabelsAPI.ProjectsLabelsDestroy(context.Background(), key, projectId).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.ProjectsLabelsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Label key name. | 
**projectId** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLabelsDestroyRequest struct via the builder pattern


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


## ProjectsLabelsList

> []ProjectLabelListResponse ProjectsLabelsList(ctx, projectId).XStackId(xStackId).Execute()

List project labels.



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
	projectId := int32(56) // int32 | ID of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.ProjectsLabelsList(context.Background(), projectId).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.ProjectsLabelsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLabelsList`: []ProjectLabelListResponse
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.ProjectsLabelsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLabelsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**[]ProjectLabelListResponse**](ProjectLabelListResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLabelsUpdate

> ProjectLabelPutResponse ProjectsLabelsUpdate(ctx, projectId).XStackId(xStackId).ProjectLabelPutRequest(projectLabelPutRequest).Execute()

Replace project labels.



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
	projectId := int32(56) // int32 | ID of the project.
	projectLabelPutRequest := *openapiclient.NewProjectLabelPutRequest([]openapiclient.ProjectLabelPutItem{*openapiclient.NewProjectLabelPutItem("Value_example")}) // ProjectLabelPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.ProjectsLabelsUpdate(context.Background(), projectId).XStackId(xStackId).ProjectLabelPutRequest(projectLabelPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.ProjectsLabelsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLabelsUpdate`: ProjectLabelPutResponse
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.ProjectsLabelsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLabelsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **projectLabelPutRequest** | [**ProjectLabelPutRequest**](ProjectLabelPutRequest.md) |  | 

### Return type

[**ProjectLabelPutResponse**](ProjectLabelPutResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

