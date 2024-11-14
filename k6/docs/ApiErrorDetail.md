# ApiErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human-readable string describing the error | 
**Code** | **string** | Service-defined error code | 
**Target** | Pointer to **string** | A string indicating the target of the error, for example, the name of the property in error | [optional] 

## Methods

### NewApiErrorDetail

`func NewApiErrorDetail(message string, code string, ) *ApiErrorDetail`

NewApiErrorDetail instantiates a new ApiErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorDetailWithDefaults

`func NewApiErrorDetailWithDefaults() *ApiErrorDetail`

NewApiErrorDetailWithDefaults instantiates a new ApiErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ApiErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.


### GetTarget

`func (o *ApiErrorDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ApiErrorDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ApiErrorDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ApiErrorDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


