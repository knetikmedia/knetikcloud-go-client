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

type CampaignResource struct {

	// Whether the campaign is active or not.  Defaults to false
	Active bool `json:"active,omitempty"`

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The unique ID for that resource
	Id int64 `json:"id,omitempty"`

	// The strategy for calculating the leaderboard. Defaults to highest score. Value MUST come from the list of available strategies from the Leaderboard Service
	LeaderboardStrategy string `json:"leaderboard_strategy,omitempty"`

	// The user friendly name of that resource. Defaults to blank string
	LongDescription string `json:"long_description,omitempty"`

	// The user friendly name of that resource
	Name string `json:"name"`

	// The name of the next challenge coming up
	NextChallenge string `json:"next_challenge,omitempty"`

	// The date/time of the next challenge coming up
	NextChallengeDate int64 `json:"next_challenge_date,omitempty"`

	// The rewards to give at the end of the campaign. When creating/updating only id is used. Reward set must be pre-existing
	RewardSet RewardSetResource `json:"reward_set,omitempty"`

	// Indicate if the rewards have been given out already
	RewardStatus string `json:"reward_status,omitempty"`

	// The user friendly name of that resource. Defaults to blank string
	ShortDescription string `json:"short_description,omitempty"`

	// A campaign template this campaign is validated against (private). May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
