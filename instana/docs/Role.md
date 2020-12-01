# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CanConfigureServiceMapping** | Pointer to **bool** |  | [optional] 
**CanConfigureEumApplications** | Pointer to **bool** |  | [optional] 
**CanConfigureMobileAppMonitoring** | Pointer to **bool** |  | [optional] 
**CanConfigureUsers** | Pointer to **bool** |  | [optional] 
**CanInstallNewAgents** | Pointer to **bool** |  | [optional] 
**CanSeeUsageInformation** | Pointer to **bool** |  | [optional] 
**CanConfigureIntegrations** | Pointer to **bool** |  | [optional] 
**CanSeeOnPremLicenseInformation** | Pointer to **bool** |  | [optional] 
**CanConfigureRoles** | Pointer to **bool** |  | [optional] 
**CanConfigureCustomAlerts** | Pointer to **bool** |  | [optional] 
**CanConfigureApiTokens** | Pointer to **bool** |  | [optional] 
**CanConfigureAgentRunMode** | Pointer to **bool** |  | [optional] 
**CanViewAuditLog** | Pointer to **bool** |  | [optional] 
**CanConfigureObjectives** | Pointer to **bool** |  | [optional] 
**CanConfigureAgents** | Pointer to **bool** |  | [optional] 
**CanConfigureAuthenticationMethods** | Pointer to **bool** |  | [optional] 
**CanConfigureApplications** | Pointer to **bool** |  | [optional] 
**CanConfigureTeams** | Pointer to **bool** |  | [optional] 
**RestrictedAccess** | Pointer to **bool** |  | [optional] 
**CanConfigureReleases** | Pointer to **bool** |  | [optional] 
**CanConfigureLogManagement** | Pointer to **bool** |  | [optional] 
**CanCreatePublicCustomDashboards** | Pointer to **bool** |  | [optional] 
**CanViewLogs** | Pointer to **bool** |  | [optional] 
**CanViewTraceDetails** | Pointer to **bool** |  | [optional] 
**CanConfigureSessionSettings** | Pointer to **bool** |  | [optional] 
**CanConfigureServiceLevelIndicators** | Pointer to **bool** |  | [optional] 
**CanConfigureGlobalAlertPayload** | Pointer to **bool** |  | [optional] 
**CanViewAccountAndBillingInformation** | Pointer to **bool** |  | [optional] 

## Methods

### NewRole

