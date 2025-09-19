# AuthenticationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **int32** | The ID of the Grafana stack matching the provided URL. | 
**DefaultProjectId** | **int32** | The ID of the default project in the stack. | 

## Methods

### NewAuthenticationResponse

`func NewAuthenticationResponse(stackId int32, defaultProjectId int32, ) *AuthenticationResponse`

NewAuthenticationResponse instantiates a new AuthenticationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationResponseWithDefaults

`func NewAuthenticationResponseWithDefaults() *AuthenticationResponse`

NewAuthenticationResponseWithDefaults instantiates a new AuthenticationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *AuthenticationResponse) GetStackId() int32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *AuthenticationResponse) GetStackIdOk() (*int32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *AuthenticationResponse) SetStackId(v int32)`

SetStackId sets StackId field to given value.


### GetDefaultProjectId

`func (o *AuthenticationResponse) GetDefaultProjectId() int32`

GetDefaultProjectId returns the DefaultProjectId field if non-nil, zero value otherwise.

### GetDefaultProjectIdOk

`func (o *AuthenticationResponse) GetDefaultProjectIdOk() (*int32, bool)`

GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProjectId

`func (o *AuthenticationResponse) SetDefaultProjectId(v int32)`

SetDefaultProjectId sets DefaultProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


