# \StaticIpsAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StaticIpsAcquireCreate**](StaticIpsAPI.md#StaticIpsAcquireCreate) | **Post** /static_ips/acquire | 
[**StaticIpsList**](StaticIpsAPI.md#StaticIpsList) | **Get** /static_ips | 
[**StaticIpsReleaseCreate**](StaticIpsAPI.md#StaticIpsReleaseCreate) | **Post** /static_ips/release | 



## StaticIpsAcquireCreate

> StaticIPList StaticIpsAcquireCreate(ctx).XStackId(xStackId).StaticIPAcquireRequest(staticIPAcquireRequest).Execute()





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
	staticIPAcquireRequest := *openapiclient.NewStaticIPAcquireRequest([]openapiclient.StaticIPAcquireLoadZoneRequest{*openapiclient.NewStaticIPAcquireLoadZoneRequest("LoadZone_example", int32(123))}) // StaticIPAcquireRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StaticIpsAPI.StaticIpsAcquireCreate(context.Background()).XStackId(xStackId).StaticIPAcquireRequest(staticIPAcquireRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticIpsAPI.StaticIpsAcquireCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StaticIpsAcquireCreate`: StaticIPList
	fmt.Fprintf(os.Stdout, "Response from `StaticIpsAPI.StaticIpsAcquireCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStaticIpsAcquireCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **staticIPAcquireRequest** | [**StaticIPAcquireRequest**](StaticIPAcquireRequest.md) |  | 

### Return type

[**StaticIPList**](StaticIPList.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StaticIpsList

> PaginatedStaticIPList StaticIpsList(ctx).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()





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
	orderby := "load_zone desc,provisioning_status" // string | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending `desc` specifier. Available fields: - load_zone - provisioning_status (optional) (default to "load_zone")
	skip := int32(56) // int32 | The initial index from which to return the results. (optional)
	top := int32(56) // int32 | Number of results to return per page. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StaticIpsAPI.StaticIpsList(context.Background()).XStackId(xStackId).Count(count).Orderby(orderby).Skip(skip).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticIpsAPI.StaticIpsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StaticIpsList`: PaginatedStaticIPList
	fmt.Fprintf(os.Stdout, "Response from `StaticIpsAPI.StaticIpsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStaticIpsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **count** | **bool** | Include collection length in the response object as &#39;@count&#39;. | 
 **orderby** | **string** | Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending &#x60;desc&#x60; specifier. Available fields: - load_zone - provisioning_status | [default to &quot;load_zone&quot;]
 **skip** | **int32** | The initial index from which to return the results. | 
 **top** | **int32** | Number of results to return per page. | [default to 1000]

### Return type

[**PaginatedStaticIPList**](PaginatedStaticIPList.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StaticIpsReleaseCreate

> StaticIPList StaticIpsReleaseCreate(ctx).XStackId(xStackId).StaticIPReleaseRequest(staticIPReleaseRequest).Execute()





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
	staticIPReleaseRequest := *openapiclient.NewStaticIPReleaseRequest([]int32{int32(123)}) // StaticIPReleaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StaticIpsAPI.StaticIpsReleaseCreate(context.Background()).XStackId(xStackId).StaticIPReleaseRequest(staticIPReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticIpsAPI.StaticIpsReleaseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StaticIpsReleaseCreate`: StaticIPList
	fmt.Fprintf(os.Stdout, "Response from `StaticIpsAPI.StaticIpsReleaseCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStaticIpsReleaseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackId** | **string** | Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.  | 
 **staticIPReleaseRequest** | [**StaticIPReleaseRequest**](StaticIPReleaseRequest.md) |  | 

### Return type

[**StaticIPList**](StaticIPList.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

