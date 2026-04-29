# StartLocalExecutionTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | **map[string]interface{}** | Parsed k6 script options. | 
**MaxVus** | **int32** | Maximum number of virtual users during the test. | 
**TotalDuration** | **int32** | Total duration of the test in seconds. | 
**Instances** | Pointer to **NullableInt32** | Number of load generator instances that will run the test (k6-operator only). | [optional] 
**ArchiveSize** | Pointer to **NullableInt32** | K6 archive file size in bytes or null if not uploading. | [optional] 

## Methods

### NewStartLocalExecutionTestRequest

`func NewStartLocalExecutionTestRequest(options map[string]interface{}, maxVus int32, totalDuration int32, ) *StartLocalExecutionTestRequest`

NewStartLocalExecutionTestRequest instantiates a new StartLocalExecutionTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartLocalExecutionTestRequestWithDefaults

`func NewStartLocalExecutionTestRequestWithDefaults() *StartLocalExecutionTestRequest`

NewStartLocalExecutionTestRequestWithDefaults instantiates a new StartLocalExecutionTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *StartLocalExecutionTestRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StartLocalExecutionTestRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StartLocalExecutionTestRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### GetMaxVus

`func (o *StartLocalExecutionTestRequest) GetMaxVus() int32`

GetMaxVus returns the MaxVus field if non-nil, zero value otherwise.

### GetMaxVusOk

`func (o *StartLocalExecutionTestRequest) GetMaxVusOk() (*int32, bool)`

GetMaxVusOk returns a tuple with the MaxVus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVus

`func (o *StartLocalExecutionTestRequest) SetMaxVus(v int32)`

SetMaxVus sets MaxVus field to given value.


### GetTotalDuration

`func (o *StartLocalExecutionTestRequest) GetTotalDuration() int32`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *StartLocalExecutionTestRequest) GetTotalDurationOk() (*int32, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *StartLocalExecutionTestRequest) SetTotalDuration(v int32)`

SetTotalDuration sets TotalDuration field to given value.


### GetInstances

`func (o *StartLocalExecutionTestRequest) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *StartLocalExecutionTestRequest) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *StartLocalExecutionTestRequest) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *StartLocalExecutionTestRequest) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *StartLocalExecutionTestRequest) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *StartLocalExecutionTestRequest) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetArchiveSize

`func (o *StartLocalExecutionTestRequest) GetArchiveSize() int32`

GetArchiveSize returns the ArchiveSize field if non-nil, zero value otherwise.

### GetArchiveSizeOk

`func (o *StartLocalExecutionTestRequest) GetArchiveSizeOk() (*int32, bool)`

GetArchiveSizeOk returns a tuple with the ArchiveSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveSize

`func (o *StartLocalExecutionTestRequest) SetArchiveSize(v int32)`

SetArchiveSize sets ArchiveSize field to given value.

### HasArchiveSize

`func (o *StartLocalExecutionTestRequest) HasArchiveSize() bool`

HasArchiveSize returns a boolean if a field has been set.

### SetArchiveSizeNil

`func (o *StartLocalExecutionTestRequest) SetArchiveSizeNil(b bool)`

 SetArchiveSizeNil sets the value for ArchiveSize to be an explicit nil

### UnsetArchiveSize
`func (o *StartLocalExecutionTestRequest) UnsetArchiveSize()`

UnsetArchiveSize ensures that no value is present for ArchiveSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


