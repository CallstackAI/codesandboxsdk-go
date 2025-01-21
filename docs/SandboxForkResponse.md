# SandboxForkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error3**](Error3.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**SandboxForkResponseAllOfData**](SandboxForkResponseAllOfData.md) |  | [optional] 

## Methods

### NewSandboxForkResponse

`func NewSandboxForkResponse() *SandboxForkResponse`

NewSandboxForkResponse instantiates a new SandboxForkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxForkResponseWithDefaults

`func NewSandboxForkResponseWithDefaults() *SandboxForkResponse`

NewSandboxForkResponseWithDefaults instantiates a new SandboxForkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *SandboxForkResponse) GetErrors() []Error3`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SandboxForkResponse) GetErrorsOk() (*[]Error3, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SandboxForkResponse) SetErrors(v []Error3)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SandboxForkResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *SandboxForkResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SandboxForkResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SandboxForkResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SandboxForkResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *SandboxForkResponse) GetData() SandboxForkResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SandboxForkResponse) GetDataOk() (*SandboxForkResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SandboxForkResponse) SetData(v SandboxForkResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *SandboxForkResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


