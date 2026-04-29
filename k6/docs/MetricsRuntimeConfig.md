# MetricsRuntimeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushUrl** | **string** | URL to push metrics to. | 
**PushInterval** | **NullableString** | The time interval between periodic API calls for sending samples to the cloud ingest service. | 
**PushConcurrency** | **NullableInt32** | How many concurrent pushes will be done at the same time to the cloud. | 
**AggregationPeriod** | **NullableString** | If specified and is greater than 0, sample aggregation with that period is enabled. | 
**AggregationWaitPeriod** | **NullableString** | If aggregation is enabled, this specifies how long we&#39;ll wait for period samples to accumulate before trying to aggregate them. | 
**AggregationMinSamples** | **NullableInt32** | The minimum number of samples to aggregate. | 
**MaxSamplesPerPackage** | **NullableInt32** | The max allowed number of time series in a single batch. | 

## Methods

### NewMetricsRuntimeConfig

`func NewMetricsRuntimeConfig(pushUrl string, pushInterval NullableString, pushConcurrency NullableInt32, aggregationPeriod NullableString, aggregationWaitPeriod NullableString, aggregationMinSamples NullableInt32, maxSamplesPerPackage NullableInt32, ) *MetricsRuntimeConfig`

NewMetricsRuntimeConfig instantiates a new MetricsRuntimeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsRuntimeConfigWithDefaults

`func NewMetricsRuntimeConfigWithDefaults() *MetricsRuntimeConfig`

NewMetricsRuntimeConfigWithDefaults instantiates a new MetricsRuntimeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushUrl

`func (o *MetricsRuntimeConfig) GetPushUrl() string`

GetPushUrl returns the PushUrl field if non-nil, zero value otherwise.

### GetPushUrlOk

`func (o *MetricsRuntimeConfig) GetPushUrlOk() (*string, bool)`

GetPushUrlOk returns a tuple with the PushUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushUrl

`func (o *MetricsRuntimeConfig) SetPushUrl(v string)`

SetPushUrl sets PushUrl field to given value.


### GetPushInterval

`func (o *MetricsRuntimeConfig) GetPushInterval() string`

GetPushInterval returns the PushInterval field if non-nil, zero value otherwise.

### GetPushIntervalOk

`func (o *MetricsRuntimeConfig) GetPushIntervalOk() (*string, bool)`

GetPushIntervalOk returns a tuple with the PushInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushInterval

`func (o *MetricsRuntimeConfig) SetPushInterval(v string)`

SetPushInterval sets PushInterval field to given value.


### SetPushIntervalNil

`func (o *MetricsRuntimeConfig) SetPushIntervalNil(b bool)`

 SetPushIntervalNil sets the value for PushInterval to be an explicit nil

### UnsetPushInterval
`func (o *MetricsRuntimeConfig) UnsetPushInterval()`

UnsetPushInterval ensures that no value is present for PushInterval, not even an explicit nil
### GetPushConcurrency

`func (o *MetricsRuntimeConfig) GetPushConcurrency() int32`

GetPushConcurrency returns the PushConcurrency field if non-nil, zero value otherwise.

### GetPushConcurrencyOk

`func (o *MetricsRuntimeConfig) GetPushConcurrencyOk() (*int32, bool)`

GetPushConcurrencyOk returns a tuple with the PushConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushConcurrency

`func (o *MetricsRuntimeConfig) SetPushConcurrency(v int32)`

SetPushConcurrency sets PushConcurrency field to given value.


### SetPushConcurrencyNil

`func (o *MetricsRuntimeConfig) SetPushConcurrencyNil(b bool)`

 SetPushConcurrencyNil sets the value for PushConcurrency to be an explicit nil

### UnsetPushConcurrency
`func (o *MetricsRuntimeConfig) UnsetPushConcurrency()`

UnsetPushConcurrency ensures that no value is present for PushConcurrency, not even an explicit nil
### GetAggregationPeriod

`func (o *MetricsRuntimeConfig) GetAggregationPeriod() string`

GetAggregationPeriod returns the AggregationPeriod field if non-nil, zero value otherwise.

### GetAggregationPeriodOk

