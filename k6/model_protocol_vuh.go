/*
Grafana Cloud k6

HTTP API for interacting with Grafana Cloud k6.

API version: 1.0.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"encoding/json"
	"fmt"
)

// ProtocolVuh Number of VUH charged for the protocol part of the test run.
type ProtocolVuh struct {
	Float32 *float32
	String  *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProtocolVuh) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Float32
	err = json.Unmarshal(data, &dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			return nil // data stored in dst.Float32, return on the first match
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProtocolVuh)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProtocolVuh) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProtocolVuh struct {
	value *ProtocolVuh
	isSet bool
}

func (v NullableProtocolVuh) Get() *ProtocolVuh {
	return v.value
}

func (v *NullableProtocolVuh) Set(val *ProtocolVuh) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolVuh) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolVuh) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolVuh(val *ProtocolVuh) *NullableProtocolVuh {
	return &NullableProtocolVuh{value: val, isSet: true}
}

func (v NullableProtocolVuh) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolVuh) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
