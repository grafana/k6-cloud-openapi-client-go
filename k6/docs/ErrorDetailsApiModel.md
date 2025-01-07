# ErrorDetailsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human-readable string describing the error. | 
**Code** | **string** | Service-defined error code. | 
**Target** | Pointer to **NullableString** | A string indicating the target of the error. For example, the name of the property in error. | [optional] 

## Methods

### NewErrorDetailsApiModel

`func NewErrorDetailsApiModel(message string, code string, ) *ErrorDetailsApiModel`

NewErrorDetailsApiModel instantiates a new ErrorDetailsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailsApiModelWithDefaults

`func NewErrorDetailsApiModelWithDefaults() *ErrorDetailsApiModel`

NewErrorDetailsApiModelWithDefaults instantiates a new ErrorDetailsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorDetailsApiModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetailsApiModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetailsApiModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ErrorDetailsApiModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDetailsApiModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDetailsApiModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetTarget

`func (o *ErrorDetailsApiModel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ErrorDetailsApiModel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ErrorDetailsApiModel) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ErrorDetailsApiModel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ErrorDetailsApiModel) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ErrorDetailsApiModel) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


