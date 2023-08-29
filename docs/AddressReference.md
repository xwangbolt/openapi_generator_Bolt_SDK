# AddressReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | The type of address reference | 
**Id** | **string** |  | [readonly] 
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

### NewAddressReference

`func NewAddressReference(tag string, id string, firstName string, lastName string, streetAddress1 string, locality string, postalCode string, countryCode string, ) *AddressReference`

NewAddressReference instantiates a new AddressReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressReferenceWithDefaults

`func NewAddressReferenceWithDefaults() *AddressReference`

NewAddressReferenceWithDefaults instantiates a new AddressReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *AddressReference) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AddressReference) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AddressReference) SetTag(v string)`

SetTag sets Tag field to given value.


### GetId

`func (o *AddressReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressReference) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *AddressReference) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddressReference) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddressReference) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *AddressReference) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddressReference) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddressReference) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetCompany

`func (o *AddressReference) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressReference) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressReference) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressReference) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetStreetAddress1

`func (o *AddressReference) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *AddressReference) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *AddressReference) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.


### GetStreetAddress2

`func (o *AddressReference) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *AddressReference) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *AddressReference) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *AddressReference) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### GetLocality

`func (o *AddressReference) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AddressReference) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AddressReference) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetPostalCode

`func (o *AddressReference) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressReference) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressReference) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetRegion

`func (o *AddressReference) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddressReference) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddressReference) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddressReference) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressReference) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressReference) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressReference) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *AddressReference) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressReference) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressReference) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressReference) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *AddressReference) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressReference) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressReference) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressReference) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


