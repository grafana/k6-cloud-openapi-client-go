# TestRunStatusExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByUser** | Pointer to **int32** | The user that set the status if applicable. | [optional] 
**Message** | Pointer to **string** | Human-readable string describing the error if applicable. | [optional] 
**Code** | Pointer to **string** | Service-defined error code if applicable. | [optional] 

## Methods

### NewTestRunStatusExtra

`func NewTestRunStatusExtra() *TestRunStatusExtra`

NewTestRunStatusExtra instantiates a new TestRunStatusExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunStatusExtraWithDefaults

`func NewTestRunStatusExtraWithDefaults() *TestRunStatusExtra`

NewTestRunStatusExtraWithDefaults instantiates a new TestRunStatusExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByUser

`func (o *TestRunStatusExtra) GetByUser() int32`

GetByUser returns the ByUser field if non-nil, zero value otherwise.

### GetByUserOk

`func (o *TestRunStatusExtra) GetByUserOk() (*int32, bool)`

GetByUserOk returns a tuple with the ByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByUser

`func (o *TestRunStatusExtra) SetByUser(v int32)`

SetByUser sets ByUser field to given value.

### HasByUser

`func (o *TestRunStatusExtra) HasByUser() bool`

HasByUser returns a boolean if a field has been set.

### GetMessage

`func (o *TestRunStatusExtra) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestRunStatusExtra) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestRunStatusExtra) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestRunStatusExtra) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *TestRunStatusExtra) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TestRunStatusExtra) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TestRunStatusExtra) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TestRunStatusExtra) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


