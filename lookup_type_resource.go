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

type LookupTypeResource struct {

	// The description of the expression type
	Description string `json:"description"`

	// The variable type the key expression must be, or null if it's dependent (see description for explanation in this case)
	KeyType string `json:"key_type"`

	// The name of the expression type
	Name string `json:"name"`

	// The variable type of the value this expression returns, or null if it's dependent (see description for explanation in this case)
	ValueType string `json:"value_type"`
}
