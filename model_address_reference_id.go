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

// checks if the AddressReferenceId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressReferenceId{}

// AddressReferenceId struct for AddressReferenceId
type AddressReferenceId struct {
	// The type of address reference
	Tag string `json:".tag"`
	// The address's ID
	Id string `json:"id"`
}

// NewAddressReferenceId instantiates a new AddressReferenceId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressReferenceId(tag string, id string) *AddressReferenceId {
	this := AddressReferenceId{}
	this.Tag = tag
	this.Id = id
	return &this
}

// NewAddressReferenceIdWithDefaults instantiates a new AddressReferenceId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressReferenceIdWithDefaults() *AddressReferenceId {
	this := AddressReferenceId{}
	return &this
}

// GetTag returns the Tag field value
func (o *AddressReferenceId) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceId) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *AddressReferenceId) SetTag(v string) {
	o.Tag = v
}

// GetId returns the Id field value
func (o *AddressReferenceId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddressReferenceId) SetId(v string) {
	o.Id = v
}

func (o AddressReferenceId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressReferenceId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize[".tag"] = o.Tag
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAddressReferenceId struct {
	value *AddressReferenceId
	isSet bool
}

func (v NullableAddressReferenceId) Get() *AddressReferenceId {
	return v.value
}

func (v *NullableAddressReferenceId) Set(val *AddressReferenceId) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressReferenceId) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressReferenceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressReferenceId(val *AddressReferenceId) *NullableAddressReferenceId {
	return &NullableAddressReferenceId{value: val, isSet: true}
}

func (v NullableAddressReferenceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressReferenceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


