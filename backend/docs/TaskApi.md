# \TaskApi

All URIs are relative to *https://navigator-api.tweakwise.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskGetStatus**](TaskApi.md#TaskGetStatus) | **Get** /task/status/{key}/{taskkey} | 
[**TaskStart**](TaskApi.md#TaskStart) | **Get** /task/trigger/{key}/{taskkey} | This service starts the specified task by the taskkey.



## TaskGetStatus

> map[string]interface{} TaskGetStatus(ctx, key, taskkey).Execute()



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
    key := "key_example" // string | 
    taskkey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.TaskGetStatus(context.Background(), key, taskkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TaskGetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskGetStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TaskGetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**taskkey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGetStatusRequest struct via the builder pattern


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


## TaskStart

> map[string]interface{} TaskStart(ctx, key, taskkey).Execute()

This service starts the specified task by the taskkey.

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
    key := "key_example" // string | The instance key
    taskkey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the task

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.TaskStart(context.Background(), key, taskkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TaskStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskStart`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TaskStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The instance key | 
**taskkey** | **string** | The GUID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

