/* 
 * Knetik Platform API Documentation Latest
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type PollResponseResource struct {

	// The answer to the poll
	Answer string `json:"answer,omitempty"`

	// The date the poll was answered, in seconds since unix epoc
	AnsweredDate int64 `json:"answered_date,omitempty"`

	// The id of the poll response
	Id string `json:"id,omitempty"`

	// The id of the poll
	PollId string `json:"poll_id,omitempty"`

	// The user
	User SimpleUserResource `json:"user,omitempty"`
}