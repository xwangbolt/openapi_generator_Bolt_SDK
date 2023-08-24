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

// checks if the CartDiscount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartDiscount{}

// CartDiscount struct for CartDiscount
type CartDiscount struct {
	Amounts Amounts `json:"amounts"`
	// Discount code.
	Code *string `json:"code,omitempty"`
	// Used to provide a link to additional details, such as a landing page, associated with the discount offering.
	DetailsUrl *string `json:"details_url,omitempty"`
}

// NewCartDiscount instantiates a new CartDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartDiscount(amounts Amounts) *CartDiscount {
	this := CartDiscount{}
	this.Amounts = amounts
	return &this
}

// NewCartDiscountWithDefaults instantiates a new CartDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartDiscountWithDefaults() *CartDiscount {
	this := CartDiscount{}
	return &this
}

// GetAmounts returns the Amounts field value
func (o *CartDiscount) GetAmounts() Amounts {
	if o == nil {
		var ret Amounts
		return ret
	}

	return o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value
// and a boolean to check if the value has been set.
func (o *CartDiscount) GetAmountsOk() (*Amounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amounts, true
}

// SetAmounts sets field value
func (o *CartDiscount) SetAmounts(v Amounts) {
	o.Amounts = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CartDiscount) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartDiscount) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CartDiscount) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CartDiscount) SetCode(v string) {
	o.Code = &v
}

// GetDetailsUrl returns the DetailsUrl field value if set, zero value otherwise.
func (o *CartDiscount) GetDetailsUrl() string {
	if o == nil || IsNil(o.DetailsUrl) {
		var ret string
		return ret
	}
	return *o.DetailsUrl
}

// GetDetailsUrlOk returns a tuple with the DetailsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartDiscount) GetDetailsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DetailsUrl) {
		return nil, false
	}
	return o.DetailsUrl, true
}

// HasDetailsUrl returns a boolean if a field has been set.
func (o *CartDiscount) HasDetailsUrl() bool {
	if o != nil && !IsNil(o.DetailsUrl) {
		return true
	}

	return false
}

// SetDetailsUrl gets a reference to the given string and assigns it to the DetailsUrl field.
func (o *CartDiscount) SetDetailsUrl(v string) {
	o.DetailsUrl = &v
}

func (o CartDiscount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartDiscount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amounts"] = o.Amounts
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.DetailsUrl) {
		toSerialize["details_url"] = o.DetailsUrl
	}
	return toSerialize, nil
}

type NullableCartDiscount struct {
	value *CartDiscount
	isSet bool
}

func (v NullableCartDiscount) Get() *CartDiscount {
	return v.value
}

func (v *NullableCartDiscount) Set(val *CartDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableCartDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableCartDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartDiscount(val *CartDiscount) *NullableCartDiscount {
	return &NullableCartDiscount{value: val, isSet: true}
}

func (v NullableCartDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

