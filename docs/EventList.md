# EventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** |  | 
**EventList** | **[]string** |  | 

## Methods

### NewEventList

`func NewEventList(tag string, eventList []string, ) *EventList`

NewEventList instantiates a new EventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventListWithDefaults

`func NewEventListWithDefaults() *EventList`

NewEventListWithDefaults instantiates a new EventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *EventList) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *EventList) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *EventList) SetTag(v string)`

SetTag sets Tag field to given value.


### GetEventList

`func (o *EventList) GetEventList() []string`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *EventList) GetEventListOk() (*[]string, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *EventList) SetEventList(v []string)`

SetEventList sets EventList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


