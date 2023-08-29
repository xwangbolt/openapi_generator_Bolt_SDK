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

// checks if the AddressReferenceExplicit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressReferenceExplicit{}

// AddressReferenceExplicit struct for AddressReferenceExplicit
type AddressReferenceExplicit struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Company *string `json:"company,omitempty"`
	StreetAddress1 string `json:"street_address1"`
	StreetAddress2 *string `json:"street_address2,omitempty"`
	Locality string `json:"locality"`
	PostalCode string `json:"postal_code"`
	Region *string `json:"region,omitempty"`
	CountryCode string `json:"country_code"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	// The type of address reference
	Tag string `json:".tag"`
}

// NewAddressReferenceExplicit instantiates a new AddressReferenceExplicit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressReferenceExplicit(id string, firstName string, lastName string, streetAddress1 string, locality string, postalCode string, countryCode string, tag string) *AddressReferenceExplicit {
	this := AddressReferenceExplicit{}
	this.Id = id
	this.FirstName = firstName
	this.LastName = lastName
	this.StreetAddress1 = streetAddress1
	this.Locality = locality
	this.PostalCode = postalCode
	this.CountryCode = countryCode
	this.Tag = tag
	return &this
}

// NewAddressReferenceExplicitWithDefaults instantiates a new AddressReferenceExplicit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressReferenceExplicitWithDefaults() *AddressReferenceExplicit {
	this := AddressReferenceExplicit{}
	return &this
}

// GetId returns the Id field value
func (o *AddressReferenceExplicit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddressReferenceExplicit) SetId(v string) {
	o.Id = v
}

// GetFirstName returns the FirstName field value
func (o *AddressReferenceExplicit) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *AddressReferenceExplicit) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *AddressReferenceExplicit) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *AddressReferenceExplicit) SetLastName(v string) {
	o.LastName = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *AddressReferenceExplicit) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *AddressReferenceExplicit) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *AddressReferenceExplicit) SetCompany(v string) {
	o.Company = &v
}

// GetStreetAddress1 returns the StreetAddress1 field value
func (o *AddressReferenceExplicit) GetStreetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreetAddress1
}

// GetStreetAddress1Ok returns a tuple with the StreetAddress1 field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetStreetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreetAddress1, true
}

// SetStreetAddress1 sets field value
func (o *AddressReferenceExplicit) SetStreetAddress1(v string) {
	o.StreetAddress1 = v
}

// GetStreetAddress2 returns the StreetAddress2 field value if set, zero value otherwise.
func (o *AddressReferenceExplicit) GetStreetAddress2() string {
	if o == nil || IsNil(o.StreetAddress2) {
		var ret string
		return ret
	}
	return *o.StreetAddress2
}

// GetStreetAddress2Ok returns a tuple with the StreetAddress2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetStreetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress2) {
		return nil, false
	}
	return o.StreetAddress2, true
}

// HasStreetAddress2 returns a boolean if a field has been set.
func (o *AddressReferenceExplicit) HasStreetAddress2() bool {
	if o != nil && !IsNil(o.StreetAddress2) {
		return true
	}

	return false
}

// SetStreetAddress2 gets a reference to the given string and assigns it to the StreetAddress2 field.
func (o *AddressReferenceExplicit) SetStreetAddress2(v string) {
	o.StreetAddress2 = &v
}

// GetLocality returns the Locality field value
func (o *AddressReferenceExplicit) GetLocality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetLocalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locality, true
}

// SetLocality sets field value
func (o *AddressReferenceExplicit) SetLocality(v string) {
	o.Locality = v
}

// GetPostalCode returns the PostalCode field value
func (o *AddressReferenceExplicit) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *AddressReferenceExplicit) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AddressReferenceExplicit) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AddressReferenceExplicit) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AddressReferenceExplicit) SetRegion(v string) {
	o.Region = &v
}

// GetCountryCode returns the CountryCode field value
func (o *AddressReferenceExplicit) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *AddressReferenceExplicit) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddressReferenceExplicit) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddressReferenceExplicit) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AddressReferenceExplicit) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AddressReferenceExplicit) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AddressReferenceExplicit) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AddressReferenceExplicit) SetPhone(v string) {
	o.Phone = &v
}

// GetTag returns the Tag field value
func (o *AddressReferenceExplicit) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *AddressReferenceExplicit) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *AddressReferenceExplicit) SetTag(v string) {
	o.Tag = v
}

func (o AddressReferenceExplicit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressReferenceExplicit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["street_address1"] = o.StreetAddress1
	if !IsNil(o.StreetAddress2) {
		toSerialize["street_address2"] = o.StreetAddress2
	}
	toSerialize["locality"] = o.Locality
	toSerialize["postal_code"] = o.PostalCode
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["country_code"] = o.CountryCode
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	toSerialize[".tag"] = o.Tag
	return toSerialize, nil
}

type NullableAddressReferenceExplicit struct {
	value *AddressReferenceExplicit
	isSet bool
}

func (v NullableAddressReferenceExplicit) Get() *AddressReferenceExplicit {
	return v.value
}

func (v *NullableAddressReferenceExplicit) Set(val *AddressReferenceExplicit) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressReferenceExplicit) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressReferenceExplicit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressReferenceExplicit(val *AddressReferenceExplicit) *NullableAddressReferenceExplicit {
	return &NullableAddressReferenceExplicit{value: val, isSet: true}
}

func (v NullableAddressReferenceExplicit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressReferenceExplicit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


