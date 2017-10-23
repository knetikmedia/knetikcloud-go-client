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

type UserAchievementGroupResource struct {

	// The list of achievements associated with the group
	Achievements []UserAchievementResource `json:"achievements"`

	// The name of the group.  If used by Leveling, this will represent the level name
	GroupName string `json:"group_name"`

	// The id of the achievement progress
	Id string `json:"id,omitempty"`

	// The current progress of the user on the group
	Progress int32 `json:"progress"`

	// The id of the user whose progress is being tracked
	UserId int32 `json:"user_id"`
}
