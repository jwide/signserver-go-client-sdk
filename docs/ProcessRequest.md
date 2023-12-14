# ProcessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The input data to be processed (i.e. signed). | 
**Encoding** | Pointer to **string** | Choice of additional encoding of the data. | [optional] 
**MetaData** | Pointer to **map[string]string** | Additional request metadata for the worker. | [optional] 

## Methods

### NewProcessRequest

`func NewProcessRequest(data string, ) *ProcessRequest`

NewProcessRequest instantiates a new ProcessRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRequestWithDefaults

`func NewProcessRequestWithDefaults() *ProcessRequest`

NewProcessRequestWithDefaults instantiates a new ProcessRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProcessRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProcessRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProcessRequest) SetData(v string)`

SetData sets Data field to given value.


### GetEncoding

`func (o *ProcessRequest) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ProcessRequest) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ProcessRequest) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *ProcessRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetMetaData

`func (o *ProcessRequest) GetMetaData() map[string]string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ProcessRequest) GetMetaDataOk() (*map[string]string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ProcessRequest) SetMetaData(v map[string]string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ProcessRequest) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


