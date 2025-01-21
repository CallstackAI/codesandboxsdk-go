# SandboxGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error4**](Error4.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**SandboxGetResponseAllOfData**](SandboxGetResponseAllOfData.md) |  | [optional] 

## Methods

### NewSandboxGetResponse

`func NewSandboxGetResponse() *SandboxGetResponse`

NewSandboxGetResponse instantiates a new SandboxGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxGetResponseWithDefaults

`func NewSandboxGetResponseWithDefaults() *SandboxGetResponse`

NewSandboxGetResponseWithDefaults instantiates a new SandboxGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *SandboxGetResponse) GetErrors() []Error4`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SandboxGetResponse) GetErrorsOk() (*[]Error4, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SandboxGetResponse) SetErrors(v []Error4)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SandboxGetResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *SandboxGetResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SandboxGetResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SandboxGetResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SandboxGetResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *SandboxGetResponse) GetData() SandboxGetResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SandboxGetResponse) GetDataOk() (*SandboxGetResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SandboxGetResponse) SetData(v SandboxGetResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *SandboxGetResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


