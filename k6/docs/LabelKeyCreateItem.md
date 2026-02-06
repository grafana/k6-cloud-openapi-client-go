# LabelKeyCreateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Label key name. Case-sensitive. Unique per organization. | 
**Description** | Pointer to **NullableString** | Optional description of the label key. | [optional] 

## Methods

### NewLabelKeyCreateItem

`func NewLabelKeyCreateItem(key string, ) *LabelKeyCreateItem`

NewLabelKeyCreateItem instantiates a new LabelKeyCreateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelKeyCreateItemWithDefaults

`func NewLabelKeyCreateItemWithDefaults() *LabelKeyCreateItem`

NewLabelKeyCreateItemWithDefaults instantiates a new LabelKeyCreateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LabelKeyCreateItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LabelKeyCreateItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LabelKeyCreateItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *LabelKeyCreateItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LabelKeyCreateItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LabelKeyCreateItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LabelKeyCreateItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LabelKeyCreateItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LabelKeyCreateItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


