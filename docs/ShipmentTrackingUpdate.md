# ShipmentTrackingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingNumber** | **string** | The carrier&#39;s tracking number for the shipment. Must be prefixed with &#x60;MockBolt&#x60;. | 
**TrackingDetails** | [**[]ShipmentTrackingUpdateTrackingDetailsInner**](ShipmentTrackingUpdateTrackingDetailsInner.md) | A list of tracking updates that contain the shipment&#39;s status, location, and any unique messages. | 
**Status** | **string** | The shipment&#39;s status. | 
**DeliveryDate** | Pointer to **time.Time** | The shipment&#39;s actual or estimated delivery date. | [optional] 

## Methods

### NewShipmentTrackingUpdate

`func NewShipmentTrackingUpdate(trackingNumber string, trackingDetails []ShipmentTrackingUpdateTrackingDetailsInner, status string, ) *ShipmentTrackingUpdate`

NewShipmentTrackingUpdate instantiates a new ShipmentTrackingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentTrackingUpdateWithDefaults

`func NewShipmentTrackingUpdateWithDefaults() *ShipmentTrackingUpdate`

NewShipmentTrackingUpdateWithDefaults instantiates a new ShipmentTrackingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingNumber

`func (o *ShipmentTrackingUpdate) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ShipmentTrackingUpdate) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ShipmentTrackingUpdate) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.


### GetTrackingDetails

`func (o *ShipmentTrackingUpdate) GetTrackingDetails() []ShipmentTrackingUpdateTrackingDetailsInner`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *ShipmentTrackingUpdate) GetTrackingDetailsOk() (*[]ShipmentTrackingUpdateTrackingDetailsInner, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *ShipmentTrackingUpdate) SetTrackingDetails(v []ShipmentTrackingUpdateTrackingDetailsInner)`

SetTrackingDetails sets TrackingDetails field to given value.


### GetStatus

`func (o *ShipmentTrackingUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipmentTrackingUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipmentTrackingUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDeliveryDate

`func (o *ShipmentTrackingUpdate) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ShipmentTrackingUpdate) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ShipmentTrackingUpdate) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ShipmentTrackingUpdate) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


