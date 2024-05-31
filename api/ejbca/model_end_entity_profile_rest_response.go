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

// checks if the EndEntityProfileRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndEntityProfileRestResponse{}

// EndEntityProfileRestResponse struct for EndEntityProfileRestResponse
type EndEntityProfileRestResponse struct {
	// End Entity profile name
	Name *string `json:"name,omitempty"`
	// End Entity profile ID
	Id *int64 `json:"id,omitempty"`
	// Description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EndEntityProfileRestResponse EndEntityProfileRestResponse

// NewEndEntityProfileRestResponse instantiates a new EndEntityProfileRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndEntityProfileRestResponse() *EndEntityProfileRestResponse {
	this := EndEntityProfileRestResponse{}
	return &this
}

// NewEndEntityProfileRestResponseWithDefaults instantiates a new EndEntityProfileRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndEntityProfileRestResponseWithDefaults() *EndEntityProfileRestResponse {
	this := EndEntityProfileRestResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EndEntityProfileRestResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndEntityProfileRestResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EndEntityProfileRestResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EndEntityProfileRestResponse) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EndEntityProfileRestResponse) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndEntityProfileRestResponse) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EndEntityProfileRestResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *EndEntityProfileRestResponse) SetId(v int64) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EndEntityProfileRestResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndEntityProfileRestResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EndEntityProfileRestResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EndEntityProfileRestResponse) SetDescription(v string) {
	o.Description = &v
}

func (o EndEntityProfileRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndEntityProfileRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EndEntityProfileRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEndEntityProfileRestResponse := _EndEntityProfileRestResponse{}

	if err = json.Unmarshal(bytes, &varEndEntityProfileRestResponse); err == nil {
		*o = EndEntityProfileRestResponse(varEndEntityProfileRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndEntityProfileRestResponse struct {
	value *EndEntityProfileRestResponse
	isSet bool
}

func (v NullableEndEntityProfileRestResponse) Get() *EndEntityProfileRestResponse {
	return v.value
}

func (v *NullableEndEntityProfileRestResponse) Set(val *EndEntityProfileRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEndEntityProfileRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEndEntityProfileRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndEntityProfileRestResponse(val *EndEntityProfileRestResponse) *NullableEndEntityProfileRestResponse {
	return &NullableEndEntityProfileRestResponse{value: val, isSet: true}
}

func (v NullableEndEntityProfileRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndEntityProfileRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
