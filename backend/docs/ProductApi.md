# \ProductApi

All URIs are relative to *https://navigator-api.tweakwise.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductAddProduct**](ProductApi.md#ProductAddProduct) | **Post** /product | Creates a new product.
[**ProductDeleteProduct**](ProductApi.md#ProductDeleteProduct) | **Delete** /product/{articleNumber} | Deletes a single product based on the articlenumber.
[**ProductGetAllProducts**](ProductApi.md#ProductGetAllProducts) | **Get** /product | Returns all products.
[**ProductGetProduct**](ProductApi.md#ProductGetProduct) | **Get** /product/{articleNumber} | Returns a single product with the specified article number.
[**ProductPatchProduct**](ProductApi.md#ProductPatchProduct) | **Patch** /product/{articleNumber} | Patches a product with the specified change.
[**ProductUpdateProduct**](ProductApi.md#ProductUpdateProduct) | **Put** /product | Updates a product with the specified data.



## ProductAddProduct

> ProductAddProduct(ctx).Product(product).Execute()

Creates a new product.



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
    product := *openapiclient.NewProductApiModel() // ProductApiModel | A product to add

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.ProductAddProduct(context.Background()).Product(product).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductAddProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAddProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**ProductApiModel**](ProductApiModel.md) | A product to add | 

### Return type

 (empty response body)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDeleteProduct

> ProductDeleteProduct(ctx, articleNumber).Execute()

Deletes a single product based on the articlenumber.



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
    articleNumber := "articleNumber_example" // string | An article number to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.ProductDeleteProduct(context.Background(), articleNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductDeleteProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleNumber** | **string** | An article number to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ProductGetAllProducts

> []ProductSummary ProductGetAllProducts(ctx).Execute()

Returns all products.



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
    resp, r, err := api_client.ProductApi.ProductGetAllProducts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetAllProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetAllProducts`: []ProductSummary
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetAllProducts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductGetAllProductsRequest struct via the builder pattern


### Return type

[**[]ProductSummary**](ProductSummary.md)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetProduct

> Product ProductGetProduct(ctx, articleNumber).Execute()

Returns a single product with the specified article number.



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
    articleNumber := "articleNumber_example" // string | An article number to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.ProductGetProduct(context.Background(), articleNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleNumber** | **string** | An article number to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Product**](Product.md)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductPatchProduct

> map[string]interface{} ProductPatchProduct(ctx, articleNumber).Product(product).Execute()

Patches a product with the specified change.



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
    articleNumber := "articleNumber_example" // string | An articlenumber to patch
    product := *openapiclient.NewProductDeltaApiModel() // ProductDeltaApiModel | The product with values to patch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.ProductPatchProduct(context.Background(), articleNumber).Product(product).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductPatchProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductPatchProduct`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductPatchProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleNumber** | **string** | An articlenumber to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductPatchProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | [**ProductDeltaApiModel**](ProductDeltaApiModel.md) | The product with values to patch | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateProduct

> ProductUpdateProduct(ctx).Product(product).Execute()

Updates a product with the specified data.



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
    product := *openapiclient.NewProductApiModel() // ProductApiModel | A filled product to replace the existing product

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.ProductUpdateProduct(context.Background()).Product(product).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductUpdateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**ProductApiModel**](ProductApiModel.md) | A filled product to replace the existing product | 

### Return type

 (empty response body)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

