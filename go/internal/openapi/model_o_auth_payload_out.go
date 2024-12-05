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

// checks if the OAuthPayloadOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthPayloadOut{}

// OAuthPayloadOut struct for OAuthPayloadOut
type OAuthPayloadOut struct {
	AccessToken *string `json:"accessToken,omitempty"`
	Error *string `json:"error,omitempty"`
	RefreshToken *string `json:"refreshToken,omitempty"`
}

// NewOAuthPayloadOut instantiates a new OAuthPayloadOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthPayloadOut() *OAuthPayloadOut {
	this := OAuthPayloadOut{}
	return &this
}

// NewOAuthPayloadOutWithDefaults instantiates a new OAuthPayloadOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthPayloadOutWithDefaults() *OAuthPayloadOut {
	this := OAuthPayloadOut{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *OAuthPayloadOut) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthPayloadOut) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OAuthPayloadOut) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *OAuthPayloadOut) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OAuthPayloadOut) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthPayloadOut) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OAuthPayloadOut) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *OAuthPayloadOut) SetError(v string) {
	o.Error = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *OAuthPayloadOut) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthPayloadOut) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OAuthPayloadOut) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *OAuthPayloadOut) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o OAuthPayloadOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthPayloadOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	return toSerialize, nil
}

type NullableOAuthPayloadOut struct {
	value *OAuthPayloadOut
	isSet bool
}

func (v NullableOAuthPayloadOut) Get() *OAuthPayloadOut {
	return v.value
}

func (v *NullableOAuthPayloadOut) Set(val *OAuthPayloadOut) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthPayloadOut) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthPayloadOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthPayloadOut(val *OAuthPayloadOut) *NullableOAuthPayloadOut {
	return &NullableOAuthPayloadOut{value: val, isSet: true}
}

func (v NullableOAuthPayloadOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthPayloadOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


