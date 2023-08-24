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
	"fmt"
)

// GuestPaymentMethodInitializeRequestPaymentMethod - struct for GuestPaymentMethodInitializeRequestPaymentMethod
type GuestPaymentMethodInitializeRequestPaymentMethod struct {
	PaymentMethodPaypal *PaymentMethodPaypal
}

// PaymentMethodPaypalAsGuestPaymentMethodInitializeRequestPaymentMethod is a convenience function that returns PaymentMethodPaypal wrapped in GuestPaymentMethodInitializeRequestPaymentMethod
func PaymentMethodPaypalAsGuestPaymentMethodInitializeRequestPaymentMethod(v *PaymentMethodPaypal) GuestPaymentMethodInitializeRequestPaymentMethod {
	return GuestPaymentMethodInitializeRequestPaymentMethod{
		PaymentMethodPaypal: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GuestPaymentMethodInitializeRequestPaymentMethod) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PaymentMethodPaypal
	err = newStrictDecoder(data).Decode(&dst.PaymentMethodPaypal)
	if err == nil {
		jsonPaymentMethodPaypal, _ := json.Marshal(dst.PaymentMethodPaypal)
		if string(jsonPaymentMethodPaypal) == "{}" { // empty struct
			dst.PaymentMethodPaypal = nil
		} else {
			match++
		}
	} else {
		dst.PaymentMethodPaypal = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PaymentMethodPaypal = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GuestPaymentMethodInitializeRequestPaymentMethod)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GuestPaymentMethodInitializeRequestPaymentMethod)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GuestPaymentMethodInitializeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodPaypal != nil {
		return json.Marshal(&src.PaymentMethodPaypal)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GuestPaymentMethodInitializeRequestPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodPaypal != nil {
		return obj.PaymentMethodPaypal
	}

	// all schemas are nil
	return nil
}

type NullableGuestPaymentMethodInitializeRequestPaymentMethod struct {
	value *GuestPaymentMethodInitializeRequestPaymentMethod
	isSet bool
}

func (v NullableGuestPaymentMethodInitializeRequestPaymentMethod) Get() *GuestPaymentMethodInitializeRequestPaymentMethod {
	return v.value
}

func (v *NullableGuestPaymentMethodInitializeRequestPaymentMethod) Set(val *GuestPaymentMethodInitializeRequestPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestPaymentMethodInitializeRequestPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestPaymentMethodInitializeRequestPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestPaymentMethodInitializeRequestPaymentMethod(val *GuestPaymentMethodInitializeRequestPaymentMethod) *NullableGuestPaymentMethodInitializeRequestPaymentMethod {
	return &NullableGuestPaymentMethodInitializeRequestPaymentMethod{value: val, isSet: true}
}

func (v NullableGuestPaymentMethodInitializeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestPaymentMethodInitializeRequestPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


