# ScriptExecutionCompletedNotificationApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Event type. | 
**Error** | Pointer to [**NullableScriptExecutionError**](ScriptExecutionError.md) |  | [optional] 

## Methods

### NewScriptExecutionCompletedNotificationApiModel

`func NewScriptExecutionCompletedNotificationApiModel(eventType string, ) *ScriptExecutionCompletedNotificationApiModel`

NewScriptExecutionCompletedNotificationApiModel instantiates a new ScriptExecutionCompletedNotificationApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptExecutionCompletedNotificationApiModelWithDefaults

`func NewScriptExecutionCompletedNotificationApiModelWithDefaults() *ScriptExecutionCompletedNotificationApiModel`

NewScriptExecutionCompletedNotificationApiModelWithDefaults instantiates a new ScriptExecutionCompletedNotificationApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ScriptExecutionCompletedNotificationApiModel) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ScriptExecutionCompletedNotificationApiModel) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ScriptExecutionCompletedNotificationApiModel) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetError

`func (o *ScriptExecutionCompletedNotificationApiModel) GetError() ScriptExecutionError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScriptExecutionCompletedNotificationApiModel) GetErrorOk() (*ScriptExecutionError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScriptExecutionCompletedNotificationApiModel) SetError(v ScriptExecutionError)`

SetError sets Error field to given value.

### HasError

`func (o *ScriptExecutionCompletedNotificationApiModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ScriptExecutionCompletedNotificationApiModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ScriptExecutionCompletedNotificationApiModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


