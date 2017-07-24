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

type PredicateResource struct {

	// The arguments the operator apply to. See notes for details.
	Args []ExpressionResource `json:"args"`

	// The operator to be used in this predicate. See notes for details.
	Op string `json:"op"`

	Type_ string `json:"type,omitempty"`
}
