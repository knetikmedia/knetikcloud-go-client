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

type AnswerResource struct {

	// The answer to the question. Different 'type' values indicate different structures as the answer may be test, image, etc. See information on additional properties for the list and their structures
	Answer Property `json:"answer"`

	// Whether the answer is correct or not
	Correct bool `json:"correct"`

	// The unique ID for that resource
	Id string `json:"id,omitempty"`
}
