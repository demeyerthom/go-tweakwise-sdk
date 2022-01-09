# ProductApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleNumber** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Stock** | Pointer to **int32** |  | [optional] 
**GroupCode** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to [**[]ProductAttributeApiModel**](ProductAttributeApiModel.md) |  | [optional] 

## Methods

### NewProductApiModel

`func NewProductApiModel() *ProductApiModel`

NewProductApiModel instantiates a new ProductApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductApiModelWithDefaults

`func NewProductApiModelWithDefaults() *ProductApiModel`

NewProductApiModelWithDefaults instantiates a new ProductApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleNumber

`func (o *ProductApiModel) GetArticleNumber() string`

GetArticleNumber returns the ArticleNumber field if non-nil, zero value otherwise.

### GetArticleNumberOk

`func (o *ProductApiModel) GetArticleNumberOk() (*string, bool)`

GetArticleNumberOk returns a tuple with the ArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleNumber

`func (o *ProductApiModel) SetArticleNumber(v string)`

SetArticleNumber sets ArticleNumber field to given value.

### HasArticleNumber

`func (o *ProductApiModel) HasArticleNumber() bool`

HasArticleNumber returns a boolean if a field has been set.

### GetName

`func (o *ProductApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ProductApiModel) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductApiModel) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductApiModel) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductApiModel) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBrand

`func (o *ProductApiModel) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ProductApiModel) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ProductApiModel) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ProductApiModel) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetStock

`func (o *ProductApiModel) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ProductApiModel) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ProductApiModel) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ProductApiModel) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetGroupCode

`func (o *ProductApiModel) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *ProductApiModel) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *ProductApiModel) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.

### HasGroupCode

`func (o *ProductApiModel) HasGroupCode() bool`

HasGroupCode returns a boolean if a field has been set.

### GetImage

`func (o *ProductApiModel) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductApiModel) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductApiModel) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProductApiModel) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetUrl

`func (o *ProductApiModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductApiModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductApiModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductApiModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCategories

`func (o *ProductApiModel) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ProductApiModel) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ProductApiModel) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ProductApiModel) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAttributes

`func (o *ProductApiModel) GetAttributes() []ProductAttributeApiModel`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProductApiModel) GetAttributesOk() (*[]ProductAttributeApiModel, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProductApiModel) SetAttributes(v []ProductAttributeApiModel)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProductApiModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


