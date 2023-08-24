# PaymentMethodInitializeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cart** | [**Cart**](Cart.md) |  | 
**PaymentMethod** | [**PaymentMethodInitializeRequestPaymentMethod**](PaymentMethodInitializeRequestPaymentMethod.md) |  | 

## Methods

### NewPaymentMethodInitializeRequest

`func NewPaymentMethodInitializeRequest(cart Cart, paymentMethod PaymentMethodInitializeRequestPaymentMethod, ) *PaymentMethodInitializeRequest`

NewPaymentMethodInitializeRequest instantiates a new PaymentMethodInitializeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodInitializeRequestWithDefaults

`func NewPaymentMethodInitializeRequestWithDefaults() *PaymentMethodInitializeRequest`

NewPaymentMethodInitializeRequestWithDefaults instantiates a new PaymentMethodInitializeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCart

`func (o *PaymentMethodInitializeRequest) GetCart() Cart`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *PaymentMethodInitializeRequest) GetCartOk() (*Cart, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *PaymentMethodInitializeRequest) SetCart(v Cart)`

SetCart sets Cart field to given value.


### GetPaymentMethod

`func (o *PaymentMethodInitializeRequest) GetPaymentMethod() PaymentMethodInitializeRequestPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentMethodInitializeRequest) GetPaymentMethodOk() (*PaymentMethodInitializeRequestPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentMethodInitializeRequest) SetPaymentMethod(v PaymentMethodInitializeRequestPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


