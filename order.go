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

type Order struct {

	Ascending bool `json:"ascending,omitempty"`

	Direction string `json:"direction,omitempty"`

	IgnoreCase bool `json:"ignore_case,omitempty"`

	NullHandling string `json:"null_handling,omitempty"`

	Property string `json:"property,omitempty"`
}
