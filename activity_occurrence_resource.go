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

type ActivityOccurrenceResource struct {

	// The id of the activity
	ActivityId int64 `json:"activity_id"`

	// The id of the challenge activity (as part of the event, required if eventId set)
	ChallengeActivityId int64 `json:"challenge_activity_id,omitempty"`

	// The date this occurrence was created, unix timestamp in seconds
	CreatedDate int64 `json:"created_date,omitempty"`

	// The entitlement item required to enter the occurrence. Required if not part of an event. Must come from the set of entitlement items listed in the activity
	Entitlement ActivityEntitlementResource `json:"entitlement,omitempty"`

	// The id of the event
	EventId int64 `json:"event_id,omitempty"`

	// The id of the activity occurrence
	Id int64 `json:"id,omitempty"`

	// Indicate if the rewards have been given out already
	RewardStatus string `json:"reward_status,omitempty"`

	// The list of settings and their options available for this activity. Should be null on create if and only if part of an event
	Settings []SelectedSettingResource `json:"settings,omitempty"`

	// Whether this occurrence will be played as a simulation.
	Simulated bool `json:"simulated,omitempty"`

	// The date this occurrence was started, unix timestamp in seconds. null if not yet started
	StartDate int64 `json:"start_date,omitempty"`

	// The current status of the occurrence (default: OPEN)
	Status string `json:"status,omitempty"`

	// The date this occurrence was last updated, unix timestamp in seconds
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The list of users playing in this occurrence. Can only be set directly with ACTIVITIES_ADMIN permission
	Users []ActivityUserResource `json:"users,omitempty"`
}
