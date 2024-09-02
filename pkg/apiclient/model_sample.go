/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Sample type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sample{}

// Sample struct for Sample
type Sample struct {
	Description string `json:"description"`
	GitUrl      string `json:"gitUrl"`
	Name        string `json:"name"`
}

type _Sample Sample

// NewSample instantiates a new Sample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSample(description string, gitUrl string, name string) *Sample {
	this := Sample{}
	this.Description = description
	this.GitUrl = gitUrl
	this.Name = name
	return &this
}

// NewSampleWithDefaults instantiates a new Sample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleWithDefaults() *Sample {
	this := Sample{}
	return &this
}

// GetDescription returns the Description field value
func (o *Sample) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Sample) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Sample) SetDescription(v string) {
	o.Description = v
}

// GetGitUrl returns the GitUrl field value
func (o *Sample) GetGitUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitUrl
}

// GetGitUrlOk returns a tuple with the GitUrl field value
// and a boolean to check if the value has been set.
func (o *Sample) GetGitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitUrl, true
}

// SetGitUrl sets field value
func (o *Sample) SetGitUrl(v string) {
	o.GitUrl = v
}

// GetName returns the Name field value
func (o *Sample) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Sample) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Sample) SetName(v string) {
	o.Name = v
}

func (o Sample) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sample) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["gitUrl"] = o.GitUrl
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *Sample) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"gitUrl",
		"name",
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

	varSample := _Sample{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSample)

	if err != nil {
		return err
	}

	*o = Sample(varSample)

	return err
}

type NullableSample struct {
	value *Sample
	isSet bool
}

func (v NullableSample) Get() *Sample {
	return v.value
}

func (v *NullableSample) Set(val *Sample) {
	v.value = val
	v.isSet = true
}

func (v NullableSample) IsSet() bool {
	return v.isSet
}

func (v *NullableSample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSample(val *Sample) *NullableSample {
	return &NullableSample{value: val, isSet: true}
}

func (v NullableSample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
