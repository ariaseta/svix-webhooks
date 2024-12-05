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

// checks if the AuthTokenOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenOut{}

// AuthTokenOut struct for AuthTokenOut
type AuthTokenOut struct {
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// The key's ID
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Token string `json:"token"`
}

type _AuthTokenOut AuthTokenOut

// NewAuthTokenOut instantiates a new AuthTokenOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenOut(createdAt time.Time, id string, token string) *AuthTokenOut {
	this := AuthTokenOut{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Token = token
	return &this
}

// NewAuthTokenOutWithDefaults instantiates a new AuthTokenOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenOutWithDefaults() *AuthTokenOut {
	this := AuthTokenOut{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *AuthTokenOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AuthTokenOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AuthTokenOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AuthTokenOut) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenOut) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AuthTokenOut) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AuthTokenOut) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *AuthTokenOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthTokenOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthTokenOut) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthTokenOut) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenOut) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthTokenOut) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthTokenOut) SetName(v string) {
	o.Name = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthTokenOut) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenOut) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthTokenOut) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *AuthTokenOut) SetScopes(v []string) {
	o.Scopes = v
}

// GetToken returns the Token field value
func (o *AuthTokenOut) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AuthTokenOut) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AuthTokenOut) SetToken(v string) {
	o.Token = v
}

func (o AuthTokenOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *AuthTokenOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"token",
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

	varAuthTokenOut := _AuthTokenOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthTokenOut)

	if err != nil {
		return err
	}

	*o = AuthTokenOut(varAuthTokenOut)

	return err
}

type NullableAuthTokenOut struct {
	value *AuthTokenOut
	isSet bool
}

func (v NullableAuthTokenOut) Get() *AuthTokenOut {
	return v.value
}

func (v *NullableAuthTokenOut) Set(val *AuthTokenOut) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenOut) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenOut(val *AuthTokenOut) *NullableAuthTokenOut {
	return &NullableAuthTokenOut{value: val, isSet: true}
}

func (v NullableAuthTokenOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


