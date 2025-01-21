# VMUpdateHibernationTimeoutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error11**](Error11.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**VMUpdateHibernationTimeoutResponseAllOfData**](VMUpdateHibernationTimeoutResponseAllOfData.md) |  | [optional] 

## Methods

### NewVMUpdateHibernationTimeoutResponse

`func NewVMUpdateHibernationTimeoutResponse() *VMUpdateHibernationTimeoutResponse`

NewVMUpdateHibernationTimeoutResponse instantiates a new VMUpdateHibernationTimeoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMUpdateHibernationTimeoutResponseWithDefaults

`func NewVMUpdateHibernationTimeoutResponseWithDefaults() *VMUpdateHibernationTimeoutResponse`

NewVMUpdateHibernationTimeoutResponseWithDefaults instantiates a new VMUpdateHibernationTimeoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *VMUpdateHibernationTimeoutResponse) GetErrors() []Error11`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VMUpdateHibernationTimeoutResponse) GetErrorsOk() (*[]Error11, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VMUpdateHibernationTimeoutResponse) SetErrors(v []Error11)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *VMUpdateHibernationTimeoutResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *VMUpdateHibernationTimeoutResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *VMUpdateHibernationTimeoutResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *VMUpdateHibernationTimeoutResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *VMUpdateHibernationTimeoutResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *VMUpdateHibernationTimeoutResponse) GetData() VMUpdateHibernationTimeoutResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VMUpdateHibernationTimeoutResponse) GetDataOk() (*VMUpdateHibernationTimeoutResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VMUpdateHibernationTimeoutResponse) SetData(v VMUpdateHibernationTimeoutResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *VMUpdateHibernationTimeoutResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


