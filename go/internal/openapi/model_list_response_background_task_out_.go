/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListResponseBackgroundTaskOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseBackgroundTaskOut{}

// ListResponseBackgroundTaskOut struct for ListResponseBackgroundTaskOut
type ListResponseBackgroundTaskOut struct {
	Data []BackgroundTaskOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator"`
	PrevIterator *string `json:"prevIterator,omitempty"`
}

type _ListResponseBackgroundTaskOut ListResponseBackgroundTaskOut

// NewListResponseBackgroundTaskOut instantiates a new ListResponseBackgroundTaskOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseBackgroundTaskOut(data []BackgroundTaskOut, done bool, iterator NullableString) *ListResponseBackgroundTaskOut {
	this := ListResponseBackgroundTaskOut{}
	this.Data = data
	this.Done = done
	this.Iterator = iterator
	return &this
}

// NewListResponseBackgroundTaskOutWithDefaults instantiates a new ListResponseBackgroundTaskOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseBackgroundTaskOutWithDefaults() *ListResponseBackgroundTaskOut {
	this := ListResponseBackgroundTaskOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseBackgroundTaskOut) GetData() []BackgroundTaskOut {
	if o == nil {
		var ret []BackgroundTaskOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseBackgroundTaskOut) GetDataOk() ([]BackgroundTaskOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseBackgroundTaskOut) SetData(v []BackgroundTaskOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseBackgroundTaskOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseBackgroundTaskOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseBackgroundTaskOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListResponseBackgroundTaskOut) GetIterator() string {
	if o == nil || o.Iterator.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseBackgroundTaskOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// SetIterator sets field value
func (o *ListResponseBackgroundTaskOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise.
func (o *ListResponseBackgroundTaskOut) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator) {
		var ret string
		return ret
	}
	return *o.PrevIterator
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResponseBackgroundTaskOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil || IsNil(o.PrevIterator) {
		return nil, false
	}
	return o.PrevIterator, true
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseBackgroundTaskOut) HasPrevIterator() bool {
	if o != nil && !IsNil(o.PrevIterator) {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given string and assigns it to the PrevIterator field.
func (o *ListResponseBackgroundTaskOut) SetPrevIterator(v string) {
	o.PrevIterator = &v
}

func (o ListResponseBackgroundTaskOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseBackgroundTaskOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	toSerialize["iterator"] = o.Iterator.Get()
	if !IsNil(o.PrevIterator) {
		toSerialize["prevIterator"] = o.PrevIterator
	}
	return toSerialize, nil
}

func (o *ListResponseBackgroundTaskOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"done",
		"iterator",
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

	varListResponseBackgroundTaskOut := _ListResponseBackgroundTaskOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseBackgroundTaskOut)

	if err != nil {
		return err
	}

	*o = ListResponseBackgroundTaskOut(varListResponseBackgroundTaskOut)

	return err
}

type NullableListResponseBackgroundTaskOut struct {
	value *ListResponseBackgroundTaskOut
	isSet bool
}

func (v NullableListResponseBackgroundTaskOut) Get() *ListResponseBackgroundTaskOut {
	return v.value
}

func (v *NullableListResponseBackgroundTaskOut) Set(val *ListResponseBackgroundTaskOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseBackgroundTaskOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseBackgroundTaskOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseBackgroundTaskOut(val *ListResponseBackgroundTaskOut) *NullableListResponseBackgroundTaskOut {
	return &NullableListResponseBackgroundTaskOut{value: val, isSet: true}
}

func (v NullableListResponseBackgroundTaskOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseBackgroundTaskOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


