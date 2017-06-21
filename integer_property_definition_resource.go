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

type IntegerPropertyDefinitionResource struct {

	// A list of the fields on both the property definition and property of this type
	FieldList PropertyFieldListResource `json:"field_list,omitempty"`

	// The name of the property
	Name string `json:"name"`

	// Whether the property is required
	Required bool `json:"required"`

	// The type of the property. Used for polymorphic type recognition and thus must match an expected type with additional properties.
	Type_ string `json:"type"`

	// If provided, the maximum value
	Max int32 `json:"max,omitempty"`

	// If provided, the minimum value
	Min int32 `json:"min,omitempty"`
}
