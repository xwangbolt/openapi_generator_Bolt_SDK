# PaymentMethodCreditCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** | Credit card type | 
**BillingAddressInput** | Pointer to [**AddressReference**](AddressReference.md) |  | [optional] 
**BillingAddressId** | Pointer to **string** | The ID of credit card&#39;s billing address | [optional] 

## Methods

### NewPaymentMethodCreditCardAllOf

`func NewPaymentMethodCreditCardAllOf(tag string, type_ string, ) *PaymentMethodCreditCardAllOf`

NewPaymentMethodCreditCardAllOf instantiates a new PaymentMethodCreditCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreditCardAllOfWithDefaults

`func NewPaymentMethodCreditCardAllOfWithDefaults() *PaymentMethodCreditCardAllOf`

NewPaymentMethodCreditCardAllOfWithDefaults instantiates a new PaymentMethodCreditCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *PaymentMethodCreditCardAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PaymentMethodCreditCardAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PaymentMethodCreditCardAllOf) SetTag(v string)`

SetTag sets Tag field to given value.


### GetId

`func (o *PaymentMethodCreditCardAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodCreditCardAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodCreditCardAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodCreditCardAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethodCreditCardAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCreditCardAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCreditCardAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetBillingAddressInput

`func (o *PaymentMethodCreditCardAllOf) GetBillingAddressInput() AddressReference`

GetBillingAddressInput returns the BillingAddressInput field if non-nil, zero value otherwise.

### GetBillingAddressInputOk

`func (o *PaymentMethodCreditCardAllOf) GetBillingAddressInputOk() (*AddressReference, bool)`

GetBillingAddressInputOk returns a tuple with the BillingAddressInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressInput

`func (o *PaymentMethodCreditCardAllOf) SetBillingAddressInput(v AddressReference)`

SetBillingAddressInput sets BillingAddressInput field to given value.

### HasBillingAddressInput

`func (o *PaymentMethodCreditCardAllOf) HasBillingAddressInput() bool`

HasBillingAddressInput returns a boolean if a field has been set.

### GetBillingAddressId

`func (o *PaymentMethodCreditCardAllOf) GetBillingAddressId() string`

GetBillingAddressId returns the BillingAddressId field if non-nil, zero value otherwise.

### GetBillingAddressIdOk

`func (o *PaymentMethodCreditCardAllOf) GetBillingAddressIdOk() (*string, bool)`

GetBillingAddressIdOk returns a tuple with the BillingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressId

`func (o *PaymentMethodCreditCardAllOf) SetBillingAddressId(v string)`

SetBillingAddressId sets BillingAddressId field to given value.

### HasBillingAddressId

`func (o *PaymentMethodCreditCardAllOf) HasBillingAddressId() bool`

HasBillingAddressId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


