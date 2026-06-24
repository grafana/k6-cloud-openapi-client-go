# ValidateOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **NullableInt64** | ID of a project where the test belongs. | [optional] 
**Options** | [**Options**](Options.md) | k6 script options object to validate. | 
**K6Dependencies** | Pointer to **map[string]string** | Version of k6 and extensions to validate, as a map of dependency name to dependency version constraint. | [optional] 
**K6Version** | Pointer to **NullableInt32** | Identifier of the k6 version used to run the test. | [optional] 
**IsLocalExecution** | Pointer to **bool** | Whether the test is being executed locally. | [optional] [default to false]

## Methods

### NewValidateOptionsRequest

`func NewValidateOptionsRequest(options Options, ) *ValidateOptionsRequest`

NewValidateOptionsRequest instantiates a new ValidateOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateOptionsRequestWithDefaults

`func NewValidateOptionsRequestWithDefaults() *ValidateOptionsRequest`

NewValidateOptionsRequestWithDefaults instantiates a new ValidateOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ValidateOptionsRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ValidateOptionsRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ValidateOptionsRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ValidateOptionsRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ValidateOptionsRequest) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ValidateOptionsRequest) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOptions

`func (o *ValidateOptionsRequest) GetOptions() Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ValidateOptionsRequest) GetOptionsOk() (*Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ValidateOptionsRequest) SetOptions(v Options)`

SetOptions sets Options field to given value.


### GetK6Dependencies

`func (o *ValidateOptionsRequest) GetK6Dependencies() map[string]string`

GetK6Dependencies returns the K6Dependencies field if non-nil, zero value otherwise.

### GetK6DependenciesOk

`func (o *ValidateOptionsRequest) GetK6DependenciesOk() (*map[string]string, bool)`

GetK6DependenciesOk returns a tuple with the K6Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Dependencies

`func (o *ValidateOptionsRequest) SetK6Dependencies(v map[string]string)`

SetK6Dependencies sets K6Dependencies field to given value.

### HasK6Dependencies

`func (o *ValidateOptionsRequest) HasK6Dependencies() bool`

HasK6Dependencies returns a boolean if a field has been set.

### GetK6Version

`func (o *ValidateOptionsRequest) GetK6Version() int32`

GetK6Version returns the K6Version field if non-nil, zero value otherwise.

### GetK6VersionOk

`func (o *ValidateOptionsRequest) GetK6VersionOk() (*int32, bool)`

GetK6VersionOk returns a tuple with the K6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Version

`func (o *ValidateOptionsRequest) SetK6Version(v int32)`

SetK6Version sets K6Version field to given value.

### HasK6Version

`func (o *ValidateOptionsRequest) HasK6Version() bool`

HasK6Version returns a boolean if a field has been set.

### SetK6VersionNil

`func (o *ValidateOptionsRequest) SetK6VersionNil(b bool)`

 SetK6VersionNil sets the value for K6Version to be an explicit nil

### UnsetK6Version
`func (o *ValidateOptionsRequest) UnsetK6Version()`

UnsetK6Version ensures that no value is present for K6Version, not even an explicit nil
### GetIsLocalExecution

`func (o *ValidateOptionsRequest) GetIsLocalExecution() bool`

GetIsLocalExecution returns the IsLocalExecution field if non-nil, zero value otherwise.

### GetIsLocalExecutionOk

`func (o *ValidateOptionsRequest) GetIsLocalExecutionOk() (*bool, bool)`

GetIsLocalExecutionOk returns a tuple with the IsLocalExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalExecution

`func (o *ValidateOptionsRequest) SetIsLocalExecution(v bool)`

SetIsLocalExecution sets IsLocalExecution field to given value.

### HasIsLocalExecution

`func (o *ValidateOptionsRequest) HasIsLocalExecution() bool`

HasIsLocalExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


