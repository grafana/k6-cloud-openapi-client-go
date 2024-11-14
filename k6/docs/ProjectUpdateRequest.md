# ProjectUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Project name. | [optional] 
**Description** | Pointer to **string** | Project description. | [optional] 
**Settings** | Pointer to [**ProjectSettings**](ProjectSettings.md) | Project specific settings. | [optional] 

## Methods

### NewProjectUpdateRequest

`func NewProjectUpdateRequest() *ProjectUpdateRequest`

NewProjectUpdateRequest instantiates a new ProjectUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateRequestWithDefaults

`func NewProjectUpdateRequestWithDefaults() *ProjectUpdateRequest`

NewProjectUpdateRequestWithDefaults instantiates a new ProjectUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSettings

`func (o *ProjectUpdateRequest) GetSettings() ProjectSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProjectUpdateRequest) GetSettingsOk() (*ProjectSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProjectUpdateRequest) SetSettings(v ProjectSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ProjectUpdateRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


