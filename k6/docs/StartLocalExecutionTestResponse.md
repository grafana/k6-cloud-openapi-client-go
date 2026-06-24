# StartLocalExecutionTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestRunId** | **int64** | ID of the created test run. | 
**RuntimeConfig** | [**RuntimeConfig**](RuntimeConfig.md) |  | 
**ArchiveUploadUrl** | **NullableString** | URL to upload the k6 archive or null if not uploading. | 
**TestRunDetailsPageUrl** | **string** | URL to the Grafana web app. | 

## Methods

### NewStartLocalExecutionTestResponse

`func NewStartLocalExecutionTestResponse(testRunId int64, runtimeConfig RuntimeConfig, archiveUploadUrl NullableString, testRunDetailsPageUrl string, ) *StartLocalExecutionTestResponse`

NewStartLocalExecutionTestResponse instantiates a new StartLocalExecutionTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartLocalExecutionTestResponseWithDefaults

`func NewStartLocalExecutionTestResponseWithDefaults() *StartLocalExecutionTestResponse`

NewStartLocalExecutionTestResponseWithDefaults instantiates a new StartLocalExecutionTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestRunId

`func (o *StartLocalExecutionTestResponse) GetTestRunId() int64`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *StartLocalExecutionTestResponse) GetTestRunIdOk() (*int64, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *StartLocalExecutionTestResponse) SetTestRunId(v int64)`

SetTestRunId sets TestRunId field to given value.


### GetRuntimeConfig

`func (o *StartLocalExecutionTestResponse) GetRuntimeConfig() RuntimeConfig`

GetRuntimeConfig returns the RuntimeConfig field if non-nil, zero value otherwise.

### GetRuntimeConfigOk

`func (o *StartLocalExecutionTestResponse) GetRuntimeConfigOk() (*RuntimeConfig, bool)`

GetRuntimeConfigOk returns a tuple with the RuntimeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeConfig

`func (o *StartLocalExecutionTestResponse) SetRuntimeConfig(v RuntimeConfig)`

SetRuntimeConfig sets RuntimeConfig field to given value.


### GetArchiveUploadUrl

`func (o *StartLocalExecutionTestResponse) GetArchiveUploadUrl() string`

GetArchiveUploadUrl returns the ArchiveUploadUrl field if non-nil, zero value otherwise.

### GetArchiveUploadUrlOk

`func (o *StartLocalExecutionTestResponse) GetArchiveUploadUrlOk() (*string, bool)`

GetArchiveUploadUrlOk returns a tuple with the ArchiveUploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUploadUrl

`func (o *StartLocalExecutionTestResponse) SetArchiveUploadUrl(v string)`

SetArchiveUploadUrl sets ArchiveUploadUrl field to given value.


### SetArchiveUploadUrlNil

`func (o *StartLocalExecutionTestResponse) SetArchiveUploadUrlNil(b bool)`

 SetArchiveUploadUrlNil sets the value for ArchiveUploadUrl to be an explicit nil

### UnsetArchiveUploadUrl
`func (o *StartLocalExecutionTestResponse) UnsetArchiveUploadUrl()`

UnsetArchiveUploadUrl ensures that no value is present for ArchiveUploadUrl, not even an explicit nil
### GetTestRunDetailsPageUrl

`func (o *StartLocalExecutionTestResponse) GetTestRunDetailsPageUrl() string`

GetTestRunDetailsPageUrl returns the TestRunDetailsPageUrl field if non-nil, zero value otherwise.

### GetTestRunDetailsPageUrlOk

`func (o *StartLocalExecutionTestResponse) GetTestRunDetailsPageUrlOk() (*string, bool)`

GetTestRunDetailsPageUrlOk returns a tuple with the TestRunDetailsPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunDetailsPageUrl

`func (o *StartLocalExecutionTestResponse) SetTestRunDetailsPageUrl(v string)`

SetTestRunDetailsPageUrl sets TestRunDetailsPageUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


