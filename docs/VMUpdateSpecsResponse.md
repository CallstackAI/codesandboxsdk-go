# VMUpdateSpecsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error12**](Error12.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**VMUpdateSpecsResponseAllOfData**](VMUpdateSpecsResponseAllOfData.md) |  | [optional] 

## Methods

### NewVMUpdateSpecsResponse

`func NewVMUpdateSpecsResponse() *VMUpdateSpecsResponse`

NewVMUpdateSpecsResponse instantiates a new VMUpdateSpecsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMUpdateSpecsResponseWithDefaults

`func NewVMUpdateSpecsResponseWithDefaults() *VMUpdateSpecsResponse`

NewVMUpdateSpecsResponseWithDefaults instantiates a new VMUpdateSpecsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *VMUpdateSpecsResponse) GetErrors() []Error12`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VMUpdateSpecsResponse) GetErrorsOk() (*[]Error12, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VMUpdateSpecsResponse) SetErrors(v []Error12)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *VMUpdateSpecsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *VMUpdateSpecsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *VMUpdateSpecsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *VMUpdateSpecsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *VMUpdateSpecsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *VMUpdateSpecsResponse) GetData() VMUpdateSpecsResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VMUpdateSpecsResponse) GetDataOk() (*VMUpdateSpecsResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VMUpdateSpecsResponse) SetData(v VMUpdateSpecsResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *VMUpdateSpecsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


