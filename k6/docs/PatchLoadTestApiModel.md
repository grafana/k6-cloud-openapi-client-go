# PatchLoadTestApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name of the test within the project. | [optional] 
**BaselineTestRunId** | Pointer to **NullableInt32** | ID of a baseline test run used for results comparison. | [optional] 

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

`func (o *PatchLoadTestApiModel) GetBaselineTestRunId() int32`

GetBaselineTestRunId returns the BaselineTestRunId field if non-nil, zero value otherwise.

### GetBaselineTestRunIdOk

`func (o *PatchLoadTestApiModel) GetBaselineTestRunIdOk() (*int32, bool)`

GetBaselineTestRunIdOk returns a tuple with the BaselineTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineTestRunId

`func (o *PatchLoadTestApiModel) SetBaselineTestRunId(v int32)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


