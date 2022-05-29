# ProductDeltaApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Stock** | Pointer to **int32** |  | [optional] 
**GroupCode** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]int64** |  | [optional] 
**Attributes** | Pointer to [**[]ProductAttributeApiModel**](ProductAttributeApiModel.md) |  | [optional] 

## Methods

### NewProductDeltaApiModel

`func NewProductDeltaApiModel() *ProductDeltaApiModel`

NewProductDeltaApiModel instantiates a new ProductDeltaApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDeltaApiModelWithDefaults

`func NewProductDeltaApiModelWithDefaults() *ProductDeltaApiModel`

NewProductDeltaApiModelWithDefaults instantiates a new ProductDeltaApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductDeltaApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductDeltaApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductDeltaApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductDeltaApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ProductDeltaApiModel) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductDeltaApiModel) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductDeltaApiModel) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductDeltaApiModel) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBrand

`func (o *ProductDeltaApiModel) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ProductDeltaApiModel) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ProductDeltaApiModel) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ProductDeltaApiModel) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetStock

`func (o *ProductDeltaApiModel) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ProductDeltaApiModel) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ProductDeltaApiModel) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ProductDeltaApiModel) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetGroupCode

`func (o *ProductDeltaApiModel) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *ProductDeltaApiModel) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *ProductDeltaApiModel) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.

### HasGroupCode

`func (o *ProductDeltaApiModel) HasGroupCode() bool`

HasGroupCode returns a boolean if a field has been set.

### GetImage

`func (o *ProductDeltaApiModel) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductDeltaApiModel) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductDeltaApiModel) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProductDeltaApiModel) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetUrl

`func (o *ProductDeltaApiModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductDeltaApiModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductDeltaApiModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductDeltaApiModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCategories

`func (o *ProductDeltaApiModel) GetCategories() []int64`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ProductDeltaApiModel) GetCategoriesOk() (*[]int64, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ProductDeltaApiModel) SetCategories(v []int64)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ProductDeltaApiModel) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAttributes

`func (o *ProductDeltaApiModel) GetAttributes() []ProductAttributeApiModel`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProductDeltaApiModel) GetAttributesOk() (*[]ProductAttributeApiModel, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProductDeltaApiModel) SetAttributes(v []ProductAttributeApiModel)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProductDeltaApiModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


