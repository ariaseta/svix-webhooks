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

// checks if the BackgroundTaskFinishedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackgroundTaskFinishedEvent{}

// BackgroundTaskFinishedEvent Sent when a background task is finished.
type BackgroundTaskFinishedEvent struct {
	Data BackgroundTaskFinishedEvent2 `json:"data"`
	Type string `json:"type"`
}

type _BackgroundTaskFinishedEvent BackgroundTaskFinishedEvent

// NewBackgroundTaskFinishedEvent instantiates a new BackgroundTaskFinishedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundTaskFinishedEvent(data BackgroundTaskFinishedEvent2, type_ string) *BackgroundTaskFinishedEvent {
	this := BackgroundTaskFinishedEvent{}
	this.Data = data
	this.Type = type_
	return &this
}

// NewBackgroundTaskFinishedEventWithDefaults instantiates a new BackgroundTaskFinishedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundTaskFinishedEventWithDefaults() *BackgroundTaskFinishedEvent {
	this := BackgroundTaskFinishedEvent{}
	var type_ string = "background_task.finished"
	this.Type = type_
	return &this
}

// GetData returns the Data field value
func (o *BackgroundTaskFinishedEvent) GetData() BackgroundTaskFinishedEvent2 {
	if o == nil {
		var ret BackgroundTaskFinishedEvent2
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BackgroundTaskFinishedEvent) GetDataOk() (*BackgroundTaskFinishedEvent2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BackgroundTaskFinishedEvent) SetData(v BackgroundTaskFinishedEvent2) {
	o.Data = v
}

// GetType returns the Type field value
func (o *BackgroundTaskFinishedEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BackgroundTaskFinishedEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BackgroundTaskFinishedEvent) SetType(v string) {
	o.Type = v
}

func (o BackgroundTaskFinishedEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackgroundTaskFinishedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *BackgroundTaskFinishedEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"type",
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

	varBackgroundTaskFinishedEvent := _BackgroundTaskFinishedEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBackgroundTaskFinishedEvent)

	if err != nil {
		return err
	}

	*o = BackgroundTaskFinishedEvent(varBackgroundTaskFinishedEvent)

	return err
}

type NullableBackgroundTaskFinishedEvent struct {
	value *BackgroundTaskFinishedEvent
	isSet bool
}

func (v NullableBackgroundTaskFinishedEvent) Get() *BackgroundTaskFinishedEvent {
	return v.value
}

func (v *NullableBackgroundTaskFinishedEvent) Set(val *BackgroundTaskFinishedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundTaskFinishedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundTaskFinishedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundTaskFinishedEvent(val *BackgroundTaskFinishedEvent) *NullableBackgroundTaskFinishedEvent {
	return &NullableBackgroundTaskFinishedEvent{value: val, isSet: true}
}

func (v NullableBackgroundTaskFinishedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundTaskFinishedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


