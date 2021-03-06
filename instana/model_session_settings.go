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

// SessionSettings struct for SessionSettings
type SessionSettings struct {
	TokenLifeTimeInMillis int64 `json:"tokenLifeTimeInMillis,omitempty"`
	IdleTimeInMillis      int64 `json:"idleTimeInMillis,omitempty"`
}
