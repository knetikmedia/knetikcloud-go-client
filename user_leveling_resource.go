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

type UserLevelingResource struct {

	// The name of the last tier the user has qualified for
	LastTierName string `json:"last_tier_name"`

	// The progress level of the last tier the user has qualified for
	LastTierProgress int32 `json:"last_tier_progress"`

	// The name of the level schema
	LevelName string `json:"level_name"`

	// The name of the next tier the user can qualify for
	NextTierName string `json:"next_tier_name"`

	// The progress needed to qualify for the next tier
	NextTierProgress int32 `json:"next_tier_progress"`

	// The amount of progress the user has
	Progress int32 `json:"progress"`

	// The names of the tiers the user has qualified for
	TierNames []string `json:"tier_names"`

	// The ID of the user
	UserId int32 `json:"user_id"`
}
