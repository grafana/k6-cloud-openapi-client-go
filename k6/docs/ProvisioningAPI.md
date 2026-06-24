# \ProvisioningAPI

All URIs are relative to *https://api.k6.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StartLocalExecutionTest**](ProvisioningAPI.md#StartLocalExecutionTest) | **Post** /provisioning/v1/load_tests/{id}/start_local_execution | Start a test in local execution mode.
[**TestRunsAisdkSessionRetrieve**](ProvisioningAPI.md#TestRunsAisdkSessionRetrieve) | **Get** /provisioning/v1/test_runs/{test_run_id}/aisdk_session | Mint an ai-sdk access token for a test run.
[**TestRunsDecryptSecretRetrieve**](ProvisioningAPI.md#TestRunsDecryptSecretRetrieve) | **Get** /provisioning/v1/test_runs/{test_run_id}/decrypt_secret | Decrypt a secret for a test run.
[**TestRunsNotify**](ProvisioningAPI.md#TestRunsNotify) | **Post** /provisioning/v1/test_runs/{id}/notify | Report a test run lifecycle event.



## StartLocalExecutionTest

> StartLocalExecutionTestResponse StartLocalExecutionTest(ctx, id).StartLocalExecutionTestRequest(startLocalExecutionTestRequest).K6IdempotencyKey(k6IdempotencyKey).Execute()

Start a test in local execution mode.



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
	id := int64(789) // int64 | ID of the load test.
	startLocalExecutionTestRequest := *openapiclient.NewStartLocalExecutionTestRequest(map[string]interface{}{"key": interface{}(123)}, int32(123), int32(123)) // StartLocalExecutionTestRequest | 
	k6IdempotencyKey := "k6IdempotencyKey_example" // string | Idempotency key to prevent duplicate test starts when retrying requests. The key is valid for 10 minutes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.StartLocalExecutionTest(context.Background(), id).StartLocalExecutionTestRequest(startLocalExecutionTestRequest).K6IdempotencyKey(k6IdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.StartLocalExecutionTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartLocalExecutionTest`: StartLocalExecutionTestResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.StartLocalExecutionTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the load test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartLocalExecutionTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startLocalExecutionTestRequest** | [**StartLocalExecutionTestRequest**](StartLocalExecutionTestRequest.md) |  | 
 **k6IdempotencyKey** | **string** | Idempotency key to prevent duplicate test starts when retrying requests. The key is valid for 10 minutes. | 

### Return type

[**StartLocalExecutionTestResponse**](StartLocalExecutionTestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsAisdkSessionRetrieve

> AisdkSessionApiModel TestRunsAisdkSessionRetrieve(ctx, testRunId).Execute()

Mint an ai-sdk access token for a test run.



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
	testRunId := int64(789) // int64 | ID of the test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.TestRunsAisdkSessionRetrieve(context.Background(), testRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.TestRunsAisdkSessionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsAisdkSessionRetrieve`: AisdkSessionApiModel
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.TestRunsAisdkSessionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testRunId** | **int64** | ID of the test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsAisdkSessionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AisdkSessionApiModel**](AisdkSessionApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsDecryptSecretRetrieve

> DecryptedApiModel TestRunsDecryptSecretRetrieve(ctx, testRunId).Name(name).Execute()

Decrypt a secret for a test run.



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
	name := "name_example" // string | The name of the secret to decrypt.
	testRunId := int64(789) // int64 | ID of the test run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.TestRunsDecryptSecretRetrieve(context.Background(), testRunId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.TestRunsDecryptSecretRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunsDecryptSecretRetrieve`: DecryptedApiModel
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.TestRunsDecryptSecretRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testRunId** | **int64** | ID of the test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsDecryptSecretRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the secret to decrypt. | 


### Return type

[**DecryptedApiModel**](DecryptedApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunsNotify

> TestRunsNotify(ctx, id).ScriptExecutionCompletedNotificationApiModel(scriptExecutionCompletedNotificationApiModel).Execute()

Report a test run lifecycle event.



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
	id := int64(789) // int64 | ID of the test run.
	scriptExecutionCompletedNotificationApiModel := *openapiclient.NewScriptExecutionCompletedNotificationApiModel("EventType_example") // ScriptExecutionCompletedNotificationApiModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.TestRunsNotify(context.Background(), id).ScriptExecutionCompletedNotificationApiModel(scriptExecutionCompletedNotificationApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.TestRunsNotify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the test run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunsNotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scriptExecutionCompletedNotificationApiModel** | [**ScriptExecutionCompletedNotificationApiModel**](ScriptExecutionCompletedNotificationApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

