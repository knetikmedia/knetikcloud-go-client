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

type CategoryResource struct {

	// Whether the category is currently active. If not, it and its questions will be filtered out.
	Active bool `json:"active,omitempty"`

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The unique ID for this category
	Id string `json:"id,omitempty"`

	// The name of this category. Cannot be blank
	Name string `json:"name"`

	// A category template this category is validated against (private). May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`
}
