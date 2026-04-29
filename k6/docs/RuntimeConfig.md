# RuntimeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**MetricsRuntimeConfig**](MetricsRuntimeConfig.md) |  | 
**Traces** | [**TracesRuntimeConfig**](TracesRuntimeConfig.md) |  | 
**Files** | [**FilesRuntimeConfig**](FilesRuntimeConfig.md) |  | 
**Logs** | [**LogsRuntimeConfig**](LogsRuntimeConfig.md) |  | 
**Secrets** | [**SecretsRuntimeConfig**](SecretsRuntimeConfig.md) |  | 
**TestRunToken** | **string** | The short-lived test run token. | 

## Methods

### NewRuntimeConfig

`func NewRuntimeConfig(metrics MetricsRuntimeConfig, traces TracesRuntimeConfig, files FilesRuntimeConfig, logs LogsRuntimeConfig, secrets SecretsRuntimeConfig, testRunToken string, ) *RuntimeConfig`

NewRuntimeConfig instantiates a new RuntimeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeConfigWithDefaults

`func NewRuntimeConfigWithDefaults() *RuntimeConfig`

NewRuntimeConfigWithDefaults instantiates a new RuntimeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *RuntimeConfig) GetMetrics() MetricsRuntimeConfig`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *RuntimeConfig) GetMetricsOk() (*MetricsRuntimeConfig, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *RuntimeConfig) SetMetrics(v MetricsRuntimeConfig)`

SetMetrics sets Metrics field to given value.


### GetTraces

`func (o *RuntimeConfig) GetTraces() TracesRuntimeConfig`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *RuntimeConfig) GetTracesOk() (*TracesRuntimeConfig, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *RuntimeConfig) SetTraces(v TracesRuntimeConfig)`

SetTraces sets Traces field to given value.


### GetFiles

`func (o *RuntimeConfig) GetFiles() FilesRuntimeConfig`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RuntimeConfig) GetFilesOk() (*FilesRuntimeConfig, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RuntimeConfig) SetFiles(v FilesRuntimeConfig)`

SetFiles sets Files field to given value.


### GetLogs

`func (o *RuntimeConfig) GetLogs() LogsRuntimeConfig`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *RuntimeConfig) GetLogsOk() (*LogsRuntimeConfig, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *RuntimeConfig) SetLogs(v LogsRuntimeConfig)`

SetLogs sets Logs field to given value.


### GetSecrets

`func (o *RuntimeConfig) GetSecrets() SecretsRuntimeConfig`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *RuntimeConfig) GetSecretsOk() (*SecretsRuntimeConfig, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *RuntimeConfig) SetSecrets(v SecretsRuntimeConfig)`

SetSecrets sets Secrets field to given value.


### GetTestRunToken

`func (o *RuntimeConfig) GetTestRunToken() string`

GetTestRunToken returns the TestRunToken field if non-nil, zero value otherwise.

### GetTestRunTokenOk

`func (o *RuntimeConfig) GetTestRunTokenOk() (*string, bool)`

GetTestRunTokenOk returns a tuple with the TestRunToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunToken

`func (o *RuntimeConfig) SetTestRunToken(v string)`

SetTestRunToken sets TestRunToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


