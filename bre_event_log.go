/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type BreEventLog struct {

	// The customer of the BRE event log
	Customer string `json:"customer,omitempty"`

	// The event id of the BRE event log
	EventId int64 `json:"event_id,omitempty"`

	// The event name of the BRE event log
	EventName string `json:"event_name,omitempty"`

	// The event start date of the BRE event log
	EventStartDate int64 `json:"event_start_date,omitempty"`

	// The id of the BRE event log
	Id string `json:"id,omitempty"`

	// The event paramters of the BRE event log
	Parameters interface{} `json:"parameters,omitempty"`

	// The rules of the BRE event log
	Rules []BreRuleLog `json:"rules,omitempty"`
}
