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

type NestedCategory struct {

	// Whether the category is active
	Active bool `json:"active,omitempty"`

	// The id of the category
	Id string `json:"id"`

	// The name of the category
	Name string `json:"name,omitempty"`
}