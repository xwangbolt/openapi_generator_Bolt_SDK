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
	"time"
)

// checks if the AccountTestCreationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountTestCreationData{}

// AccountTestCreationData struct for AccountTestCreationData
type AccountTestCreationData struct {
	Email string `json:"email"`
	EmailState string `json:"email_state"`
	Phone string `json:"phone"`
	PhoneState string `json:"phone_state"`
	IsMigrated *bool `json:"is_migrated,omitempty"`
	HasAddress *bool `json:"has_address,omitempty"`
	OtpCode string `json:"otp_code"`
	OauthCode string `json:"oauth_code"`
	DeactivateAt time.Time `json:"deactivate_at"`
}

// NewAccountTestCreationData instantiates a new AccountTestCreationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTestCreationData(email string, emailState string, phone string, phoneState string, otpCode string, oauthCode string, deactivateAt time.Time) *AccountTestCreationData {
	this := AccountTestCreationData{}
	this.Email = email
	this.EmailState = emailState
	this.Phone = phone
	this.PhoneState = phoneState
	this.OtpCode = otpCode
	this.OauthCode = oauthCode
	this.DeactivateAt = deactivateAt
	return &this
}

// NewAccountTestCreationDataWithDefaults instantiates a new AccountTestCreationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTestCreationDataWithDefaults() *AccountTestCreationData {
	this := AccountTestCreationData{}
	return &this
}

// GetEmail returns the Email field value
func (o *AccountTestCreationData) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AccountTestCreationData) SetEmail(v string) {
	o.Email = v
}

// GetEmailState returns the EmailState field value
func (o *AccountTestCreationData) GetEmailState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailState
}

// GetEmailStateOk returns a tuple with the EmailState field value
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetEmailStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailState, true
}

// SetEmailState sets field value
func (o *AccountTestCreationData) SetEmailState(v string) {
	o.EmailState = v
}

// GetPhone returns the Phone field value
func (o *AccountTestCreationData) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *AccountTestCreationData) SetPhone(v string) {
	o.Phone = v
}

// GetPhoneState returns the PhoneState field value
func (o *AccountTestCreationData) GetPhoneState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneState
}

// GetPhoneStateOk returns a tuple with the PhoneState field value
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetPhoneStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneState, true
}

// SetPhoneState sets field value
func (o *AccountTestCreationData) SetPhoneState(v string) {
	o.PhoneState = v
}

// GetIsMigrated returns the IsMigrated field value if set, zero value otherwise.
func (o *AccountTestCreationData) GetIsMigrated() bool {
	if o == nil || IsNil(o.IsMigrated) {
		var ret bool
		return ret
	}
	return *o.IsMigrated
}

// GetIsMigratedOk returns a tuple with the IsMigrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetIsMigratedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMigrated) {
		return nil, false
	}
	return o.IsMigrated, true
}

// HasIsMigrated returns a boolean if a field has been set.
func (o *AccountTestCreationData) HasIsMigrated() bool {
	if o != nil && !IsNil(o.IsMigrated) {
		return true
	}

	return false
}

// SetIsMigrated gets a reference to the given bool and assigns it to the IsMigrated field.
func (o *AccountTestCreationData) SetIsMigrated(v bool) {
	o.IsMigrated = &v
}

// GetHasAddress returns the HasAddress field value if set, zero value otherwise.
func (o *AccountTestCreationData) GetHasAddress() bool {
	if o == nil || IsNil(o.HasAddress) {
		var ret bool
		return ret
	}
	return *o.HasAddress
}

// GetHasAddressOk returns a tuple with the HasAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetHasAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAddress) {
		return nil, false
	}
	return o.HasAddress, true
}

// HasHasAddress returns a boolean if a field has been set.
func (o *AccountTestCreationData) HasHasAddress() bool {
	if o != nil && !IsNil(o.HasAddress) {
		return true
	}

	return false
}

// SetHasAddress gets a reference to the given bool and assigns it to the HasAddress field.
func (o *AccountTestCreationData) SetHasAddress(v bool) {
	o.HasAddress = &v
}

// GetOtpCode returns the OtpCode field value
func (o *AccountTestCreationData) GetOtpCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OtpCode
}

// GetOtpCodeOk returns a tuple with the OtpCode field value
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetOtpCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtpCode, true
}

// SetOtpCode sets field value
func (o *AccountTestCreationData) SetOtpCode(v string) {
	o.OtpCode = v
}

// GetOauthCode returns the OauthCode field value
func (o *AccountTestCreationData) GetOauthCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OauthCode
}

// GetOauthCodeOk returns a tuple with the OauthCode field value
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetOauthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OauthCode, true
}

// SetOauthCode sets field value
func (o *AccountTestCreationData) SetOauthCode(v string) {
	o.OauthCode = v
}

// GetDeactivateAt returns the DeactivateAt field value
func (o *AccountTestCreationData) GetDeactivateAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeactivateAt
}

// GetDeactivateAtOk returns a tuple with the DeactivateAt field value
// and a boolean to check if the value has been set.
func (o *AccountTestCreationData) GetDeactivateAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeactivateAt, true
}

// SetDeactivateAt sets field value
func (o *AccountTestCreationData) SetDeactivateAt(v time.Time) {
	o.DeactivateAt = v
}

func (o AccountTestCreationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountTestCreationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: email is readOnly
	toSerialize["email_state"] = o.EmailState
	// skip: phone is readOnly
	toSerialize["phone_state"] = o.PhoneState
	if !IsNil(o.IsMigrated) {
		toSerialize["is_migrated"] = o.IsMigrated
	}
	if !IsNil(o.HasAddress) {
		toSerialize["has_address"] = o.HasAddress
	}
	// skip: otp_code is readOnly
	// skip: oauth_code is readOnly
	toSerialize["deactivate_at"] = o.DeactivateAt
	return toSerialize, nil
}

type NullableAccountTestCreationData struct {
	value *AccountTestCreationData
	isSet bool
}

func (v NullableAccountTestCreationData) Get() *AccountTestCreationData {
	return v.value
}

func (v *NullableAccountTestCreationData) Set(val *AccountTestCreationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTestCreationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTestCreationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTestCreationData(val *AccountTestCreationData) *NullableAccountTestCreationData {
	return &NullableAccountTestCreationData{value: val, isSet: true}
}

func (v NullableAccountTestCreationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTestCreationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


