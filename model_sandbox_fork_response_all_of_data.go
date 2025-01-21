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

// checks if the SandboxForkResponseAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxForkResponseAllOfData{}

// SandboxForkResponseAllOfData struct for SandboxForkResponseAllOfData
type SandboxForkResponseAllOfData struct {
	Alias string `json:"alias"`
	Id string `json:"id"`
	StartResponse NullableSandboxForkResponseAllOfDataStartResponse `json:"start_response,omitempty"`
	Title NullableString `json:"title"`
}

type _SandboxForkResponseAllOfData SandboxForkResponseAllOfData

// NewSandboxForkResponseAllOfData instantiates a new SandboxForkResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxForkResponseAllOfData(alias string, id string, title NullableString) *SandboxForkResponseAllOfData {
	this := SandboxForkResponseAllOfData{}
	this.Alias = alias
	this.Id = id
	this.Title = title
	return &this
}

// NewSandboxForkResponseAllOfDataWithDefaults instantiates a new SandboxForkResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxForkResponseAllOfDataWithDefaults() *SandboxForkResponseAllOfData {
	this := SandboxForkResponseAllOfData{}
	return &this
}

// GetAlias returns the Alias field value
func (o *SandboxForkResponseAllOfData) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *SandboxForkResponseAllOfData) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *SandboxForkResponseAllOfData) SetAlias(v string) {
	o.Alias = v
}

// GetId returns the Id field value
func (o *SandboxForkResponseAllOfData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SandboxForkResponseAllOfData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SandboxForkResponseAllOfData) SetId(v string) {
	o.Id = v
}

// GetStartResponse returns the StartResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxForkResponseAllOfData) GetStartResponse() SandboxForkResponseAllOfDataStartResponse {
	if o == nil || IsNil(o.StartResponse.Get()) {
		var ret SandboxForkResponseAllOfDataStartResponse
		return ret
	}
	return *o.StartResponse.Get()
}

// GetStartResponseOk returns a tuple with the StartResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxForkResponseAllOfData) GetStartResponseOk() (*SandboxForkResponseAllOfDataStartResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartResponse.Get(), o.StartResponse.IsSet()
}

// HasStartResponse returns a boolean if a field has been set.
func (o *SandboxForkResponseAllOfData) HasStartResponse() bool {
	if o != nil && o.StartResponse.IsSet() {
		return true
	}

	return false
}

// SetStartResponse gets a reference to the given NullableSandboxForkResponseAllOfDataStartResponse and assigns it to the StartResponse field.
func (o *SandboxForkResponseAllOfData) SetStartResponse(v SandboxForkResponseAllOfDataStartResponse) {
	o.StartResponse.Set(&v)
}
// SetStartResponseNil sets the value for StartResponse to be an explicit nil
func (o *SandboxForkResponseAllOfData) SetStartResponseNil() {
	o.StartResponse.Set(nil)
}

// UnsetStartResponse ensures that no value is present for StartResponse, not even an explicit nil
func (o *SandboxForkResponseAllOfData) UnsetStartResponse() {
	o.StartResponse.Unset()
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SandboxForkResponseAllOfData) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxForkResponseAllOfData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *SandboxForkResponseAllOfData) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o SandboxForkResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxForkResponseAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["id"] = o.Id
	if o.StartResponse.IsSet() {
		toSerialize["start_response"] = o.StartResponse.Get()
	}
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *SandboxForkResponseAllOfData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"id",
		"title",
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

	varSandboxForkResponseAllOfData := _SandboxForkResponseAllOfData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSandboxForkResponseAllOfData)

	if err != nil {
		return err
	}

	*o = SandboxForkResponseAllOfData(varSandboxForkResponseAllOfData)

	return err
}

type NullableSandboxForkResponseAllOfData struct {
	value *SandboxForkResponseAllOfData
	isSet bool
}

func (v NullableSandboxForkResponseAllOfData) Get() *SandboxForkResponseAllOfData {
	return v.value
}

func (v *NullableSandboxForkResponseAllOfData) Set(val *SandboxForkResponseAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxForkResponseAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxForkResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxForkResponseAllOfData(val *SandboxForkResponseAllOfData) *NullableSandboxForkResponseAllOfData {
	return &NullableSandboxForkResponseAllOfData{value: val, isSet: true}
}

func (v NullableSandboxForkResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxForkResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


