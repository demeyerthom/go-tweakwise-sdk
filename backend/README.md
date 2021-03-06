# Go API client for backend

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./backend"
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
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://navigator-api.tweakwise.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CategoryApi* | [**CategoryAddCategory**](docs/CategoryApi.md#categoryaddcategory) | **Post** /category | Creates a new category.
*CategoryApi* | [**CategoryDeleteCategory**](docs/CategoryApi.md#categorydeletecategory) | **Delete** /category/{categoryId} | Deletes a single category based on the category Id.
*CategoryApi* | [**CategoryGetAllCategories**](docs/CategoryApi.md#categorygetallcategories) | **Get** /category | Returns all categories.
*CategoryApi* | [**CategoryGetCategory**](docs/CategoryApi.md#categorygetcategory) | **Get** /category/{categoryId} | Returns a single category with the specified categoryId.
*CategoryApi* | [**CategoryGetCategoryByKey**](docs/CategoryApi.md#categorygetcategorybykey) | **Get** /category/getbykey/{key} | Returns a single category with the specified key.
*CategoryApi* | [**CategoryPatchCategory**](docs/CategoryApi.md#categorypatchcategory) | **Patch** /category/{categoryId} | Patches a category with the specified change.
*CategoryApi* | [**CategoryUpdateCategory**](docs/CategoryApi.md#categoryupdatecategory) | **Put** /category | Update an existing category.
*InstanceStatsApi* | [**InstanceStatsTotalCategories**](docs/InstanceStatsApi.md#instancestatstotalcategories) | **Get** /stats/totalcategories | Get the number of total categories.
*InstanceStatsApi* | [**InstanceStatsTotalProducts**](docs/InstanceStatsApi.md#instancestatstotalproducts) | **Get** /stats/totalproducts | Get the number of total products.
*ProductApi* | [**ProductAddProduct**](docs/ProductApi.md#productaddproduct) | **Post** /product | Creates a new product.
*ProductApi* | [**ProductDeleteProduct**](docs/ProductApi.md#productdeleteproduct) | **Delete** /product/{articleNumber} | Deletes a single product based on the articlenumber.
*ProductApi* | [**ProductGetAllProducts**](docs/ProductApi.md#productgetallproducts) | **Get** /product | Returns all products.
*ProductApi* | [**ProductGetProduct**](docs/ProductApi.md#productgetproduct) | **Get** /product/{articleNumber} | Returns a single product with the specified article number.
*ProductApi* | [**ProductPatchProduct**](docs/ProductApi.md#productpatchproduct) | **Patch** /product/{articleNumber} | Patches a product with the specified change.
*ProductApi* | [**ProductUpdateProduct**](docs/ProductApi.md#productupdateproduct) | **Put** /product | Updates a product with the specified data.
*Shopware6Api* | [**Shopware6Channels**](docs/Shopware6Api.md#shopware6channels) | **Get** /shopware6/channels | Get sales channels by shop
*Shopware6Api* | [**Shopware6CheckConnection**](docs/Shopware6Api.md#shopware6checkconnection) | **Get** /shopware6/check-connection | Check whether a connection exists for this shopId
*Shopware6Api* | [**Shopware6Connect**](docs/Shopware6Api.md#shopware6connect) | **Post** /shopware6/connect | Connects Shopware6 store to the Tweakwise Instance
*Shopware6Api* | [**Shopware6GetActiveChannels**](docs/Shopware6Api.md#shopware6getactivechannels) | **Get** /shopware6/active-channel | Returns an active channelID
*Shopware6Api* | [**Shopware6ScheduleFullSync**](docs/Shopware6Api.md#shopware6schedulefullsync) | **Post** /shopware6/schedule-full-sync | Schedules a full sync of the Shopware6 store into Tweakwise
*Shopware6Api* | [**Shopware6UpdateActiveChannels**](docs/Shopware6Api.md#shopware6updateactivechannels) | **Put** /shopware6/active-channel | Updates an active channelID
*TaskApi* | [**TaskGetStatus**](docs/TaskApi.md#taskgetstatus) | **Get** /task/status/{key}/{taskkey} | 
*TaskApi* | [**TaskStart**](docs/TaskApi.md#taskstart) | **Get** /task/trigger/{key}/{taskkey} | This service starts the specified task by the taskkey.
*UserApi* | [**UserGetInstances**](docs/UserApi.md#usergetinstances) | **Get** /user/instances | Returns all instances for current Tweakwise Account User


## Documentation For Models

 - [ActiveChannelModel](docs/ActiveChannelModel.md)
 - [Category](docs/Category.md)
 - [CategoryApiModel](docs/CategoryApiModel.md)
 - [CategoryDeltaApiModel](docs/CategoryDeltaApiModel.md)
 - [ConnectModel](docs/ConnectModel.md)
 - [Product](docs/Product.md)
 - [ProductApiModel](docs/ProductApiModel.md)
 - [ProductAttribute](docs/ProductAttribute.md)
 - [ProductAttributeApiModel](docs/ProductAttributeApiModel.md)
 - [ProductDeltaApiModel](docs/ProductDeltaApiModel.md)
 - [ProductSummary](docs/ProductSummary.md)
 - [SalesChannel](docs/SalesChannel.md)
 - [UserInstance](docs/UserInstance.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorContainer](docs/ValidationErrorContainer.md)


## Documentation For Authorization



### apiKeyDefinition

- **Type**: API key
- **API key parameter name**: TWN-Authentication
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: TWN-Authentication and passed in as the auth context for each request.


### instanceKeyDefinition

- **Type**: API key
- **API key parameter name**: TWN-InstanceKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: TWN-InstanceKey and passed in as the auth context for each request.


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



