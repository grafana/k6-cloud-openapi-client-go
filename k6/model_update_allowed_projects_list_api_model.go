/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.5.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateAllowedProjectsListApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAllowedProjectsListApiModel{}

// UpdateAllowedProjectsListApiModel struct for UpdateAllowedProjectsListApiModel
type UpdateAllowedProjectsListApiModel struct {
	// List of the projects that will become allowed to use load zone.
	Value []AllowedProjectToUpdateApiModel `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAllowedProjectsListApiModel UpdateAllowedProjectsListApiModel

// NewUpdateAllowedProjectsListApiModel instantiates a new UpdateAllowedProjectsListApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAllowedProjectsListApiModel(value []AllowedProjectToUpdateApiModel) *UpdateAllowedProjectsListApiModel {
	this := UpdateAllowedProjectsListApiModel{}
	this.Value = value
	return &this
}

// NewUpdateAllowedProjectsListApiModelWithDefaults instantiates a new UpdateAllowedProjectsListApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAllowedProjectsListApiModelWithDefaults() *UpdateAllowedProjectsListApiModel {
	this := UpdateAllowedProjectsListApiModel{}
	return &this
}

// GetValue returns the Value field value
func (o *UpdateAllowedProjectsListApiModel) GetValue() []AllowedProjectToUpdateApiModel {
	if o == nil {
		var ret []AllowedProjectToUpdateApiModel
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateAllowedProjectsListApiModel) GetValueOk() ([]AllowedProjectToUpdateApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *UpdateAllowedProjectsListApiModel) SetValue(v []AllowedProjectToUpdateApiModel) {
	o.Value = v
}

func (o UpdateAllowedProjectsListApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAllowedProjectsListApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAllowedProjectsListApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateAllowedProjectsListApiModel := _UpdateAllowedProjectsListApiModel{}

	err = json.Unmarshal(data, &varUpdateAllowedProjectsListApiModel)

	if err != nil {
		return err
	}

	*o = UpdateAllowedProjectsListApiModel(varUpdateAllowedProjectsListApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAllowedProjectsListApiModel struct {
	value *UpdateAllowedProjectsListApiModel
	isSet bool
}

func (v NullableUpdateAllowedProjectsListApiModel) Get() *UpdateAllowedProjectsListApiModel {
	return v.value
}

func (v *NullableUpdateAllowedProjectsListApiModel) Set(val *UpdateAllowedProjectsListApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAllowedProjectsListApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAllowedProjectsListApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAllowedProjectsListApiModel(val *UpdateAllowedProjectsListApiModel) *NullableUpdateAllowedProjectsListApiModel {
	return &NullableUpdateAllowedProjectsListApiModel{value: val, isSet: true}
}

func (v NullableUpdateAllowedProjectsListApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAllowedProjectsListApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


