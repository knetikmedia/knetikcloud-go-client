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

type UserActivityResults struct {

	// Any currency rewarded to this user
	CurrencyRewards []RewardCurrencyResource `json:"currency_rewards,omitempty"`

	// Any items rewarded to this user
	ItemRewards []RewardItemResource `json:"item_rewards,omitempty"`

	// The position of the user in the leaderboard. Null means non-compete or disqualification
	Rank int64 `json:"rank,omitempty"`

	// The raw score in this leaderboard. Null means non-compete or disqualification
	Score int64 `json:"score,omitempty"`

	// Any tags for the metric. Each unique tag will translate into a unique leaderboard. Maximum 10 tags and 50 characters each
	Tags []string `json:"tags,omitempty"`

	// The number of users tied at this rank, including this user. 1 means no tie
	Ties int32 `json:"ties,omitempty"`

	// The date this score was recorded or updated. Unix timestamp in seconds
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The player for this entry
	User *SimpleUserResource `json:"user"`
}
