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

type BatchResult struct {

	// List of batch responses.  Returns in the order requested
	BatchReturn []BatchReturn `json:"batch_return,omitempty"`

	// The date the batch call started processing
	CreatedDate int64 `json:"created_date,omitempty"`

	// The token to use at the /batch/{token} endpoint if the request times out
	Id string `json:"id,omitempty"`

	// The date the batch call finished processing
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
