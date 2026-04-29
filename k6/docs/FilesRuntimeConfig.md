# FilesRuntimeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushUrl** | **string** | URL to generate a pre-signed URLs for file uploads. | 
**ScreenshotsBasePath** | **string** | Base path for screenshot uploads. | 
**Headers** | **map[string]string** | HTTP headers to include in requests. | 
**Method** | **string** | HTTP method to use for screenshot uploads. | 

## Methods

### NewFilesRuntimeConfig

`func NewFilesRuntimeConfig(pushUrl string, screenshotsBasePath string, headers map[string]string, method string, ) *FilesRuntimeConfig`

NewFilesRuntimeConfig instantiates a new FilesRuntimeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesRuntimeConfigWithDefaults

`func NewFilesRuntimeConfigWithDefaults() *FilesRuntimeConfig`

NewFilesRuntimeConfigWithDefaults instantiates a new FilesRuntimeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushUrl

`func (o *FilesRuntimeConfig) GetPushUrl() string`

GetPushUrl returns the PushUrl field if non-nil, zero value otherwise.

### GetPushUrlOk

`func (o *FilesRuntimeConfig) GetPushUrlOk() (*string, bool)`

GetPushUrlOk returns a tuple with the PushUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushUrl

`func (o *FilesRuntimeConfig) SetPushUrl(v string)`

SetPushUrl sets PushUrl field to given value.


### GetScreenshotsBasePath

`func (o *FilesRuntimeConfig) GetScreenshotsBasePath() string`

GetScreenshotsBasePath returns the ScreenshotsBasePath field if non-nil, zero value otherwise.

### GetScreenshotsBasePathOk

`func (o *FilesRuntimeConfig) GetScreenshotsBasePathOk() (*string, bool)`

GetScreenshotsBasePathOk returns a tuple with the ScreenshotsBasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotsBasePath

`func (o *FilesRuntimeConfig) SetScreenshotsBasePath(v string)`

SetScreenshotsBasePath sets ScreenshotsBasePath field to given value.


### GetHeaders

`func (o *FilesRuntimeConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FilesRuntimeConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FilesRuntimeConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetMethod

`func (o *FilesRuntimeConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FilesRuntimeConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FilesRuntimeConfig) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


