# Product

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
**Categories** | Pointer to **[]int64** |  | [optional] 
**Attributes** | Pointer to [**[]ProductAttribute**](ProductAttribute.md) |  | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleNumber

`func (o *Product) GetArticleNumber() string`

GetArticleNumber returns the ArticleNumber field if non-nil, zero value otherwise.

### GetArticleNumberOk

`func (o *Product) GetArticleNumberOk() (*string, bool)`

GetArticleNumberOk returns a tuple with the ArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleNumber

`func (o *Product) SetArticleNumber(v string)`

SetArticleNumber sets ArticleNumber field to given value.

### HasArticleNumber

`func (o *Product) HasArticleNumber() bool`

HasArticleNumber returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Product) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Product) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBrand

`func (o *Product) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Product) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Product) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Product) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetStock

`func (o *Product) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *Product) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *Product) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *Product) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetGroupCode

`func (o *Product) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *Product) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *Product) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.

### HasGroupCode

`func (o *Product) HasGroupCode() bool`

HasGroupCode returns a boolean if a field has been set.

### GetImage

`func (o *Product) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Product) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Product) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Product) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetUrl

`func (o *Product) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Product) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Product) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Product) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCategories

`func (o *Product) GetCategories() []int64`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Product) GetCategoriesOk() (*[]int64, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Product) SetCategories(v []int64)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Product) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAttributes

`func (o *Product) GetAttributes() []ProductAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Product) GetAttributesOk() (*[]ProductAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Product) SetAttributes(v []ProductAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Product) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


