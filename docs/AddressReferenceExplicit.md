# AddressReferenceExplicit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | The type of address reference | 
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

## Methods

### NewAddressReferenceExplicit

`func NewAddressReferenceExplicit(tag string, firstName string, lastName string, streetAddress1 string, locality string, postalCode string, countryCode string, ) *AddressReferenceExplicit`

NewAddressReferenceExplicit instantiates a new AddressReferenceExplicit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressReferenceExplicitWithDefaults

`func NewAddressReferenceExplicitWithDefaults() *AddressReferenceExplicit`

NewAddressReferenceExplicitWithDefaults instantiates a new AddressReferenceExplicit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *AddressReferenceExplicit) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AddressReferenceExplicit) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AddressReferenceExplicit) SetTag(v string)`

SetTag sets Tag field to given value.


### GetId

`func (o *AddressReferenceExplicit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressReferenceExplicit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressReferenceExplicit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddressReferenceExplicit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *AddressReferenceExplicit) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddressReferenceExplicit) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddressReferenceExplicit) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *AddressReferenceExplicit) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddressReferenceExplicit) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddressReferenceExplicit) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetCompany

`func (o *AddressReferenceExplicit) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressReferenceExplicit) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressReferenceExplicit) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressReferenceExplicit) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetStreetAddress1

`func (o *AddressReferenceExplicit) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *AddressReferenceExplicit) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *AddressReferenceExplicit) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.


### GetStreetAddress2

`func (o *AddressReferenceExplicit) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *AddressReferenceExplicit) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *AddressReferenceExplicit) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *AddressReferenceExplicit) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### GetLocality

`func (o *AddressReferenceExplicit) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AddressReferenceExplicit) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AddressReferenceExplicit) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetPostalCode

`func (o *AddressReferenceExplicit) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressReferenceExplicit) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressReferenceExplicit) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetRegion

`func (o *AddressReferenceExplicit) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddressReferenceExplicit) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddressReferenceExplicit) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddressReferenceExplicit) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressReferenceExplicit) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressReferenceExplicit) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressReferenceExplicit) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *AddressReferenceExplicit) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressReferenceExplicit) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressReferenceExplicit) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressReferenceExplicit) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *AddressReferenceExplicit) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressReferenceExplicit) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressReferenceExplicit) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressReferenceExplicit) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


