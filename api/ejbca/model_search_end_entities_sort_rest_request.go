/*
Copyright 2024 Keyfactor

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"encoding/json"
)

// checks if the SearchEndEntitiesSortRestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchEndEntitiesSortRestRequest{}

// SearchEndEntitiesSortRestRequest Use one of allowed values as property and operation. Available propertiesUSERNAME  SUBJECT_DN  SUBJECT_ALT_NAME  END_ENTITY_PROFILE(by databse identifier, not user-given name)  CERTIFICATE_PROFILE(by identifier)  CA(by identifier)  STATUS  UPDATE_TIME  CREATED_DATE   Available operationsASC  DESC
type SearchEndEntitiesSortRestRequest struct {
	// Sorted by
	Property *string `json:"property,omitempty"`
	// Sort ascending or descending. 'ASC' for ascending, 'DESC' for descending.
	Operation            *string `json:"operation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchEndEntitiesSortRestRequest SearchEndEntitiesSortRestRequest

// NewSearchEndEntitiesSortRestRequest instantiates a new SearchEndEntitiesSortRestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEndEntitiesSortRestRequest() *SearchEndEntitiesSortRestRequest {
	this := SearchEndEntitiesSortRestRequest{}
	return &this
}

// NewSearchEndEntitiesSortRestRequestWithDefaults instantiates a new SearchEndEntitiesSortRestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchEndEntitiesSortRestRequestWithDefaults() *SearchEndEntitiesSortRestRequest {
	this := SearchEndEntitiesSortRestRequest{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SearchEndEntitiesSortRestRequest) GetProperty() string {
	if o == nil || isNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEndEntitiesSortRestRequest) GetPropertyOk() (*string, bool) {
	if o == nil || isNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SearchEndEntitiesSortRestRequest) HasProperty() bool {
	if o != nil && !isNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SearchEndEntitiesSortRestRequest) SetProperty(v string) {
	o.Property = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *SearchEndEntitiesSortRestRequest) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEndEntitiesSortRestRequest) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *SearchEndEntitiesSortRestRequest) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *SearchEndEntitiesSortRestRequest) SetOperation(v string) {
	o.Operation = &v
}

func (o SearchEndEntitiesSortRestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchEndEntitiesSortRestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchEndEntitiesSortRestRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSearchEndEntitiesSortRestRequest := _SearchEndEntitiesSortRestRequest{}

	if err = json.Unmarshal(bytes, &varSearchEndEntitiesSortRestRequest); err == nil {
		*o = SearchEndEntitiesSortRestRequest(varSearchEndEntitiesSortRestRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "property")
		delete(additionalProperties, "operation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchEndEntitiesSortRestRequest struct {
	value *SearchEndEntitiesSortRestRequest
	isSet bool
}

func (v NullableSearchEndEntitiesSortRestRequest) Get() *SearchEndEntitiesSortRestRequest {
	return v.value
}

func (v *NullableSearchEndEntitiesSortRestRequest) Set(val *SearchEndEntitiesSortRestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEndEntitiesSortRestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEndEntitiesSortRestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEndEntitiesSortRestRequest(val *SearchEndEntitiesSortRestRequest) *NullableSearchEndEntitiesSortRestRequest {
	return &NullableSearchEndEntitiesSortRestRequest{value: val, isSet: true}
}

func (v NullableSearchEndEntitiesSortRestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEndEntitiesSortRestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
