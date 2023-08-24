# \WebhooksApi

All URIs are relative to *https://api.xwang.dev.bolt.me/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksCreate**](WebhooksApi.md#WebhooksCreate) | **Put** /webhooks | Create a webhook to subscribe to certain events
[**WebhooksDelete**](WebhooksApi.md#WebhooksDelete) | **Delete** /webhooks/{id} | Delete an existing webhook
[**WebhooksGet**](WebhooksApi.md#WebhooksGet) | **Get** /webhooks/{id} | Retrieve information for a specific webhook
[**WebhooksGetAll**](WebhooksApi.md#WebhooksGetAll) | **Get** /webhooks | Retrieve information about all existing webhooks



## WebhooksCreate

> Webhook WebhooksCreate(ctx).Webhook(webhook).Execute()

Create a webhook to subscribe to certain events



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
    webhook := *openapiclient.NewWebhook(openapiclient.webhook_event{EventGroup: openapiclient.NewEventGroup("group", "all")}, "https://www.example.com/webhook") // Webhook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.WebhooksCreate(context.Background()).Webhook(webhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhooksCreate`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhooksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksDelete

> WebhooksDelete(ctx, id).Execute()

Delete an existing webhook



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
    id := "wh_za7VbYcSQU2zRgGQXQAm-g" // string | The ID of the webhook to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.WebhooksDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the webhook to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksGet

> Webhook WebhooksGet(ctx, id).Execute()

Retrieve information for a specific webhook



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
    id := "wh_za7VbYcSQU2zRgGQXQAm-g" // string | The ID of the webhook whose information to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.WebhooksGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhooksGet`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the webhook whose information to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksGetAll

> WebhooksGetAll200Response WebhooksGetAll(ctx).XPublishableKey(xPublishableKey).Execute()

Retrieve information about all existing webhooks



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
    resp, r, err := apiClient.WebhooksApi.WebhooksGetAll(context.Background()).XPublishableKey(xPublishableKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhooksGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhooksGetAll`: WebhooksGetAll200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhooksGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPublishableKey** | **string** | The publicly viewable identifier used to identify a merchant division. | 

### Return type

[**WebhooksGetAll200Response**](WebhooksGetAll200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

