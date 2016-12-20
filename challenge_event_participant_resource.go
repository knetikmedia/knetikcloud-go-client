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

type ChallengeEventParticipantResource struct {

	Email string `json:"email,omitempty"`

	Fullname string `json:"fullname,omitempty"`

	Score int64 `json:"score,omitempty"`

	UserId int32 `json:"user_id,omitempty"`

	Username string `json:"username,omitempty"`
}
