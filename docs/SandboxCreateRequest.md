# SandboxCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional text description of the sandbox. Defaults to no description. | [optional] 
**Entry** | Pointer to **string** | Filename of the entrypoint of the sandbox. | [optional] 
**ExternalResources** | Pointer to **[]string** | Array of strings with external resources to load. | [optional] 
**Files** | [**map[string]SandboxCreateRequestFilesValue**](SandboxCreateRequestFilesValue.md) | Map of &#x60;path &#x3D;&gt; file&#x60; where each file is a map. | 
**IsFrozen** | Pointer to **bool** | Whether changes to the sandbox are disallowed. Defaults to &#x60;false&#x60;. | [optional] [default to false]
**NpmDependencies** | Pointer to **map[string]string** | Map of dependencies and their version specifications. | [optional] 
**Path** | Pointer to **string** | Path to the collection where the new sandbox should be stored. Defaults to \&quot;/\&quot;. If no collection exists at the given path, it will be created. | [optional] [default to "/"]
**Privacy** | Pointer to **int32** | 0 for public, 1 for unlisted, and 2 for private. Privacy is subject to certain restrictions (team minimum setting, drafts must be private, etc.). Defaults to public. | [optional] [default to 0]
**Runtime** | Pointer to **string** | Runtime to use for the sandbox. Defaults to &#x60;\&quot;browser\&quot;&#x60;. | [optional] [default to "browser"]
**Tags** | Pointer to **[]string** | List of string tags to apply to the sandbox. Only the first ten will be used. Defaults to no tags. | [optional] [default to []]
**Template** | Pointer to **string** | Name of the template from which the sandbox is derived (for example, &#x60;\&quot;static\&quot;&#x60;). | [optional] 
**Title** | Pointer to **string** | Sandbox title. Maximum 255 characters. Defaults to no title. | [optional] [default to ""]

## Methods

### NewSandboxCreateRequest

`func NewSandboxCreateRequest(files map[string]SandboxCreateRequestFilesValue, ) *SandboxCreateRequest`

NewSandboxCreateRequest instantiates a new SandboxCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCreateRequestWithDefaults

`func NewSandboxCreateRequestWithDefaults() *SandboxCreateRequest`

NewSandboxCreateRequestWithDefaults instantiates a new SandboxCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SandboxCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SandboxCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SandboxCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SandboxCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntry

`func (o *SandboxCreateRequest) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *SandboxCreateRequest) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *SandboxCreateRequest) SetEntry(v string)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *SandboxCreateRequest) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetExternalResources

`func (o *SandboxCreateRequest) GetExternalResources() []string`

GetExternalResources returns the ExternalResources field if non-nil, zero value otherwise.

### GetExternalResourcesOk

`func (o *SandboxCreateRequest) GetExternalResourcesOk() (*[]string, bool)`

GetExternalResourcesOk returns a tuple with the ExternalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalResources

`func (o *SandboxCreateRequest) SetExternalResources(v []string)`

SetExternalResources sets ExternalResources field to given value.

### HasExternalResources

`func (o *SandboxCreateRequest) HasExternalResources() bool`

HasExternalResources returns a boolean if a field has been set.

### GetFiles

`func (o *SandboxCreateRequest) GetFiles() map[string]SandboxCreateRequestFilesValue`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SandboxCreateRequest) GetFilesOk() (*map[string]SandboxCreateRequestFilesValue, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SandboxCreateRequest) SetFiles(v map[string]SandboxCreateRequestFilesValue)`

SetFiles sets Files field to given value.


### GetIsFrozen

`func (o *SandboxCreateRequest) GetIsFrozen() bool`

GetIsFrozen returns the IsFrozen field if non-nil, zero value otherwise.

### GetIsFrozenOk

`func (o *SandboxCreateRequest) GetIsFrozenOk() (*bool, bool)`

GetIsFrozenOk returns a tuple with the IsFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFrozen

`func (o *SandboxCreateRequest) SetIsFrozen(v bool)`

SetIsFrozen sets IsFrozen field to given value.

### HasIsFrozen

`func (o *SandboxCreateRequest) HasIsFrozen() bool`

HasIsFrozen returns a boolean if a field has been set.

### GetNpmDependencies

`func (o *SandboxCreateRequest) GetNpmDependencies() map[string]string`

GetNpmDependencies returns the NpmDependencies field if non-nil, zero value otherwise.

### GetNpmDependenciesOk

`func (o *SandboxCreateRequest) GetNpmDependenciesOk() (*map[string]string, bool)`

GetNpmDependenciesOk returns a tuple with the NpmDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmDependencies

`func (o *SandboxCreateRequest) SetNpmDependencies(v map[string]string)`

SetNpmDependencies sets NpmDependencies field to given value.

### HasNpmDependencies

`func (o *SandboxCreateRequest) HasNpmDependencies() bool`

HasNpmDependencies returns a boolean if a field has been set.

### GetPath

`func (o *SandboxCreateRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SandboxCreateRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SandboxCreateRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SandboxCreateRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrivacy

`func (o *SandboxCreateRequest) GetPrivacy() int32`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *SandboxCreateRequest) GetPrivacyOk() (*int32, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *SandboxCreateRequest) SetPrivacy(v int32)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *SandboxCreateRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetRuntime

`func (o *SandboxCreateRequest) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SandboxCreateRequest) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SandboxCreateRequest) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *SandboxCreateRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTags

`func (o *SandboxCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SandboxCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SandboxCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SandboxCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *SandboxCreateRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SandboxCreateRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SandboxCreateRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SandboxCreateRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *SandboxCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SandboxCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SandboxCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SandboxCreateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


