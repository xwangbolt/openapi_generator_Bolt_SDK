# \PaymentsAPI

All URIs are relative to *https://api.xwang.dev.bolt.me/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GuestPaymentsInitialize**](PaymentsAPI.md#GuestPaymentsInitialize) | **Post** /guest/payments | Initialize a Bolt payment for guest shoppers
[**PaymentsInitialize**](PaymentsAPI.md#PaymentsInitialize) | **Post** /payments | Initialize a Bolt payment for logged in shoppers



## GuestPaymentsInitialize

> PaymentMethodInitializeResponse GuestPaymentsInitialize(ctx).XPublishableKey(xPublishableKey).GuestPaymentMethodInitializeRequest(guestPaymentMethodInitializeRequest).Execute()

Initialize a Bolt payment for guest shoppers



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
    xPublishableKey := "xPublishableKey_example" // string | The publicly viewable identifier used to identify a merchant division.
    guestPaymentMethodInitializeRequest := *openapiclient.NewGuestPaymentMethodInitializeRequest(*openapiclient.NewCart(*openapiclient.NewAmounts(int64(900), "USD"), "order_100"), openapiclient.guest_payment_method_initialize_request_payment_method{PaymentMethodPaypal: openapiclient.NewPaymentMethodPaypal("paypal", "www.example.com/handle_paypal_success", "www.example.com/handle_paypal_cancel")}) // GuestPaymentMethodInitializeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.GuestPaymentsInitialize(context.Background()).XPublishableKey(xPublishableKey).GuestPaymentMethodInitializeRequest(guestPaymentMethodInitializeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GuestPaymentsInitialize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuestPaymentsInitialize`: PaymentMethodInitializeResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GuestPaymentsInitialize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuestPaymentsInitializeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 
 **guestPaymentMethodInitializeRequest** | [**GuestPaymentMethodInitializeRequest**](GuestPaymentMethodInitializeRequest.md) |  | 

### Return type

[**PaymentMethodInitializeResponse**](PaymentMethodInitializeResponse.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsInitialize

> PaymentMethodInitializeResponse PaymentsInitialize(ctx).XPublishableKey(xPublishableKey).PaymentMethodInitializeRequest(paymentMethodInitializeRequest).Execute()

Initialize a Bolt payment for logged in shoppers



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
    xPublishableKey := "xPublishableKey_example" // string | The publicly viewable identifier used to identify a merchant division.
    paymentMethodInitializeRequest := *openapiclient.NewPaymentMethodInitializeRequest(*openapiclient.NewCart(*openapiclient.NewAmounts(int64(900), "USD"), "order_100"), openapiclient.payment_method_initialize_request_payment_method{PaymentMethodSavedPaymentMethod: openapiclient.NewPaymentMethodSavedPaymentMethod("saved_payment_method", "id")}) // PaymentMethodInitializeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.PaymentsInitialize(context.Background()).XPublishableKey(xPublishableKey).PaymentMethodInitializeRequest(paymentMethodInitializeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsInitialize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsInitialize`: PaymentMethodInitializeResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsInitialize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsInitializeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 
 **paymentMethodInitializeRequest** | [**PaymentMethodInitializeRequest**](PaymentMethodInitializeRequest.md) |  | 

### Return type

[**PaymentMethodInitializeResponse**](PaymentMethodInitializeResponse.md)

### Authorization

[api-key](../README.md#api-key), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

