# LogsRuntimeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushUrl** | **string** | URL to push logs to. | 
**TailUrl** | **string** | URL to tail logs from. | 
**Limit** | **int32** | Maximum number of log lines to push per interval. | 
**PushPeriodSeconds** | **string** | How often logs are pushed to the ingest service, for example &#39;3s&#39;. | 
**Level** | **string** | Minimum log level to collect. | 
**MessageMaxSize** | **int32** | Maximum size in bytes of a single log message. | 
**AllowedLabels** | **[]string** | Log stream labels that are allowed to be forwarded. | 

## Methods

### NewLogsRuntimeConfig

`func NewLogsRuntimeConfig(pushUrl string, tailUrl string, limit int32, pushPeriodSeconds string, level string, messageMaxSize int32, allowedLabels []string, ) *LogsRuntimeConfig`

NewLogsRuntimeConfig instantiates a new LogsRuntimeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsRuntimeConfigWithDefaults

`func NewLogsRuntimeConfigWithDefaults() *LogsRuntimeConfig`

NewLogsRuntimeConfigWithDefaults instantiates a new LogsRuntimeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushUrl

`func (o *LogsRuntimeConfig) GetPushUrl() string`

GetPushUrl returns the PushUrl field if non-nil, zero value otherwise.

### GetPushUrlOk

`func (o *LogsRuntimeConfig) GetPushUrlOk() (*string, bool)`

GetPushUrlOk returns a tuple with the PushUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushUrl

`func (o *LogsRuntimeConfig) SetPushUrl(v string)`

SetPushUrl sets PushUrl field to given value.


### GetTailUrl

`func (o *LogsRuntimeConfig) GetTailUrl() string`

GetTailUrl returns the TailUrl field if non-nil, zero value otherwise.

### GetTailUrlOk

`func (o *LogsRuntimeConfig) GetTailUrlOk() (*string, bool)`

GetTailUrlOk returns a tuple with the TailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTailUrl

`func (o *LogsRuntimeConfig) SetTailUrl(v string)`

SetTailUrl sets TailUrl field to given value.


### GetLimit

`func (o *LogsRuntimeConfig) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogsRuntimeConfig) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LogsRuntimeConfig) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPushPeriodSeconds

`func (o *LogsRuntimeConfig) GetPushPeriodSeconds() string`

GetPushPeriodSeconds returns the PushPeriodSeconds field if non-nil, zero value otherwise.

### GetPushPeriodSecondsOk

`func (o *LogsRuntimeConfig) GetPushPeriodSecondsOk() (*string, bool)`

GetPushPeriodSecondsOk returns a tuple with the PushPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushPeriodSeconds

`func (o *LogsRuntimeConfig) SetPushPeriodSeconds(v string)`

SetPushPeriodSeconds sets PushPeriodSeconds field to given value.


### GetLevel

`func (o *LogsRuntimeConfig) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogsRuntimeConfig) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogsRuntimeConfig) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMessageMaxSize

`func (o *LogsRuntimeConfig) GetMessageMaxSize() int32`

GetMessageMaxSize returns the MessageMaxSize field if non-nil, zero value otherwise.

### GetMessageMaxSizeOk

`func (o *LogsRuntimeConfig) GetMessageMaxSizeOk() (*int32, bool)`

GetMessageMaxSizeOk returns a tuple with the MessageMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageMaxSize

`func (o *LogsRuntimeConfig) SetMessageMaxSize(v int32)`

SetMessageMaxSize sets MessageMaxSize field to given value.


### GetAllowedLabels

`func (o *LogsRuntimeConfig) GetAllowedLabels() []string`

GetAllowedLabels returns the AllowedLabels field if non-nil, zero value otherwise.

### GetAllowedLabelsOk

`func (o *LogsRuntimeConfig) GetAllowedLabelsOk() (*[]string, bool)`

GetAllowedLabelsOk returns a tuple with the AllowedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLabels

`func (o *LogsRuntimeConfig) SetAllowedLabels(v []string)`

SetAllowedLabels sets AllowedLabels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


