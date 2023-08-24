# CartDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amounts** | [**Amounts**](Amounts.md) |  | 
**Code** | Pointer to **string** | Discount code. | [optional] 
**DetailsUrl** | Pointer to **string** | Used to provide a link to additional details, such as a landing page, associated with the discount offering. | [optional] 

## Methods

### NewCartDiscount

`func NewCartDiscount(amounts Amounts, ) *CartDiscount`

NewCartDiscount instantiates a new CartDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartDiscountWithDefaults

`func NewCartDiscountWithDefaults() *CartDiscount`

NewCartDiscountWithDefaults instantiates a new CartDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmounts

`func (o *CartDiscount) GetAmounts() Amounts`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *CartDiscount) GetAmountsOk() (*Amounts, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *CartDiscount) SetAmounts(v Amounts)`

SetAmounts sets Amounts field to given value.


### GetCode

`func (o *CartDiscount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CartDiscount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CartDiscount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CartDiscount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetailsUrl

`func (o *CartDiscount) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *CartDiscount) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *CartDiscount) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *CartDiscount) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


