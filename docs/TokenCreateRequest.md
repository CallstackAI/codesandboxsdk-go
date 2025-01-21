# TokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersion** | Pointer to **NullableString** | API Version to use, formatted as YYYY-MM-DD. Defaults to the latest version at time of creation. | [optional] 
**Description** | Pointer to **string** | Description of this token, for instance where it will be used. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | UTC Timestamp until when this token is valid. Omitting this field will create a token without an expiry. | [optional] 
**Scopes** | Pointer to **[]string** | Which scopes to grant this token. The given scopes will replace the current scopes, revoking any that are no longer present in the list. | [optional] 

## Methods

### NewTokenCreateRequest

`func NewTokenCreateRequest() *TokenCreateRequest`

NewTokenCreateRequest instantiates a new TokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateRequestWithDefaults

`func NewTokenCreateRequestWithDefaults() *TokenCreateRequest`

NewTokenCreateRequestWithDefaults instantiates a new TokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersion

`func (o *TokenCreateRequest) GetDefaultVersion() string`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *TokenCreateRequest) GetDefaultVersionOk() (*string, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *TokenCreateRequest) SetDefaultVersion(v string)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *TokenCreateRequest) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### SetDefaultVersionNil

`func (o *TokenCreateRequest) SetDefaultVersionNil(b bool)`

 SetDefaultVersionNil sets the value for DefaultVersion to be an explicit nil

### UnsetDefaultVersion
`func (o *TokenCreateRequest) UnsetDefaultVersion()`

UnsetDefaultVersion ensures that no value is present for DefaultVersion, not even an explicit nil
### GetDescription

`func (o *TokenCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenCreateRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenCreateRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenCreateRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenCreateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *TokenCreateRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *TokenCreateRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetScopes

`func (o *TokenCreateRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenCreateRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenCreateRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenCreateRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


