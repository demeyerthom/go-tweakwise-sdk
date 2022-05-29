# \UserApi

All URIs are relative to *https://navigator-api.tweakwise.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGetInstances**](UserApi.md#UserGetInstances) | **Get** /user/instances | Returns all instances for current Tweakwise Account User



## UserGetInstances

> []UserInstance UserGetInstances(ctx).Execute()

Returns all instances for current Tweakwise Account User

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
    resp, r, err := api_client.UserApi.UserGetInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetInstances`: []UserInstance
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetInstancesRequest struct via the builder pattern


### Return type

[**[]UserInstance**](UserInstance.md)

### Authorization

[apiKeyDefinition](../README.md#apiKeyDefinition), [instanceKeyDefinition](../README.md#instanceKeyDefinition)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

