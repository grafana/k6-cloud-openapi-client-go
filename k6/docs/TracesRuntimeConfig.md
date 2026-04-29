# TracesRuntimeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushUrl** | **string** | URL to push traces to. | 
**TracesMetadata** | **map[string]string** | Metadata labels attached to all traces. | 
**Proto** | **string** | Protocol to use for pushing traces. | 

## Methods

### NewTracesRuntimeConfig

`func NewTracesRuntimeConfig(pushUrl string, tracesMetadata map[string]string, proto string, ) *TracesRuntimeConfig`

NewTracesRuntimeConfig instantiates a new TracesRuntimeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracesRuntimeConfigWithDefaults

`func NewTracesRuntimeConfigWithDefaults() *TracesRuntimeConfig`

NewTracesRuntimeConfigWithDefaults instantiates a new TracesRuntimeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushUrl

`func (o *TracesRuntimeConfig) GetPushUrl() string`

GetPushUrl returns the PushUrl field if non-nil, zero value otherwise.

### GetPushUrlOk

`func (o *TracesRuntimeConfig) GetPushUrlOk() (*string, bool)`

GetPushUrlOk returns a tuple with the PushUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushUrl

`func (o *TracesRuntimeConfig) SetPushUrl(v string)`

SetPushUrl sets PushUrl field to given value.


### GetTracesMetadata

`func (o *TracesRuntimeConfig) GetTracesMetadata() map[string]string`

GetTracesMetadata returns the TracesMetadata field if non-nil, zero value otherwise.

### GetTracesMetadataOk

`func (o *TracesRuntimeConfig) GetTracesMetadataOk() (*map[string]string, bool)`

GetTracesMetadataOk returns a tuple with the TracesMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracesMetadata

`func (o *TracesRuntimeConfig) SetTracesMetadata(v map[string]string)`

SetTracesMetadata sets TracesMetadata field to given value.


### GetProto

`func (o *TracesRuntimeConfig) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *TracesRuntimeConfig) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *TracesRuntimeConfig) SetProto(v string)`

SetProto sets Proto field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


