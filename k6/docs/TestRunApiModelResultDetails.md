# TestRunApiModelResultDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result: &#x60;passed&#x60;, &#x60;failed&#x60; or &#x60;error&#x60;. | 
**Message** | Pointer to **string** | Human-readable description of why the test failed or the description of the error if applicable. | [optional] 
**Code** | Pointer to **int32** | Service-defined error code if applicable. | [optional] 

## Methods

### NewTestRunApiModelResultDetails

`func NewTestRunApiModelResultDetails(type_ string, ) *TestRunApiModelResultDetails`

NewTestRunApiModelResultDetails instantiates a new TestRunApiModelResultDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunApiModelResultDetailsWithDefaults

`func NewTestRunApiModelResultDetailsWithDefaults() *TestRunApiModelResultDetails`

NewTestRunApiModelResultDetailsWithDefaults instantiates a new TestRunApiModelResultDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TestRunApiModelResultDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestRunApiModelResultDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestRunApiModelResultDetails) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *TestRunApiModelResultDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestRunApiModelResultDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestRunApiModelResultDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestRunApiModelResultDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *TestRunApiModelResultDetails) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TestRunApiModelResultDetails) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TestRunApiModelResultDetails) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *TestRunApiModelResultDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


