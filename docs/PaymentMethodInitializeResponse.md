# PaymentMethodInitializeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**PaymentMethodInitializeResponseAction**](PaymentMethodInitializeResponseAction.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodInitializeResponse

`func NewPaymentMethodInitializeResponse() *PaymentMethodInitializeResponse`

NewPaymentMethodInitializeResponse instantiates a new PaymentMethodInitializeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodInitializeResponseWithDefaults

`func NewPaymentMethodInitializeResponseWithDefaults() *PaymentMethodInitializeResponse`

NewPaymentMethodInitializeResponseWithDefaults instantiates a new PaymentMethodInitializeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PaymentMethodInitializeResponse) GetAction() PaymentMethodInitializeResponseAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PaymentMethodInitializeResponse) GetActionOk() (*PaymentMethodInitializeResponseAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PaymentMethodInitializeResponse) SetAction(v PaymentMethodInitializeResponseAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PaymentMethodInitializeResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethodInitializeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodInitializeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodInitializeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodInitializeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentMethodInitializeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentMethodInitializeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentMethodInitializeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentMethodInitializeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


