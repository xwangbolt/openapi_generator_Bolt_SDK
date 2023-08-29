/*
Bolt API Reference

A comprehensive Bolt API reference for interacting with Transactions, Orders, Product Catalog, Configuration, Testing, and much more.

API version: 3.0.1
Contact: api-help@bolt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AccountExistsIdentifierParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountExistsIdentifierParameter{}

// AccountExistsIdentifierParameter struct for AccountExistsIdentifierParameter
type AccountExistsIdentifierParameter struct {
	// The type of identifier
	Type string `json:"type"`
	// The value of the identifier. The value must be valid for the specified `identifier_type`
	Value string `json:"value"`
}

// NewAccountExistsIdentifierParameter instantiates a new AccountExistsIdentifierParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountExistsIdentifierParameter(type_ string, value string) *AccountExistsIdentifierParameter {
	this := AccountExistsIdentifierParameter{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAccountExistsIdentifierParameterWithDefaults instantiates a new AccountExistsIdentifierParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountExistsIdentifierParameterWithDefaults() *AccountExistsIdentifierParameter {
	this := AccountExistsIdentifierParameter{}
	return &this
}

// GetType returns the Type field value
func (o *AccountExistsIdentifierParameter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountExistsIdentifierParameter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountExistsIdentifierParameter) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AccountExistsIdentifierParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AccountExistsIdentifierParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AccountExistsIdentifierParameter) SetValue(v string) {
	o.Value = v
}

func (o AccountExistsIdentifierParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountExistsIdentifierParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableAccountExistsIdentifierParameter struct {
	value *AccountExistsIdentifierParameter
	isSet bool
}

func (v NullableAccountExistsIdentifierParameter) Get() *AccountExistsIdentifierParameter {
	return v.value
}

func (v *NullableAccountExistsIdentifierParameter) Set(val *AccountExistsIdentifierParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountExistsIdentifierParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountExistsIdentifierParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountExistsIdentifierParameter(val *AccountExistsIdentifierParameter) *NullableAccountExistsIdentifierParameter {
	return &NullableAccountExistsIdentifierParameter{value: val, isSet: true}
}

func (v NullableAccountExistsIdentifierParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountExistsIdentifierParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


