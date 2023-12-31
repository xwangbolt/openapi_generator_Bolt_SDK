# Go API client for openapi

A comprehensive Bolt API reference for interacting with Transactions, Orders, Product Catalog, Configuration, Testing, and much more.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://help.bolt.com/api-bolt/](https://help.bolt.com/api-bolt/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.xwang.dev.bolt.me/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountAPI* | [**AccountAddPaymentMethod**](docs/AccountAPI.md#accountaddpaymentmethod) | **Post** /account/payment-methods | Add a payment method to a shopper&#39;s Bolt account Wallet.
*AccountAPI* | [**AccountAddressCreate**](docs/AccountAPI.md#accountaddresscreate) | **Post** /account/addresses | Add an address
*AccountAPI* | [**AccountAddressDelete**](docs/AccountAPI.md#accountaddressdelete) | **Delete** /account/addresses/{id} | Delete an existing address
*AccountAPI* | [**AccountAddressEdit**](docs/AccountAPI.md#accountaddressedit) | **Put** /account/addresses/{id} | Edit an existing address
*AccountAPI* | [**AccountExists**](docs/AccountAPI.md#accountexists) | **Get** /account/exists | Determine the existence of a Bolt account
*AccountAPI* | [**AccountGet**](docs/AccountAPI.md#accountget) | **Get** /account | Retrieve account details
*AccountAPI* | [**AccountPaymentMethodDelete**](docs/AccountAPI.md#accountpaymentmethoddelete) | **Delete** /account/payment-methods/{id} | Delete an existing payment method
*ConfigurationAPI* | [**MerchantCallbacksGet**](docs/ConfigurationAPI.md#merchantcallbacksget) | **Get** /merchant/callbacks | Retrieve callback URLs for the merchant
*ConfigurationAPI* | [**MerchantCallbacksUpdate**](docs/ConfigurationAPI.md#merchantcallbacksupdate) | **Patch** /merchant/callbacks | Update callback URLs for the merchant
*ConfigurationAPI* | [**MerchantIdentifiersGet**](docs/ConfigurationAPI.md#merchantidentifiersget) | **Get** /merchant/identifiers | Retrieve identifiers for the merchant
*PaymentsAPI* | [**GuestPaymentsInitialize**](docs/PaymentsAPI.md#guestpaymentsinitialize) | **Post** /guest/payments | Initialize a Bolt payment for guest shoppers
*PaymentsAPI* | [**PaymentsInitialize**](docs/PaymentsAPI.md#paymentsinitialize) | **Post** /payments | Initialize a Bolt payment for logged in shoppers
*TestingAPI* | [**TestingAccountCreate**](docs/TestingAPI.md#testingaccountcreate) | **Post** /testing/accounts | Create a test account
*TestingAPI* | [**TestingCreditCardGet**](docs/TestingAPI.md#testingcreditcardget) | **Get** /testing/credit-cards | Retrieve a test credit card, including its token
*TestingAPI* | [**TestingShipmentTrackingCreate**](docs/TestingAPI.md#testingshipmenttrackingcreate) | **Post** /testing/shipments | Simulate a shipment tracking update
*WebhooksAPI* | [**WebhooksCreate**](docs/WebhooksAPI.md#webhookscreate) | **Put** /webhooks | Create a webhook to subscribe to certain events
*WebhooksAPI* | [**WebhooksDelete**](docs/WebhooksAPI.md#webhooksdelete) | **Delete** /webhooks/{id} | Delete an existing webhook
*WebhooksAPI* | [**WebhooksGet**](docs/WebhooksAPI.md#webhooksget) | **Get** /webhooks/{id} | Retrieve information for a specific webhook
*WebhooksAPI* | [**WebhooksGetAll**](docs/WebhooksAPI.md#webhooksgetall) | **Get** /webhooks | Retrieve information about all existing webhooks


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountAddressCreate400Response](docs/AccountAddressCreate400Response.md)
 - [AccountExistsIdentifierParameter](docs/AccountExistsIdentifierParameter.md)
 - [AccountTestCreationData](docs/AccountTestCreationData.md)
 - [Address](docs/Address.md)
 - [AddressErrorInvalidPostalCode](docs/AddressErrorInvalidPostalCode.md)
 - [AddressErrorInvalidRegion](docs/AddressErrorInvalidRegion.md)
 - [AddressListing](docs/AddressListing.md)
 - [AddressReference](docs/AddressReference.md)
 - [AddressReferenceExplicit](docs/AddressReferenceExplicit.md)
 - [AddressReferenceId](docs/AddressReferenceId.md)
 - [Amounts](docs/Amounts.md)
 - [CallbackUrlErrorInvalidUrl](docs/CallbackUrlErrorInvalidUrl.md)
 - [CallbackUrls](docs/CallbackUrls.md)
 - [Cart](docs/Cart.md)
 - [CartDiscount](docs/CartDiscount.md)
 - [CartItem](docs/CartItem.md)
 - [CartShipment](docs/CartShipment.md)
 - [CreditCard](docs/CreditCard.md)
 - [EventGroup](docs/EventGroup.md)
 - [EventList](docs/EventList.md)
 - [GuestPaymentMethodInitializeRequest](docs/GuestPaymentMethodInitializeRequest.md)
 - [GuestPaymentMethodInitializeRequestPaymentMethod](docs/GuestPaymentMethodInitializeRequestPaymentMethod.md)
 - [Identifiers](docs/Identifiers.md)
 - [IdentifiersMerchantDivisionsInner](docs/IdentifiersMerchantDivisionsInner.md)
 - [MerchantCallbacksUpdate400Response](docs/MerchantCallbacksUpdate400Response.md)
 - [ModelError](docs/ModelError.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [PaymentMethodCreditCard](docs/PaymentMethodCreditCard.md)
 - [PaymentMethodInitializeRequest](docs/PaymentMethodInitializeRequest.md)
 - [PaymentMethodInitializeRequestPaymentMethod](docs/PaymentMethodInitializeRequestPaymentMethod.md)
 - [PaymentMethodInitializeResponse](docs/PaymentMethodInitializeResponse.md)
 - [PaymentMethodInitializeResponseAction](docs/PaymentMethodInitializeResponseAction.md)
 - [PaymentMethodPaypal](docs/PaymentMethodPaypal.md)
 - [PaymentMethodSavedPaymentMethod](docs/PaymentMethodSavedPaymentMethod.md)
 - [Profile](docs/Profile.md)
 - [ShipmentTrackingUpdate](docs/ShipmentTrackingUpdate.md)
 - [ShipmentTrackingUpdateTrackingDetailsInner](docs/ShipmentTrackingUpdateTrackingDetailsInner.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookEvent](docs/WebhookEvent.md)
 - [WebhooksGetAll200Response](docs/WebhooksGetAll200Response.md)
 - [WebhooksPostRequest](docs/WebhooksPostRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api-key

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Key and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"X-API-Key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### oauth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: /v1/oauth/authorize
- **Scopes**: 
 - **bolt.account.manage**: This scope grants permissions to perform read/edit/delete actions on Bolt Account data
 - **bolt.account.view**: This scope grants permissions to perform read only actions on Bolt Account data
 - **openid**: This scope grants permissions that enable Bolt SSO by granting an id token JWT that stores account data

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

api-help@bolt.com

