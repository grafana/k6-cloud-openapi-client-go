# LabelKeyPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | New label key name. | [optional] 
**Description** | Pointer to **NullableString** | New label key description. | [optional] 

## Methods

### NewLabelKeyPatchRequest

`func NewLabelKeyPatchRequest() *LabelKeyPatchRequest`

NewLabelKeyPatchRequest instantiates a new LabelKeyPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelKeyPatchRequestWithDefaults

`func NewLabelKeyPatchRequestWithDefaults() *LabelKeyPatchRequest`

NewLabelKeyPatchRequestWithDefaults instantiates a new LabelKeyPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LabelKeyPatchRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LabelKeyPatchRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LabelKeyPatchRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LabelKeyPatchRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *LabelKeyPatchRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *LabelKeyPatchRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetDescription

`func (o *LabelKeyPatchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LabelKeyPatchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LabelKeyPatchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LabelKeyPatchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LabelKeyPatchRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LabelKeyPatchRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