`func (o *MetricsRuntimeConfig) GetAggregationPeriodOk() (*string, bool)`

GetAggregationPeriodOk returns a tuple with the AggregationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPeriod

`func (o *MetricsRuntimeConfig) SetAggregationPeriod(v string)`

SetAggregationPeriod sets AggregationPeriod field to given value.


### SetAggregationPeriodNil

`func (o *MetricsRuntimeConfig) SetAggregationPeriodNil(b bool)`

 SetAggregationPeriodNil sets the value for AggregationPeriod to be an explicit nil

### UnsetAggregationPeriod
`func (o *MetricsRuntimeConfig) UnsetAggregationPeriod()`

UnsetAggregationPeriod ensures that no value is present for AggregationPeriod, not even an explicit nil
### GetAggregationWaitPeriod

`func (o *MetricsRuntimeConfig) GetAggregationWaitPeriod() string`

GetAggregationWaitPeriod returns the AggregationWaitPeriod field if non-nil, zero value otherwise.

### GetAggregationWaitPeriodOk

`func (o *MetricsRuntimeConfig) GetAggregationWaitPeriodOk() (*string, bool)`

GetAggregationWaitPeriodOk returns a tuple with the AggregationWaitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationWaitPeriod

`func (o *MetricsRuntimeConfig) SetAggregationWaitPeriod(v string)`

SetAggregationWaitPeriod sets AggregationWaitPeriod field to given value.


### SetAggregationWaitPeriodNil

`func (o *MetricsRuntimeConfig) SetAggregationWaitPeriodNil(b bool)`

 SetAggregationWaitPeriodNil sets the value for AggregationWaitPeriod to be an explicit nil

### UnsetAggregationWaitPeriod
`func (o *MetricsRuntimeConfig) UnsetAggregationWaitPeriod()`

UnsetAggregationWaitPeriod ensures that no value is present for AggregationWaitPeriod, not even an explicit nil
### GetAggregationMinSamples

`func (o *MetricsRuntimeConfig) GetAggregationMinSamples() int32`

GetAggregationMinSamples returns the AggregationMinSamples field if non-nil, zero value otherwise.

### GetAggregationMinSamplesOk

`func (o *MetricsRuntimeConfig) GetAggregationMinSamplesOk() (*int32, bool)`

GetAggregationMinSamplesOk returns a tuple with the AggregationMinSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMinSamples

`func (o *MetricsRuntimeConfig) SetAggregationMinSamples(v int32)`

SetAggregationMinSamples sets AggregationMinSamples field to given value.


### SetAggregationMinSamplesNil

`func (o *MetricsRuntimeConfig) SetAggregationMinSamplesNil(b bool)`

 SetAggregationMinSamplesNil sets the value for AggregationMinSamples to be an explicit nil

### UnsetAggregationMinSamples
`func (o *MetricsRuntimeConfig) UnsetAggregationMinSamples()`

UnsetAggregationMinSamples ensures that no value is present for AggregationMinSamples, not even an explicit nil
### GetMaxSamplesPerPackage

`func (o *MetricsRuntimeConfig) GetMaxSamplesPerPackage() int32`

GetMaxSamplesPerPackage returns the MaxSamplesPerPackage field if non-nil, zero value otherwise.

### GetMaxSamplesPerPackageOk

`func (o *MetricsRuntimeConfig) GetMaxSamplesPerPackageOk() (*int32, bool)`

GetMaxSamplesPerPackageOk returns a tuple with the MaxSamplesPerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSamplesPerPackage

`func (o *MetricsRuntimeConfig) SetMaxSamplesPerPackage(v int32)`

SetMaxSamplesPerPackage sets MaxSamplesPerPackage field to given value.


### SetMaxSamplesPerPackageNil

`func (o *MetricsRuntimeConfig) SetMaxSamplesPerPackageNil(b bool)`

 SetMaxSamplesPerPackageNil sets the value for MaxSamplesPerPackage to be an explicit nil

### UnsetMaxSamplesPerPackage
`func (o *MetricsRuntimeConfig) UnsetMaxSamplesPerPackage()`

UnsetMaxSamplesPerPackage ensures that no value is present for MaxSamplesPerPackage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


