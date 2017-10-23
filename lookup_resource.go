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

// Expressions are instructions for the rule engine to resolve certain values. For example instead of `user 1` it'd state `user provided by the event`. Full list and definitions available at GET /bre/expressions.
type LookupResource struct {

	Definition string `json:"definition,omitempty"`

	LookupKey *ExpressionResource `json:"lookup_key,omitempty"`

	RequiredKeyType string `json:"required_key_type,omitempty"`

	Type_ string `json:"type,omitempty"`

	ValueType string `json:"value_type,omitempty"`
}
