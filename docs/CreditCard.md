# CreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | The credit card network. | 
**Bin** | **string** | The Bank Identification Number for the credit card. This is typically the first 4-6 digits of the credit card number. | 
**Last4** | **string** | The last 4 digits of the credit card number. | 
**Expiration** | **string** | The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE | 
**Token** | **string** | The Bolt token associated to the credit card. | 

## Methods

### NewCreditCard

`func NewCreditCard(network string, bin string, last4 string, expiration string, token string, ) *CreditCard`

NewCreditCard instantiates a new CreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardWithDefaults

`func NewCreditCardWithDefaults() *CreditCard`

NewCreditCardWithDefaults instantiates a new CreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *CreditCard) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreditCard) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreditCard) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetBin

`func (o *CreditCard) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *CreditCard) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *CreditCard) SetBin(v string)`

SetBin sets Bin field to given value.


### GetLast4

`func (o *CreditCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *CreditCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *CreditCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetExpiration

`func (o *CreditCard) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CreditCard) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CreditCard) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetToken

`func (o *CreditCard) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreditCard) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreditCard) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


