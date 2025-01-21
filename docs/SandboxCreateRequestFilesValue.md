# SandboxCreateRequestFilesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryContent** | Pointer to **string** | If the file has binary (non plain-text) contents, place the base-64 encoded contents in this key. Should be empty or missing if &#x60;is_binary&#x60; is &#x60;false&#x60;. | [optional] 
**Code** | Pointer to **string** | If the file is non-binary in nature, place the (escaped) plain text contents in this key. Should be empty or missing if &#x60;is_binary&#x60; is &#x60;true&#x60;. | [optional] 
**IsBinary** | Pointer to **bool** | Whether the file contains binary contents. | [optional] [default to false]

## Methods

### NewSandboxCreateRequestFilesValue

`func NewSandboxCreateRequestFilesValue() *SandboxCreateRequestFilesValue`

NewSandboxCreateRequestFilesValue instantiates a new SandboxCreateRequestFilesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCreateRequestFilesValueWithDefaults

`func NewSandboxCreateRequestFilesValueWithDefaults() *SandboxCreateRequestFilesValue`

NewSandboxCreateRequestFilesValueWithDefaults instantiates a new SandboxCreateRequestFilesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryContent

`func (o *SandboxCreateRequestFilesValue) GetBinaryContent() string`

GetBinaryContent returns the BinaryContent field if non-nil, zero value otherwise.

### GetBinaryContentOk

`func (o *SandboxCreateRequestFilesValue) GetBinaryContentOk() (*string, bool)`

GetBinaryContentOk returns a tuple with the BinaryContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryContent

`func (o *SandboxCreateRequestFilesValue) SetBinaryContent(v string)`

SetBinaryContent sets BinaryContent field to given value.

### HasBinaryContent

`func (o *SandboxCreateRequestFilesValue) HasBinaryContent() bool`

HasBinaryContent returns a boolean if a field has been set.

### GetCode

`func (o *SandboxCreateRequestFilesValue) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SandboxCreateRequestFilesValue) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SandboxCreateRequestFilesValue) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SandboxCreateRequestFilesValue) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIsBinary

`func (o *SandboxCreateRequestFilesValue) GetIsBinary() bool`

GetIsBinary returns the IsBinary field if non-nil, zero value otherwise.

### GetIsBinaryOk

`func (o *SandboxCreateRequestFilesValue) GetIsBinaryOk() (*bool, bool)`

GetIsBinaryOk returns a tuple with the IsBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBinary

`func (o *SandboxCreateRequestFilesValue) SetIsBinary(v bool)`

SetIsBinary sets IsBinary field to given value.

### HasIsBinary

`func (o *SandboxCreateRequestFilesValue) HasIsBinary() bool`

HasIsBinary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


