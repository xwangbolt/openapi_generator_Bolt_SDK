# CartShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**AddressReference**](AddressReference.md) |  | [optional] 
**Cost** | Pointer to [**Amounts**](Amounts.md) |  | [optional] 
**Carrier** | Pointer to **string** | The name of the carrier selected. | [optional] 

## Methods

### NewCartShipment

`func NewCartShipment() *CartShipment`

NewCartShipment instantiates a new CartShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartShipmentWithDefaults

`func NewCartShipmentWithDefaults() *CartShipment`

NewCartShipmentWithDefaults instantiates a new CartShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CartShipment) GetAddress() AddressReference`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CartShipment) GetAddressOk() (*AddressReference, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CartShipment) SetAddress(v AddressReference)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CartShipment) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCost

`func (o *CartShipment) GetCost() Amounts`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CartShipment) GetCostOk() (*Amounts, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CartShipment) SetCost(v Amounts)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CartShipment) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCarrier

`func (o *CartShipment) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *CartShipment) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *CartShipment) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *CartShipment) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


