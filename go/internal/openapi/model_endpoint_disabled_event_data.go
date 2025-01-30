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

// checks if the EndpointDisabledEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointDisabledEventData{}

// EndpointDisabledEventData Sent when an endpoint has been automatically disabled after continuous failures, or manually via an API call.
type EndpointDisabledEventData struct {
	// The app's ID
	AppId string `json:"appId"`
	// The app's UID
	AppUid *string `json:"appUid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The ep's ID
	EndpointId string `json:"endpointId"`
	// The ep's UID
	EndpointUid *string `json:"endpointUid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	FailSince *time.Time `json:"failSince,omitempty"`
	Trigger *EndpointDisabledTrigger `json:"trigger,omitempty"`
}

type _EndpointDisabledEventData EndpointDisabledEventData

// NewEndpointDisabledEventData instantiates a new EndpointDisabledEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDisabledEventData(appId string, endpointId string) *EndpointDisabledEventData {
	this := EndpointDisabledEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	return &this
}

// NewEndpointDisabledEventDataWithDefaults instantiates a new EndpointDisabledEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDisabledEventDataWithDefaults() *EndpointDisabledEventData {
	this := EndpointDisabledEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *EndpointDisabledEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEventData) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *EndpointDisabledEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise.
func (o *EndpointDisabledEventData) GetAppUid() string {
	if o == nil || IsNil(o.AppUid) {
		var ret string
		return ret
	}
	return *o.AppUid
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEventData) GetAppUidOk() (*string, bool) {
	if o == nil || IsNil(o.AppUid) {
		return nil, false
	}
	return o.AppUid, true
}

// HasAppUid returns a boolean if a field has been set.
func (o *EndpointDisabledEventData) HasAppUid() bool {
	if o != nil && !IsNil(o.AppUid) {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given string and assigns it to the AppUid field.
func (o *EndpointDisabledEventData) SetAppUid(v string) {
	o.AppUid = &v
}

// GetEndpointId returns the EndpointId field value
func (o *EndpointDisabledEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *EndpointDisabledEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetEndpointUid returns the EndpointUid field value if set, zero value otherwise.
func (o *EndpointDisabledEventData) GetEndpointUid() string {
	if o == nil || IsNil(o.EndpointUid) {
		var ret string
		return ret
	}
	return *o.EndpointUid
}

// GetEndpointUidOk returns a tuple with the EndpointUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEventData) GetEndpointUidOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointUid) {
		return nil, false
	}
	return o.EndpointUid, true
}

// HasEndpointUid returns a boolean if a field has been set.
func (o *EndpointDisabledEventData) HasEndpointUid() bool {
	if o != nil && !IsNil(o.EndpointUid) {
		return true
	}

	return false
}

// SetEndpointUid gets a reference to the given string and assigns it to the EndpointUid field.
func (o *EndpointDisabledEventData) SetEndpointUid(v string) {
	o.EndpointUid = &v
}

// GetFailSince returns the FailSince field value if set, zero value otherwise.
func (o *EndpointDisabledEventData) GetFailSince() time.Time {
	if o == nil || IsNil(o.FailSince) {
		var ret time.Time
		return ret
	}
	return *o.FailSince
}

// GetFailSinceOk returns a tuple with the FailSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEventData) GetFailSinceOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FailSince) {
		return nil, false
	}
	return o.FailSince, true
}

// HasFailSince returns a boolean if a field has been set.
func (o *EndpointDisabledEventData) HasFailSince() bool {
	if o != nil && !IsNil(o.FailSince) {
		return true
	}

	return false
}

// SetFailSince gets a reference to the given time.Time and assigns it to the FailSince field.
func (o *EndpointDisabledEventData) SetFailSince(v time.Time) {
	o.FailSince = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *EndpointDisabledEventData) GetTrigger() EndpointDisabledTrigger {
	if o == nil || IsNil(o.Trigger) {
		var ret EndpointDisabledTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEventData) GetTriggerOk() (*EndpointDisabledTrigger, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *EndpointDisabledEventData) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given EndpointDisabledTrigger and assigns it to the Trigger field.
func (o *EndpointDisabledEventData) SetTrigger(v EndpointDisabledTrigger) {
	o.Trigger = &v
}

func (o EndpointDisabledEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointDisabledEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if !IsNil(o.AppUid) {
		toSerialize["appUid"] = o.AppUid
	}
	toSerialize["endpointId"] = o.EndpointId
	if !IsNil(o.EndpointUid) {
		toSerialize["endpointUid"] = o.EndpointUid
	}
	if !IsNil(o.FailSince) {
		toSerialize["failSince"] = o.FailSince
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	return toSerialize, nil
}

func (o *EndpointDisabledEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appId",
		"endpointId",
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

	varEndpointDisabledEventData := _EndpointDisabledEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointDisabledEventData)

	if err != nil {
		return err
	}

	*o = EndpointDisabledEventData(varEndpointDisabledEventData)

	return err
}

type NullableEndpointDisabledEventData struct {
	value *EndpointDisabledEventData
	isSet bool
}

func (v NullableEndpointDisabledEventData) Get() *EndpointDisabledEventData {
	return v.value
}

func (v *NullableEndpointDisabledEventData) Set(val *EndpointDisabledEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDisabledEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDisabledEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDisabledEventData(val *EndpointDisabledEventData) *NullableEndpointDisabledEventData {
	return &NullableEndpointDisabledEventData{value: val, isSet: true}
}

func (v NullableEndpointDisabledEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDisabledEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


