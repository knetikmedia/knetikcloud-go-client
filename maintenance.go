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

type Maintenance struct {

	// Whether access to the system has been locked
	AccessLocked bool `json:"access_locked"`

	// A simple object of any schema for client side use and processing
	Details *interface{} `json:"details,omitempty"`

	// User displayable message about the maintenance
	Message string `json:"message"`
}
