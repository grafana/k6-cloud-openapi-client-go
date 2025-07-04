# ValidateOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **NullableInt32** | ID of a project where the test belongs. | [optional] 
**Options** | [**Options**](Options.md) | k6 script options object to validate. | 
**K6Dependencies** | Pointer to **map[string]string** | Version of k6 and extensions to validate, as a map of dependency name to dependency version constraint. | [optional] 

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

`func (o *ValidateOptionsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ValidateOptionsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ValidateOptionsRequest) SetProjectId(v int32)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


