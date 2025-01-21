# WorkspaceCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error13**](Error13.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**WorkspaceCreateResponseAllOfData**](WorkspaceCreateResponseAllOfData.md) |  | [optional] 

## Methods

### NewWorkspaceCreateResponse

`func NewWorkspaceCreateResponse() *WorkspaceCreateResponse`

NewWorkspaceCreateResponse instantiates a new WorkspaceCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceCreateResponseWithDefaults

`func NewWorkspaceCreateResponseWithDefaults() *WorkspaceCreateResponse`

NewWorkspaceCreateResponseWithDefaults instantiates a new WorkspaceCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *WorkspaceCreateResponse) GetErrors() []Error13`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WorkspaceCreateResponse) GetErrorsOk() (*[]Error13, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WorkspaceCreateResponse) SetErrors(v []Error13)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WorkspaceCreateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *WorkspaceCreateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WorkspaceCreateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WorkspaceCreateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *WorkspaceCreateResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *WorkspaceCreateResponse) GetData() WorkspaceCreateResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WorkspaceCreateResponse) GetDataOk() (*WorkspaceCreateResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WorkspaceCreateResponse) SetData(v WorkspaceCreateResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *WorkspaceCreateResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


