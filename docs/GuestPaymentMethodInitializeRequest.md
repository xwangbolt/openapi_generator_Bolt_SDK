# GuestPaymentMethodInitializeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cart** | [**Cart**](Cart.md) |  | 
**PaymentMethod** | [**GuestPaymentMethodInitializeRequestPaymentMethod**](GuestPaymentMethodInitializeRequestPaymentMethod.md) |  | 

## Methods

### NewGuestPaymentMethodInitializeRequest

`func NewGuestPaymentMethodInitializeRequest(cart Cart, paymentMethod GuestPaymentMethodInitializeRequestPaymentMethod, ) *GuestPaymentMethodInitializeRequest`

NewGuestPaymentMethodInitializeRequest instantiates a new GuestPaymentMethodInitializeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestPaymentMethodInitializeRequestWithDefaults

`func NewGuestPaymentMethodInitializeRequestWithDefaults() *GuestPaymentMethodInitializeRequest`

NewGuestPaymentMethodInitializeRequestWithDefaults instantiates a new GuestPaymentMethodInitializeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCart

`func (o *GuestPaymentMethodInitializeRequest) GetCart() Cart`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *GuestPaymentMethodInitializeRequest) GetCartOk() (*Cart, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *GuestPaymentMethodInitializeRequest) SetCart(v Cart)`

SetCart sets Cart field to given value.


### GetPaymentMethod

`func (o *GuestPaymentMethodInitializeRequest) GetPaymentMethod() GuestPaymentMethodInitializeRequestPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GuestPaymentMethodInitializeRequest) GetPaymentMethodOk() (*GuestPaymentMethodInitializeRequestPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GuestPaymentMethodInitializeRequest) SetPaymentMethod(v GuestPaymentMethodInitializeRequestPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


