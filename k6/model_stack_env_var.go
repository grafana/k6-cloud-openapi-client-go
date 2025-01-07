/*


HTTP API for interacting with k6 Cloud.

API version: 0.0.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"time"
)

// checks if the StackEnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StackEnvVar{}

// StackEnvVar struct for StackEnvVar
type StackEnvVar struct {
	// Id of the variable.
	Id int32 `json:"id"`
	// Name of the variable.
	Name string `json:"name"`
	// Variable value.
	Value string `json:"value"`
	// Description of the variable.
	Description NullableString `json:"description,omitempty"`
	// The date when the variable was created.
	Created time.Time `json:"created"`
	// The date when the variable was last time updated.
	Updated              time.Time `json:"updated"`
	AdditionalProperties map[string]interface{}
}

type _StackEnvVar StackEnvVar

// NewStackEnvVar instantiates a new StackEnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackEnvVar(id int32, name string, value string, created time.Time, updated time.Time) *StackEnvVar {
	this := StackEnvVar{}
	this.Id = id
	this.Name = name
	this.Value = value
	this.Created = created
	this.Updated = updated
	return &this
}

// NewStackEnvVarWithDefaults instantiates a new StackEnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackEnvVarWithDefaults() *StackEnvVar {
	this := StackEnvVar{}
	return &this
}

// GetId returns the Id field value
func (o *StackEnvVar) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StackEnvVar) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StackEnvVar) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *StackEnvVar) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StackEnvVar) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StackEnvVar) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *StackEnvVar) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StackEnvVar) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StackEnvVar) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StackEnvVar) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StackEnvVar) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StackEnvVar) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StackEnvVar) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StackEnvVar) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StackEnvVar) UnsetDescription() {
	o.Description.Unset()
}

// GetCreated returns the Created field value
func (o *StackEnvVar) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *StackEnvVar) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *StackEnvVar) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *StackEnvVar) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *StackEnvVar) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *StackEnvVar) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o StackEnvVar) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StackEnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StackEnvVar) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varStackEnvVar := _StackEnvVar{}

	err = json.Unmarshal(data, &varStackEnvVar)

	if err != nil {
		return err
	}

	*o = StackEnvVar(varStackEnvVar)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStackEnvVar struct {
	value *StackEnvVar
	isSet bool
}

func (v NullableStackEnvVar) Get() *StackEnvVar {
	return v.value
}

func (v *NullableStackEnvVar) Set(val *StackEnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableStackEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableStackEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackEnvVar(val *StackEnvVar) *NullableStackEnvVar {
	return &NullableStackEnvVar{value: val, isSet: true}
}

func (v NullableStackEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
