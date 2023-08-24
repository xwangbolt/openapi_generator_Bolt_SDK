# Cart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amounts** | [**Amounts**](Amounts.md) |  | 
**OrderReference** | **string** | This value is used by Bolt as an external reference to a given order. This reference must be unique per successful transaction. | 
**OrderDescription** | Pointer to **string** | Used optionally to pass additional information like order numbers or other IDs as needed. | [optional] 
**DisplayId** | Pointer to **string** | This field corresponds to the merchant&#39;s order reference associated with this Bolt transaction. | [optional] 
**Shipments** | Pointer to [**[]CartShipment**](CartShipment.md) |  | [optional] 
**Discounts** | Pointer to [**[]CartDiscount**](CartDiscount.md) |  | [optional] 
**Items** | Pointer to [**[]CartItem**](CartItem.md) |  | [optional] 

## Methods

### NewCart

`func NewCart(amounts Amounts, orderReference string, ) *Cart`

NewCart instantiates a new Cart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartWithDefaults

`func NewCartWithDefaults() *Cart`

NewCartWithDefaults instantiates a new Cart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmounts

`func (o *Cart) GetAmounts() Amounts`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *Cart) GetAmountsOk() (*Amounts, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *Cart) SetAmounts(v Amounts)`

SetAmounts sets Amounts field to given value.


### GetOrderReference

`func (o *Cart) GetOrderReference() string`

GetOrderReference returns the OrderReference field if non-nil, zero value otherwise.

### GetOrderReferenceOk

`func (o *Cart) GetOrderReferenceOk() (*string, bool)`

GetOrderReferenceOk returns a tuple with the OrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReference

`func (o *Cart) SetOrderReference(v string)`

SetOrderReference sets OrderReference field to given value.


### GetOrderDescription

`func (o *Cart) GetOrderDescription() string`

GetOrderDescription returns the OrderDescription field if non-nil, zero value otherwise.

### GetOrderDescriptionOk

`func (o *Cart) GetOrderDescriptionOk() (*string, bool)`

GetOrderDescriptionOk returns a tuple with the OrderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDescription

`func (o *Cart) SetOrderDescription(v string)`

SetOrderDescription sets OrderDescription field to given value.

### HasOrderDescription

`func (o *Cart) HasOrderDescription() bool`

HasOrderDescription returns a boolean if a field has been set.

### GetDisplayId

`func (o *Cart) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Cart) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Cart) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *Cart) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetShipments

`func (o *Cart) GetShipments() []CartShipment`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *Cart) GetShipmentsOk() (*[]CartShipment, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *Cart) SetShipments(v []CartShipment)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *Cart) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetDiscounts

`func (o *Cart) GetDiscounts() []CartDiscount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *Cart) GetDiscountsOk() (*[]CartDiscount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *Cart) SetDiscounts(v []CartDiscount)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *Cart) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetItems

`func (o *Cart) GetItems() []CartItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Cart) GetItemsOk() (*[]CartItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Cart) SetItems(v []CartItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Cart) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


