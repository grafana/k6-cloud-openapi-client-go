# StaticIPAcquireLoadZoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadZone** | **string** | Name of the load zone. | 
**Count** | **int32** | Number of Static IPs to acquire. | 

## Methods

### NewStaticIPAcquireLoadZoneRequest

`func NewStaticIPAcquireLoadZoneRequest(loadZone string, count int32, ) *StaticIPAcquireLoadZoneRequest`

NewStaticIPAcquireLoadZoneRequest instantiates a new StaticIPAcquireLoadZoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticIPAcquireLoadZoneRequestWithDefaults

`func NewStaticIPAcquireLoadZoneRequestWithDefaults() *StaticIPAcquireLoadZoneRequest`

NewStaticIPAcquireLoadZoneRequestWithDefaults instantiates a new StaticIPAcquireLoadZoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadZone

`func (o *StaticIPAcquireLoadZoneRequest) GetLoadZone() string`

GetLoadZone returns the LoadZone field if non-nil, zero value otherwise.

### GetLoadZoneOk

`func (o *StaticIPAcquireLoadZoneRequest) GetLoadZoneOk() (*string, bool)`

GetLoadZoneOk returns a tuple with the LoadZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadZone

`func (o *StaticIPAcquireLoadZoneRequest) SetLoadZone(v string)`

SetLoadZone sets LoadZone field to given value.


### GetCount

`func (o *StaticIPAcquireLoadZoneRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StaticIPAcquireLoadZoneRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StaticIPAcquireLoadZoneRequest) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


