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

type ActionResource struct {

	// The category the action is in. All customer specific actions are in the 'custom' category
	Category string `json:"category,omitempty"`

	// The description of the action
	Description string `json:"description,omitempty"`

	// The name of the action. Used as the unique id for reference
	Name string `json:"name,omitempty"`

	// The variables required for the action
	Variables []ActionVariableResource `json:"variables,omitempty"`
}