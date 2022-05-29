# ValidationErrorContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message | [optional] [readonly] 
**Errors** | Pointer to [**[]ValidationError**](ValidationError.md) | List of validation Errors | [optional] 

## Methods

### NewValidationErrorContainer

`func NewValidationErrorContainer() *ValidationErrorContainer`

NewValidationErrorContainer instantiates a new ValidationErrorContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorContainerWithDefaults

`func NewValidationErrorContainerWithDefaults() *ValidationErrorContainer`

NewValidationErrorContainerWithDefaults instantiates a new ValidationErrorContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ValidationErrorContainer) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorContainer) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorContainer) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationErrorContainer) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *ValidationErrorContainer) GetErrors() []ValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationErrorContainer) GetErrorsOk() (*[]ValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationErrorContainer) SetErrors(v []ValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidationErrorContainer) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


