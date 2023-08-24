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

// checks if the CartItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartItem{}

// CartItem struct for CartItem
type CartItem struct {
	// The name of a given item.
	Name string `json:"name"`
	// This value is used by Bolt as an external reference to a given item.
	Reference string `json:"reference"`
	// A human-readable description of this cart item.
	Description *string `json:"description,omitempty"`
	// The total amount, in cents, of the item including its taxes if applicable.
	TotalAmount int64 `json:"total_amount"`
	// The price of one unit of the item; for example, the price of one pack of socks.
	UnitPrice int64 `json:"unit_price"`
	// The number of units that comprise this cart item.
	Quantity int64 `json:"quantity"`
	// Used to provide a link to the image associated with the item.
	ImageUrl *string `json:"image_url,omitempty"`
}

// NewCartItem instantiates a new CartItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartItem(name string, reference string, totalAmount int64, unitPrice int64, quantity int64) *CartItem {
	this := CartItem{}
	this.Name = name
	this.Reference = reference
	this.TotalAmount = totalAmount
	this.UnitPrice = unitPrice
	this.Quantity = quantity
	return &this
}

// NewCartItemWithDefaults instantiates a new CartItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartItemWithDefaults() *CartItem {
	this := CartItem{}
	return &this
}

// GetName returns the Name field value
func (o *CartItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CartItem) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value
func (o *CartItem) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *CartItem) SetReference(v string) {
	o.Reference = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CartItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CartItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CartItem) SetDescription(v string) {
	o.Description = &v
}

// GetTotalAmount returns the TotalAmount field value
func (o *CartItem) GetTotalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetTotalAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *CartItem) SetTotalAmount(v int64) {
	o.TotalAmount = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *CartItem) GetUnitPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetUnitPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *CartItem) SetUnitPrice(v int64) {
	o.UnitPrice = v
}

// GetQuantity returns the Quantity field value
func (o *CartItem) GetQuantity() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetQuantityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CartItem) SetQuantity(v int64) {
	o.Quantity = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *CartItem) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CartItem) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *CartItem) SetImageUrl(v string) {
	o.ImageUrl = &v
}

func (o CartItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["reference"] = o.Reference
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["total_amount"] = o.TotalAmount
	toSerialize["unit_price"] = o.UnitPrice
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	return toSerialize, nil
}

type NullableCartItem struct {
	value *CartItem
	isSet bool
}

func (v NullableCartItem) Get() *CartItem {
	return v.value
}

func (v *NullableCartItem) Set(val *CartItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCartItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCartItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartItem(val *CartItem) *NullableCartItem {
	return &NullableCartItem{value: val, isSet: true}
}

func (v NullableCartItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


