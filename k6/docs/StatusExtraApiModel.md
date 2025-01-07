# StatusExtraApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByUser** | **NullableString** | Email of the user that set the status if applicable. | 
**Message** | **NullableString** | Human-readable string describing the error if applicable. | 
**Code** | **NullableInt32** | Service-defined error code if applicable. | 

## Methods

### NewStatusExtraApiModel

`func NewStatusExtraApiModel(byUser NullableString, message NullableString, code NullableInt32, ) *StatusExtraApiModel`

NewStatusExtraApiModel instantiates a new StatusExtraApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusExtraApiModelWithDefaults

`func NewStatusExtraApiModelWithDefaults() *StatusExtraApiModel`

NewStatusExtraApiModelWithDefaults instantiates a new StatusExtraApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByUser

`func (o *StatusExtraApiModel) GetByUser() string`

GetByUser returns the ByUser field if non-nil, zero value otherwise.

### GetByUserOk

`func (o *StatusExtraApiModel) GetByUserOk() (*string, bool)`

GetByUserOk returns a tuple with the ByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByUser

`func (o *StatusExtraApiModel) SetByUser(v string)`

SetByUser sets ByUser field to given value.


### SetByUserNil

`func (o *StatusExtraApiModel) SetByUserNil(b bool)`

 SetByUserNil sets the value for ByUser to be an explicit nil

### UnsetByUser
`func (o *StatusExtraApiModel) UnsetByUser()`

UnsetByUser ensures that no value is present for ByUser, not even an explicit nil
### GetMessage

`func (o *StatusExtraApiModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusExtraApiModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusExtraApiModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *StatusExtraApiModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *StatusExtraApiModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCode

`func (o *StatusExtraApiModel) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StatusExtraApiModel) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StatusExtraApiModel) SetCode(v int32)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *StatusExtraApiModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *StatusExtraApiModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


