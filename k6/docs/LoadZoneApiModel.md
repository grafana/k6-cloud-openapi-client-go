# LoadZoneApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the load zone. | 
**Name** | **string** | Name of the load zone. | 
**K6LoadZoneId** | **string** | ID used to identify the load zone in the k6 scripts. | 

## Methods

### NewLoadZoneApiModel

`func NewLoadZoneApiModel(id int32, name string, k6LoadZoneId string, ) *LoadZoneApiModel`

NewLoadZoneApiModel instantiates a new LoadZoneApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadZoneApiModelWithDefaults

`func NewLoadZoneApiModelWithDefaults() *LoadZoneApiModel`

NewLoadZoneApiModelWithDefaults instantiates a new LoadZoneApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadZoneApiModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadZoneApiModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadZoneApiModel) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *LoadZoneApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadZoneApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadZoneApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetK6LoadZoneId

`func (o *LoadZoneApiModel) GetK6LoadZoneId() string`

GetK6LoadZoneId returns the K6LoadZoneId field if non-nil, zero value otherwise.

### GetK6LoadZoneIdOk

`func (o *LoadZoneApiModel) GetK6LoadZoneIdOk() (*string, bool)`

GetK6LoadZoneIdOk returns a tuple with the K6LoadZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6LoadZoneId

`func (o *LoadZoneApiModel) SetK6LoadZoneId(v string)`

SetK6LoadZoneId sets K6LoadZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


