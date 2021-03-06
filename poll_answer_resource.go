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

type PollAnswerResource struct {

	// The number of users that selected this answer
	Count int32 `json:"count,omitempty"`

	// The key to the answer (for code reference)
	Key string `json:"key"`

	// The text of the answer (for user display)
	Text string `json:"text"`
}
