/*
CodeSandbox API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-07-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csdkgo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MetaInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaInformation{}

// MetaInformation struct for MetaInformation
type MetaInformation struct {
	Api MetaInformationApi `json:"api"`
	Auth *MetaInformationAuth `json:"auth,omitempty"`
}

type _MetaInformation MetaInformation

// NewMetaInformation instantiates a new MetaInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaInformation(api MetaInformationApi) *MetaInformation {
	this := MetaInformation{}
	this.Api = api
	return &this
}

// NewMetaInformationWithDefaults instantiates a new MetaInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaInformationWithDefaults() *MetaInformation {
	this := MetaInformation{}
	return &this
}

// GetApi returns the Api field value
func (o *MetaInformation) GetApi() MetaInformationApi {
	if o == nil {
		var ret MetaInformationApi
		return ret
	}

	return o.Api
}

// GetApiOk returns a tuple with the Api field value
// and a boolean to check if the value has been set.
func (o *MetaInformation) GetApiOk() (*MetaInformationApi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Api, true
}

// SetApi sets field value
func (o *MetaInformation) SetApi(v MetaInformationApi) {
	o.Api = v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *MetaInformation) GetAuth() MetaInformationAuth {
	if o == nil || IsNil(o.Auth) {
		var ret MetaInformationAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaInformation) GetAuthOk() (*MetaInformationAuth, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *MetaInformation) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given MetaInformationAuth and assigns it to the Auth field.
func (o *MetaInformation) SetAuth(v MetaInformationAuth) {
	o.Auth = &v
}

func (o MetaInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api"] = o.Api
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	return toSerialize, nil
}

func (o *MetaInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"api",
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

	varMetaInformation := _MetaInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetaInformation)

	if err != nil {
		return err
	}

	*o = MetaInformation(varMetaInformation)

	return err
}

type NullableMetaInformation struct {
	value *MetaInformation
	isSet bool
}

func (v NullableMetaInformation) Get() *MetaInformation {
	return v.value
}

func (v *NullableMetaInformation) Set(val *MetaInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaInformation(val *MetaInformation) *NullableMetaInformation {
	return &NullableMetaInformation{value: val, isSet: true}
}

func (v NullableMetaInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


