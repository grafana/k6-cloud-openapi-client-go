# AisdkSessionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Short-lived access token (TTL matches the agentic test cap). | 
**ExpiresAt** | **time.Time** | ISO 8601 timestamp when the access token expires (UTC). | 
**BaseUrl** | **string** | Region-matched URL of the hosted ai-sdk endpoint. | 
**Namespace** | **string** | Stack-scoped namespace claim on the access token (e.g. &#39;stacks-12345&#39;). | 

## Methods

### NewAisdkSessionApiModel

`func NewAisdkSessionApiModel(accessToken string, expiresAt time.Time, baseUrl string, namespace string, ) *AisdkSessionApiModel`

NewAisdkSessionApiModel instantiates a new AisdkSessionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAisdkSessionApiModelWithDefaults

`func NewAisdkSessionApiModelWithDefaults() *AisdkSessionApiModel`

NewAisdkSessionApiModelWithDefaults instantiates a new AisdkSessionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AisdkSessionApiModel) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AisdkSessionApiModel) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AisdkSessionApiModel) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetExpiresAt

`func (o *AisdkSessionApiModel) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AisdkSessionApiModel) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AisdkSessionApiModel) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetBaseUrl

`func (o *AisdkSessionApiModel) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AisdkSessionApiModel) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AisdkSessionApiModel) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetNamespace

`func (o *AisdkSessionApiModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AisdkSessionApiModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AisdkSessionApiModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


