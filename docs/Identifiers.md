# Identifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** |  | 
**MerchantDivisions** | [**[]IdentifiersMerchantDivisionsInner**](IdentifiersMerchantDivisionsInner.md) |  | 
**SigningSecret** | **string** |  | 

## Methods

### NewIdentifiers

`func NewIdentifiers(merchantId string, merchantDivisions []IdentifiersMerchantDivisionsInner, signingSecret string, ) *Identifiers`

NewIdentifiers instantiates a new Identifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifiersWithDefaults

`func NewIdentifiersWithDefaults() *Identifiers`

NewIdentifiersWithDefaults instantiates a new Identifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *Identifiers) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *Identifiers) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *Identifiers) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantDivisions

`func (o *Identifiers) GetMerchantDivisions() []IdentifiersMerchantDivisionsInner`

GetMerchantDivisions returns the MerchantDivisions field if non-nil, zero value otherwise.

### GetMerchantDivisionsOk

`func (o *Identifiers) GetMerchantDivisionsOk() (*[]IdentifiersMerchantDivisionsInner, bool)`

GetMerchantDivisionsOk returns a tuple with the MerchantDivisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantDivisions

`func (o *Identifiers) SetMerchantDivisions(v []IdentifiersMerchantDivisionsInner)`

SetMerchantDivisions sets MerchantDivisions field to given value.


### GetSigningSecret

`func (o *Identifiers) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *Identifiers) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *Identifiers) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


