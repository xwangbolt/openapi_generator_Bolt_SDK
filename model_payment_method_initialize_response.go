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

// checks if the PaymentMethodInitializeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodInitializeResponse{}

// PaymentMethodInitializeResponse struct for PaymentMethodInitializeResponse
type PaymentMethodInitializeResponse struct {
	Action *PaymentMethodInitializeResponseAction `json:"action,omitempty"`
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewPaymentMethodInitializeResponse instantiates a new PaymentMethodInitializeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodInitializeResponse() *PaymentMethodInitializeResponse {
	this := PaymentMethodInitializeResponse{}
	return &this
}

// NewPaymentMethodInitializeResponseWithDefaults instantiates a new PaymentMethodInitializeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodInitializeResponseWithDefaults() *PaymentMethodInitializeResponse {
	this := PaymentMethodInitializeResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PaymentMethodInitializeResponse) GetAction() PaymentMethodInitializeResponseAction {
	if o == nil || IsNil(o.Action) {
		var ret PaymentMethodInitializeResponseAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodInitializeResponse) GetActionOk() (*PaymentMethodInitializeResponseAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PaymentMethodInitializeResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given PaymentMethodInitializeResponseAction and assigns it to the Action field.
func (o *PaymentMethodInitializeResponse) SetAction(v PaymentMethodInitializeResponseAction) {
	o.Action = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentMethodInitializeResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodInitializeResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethodInitializeResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentMethodInitializeResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentMethodInitializeResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodInitializeResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentMethodInitializeResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentMethodInitializeResponse) SetStatus(v string) {
	o.Status = &v
}

func (o PaymentMethodInitializeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodInitializeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePaymentMethodInitializeResponse struct {
	value *PaymentMethodInitializeResponse
	isSet bool
}

func (v NullablePaymentMethodInitializeResponse) Get() *PaymentMethodInitializeResponse {
	return v.value
}

func (v *NullablePaymentMethodInitializeResponse) Set(val *PaymentMethodInitializeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodInitializeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodInitializeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodInitializeResponse(val *PaymentMethodInitializeResponse) *NullablePaymentMethodInitializeResponse {
	return &NullablePaymentMethodInitializeResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodInitializeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodInitializeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


