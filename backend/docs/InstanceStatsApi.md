# \InstanceStatsApi

All URIs are relative to *https://navigator-api.tweakwise.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceStatsTotalCategories**](InstanceStatsApi.md#InstanceStatsTotalCategories) | **Get** /stats/totalcategories | Get the number of total categories.
[**InstanceStatsTotalProducts**](InstanceStatsApi.md#InstanceStatsTotalProducts) | **Get** /stats/totalproducts | Get the number of total products.



## InstanceStatsTotalCategories

> map[string]interface{} InstanceStatsTotalCategories(ctx).Execute()

Get the number of total categories.

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
    resp, r, err := api_client.InstanceStatsApi.InstanceStatsTotalCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceStatsApi.InstanceStatsTotalCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstanceStatsTotalCategories`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstanceStatsApi.InstanceStatsTotalCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceStatsTotalCategoriesRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceStatsTotalProducts

> map[string]interface{} InstanceStatsTotalProducts(ctx).Execute()

Get the number of total products.

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
    resp, r, err := api_client.InstanceStatsApi.InstanceStatsTotalProducts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceStatsApi.InstanceStatsTotalProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstanceStatsTotalProducts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstanceStatsApi.InstanceStatsTotalProducts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceStatsTotalProductsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

