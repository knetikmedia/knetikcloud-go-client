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

type FulfillmentType struct {

	// Whether the type is core and cannot be altered/deleted, read-only
	Core bool `json:"core,omitempty"`

	// A description of the type
	Description string `json:"description,omitempty"`

	// The unique id of the type, read-only
	Id int32 `json:"id,omitempty"`

	// The name of the type
	Name string `json:"name"`
}
