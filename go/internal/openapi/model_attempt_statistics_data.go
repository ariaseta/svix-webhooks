/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AttemptStatisticsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttemptStatisticsData{}

// AttemptStatisticsData struct for AttemptStatisticsData
type AttemptStatisticsData struct {
	FailureCount []int32 `json:"failureCount,omitempty"`
	SuccessCount []int32 `json:"successCount,omitempty"`
}

// NewAttemptStatisticsData instantiates a new AttemptStatisticsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttemptStatisticsData() *AttemptStatisticsData {
	this := AttemptStatisticsData{}
	return &this
}

// NewAttemptStatisticsDataWithDefaults instantiates a new AttemptStatisticsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttemptStatisticsDataWithDefaults() *AttemptStatisticsData {
	this := AttemptStatisticsData{}
	return &this
}

// GetFailureCount returns the FailureCount field value if set, zero value otherwise.
func (o *AttemptStatisticsData) GetFailureCount() []int32 {
	if o == nil || IsNil(o.FailureCount) {
		var ret []int32
		return ret
	}
	return o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttemptStatisticsData) GetFailureCountOk() ([]int32, bool) {
	if o == nil || IsNil(o.FailureCount) {
		return nil, false
	}
	return o.FailureCount, true
}

// HasFailureCount returns a boolean if a field has been set.
func (o *AttemptStatisticsData) HasFailureCount() bool {
	if o != nil && !IsNil(o.FailureCount) {
		return true
	}

	return false
}

// SetFailureCount gets a reference to the given []int32 and assigns it to the FailureCount field.
func (o *AttemptStatisticsData) SetFailureCount(v []int32) {
	o.FailureCount = v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *AttemptStatisticsData) GetSuccessCount() []int32 {
	if o == nil || IsNil(o.SuccessCount) {
		var ret []int32
		return ret
	}
	return o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttemptStatisticsData) GetSuccessCountOk() ([]int32, bool) {
	if o == nil || IsNil(o.SuccessCount) {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *AttemptStatisticsData) HasSuccessCount() bool {
	if o != nil && !IsNil(o.SuccessCount) {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given []int32 and assigns it to the SuccessCount field.
func (o *AttemptStatisticsData) SetSuccessCount(v []int32) {
	o.SuccessCount = v
}

func (o AttemptStatisticsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttemptStatisticsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailureCount) {
		toSerialize["failureCount"] = o.FailureCount
	}
	if !IsNil(o.SuccessCount) {
		toSerialize["successCount"] = o.SuccessCount
	}
	return toSerialize, nil
}

type NullableAttemptStatisticsData struct {
	value *AttemptStatisticsData
	isSet bool
}

func (v NullableAttemptStatisticsData) Get() *AttemptStatisticsData {
	return v.value
}

func (v *NullableAttemptStatisticsData) Set(val *AttemptStatisticsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttemptStatisticsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttemptStatisticsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttemptStatisticsData(val *AttemptStatisticsData) *NullableAttemptStatisticsData {
	return &NullableAttemptStatisticsData{value: val, isSet: true}
}

func (v NullableAttemptStatisticsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttemptStatisticsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


