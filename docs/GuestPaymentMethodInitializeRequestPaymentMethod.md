# GuestPaymentMethodInitializeRequestPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** |  | 
**Success** | **string** | Redirect URL for successful PayPal transaction. | 
**Cancel** | **string** | Redirect URL for canceled PayPal transaction. | 

## Methods

### NewGuestPaymentMethodInitializeRequestPaymentMethod

`func NewGuestPaymentMethodInitializeRequestPaymentMethod(tag string, success string, cancel string, ) *GuestPaymentMethodInitializeRequestPaymentMethod`

NewGuestPaymentMethodInitializeRequestPaymentMethod instantiates a new GuestPaymentMethodInitializeRequestPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestPaymentMethodInitializeRequestPaymentMethodWithDefaults

`func NewGuestPaymentMethodInitializeRequestPaymentMethodWithDefaults() *GuestPaymentMethodInitializeRequestPaymentMethod`

NewGuestPaymentMethodInitializeRequestPaymentMethodWithDefaults instantiates a new GuestPaymentMethodInitializeRequestPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) SetTag(v string)`

SetTag sets Tag field to given value.


### GetSuccess

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) GetSuccess() string`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) GetSuccessOk() (*string, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) SetSuccess(v string)`

SetSuccess sets Success field to given value.


### GetCancel

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) GetCancel() string`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) GetCancelOk() (*string, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *GuestPaymentMethodInitializeRequestPaymentMethod) SetCancel(v string)`

SetCancel sets Cancel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


