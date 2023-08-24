# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** |  | 
**EventGroup** | **string** |  | 
**EventList** | **[]string** |  | 

## Methods

### NewWebhookEvent

`func NewWebhookEvent(tag string, eventGroup string, eventList []string, ) *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *WebhookEvent) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *WebhookEvent) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *WebhookEvent) SetTag(v string)`

SetTag sets Tag field to given value.


### GetEventGroup

`func (o *WebhookEvent) GetEventGroup() string`

GetEventGroup returns the EventGroup field if non-nil, zero value otherwise.

### GetEventGroupOk

`func (o *WebhookEvent) GetEventGroupOk() (*string, bool)`

GetEventGroupOk returns a tuple with the EventGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventGroup

`func (o *WebhookEvent) SetEventGroup(v string)`

SetEventGroup sets EventGroup field to given value.


### GetEventList

`func (o *WebhookEvent) GetEventList() []string`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *WebhookEvent) GetEventListOk() (*[]string, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *WebhookEvent) SetEventList(v []string)`

SetEventList sets EventList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


