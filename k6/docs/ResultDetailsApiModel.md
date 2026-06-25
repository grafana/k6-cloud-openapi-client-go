# ResultDetailsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result: &#x60;passed&#x60;, &#x60;failed&#x60; or &#x60;error&#x60;. | 
**Message** | Pointer to **NullableString** | Human-readable description of why the test failed or the description of the error if applicable. | [optional] 
**Code** | Pointer to **NullableInt32** | Service-defined error code if applicable. | [optional] 

## Methods

### NewResultDetailsApiModel

`func NewResultDetailsApiModel(type_ string, ) *ResultDetailsApiModel`

NewResultDetailsApiModel instantiates a new ResultDetailsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultDetailsApiModelWithDefaults

`func NewResultDetailsApiModelWithDefaults() *ResultDetailsApiModel`

NewResultDetailsApiModelWithDefaults instantiates a new ResultDetailsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResultDetailsApiModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResultDetailsApiModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResultDetailsApiModel) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *ResultDetailsApiModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResultDetailsApiModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResultDetailsApiModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResultDetailsApiModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ResultDetailsApiModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ResultDetailsApiModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCode

`func (o *ResultDetailsApiModel) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResultDetailsApiModel) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResultDetailsApiModel) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResultDetailsApiModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ResultDetailsApiModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ResultDetailsApiModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


