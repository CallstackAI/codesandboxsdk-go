# SandboxListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error5**](Error5.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**SandboxListResponseAllOfData**](SandboxListResponseAllOfData.md) |  | [optional] 

## Methods

### NewSandboxListResponse

`func NewSandboxListResponse() *SandboxListResponse`

NewSandboxListResponse instantiates a new SandboxListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxListResponseWithDefaults

`func NewSandboxListResponseWithDefaults() *SandboxListResponse`

NewSandboxListResponseWithDefaults instantiates a new SandboxListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *SandboxListResponse) GetErrors() []Error5`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SandboxListResponse) GetErrorsOk() (*[]Error5, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SandboxListResponse) SetErrors(v []Error5)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SandboxListResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *SandboxListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SandboxListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SandboxListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SandboxListResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *SandboxListResponse) GetData() SandboxListResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SandboxListResponse) GetDataOk() (*SandboxListResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SandboxListResponse) SetData(v SandboxListResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *SandboxListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


