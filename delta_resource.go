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

type DeltaResource struct {

	// The id of the category for question
	CategoryId string `json:"category_id,omitempty"`

	// The media type of the question
	MediaType string `json:"media_type,omitempty"`

	// The id of the question
	QuestionId string `json:"question_id,omitempty"`

	// Whether the question was updated or removed
	State string `json:"state,omitempty"`

	// The tags for the question
	Tags []string `json:"tags,omitempty"`

	// The date this question was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
