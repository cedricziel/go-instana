# Go API client for instana

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.192.86
- Package version: 1.192.86
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [http://instana.com](http://instana.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./instana"
```

## Documentation for API Endpoints

All URIs are relative to *https://unit-tenant.instana.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APITokenApi* | [**DeleteApiToken**](docs/APITokenApi.md#deleteapitoken) | **Delete** /api/settings/api-tokens/{internalId} | Delete API token
*APITokenApi* | [**GetApiToken**](docs/APITokenApi.md#getapitoken) | **Get** /api/settings/api-tokens/{internalId} | API token
*APITokenApi* | [**GetApiTokens**](docs/APITokenApi.md#getapitokens) | **Get** /api/settings/api-tokens | All API tokens
*APITokenApi* | [**PostApiToken**](docs/APITokenApi.md#postapitoken) | **Post** /api/settings/api-tokens | Create an API token
*APITokenApi* | [**PutApiToken**](docs/APITokenApi.md#putapitoken) | **Put** /api/settings/api-tokens/{internalId} | Create or update an API token
*ApplicationAnalyzeApi* | [**GetCallGroup**](docs/ApplicationAnalyzeApi.md#getcallgroup) | **Post** /api/application-monitoring/analyze/call-groups | Get grouped call metrics
*ApplicationAnalyzeApi* | [**GetTrace**](docs/ApplicationAnalyzeApi.md#gettrace) | **Get** /api/application-monitoring/analyze/traces/{id} | Get trace detail
*ApplicationAnalyzeApi* | [**GetTraceGroups**](docs/ApplicationAnalyzeApi.md#gettracegroups) | **Post** /api/application-monitoring/analyze/trace-groups | Get grouped trace metrics
*ApplicationAnalyzeApi* | [**GetTraces**](docs/ApplicationAnalyzeApi.md#gettraces) | **Post** /api/application-monitoring/analyze/traces | Get all traces
*ApplicationCatalogApi* | [**GetApplicationCatalogMetrics**](docs/ApplicationCatalogApi.md#getapplicationcatalogmetrics) | **Get** /api/application-monitoring/catalog/metrics | Get Metric catalog
*ApplicationCatalogApi* | [**GetApplicationTagCatalog**](docs/ApplicationCatalogApi.md#getapplicationtagcatalog) | **Get** /api/application-monitoring/catalog | Get application tag catalog
*ApplicationCatalogApi* | [**GetApplicationTags**](docs/ApplicationCatalogApi.md#getapplicationtags) | **Get** /api/application-monitoring/catalog/tags | Get application tags
*ApplicationMetricsApi* | [**GetApplicationMetrics**](docs/ApplicationMetricsApi.md#getapplicationmetrics) | **Post** /api/application-monitoring/metrics/applications | Get Application Metrics
*ApplicationMetricsApi* | [**GetEndpointsMetrics**](docs/ApplicationMetricsApi.md#getendpointsmetrics) | **Post** /api/application-monitoring/metrics/endpoints | Get Endpoint metrics
*ApplicationMetricsApi* | [**GetServicesMetrics**](docs/ApplicationMetricsApi.md#getservicesmetrics) | **Post** /api/application-monitoring/metrics/services | Get Service metrics
*ApplicationResourcesApi* | [**ApplicationResourcesEndpoints**](docs/ApplicationResourcesApi.md#applicationresourcesendpoints) | **Get** /api/application-monitoring/applications/services/endpoints | Get endpoints
*ApplicationResourcesApi* | [**GetApplicationServices**](docs/ApplicationResourcesApi.md#getapplicationservices) | **Get** /api/application-monitoring/applications/services | Get applications/services
*ApplicationResourcesApi* | [**GetApplications**](docs/ApplicationResourcesApi.md#getapplications) | **Get** /api/application-monitoring/applications | Get applications
*ApplicationResourcesApi* | [**GetServices**](docs/ApplicationResourcesApi.md#getservices) | **Get** /api/application-monitoring/services | Get services
*ApplicationSettingsApi* | [**AddApplicationConfig**](docs/ApplicationSettingsApi.md#addapplicationconfig) | **Post** /api/application-monitoring/settings/application | Add application configuration
*ApplicationSettingsApi* | [**AddServiceConfig**](docs/ApplicationSettingsApi.md#addserviceconfig) | **Post** /api/application-monitoring/settings/service | Add service configuration
*ApplicationSettingsApi* | [**CreateEndpointConfig**](docs/ApplicationSettingsApi.md#createendpointconfig) | **Post** /api/application-monitoring/settings/http-endpoint | Create endpoint configuration
*ApplicationSettingsApi* | [**DeleteApplicationConfig**](docs/ApplicationSettingsApi.md#deleteapplicationconfig) | **Delete** /api/application-monitoring/settings/application/{id} | Delete application configuration
*ApplicationSettingsApi* | [**DeleteEndpointConfig**](docs/ApplicationSettingsApi.md#deleteendpointconfig) | **Delete** /api/application-monitoring/settings/http-endpoint/{id} | Delete endpoint configuration
*ApplicationSettingsApi* | [**DeleteServiceConfig**](docs/ApplicationSettingsApi.md#deleteserviceconfig) | **Delete** /api/application-monitoring/settings/service/{id} | Delete service configuration
*ApplicationSettingsApi* | [**GetApplicationConfig**](docs/ApplicationSettingsApi.md#getapplicationconfig) | **Get** /api/application-monitoring/settings/application/{id} | Application configuration
*ApplicationSettingsApi* | [**GetApplicationConfigs**](docs/ApplicationSettingsApi.md#getapplicationconfigs) | **Get** /api/application-monitoring/settings/application | All Application configurations
*ApplicationSettingsApi* | [**GetEndpointConfig**](docs/ApplicationSettingsApi.md#getendpointconfig) | **Get** /api/application-monitoring/settings/http-endpoint/{id} | Endpoint configuration
*ApplicationSettingsApi* | [**GetEndpointConfigs**](docs/ApplicationSettingsApi.md#getendpointconfigs) | **Get** /api/application-monitoring/settings/http-endpoint | All Endpoint configurations
*ApplicationSettingsApi* | [**GetServiceConfig**](docs/ApplicationSettingsApi.md#getserviceconfig) | **Get** /api/application-monitoring/settings/service/{id} | Service configuration
*ApplicationSettingsApi* | [**GetServiceConfigs**](docs/ApplicationSettingsApi.md#getserviceconfigs) | **Get** /api/application-monitoring/settings/service | All service configurations
*ApplicationSettingsApi* | [**OrderServiceConfig**](docs/ApplicationSettingsApi.md#orderserviceconfig) | **Put** /api/application-monitoring/settings/service/order | Order of service configuration
*ApplicationSettingsApi* | [**PutApplicationConfig**](docs/ApplicationSettingsApi.md#putapplicationconfig) | **Put** /api/application-monitoring/settings/application/{id} | Update application configuration
*ApplicationSettingsApi* | [**PutServiceConfig**](docs/ApplicationSettingsApi.md#putserviceconfig) | **Put** /api/application-monitoring/settings/service/{id} | Update service configuration
*ApplicationSettingsApi* | [**ReplaceAll**](docs/ApplicationSettingsApi.md#replaceall) | **Put** /api/application-monitoring/settings/service | Replace all service configurations
*ApplicationSettingsApi* | [**UpdateEndpointConfig**](docs/ApplicationSettingsApi.md#updateendpointconfig) | **Put** /api/application-monitoring/settings/http-endpoint/{id} | Update endpoint configuration
*AuditLogApi* | [**GetAuditLogs**](docs/AuditLogApi.md#getauditlogs) | **Get** /api/settings/auditlog | Audit log
*CustomDashboardsApi* | [**AddCustomDashboard**](docs/CustomDashboardsApi.md#addcustomdashboard) | **Post** /api/custom-dashboard | Add custom dashboard
*CustomDashboardsApi* | [**DeleteCustomDashboard**](docs/CustomDashboardsApi.md#deletecustomdashboard) | **Delete** /api/custom-dashboard/{customDashboardId} | Remove custom dashboard
*CustomDashboardsApi* | [**GetCustomDashboard**](docs/CustomDashboardsApi.md#getcustomdashboard) | **Get** /api/custom-dashboard/{customDashboardId} | Get custom dashboard
*CustomDashboardsApi* | [**GetCustomDashboards**](docs/CustomDashboardsApi.md#getcustomdashboards) | **Get** /api/custom-dashboard | Get accessible custom dashboards
*CustomDashboardsApi* | [**GetShareableApiTokens**](docs/CustomDashboardsApi.md#getshareableapitokens) | **Get** /api/custom-dashboard/shareable-api-tokens | Get all API tokens.
*CustomDashboardsApi* | [**GetShareableUsers**](docs/CustomDashboardsApi.md#getshareableusers) | **Get** /api/custom-dashboard/shareable-users | Get all users (without invitations).
*CustomDashboardsApi* | [**UpdateCustomDashboard**](docs/CustomDashboardsApi.md#updatecustomdashboard) | **Put** /api/custom-dashboard/{customDashboardId} | Update custom dashboard
*EventSettingsApi* | [**CreateWebsiteAlertConfig**](docs/EventSettingsApi.md#createwebsitealertconfig) | **Post** /api/events/settings/website-alert-configs | Create Website Alert Config
*EventSettingsApi* | [**DeleteAlert**](docs/EventSettingsApi.md#deletealert) | **Delete** /api/events/settings/alerts/{id} | Delete alerting
*EventSettingsApi* | [**DeleteAlertingChannel**](docs/EventSettingsApi.md#deletealertingchannel) | **Delete** /api/events/settings/alertingChannels/{id} | Delete alerting channel
*EventSettingsApi* | [**DeleteBuiltInEventSpecification**](docs/EventSettingsApi.md#deletebuiltineventspecification) | **Delete** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Delete built-in event specification
*EventSettingsApi* | [**DeleteCustomEventSpecification**](docs/EventSettingsApi.md#deletecustomeventspecification) | **Delete** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Delete custom event specification
*EventSettingsApi* | [**DeleteCustomPayloadConfiguration**](docs/EventSettingsApi.md#deletecustompayloadconfiguration) | **Delete** /api/events/settings/custom-payload-configurations | Delete Custom Payload Configuration
*EventSettingsApi* | [**DeleteWebsiteAlertConfig**](docs/EventSettingsApi.md#deletewebsitealertconfig) | **Delete** /api/events/settings/website-alert-configs/{id} | Delete Website Alert Config
*EventSettingsApi* | [**DisableBuiltInEventSpecification**](docs/EventSettingsApi.md#disablebuiltineventspecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/disable | Disable built-in event specification
*EventSettingsApi* | [**DisableCustomEventSpecification**](docs/EventSettingsApi.md#disablecustomeventspecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/disable | Disable custom event specification
*EventSettingsApi* | [**DisableWebsiteAlertConfig**](docs/EventSettingsApi.md#disablewebsitealertconfig) | **Put** /api/events/settings/website-alert-configs/{id}/disable | Disable Website Alert Config
*EventSettingsApi* | [**EnableBuiltInEventSpecification**](docs/EventSettingsApi.md#enablebuiltineventspecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/enable | Enable built-in event specification
*EventSettingsApi* | [**EnableCustomEventSpecification**](docs/EventSettingsApi.md#enablecustomeventspecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/enable | Enable custom event specification
*EventSettingsApi* | [**EnableWebsiteAlertConfig**](docs/EventSettingsApi.md#enablewebsitealertconfig) | **Put** /api/events/settings/website-alert-configs/{id}/enable | Enable Website Alert Config
*EventSettingsApi* | [**FindActiveWebsiteAlertConfigs**](docs/EventSettingsApi.md#findactivewebsitealertconfigs) | **Get** /api/events/settings/website-alert-configs | All Website Alert Configs
*EventSettingsApi* | [**FindWebsiteAlertConfig**](docs/EventSettingsApi.md#findwebsitealertconfig) | **Get** /api/events/settings/website-alert-configs/{id} | Get Website Alert Config
*EventSettingsApi* | [**FindWebsiteAlertConfigVersions**](docs/EventSettingsApi.md#findwebsitealertconfigversions) | **Get** /api/events/settings/website-alert-configs/{id}/versions | Get versions of Website Alert Config
*EventSettingsApi* | [**GetAlert**](docs/EventSettingsApi.md#getalert) | **Get** /api/events/settings/alerts/{id} | Alerting
*EventSettingsApi* | [**GetAlertingChannel**](docs/EventSettingsApi.md#getalertingchannel) | **Get** /api/events/settings/alertingChannels/{id} | Alerting channel
*EventSettingsApi* | [**GetAlertingChannels**](docs/EventSettingsApi.md#getalertingchannels) | **Get** /api/events/settings/alertingChannels | All alerting channels
*EventSettingsApi* | [**GetAlertingChannelsOverview**](docs/EventSettingsApi.md#getalertingchannelsoverview) | **Get** /api/events/settings/alertingChannels/infos | Overview over all alerting channels
*EventSettingsApi* | [**GetAlertingConfigurationInfos**](docs/EventSettingsApi.md#getalertingconfigurationinfos) | **Get** /api/events/settings/alerts/infos | All alerting configuration info
*EventSettingsApi* | [**GetAlerts**](docs/EventSettingsApi.md#getalerts) | **Get** /api/events/settings/alerts | All Alerting
*EventSettingsApi* | [**GetBuiltInEventSpecification**](docs/EventSettingsApi.md#getbuiltineventspecification) | **Get** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Built-in event specifications
*EventSettingsApi* | [**GetBuiltInEventSpecifications**](docs/EventSettingsApi.md#getbuiltineventspecifications) | **Get** /api/events/settings/event-specifications/built-in | All built-in event specification
*EventSettingsApi* | [**GetCustomEventSpecification**](docs/EventSettingsApi.md#getcustomeventspecification) | **Get** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Custom event specification
*EventSettingsApi* | [**GetCustomEventSpecifications**](docs/EventSettingsApi.md#getcustomeventspecifications) | **Get** /api/events/settings/event-specifications/custom | All custom event specifications
*EventSettingsApi* | [**GetCustomPayloadConfigurations**](docs/EventSettingsApi.md#getcustompayloadconfigurations) | **Get** /api/events/settings/custom-payload-configurations | Get Custom Payload Configuration
*EventSettingsApi* | [**GetCustomPayloadTagCatalog**](docs/EventSettingsApi.md#getcustompayloadtagcatalog) | **Get** /api/events/settings/custom-payload-configurations/catalog | Get tag catalog for custom payload in alerting
*EventSettingsApi* | [**GetEventSpecificationInfos**](docs/EventSettingsApi.md#geteventspecificationinfos) | **Get** /api/events/settings/event-specifications/infos | Summary of all built-in and custom event specifications
*EventSettingsApi* | [**GetEventSpecificationInfosByIds**](docs/EventSettingsApi.md#geteventspecificationinfosbyids) | **Post** /api/events/settings/event-specifications/infos | All built-in and custom event specifications
*EventSettingsApi* | [**GetSystemRules**](docs/EventSettingsApi.md#getsystemrules) | **Get** /api/events/settings/event-specifications/custom/systemRules | All system rules for custom event specifications
*EventSettingsApi* | [**PutAlert**](docs/EventSettingsApi.md#putalert) | **Put** /api/events/settings/alerts/{id} | Update alerting
*EventSettingsApi* | [**PutAlertingChannel**](docs/EventSettingsApi.md#putalertingchannel) | **Put** /api/events/settings/alertingChannels/{id} | Update alerting channel
*EventSettingsApi* | [**PutCustomEventSpecification**](docs/EventSettingsApi.md#putcustomeventspecification) | **Put** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Create or Update custom event specification
*EventSettingsApi* | [**SendTestAlerting**](docs/EventSettingsApi.md#sendtestalerting) | **Put** /api/events/settings/alertingChannels/test | Test alerting channel
*EventSettingsApi* | [**UpdateWebsiteAlertConfig**](docs/EventSettingsApi.md#updatewebsitealertconfig) | **Post** /api/events/settings/website-alert-configs/{id} | Update Website Alert Config
*EventSettingsApi* | [**UpsertCustomPayloadConfiguration**](docs/EventSettingsApi.md#upsertcustompayloadconfiguration) | **Put** /api/events/settings/custom-payload-configurations | Create / Update Custom Payload Configuration
*EventsApi* | [**GetEvent**](docs/EventsApi.md#getevent) | **Get** /api/events/{eventId} | Get Event
*EventsApi* | [**GetEvents**](docs/EventsApi.md#getevents) | **Get** /api/events | Get alerts
*GroupsApi* | [**CreateGroup**](docs/GroupsApi.md#creategroup) | **Post** /api/settings/group | Create group
*GroupsApi* | [**DeleteGroup**](docs/GroupsApi.md#deletegroup) | **Delete** /api/settings/group/{id} | Delete groups
*GroupsApi* | [**GetGroup**](docs/GroupsApi.md#getgroup) | **Get** /api/settings/group/{id} | Get group
*GroupsApi* | [**GetGroups**](docs/GroupsApi.md#getgroups) | **Get** /api/settings/groups | All groups
*GroupsApi* | [**GetGroupsByUser**](docs/GroupsApi.md#getgroupsbyuser) | **Get** /api/settings/groups/user/{email} | Get groups
*GroupsApi* | [**UpdateGroup**](docs/GroupsApi.md#updategroup) | **Put** /api/settings/group/{id} | Update group
*GroupsApi* | [**UpdateGroups**](docs/GroupsApi.md#updategroups) | **Put** /api/settings/groups | Create groups
*HealthApi* | [**GetHealthState**](docs/HealthApi.md#gethealthstate) | **Get** /api/instana/health | Basic health traffic light
*HealthApi* | [**GetVersion**](docs/HealthApi.md#getversion) | **Get** /api/instana/version | API version information
*HostAgentApi* | [**GetAgentSnapshot**](docs/HostAgentApi.md#getagentsnapshot) | **Get** /api/host-agent/{id} | Get host agent snapshot details
*HostAgentApi* | [**GetLogs**](docs/HostAgentApi.md#getlogs) | **Get** /api/host-agent/{hostId}/logs | Agent download logs
*HostAgentApi* | [**Search**](docs/HostAgentApi.md#search) | **Get** /api/host-agent | Query host agent snapshots
*HostAgentApi* | [**Update**](docs/HostAgentApi.md#update) | **Post** /api/host-agent/{hostId}/update | Agent update
*HostAgentApi* | [**UpdateConfigurationByHost**](docs/HostAgentApi.md#updateconfigurationbyhost) | **Post** /api/host-agent/{hostId}/configuration | Update agent configuration by host
*HostAgentApi* | [**UpdateConfigurationByQuery**](docs/HostAgentApi.md#updateconfigurationbyquery) | **Post** /api/host-agent/configuration | Update agent configuration by query
*InfrastructureCatalogApi* | [**GetInfrastructureCatalogMetrics**](docs/InfrastructureCatalogApi.md#getinfrastructurecatalogmetrics) | **Get** /api/infrastructure-monitoring/catalog/metrics/{plugin} | Get metric catalog
*InfrastructureCatalogApi* | [**GetInfrastructureCatalogPlugins**](docs/InfrastructureCatalogApi.md#getinfrastructurecatalogplugins) | **Get** /api/infrastructure-monitoring/catalog/plugins | Get plugin catalog
*InfrastructureCatalogApi* | [**GetInfrastructureCatalogPluginsWithCustomMetrics**](docs/InfrastructureCatalogApi.md#getinfrastructurecatalogpluginswithcustommetrics) | **Get** /api/infrastructure-monitoring/catalog/plugins-with-custom-metrics | Get all plugins with custom metrics catalog
*InfrastructureCatalogApi* | [**GetInfrastructureCatalogSearchFields**](docs/InfrastructureCatalogApi.md#getinfrastructurecatalogsearchfields) | **Get** /api/infrastructure-monitoring/catalog/search | get search field catalog
*InfrastructureMetricsApi* | [**GetInfrastructureMetrics**](docs/InfrastructureMetricsApi.md#getinfrastructuremetrics) | **Post** /api/infrastructure-monitoring/metrics | Get infrastructure metrics
*InfrastructureMetricsApi* | [**GetSnapshot**](docs/InfrastructureMetricsApi.md#getsnapshot) | **Get** /api/infrastructure-monitoring/snapshots/{id} | Get snapshot details
*InfrastructureMetricsApi* | [**GetSnapshots**](docs/InfrastructureMetricsApi.md#getsnapshots) | **Get** /api/infrastructure-monitoring/snapshots | Search snapshots
*InfrastructureResourcesApi* | [**GetInfrastructureViewTree**](docs/InfrastructureResourcesApi.md#getinfrastructureviewtree) | **Get** /api/infrastructure-monitoring/graph/views | Get view tree
*InfrastructureResourcesApi* | [**GetMonitoringState**](docs/InfrastructureResourcesApi.md#getmonitoringstate) | **Get** /api/infrastructure-monitoring/monitoring-state | Monitored host count
*InfrastructureResourcesApi* | [**GetRelatedHosts**](docs/InfrastructureResourcesApi.md#getrelatedhosts) | **Get** /api/infrastructure-monitoring/graph/related-hosts/{snapshotId} | Related hosts
*InfrastructureResourcesApi* | [**SoftwareVersions**](docs/InfrastructureResourcesApi.md#softwareversions) | **Get** /api/infrastructure-monitoring/software/versions | Get installed software
*MaintenanceConfigurationApi* | [**DeleteMaintenanceConfig**](docs/MaintenanceConfigurationApi.md#deletemaintenanceconfig) | **Delete** /api/settings/maintenance/{id} | Delete maintenance configuration
*MaintenanceConfigurationApi* | [**GetMaintenanceConfig**](docs/MaintenanceConfigurationApi.md#getmaintenanceconfig) | **Get** /api/settings/maintenance/{id} | Maintenance configuration
*MaintenanceConfigurationApi* | [**GetMaintenanceConfigs**](docs/MaintenanceConfigurationApi.md#getmaintenanceconfigs) | **Get** /api/settings/maintenance | All maintenance configurations
*MaintenanceConfigurationApi* | [**PutMaintenanceConfig**](docs/MaintenanceConfigurationApi.md#putmaintenanceconfig) | **Put** /api/settings/maintenance/{id} | Create or update maintenance configuration
*ReleasesApi* | [**DeleteRelease**](docs/ReleasesApi.md#deleterelease) | **Delete** /api/releases/{releaseId} | Delete release
*ReleasesApi* | [**GetAllReleases**](docs/ReleasesApi.md#getallreleases) | **Get** /api/releases | Get all releases
*ReleasesApi* | [**GetRelease**](docs/ReleasesApi.md#getrelease) | **Get** /api/releases/{releaseId} | Get release
*ReleasesApi* | [**PostRelease**](docs/ReleasesApi.md#postrelease) | **Post** /api/releases | Create release
*ReleasesApi* | [**PutRelease**](docs/ReleasesApi.md#putrelease) | **Put** /api/releases/{releaseId} | Update release
*SLIReportApi* | [**GetSli**](docs/SLIReportApi.md#getsli) | **Get** /api/sli/report/{sliId} | Generate SLI report
*SLISettingsApi* | [**CreateSli**](docs/SLISettingsApi.md#createsli) | **Post** /api/settings/sli | Create SLI Config
*SLISettingsApi* | [**Delete**](docs/SLISettingsApi.md#delete) | **Delete** /api/settings/sli/{id} | Delete SLI Config
*SLISettingsApi* | [**GetSli1**](docs/SLISettingsApi.md#getsli1) | **Get** /api/settings/sli/{id} | Get SLI Config
*SLISettingsApi* | [**GetSli2**](docs/SLISettingsApi.md#getsli2) | **Get** /api/settings/sli | Get All SLI Configs
*SLISettingsApi* | [**PutSli**](docs/SLISettingsApi.md#putsli) | **Put** /api/settings/sli/{id} | Update SLI Config
*SessionSettingsApi* | [**DeleteSessionSettings**](docs/SessionSettingsApi.md#deletesessionsettings) | **Delete** /api/settings/session | Delete session settings
*SessionSettingsApi* | [**GetSessionSettings**](docs/SessionSettingsApi.md#getsessionsettings) | **Get** /api/settings/session | Session settings
*SessionSettingsApi* | [**SetSessionSettings**](docs/SessionSettingsApi.md#setsessionsettings) | **Put** /api/settings/session | Configure session settings
*SyntheticCallsApi* | [**DeleteSyntheticCall**](docs/SyntheticCallsApi.md#deletesyntheticcall) | **Delete** /api/settings/synthetic-calls | Delete synthetic call configurations
*SyntheticCallsApi* | [**GetSyntheticCalls**](docs/SyntheticCallsApi.md#getsyntheticcalls) | **Get** /api/settings/synthetic-calls | Synthetic call configurations
*SyntheticCallsApi* | [**UpdateSyntheticCall**](docs/SyntheticCallsApi.md#updatesyntheticcall) | **Put** /api/settings/synthetic-calls | Update synthetic call configurations
*UsageApi* | [**GetAllUsage**](docs/UsageApi.md#getallusage) | **Get** /api/instana/usage/api | API usage by customer
*UsageApi* | [**GetHostsPerDay**](docs/UsageApi.md#gethostsperday) | **Get** /api/instana/usage/hosts/{day}/{month}/{year} | Host count day / month / year
*UsageApi* | [**GetHostsPerMonth**](docs/UsageApi.md#gethostspermonth) | **Get** /api/instana/usage/hosts/{month}/{year} | Host count month / year
*UsageApi* | [**GetUsagePerDay**](docs/UsageApi.md#getusageperday) | **Get** /api/instana/usage/api/{day}/{month}/{year} | API usage day / month / year
*UsageApi* | [**GetUsagePerMonth**](docs/UsageApi.md#getusagepermonth) | **Get** /api/instana/usage/api/{month}/{year} | API usage month / year
*UserApi* | [**DeleteRole**](docs/UserApi.md#deleterole) | **Delete** /api/settings/roles/{roleId} | Delete role
*UserApi* | [**GetInvitations**](docs/UserApi.md#getinvitations) | **Get** /api/settings/users/invitations | All pending invitations
*UserApi* | [**GetRole**](docs/UserApi.md#getrole) | **Get** /api/settings/roles/{roleId} | Role
*UserApi* | [**GetRoles**](docs/UserApi.md#getroles) | **Get** /api/settings/roles | All roles
*UserApi* | [**GetUsers**](docs/UserApi.md#getusers) | **Get** /api/settings/users | All users (without invitations)
*UserApi* | [**GetUsersIncludingInvitations**](docs/UserApi.md#getusersincludinginvitations) | **Get** /api/settings/users/overview | All users (incl. invitations)
*UserApi* | [**PutRole**](docs/UserApi.md#putrole) | **Put** /api/settings/roles/{roleId} | Create or update role
*UserApi* | [**RemoveUserFromTenant**](docs/UserApi.md#removeuserfromtenant) | **Delete** /api/settings/users/{userId} | Remove user from tenant
*UserApi* | [**RevokePendingInvitations**](docs/UserApi.md#revokependinginvitations) | **Delete** /api/settings/users/invitations | Revoke pending invitation
*UserApi* | [**SendInvitation**](docs/UserApi.md#sendinvitation) | **Post** /api/settings/users/invitations | Send user invitation
*UserApi* | [**SetRole**](docs/UserApi.md#setrole) | **Put** /api/settings/users/{userId}/role | Add user to role
*WebsiteAnalyzeApi* | [**GetBeaconGroups**](docs/WebsiteAnalyzeApi.md#getbeacongroups) | **Post** /api/website-monitoring/analyze/beacon-groups | Get grouped beacon metrics
*WebsiteAnalyzeApi* | [**GetBeacons**](docs/WebsiteAnalyzeApi.md#getbeacons) | **Post** /api/website-monitoring/analyze/beacons | Get all beacons
*WebsiteCatalogApi* | [**GetWebsiteCatalogMetrics**](docs/WebsiteCatalogApi.md#getwebsitecatalogmetrics) | **Get** /api/website-monitoring/catalog/metrics | Metric catalog
*WebsiteCatalogApi* | [**GetWebsiteCatalogTags**](docs/WebsiteCatalogApi.md#getwebsitecatalogtags) | **Get** /api/website-monitoring/catalog/tags | Get all existing website tags
*WebsiteCatalogApi* | [**GetWebsiteTagCatalog**](docs/WebsiteCatalogApi.md#getwebsitetagcatalog) | **Get** /api/website-monitoring/catalog | Get website tag catalog
*WebsiteConfigurationApi* | [**Delete1**](docs/WebsiteConfigurationApi.md#delete1) | **Delete** /api/website-monitoring/config/{websiteId} | Remove website
*WebsiteConfigurationApi* | [**Get**](docs/WebsiteConfigurationApi.md#get) | **Get** /api/website-monitoring/config/{websiteId} | Get configured website
*WebsiteConfigurationApi* | [**GetWebsites**](docs/WebsiteConfigurationApi.md#getwebsites) | **Get** /api/website-monitoring/config | Get configured websites
*WebsiteConfigurationApi* | [**Post**](docs/WebsiteConfigurationApi.md#post) | **Post** /api/website-monitoring/config | Configure new website
*WebsiteConfigurationApi* | [**Rename**](docs/WebsiteConfigurationApi.md#rename) | **Put** /api/website-monitoring/config/{websiteId} | Rename website
*WebsiteMetricsApi* | [**GetBeaconMetrics**](docs/WebsiteMetricsApi.md#getbeaconmetrics) | **Post** /api/website-monitoring/metrics | Get beacon metrics
*WebsiteMetricsApi* | [**GetPageLoad**](docs/WebsiteMetricsApi.md#getpageload) | **Get** /api/website-monitoring/page-load | Get page load


## Documentation For Models

 - [AbstractIntegration](docs/AbstractIntegration.md)
 - [AbstractRule](docs/AbstractRule.md)
 - [AccessRule](docs/AccessRule.md)
 - [AgentConfigurationUpdate](docs/AgentConfigurationUpdate.md)
 - [AlertingConfiguration](docs/AlertingConfiguration.md)
 - [AlertingConfigurationWithLastUpdated](docs/AlertingConfigurationWithLastUpdated.md)
 - [ApiToken](docs/ApiToken.md)
 - [AppDataMetricConfiguration](docs/AppDataMetricConfiguration.md)
 - [Application](docs/Application.md)
 - [ApplicationConfig](docs/ApplicationConfig.md)
 - [ApplicationItem](docs/ApplicationItem.md)
 - [ApplicationMetricResult](docs/ApplicationMetricResult.md)
 - [ApplicationResult](docs/ApplicationResult.md)
 - [ApplicationScope](docs/ApplicationScope.md)
 - [ApplicationScopeWithMetadata](docs/ApplicationScopeWithMetadata.md)
 - [ApplicationSliEntity](docs/ApplicationSliEntity.md)
 - [ApplicationSliEntityAllOf](docs/ApplicationSliEntityAllOf.md)
 - [AuditLogEntry](docs/AuditLogEntry.md)
 - [AuditLogResponse](docs/AuditLogResponse.md)
 - [AvailabilitySliEntity](docs/AvailabilitySliEntity.md)
 - [AvailabilitySliEntityAllOf](docs/AvailabilitySliEntityAllOf.md)
 - [BeaconGroupsResult](docs/BeaconGroupsResult.md)
 - [BeaconResult](docs/BeaconResult.md)
 - [BinaryOperatorDto](docs/BinaryOperatorDto.md)
 - [BinaryOperatorDtoAllOf](docs/BinaryOperatorDtoAllOf.md)
 - [BuiltInEventSpecification](docs/BuiltInEventSpecification.md)
 - [BuiltInEventSpecificationWithLastUpdated](docs/BuiltInEventSpecificationWithLastUpdated.md)
 - [CallGroupsItem](docs/CallGroupsItem.md)
 - [CallGroupsResult](docs/CallGroupsResult.md)
 - [CloudfoundryPhysicalContext](docs/CloudfoundryPhysicalContext.md)
 - [ConfigVersion](docs/ConfigVersion.md)
 - [CursorPagination](docs/CursorPagination.md)
 - [CustomDashboard](docs/CustomDashboard.md)
 - [CustomEventSpecification](docs/CustomEventSpecification.md)
 - [CustomEventSpecificationWithLastUpdated](docs/CustomEventSpecificationWithLastUpdated.md)
 - [CustomPayloadConfiguration](docs/CustomPayloadConfiguration.md)
 - [CustomPayloadField](docs/CustomPayloadField.md)
 - [CustomPayloadWithLastUpdated](docs/CustomPayloadWithLastUpdated.md)
 - [DynamicField](docs/DynamicField.md)
 - [DynamicFieldAllOf](docs/DynamicFieldAllOf.md)
 - [DynamicFieldValue](docs/DynamicFieldValue.md)
 - [EmailIntegration](docs/EmailIntegration.md)
 - [EmailIntegrationAllOf](docs/EmailIntegrationAllOf.md)
 - [Endpoint](docs/Endpoint.md)
 - [EndpointItem](docs/EndpointItem.md)
 - [EndpointMetricResult](docs/EndpointMetricResult.md)
 - [EndpointResult](docs/EndpointResult.md)
 - [EntityVerificationRule](docs/EntityVerificationRule.md)
 - [EntityVerificationRuleAllOf](docs/EntityVerificationRuleAllOf.md)
 - [EventFilteringConfiguration](docs/EventFilteringConfiguration.md)
 - [EventResult](docs/EventResult.md)
 - [EventSpecificationInfo](docs/EventSpecificationInfo.md)
 - [FixedHttpPathSegmentMatchingRule](docs/FixedHttpPathSegmentMatchingRule.md)
 - [FixedHttpPathSegmentMatchingRuleAllOf](docs/FixedHttpPathSegmentMatchingRuleAllOf.md)
 - [FullTrace](docs/FullTrace.md)
 - [GetApplications](docs/GetApplications.md)
 - [GetCallGroups](docs/GetCallGroups.md)
 - [GetCombinedMetrics](docs/GetCombinedMetrics.md)
 - [GetEndpoints](docs/GetEndpoints.md)
 - [GetServices](docs/GetServices.md)
 - [GetTraceGroups](docs/GetTraceGroups.md)
 - [GetTraces](docs/GetTraces.md)
 - [GetWebsiteBeaconGroups](docs/GetWebsiteBeaconGroups.md)
 - [GetWebsiteBeacons](docs/GetWebsiteBeacons.md)
 - [GetWebsiteMetrics](docs/GetWebsiteMetrics.md)
 - [GoogleChatIntegration](docs/GoogleChatIntegration.md)
 - [GoogleChatIntegrationAllOf](docs/GoogleChatIntegrationAllOf.md)
 - [Group](docs/Group.md)
 - [GroupMember](docs/GroupMember.md)
 - [GroupPermission](docs/GroupPermission.md)
 - [HealthState](docs/HealthState.md)
 - [HistoricBaseline](docs/HistoricBaseline.md)
 - [HistoricBaselineAllOf](docs/HistoricBaselineAllOf.md)
 - [HttpEndpointConfig](docs/HttpEndpointConfig.md)
 - [HttpEndpointRule](docs/HttpEndpointRule.md)
 - [HttpPathSegmentMatchingRule](docs/HttpPathSegmentMatchingRule.md)
 - [HyperParam](docs/HyperParam.md)
 - [InfrastructureMetricResult](docs/InfrastructureMetricResult.md)
 - [InfrastructureSliEntity](docs/InfrastructureSliEntity.md)
 - [InfrastructureSliEntityAllOf](docs/InfrastructureSliEntityAllOf.md)
 - [IngestionOffsetCursor](docs/IngestionOffsetCursor.md)
 - [InstanaVersionInfo](docs/InstanaVersionInfo.md)
 - [IntegrationOverview](docs/IntegrationOverview.md)
 - [KubernetesPhysicalContext](docs/KubernetesPhysicalContext.md)
 - [LogEntryActor](docs/LogEntryActor.md)
 - [MaintenanceConfig](docs/MaintenanceConfig.md)
 - [MaintenanceConfigWithLastUpdated](docs/MaintenanceConfigWithLastUpdated.md)
 - [MaintenanceWindow](docs/MaintenanceWindow.md)
 - [MatchAllHttpPathSegmentMatchingRule](docs/MatchAllHttpPathSegmentMatchingRule.md)
 - [MatchExpressionDto](docs/MatchExpressionDto.md)
 - [Member](docs/Member.md)
 - [MetricConfig](docs/MetricConfig.md)
 - [MetricConfiguration](docs/MetricConfiguration.md)
 - [MetricDescription](docs/MetricDescription.md)
 - [MetricInstance](docs/MetricInstance.md)
 - [MetricItem](docs/MetricItem.md)
 - [MetricPattern](docs/MetricPattern.md)
 - [MonitoringState](docs/MonitoringState.md)
 - [NewApplicationConfig](docs/NewApplicationConfig.md)
 - [Office365Integration](docs/Office365Integration.md)
 - [OpsgenieIntegration](docs/OpsgenieIntegration.md)
 - [OpsgenieIntegrationAllOf](docs/OpsgenieIntegrationAllOf.md)
 - [Order](docs/Order.md)
 - [PagerdutyIntegration](docs/PagerdutyIntegration.md)
 - [PagerdutyIntegrationAllOf](docs/PagerdutyIntegrationAllOf.md)
 - [Pagination](docs/Pagination.md)
 - [PathParameterHttpPathSegmentMatchingRule](docs/PathParameterHttpPathSegmentMatchingRule.md)
 - [PermissionSet](docs/PermissionSet.md)
 - [PhysicalContext](docs/PhysicalContext.md)
 - [PluginResult](docs/PluginResult.md)
 - [PrometheusWebhookIntegration](docs/PrometheusWebhookIntegration.md)
 - [PrometheusWebhookIntegrationAllOf](docs/PrometheusWebhookIntegrationAllOf.md)
 - [RbacGroup](docs/RbacGroup.md)
 - [Release](docs/Release.md)
 - [ReleaseScope](docs/ReleaseScope.md)
 - [ReleaseWithMetadata](docs/ReleaseWithMetadata.md)
 - [Role](docs/Role.md)
 - [RuleInput](docs/RuleInput.md)
 - [ScopedGroup](docs/ScopedGroup.md)
 - [SearchFieldResult](docs/SearchFieldResult.md)
 - [Service](docs/Service.md)
 - [ServiceConfig](docs/ServiceConfig.md)
 - [ServiceItem](docs/ServiceItem.md)
 - [ServiceMatchingRule](docs/ServiceMatchingRule.md)
 - [ServiceMetricResult](docs/ServiceMetricResult.md)
 - [ServiceResult](docs/ServiceResult.md)
 - [ServiceScope](docs/ServiceScope.md)
 - [ServiceScopeWithMetadata](docs/ServiceScopeWithMetadata.md)
 - [ServiceScopedTo](docs/ServiceScopedTo.md)
 - [ServiceScopedToWithMetadata](docs/ServiceScopedToWithMetadata.md)
 - [SessionSettings](docs/SessionSettings.md)
 - [SlackIntegration](docs/SlackIntegration.md)
 - [SlackIntegrationAllOf](docs/SlackIntegrationAllOf.md)
 - [SliConfiguration](docs/SliConfiguration.md)
 - [SliConfigurationWithLastUpdated](docs/SliConfigurationWithLastUpdated.md)
 - [SliEntity](docs/SliEntity.md)
 - [SliReport](docs/SliReport.md)
 - [SlownessWebsiteAlertRule](docs/SlownessWebsiteAlertRule.md)
 - [SlownessWebsiteAlertRuleAllOf](docs/SlownessWebsiteAlertRuleAllOf.md)
 - [SnapshotItem](docs/SnapshotItem.md)
 - [SnapshotPreview](docs/SnapshotPreview.md)
 - [SnapshotResult](docs/SnapshotResult.md)
 - [SoftwareUser](docs/SoftwareUser.md)
 - [SoftwareVersion](docs/SoftwareVersion.md)
 - [Span](docs/Span.md)
 - [SpanRelation](docs/SpanRelation.md)
 - [SpecificJsErrorsWebsiteAlertRule](docs/SpecificJsErrorsWebsiteAlertRule.md)
 - [SpecificJsErrorsWebsiteAlertRuleAllOf](docs/SpecificJsErrorsWebsiteAlertRuleAllOf.md)
 - [SplunkIntegration](docs/SplunkIntegration.md)
 - [SplunkIntegrationAllOf](docs/SplunkIntegrationAllOf.md)
 - [StackTraceItem](docs/StackTraceItem.md)
 - [StackTraceLine](docs/StackTraceLine.md)
 - [StaticBooleanField](docs/StaticBooleanField.md)
 - [StaticBooleanFieldAllOf](docs/StaticBooleanFieldAllOf.md)
 - [StaticNumberField](docs/StaticNumberField.md)
 - [StaticNumberFieldAllOf](docs/StaticNumberFieldAllOf.md)
 - [StaticStringField](docs/StaticStringField.md)
 - [StaticStringFieldAllOf](docs/StaticStringFieldAllOf.md)
 - [StaticThreshold](docs/StaticThreshold.md)
 - [StaticThresholdAllOf](docs/StaticThresholdAllOf.md)
 - [StatusCodeWebsiteAlertRule](docs/StatusCodeWebsiteAlertRule.md)
 - [SyntheticCallConfig](docs/SyntheticCallConfig.md)
 - [SyntheticCallRule](docs/SyntheticCallRule.md)
 - [SyntheticCallWithDefaultsConfig](docs/SyntheticCallWithDefaultsConfig.md)
 - [SystemRule](docs/SystemRule.md)
 - [SystemRuleAllOf](docs/SystemRuleAllOf.md)
 - [SystemRuleLabel](docs/SystemRuleLabel.md)
 - [Tag](docs/Tag.md)
 - [TagCatalog](docs/TagCatalog.md)
 - [TagFilter](docs/TagFilter.md)
 - [TagMatcherDto](docs/TagMatcherDto.md)
 - [TagMatcherDtoAllOf](docs/TagMatcherDtoAllOf.md)
 - [TagTreeLevel](docs/TagTreeLevel.md)
 - [TagTreeNode](docs/TagTreeNode.md)
 - [TagTreeTag](docs/TagTreeTag.md)
 - [TagTreeTagAllOf](docs/TagTreeTagAllOf.md)
 - [Threshold](docs/Threshold.md)
 - [ThresholdRule](docs/ThresholdRule.md)
 - [ThresholdRuleAllOf](docs/ThresholdRuleAllOf.md)
 - [ThroughputWebsiteAlertRule](docs/ThroughputWebsiteAlertRule.md)
 - [TimeFrame](docs/TimeFrame.md)
 - [Trace](docs/Trace.md)
 - [TraceGroupsItem](docs/TraceGroupsItem.md)
 - [TraceGroupsResult](docs/TraceGroupsResult.md)
 - [TraceItem](docs/TraceItem.md)
 - [TraceResult](docs/TraceResult.md)
 - [TreeNode](docs/TreeNode.md)
 - [TreeNodeResult](docs/TreeNodeResult.md)
 - [UnsupportedHttpPathSegmentMatchingRule](docs/UnsupportedHttpPathSegmentMatchingRule.md)
 - [UnsupportedHttpPathSegmentMatchingRuleAllOf](docs/UnsupportedHttpPathSegmentMatchingRuleAllOf.md)
 - [UsageResult](docs/UsageResult.md)
 - [UsageResultItems](docs/UsageResultItems.md)
 - [UserImpactWebsiteTimeThreshold](docs/UserImpactWebsiteTimeThreshold.md)
 - [UserImpactWebsiteTimeThresholdAllOf](docs/UserImpactWebsiteTimeThresholdAllOf.md)
 - [UserResult](docs/UserResult.md)
 - [UsersResult](docs/UsersResult.md)
 - [ValidatedAlertingChannelInputInfo](docs/ValidatedAlertingChannelInputInfo.md)
 - [ValidatedAlertingConfiguration](docs/ValidatedAlertingConfiguration.md)
 - [ValidatedMaintenanceConfigWithStatus](docs/ValidatedMaintenanceConfigWithStatus.md)
 - [VictorOpsIntegration](docs/VictorOpsIntegration.md)
 - [VictorOpsIntegrationAllOf](docs/VictorOpsIntegrationAllOf.md)
 - [ViolationsInPeriodWebsiteTimeThreshold](docs/ViolationsInPeriodWebsiteTimeThreshold.md)
 - [ViolationsInPeriodWebsiteTimeThresholdAllOf](docs/ViolationsInPeriodWebsiteTimeThresholdAllOf.md)
 - [ViolationsInSequenceWebsiteTimeThreshold](docs/ViolationsInSequenceWebsiteTimeThreshold.md)
 - [WebexTeamsWebhookIntegration](docs/WebexTeamsWebhookIntegration.md)
 - [WebhookIntegration](docs/WebhookIntegration.md)
 - [WebhookIntegrationAllOf](docs/WebhookIntegrationAllOf.md)
 - [Website](docs/Website.md)
 - [WebsiteAlertConfig](docs/WebsiteAlertConfig.md)
 - [WebsiteAlertConfigWithMetadata](docs/WebsiteAlertConfigWithMetadata.md)
 - [WebsiteAlertRule](docs/WebsiteAlertRule.md)
 - [WebsiteBeaconGroupsItem](docs/WebsiteBeaconGroupsItem.md)
 - [WebsiteBeaconTagGroup](docs/WebsiteBeaconTagGroup.md)
 - [WebsiteBeaconsItem](docs/WebsiteBeaconsItem.md)
 - [WebsiteMetricResult](docs/WebsiteMetricResult.md)
 - [WebsiteMonitoringBeacon](docs/WebsiteMonitoringBeacon.md)
 - [WebsiteMonitoringMetricDescription](docs/WebsiteMonitoringMetricDescription.md)
 - [WebsiteMonitoringMetricsConfiguration](docs/WebsiteMonitoringMetricsConfiguration.md)
 - [WebsiteTimeThreshold](docs/WebsiteTimeThreshold.md)
 - [Widget](docs/Widget.md)


## Documentation For Authorization



## ApiKeyAuth

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```



## Author

support@instana.com

