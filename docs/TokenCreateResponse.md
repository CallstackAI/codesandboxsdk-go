# TokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error6**](Error6.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**TokenCreateResponseAllOfData**](TokenCreateResponseAllOfData.md) |  | [optional] 

## Methods

### NewTokenCreateResponse

`func NewTokenCreateResponse() *TokenCreateResponse`

NewTokenCreateResponse instantiates a new TokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateResponseWithDefaults

`func NewTokenCreateResponseWithDefaults() *TokenCreateResponse`

NewTokenCreateResponseWithDefaults instantiates a new TokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *TokenCreateResponse) GetErrors() []Error6`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TokenCreateResponse) GetErrorsOk() (*[]Error6, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TokenCreateResponse) SetErrors(v []Error6)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TokenCreateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *TokenCreateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TokenCreateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TokenCreateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TokenCreateResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *TokenCreateResponse) GetData() TokenCreateResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenCreateResponse) GetDataOk() (*TokenCreateResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenCreateResponse) SetData(v TokenCreateResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *TokenCreateResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


