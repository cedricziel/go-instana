# AuditLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Entries** | Pointer to [**[]AuditLogEntry**](AuditLogEntry.md) |  | [optional] 

## Methods

### NewAuditLogResponse

`func NewAuditLogResponse() *AuditLogResponse`

NewAuditLogResponse instantiates a new AuditLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogResponseWithDefaults

`func NewAuditLogResponseWithDefaults() *AuditLogResponse`

NewAuditLogResponseWithDefaults instantiates a new AuditLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AuditLogResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuditLogResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuditLogResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AuditLogResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetEntries

`func (o *AuditLogResponse) GetEntries() []AuditLogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AuditLogResponse) GetEntriesOk() (*[]AuditLogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AuditLogResponse) SetEntries(v []AuditLogEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AuditLogResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


