# StackEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the variable. | [readonly] 
**Name** | **string** | Name of the variable. | 
**Value** | **string** | Variable value. | 
**Description** | Pointer to **string** | Description of the variable. | [optional] 
**Created** | **time.Time** | The date when the variable was created. | [readonly] 
**Updated** | **time.Time** | The date when the variable was last time updated. | [readonly] 

## Methods

### NewStackEnvVar

`func NewStackEnvVar(id int32, name string, value string, created time.Time, updated time.Time, ) *StackEnvVar`

NewStackEnvVar instantiates a new StackEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackEnvVarWithDefaults

`func NewStackEnvVarWithDefaults() *StackEnvVar`

NewStackEnvVarWithDefaults instantiates a new StackEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackEnvVar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackEnvVar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackEnvVar) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StackEnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackEnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackEnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *StackEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StackEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StackEnvVar) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *StackEnvVar) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StackEnvVar) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StackEnvVar) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StackEnvVar) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *StackEnvVar) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StackEnvVar) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StackEnvVar) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *StackEnvVar) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *StackEnvVar) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *StackEnvVar) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


