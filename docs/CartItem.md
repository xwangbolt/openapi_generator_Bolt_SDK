# CartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of a given item. | 
**Reference** | **string** | This value is used by Bolt as an external reference to a given item. | 
**Description** | Pointer to **string** | A human-readable description of this cart item. | [optional] 
**TotalAmount** | **int64** | The total amount, in cents, of the item including its taxes if applicable. | 
**UnitPrice** | **int64** | The price of one unit of the item; for example, the price of one pack of socks. | 
**Quantity** | **int64** | The number of units that comprise this cart item. | 
**ImageUrl** | Pointer to **string** | Used to provide a link to the image associated with the item. | [optional] 

## Methods

### NewCartItem

`func NewCartItem(name string, reference string, totalAmount int64, unitPrice int64, quantity int64, ) *CartItem`

NewCartItem instantiates a new CartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemWithDefaults

`func NewCartItemWithDefaults() *CartItem`

NewCartItemWithDefaults instantiates a new CartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CartItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartItem) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *CartItem) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CartItem) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CartItem) SetReference(v string)`

SetReference sets Reference field to given value.


### GetDescription

`func (o *CartItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CartItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CartItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CartItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTotalAmount

`func (o *CartItem) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CartItem) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CartItem) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.


### GetUnitPrice

`func (o *CartItem) GetUnitPrice() int64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CartItem) GetUnitPriceOk() (*int64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CartItem) SetUnitPrice(v int64)`

SetUnitPrice sets UnitPrice field to given value.


### GetQuantity

`func (o *CartItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetImageUrl

`func (o *CartItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CartItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CartItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CartItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


