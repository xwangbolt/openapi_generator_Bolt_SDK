# AccountExistsIdentifierParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of identifier | 
**Value** | **string** | The value of the identifier. The value must be valid for the specified &#x60;identifier_type&#x60; | 

## Methods

### NewAccountExistsIdentifierParameter

`func NewAccountExistsIdentifierParameter(type_ string, value string, ) *AccountExistsIdentifierParameter`

NewAccountExistsIdentifierParameter instantiates a new AccountExistsIdentifierParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountExistsIdentifierParameterWithDefaults

`func NewAccountExistsIdentifierParameterWithDefaults() *AccountExistsIdentifierParameter`

NewAccountExistsIdentifierParameterWithDefaults instantiates a new AccountExistsIdentifierParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountExistsIdentifierParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountExistsIdentifierParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountExistsIdentifierParameter) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *AccountExistsIdentifierParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AccountExistsIdentifierParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AccountExistsIdentifierParameter) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


