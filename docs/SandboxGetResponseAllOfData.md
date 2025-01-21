# SandboxGetResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**IsFrozen** | **bool** |  | 
**Privacy** | **int32** |  | 
**Tags** | **[]string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewSandboxGetResponseAllOfData

`func NewSandboxGetResponseAllOfData(createdAt time.Time, id string, isFrozen bool, privacy int32, tags []string, updatedAt time.Time, ) *SandboxGetResponseAllOfData`

NewSandboxGetResponseAllOfData instantiates a new SandboxGetResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxGetResponseAllOfDataWithDefaults

`func NewSandboxGetResponseAllOfDataWithDefaults() *SandboxGetResponseAllOfData`

NewSandboxGetResponseAllOfDataWithDefaults instantiates a new SandboxGetResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SandboxGetResponseAllOfData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SandboxGetResponseAllOfData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SandboxGetResponseAllOfData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *SandboxGetResponseAllOfData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SandboxGetResponseAllOfData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SandboxGetResponseAllOfData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SandboxGetResponseAllOfData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SandboxGetResponseAllOfData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SandboxGetResponseAllOfData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *SandboxGetResponseAllOfData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxGetResponseAllOfData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxGetResponseAllOfData) SetId(v string)`

SetId sets Id field to given value.


### GetIsFrozen

`func (o *SandboxGetResponseAllOfData) GetIsFrozen() bool`

GetIsFrozen returns the IsFrozen field if non-nil, zero value otherwise.

### GetIsFrozenOk

`func (o *SandboxGetResponseAllOfData) GetIsFrozenOk() (*bool, bool)`

GetIsFrozenOk returns a tuple with the IsFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFrozen

`func (o *SandboxGetResponseAllOfData) SetIsFrozen(v bool)`

SetIsFrozen sets IsFrozen field to given value.


### GetPrivacy

`func (o *SandboxGetResponseAllOfData) GetPrivacy() int32`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *SandboxGetResponseAllOfData) GetPrivacyOk() (*int32, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *SandboxGetResponseAllOfData) SetPrivacy(v int32)`

SetPrivacy sets Privacy field to given value.


### GetTags

`func (o *SandboxGetResponseAllOfData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SandboxGetResponseAllOfData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SandboxGetResponseAllOfData) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTitle

`func (o *SandboxGetResponseAllOfData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SandboxGetResponseAllOfData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SandboxGetResponseAllOfData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SandboxGetResponseAllOfData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SandboxGetResponseAllOfData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SandboxGetResponseAllOfData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUpdatedAt

`func (o *SandboxGetResponseAllOfData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SandboxGetResponseAllOfData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SandboxGetResponseAllOfData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


