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

// checks if the CertificatesRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificatesRestResponse{}

// CertificatesRestResponse struct for CertificatesRestResponse
type CertificatesRestResponse struct {
	Certificates         []CertificateRestResponse `json:"certificates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificatesRestResponse CertificatesRestResponse

// NewCertificatesRestResponse instantiates a new CertificatesRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificatesRestResponse() *CertificatesRestResponse {
	this := CertificatesRestResponse{}
	return &this
}

// NewCertificatesRestResponseWithDefaults instantiates a new CertificatesRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatesRestResponseWithDefaults() *CertificatesRestResponse {
	this := CertificatesRestResponse{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *CertificatesRestResponse) GetCertificates() []CertificateRestResponse {
	if o == nil || isNil(o.Certificates) {
		var ret []CertificateRestResponse
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatesRestResponse) GetCertificatesOk() ([]CertificateRestResponse, bool) {
	if o == nil || isNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *CertificatesRestResponse) HasCertificates() bool {
	if o != nil && !isNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []CertificateRestResponse and assigns it to the Certificates field.
func (o *CertificatesRestResponse) SetCertificates(v []CertificateRestResponse) {
	o.Certificates = v
}

func (o CertificatesRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificatesRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificatesRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCertificatesRestResponse := _CertificatesRestResponse{}

	if err = json.Unmarshal(bytes, &varCertificatesRestResponse); err == nil {
		*o = CertificatesRestResponse(varCertificatesRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "certificates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificatesRestResponse struct {
	value *CertificatesRestResponse
	isSet bool
}

func (v NullableCertificatesRestResponse) Get() *CertificatesRestResponse {
	return v.value
}

func (v *NullableCertificatesRestResponse) Set(val *CertificatesRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatesRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatesRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatesRestResponse(val *CertificatesRestResponse) *NullableCertificatesRestResponse {
	return &NullableCertificatesRestResponse{value: val, isSet: true}
}

func (v NullableCertificatesRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatesRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
