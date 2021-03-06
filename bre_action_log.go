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

type BreActionLog struct {

	// The name of the action
	Name string `json:"name,omitempty"`

	// The runtime of the action in milliseconds
	Runtime int64 `json:"runtime,omitempty"`

	// The status of the action (ran, failed)
	Status string `json:"status,omitempty"`
}
