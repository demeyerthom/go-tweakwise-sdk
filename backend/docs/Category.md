# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int64** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rank** | Pointer to **int64** |  | [optional] 
**Parents** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCategory

`func NewCategory() *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *Category) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Category) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Category) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *Category) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetKey

`func (o *Category) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Category) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Category) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Category) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Category) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRank

`func (o *Category) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Category) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Category) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *Category) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetParents

`func (o *Category) GetParents() []int64`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Category) GetParentsOk() (*[]int64, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Category) SetParents(v []int64)`

SetParents sets Parents field to given value.

### HasParents

`func (o *Category) HasParents() bool`

HasParents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


