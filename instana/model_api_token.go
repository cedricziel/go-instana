/*
 * Introduction to Instana public APIs
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.192.86
 * Contact: support@instana.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package instana

// ApiToken struct for ApiToken
type ApiToken struct {
	Id                                  string `json:"id,omitempty"`
	AccessGrantingToken                 string `json:"accessGrantingToken"`
	InternalId                          string `json:"internalId"`
	Name                                string `json:"name"`
	CanConfigureServiceMapping          bool   `json:"canConfigureServiceMapping,omitempty"`
	CanConfigureEumApplications         bool   `json:"canConfigureEumApplications,omitempty"`
	CanConfigureMobileAppMonitoring     bool   `json:"canConfigureMobileAppMonitoring,omitempty"`
	CanConfigureUsers                   bool   `json:"canConfigureUsers,omitempty"`
	CanInstallNewAgents                 bool   `json:"canInstallNewAgents,omitempty"`
	CanSeeUsageInformation              bool   `json:"canSeeUsageInformation,omitempty"`
	CanConfigureIntegrations            bool   `json:"canConfigureIntegrations,omitempty"`
	CanSeeOnPremLicenseInformation      bool   `json:"canSeeOnPremLicenseInformation,omitempty"`
	CanConfigureRoles                   bool   `json:"canConfigureRoles,omitempty"`
	CanConfigureCustomAlerts            bool   `json:"canConfigureCustomAlerts,omitempty"`
	CanConfigureApiTokens               bool   `json:"canConfigureApiTokens,omitempty"`
	CanConfigureAgentRunMode            bool   `json:"canConfigureAgentRunMode,omitempty"`
	CanViewAuditLog                     bool   `json:"canViewAuditLog,omitempty"`
	CanConfigureObjectives              bool   `json:"canConfigureObjectives,omitempty"`
	CanConfigureAgents                  bool   `json:"canConfigureAgents,omitempty"`
	CanConfigureAuthenticationMethods   bool   `json:"canConfigureAuthenticationMethods,omitempty"`
	CanConfigureApplications            bool   `json:"canConfigureApplications,omitempty"`
	CanConfigureTeams                   bool   `json:"canConfigureTeams,omitempty"`
	CanConfigureReleases                bool   `json:"canConfigureReleases,omitempty"`
	CanConfigureLogManagement           bool   `json:"canConfigureLogManagement,omitempty"`
	CanCreatePublicCustomDashboards     bool   `json:"canCreatePublicCustomDashboards,omitempty"`
	CanViewLogs                         bool   `json:"canViewLogs,omitempty"`
	CanViewTraceDetails                 bool   `json:"canViewTraceDetails,omitempty"`
	CanConfigureSessionSettings         bool   `json:"canConfigureSessionSettings,omitempty"`
	CanConfigureServiceLevelIndicators  bool   `json:"canConfigureServiceLevelIndicators,omitempty"`
	CanConfigureGlobalAlertPayload      bool   `json:"canConfigureGlobalAlertPayload,omitempty"`
	CanViewAccountAndBillingInformation bool   `json:"canViewAccountAndBillingInformation,omitempty"`
}
