# AccountAddressCreate400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | The type of error returned | 
**Message** | **string** | A human-readable error message, which might include information specific to the request that was made.  | 

## Methods

### NewAccountAddressCreate400Response

`func NewAccountAddressCreate400Response(tag string, message string, ) *AccountAddressCreate400Response`

NewAccountAddressCreate400Response instantiates a new AccountAddressCreate400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddressCreate400ResponseWithDefaults

`func NewAccountAddressCreate400ResponseWithDefaults() *AccountAddressCreate400Response`

NewAccountAddressCreate400ResponseWithDefaults instantiates a new AccountAddressCreate400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *AccountAddressCreate400Response) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AccountAddressCreate400Response) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AccountAddressCreate400Response) SetTag(v string)`

SetTag sets Tag field to given value.


### GetMessage

`func (o *AccountAddressCreate400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccountAddressCreate400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccountAddressCreate400Response) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


