# \LoadZonesAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadZonesAllowedProjectsRetrieve**](LoadZonesAPI.md#LoadZonesAllowedProjectsRetrieve) | **Get** /cloud/v6/load_zones/{id}/allowed_projects | List projects allowed to use a given private load zone.
[**LoadZonesAllowedProjectsUpdate**](LoadZonesAPI.md#LoadZonesAllowedProjectsUpdate) | **Put** /cloud/v6/load_zones/{id}/allowed_projects | Update the list of projects allowed to use a given private load zone.
[**LoadZonesList**](LoadZonesAPI.md#LoadZonesList) | **Get** /cloud/v6/load_zones | List all load zones.
[**ProjectsAllowedLoadZonesRetrieve**](LoadZonesAPI.md#ProjectsAllowedLoadZonesRetrieve) | **Get** /cloud/v6/projects/{id}/allowed_load_zones | List private load zones that can be used by a given project.
[**ProjectsAllowedLoadZonesUpdate**](LoadZonesAPI.md#ProjectsAllowedLoadZonesUpdate) | **Put** /cloud/v6/projects/{id}/allowed_load_zones | Update the list of private load zones that can be used by a given project.



## LoadZonesAllowedProjectsRetrieve

> AllowedProjectsListApiModel LoadZonesAllowedProjectsRetrieve(ctx, id).XStackId(xStackId).Execute()

List projects allowed to use a given private load zone.



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
	id := int32(56) // int32 | ID of the load zone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadZonesAPI.LoadZonesAllowedProjectsRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadZonesAPI.LoadZonesAllowedProjectsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadZonesAllowedProjectsRetrieve`: AllowedProjectsListApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadZonesAPI.LoadZonesAllowedProjectsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadZonesAllowedProjectsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**AllowedProjectsListApiModel**](AllowedProjectsListApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadZonesAllowedProjectsUpdate

> AllowedProjectsListApiModel LoadZonesAllowedProjectsUpdate(ctx, id).XStackId(xStackId).UpdateAllowedProjectsListApiModel(updateAllowedProjectsListApiModel).Execute()

Update the list of projects allowed to use a given private load zone.



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
	id := int32(56) // int32 | ID of the load zone.
	updateAllowedProjectsListApiModel := *openapiclient.NewUpdateAllowedProjectsListApiModel([]openapiclient.AllowedProjectToUpdateApiModel{*openapiclient.NewAllowedProjectToUpdateApiModel(int32(123))}) // UpdateAllowedProjectsListApiModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadZonesAPI.LoadZonesAllowedProjectsUpdate(context.Background(), id).XStackId(xStackId).UpdateAllowedProjectsListApiModel(updateAllowedProjectsListApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadZonesAPI.LoadZonesAllowedProjectsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadZonesAllowedProjectsUpdate`: AllowedProjectsListApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadZonesAPI.LoadZonesAllowedProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the load zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadZonesAllowedProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **updateAllowedProjectsListApiModel** | [**UpdateAllowedProjectsListApiModel**](UpdateAllowedProjectsListApiModel.md) |  | 

### Return type

[**AllowedProjectsListApiModel**](AllowedProjectsListApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadZonesList

> LoadZonesListApiModel LoadZonesList(ctx).XStackId(xStackId).K6LoadZoneId(k6LoadZoneId).Execute()

List all load zones.



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
	k6LoadZoneId := "k6LoadZoneId_example" // string | Filter results by k6 load zone ID (exact match). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadZonesAPI.LoadZonesList(context.Background()).XStackId(xStackId).K6LoadZoneId(k6LoadZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadZonesAPI.LoadZonesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadZonesList`: LoadZonesListApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadZonesAPI.LoadZonesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoadZonesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 
 **k6LoadZoneId** | **string** | Filter results by k6 load zone ID (exact match). | 

### Return type

[**LoadZonesListApiModel**](LoadZonesListApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsAllowedLoadZonesRetrieve

> AllowedLoadZonesListApiModel ProjectsAllowedLoadZonesRetrieve(ctx, id).XStackId(xStackId).Execute()

List private load zones that can be used by a given project.



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
	resp, r, err := apiClient.LoadZonesAPI.ProjectsAllowedLoadZonesRetrieve(context.Background(), id).XStackId(xStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadZonesAPI.ProjectsAllowedLoadZonesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsAllowedLoadZonesRetrieve`: AllowedLoadZonesListApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadZonesAPI.ProjectsAllowedLoadZonesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsAllowedLoadZonesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 


### Return type

[**AllowedLoadZonesListApiModel**](AllowedLoadZonesListApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsAllowedLoadZonesUpdate

> AllowedLoadZonesListApiModel ProjectsAllowedLoadZonesUpdate(ctx, id).XStackId(xStackId).UpdateAllowedLoadZonesListApiModel(updateAllowedLoadZonesListApiModel).Execute()

Update the list of private load zones that can be used by a given project.



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
	updateAllowedLoadZonesListApiModel := *openapiclient.NewUpdateAllowedLoadZonesListApiModel([]openapiclient.AllowedLoadZoneToUpdateApiModel{*openapiclient.NewAllowedLoadZoneToUpdateApiModel(int32(123))}) // UpdateAllowedLoadZonesListApiModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadZonesAPI.ProjectsAllowedLoadZonesUpdate(context.Background(), id).XStackId(xStackId).UpdateAllowedLoadZonesListApiModel(updateAllowedLoadZonesListApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadZonesAPI.ProjectsAllowedLoadZonesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsAllowedLoadZonesUpdate`: AllowedLoadZonesListApiModel
	fmt.Fprintf(os.Stdout, "Response from `LoadZonesAPI.ProjectsAllowedLoadZonesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsAllowedLoadZonesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **int32** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack. | 

 **updateAllowedLoadZonesListApiModel** | [**UpdateAllowedLoadZonesListApiModel**](UpdateAllowedLoadZonesListApiModel.md) |  | 

### Return type

[**AllowedLoadZonesListApiModel**](AllowedLoadZonesListApiModel.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

