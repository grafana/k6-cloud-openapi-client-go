# ProjectLabelPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**[]ProjectLabelPutItem**](ProjectLabelPutItem.md) | Array of labels to assign (0-30). | 

## Methods

### NewProjectLabelPutRequest

`func NewProjectLabelPutRequest(value []ProjectLabelPutItem, ) *ProjectLabelPutRequest`

NewProjectLabelPutRequest instantiates a new ProjectLabelPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectLabelPutRequestWithDefaults

`func NewProjectLabelPutRequestWithDefaults() *ProjectLabelPutRequest`

NewProjectLabelPutRequestWithDefaults instantiates a new ProjectLabelPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ProjectLabelPutRequest) GetValue() []ProjectLabelPutItem`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectLabelPutRequest) GetValueOk() (*[]ProjectLabelPutItem, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectLabelPutRequest) SetValue(v []ProjectLabelPutItem)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


