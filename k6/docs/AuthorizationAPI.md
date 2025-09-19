# \AuthorizationAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Auth**](AuthorizationAPI.md#Auth) | **Get** /cloud/v6/auth | Validate k6 Cloud API token and access to the stack.



## Auth

> AuthenticationResponse Auth(ctx).XStackUrl(xStackUrl).Execute()

Validate k6 Cloud API token and access to the stack.



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
	xStackUrl := "xStackUrl_example" // string | The URL of the Grafana stack to authenticate to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.Auth(context.Background()).XStackUrl(xStackUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.Auth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Auth`: AuthenticationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.Auth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStackUrl** | **string** | The URL of the Grafana stack to authenticate to. | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

[k6ApiToken](../README.md#k6ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

