# ProcessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | The resulting data (i.e the signature) in Base64 encoding | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**ArchiveId** | Pointer to **string** |  | [optional] 
**SignerCertificate** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProcessResponse

`func NewProcessResponse() *ProcessResponse`

NewProcessResponse instantiates a new ProcessResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessResponseWithDefaults

`func NewProcessResponseWithDefaults() *ProcessResponse`

NewProcessResponseWithDefaults instantiates a new ProcessResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProcessResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProcessResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProcessResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ProcessResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ProcessResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ProcessResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetArchiveId

`func (o *ProcessResponse) GetArchiveId() string`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *ProcessResponse) GetArchiveIdOk() (*string, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *ProcessResponse) SetArchiveId(v string)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *ProcessResponse) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetSignerCertificate

`func (o *ProcessResponse) GetSignerCertificate() string`

GetSignerCertificate returns the SignerCertificate field if non-nil, zero value otherwise.

### GetSignerCertificateOk

`func (o *ProcessResponse) GetSignerCertificateOk() (*string, bool)`

GetSignerCertificateOk returns a tuple with the SignerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCertificate

`func (o *ProcessResponse) SetSignerCertificate(v string)`

SetSignerCertificate sets SignerCertificate field to given value.

### HasSignerCertificate

`func (o *ProcessResponse) HasSignerCertificate() bool`

HasSignerCertificate returns a boolean if a field has been set.

### GetMetaData

`func (o *ProcessResponse) GetMetaData() map[string]string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ProcessResponse) GetMetaDataOk() (*map[string]string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ProcessResponse) SetMetaData(v map[string]string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ProcessResponse) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


