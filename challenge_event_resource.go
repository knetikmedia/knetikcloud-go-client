/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ChallengeEventResource struct {

	// The id of the challenge
	ChallengeId int64 `json:"challenge_id,omitempty"`

	// The end date in seconds
	EndDate int64 `json:"end_date,omitempty"`

	// The id of the challenge event
	Id int64 `json:"id,omitempty"`

	// Indicate if the rewards have been given out already 
	RewardStatus string `json:"reward_status,omitempty"`

	// The start date in seconds
	StartDate int64 `json:"start_date,omitempty"`
}
