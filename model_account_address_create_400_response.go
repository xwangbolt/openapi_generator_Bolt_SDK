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

// AccountAddressCreate400Response - struct for AccountAddressCreate400Response
type AccountAddressCreate400Response struct {
	AddressErrorInvalidPostalCode *AddressErrorInvalidPostalCode
	AddressErrorInvalidRegion *AddressErrorInvalidRegion
}

// AddressErrorInvalidPostalCodeAsAccountAddressCreate400Response is a convenience function that returns AddressErrorInvalidPostalCode wrapped in AccountAddressCreate400Response
func AddressErrorInvalidPostalCodeAsAccountAddressCreate400Response(v *AddressErrorInvalidPostalCode) AccountAddressCreate400Response {
	return AccountAddressCreate400Response{
		AddressErrorInvalidPostalCode: v,
	}
}

// AddressErrorInvalidRegionAsAccountAddressCreate400Response is a convenience function that returns AddressErrorInvalidRegion wrapped in AccountAddressCreate400Response
func AddressErrorInvalidRegionAsAccountAddressCreate400Response(v *AddressErrorInvalidRegion) AccountAddressCreate400Response {
	return AccountAddressCreate400Response{
		AddressErrorInvalidRegion: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccountAddressCreate400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddressErrorInvalidPostalCode
	err = newStrictDecoder(data).Decode(&dst.AddressErrorInvalidPostalCode)
	if err == nil {
		jsonAddressErrorInvalidPostalCode, _ := json.Marshal(dst.AddressErrorInvalidPostalCode)
		if string(jsonAddressErrorInvalidPostalCode) == "{}" { // empty struct
			dst.AddressErrorInvalidPostalCode = nil
		} else {
			match++
		}
	} else {
		dst.AddressErrorInvalidPostalCode = nil
	}

	// try to unmarshal data into AddressErrorInvalidRegion
	err = newStrictDecoder(data).Decode(&dst.AddressErrorInvalidRegion)
	if err == nil {
		jsonAddressErrorInvalidRegion, _ := json.Marshal(dst.AddressErrorInvalidRegion)
		if string(jsonAddressErrorInvalidRegion) == "{}" { // empty struct
			dst.AddressErrorInvalidRegion = nil
		} else {
			match++
		}
	} else {
		dst.AddressErrorInvalidRegion = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddressErrorInvalidPostalCode = nil
		dst.AddressErrorInvalidRegion = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AccountAddressCreate400Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AccountAddressCreate400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccountAddressCreate400Response) MarshalJSON() ([]byte, error) {
	if src.AddressErrorInvalidPostalCode != nil {
		return json.Marshal(&src.AddressErrorInvalidPostalCode)
	}

	if src.AddressErrorInvalidRegion != nil {
		return json.Marshal(&src.AddressErrorInvalidRegion)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccountAddressCreate400Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddressErrorInvalidPostalCode != nil {
		return obj.AddressErrorInvalidPostalCode
	}

	if obj.AddressErrorInvalidRegion != nil {
		return obj.AddressErrorInvalidRegion
	}

	// all schemas are nil
	return nil
}

type NullableAccountAddressCreate400Response struct {
	value *AccountAddressCreate400Response
	isSet bool
}

func (v NullableAccountAddressCreate400Response) Get() *AccountAddressCreate400Response {
	return v.value
}

func (v *NullableAccountAddressCreate400Response) Set(val *AccountAddressCreate400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAddressCreate400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAddressCreate400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAddressCreate400Response(val *AccountAddressCreate400Response) *NullableAccountAddressCreate400Response {
	return &NullableAccountAddressCreate400Response{value: val, isSet: true}
}

func (v NullableAccountAddressCreate400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAddressCreate400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

