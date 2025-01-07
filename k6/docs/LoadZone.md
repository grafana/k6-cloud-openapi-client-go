# LoadZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the load zone. | [readonly] 
**Name** | **NullableString** | Unique load zone name used in k6 script. | [readonly] 
**Title** | **string** | Human readable load zone name. | [readonly] 
**Private** | **bool** | If the load zone is a private load zone owned by the organization. | 
**ProviderName** | **string** | Name of the load zone provider. Usually, &#x60;amazon&#x60; for public load zones and &#x60;k8s&#x60; for the private ones. | [readonly] 
**ProviderDetails** | **map[string]interface{}** | Provider specific details. | 

## Methods

### NewLoadZone

`func NewLoadZone(id int32, name NullableString, title string, private bool, providerName string, providerDetails map[string]interface{}, ) *LoadZone`

NewLoadZone instantiates a new LoadZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadZoneWithDefaults

`func NewLoadZoneWithDefaults() *LoadZone`

NewLoadZoneWithDefaults instantiates a new LoadZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadZone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadZone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadZone) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *LoadZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadZone) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *LoadZone) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LoadZone) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *LoadZone) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LoadZone) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LoadZone) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrivate

`func (o *LoadZone) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *LoadZone) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *LoadZone) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetProviderName

`func (o *LoadZone) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *LoadZone) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *LoadZone) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetProviderDetails

`func (o *LoadZone) GetProviderDetails() map[string]interface{}`

GetProviderDetails returns the ProviderDetails field if non-nil, zero value otherwise.

### GetProviderDetailsOk

`func (o *LoadZone) GetProviderDetailsOk() (*map[string]interface{}, bool)`

GetProviderDetailsOk returns a tuple with the ProviderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDetails

`func (o *LoadZone) SetProviderDetails(v map[string]interface{})`

SetProviderDetails sets ProviderDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


