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

// checks if the VMUpdateSpecsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMUpdateSpecsRequest{}

// VMUpdateSpecsRequest struct for VMUpdateSpecsRequest
type VMUpdateSpecsRequest struct {
	// Determines which specs to update the VM with.  Not all tiers will be available depending on the workspace subscription status, and higher tiers incur higher costs. Please see codesandbox.io/pricing for details on specs and costs. 
	Tier string `json:"tier"`
}

type _VMUpdateSpecsRequest VMUpdateSpecsRequest

// NewVMUpdateSpecsRequest instantiates a new VMUpdateSpecsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMUpdateSpecsRequest(tier string) *VMUpdateSpecsRequest {
	this := VMUpdateSpecsRequest{}
	this.Tier = tier
	return &this
}

// NewVMUpdateSpecsRequestWithDefaults instantiates a new VMUpdateSpecsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMUpdateSpecsRequestWithDefaults() *VMUpdateSpecsRequest {
	this := VMUpdateSpecsRequest{}
	return &this
}

// GetTier returns the Tier field value
func (o *VMUpdateSpecsRequest) GetTier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *VMUpdateSpecsRequest) GetTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *VMUpdateSpecsRequest) SetTier(v string) {
	o.Tier = v
}

func (o VMUpdateSpecsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMUpdateSpecsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tier"] = o.Tier
	return toSerialize, nil
}

func (o *VMUpdateSpecsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varVMUpdateSpecsRequest := _VMUpdateSpecsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVMUpdateSpecsRequest)

	if err != nil {
		return err
	}

	*o = VMUpdateSpecsRequest(varVMUpdateSpecsRequest)

	return err
}

type NullableVMUpdateSpecsRequest struct {
	value *VMUpdateSpecsRequest
	isSet bool
}

func (v NullableVMUpdateSpecsRequest) Get() *VMUpdateSpecsRequest {
	return v.value
}

func (v *NullableVMUpdateSpecsRequest) Set(val *VMUpdateSpecsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVMUpdateSpecsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVMUpdateSpecsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMUpdateSpecsRequest(val *VMUpdateSpecsRequest) *NullableVMUpdateSpecsRequest {
	return &NullableVMUpdateSpecsRequest{value: val, isSet: true}
}

func (v NullableVMUpdateSpecsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMUpdateSpecsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


