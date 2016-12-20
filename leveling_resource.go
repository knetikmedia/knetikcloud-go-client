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

type LevelingResource struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The date the leveling schema was created
	CreatedDate int64 `json:"created_date,omitempty"`

	// The description of the leveling schema
	Description string `json:"description,omitempty"`

	// The name of the leveling schema.  IMMUTABLE
	Name string `json:"name,omitempty"`

	// A set of tiers that contain experience boundaries
	Tiers []TierResource `json:"tiers,omitempty"`

	// The date the leveling schema was updated
	UpdatedDate int64 `json:"updated_date,omitempty"`
}