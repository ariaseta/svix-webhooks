/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the StreamOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamOut{}

// StreamOut struct for StreamOut
type StreamOut struct {
	CreatedAt time.Time `json:"createdAt"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	Uid *string `json:"uid,omitempty" validate:"regexp=^(?!strm_)[a-zA-Z0-9_-]+$"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type _StreamOut StreamOut

// NewStreamOut instantiates a new StreamOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamOut(createdAt time.Time, id string, updatedAt time.Time) *StreamOut {
	this := StreamOut{}
	this.CreatedAt = createdAt
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewStreamOutWithDefaults instantiates a new StreamOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamOutWithDefaults() *StreamOut {
	this := StreamOut{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *StreamOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *StreamOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *StreamOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StreamOut) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamOut) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StreamOut) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StreamOut) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *StreamOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StreamOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StreamOut) SetId(v string) {
	o.Id = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *StreamOut) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamOut) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *StreamOut) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *StreamOut) SetUid(v string) {
	o.Uid = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *StreamOut) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *StreamOut) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *StreamOut) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o StreamOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *StreamOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"updatedAt",
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

	varStreamOut := _StreamOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStreamOut)

	if err != nil {
		return err
	}

	*o = StreamOut(varStreamOut)

	return err
}

type NullableStreamOut struct {
	value *StreamOut
	isSet bool
}

func (v NullableStreamOut) Get() *StreamOut {
	return v.value
}

func (v *NullableStreamOut) Set(val *StreamOut) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamOut) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamOut(val *StreamOut) *NullableStreamOut {
	return &NullableStreamOut{value: val, isSet: true}
}

func (v NullableStreamOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


