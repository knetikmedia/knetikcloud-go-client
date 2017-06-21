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

type Consumable struct {

	Description string `json:"description,omitempty"`

	// Used for polymorphic type recognition and thus must match an expected type with additional properties
	TypeHint string `json:"type_hint,omitempty"`

	// The maximum number of times an item can be used
	MaxUse int32 `json:"max_use,omitempty"`
}
