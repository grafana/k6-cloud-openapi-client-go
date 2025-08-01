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

// checks if the MoveLoadTestApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveLoadTestApiModel{}

// MoveLoadTestApiModel struct for MoveLoadTestApiModel
type MoveLoadTestApiModel struct {
	// ID of the destination project.
	ProjectId int32 `json:"project_id"`
	AdditionalProperties map[string]interface{}
}

type _MoveLoadTestApiModel MoveLoadTestApiModel

// NewMoveLoadTestApiModel instantiates a new MoveLoadTestApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveLoadTestApiModel(projectId int32) *MoveLoadTestApiModel {
	this := MoveLoadTestApiModel{}
	this.ProjectId = projectId
	return &this
}

// NewMoveLoadTestApiModelWithDefaults instantiates a new MoveLoadTestApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveLoadTestApiModelWithDefaults() *MoveLoadTestApiModel {
	this := MoveLoadTestApiModel{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *MoveLoadTestApiModel) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *MoveLoadTestApiModel) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *MoveLoadTestApiModel) SetProjectId(v int32) {
	o.ProjectId = v
}

func (o MoveLoadTestApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveLoadTestApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MoveLoadTestApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
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

	varMoveLoadTestApiModel := _MoveLoadTestApiModel{}

	err = json.Unmarshal(data, &varMoveLoadTestApiModel)

	if err != nil {
		return err
	}

	*o = MoveLoadTestApiModel(varMoveLoadTestApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoveLoadTestApiModel struct {
	value *MoveLoadTestApiModel
	isSet bool
}

func (v NullableMoveLoadTestApiModel) Get() *MoveLoadTestApiModel {
	return v.value
}

func (v *NullableMoveLoadTestApiModel) Set(val *MoveLoadTestApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveLoadTestApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveLoadTestApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveLoadTestApiModel(val *MoveLoadTestApiModel) *NullableMoveLoadTestApiModel {
	return &NullableMoveLoadTestApiModel{value: val, isSet: true}
}

func (v NullableMoveLoadTestApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveLoadTestApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


