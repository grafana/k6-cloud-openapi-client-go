# PatchedStackEnvVarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the variable. | [optional] 
**Value** | Pointer to **string** | Variable value. | [optional] 
**Description** | Pointer to **NullableString** | Description of the variable. | [optional] 

## Methods

### NewPatchedStackEnvVarRequest

`func NewPatchedStackEnvVarRequest() *PatchedStackEnvVarRequest`

NewPatchedStackEnvVarRequest instantiates a new PatchedStackEnvVarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedStackEnvVarRequestWithDefaults

`func NewPatchedStackEnvVarRequestWithDefaults() *PatchedStackEnvVarRequest`

NewPatchedStackEnvVarRequestWithDefaults instantiates a new PatchedStackEnvVarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedStackEnvVarRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedStackEnvVarRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedStackEnvVarRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedStackEnvVarRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *PatchedStackEnvVarRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedStackEnvVarRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedStackEnvVarRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedStackEnvVarRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedStackEnvVarRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedStackEnvVarRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedStackEnvVarRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedStackEnvVarRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedStackEnvVarRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedStackEnvVarRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


