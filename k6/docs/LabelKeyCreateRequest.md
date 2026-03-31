# LabelKeyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**[]LabelKeyCreateItem**](LabelKeyCreateItem.md) | Array of label keys to create (1-100). | 

## Methods

### NewLabelKeyCreateRequest

`func NewLabelKeyCreateRequest(value []LabelKeyCreateItem, ) *LabelKeyCreateRequest`

NewLabelKeyCreateRequest instantiates a new LabelKeyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelKeyCreateRequestWithDefaults

`func NewLabelKeyCreateRequestWithDefaults() *LabelKeyCreateRequest`

NewLabelKeyCreateRequestWithDefaults instantiates a new LabelKeyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *LabelKeyCreateRequest) GetValue() []LabelKeyCreateItem`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LabelKeyCreateRequest) GetValueOk() (*[]LabelKeyCreateItem, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LabelKeyCreateRequest) SetValue(v []LabelKeyCreateItem)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


