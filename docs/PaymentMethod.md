# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**BillingAddress** | [**AddressReference**](AddressReference.md) |  | 
**BillingAddressId** | Pointer to **string** | The ID of credit card&#39;s billing address | [optional] [readonly] 
**Network** | **string** | The credit card network. | 
**Bin** | **string** | The Bank Identification Number for the credit card. This is typically the first 4-6 digits of the credit card number. | 
**Last4** | **string** | The last 4 digits of the credit card number. | 
**Expiration** | **string** | The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE | 
**Token** | **string** | The Bolt token associated to the credit card. | 

## Methods

### NewPaymentMethod

`func NewPaymentMethod(tag string, billingAddress AddressReference, network string, bin string, last4 string, expiration string, token string, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *PaymentMethod) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PaymentMethod) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PaymentMethod) SetTag(v string)`

SetTag sets Tag field to given value.


### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBillingAddress

`func (o *PaymentMethod) GetBillingAddress() AddressReference`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentMethod) GetBillingAddressOk() (*AddressReference, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentMethod) SetBillingAddress(v AddressReference)`

SetBillingAddress sets BillingAddress field to given value.


### GetBillingAddressId

`func (o *PaymentMethod) GetBillingAddressId() string`

GetBillingAddressId returns the BillingAddressId field if non-nil, zero value otherwise.

### GetBillingAddressIdOk

`func (o *PaymentMethod) GetBillingAddressIdOk() (*string, bool)`

GetBillingAddressIdOk returns a tuple with the BillingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressId

`func (o *PaymentMethod) SetBillingAddressId(v string)`

SetBillingAddressId sets BillingAddressId field to given value.

### HasBillingAddressId

`func (o *PaymentMethod) HasBillingAddressId() bool`

HasBillingAddressId returns a boolean if a field has been set.

### GetNetwork

`func (o *PaymentMethod) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PaymentMethod) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PaymentMethod) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetBin

`func (o *PaymentMethod) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *PaymentMethod) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *PaymentMethod) SetBin(v string)`

SetBin sets Bin field to given value.


### GetLast4

`func (o *PaymentMethod) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethod) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethod) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetExpiration

`func (o *PaymentMethod) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PaymentMethod) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PaymentMethod) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetToken

`func (o *PaymentMethod) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentMethod) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentMethod) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


