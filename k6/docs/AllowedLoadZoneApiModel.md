# AllowedLoadZoneApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the load zone. | 
**Name** | **string** | Name of the load zone. | 
**K6LoadZoneId** | **string** | ID used to identify the load zone in the k6 scripts. | 
**Available** | **bool** | Whether the load zone can be used to start tests. | 
**CustomLoadRunnerImage** | **NullableString** | Custom load runner image. Only set for private load zones. | 
**Public** | **bool** | Whether the load zone is public or private. | 

## Methods

### NewAllowedLoadZoneApiModel

`func NewAllowedLoadZoneApiModel(id int64, name string, k6LoadZoneId string, available bool, customLoadRunnerImage NullableString, public bool, ) *AllowedLoadZoneApiModel`

NewAllowedLoadZoneApiModel instantiates a new AllowedLoadZoneApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedLoadZoneApiModelWithDefaults

`func NewAllowedLoadZoneApiModelWithDefaults() *AllowedLoadZoneApiModel`

NewAllowedLoadZoneApiModelWithDefaults instantiates a new AllowedLoadZoneApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedLoadZoneApiModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedLoadZoneApiModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedLoadZoneApiModel) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *AllowedLoadZoneApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllowedLoadZoneApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllowedLoadZoneApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetK6LoadZoneId

`func (o *AllowedLoadZoneApiModel) GetK6LoadZoneId() string`

GetK6LoadZoneId returns the K6LoadZoneId field if non-nil, zero value otherwise.

### GetK6LoadZoneIdOk

`func (o *AllowedLoadZoneApiModel) GetK6LoadZoneIdOk() (*string, bool)`

GetK6LoadZoneIdOk returns a tuple with the K6LoadZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6LoadZoneId

`func (o *AllowedLoadZoneApiModel) SetK6LoadZoneId(v string)`

SetK6LoadZoneId sets K6LoadZoneId field to given value.


### GetAvailable

`func (o *AllowedLoadZoneApiModel) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AllowedLoadZoneApiModel) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AllowedLoadZoneApiModel) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCustomLoadRunnerImage

`func (o *AllowedLoadZoneApiModel) GetCustomLoadRunnerImage() string`

GetCustomLoadRunnerImage returns the CustomLoadRunnerImage field if non-nil, zero value otherwise.

### GetCustomLoadRunnerImageOk

`func (o *AllowedLoadZoneApiModel) GetCustomLoadRunnerImageOk() (*string, bool)`

GetCustomLoadRunnerImageOk returns a tuple with the CustomLoadRunnerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLoadRunnerImage

`func (o *AllowedLoadZoneApiModel) SetCustomLoadRunnerImage(v string)`

SetCustomLoadRunnerImage sets CustomLoadRunnerImage field to given value.


### SetCustomLoadRunnerImageNil

`func (o *AllowedLoadZoneApiModel) SetCustomLoadRunnerImageNil(b bool)`

 SetCustomLoadRunnerImageNil sets the value for CustomLoadRunnerImage to be an explicit nil

### UnsetCustomLoadRunnerImage
`func (o *AllowedLoadZoneApiModel) UnsetCustomLoadRunnerImage()`

UnsetCustomLoadRunnerImage ensures that no value is present for CustomLoadRunnerImage, not even an explicit nil
### GetPublic

`func (o *AllowedLoadZoneApiModel) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *AllowedLoadZoneApiModel) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *AllowedLoadZoneApiModel) SetPublic(v bool)`

SetPublic sets Public field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


