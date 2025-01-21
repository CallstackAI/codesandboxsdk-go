# TokenUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersion** | Pointer to **NullableString** | API Version to use, formatted as YYYY-MM-DD | [optional] 
**Description** | Pointer to **NullableString** | Description of this token, for instance where it will be used. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | Expiry time of this token. Cannot be set in the past, and cannot be changed for tokens that have already expired. Pass null to make this token never expire. | [optional] 
**Scopes** | Pointer to **[]string** | Which scopes to grant this token. The given scopes will replace the current scopes, revoking any that are no longer present in the list. | [optional] 

## Methods

### NewTokenUpdateRequest

`func NewTokenUpdateRequest() *TokenUpdateRequest`

NewTokenUpdateRequest instantiates a new TokenUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUpdateRequestWithDefaults

`func NewTokenUpdateRequestWithDefaults() *TokenUpdateRequest`

NewTokenUpdateRequestWithDefaults instantiates a new TokenUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersion

`func (o *TokenUpdateRequest) GetDefaultVersion() string`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *TokenUpdateRequest) GetDefaultVersionOk() (*string, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *TokenUpdateRequest) SetDefaultVersion(v string)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *TokenUpdateRequest) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### SetDefaultVersionNil

`func (o *TokenUpdateRequest) SetDefaultVersionNil(b bool)`

 SetDefaultVersionNil sets the value for DefaultVersion to be an explicit nil

### UnsetDefaultVersion
`func (o *TokenUpdateRequest) UnsetDefaultVersion()`

UnsetDefaultVersion ensures that no value is present for DefaultVersion, not even an explicit nil
### GetDescription

`func (o *TokenUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TokenUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TokenUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpiresAt

`func (o *TokenUpdateRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenUpdateRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenUpdateRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenUpdateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *TokenUpdateRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *TokenUpdateRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetScopes

`func (o *TokenUpdateRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenUpdateRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenUpdateRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenUpdateRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


