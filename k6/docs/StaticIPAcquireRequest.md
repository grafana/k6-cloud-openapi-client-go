# StaticIPAcquireRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadZones** | [**[]StaticIPAcquireLoadZoneRequest**](StaticIPAcquireLoadZoneRequest.md) | List of the load zones with the number of IPs to acquire. | 

## Methods

### NewStaticIPAcquireRequest

`func NewStaticIPAcquireRequest(loadZones []StaticIPAcquireLoadZoneRequest, ) *StaticIPAcquireRequest`

NewStaticIPAcquireRequest instantiates a new StaticIPAcquireRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticIPAcquireRequestWithDefaults

`func NewStaticIPAcquireRequestWithDefaults() *StaticIPAcquireRequest`

NewStaticIPAcquireRequestWithDefaults instantiates a new StaticIPAcquireRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadZones

`func (o *StaticIPAcquireRequest) GetLoadZones() []StaticIPAcquireLoadZoneRequest`

GetLoadZones returns the LoadZones field if non-nil, zero value otherwise.

### GetLoadZonesOk

`func (o *StaticIPAcquireRequest) GetLoadZonesOk() (*[]StaticIPAcquireLoadZoneRequest, bool)`

GetLoadZonesOk returns a tuple with the LoadZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadZones

`func (o *StaticIPAcquireRequest) SetLoadZones(v []StaticIPAcquireLoadZoneRequest)`

SetLoadZones sets LoadZones field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


