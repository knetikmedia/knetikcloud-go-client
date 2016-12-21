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

type LeaderboardEntryResource struct {

	// The position of the user in the leaderboard. Null means non-compete or disqualification
	Rank int64 `json:"rank,omitempty"`

	// The raw score in this leaderboard. Null means non-compete or disqualification
	Score int64 `json:"score,omitempty"`

	// The player for this entry
	User SimpleUserResource `json:"user,omitempty"`
}
