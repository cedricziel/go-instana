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

// WebsiteAlertConfigWithMetadata struct for WebsiteAlertConfigWithMetadata
type WebsiteAlertConfigWithMetadata struct {
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	WebsiteId       string               `json:"websiteId"`
	Severity        int32                `json:"severity,omitempty"`
	Triggering      bool                 `json:"triggering,omitempty"`
	TagFilters      []TagFilter          `json:"tagFilters"`
	Rule            WebsiteAlertRule     `json:"rule"`
	Baseline        HistoricBaseline     `json:"baseline,omitempty"`
	AlertChannelIds []string             `json:"alertChannelIds"`
	Granularity     int32                `json:"granularity,omitempty"`
	TimeThreshold   WebsiteTimeThreshold `json:"timeThreshold"`
	Id              string               `json:"id"`
	Created         int64                `json:"created,omitempty"`
	ReadOnly        bool                 `json:"readOnly,omitempty"`
	Enabled         bool                 `json:"enabled,omitempty"`
	Threshold       Threshold            `json:"threshold"`
}
