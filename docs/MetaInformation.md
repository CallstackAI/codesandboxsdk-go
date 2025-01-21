# MetaInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | [**MetaInformationApi**](MetaInformationApi.md) |  | 
**Auth** | Pointer to [**MetaInformationAuth**](MetaInformationAuth.md) |  | [optional] 

## Methods

### NewMetaInformation

`func NewMetaInformation(api MetaInformationApi, ) *MetaInformation`

NewMetaInformation instantiates a new MetaInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaInformationWithDefaults

`func NewMetaInformationWithDefaults() *MetaInformation`

NewMetaInformationWithDefaults instantiates a new MetaInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *MetaInformation) GetApi() MetaInformationApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *MetaInformation) GetApiOk() (*MetaInformationApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *MetaInformation) SetApi(v MetaInformationApi)`

SetApi sets Api field to given value.


### GetAuth

`func (o *MetaInformation) GetAuth() MetaInformationAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *MetaInformation) GetAuthOk() (*MetaInformationAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *MetaInformation) SetAuth(v MetaInformationAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *MetaInformation) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


