# PaymentMethodCreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | The credit card network. | 
**Bin** | **string** | The Bank Identification Number for the credit card. This is typically the first 4-6 digits of the credit card number. | 
**Last4** | **string** | The last 4 digits of the credit card number. | 
**Expiration** | **string** | The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE | 
**Token** | **string** | The Bolt token associated to the credit card. | 
**Tag** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**BillingAddress** | [**AddressReference**](AddressReference.md) |  | 
**BillingAddressId** | Pointer to **string** | The ID of credit card&#39;s billing address | [optional] [readonly] 

## Methods

### NewPaymentMethodCreditCard

`func NewPaymentMethodCreditCard(network string, bin string, last4 string, expiration string, token string, tag string, billingAddress AddressReference, ) *PaymentMethodCreditCard`

NewPaymentMethodCreditCard instantiates a new PaymentMethodCreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreditCardWithDefaults

`func NewPaymentMethodCreditCardWithDefaults() *PaymentMethodCreditCard`

NewPaymentMethodCreditCardWithDefaults instantiates a new PaymentMethodCreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *PaymentMethodCreditCard) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PaymentMethodCreditCard) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PaymentMethodCreditCard) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetBin

`func (o *PaymentMethodCreditCard) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *PaymentMethodCreditCard) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *PaymentMethodCreditCard) SetBin(v string)`

SetBin sets Bin field to given value.


### GetLast4

`func (o *PaymentMethodCreditCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethodCreditCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethodCreditCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetExpiration

`func (o *PaymentMethodCreditCard) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PaymentMethodCreditCard) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PaymentMethodCreditCard) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetToken

`func (o *PaymentMethodCreditCard) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentMethodCreditCard) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentMethodCreditCard) SetToken(v string)`

SetToken sets Token field to given value.


### GetTag

`func (o *PaymentMethodCreditCard) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PaymentMethodCreditCard) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PaymentMethodCreditCard) SetTag(v string)`

SetTag sets Tag field to given value.


### GetId

`func (o *PaymentMethodCreditCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodCreditCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodCreditCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodCreditCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBillingAddress

`func (o *PaymentMethodCreditCard) GetBillingAddress() AddressReference`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentMethodCreditCard) GetBillingAddressOk() (*AddressReference, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentMethodCreditCard) SetBillingAddress(v AddressReference)`

SetBillingAddress sets BillingAddress field to given value.


### GetBillingAddressId

`func (o *PaymentMethodCreditCard) GetBillingAddressId() string`

GetBillingAddressId returns the BillingAddressId field if non-nil, zero value otherwise.

### GetBillingAddressIdOk

`func (o *PaymentMethodCreditCard) GetBillingAddressIdOk() (*string, bool)`

GetBillingAddressIdOk returns a tuple with the BillingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressId

`func (o *PaymentMethodCreditCard) SetBillingAddressId(v string)`

SetBillingAddressId sets BillingAddressId field to given value.

### HasBillingAddressId

`func (o *PaymentMethodCreditCard) HasBillingAddressId() bool`

HasBillingAddressId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


