# \Shopware6Api

All URIs are relative to *https://navigator-api.tweakwise.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Shopware6Channels**](Shopware6Api.md#Shopware6Channels) | **Get** /shopware6/channels | Get sales channels by shop
[**Shopware6CheckConnection**](Shopware6Api.md#Shopware6CheckConnection) | **Get** /shopware6/check-connection | Check whether a connection exists for this shopId
[**Shopware6Connect**](Shopware6Api.md#Shopware6Connect) | **Post** /shopware6/connect | Connects Shopware6 store to the Tweakwise Instance
[**Shopware6GetActiveChannels**](Shopware6Api.md#Shopware6GetActiveChannels) | **Get** /shopware6/active-channel | Returns an active channelID
[**Shopware6ScheduleFullSync**](Shopware6Api.md#Shopware6ScheduleFullSync) | **Post** /shopware6/schedule-full-sync | Schedules a full sync of the Shopware6 store into Tweakwise
[**Shopware6UpdateActiveChannels**](Shopware6Api.md#Shopware6UpdateActiveChannels) | **Put** /shopware6/active-channel | Updates an active channelID



## Shopware6Channels

> []SalesChannel Shopware6Channels(ctx).ShopId(shopId).Execute()

Get sales channels by shop

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shopId := "shopId_example" // string | Shopware shop id the sales channels need to be provided for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Shopware6Api.Shopware6Channels(context.Background()).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Shopware6Api.Shopware6Channels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Shopware6Channels`: []SalesChannel
    fmt.Fprintf(os.Stdout, "Response from `Shopware6Api.Shopware6Channels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShopware6ChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shopId** | **string** | Shopware shop id the sales channels need to be provided for | 

### Return type

[**[]SalesChannel**](SalesChannel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Shopware6CheckConnection

> map[string]interface{} Shopware6CheckConnection(ctx).ShopId(shopId).Execute()

Check whether a connection exists for this shopId

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shopId := "shopId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Shopware6Api.Shopware6CheckConnection(context.Background()).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Shopware6Api.Shopware6CheckConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Shopware6CheckConnection`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Shopware6Api.Shopware6CheckConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShopware6CheckConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shopId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Shopware6Connect

> Shopware6Connect(ctx).Model(model).Execute()

Connects Shopware6 store to the Tweakwise Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    model := *openapiclient.NewConnectModel() // ConnectModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Shopware6Api.Shopware6Connect(context.Background()).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Shopware6Api.Shopware6Connect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShopware6ConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**ConnectModel**](ConnectModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Shopware6GetActiveChannels

> Shopware6GetActiveChannels(ctx).ShopId(shopId).Execute()

Returns an active channelID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shopId := "shopId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Shopware6Api.Shopware6GetActiveChannels(context.Background()).ShopId(shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Shopware6Api.Shopware6GetActiveChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShopware6GetActiveChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shopId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Shopware6ScheduleFullSync

> Shopware6ScheduleFullSync(ctx).Execute()

Schedules a full sync of the Shopware6 store into Tweakwise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Shopware6Api.Shopware6ScheduleFullSync(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Shopware6Api.Shopware6ScheduleFullSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShopware6ScheduleFullSyncRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Shopware6UpdateActiveChannels

> Shopware6UpdateActiveChannels(ctx).ShopId(shopId).Model(model).Execute()

Updates an active channelID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shopId := "shopId_example" // string | 
    model := *openapiclient.NewActiveChannelModel("ChannelId_example") // ActiveChannelModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Shopware6Api.Shopware6UpdateActiveChannels(context.Background()).ShopId(shopId).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Shopware6Api.Shopware6UpdateActiveChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShopware6UpdateActiveChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shopId** | **string** |  | 
 **model** | [**ActiveChannelModel**](ActiveChannelModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

