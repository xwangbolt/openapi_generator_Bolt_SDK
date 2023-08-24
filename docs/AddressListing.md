# AddressListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Company** | Pointer to **string** |  | [optional] 
**StreetAddress1** | **string** |  | 
**StreetAddress2** | Pointer to **string** |  | [optional] 
**Locality** | **string** |  | 
**PostalCode** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**CountryCode** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddressListing

`func NewAddressListing(firstName string, lastName string, streetAddress1 string, locality string, postalCode string, countryCode string, ) *AddressListing`

NewAddressListing instantiates a new AddressListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressListingWithDefaults

`func NewAddressListingWithDefaults() *AddressListing`

NewAddressListingWithDefaults instantiates a new AddressListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddressListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *AddressListing) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddressListing) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddressListing) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *AddressListing) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddressListing) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddressListing) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetCompany

`func (o *AddressListing) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressListing) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressListing) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressListing) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetStreetAddress1

`func (o *AddressListing) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *AddressListing) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *AddressListing) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.


### GetStreetAddress2

`func (o *AddressListing) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *AddressListing) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *AddressListing) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *AddressListing) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### GetLocality

`func (o *AddressListing) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AddressListing) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AddressListing) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetPostalCode

`func (o *AddressListing) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressListing) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressListing) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetRegion

`func (o *AddressListing) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddressListing) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddressListing) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddressListing) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressListing) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressListing) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressListing) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *AddressListing) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressListing) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressListing) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressListing) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *AddressListing) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressListing) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressListing) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressListing) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetIsDefault

`func (o *AddressListing) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AddressListing) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AddressListing) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AddressListing) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


