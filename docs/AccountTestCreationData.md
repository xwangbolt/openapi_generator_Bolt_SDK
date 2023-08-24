# AccountTestCreationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | [readonly] 
**EmailState** | **string** |  | 
**Phone** | **string** |  | [readonly] 
**PhoneState** | **string** |  | 
**IsMigrated** | Pointer to **bool** |  | [optional] 
**HasAddress** | Pointer to **bool** |  | [optional] 
**OtpCode** | **string** |  | [readonly] 
**OauthCode** | **string** |  | [readonly] 
**DeactivateAt** | **time.Time** |  | 

## Methods

### NewAccountTestCreationData

`func NewAccountTestCreationData(email string, emailState string, phone string, phoneState string, otpCode string, oauthCode string, deactivateAt time.Time, ) *AccountTestCreationData`

NewAccountTestCreationData instantiates a new AccountTestCreationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTestCreationDataWithDefaults

`func NewAccountTestCreationDataWithDefaults() *AccountTestCreationData`

NewAccountTestCreationDataWithDefaults instantiates a new AccountTestCreationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountTestCreationData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountTestCreationData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountTestCreationData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailState

`func (o *AccountTestCreationData) GetEmailState() string`

GetEmailState returns the EmailState field if non-nil, zero value otherwise.

### GetEmailStateOk

`func (o *AccountTestCreationData) GetEmailStateOk() (*string, bool)`

GetEmailStateOk returns a tuple with the EmailState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailState

`func (o *AccountTestCreationData) SetEmailState(v string)`

SetEmailState sets EmailState field to given value.


### GetPhone

`func (o *AccountTestCreationData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AccountTestCreationData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AccountTestCreationData) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPhoneState

`func (o *AccountTestCreationData) GetPhoneState() string`

GetPhoneState returns the PhoneState field if non-nil, zero value otherwise.

### GetPhoneStateOk

`func (o *AccountTestCreationData) GetPhoneStateOk() (*string, bool)`

GetPhoneStateOk returns a tuple with the PhoneState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneState

`func (o *AccountTestCreationData) SetPhoneState(v string)`

SetPhoneState sets PhoneState field to given value.


### GetIsMigrated

`func (o *AccountTestCreationData) GetIsMigrated() bool`

GetIsMigrated returns the IsMigrated field if non-nil, zero value otherwise.

### GetIsMigratedOk

`func (o *AccountTestCreationData) GetIsMigratedOk() (*bool, bool)`

GetIsMigratedOk returns a tuple with the IsMigrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMigrated

`func (o *AccountTestCreationData) SetIsMigrated(v bool)`

SetIsMigrated sets IsMigrated field to given value.

### HasIsMigrated

`func (o *AccountTestCreationData) HasIsMigrated() bool`

HasIsMigrated returns a boolean if a field has been set.

### GetHasAddress

`func (o *AccountTestCreationData) GetHasAddress() bool`

GetHasAddress returns the HasAddress field if non-nil, zero value otherwise.

### GetHasAddressOk

`func (o *AccountTestCreationData) GetHasAddressOk() (*bool, bool)`

GetHasAddressOk returns a tuple with the HasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAddress

`func (o *AccountTestCreationData) SetHasAddress(v bool)`

SetHasAddress sets HasAddress field to given value.

### HasHasAddress

`func (o *AccountTestCreationData) HasHasAddress() bool`

HasHasAddress returns a boolean if a field has been set.

### GetOtpCode

`func (o *AccountTestCreationData) GetOtpCode() string`

GetOtpCode returns the OtpCode field if non-nil, zero value otherwise.

### GetOtpCodeOk

`func (o *AccountTestCreationData) GetOtpCodeOk() (*string, bool)`

GetOtpCodeOk returns a tuple with the OtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpCode

`func (o *AccountTestCreationData) SetOtpCode(v string)`

SetOtpCode sets OtpCode field to given value.


### GetOauthCode

`func (o *AccountTestCreationData) GetOauthCode() string`

GetOauthCode returns the OauthCode field if non-nil, zero value otherwise.

### GetOauthCodeOk

`func (o *AccountTestCreationData) GetOauthCodeOk() (*string, bool)`

GetOauthCodeOk returns a tuple with the OauthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthCode

`func (o *AccountTestCreationData) SetOauthCode(v string)`

SetOauthCode sets OauthCode field to given value.


### GetDeactivateAt

`func (o *AccountTestCreationData) GetDeactivateAt() time.Time`

GetDeactivateAt returns the DeactivateAt field if non-nil, zero value otherwise.

### GetDeactivateAtOk

`func (o *AccountTestCreationData) GetDeactivateAtOk() (*time.Time, bool)`

GetDeactivateAtOk returns a tuple with the DeactivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivateAt

`func (o *AccountTestCreationData) SetDeactivateAt(v time.Time)`

SetDeactivateAt sets DeactivateAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


