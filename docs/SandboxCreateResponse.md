# SandboxCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error2**](Error2.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**SandboxCreateResponseAllOfData**](SandboxCreateResponseAllOfData.md) |  | [optional] 

## Methods

### NewSandboxCreateResponse

`func NewSandboxCreateResponse() *SandboxCreateResponse`

NewSandboxCreateResponse instantiates a new SandboxCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCreateResponseWithDefaults

`func NewSandboxCreateResponseWithDefaults() *SandboxCreateResponse`

NewSandboxCreateResponseWithDefaults instantiates a new SandboxCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *SandboxCreateResponse) GetErrors() []Error2`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SandboxCreateResponse) GetErrorsOk() (*[]Error2, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SandboxCreateResponse) SetErrors(v []Error2)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SandboxCreateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *SandboxCreateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SandboxCreateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SandboxCreateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SandboxCreateResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *SandboxCreateResponse) GetData() SandboxCreateResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SandboxCreateResponse) GetDataOk() (*SandboxCreateResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SandboxCreateResponse) SetData(v SandboxCreateResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *SandboxCreateResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


