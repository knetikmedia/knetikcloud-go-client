/* 
 * Knetik Platform API Documentation Latest
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Maintenance struct {

	// Whether access to the system has been locked
	AccessLocked bool `json:"access_locked,omitempty"`

	// A simple object of any schema for client side use and processing
	Details interface{} `json:"details,omitempty"`

	// User displayable message about the maintenance
	Message string `json:"message,omitempty"`
}
