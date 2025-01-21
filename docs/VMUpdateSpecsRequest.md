# VMUpdateSpecsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | **string** | Determines which specs to update the VM with.  Not all tiers will be available depending on the workspace subscription status, and higher tiers incur higher costs. Please see codesandbox.io/pricing for details on specs and costs.  | 

## Methods

### NewVMUpdateSpecsRequest

`func NewVMUpdateSpecsRequest(tier string, ) *VMUpdateSpecsRequest`

NewVMUpdateSpecsRequest instantiates a new VMUpdateSpecsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMUpdateSpecsRequestWithDefaults

`func NewVMUpdateSpecsRequestWithDefaults() *VMUpdateSpecsRequest`

NewVMUpdateSpecsRequestWithDefaults instantiates a new VMUpdateSpecsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *VMUpdateSpecsRequest) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *VMUpdateSpecsRequest) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *VMUpdateSpecsRequest) SetTier(v string)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


