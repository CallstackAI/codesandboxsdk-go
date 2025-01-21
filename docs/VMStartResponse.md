# VMStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error10**](Error10.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**VMStartResponseAllOfData**](VMStartResponseAllOfData.md) |  | [optional] 

## Methods

### NewVMStartResponse

`func NewVMStartResponse() *VMStartResponse`

NewVMStartResponse instantiates a new VMStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMStartResponseWithDefaults

`func NewVMStartResponseWithDefaults() *VMStartResponse`

NewVMStartResponseWithDefaults instantiates a new VMStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *VMStartResponse) GetErrors() []Error10`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VMStartResponse) GetErrorsOk() (*[]Error10, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VMStartResponse) SetErrors(v []Error10)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *VMStartResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *VMStartResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *VMStartResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *VMStartResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *VMStartResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *VMStartResponse) GetData() VMStartResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VMStartResponse) GetDataOk() (*VMStartResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VMStartResponse) SetData(v VMStartResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *VMStartResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


