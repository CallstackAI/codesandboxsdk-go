# TokenUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error7**](Error7.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**TokenUpdateResponseAllOfData**](TokenUpdateResponseAllOfData.md) |  | [optional] 

## Methods

### NewTokenUpdateResponse

`func NewTokenUpdateResponse() *TokenUpdateResponse`

NewTokenUpdateResponse instantiates a new TokenUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUpdateResponseWithDefaults

`func NewTokenUpdateResponseWithDefaults() *TokenUpdateResponse`

NewTokenUpdateResponseWithDefaults instantiates a new TokenUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *TokenUpdateResponse) GetErrors() []Error7`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TokenUpdateResponse) GetErrorsOk() (*[]Error7, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TokenUpdateResponse) SetErrors(v []Error7)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TokenUpdateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *TokenUpdateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TokenUpdateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TokenUpdateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TokenUpdateResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *TokenUpdateResponse) GetData() TokenUpdateResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenUpdateResponse) GetDataOk() (*TokenUpdateResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenUpdateResponse) SetData(v TokenUpdateResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *TokenUpdateResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


