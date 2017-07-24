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

type BreGlobalScopeDefinition struct {

	// The name of the scoping parameter. This is used as the unique identifier of this scope
	Name string `json:"name"`

	// The variable type of this scoping parameter. See Bre Variables endpoint for list
	Type_ string `json:"type"`
}
