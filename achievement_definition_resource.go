/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://knetikcloud.com.
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type AchievementDefinitionResource struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this resource type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The description of the achievement. Must be at least 2 characters in length.
	Description string `json:"description,omitempty"`

	// Whether the achievement is hidden from the user
	Hidden bool `json:"hidden"`

	// The name of the achievement. Must be at least 6 characters in length. IMMUTABLE
	Name string `json:"name"`

	// The required progress for the achievement definition
	RequiredProgress int32 `json:"required_progress"`

	// The id of the rule generated for this achievement
	RuleId string `json:"rule_id,omitempty"`

	// The tags for the achievement definition
	Tags []string `json:"tags,omitempty"`

	// An achievement template this achievement is validated against (private). May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`

	// The name of the trigger event associated with this achievement
	TriggerEventName string `json:"trigger_event_name,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
