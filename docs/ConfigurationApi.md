# \ConfigurationAPI

All URIs are relative to *https://api.xwang.dev.bolt.me/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantCallbacksGet**](ConfigurationAPI.md#MerchantCallbacksGet) | **Get** /merchant/callbacks | Retrieve callback URLs for the merchant
[**MerchantCallbacksUpdate**](ConfigurationAPI.md#MerchantCallbacksUpdate) | **Patch** /merchant/callbacks | Update callback URLs for the merchant
[**MerchantIdentifiersGet**](ConfigurationAPI.md#MerchantIdentifiersGet) | **Get** /merchant/identifiers | Retrieve identifiers for the merchant



## MerchantCallbacksGet

> CallbackUrls MerchantCallbacksGet(ctx).XPublishableKey(xPublishableKey).Execute()

Retrieve callback URLs for the merchant



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
    resp, r, err := apiClient.ConfigurationAPI.MerchantCallbacksGet(context.Background()).XPublishableKey(xPublishableKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.MerchantCallbacksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MerchantCallbacksGet`: CallbackUrls
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.MerchantCallbacksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantCallbacksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 

### Return type

[**CallbackUrls**](CallbackUrls.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantCallbacksUpdate

> CallbackUrls MerchantCallbacksUpdate(ctx).XPublishableKey(xPublishableKey).CallbackUrls(callbackUrls).Execute()

Update callback URLs for the merchant



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
    callbackUrls := *openapiclient.NewCallbackUrls() // CallbackUrls | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationAPI.MerchantCallbacksUpdate(context.Background()).XPublishableKey(xPublishableKey).CallbackUrls(callbackUrls).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.MerchantCallbacksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MerchantCallbacksUpdate`: CallbackUrls
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.MerchantCallbacksUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantCallbacksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 
 **callbackUrls** | [**CallbackUrls**](CallbackUrls.md) |  | 

### Return type

[**CallbackUrls**](CallbackUrls.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantIdentifiersGet

> Identifiers MerchantIdentifiersGet(ctx).Execute()

Retrieve identifiers for the merchant



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
    resp, r, err := apiClient.ConfigurationAPI.MerchantIdentifiersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.MerchantIdentifiersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MerchantIdentifiersGet`: Identifiers
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.MerchantIdentifiersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantIdentifiersGetRequest struct via the builder pattern


### Return type

[**Identifiers**](Identifiers.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

