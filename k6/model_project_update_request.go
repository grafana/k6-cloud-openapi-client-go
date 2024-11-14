/*


HTTP API for interacting with k6 Cloud.

API version: 0.0.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
)

// checks if the ProjectUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectUpdateRequest{}

// ProjectUpdateRequest struct for ProjectUpdateRequest
type ProjectUpdateRequest struct {
	// Project name.
	Name *string `json:"name,omitempty"`
	// Project description.
	Description *string `json:"description,omitempty"`
	// Project specific settings.
	Settings             *ProjectSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectUpdateRequest ProjectUpdateRequest

// NewProjectUpdateRequest instantiates a new ProjectUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUpdateRequest() *ProjectUpdateRequest {
	this := ProjectUpdateRequest{}
	return &this
}

// NewProjectUpdateRequestWithDefaults instantiates a new ProjectUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUpdateRequestWithDefaults() *ProjectUpdateRequest {
	this := ProjectUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ProjectUpdateRequest) GetSettings() ProjectSettings {
	if o == nil || IsNil(o.Settings) {
		var ret ProjectSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateRequest) GetSettingsOk() (*ProjectSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ProjectUpdateRequest) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ProjectSettings and assigns it to the Settings field.
func (o *ProjectUpdateRequest) SetSettings(v ProjectSettings) {
	o.Settings = &v
}

func (o ProjectUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varProjectUpdateRequest := _ProjectUpdateRequest{}

	err = json.Unmarshal(data, &varProjectUpdateRequest)

	if err != nil {
		return err
	}

	*o = ProjectUpdateRequest(varProjectUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectUpdateRequest struct {
	value *ProjectUpdateRequest
	isSet bool
}

func (v NullableProjectUpdateRequest) Get() *ProjectUpdateRequest {
	return v.value
}

func (v *NullableProjectUpdateRequest) Set(val *ProjectUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUpdateRequest(val *ProjectUpdateRequest) *NullableProjectUpdateRequest {
	return &NullableProjectUpdateRequest{value: val, isSet: true}
}

func (v NullableProjectUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
