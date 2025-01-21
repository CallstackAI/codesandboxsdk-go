# SandboxForkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Sandbox description. Maximum 255 characters. Defaults to description of original sandbox. | [optional] [default to "[Original description]"]
**IsFrozen** | Pointer to **bool** | Sandbox frozen status. When true, edits to the sandbox are restricted. Defaults to frozen status of the original sandbox. | [optional] [default to false]
**Path** | Pointer to **string** | Path to the collection where the new sandbox should be stored. Defaults to \&quot;/\&quot;. If no collection exists at the given path, it will be created. | [optional] [default to "/"]
**Privacy** | Pointer to **int32** | Sandbox privacy. 0 for public, 1 for unlisted, and 2 for private. Subject to the minimum privacy restrictions of the workspace. Defaults to the privacy of the original sandbox. | [optional] [default to 0]
**StartOptions** | Pointer to [**SandboxForkRequestStartOptions**](SandboxForkRequestStartOptions.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags to set on the new sandbox, if any. Will not inherit tags from the source sandbox. | [optional] [default to []]
**Title** | Pointer to **string** | Sandbox title. Maximum 255 characters. Defaults to title of original sandbox with (forked). | [optional] [default to "[Original title]"]

## Methods

### NewSandboxForkRequest

`func NewSandboxForkRequest() *SandboxForkRequest`

NewSandboxForkRequest instantiates a new SandboxForkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxForkRequestWithDefaults

`func NewSandboxForkRequestWithDefaults() *SandboxForkRequest`

NewSandboxForkRequestWithDefaults instantiates a new SandboxForkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SandboxForkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SandboxForkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SandboxForkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SandboxForkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsFrozen

`func (o *SandboxForkRequest) GetIsFrozen() bool`

GetIsFrozen returns the IsFrozen field if non-nil, zero value otherwise.

### GetIsFrozenOk

`func (o *SandboxForkRequest) GetIsFrozenOk() (*bool, bool)`

GetIsFrozenOk returns a tuple with the IsFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFrozen

`func (o *SandboxForkRequest) SetIsFrozen(v bool)`

SetIsFrozen sets IsFrozen field to given value.

### HasIsFrozen

`func (o *SandboxForkRequest) HasIsFrozen() bool`

HasIsFrozen returns a boolean if a field has been set.

### GetPath

`func (o *SandboxForkRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SandboxForkRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SandboxForkRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SandboxForkRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrivacy

`func (o *SandboxForkRequest) GetPrivacy() int32`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *SandboxForkRequest) GetPrivacyOk() (*int32, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *SandboxForkRequest) SetPrivacy(v int32)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *SandboxForkRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetStartOptions

`func (o *SandboxForkRequest) GetStartOptions() SandboxForkRequestStartOptions`

GetStartOptions returns the StartOptions field if non-nil, zero value otherwise.

### GetStartOptionsOk

`func (o *SandboxForkRequest) GetStartOptionsOk() (*SandboxForkRequestStartOptions, bool)`

GetStartOptionsOk returns a tuple with the StartOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOptions

`func (o *SandboxForkRequest) SetStartOptions(v SandboxForkRequestStartOptions)`

SetStartOptions sets StartOptions field to given value.

### HasStartOptions

`func (o *SandboxForkRequest) HasStartOptions() bool`

HasStartOptions returns a boolean if a field has been set.

### GetTags

`func (o *SandboxForkRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SandboxForkRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SandboxForkRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SandboxForkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *SandboxForkRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SandboxForkRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SandboxForkRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SandboxForkRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


