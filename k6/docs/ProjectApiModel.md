# ProjectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Project ID. | 
**Name** | **string** | Project name. | 
**IsDefault** | **bool** | Use this project as default for running tests when no explicit project ID is provided. | 
**GrafanaFolderUid** | **NullableString** | Grafana folder UID. | 
**Created** | **time.Time** | The date when the project was created. | 
**Updated** | **time.Time** | The date when the project was last updated. | 

## Methods

### NewProjectApiModel

`func NewProjectApiModel(id int32, name string, isDefault bool, grafanaFolderUid NullableString, created time.Time, updated time.Time, ) *ProjectApiModel`

NewProjectApiModel instantiates a new ProjectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectApiModelWithDefaults

`func NewProjectApiModelWithDefaults() *ProjectApiModel`

NewProjectApiModelWithDefaults instantiates a new ProjectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectApiModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectApiModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectApiModel) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *ProjectApiModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProjectApiModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProjectApiModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetGrafanaFolderUid

`func (o *ProjectApiModel) GetGrafanaFolderUid() string`

GetGrafanaFolderUid returns the GrafanaFolderUid field if non-nil, zero value otherwise.

### GetGrafanaFolderUidOk

`func (o *ProjectApiModel) GetGrafanaFolderUidOk() (*string, bool)`

GetGrafanaFolderUidOk returns a tuple with the GrafanaFolderUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaFolderUid

`func (o *ProjectApiModel) SetGrafanaFolderUid(v string)`

SetGrafanaFolderUid sets GrafanaFolderUid field to given value.


### SetGrafanaFolderUidNil

`func (o *ProjectApiModel) SetGrafanaFolderUidNil(b bool)`

 SetGrafanaFolderUidNil sets the value for GrafanaFolderUid to be an explicit nil

### UnsetGrafanaFolderUid
`func (o *ProjectApiModel) UnsetGrafanaFolderUid()`

UnsetGrafanaFolderUid ensures that no value is present for GrafanaFolderUid, not even an explicit nil
### GetCreated

`func (o *ProjectApiModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProjectApiModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProjectApiModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *ProjectApiModel) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ProjectApiModel) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ProjectApiModel) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


