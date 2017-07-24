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

type BreTriggerParameterDefinition struct {

	// The name of the parameter. This is used as the unique identifier of this parameter
	Name string `json:"name"`

	// Whether this parameter can be left off when firing the event. Default false
	Optional bool `json:"optional,omitempty"`

	// The variable type of this parameter. See Bre Variables endpoint for list
	Type_ string `json:"type"`
}
