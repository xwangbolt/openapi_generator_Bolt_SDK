# Amounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | The total amount, in cents, including its items and taxes (if applicable). | 
**Currency** | **string** |  | 
**Tax** | Pointer to **int64** | The total tax amount, in cents for all of the items associated with the cart. | [optional] 

## Methods

### NewAmounts

`func NewAmounts(total int64, currency string, ) *Amounts`

NewAmounts instantiates a new Amounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountsWithDefaults

`func NewAmountsWithDefaults() *Amounts`

NewAmountsWithDefaults instantiates a new Amounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *Amounts) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Amounts) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Amounts) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCurrency

`func (o *Amounts) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Amounts) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Amounts) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTax

`func (o *Amounts) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *Amounts) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *Amounts) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *Amounts) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


