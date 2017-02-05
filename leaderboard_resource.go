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

type LeaderboardResource struct {

	// The paginated list of user results, in order from best to worst
	Entries []LeaderboardEntryResource `json:"entries,omitempty"`

	// The id of the leaderboard
	Id int64 `json:"id,omitempty"`

	// The name of the strategy that defines how entries are stored and compared
	Strategy string `json:"strategy,omitempty"`
}
