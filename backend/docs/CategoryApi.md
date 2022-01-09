# \CategoryApi

All URIs are relative to *https://navigator-api.tweakwise.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoryAddCategory**](CategoryApi.md#CategoryAddCategory) | **Post** /category | Creates a new category.
[**CategoryDeleteCategory**](CategoryApi.md#CategoryDeleteCategory) | **Delete** /category/{categoryKey} | Deletes a single category based on the category Id.
[**CategoryDeleteImportCategory**](CategoryApi.md#CategoryDeleteImportCategory) | **Delete** /import/category/{categoryKey} | Deletes a single category.
[**CategoryGetAllCategories**](CategoryApi.md#CategoryGetAllCategories) | **Get** /category | Returns all categories.
[**CategoryGetCategory**](CategoryApi.md#CategoryGetCategory) | **Get** /category/{categoryKey} | Returns a single category with the specified categorykey.
[**CategoryPatchCategory**](CategoryApi.md#CategoryPatchCategory) | **Patch** /category/{categoryKey} | Patches a category with the specified change.
[**CategoryUpdateCategory**](CategoryApi.md#CategoryUpdateCategory) | **Put** /category | Updates a category with the specified data.



## CategoryAddCategory

> map[string]interface{} CategoryAddCategory(ctx).Category(category).Execute()

Creates a new category.



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
    category := *openapiclient.NewCategoryApiModel() // CategoryApiModel | A category to add

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryAddCategory(context.Background()).Category(category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryAddCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryAddCategory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryAddCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryAddCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**CategoryApiModel**](CategoryApiModel.md) | A category to add | 

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


## CategoryDeleteCategory

> CategoryDeleteCategory(ctx, categoryKey).Execute()

Deletes a single category based on the category Id.



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
    categoryKey := "categoryKey_example" // string | The category to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryDeleteCategory(context.Background(), categoryKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryDeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryKey** | **string** | The category to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryDeleteCategoryRequest struct via the builder pattern


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


## CategoryDeleteImportCategory

> CategoryDeleteImportCategory(ctx, categoryKey).Execute()

Deletes a single category.



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
    categoryKey := "categoryKey_example" // string | The category to delete permanently

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryDeleteImportCategory(context.Background(), categoryKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryDeleteImportCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryKey** | **string** | The category to delete permanently | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryDeleteImportCategoryRequest struct via the builder pattern


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


## CategoryGetAllCategories

> map[string]interface{} CategoryGetAllCategories(ctx).Execute()

Returns all categories.



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
    resp, r, err := api_client.CategoryApi.CategoryGetAllCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryGetAllCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryGetAllCategories`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryGetAllCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryGetAllCategoriesRequest struct via the builder pattern


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


## CategoryGetCategory

> map[string]interface{} CategoryGetCategory(ctx, categoryKey).Execute()

Returns a single category with the specified categorykey.



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
    categoryKey := "categoryKey_example" // string | The categorykey/id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryGetCategory(context.Background(), categoryKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryGetCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryGetCategory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryGetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryKey** | **string** | The categorykey/id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CategoryPatchCategory

> map[string]interface{} CategoryPatchCategory(ctx, categoryKey).CategoryDelta(categoryDelta).Execute()

Patches a category with the specified change.



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
    categoryKey := "categoryKey_example" // string | The category to patch
    categoryDelta := *openapiclient.NewCategoryDeltaApiModel() // CategoryDeltaApiModel | The category with values to patch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryPatchCategory(context.Background(), categoryKey).CategoryDelta(categoryDelta).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryPatchCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryPatchCategory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryPatchCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryKey** | **string** | The category to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryPatchCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryDelta** | [**CategoryDeltaApiModel**](CategoryDeltaApiModel.md) | The category with values to patch | 

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


## CategoryUpdateCategory

> map[string]interface{} CategoryUpdateCategory(ctx).Category(category).Execute()

Updates a category with the specified data.



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
    category := *openapiclient.NewCategoryApiModel() // CategoryApiModel | A filled category to replace the existing category

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryUpdateCategory(context.Background()).Category(category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryUpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryUpdateCategory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryUpdateCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**CategoryApiModel**](CategoryApiModel.md) | A filled category to replace the existing category | 

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

