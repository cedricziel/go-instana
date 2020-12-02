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

// EntityVerificationRule struct for EntityVerificationRule
type EntityVerificationRule struct {
	// Values: `\"THRESHOLD\"`  `\"SYSTEM\"`  `\"ENTITY_VERIFICATION\"`
	RuleType            string `json:"ruleType"`
	Severity            int32  `json:"severity,omitempty"`
	MatchingEntityType  string `json:"matchingEntityType"`
	MatchingOperator    string `json:"matchingOperator"`
	MatchingEntityLabel string `json:"matchingEntityLabel"`
	OfflineDuration     int64  `json:"offlineDuration,omitempty"`
}
