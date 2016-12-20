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

type BareChallengeActivityResource struct {

	// The id of the activity
	ActivityId int64 `json:"activity_id,omitempty"`

	// The id of the challenge
	ChallengeId int64 `json:"challenge_id,omitempty"`

	// The unique ID for this resource
	Id int64 `json:"id,omitempty"`
}