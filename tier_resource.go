/* 
 * Knetik Platform API Documentation Latest
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type TierResource struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The name of the tier
	Name string `json:"name,omitempty"`

	// The required progress for the tier
	RequiredProgress int32 `json:"required_progress,omitempty"`

	// The name of the triggered event
	TriggerEventName string `json:"trigger_event_name,omitempty"`
}
