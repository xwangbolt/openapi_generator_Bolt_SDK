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

// MerchantCallbacksUpdate400Response - struct for MerchantCallbacksUpdate400Response
type MerchantCallbacksUpdate400Response struct {
	CallbackUrlErrorInvalidUrl *CallbackUrlErrorInvalidUrl
}

// CallbackUrlErrorInvalidUrlAsMerchantCallbacksUpdate400Response is a convenience function that returns CallbackUrlErrorInvalidUrl wrapped in MerchantCallbacksUpdate400Response
func CallbackUrlErrorInvalidUrlAsMerchantCallbacksUpdate400Response(v *CallbackUrlErrorInvalidUrl) MerchantCallbacksUpdate400Response {
	return MerchantCallbacksUpdate400Response{
		CallbackUrlErrorInvalidUrl: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MerchantCallbacksUpdate400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CallbackUrlErrorInvalidUrl
	err = newStrictDecoder(data).Decode(&dst.CallbackUrlErrorInvalidUrl)
	if err == nil {
		jsonCallbackUrlErrorInvalidUrl, _ := json.Marshal(dst.CallbackUrlErrorInvalidUrl)
		if string(jsonCallbackUrlErrorInvalidUrl) == "{}" { // empty struct
			dst.CallbackUrlErrorInvalidUrl = nil
		} else {
			match++
		}
	} else {
		dst.CallbackUrlErrorInvalidUrl = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CallbackUrlErrorInvalidUrl = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MerchantCallbacksUpdate400Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MerchantCallbacksUpdate400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MerchantCallbacksUpdate400Response) MarshalJSON() ([]byte, error) {
	if src.CallbackUrlErrorInvalidUrl != nil {
		return json.Marshal(&src.CallbackUrlErrorInvalidUrl)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MerchantCallbacksUpdate400Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CallbackUrlErrorInvalidUrl != nil {
		return obj.CallbackUrlErrorInvalidUrl
	}

	// all schemas are nil
	return nil
}

type NullableMerchantCallbacksUpdate400Response struct {
	value *MerchantCallbacksUpdate400Response
	isSet bool
}

func (v NullableMerchantCallbacksUpdate400Response) Get() *MerchantCallbacksUpdate400Response {
	return v.value
}

func (v *NullableMerchantCallbacksUpdate400Response) Set(val *MerchantCallbacksUpdate400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCallbacksUpdate400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCallbacksUpdate400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCallbacksUpdate400Response(val *MerchantCallbacksUpdate400Response) *NullableMerchantCallbacksUpdate400Response {
	return &NullableMerchantCallbacksUpdate400Response{value: val, isSet: true}
}

func (v NullableMerchantCallbacksUpdate400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCallbacksUpdate400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

