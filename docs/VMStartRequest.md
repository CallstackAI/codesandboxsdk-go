# VMStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HibernationTimeoutSeconds** | Pointer to **int32** | The time in seconds after which the VM will hibernate due to inactivity. Must be a positive integer between 1 and 86400 (24 hours). Defaults to 300 seconds (5 minutes) if not specified.  | [optional] 
**Ipcountry** | Pointer to **string** | This determines in which cluster, closest to the given country the VM will be started in. The format is ISO-3166-1 alpha-2. If not set, the VM will be started closest to the caller of this API. This will only be applied when a VM is run for the first time, and will only serve as a hint (e.g. if the template of this sandbox runs in EU cluster, this sandbox will also run in the EU cluster). | [optional] 
**Tier** | Pointer to **string** | Determines which specs to start the VM with. If not specified, the VM will start with the default specs for the workspace.  You can only specify a VM tier when starting a VM that is inside your workspace. Specifying a VM tier for someone else&#39;s sandbox will return an error.  Not all tiers will be available depending on the workspace subscription status, and higher tiers incur higher costs. Please see codesandbox.io/pricing for details on specs and costs.  | [optional] 

## Methods

### NewVMStartRequest

`func NewVMStartRequest() *VMStartRequest`

NewVMStartRequest instantiates a new VMStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMStartRequestWithDefaults

`func NewVMStartRequestWithDefaults() *VMStartRequest`

NewVMStartRequestWithDefaults instantiates a new VMStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHibernationTimeoutSeconds

`func (o *VMStartRequest) GetHibernationTimeoutSeconds() int32`

GetHibernationTimeoutSeconds returns the HibernationTimeoutSeconds field if non-nil, zero value otherwise.

### GetHibernationTimeoutSecondsOk

`func (o *VMStartRequest) GetHibernationTimeoutSecondsOk() (*int32, bool)`

GetHibernationTimeoutSecondsOk returns a tuple with the HibernationTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHibernationTimeoutSeconds

`func (o *VMStartRequest) SetHibernationTimeoutSeconds(v int32)`

SetHibernationTimeoutSeconds sets HibernationTimeoutSeconds field to given value.

### HasHibernationTimeoutSeconds

`func (o *VMStartRequest) HasHibernationTimeoutSeconds() bool`

HasHibernationTimeoutSeconds returns a boolean if a field has been set.

### GetIpcountry

`func (o *VMStartRequest) GetIpcountry() string`

GetIpcountry returns the Ipcountry field if non-nil, zero value otherwise.

### GetIpcountryOk

`func (o *VMStartRequest) GetIpcountryOk() (*string, bool)`

GetIpcountryOk returns a tuple with the Ipcountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpcountry

`func (o *VMStartRequest) SetIpcountry(v string)`

SetIpcountry sets Ipcountry field to given value.

### HasIpcountry

`func (o *VMStartRequest) HasIpcountry() bool`

HasIpcountry returns a boolean if a field has been set.

### GetTier

`func (o *VMStartRequest) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *VMStartRequest) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *VMStartRequest) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *VMStartRequest) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


