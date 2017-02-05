/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ActionVariableResource struct {

	// The name of the variable
	Name string `json:"name"`

	// Whether this variable is optional and can be left out/null
	Optional bool `json:"optional"`

	// The type of the variable (see variable type endpoint for list)
	Type_ string `json:"type"`
}
