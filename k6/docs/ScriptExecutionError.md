# ScriptExecutionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Error code. | 
**Reason** | **string** | Human-readable error description shown to the user. | 
**Detail** | Pointer to **NullableString** | Optional internal debug information. | [optional] 

## Methods

### NewScriptExecutionError

`func NewScriptExecutionError(code int32, reason string, ) *ScriptExecutionError`

NewScriptExecutionError instantiates a new ScriptExecutionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptExecutionErrorWithDefaults

`func NewScriptExecutionErrorWithDefaults() *ScriptExecutionError`

NewScriptExecutionErrorWithDefaults instantiates a new ScriptExecutionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ScriptExecutionError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ScriptExecutionError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ScriptExecutionError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetReason

`func (o *ScriptExecutionError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ScriptExecutionError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ScriptExecutionError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDetail

`func (o *ScriptExecutionError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ScriptExecutionError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ScriptExecutionError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ScriptExecutionError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ScriptExecutionError) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ScriptExecutionError) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


