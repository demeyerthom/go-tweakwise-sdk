# \CategoryApi

All URIs are relative to *https://navigator-api.tweakwise.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoryAddCategory**](CategoryApi.md#CategoryAddCategory) | **Post** /category | Creates a new category.
[**CategoryDeleteCategory**](CategoryApi.md#CategoryDeleteCategory) | **Delete** /category/{categoryId} | Deletes a single category based on the category Id.
[**CategoryGetAllCategories**](CategoryApi.md#CategoryGetAllCategories) | **Get** /category | Returns all categories.
[**CategoryGetCategory**](CategoryApi.md#CategoryGetCategory) | **Get** /category/{categoryId} | Returns a single category with the specified categoryId.
[**CategoryGetCategoryByKey**](CategoryApi.md#CategoryGetCategoryByKey) | **Get** /category/getbykey/{key} | Returns a single category with the specified key.
[**CategoryPatchCategory**](CategoryApi.md#CategoryPatchCategory) | **Patch** /category/{categoryId} | Patches a category with the specified change.
[**CategoryUpdateCategory**](CategoryApi.md#CategoryUpdateCategory) | **Put** /category | Update an existing category.



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
    category := *openapiclient.NewCategoryApiModel() // CategoryApiModel | Category object that needs to be added to the store

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
 **category** | [**CategoryApiModel**](CategoryApiModel.md) | Category object that needs to be added to the store | 

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

> map[string]interface{} CategoryDeleteCategory(ctx, categoryId).Execute()

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
    categoryId := int64(789) // int64 | ID of the category to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryDeleteCategory(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryDeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryDeleteCategory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryDeleteCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | ID of the category to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryDeleteCategoryRequest struct via the builder pattern


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


## CategoryGetAllCategories

> []Category CategoryGetAllCategories(ctx).Execute()

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
    // response from `CategoryGetAllCategories`: []Category
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryGetAllCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryGetAllCategoriesRequest struct via the builder pattern


### Return type

[**[]Category**](Category.md)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryGetCategory

> CategoryApiModel CategoryGetCategory(ctx, categoryId).Execute()

Returns a single category with the specified categoryId.



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
    categoryId := int64(789) // int64 | The categoryId/id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryGetCategory(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryGetCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryGetCategory`: CategoryApiModel
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryGetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | The categoryId/id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CategoryApiModel**](CategoryApiModel.md)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryGetCategoryByKey

> Category CategoryGetCategoryByKey(ctx, key).Execute()

Returns a single category with the specified key.



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
    key := "key_example" // string | an unique reference to get a category

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryGetCategoryByKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryApi.CategoryGetCategoryByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoryGetCategoryByKey`: Category
    fmt.Fprintf(os.Stdout, "Response from `CategoryApi.CategoryGetCategoryByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | an unique reference to get a category | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryGetCategoryByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Category**](Category.md)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryPatchCategory

> map[string]interface{} CategoryPatchCategory(ctx, categoryId).CategoryDelta(categoryDelta).Execute()

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
    categoryId := int64(789) // int64 | The ID to patch
    categoryDelta := *openapiclient.NewCategoryDeltaApiModel() // CategoryDeltaApiModel | A category object with only the values that needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryApi.CategoryPatchCategory(context.Background(), categoryId).CategoryDelta(categoryDelta).Execute()
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
**categoryId** | **int64** | The ID to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryPatchCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryDelta** | [**CategoryDeltaApiModel**](CategoryDeltaApiModel.md) | A category object with only the values that needs to be updated. | 

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

Update an existing category.



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
    category := *openapiclient.NewCategoryApiModel() // CategoryApiModel | A category object with changes to update the given category

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
 **category** | [**CategoryApiModel**](CategoryApiModel.md) | A category object with changes to update the given category | 

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

