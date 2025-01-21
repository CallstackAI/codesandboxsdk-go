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

// checks if the VMUpdateSpecsResponseAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMUpdateSpecsResponseAllOfData{}

// VMUpdateSpecsResponseAllOfData struct for VMUpdateSpecsResponseAllOfData
type VMUpdateSpecsResponseAllOfData struct {
	Id string `json:"id"`
	Tier string `json:"tier"`
}

type _VMUpdateSpecsResponseAllOfData VMUpdateSpecsResponseAllOfData

// NewVMUpdateSpecsResponseAllOfData instantiates a new VMUpdateSpecsResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMUpdateSpecsResponseAllOfData(id string, tier string) *VMUpdateSpecsResponseAllOfData {
	this := VMUpdateSpecsResponseAllOfData{}
	this.Id = id
	this.Tier = tier
	return &this
}

// NewVMUpdateSpecsResponseAllOfDataWithDefaults instantiates a new VMUpdateSpecsResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMUpdateSpecsResponseAllOfDataWithDefaults() *VMUpdateSpecsResponseAllOfData {
	this := VMUpdateSpecsResponseAllOfData{}
	return &this
}

// GetId returns the Id field value
func (o *VMUpdateSpecsResponseAllOfData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMUpdateSpecsResponseAllOfData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMUpdateSpecsResponseAllOfData) SetId(v string) {
	o.Id = v
}

// GetTier returns the Tier field value
func (o *VMUpdateSpecsResponseAllOfData) GetTier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *VMUpdateSpecsResponseAllOfData) GetTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *VMUpdateSpecsResponseAllOfData) SetTier(v string) {
	o.Tier = v
}

func (o VMUpdateSpecsResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMUpdateSpecsResponseAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["tier"] = o.Tier
	return toSerialize, nil
}

func (o *VMUpdateSpecsResponseAllOfData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"tier",
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

	varVMUpdateSpecsResponseAllOfData := _VMUpdateSpecsResponseAllOfData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVMUpdateSpecsResponseAllOfData)

	if err != nil {
		return err
	}

	*o = VMUpdateSpecsResponseAllOfData(varVMUpdateSpecsResponseAllOfData)

	return err
}

type NullableVMUpdateSpecsResponseAllOfData struct {
	value *VMUpdateSpecsResponseAllOfData
	isSet bool
}

func (v NullableVMUpdateSpecsResponseAllOfData) Get() *VMUpdateSpecsResponseAllOfData {
	return v.value
}

func (v *NullableVMUpdateSpecsResponseAllOfData) Set(val *VMUpdateSpecsResponseAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableVMUpdateSpecsResponseAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableVMUpdateSpecsResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMUpdateSpecsResponseAllOfData(val *VMUpdateSpecsResponseAllOfData) *NullableVMUpdateSpecsResponseAllOfData {
	return &NullableVMUpdateSpecsResponseAllOfData{value: val, isSet: true}
}

func (v NullableVMUpdateSpecsResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMUpdateSpecsResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


