# ValidateOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** | Id of a project where the test belongs | [optional] 
**Options** | **map[string]interface{}** | k6 script options object to validate | 

## Methods

### NewValidateOptionsRequest

`func NewValidateOptionsRequest(options map[string]interface{}, ) *ValidateOptionsRequest`

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

### GetOptions

`func (o *ValidateOptionsRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ValidateOptionsRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ValidateOptionsRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


