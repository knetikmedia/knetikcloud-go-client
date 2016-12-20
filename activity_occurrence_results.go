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

type ActivityOccurrenceResults struct {

	// The game results for each user. Include all users that played (paid to get in) even if they were eliminated without a result. A null metric is allowed
	Users []UserActivityResultsResource `json:"users,omitempty"`
}