`func NewRole(id string, name string, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetCanConfigureServiceMapping

`func (o *Role) GetCanConfigureServiceMapping() bool`

GetCanConfigureServiceMapping returns the CanConfigureServiceMapping field if non-nil, zero value otherwise.

### GetCanConfigureServiceMappingOk

`func (o *Role) GetCanConfigureServiceMappingOk() (*bool, bool)`

GetCanConfigureServiceMappingOk returns a tuple with the CanConfigureServiceMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureServiceMapping

`func (o *Role) SetCanConfigureServiceMapping(v bool)`

SetCanConfigureServiceMapping sets CanConfigureServiceMapping field to given value.

### HasCanConfigureServiceMapping

`func (o *Role) HasCanConfigureServiceMapping() bool`

HasCanConfigureServiceMapping returns a boolean if a field has been set.

### GetCanConfigureEumApplications

`func (o *Role) GetCanConfigureEumApplications() bool`

GetCanConfigureEumApplications returns the CanConfigureEumApplications field if non-nil, zero value otherwise.

### GetCanConfigureEumApplicationsOk

`func (o *Role) GetCanConfigureEumApplicationsOk() (*bool, bool)`

GetCanConfigureEumApplicationsOk returns a tuple with the CanConfigureEumApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureEumApplications

`func (o *Role) SetCanConfigureEumApplications(v bool)`

SetCanConfigureEumApplications sets CanConfigureEumApplications field to given value.

### HasCanConfigureEumApplications

`func (o *Role) HasCanConfigureEumApplications() bool`

HasCanConfigureEumApplications returns a boolean if a field has been set.

### GetCanConfigureMobileAppMonitoring

`func (o *Role) GetCanConfigureMobileAppMonitoring() bool`

GetCanConfigureMobileAppMonitoring returns the CanConfigureMobileAppMonitoring field if non-nil, zero value otherwise.

### GetCanConfigureMobileAppMonitoringOk

`func (o *Role) GetCanConfigureMobileAppMonitoringOk() (*bool, bool)`

GetCanConfigureMobileAppMonitoringOk returns a tuple with the CanConfigureMobileAppMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureMobileAppMonitoring

`func (o *Role) SetCanConfigureMobileAppMonitoring(v bool)`

SetCanConfigureMobileAppMonitoring sets CanConfigureMobileAppMonitoring field to given value.

### HasCanConfigureMobileAppMonitoring

`func (o *Role) HasCanConfigureMobileAppMonitoring() bool`

HasCanConfigureMobileAppMonitoring returns a boolean if a field has been set.

### GetCanConfigureUsers

`func (o *Role) GetCanConfigureUsers() bool`

GetCanConfigureUsers returns the CanConfigureUsers field if non-nil, zero value otherwise.

### GetCanConfigureUsersOk

`func (o *Role) GetCanConfigureUsersOk() (*bool, bool)`

GetCanConfigureUsersOk returns a tuple with the CanConfigureUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureUsers

`func (o *Role) SetCanConfigureUsers(v bool)`

SetCanConfigureUsers sets CanConfigureUsers field to given value.

### HasCanConfigureUsers

`func (o *Role) HasCanConfigureUsers() bool`

HasCanConfigureUsers returns a boolean if a field has been set.

### GetCanInstallNewAgents

`func (o *Role) GetCanInstallNewAgents() bool`

GetCanInstallNewAgents returns the CanInstallNewAgents field if non-nil, zero value otherwise.

### GetCanInstallNewAgentsOk

`func (o *Role) GetCanInstallNewAgentsOk() (*bool, bool)`

GetCanInstallNewAgentsOk returns a tuple with the CanInstallNewAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInstallNewAgents

`func (o *Role) SetCanInstallNewAgents(v bool)`

SetCanInstallNewAgents sets CanInstallNewAgents field to given value.

### HasCanInstallNewAgents

`func (o *Role) HasCanInstallNewAgents() bool`

HasCanInstallNewAgents returns a boolean if a field has been set.

### GetCanSeeUsageInformation

`func (o *Role) GetCanSeeUsageInformation() bool`

GetCanSeeUsageInformation returns the CanSeeUsageInformation field if non-nil, zero value otherwise.

### GetCanSeeUsageInformationOk

`func (o *Role) GetCanSeeUsageInformationOk() (*bool, bool)`

GetCanSeeUsageInformationOk returns a tuple with the CanSeeUsageInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeeUsageInformation

`func (o *Role) SetCanSeeUsageInformation(v bool)`

SetCanSeeUsageInformation sets CanSeeUsageInformation field to given value.

### HasCanSeeUsageInformation

`func (o *Role) HasCanSeeUsageInformation() bool`

HasCanSeeUsageInformation returns a boolean if a field has been set.

### GetCanConfigureIntegrations

`func (o *Role) GetCanConfigureIntegrations() bool`

GetCanConfigureIntegrations returns the CanConfigureIntegrations field if non-nil, zero value otherwise.

### GetCanConfigureIntegrationsOk

`func (o *Role) GetCanConfigureIntegrationsOk() (*bool, bool)`

GetCanConfigureIntegrationsOk returns a tuple with the CanConfigureIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureIntegrations

`func (o *Role) SetCanConfigureIntegrations(v bool)`

SetCanConfigureIntegrations sets CanConfigureIntegrations field to given value.

### HasCanConfigureIntegrations

`func (o *Role) HasCanConfigureIntegrations() bool`

HasCanConfigureIntegrations returns a boolean if a field has been set.

### GetCanSeeOnPremLicenseInformation

`func (o *Role) GetCanSeeOnPremLicenseInformation() bool`

GetCanSeeOnPremLicenseInformation returns the CanSeeOnPremLicenseInformation field if non-nil, zero value otherwise.

### GetCanSeeOnPremLicenseInformationOk

`func (o *Role) GetCanSeeOnPremLicenseInformationOk() (*bool, bool)`

GetCanSeeOnPremLicenseInformationOk returns a tuple with the CanSeeOnPremLicenseInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeeOnPremLicenseInformation

`func (o *Role) SetCanSeeOnPremLicenseInformation(v bool)`

SetCanSeeOnPremLicenseInformation sets CanSeeOnPremLicenseInformation field to given value.

### HasCanSeeOnPremLicenseInformation

`func (o *Role) HasCanSeeOnPremLicenseInformation() bool`

HasCanSeeOnPremLicenseInformation returns a boolean if a field has been set.

### GetCanConfigureRoles

`func (o *Role) GetCanConfigureRoles() bool`

GetCanConfigureRoles returns the CanConfigureRoles field if non-nil, zero value otherwise.

### GetCanConfigureRolesOk

`func (o *Role) GetCanConfigureRolesOk() (*bool, bool)`

GetCanConfigureRolesOk returns a tuple with the CanConfigureRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureRoles

`func (o *Role) SetCanConfigureRoles(v bool)`

SetCanConfigureRoles sets CanConfigureRoles field to given value.

### HasCanConfigureRoles

`func (o *Role) HasCanConfigureRoles() bool`

HasCanConfigureRoles returns a boolean if a field has been set.

### GetCanConfigureCustomAlerts

`func (o *Role) GetCanConfigureCustomAlerts() bool`

GetCanConfigureCustomAlerts returns the CanConfigureCustomAlerts field if non-nil, zero value otherwise.

### GetCanConfigureCustomAlertsOk

`func (o *Role) GetCanConfigureCustomAlertsOk() (*bool, bool)`

GetCanConfigureCustomAlertsOk returns a tuple with the CanConfigureCustomAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureCustomAlerts

`func (o *Role) SetCanConfigureCustomAlerts(v bool)`

SetCanConfigureCustomAlerts sets CanConfigureCustomAlerts field to given value.

### HasCanConfigureCustomAlerts

`func (o *Role) HasCanConfigureCustomAlerts() bool`

HasCanConfigureCustomAlerts returns a boolean if a field has been set.

### GetCanConfigureApiTokens

`func (o *Role) GetCanConfigureApiTokens() bool`

GetCanConfigureApiTokens returns the CanConfigureApiTokens field if non-nil, zero value otherwise.

### GetCanConfigureApiTokensOk

`func (o *Role) GetCanConfigureApiTokensOk() (*bool, bool)`

GetCanConfigureApiTokensOk returns a tuple with the CanConfigureApiTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureApiTokens

`func (o *Role) SetCanConfigureApiTokens(v bool)`

SetCanConfigureApiTokens sets CanConfigureApiTokens field to given value.

### HasCanConfigureApiTokens

`func (o *Role) HasCanConfigureApiTokens() bool`

HasCanConfigureApiTokens returns a boolean if a field has been set.

### GetCanConfigureAgentRunMode

`func (o *Role) GetCanConfigureAgentRunMode() bool`

GetCanConfigureAgentRunMode returns the CanConfigureAgentRunMode field if non-nil, zero value otherwise.

### GetCanConfigureAgentRunModeOk

`func (o *Role) GetCanConfigureAgentRunModeOk() (*bool, bool)`

GetCanConfigureAgentRunModeOk returns a tuple with the CanConfigureAgentRunMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureAgentRunMode

`func (o *Role) SetCanConfigureAgentRunMode(v bool)`

SetCanConfigureAgentRunMode sets CanConfigureAgentRunMode field to given value.

### HasCanConfigureAgentRunMode

`func (o *Role) HasCanConfigureAgentRunMode() bool`

HasCanConfigureAgentRunMode returns a boolean if a field has been set.

### GetCanViewAuditLog

`func (o *Role) GetCanViewAuditLog() bool`

GetCanViewAuditLog returns the CanViewAuditLog field if non-nil, zero value otherwise.

### GetCanViewAuditLogOk

`func (o *Role) GetCanViewAuditLogOk() (*bool, bool)`

GetCanViewAuditLogOk returns a tuple with the CanViewAuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewAuditLog

`func (o *Role) SetCanViewAuditLog(v bool)`

SetCanViewAuditLog sets CanViewAuditLog field to given value.

### HasCanViewAuditLog

`func (o *Role) HasCanViewAuditLog() bool`

HasCanViewAuditLog returns a boolean if a field has been set.

### GetCanConfigureObjectives

`func (o *Role) GetCanConfigureObjectives() bool`

GetCanConfigureObjectives returns the CanConfigureObjectives field if non-nil, zero value otherwise.

### GetCanConfigureObjectivesOk

`func (o *Role) GetCanConfigureObjectivesOk() (*bool, bool)`

GetCanConfigureObjectivesOk returns a tuple with the CanConfigureObjectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureObjectives

`func (o *Role) SetCanConfigureObjectives(v bool)`

SetCanConfigureObjectives sets CanConfigureObjectives field to given value.

### HasCanConfigureObjectives

`func (o *Role) HasCanConfigureObjectives() bool`

HasCanConfigureObjectives returns a boolean if a field has been set.

### GetCanConfigureAgents

`func (o *Role) GetCanConfigureAgents() bool`

GetCanConfigureAgents returns the CanConfigureAgents field if non-nil, zero value otherwise.

### GetCanConfigureAgentsOk

`func (o *Role) GetCanConfigureAgentsOk() (*bool, bool)`

GetCanConfigureAgentsOk returns a tuple with the CanConfigureAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureAgents

`func (o *Role) SetCanConfigureAgents(v bool)`

SetCanConfigureAgents sets CanConfigureAgents field to given value.

### HasCanConfigureAgents

`func (o *Role) HasCanConfigureAgents() bool`

HasCanConfigureAgents returns a boolean if a field has been set.

### GetCanConfigureAuthenticationMethods

`func (o *Role) GetCanConfigureAuthenticationMethods() bool`

GetCanConfigureAuthenticationMethods returns the CanConfigureAuthenticationMethods field if non-nil, zero value otherwise.

### GetCanConfigureAuthenticationMethodsOk

`func (o *Role) GetCanConfigureAuthenticationMethodsOk() (*bool, bool)`

GetCanConfigureAuthenticationMethodsOk returns a tuple with the CanConfigureAuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureAuthenticationMethods

`func (o *Role) SetCanConfigureAuthenticationMethods(v bool)`

SetCanConfigureAuthenticationMethods sets CanConfigureAuthenticationMethods field to given value.

### HasCanConfigureAuthenticationMethods

`func (o *Role) HasCanConfigureAuthenticationMethods() bool`

HasCanConfigureAuthenticationMethods returns a boolean if a field has been set.

### GetCanConfigureApplications

`func (o *Role) GetCanConfigureApplications() bool`

GetCanConfigureApplications returns the CanConfigureApplications field if non-nil, zero value otherwise.

### GetCanConfigureApplicationsOk

`func (o *Role) GetCanConfigureApplicationsOk() (*bool, bool)`

GetCanConfigureApplicationsOk returns a tuple with the CanConfigureApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureApplications

`func (o *Role) SetCanConfigureApplications(v bool)`

SetCanConfigureApplications sets CanConfigureApplications field to given value.

### HasCanConfigureApplications

`func (o *Role) HasCanConfigureApplications() bool`

HasCanConfigureApplications returns a boolean if a field has been set.

### GetCanConfigureTeams

`func (o *Role) GetCanConfigureTeams() bool`

GetCanConfigureTeams returns the CanConfigureTeams field if non-nil, zero value otherwise.

### GetCanConfigureTeamsOk

`func (o *Role) GetCanConfigureTeamsOk() (*bool, bool)`

GetCanConfigureTeamsOk returns a tuple with the CanConfigureTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureTeams

`func (o *Role) SetCanConfigureTeams(v bool)`

SetCanConfigureTeams sets CanConfigureTeams field to given value.

### HasCanConfigureTeams

`func (o *Role) HasCanConfigureTeams() bool`

HasCanConfigureTeams returns a boolean if a field has been set.

### GetRestrictedAccess

`func (o *Role) GetRestrictedAccess() bool`

GetRestrictedAccess returns the RestrictedAccess field if non-nil, zero value otherwise.

### GetRestrictedAccessOk

`func (o *Role) GetRestrictedAccessOk() (*bool, bool)`

GetRestrictedAccessOk returns a tuple with the RestrictedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedAccess

`func (o *Role) SetRestrictedAccess(v bool)`

SetRestrictedAccess sets RestrictedAccess field to given value.

### HasRestrictedAccess

`func (o *Role) HasRestrictedAccess() bool`

HasRestrictedAccess returns a boolean if a field has been set.

### GetCanConfigureReleases

`func (o *Role) GetCanConfigureReleases() bool`

GetCanConfigureReleases returns the CanConfigureReleases field if non-nil, zero value otherwise.

### GetCanConfigureReleasesOk

`func (o *Role) GetCanConfigureReleasesOk() (*bool, bool)`

GetCanConfigureReleasesOk returns a tuple with the CanConfigureReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureReleases

`func (o *Role) SetCanConfigureReleases(v bool)`

SetCanConfigureReleases sets CanConfigureReleases field to given value.

### HasCanConfigureReleases

`func (o *Role) HasCanConfigureReleases() bool`

HasCanConfigureReleases returns a boolean if a field has been set.

### GetCanConfigureLogManagement

`func (o *Role) GetCanConfigureLogManagement() bool`

GetCanConfigureLogManagement returns the CanConfigureLogManagement field if non-nil, zero value otherwise.

### GetCanConfigureLogManagementOk

`func (o *Role) GetCanConfigureLogManagementOk() (*bool, bool)`

GetCanConfigureLogManagementOk returns a tuple with the CanConfigureLogManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureLogManagement

`func (o *Role) SetCanConfigureLogManagement(v bool)`

SetCanConfigureLogManagement sets CanConfigureLogManagement field to given value.

### HasCanConfigureLogManagement

`func (o *Role) HasCanConfigureLogManagement() bool`

HasCanConfigureLogManagement returns a boolean if a field has been set.

### GetCanCreatePublicCustomDashboards

`func (o *Role) GetCanCreatePublicCustomDashboards() bool`

GetCanCreatePublicCustomDashboards returns the CanCreatePublicCustomDashboards field if non-nil, zero value otherwise.

### GetCanCreatePublicCustomDashboardsOk

`func (o *Role) GetCanCreatePublicCustomDashboardsOk() (*bool, bool)`

GetCanCreatePublicCustomDashboardsOk returns a tuple with the CanCreatePublicCustomDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreatePublicCustomDashboards

`func (o *Role) SetCanCreatePublicCustomDashboards(v bool)`

SetCanCreatePublicCustomDashboards sets CanCreatePublicCustomDashboards field to given value.

### HasCanCreatePublicCustomDashboards

`func (o *Role) HasCanCreatePublicCustomDashboards() bool`

HasCanCreatePublicCustomDashboards returns a boolean if a field has been set.

### GetCanViewLogs

`func (o *Role) GetCanViewLogs() bool`

GetCanViewLogs returns the CanViewLogs field if non-nil, zero value otherwise.

### GetCanViewLogsOk

`func (o *Role) GetCanViewLogsOk() (*bool, bool)`

GetCanViewLogsOk returns a tuple with the CanViewLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewLogs

`func (o *Role) SetCanViewLogs(v bool)`

SetCanViewLogs sets CanViewLogs field to given value.

### HasCanViewLogs

`func (o *Role) HasCanViewLogs() bool`

HasCanViewLogs returns a boolean if a field has been set.

### GetCanViewTraceDetails

`func (o *Role) GetCanViewTraceDetails() bool`

GetCanViewTraceDetails returns the CanViewTraceDetails field if non-nil, zero value otherwise.

### GetCanViewTraceDetailsOk

`func (o *Role) GetCanViewTraceDetailsOk() (*bool, bool)`

GetCanViewTraceDetailsOk returns a tuple with the CanViewTraceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewTraceDetails

`func (o *Role) SetCanViewTraceDetails(v bool)`

SetCanViewTraceDetails sets CanViewTraceDetails field to given value.

### HasCanViewTraceDetails

`func (o *Role) HasCanViewTraceDetails() bool`

HasCanViewTraceDetails returns a boolean if a field has been set.

### GetCanConfigureSessionSettings

`func (o *Role) GetCanConfigureSessionSettings() bool`

GetCanConfigureSessionSettings returns the CanConfigureSessionSettings field if non-nil, zero value otherwise.

### GetCanConfigureSessionSettingsOk

`func (o *Role) GetCanConfigureSessionSettingsOk() (*bool, bool)`

GetCanConfigureSessionSettingsOk returns a tuple with the CanConfigureSessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureSessionSettings

`func (o *Role) SetCanConfigureSessionSettings(v bool)`

SetCanConfigureSessionSettings sets CanConfigureSessionSettings field to given value.

### HasCanConfigureSessionSettings

`func (o *Role) HasCanConfigureSessionSettings() bool`

HasCanConfigureSessionSettings returns a boolean if a field has been set.

### GetCanConfigureServiceLevelIndicators

`func (o *Role) GetCanConfigureServiceLevelIndicators() bool`

GetCanConfigureServiceLevelIndicators returns the CanConfigureServiceLevelIndicators field if non-nil, zero value otherwise.

### GetCanConfigureServiceLevelIndicatorsOk

`func (o *Role) GetCanConfigureServiceLevelIndicatorsOk() (*bool, bool)`

GetCanConfigureServiceLevelIndicatorsOk returns a tuple with the CanConfigureServiceLevelIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureServiceLevelIndicators

`func (o *Role) SetCanConfigureServiceLevelIndicators(v bool)`

SetCanConfigureServiceLevelIndicators sets CanConfigureServiceLevelIndicators field to given value.

### HasCanConfigureServiceLevelIndicators

`func (o *Role) HasCanConfigureServiceLevelIndicators() bool`

HasCanConfigureServiceLevelIndicators returns a boolean if a field has been set.

### GetCanConfigureGlobalAlertPayload

`func (o *Role) GetCanConfigureGlobalAlertPayload() bool`

GetCanConfigureGlobalAlertPayload returns the CanConfigureGlobalAlertPayload field if non-nil, zero value otherwise.

### GetCanConfigureGlobalAlertPayloadOk

`func (o *Role) GetCanConfigureGlobalAlertPayloadOk() (*bool, bool)`

GetCanConfigureGlobalAlertPayloadOk returns a tuple with the CanConfigureGlobalAlertPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureGlobalAlertPayload

`func (o *Role) SetCanConfigureGlobalAlertPayload(v bool)`

SetCanConfigureGlobalAlertPayload sets CanConfigureGlobalAlertPayload field to given value.

### HasCanConfigureGlobalAlertPayload

`func (o *Role) HasCanConfigureGlobalAlertPayload() bool`

HasCanConfigureGlobalAlertPayload returns a boolean if a field has been set.

### GetCanViewAccountAndBillingInformation

`func (o *Role) GetCanViewAccountAndBillingInformation() bool`

GetCanViewAccountAndBillingInformation returns the CanViewAccountAndBillingInformation field if non-nil, zero value otherwise.

### GetCanViewAccountAndBillingInformationOk

`func (o *Role) GetCanViewAccountAndBillingInformationOk() (*bool, bool)`

GetCanViewAccountAndBillingInformationOk returns a tuple with the CanViewAccountAndBillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewAccountAndBillingInformation

`func (o *Role) SetCanViewAccountAndBillingInformation(v bool)`

SetCanViewAccountAndBillingInformation sets CanViewAccountAndBillingInformation field to given value.

### HasCanViewAccountAndBillingInformation

`func (o *Role) HasCanViewAccountAndBillingInformation() bool`

HasCanViewAccountAndBillingInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


