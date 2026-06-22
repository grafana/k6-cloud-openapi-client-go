# ProjectLabelPutItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **NullableInt64** | Label key ID. Mutually exclusive with &#39;key&#39;. | [optional] 
**Key** | Pointer to **NullableString** | Label key name. Mutually exclusive with &#39;key_id&#39;. | [optional] 
**Value** | **string** | Label value. | 

## Methods

### NewProjectLabelPutItem

`func NewProjectLabelPutItem(value string, ) *ProjectLabelPutItem`

NewProjectLabelPutItem instantiates a new ProjectLabelPutItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectLabelPutItemWithDefaults

`func NewProjectLabelPutItemWithDefaults() *ProjectLabelPutItem`

NewProjectLabelPutItemWithDefaults instantiates a new ProjectLabelPutItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *ProjectLabelPutItem) GetKeyId() int64`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ProjectLabelPutItem) GetKeyIdOk() (*int64, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ProjectLabelPutItem) SetKeyId(v int64)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ProjectLabelPutItem) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### SetKeyIdNil

`func (o *ProjectLabelPutItem) SetKeyIdNil(b bool)`

 SetKeyIdNil sets the value for KeyId to be an explicit nil

### UnsetKeyId
`func (o *ProjectLabelPutItem) UnsetKeyId()`

UnsetKeyId ensures that no value is present for KeyId, not even an explicit nil
### GetKey

`func (o *ProjectLabelPutItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectLabelPutItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectLabelPutItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectLabelPutItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ProjectLabelPutItem) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProjectLabelPutItem) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *ProjectLabelPutItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectLabelPutItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectLabelPutItem) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


