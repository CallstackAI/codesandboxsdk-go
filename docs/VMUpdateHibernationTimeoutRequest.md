# VMUpdateHibernationTimeoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HibernationTimeoutSeconds** | **int32** | The new hibernation timeout in seconds.  Must be greater than 0 and less than or equal to 86400 (24 hours).  | 

## Methods

### NewVMUpdateHibernationTimeoutRequest

`func NewVMUpdateHibernationTimeoutRequest(hibernationTimeoutSeconds int32, ) *VMUpdateHibernationTimeoutRequest`

NewVMUpdateHibernationTimeoutRequest instantiates a new VMUpdateHibernationTimeoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMUpdateHibernationTimeoutRequestWithDefaults

`func NewVMUpdateHibernationTimeoutRequestWithDefaults() *VMUpdateHibernationTimeoutRequest`

NewVMUpdateHibernationTimeoutRequestWithDefaults instantiates a new VMUpdateHibernationTimeoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHibernationTimeoutSeconds

`func (o *VMUpdateHibernationTimeoutRequest) GetHibernationTimeoutSeconds() int32`

GetHibernationTimeoutSeconds returns the HibernationTimeoutSeconds field if non-nil, zero value otherwise.

### GetHibernationTimeoutSecondsOk

`func (o *VMUpdateHibernationTimeoutRequest) GetHibernationTimeoutSecondsOk() (*int32, bool)`

GetHibernationTimeoutSecondsOk returns a tuple with the HibernationTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHibernationTimeoutSeconds

`func (o *VMUpdateHibernationTimeoutRequest) SetHibernationTimeoutSeconds(v int32)`

SetHibernationTimeoutSeconds sets HibernationTimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


