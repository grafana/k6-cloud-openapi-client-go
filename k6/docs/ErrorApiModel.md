# ErrorApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human-readable string describing the error. | 
**Code** | **string** | Service-defined error code. | 
**Target** | Pointer to **NullableString** | A string indicating the target of the error. For example, the name of the property in error. | [optional] 
**Details** | Pointer to [**[]ErrorDetailsApiModel**](ErrorDetailsApiModel.md) | Array of objects with more specific error information when applicable. | [optional] 

## Methods

### NewErrorApiModel

`func NewErrorApiModel(message string, code string, ) *ErrorApiModel`

NewErrorApiModel instantiates a new ErrorApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorApiModelWithDefaults

`func NewErrorApiModelWithDefaults() *ErrorApiModel`

NewErrorApiModelWithDefaults instantiates a new ErrorApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorApiModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorApiModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorApiModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ErrorApiModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorApiModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorApiModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetTarget

`func (o *ErrorApiModel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ErrorApiModel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ErrorApiModel) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ErrorApiModel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ErrorApiModel) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ErrorApiModel) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetDetails

`func (o *ErrorApiModel) GetDetails() []ErrorDetailsApiModel`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorApiModel) GetDetailsOk() (*[]ErrorDetailsApiModel, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorApiModel) SetDetails(v []ErrorDetailsApiModel)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorApiModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ErrorApiModel) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ErrorApiModel) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


