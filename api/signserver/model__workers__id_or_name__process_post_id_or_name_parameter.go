/*
Copyright 2022 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.

SignServer REST Interface

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package signserver

import (
	"encoding/json"
	"fmt"
)

// WorkersIdOrNameProcessPostIdOrNameParameter struct for WorkersIdOrNameProcessPostIdOrNameParameter
type WorkersIdOrNameProcessPostIdOrNameParameter struct {
	int32 *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *WorkersIdOrNameProcessPostIdOrNameParameter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32);
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(WorkersIdOrNameProcessPostIdOrNameParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *WorkersIdOrNameProcessPostIdOrNameParameter) MarshalJSON() ([]byte, error) {
	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableWorkersIdOrNameProcessPostIdOrNameParameter struct {
	value *WorkersIdOrNameProcessPostIdOrNameParameter
	isSet bool
}

func (v NullableWorkersIdOrNameProcessPostIdOrNameParameter) Get() *WorkersIdOrNameProcessPostIdOrNameParameter {
	return v.value
}

func (v *NullableWorkersIdOrNameProcessPostIdOrNameParameter) Set(val *WorkersIdOrNameProcessPostIdOrNameParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkersIdOrNameProcessPostIdOrNameParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkersIdOrNameProcessPostIdOrNameParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkersIdOrNameProcessPostIdOrNameParameter(val *WorkersIdOrNameProcessPostIdOrNameParameter) *NullableWorkersIdOrNameProcessPostIdOrNameParameter {
	return &NullableWorkersIdOrNameProcessPostIdOrNameParameter{value: val, isSet: true}
}

func (v NullableWorkersIdOrNameProcessPostIdOrNameParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkersIdOrNameProcessPostIdOrNameParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

