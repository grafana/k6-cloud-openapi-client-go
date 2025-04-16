/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.2.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// checks if the ErrorDetailsApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDetailsApiModel{}

// ErrorDetailsApiModel struct for ErrorDetailsApiModel
type ErrorDetailsApiModel struct {
	// Human-readable string describing the error.
	Message string `json:"message"`
	// Service-defined error code.
	Code string `json:"code"`
	// A string indicating the target of the error. For example, the name of the property in error.
	Target               NullableString `json:"target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorDetailsApiModel ErrorDetailsApiModel

// NewErrorDetailsApiModel instantiates a new ErrorDetailsApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetailsApiModel(message string, code string) *ErrorDetailsApiModel {
	this := ErrorDetailsApiModel{}
	this.Message = message
	this.Code = code
	return &this
}

// NewErrorDetailsApiModelWithDefaults instantiates a new ErrorDetailsApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailsApiModelWithDefaults() *ErrorDetailsApiModel {
	this := ErrorDetailsApiModel{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorDetailsApiModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailsApiModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorDetailsApiModel) SetMessage(v string) {
	o.Message = v
}

// GetCode returns the Code field value
func (o *ErrorDetailsApiModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailsApiModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorDetailsApiModel) SetCode(v string) {
	o.Code = v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorDetailsApiModel) GetTarget() string {
	if o == nil || IsNil(o.Target.Get()) {
		var ret string
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorDetailsApiModel) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *ErrorDetailsApiModel) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableString and assigns it to the Target field.
func (o *ErrorDetailsApiModel) SetTarget(v string) {
	o.Target.Set(&v)
}

// SetTargetNil sets the value for Target to be an explicit nil
func (o *ErrorDetailsApiModel) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *ErrorDetailsApiModel) UnsetTarget() {
	o.Target.Unset()
}

func (o ErrorDetailsApiModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDetailsApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["code"] = o.Code
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorDetailsApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorDetailsApiModel := _ErrorDetailsApiModel{}

	err = json.Unmarshal(data, &varErrorDetailsApiModel)

	if err != nil {
		return err
	}

	*o = ErrorDetailsApiModel(varErrorDetailsApiModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "code")
		delete(additionalProperties, "target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorDetailsApiModel struct {
	value *ErrorDetailsApiModel
	isSet bool
}

func (v NullableErrorDetailsApiModel) Get() *ErrorDetailsApiModel {
	return v.value
}

func (v *NullableErrorDetailsApiModel) Set(val *ErrorDetailsApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetailsApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetailsApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetailsApiModel(val *ErrorDetailsApiModel) *NullableErrorDetailsApiModel {
	return &NullableErrorDetailsApiModel{value: val, isSet: true}
}

func (v NullableErrorDetailsApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetailsApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
