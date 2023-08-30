# AccountExistsIdentifierParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierType** | **string** | The type of identifier | 
**IdentifierValue** | **string** | The value of the identifier. The value must be valid for the specified &#x60;identifier_type&#x60; | 

## Methods

### NewAccountExistsIdentifierParameter

`func NewAccountExistsIdentifierParameter(identifierType string, identifierValue string, ) *AccountExistsIdentifierParameter`

NewAccountExistsIdentifierParameter instantiates a new AccountExistsIdentifierParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountExistsIdentifierParameterWithDefaults

`func NewAccountExistsIdentifierParameterWithDefaults() *AccountExistsIdentifierParameter`

NewAccountExistsIdentifierParameterWithDefaults instantiates a new AccountExistsIdentifierParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierType

`func (o *AccountExistsIdentifierParameter) GetIdentifierType() string`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *AccountExistsIdentifierParameter) GetIdentifierTypeOk() (*string, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *AccountExistsIdentifierParameter) SetIdentifierType(v string)`

SetIdentifierType sets IdentifierType field to given value.


### GetIdentifierValue

`func (o *AccountExistsIdentifierParameter) GetIdentifierValue() string`

GetIdentifierValue returns the IdentifierValue field if non-nil, zero value otherwise.

### GetIdentifierValueOk

`func (o *AccountExistsIdentifierParameter) GetIdentifierValueOk() (*string, bool)`

GetIdentifierValueOk returns a tuple with the IdentifierValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierValue

`func (o *AccountExistsIdentifierParameter) SetIdentifierValue(v string)`

SetIdentifierValue sets IdentifierValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


