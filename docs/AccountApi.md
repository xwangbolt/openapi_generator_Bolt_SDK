# \AccountApi

All URIs are relative to *https://api.xwang.dev.bolt.me/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountAddPaymentMethod**](AccountApi.md#AccountAddPaymentMethod) | **Post** /account/payment-methods | Add a payment method to a shopper&#39;s Bolt account Wallet.
[**AccountAddressCreate**](AccountApi.md#AccountAddressCreate) | **Post** /account/addresses | Add an address
[**AccountAddressDelete**](AccountApi.md#AccountAddressDelete) | **Delete** /account/addresses/{id} | Delete an existing address
[**AccountAddressEdit**](AccountApi.md#AccountAddressEdit) | **Put** /account/addresses/{id} | Edit an existing address
[**AccountExists**](AccountApi.md#AccountExists) | **Get** /account/exists | Determine the existence of a Bolt account
[**AccountGet**](AccountApi.md#AccountGet) | **Get** /account | Retrieve account details



## AccountAddPaymentMethod

> PaymentMethod AccountAddPaymentMethod(ctx).XPublishableKey(xPublishableKey).PaymentMethod(paymentMethod).Execute()

Add a payment method to a shopper's Bolt account Wallet.



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
    paymentMethod := openapiclient.payment-method{PaymentMethodCreditCard: openapiclient.NewPaymentMethodCreditCard("credit_card", "credit", "visa", "411111", "1004", "2025-03", "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0")} // PaymentMethod | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.AccountAddPaymentMethod(context.Background()).XPublishableKey(xPublishableKey).PaymentMethod(paymentMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountAddPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountAddPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountAddPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountAddPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 
 **paymentMethod** | [**PaymentMethod**](PaymentMethod.md) |  | 

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[api-key](../README.md#api-key), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountAddressCreate

> AddressListing AccountAddressCreate(ctx).XPublishableKey(xPublishableKey).AddressListing(addressListing).Execute()

Add an address



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
    addressListing := *openapiclient.NewAddressListing("Alice", "Baker", "535 Mission St, Ste 1401", "San Francisco", "94105", "US") // AddressListing | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.AccountAddressCreate(context.Background()).XPublishableKey(xPublishableKey).AddressListing(addressListing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountAddressCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountAddressCreate`: AddressListing
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountAddressCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountAddressCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 
 **addressListing** | [**AddressListing**](AddressListing.md) |  | 

### Return type

[**AddressListing**](AddressListing.md)

### Authorization

[api-key](../README.md#api-key), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountAddressDelete

> AccountAddressDelete(ctx, id).XPublishableKey(xPublishableKey).Execute()

Delete an existing address



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
    id := "D4g3h5tBuVYK9" // string | The ID of the address to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountApi.AccountAddressDelete(context.Background(), id).XPublishableKey(xPublishableKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountAddressDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the address to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountAddressDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountAddressEdit

> AddressListing AccountAddressEdit(ctx, id).XPublishableKey(xPublishableKey).AddressListing(addressListing).Execute()

Edit an existing address



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
    id := "D4g3h5tBuVYK9" // string | The ID of the address to edit
    addressListing := *openapiclient.NewAddressListing("Alice", "Baker", "535 Mission St, Ste 1401", "San Francisco", "94105", "US") // AddressListing | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.AccountAddressEdit(context.Background(), id).XPublishableKey(xPublishableKey).AddressListing(addressListing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountAddressEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountAddressEdit`: AddressListing
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountAddressEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the address to edit | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountAddressEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 

 **addressListing** | [**AddressListing**](AddressListing.md) |  | 

### Return type

[**AddressListing**](AddressListing.md)

### Authorization

[api-key](../README.md#api-key), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountExists

> AccountExists(ctx).XPublishableKey(xPublishableKey).Identifier(identifier).Execute()

Determine the existence of a Bolt account



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
    identifier := map[string][]openapiclient.AccountExistsIdentifierParameter{ ... } // AccountExistsIdentifierParameter | A type and value combination that defines the identifier used to detect the existence of an account. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountApi.AccountExists(context.Background()).XPublishableKey(xPublishableKey).Identifier(identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 
 **identifier** | [**AccountExistsIdentifierParameter**](AccountExistsIdentifierParameter.md) | A type and value combination that defines the identifier used to detect the existence of an account.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGet

> Account AccountGet(ctx).XPublishableKey(xPublishableKey).Execute()

Retrieve account details



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.AccountGet(context.Background()).XPublishableKey(xPublishableKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGet`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 

### Return type

[**Account**](Account.md)

### Authorization

[api-key](../README.md#api-key), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

