# \TestingApi

All URIs are relative to *https://api.xwang.dev.bolt.me/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestingAccountCreate**](TestingApi.md#TestingAccountCreate) | **Post** /testing/accounts | Create a test account
[**TestingCreditCardGet**](TestingApi.md#TestingCreditCardGet) | **Get** /testing/credit-cards | Retrieve a test credit card, including its token
[**TestingShipmentTrackingCreate**](TestingApi.md#TestingShipmentTrackingCreate) | **Post** /testing/shipments | Simulate a shipment tracking update



## TestingAccountCreate

> AccountTestCreationData TestingAccountCreate(ctx).AccountTestCreationData(accountTestCreationData).Execute()

Create a test account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountTestCreationData := *openapiclient.NewAccountTestCreationData("alice@example.com", "unverified", "+14155550199", "verified", "123456", "7GSjMRSHs6Ak7C_zvVW6P2IhZOHxMK7HZKW1fMX85ms", time.Now()) // AccountTestCreationData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestingApi.TestingAccountCreate(context.Background()).AccountTestCreationData(accountTestCreationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestingApi.TestingAccountCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestingAccountCreate`: AccountTestCreationData
    fmt.Fprintf(os.Stdout, "Response from `TestingApi.TestingAccountCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestingAccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountTestCreationData** | [**AccountTestCreationData**](AccountTestCreationData.md) |  | 

### Return type

[**AccountTestCreationData**](AccountTestCreationData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestingCreditCardGet

> CreditCard TestingCreditCardGet(ctx).Execute()

Retrieve a test credit card, including its token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestingApi.TestingCreditCardGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestingApi.TestingCreditCardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestingCreditCardGet`: CreditCard
    fmt.Fprintf(os.Stdout, "Response from `TestingApi.TestingCreditCardGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestingCreditCardGetRequest struct via the builder pattern


### Return type

[**CreditCard**](CreditCard.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestingShipmentTrackingCreate

> TestingShipmentTrackingCreate(ctx).ShipmentTrackingUpdate(shipmentTrackingUpdate).Execute()

Simulate a shipment tracking update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    shipmentTrackingUpdate := *openapiclient.NewShipmentTrackingUpdate("MockBolt-143292", []openapiclient.ShipmentTrackingUpdateTrackingDetailsInner{*openapiclient.NewShipmentTrackingUpdateTrackingDetailsInner()}, "in_transit") // ShipmentTrackingUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestingApi.TestingShipmentTrackingCreate(context.Background()).ShipmentTrackingUpdate(shipmentTrackingUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestingApi.TestingShipmentTrackingCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestingShipmentTrackingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentTrackingUpdate** | [**ShipmentTrackingUpdate**](ShipmentTrackingUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

