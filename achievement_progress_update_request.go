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

type AchievementProgressUpdateRequest struct {

	// Whether to add one to the current progress instead of setting it to progress_value. Default: false
	IncrementValue bool `json:"increment_value,omitempty"`

	// The amount of progress towards earning the achievement. The max/target depends on the achievement. Required if increment_value is false/missing.
	ProgressValue int32 `json:"progress_value,omitempty"`
}
