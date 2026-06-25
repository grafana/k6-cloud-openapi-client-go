# PatchLoadTestApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name of the test within the project. | [optional] 
**BaselineTestRunId** | Pointer to **NullableInt64** | ID of a baseline test run used for results comparison. Deprecated: baselines are being replaced by the star/unstar test run APIs and this field is scheduled for removal on 2026-09-01. | [optional] 
**K6Version** | Pointer to **NullableInt32** | Identifier of the k6 version used to run the test. | [optional] 

## Methods

### NewPatchLoadTestApiModel

`func NewPatchLoadTestApiModel() *PatchLoadTestApiModel`

NewPatchLoadTestApiModel instantiates a new PatchLoadTestApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLoadTestApiModelWithDefaults

`func NewPatchLoadTestApiModelWithDefaults() *PatchLoadTestApiModel`

NewPatchLoadTestApiModelWithDefaults instantiates a new PatchLoadTestApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchLoadTestApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchLoadTestApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchLoadTestApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchLoadTestApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaselineTestRunId

`func (o *PatchLoadTestApiModel) GetBaselineTestRunId() int64`

GetBaselineTestRunId returns the BaselineTestRunId field if non-nil, zero value otherwise.

### GetBaselineTestRunIdOk

`func (o *PatchLoadTestApiModel) GetBaselineTestRunIdOk() (*int64, bool)`

GetBaselineTestRunIdOk returns a tuple with the BaselineTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineTestRunId

`func (o *PatchLoadTestApiModel) SetBaselineTestRunId(v int64)`

SetBaselineTestRunId sets BaselineTestRunId field to given value.

### HasBaselineTestRunId

`func (o *PatchLoadTestApiModel) HasBaselineTestRunId() bool`

HasBaselineTestRunId returns a boolean if a field has been set.

### SetBaselineTestRunIdNil

`func (o *PatchLoadTestApiModel) SetBaselineTestRunIdNil(b bool)`

 SetBaselineTestRunIdNil sets the value for BaselineTestRunId to be an explicit nil

### UnsetBaselineTestRunId
`func (o *PatchLoadTestApiModel) UnsetBaselineTestRunId()`

UnsetBaselineTestRunId ensures that no value is present for BaselineTestRunId, not even an explicit nil
### GetK6Version

`func (o *PatchLoadTestApiModel) GetK6Version() int32`

GetK6Version returns the K6Version field if non-nil, zero value otherwise.

### GetK6VersionOk

`func (o *PatchLoadTestApiModel) GetK6VersionOk() (*int32, bool)`

GetK6VersionOk returns a tuple with the K6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6Version

`func (o *PatchLoadTestApiModel) SetK6Version(v int32)`

SetK6Version sets K6Version field to given value.

### HasK6Version

`func (o *PatchLoadTestApiModel) HasK6Version() bool`

HasK6Version returns a boolean if a field has been set.

### SetK6VersionNil

`func (o *PatchLoadTestApiModel) SetK6VersionNil(b bool)`

 SetK6VersionNil sets the value for K6Version to be an explicit nil

### UnsetK6Version
`func (o *PatchLoadTestApiModel) UnsetK6Version()`

UnsetK6Version ensures that no value is present for K6Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


