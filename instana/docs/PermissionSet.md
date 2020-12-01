# PermissionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Permissions** | **[]string** |  | 
**ApplicationIds** | **[]string** |  | 
**KubernetesClusterUUIDs** | **[]string** |  | 
**KubernetesNamespaceUIDs** | **[]string** |  | 
**WebsiteIds** | **[]string** |  | 
**MobileAppIds** | **[]string** |  | 
**InfraDfqFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewPermissionSet

`func NewPermissionSet(name string, permissions []string, applicationIds []string, kubernetesClusterUUIDs []string, kubernetesNamespaceUIDs []string, websiteIds []string, mobileAppIds []string, ) *PermissionSet`

NewPermissionSet instantiates a new PermissionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSetWithDefaults

`func NewPermissionSetWithDefaults() *PermissionSet`

NewPermissionSetWithDefaults instantiates a new PermissionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PermissionSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionSet) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *PermissionSet) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionSet) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionSet) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetApplicationIds

`func (o *PermissionSet) GetApplicationIds() []string`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *PermissionSet) GetApplicationIdsOk() (*[]string, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *PermissionSet) SetApplicationIds(v []string)`

SetApplicationIds sets ApplicationIds field to given value.


### GetKubernetesClusterUUIDs

`func (o *PermissionSet) GetKubernetesClusterUUIDs() []string`

GetKubernetesClusterUUIDs returns the KubernetesClusterUUIDs field if non-nil, zero value otherwise.

### GetKubernetesClusterUUIDsOk

`func (o *PermissionSet) GetKubernetesClusterUUIDsOk() (*[]string, bool)`

GetKubernetesClusterUUIDsOk returns a tuple with the KubernetesClusterUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterUUIDs

`func (o *PermissionSet) SetKubernetesClusterUUIDs(v []string)`

SetKubernetesClusterUUIDs sets KubernetesClusterUUIDs field to given value.


### GetKubernetesNamespaceUIDs

`func (o *PermissionSet) GetKubernetesNamespaceUIDs() []string`

GetKubernetesNamespaceUIDs returns the KubernetesNamespaceUIDs field if non-nil, zero value otherwise.

### GetKubernetesNamespaceUIDsOk

`func (o *PermissionSet) GetKubernetesNamespaceUIDsOk() (*[]string, bool)`

GetKubernetesNamespaceUIDsOk returns a tuple with the KubernetesNamespaceUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespaceUIDs

`func (o *PermissionSet) SetKubernetesNamespaceUIDs(v []string)`

SetKubernetesNamespaceUIDs sets KubernetesNamespaceUIDs field to given value.


### GetWebsiteIds

`func (o *PermissionSet) GetWebsiteIds() []string`

GetWebsiteIds returns the WebsiteIds field if non-nil, zero value otherwise.

### GetWebsiteIdsOk

`func (o *PermissionSet) GetWebsiteIdsOk() (*[]string, bool)`

GetWebsiteIdsOk returns a tuple with the WebsiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteIds

`func (o *PermissionSet) SetWebsiteIds(v []string)`

SetWebsiteIds sets WebsiteIds field to given value.


### GetMobileAppIds

`func (o *PermissionSet) GetMobileAppIds() []string`

GetMobileAppIds returns the MobileAppIds field if non-nil, zero value otherwise.

### GetMobileAppIdsOk

`func (o *PermissionSet) GetMobileAppIdsOk() (*[]string, bool)`

GetMobileAppIdsOk returns a tuple with the MobileAppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppIds

`func (o *PermissionSet) SetMobileAppIds(v []string)`

SetMobileAppIds sets MobileAppIds field to given value.


### GetInfraDfqFilter

`func (o *PermissionSet) GetInfraDfqFilter() string`

GetInfraDfqFilter returns the InfraDfqFilter field if non-nil, zero value otherwise.

### GetInfraDfqFilterOk

`func (o *PermissionSet) GetInfraDfqFilterOk() (*string, bool)`

GetInfraDfqFilterOk returns a tuple with the InfraDfqFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraDfqFilter

`func (o *PermissionSet) SetInfraDfqFilter(v string)`

SetInfraDfqFilter sets InfraDfqFilter field to given value.

### HasInfraDfqFilter

`func (o *PermissionSet) HasInfraDfqFilter() bool`

HasInfraDfqFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


