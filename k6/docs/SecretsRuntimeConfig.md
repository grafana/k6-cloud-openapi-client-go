# SecretsRuntimeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | URL template for the secrets decrypt endpoint. | 
**ResponsePath** | **string** | JSON path to extract the secret value from the response. | 

## Methods

### NewSecretsRuntimeConfig

`func NewSecretsRuntimeConfig(endpoint string, responsePath string, ) *SecretsRuntimeConfig`

NewSecretsRuntimeConfig instantiates a new SecretsRuntimeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsRuntimeConfigWithDefaults

`func NewSecretsRuntimeConfigWithDefaults() *SecretsRuntimeConfig`

NewSecretsRuntimeConfigWithDefaults instantiates a new SecretsRuntimeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *SecretsRuntimeConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SecretsRuntimeConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SecretsRuntimeConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetResponsePath

`func (o *SecretsRuntimeConfig) GetResponsePath() string`

GetResponsePath returns the ResponsePath field if non-nil, zero value otherwise.

### GetResponsePathOk

`func (o *SecretsRuntimeConfig) GetResponsePathOk() (*string, bool)`

GetResponsePathOk returns a tuple with the ResponsePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePath

`func (o *SecretsRuntimeConfig) SetResponsePath(v string)`

SetResponsePath sets ResponsePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


