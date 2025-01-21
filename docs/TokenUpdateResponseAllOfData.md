# TokenUpdateResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**Scopes** | **[]string** |  | 
**TeamId** | **string** |  | 
**TokenId** | **string** |  | 

## Methods

### NewTokenUpdateResponseAllOfData

`func NewTokenUpdateResponseAllOfData(description NullableString, scopes []string, teamId string, tokenId string, ) *TokenUpdateResponseAllOfData`

NewTokenUpdateResponseAllOfData instantiates a new TokenUpdateResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUpdateResponseAllOfDataWithDefaults

`func NewTokenUpdateResponseAllOfDataWithDefaults() *TokenUpdateResponseAllOfData`

NewTokenUpdateResponseAllOfDataWithDefaults instantiates a new TokenUpdateResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TokenUpdateResponseAllOfData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenUpdateResponseAllOfData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenUpdateResponseAllOfData) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TokenUpdateResponseAllOfData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TokenUpdateResponseAllOfData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpiresAt

`func (o *TokenUpdateResponseAllOfData) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenUpdateResponseAllOfData) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenUpdateResponseAllOfData) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenUpdateResponseAllOfData) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *TokenUpdateResponseAllOfData) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *TokenUpdateResponseAllOfData) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetScopes

`func (o *TokenUpdateResponseAllOfData) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenUpdateResponseAllOfData) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenUpdateResponseAllOfData) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetTeamId

`func (o *TokenUpdateResponseAllOfData) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TokenUpdateResponseAllOfData) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TokenUpdateResponseAllOfData) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTokenId

`func (o *TokenUpdateResponseAllOfData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenUpdateResponseAllOfData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenUpdateResponseAllOfData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


