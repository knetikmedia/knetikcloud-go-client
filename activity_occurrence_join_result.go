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

type ActivityOccurrenceJoinResult struct {

	// The details on the entitlement object needed to enter the occurrence (if any)
	Entitlement ActivityEntitlementResource `json:"entitlement,omitempty"`

	// Zero if the user was/could be added to the occurrence. Jsapi error code indicating the reason of the failure otherwise
	ErrorCode int32 `json:"error_code,omitempty"`

	// An error message if failure
	Message string `json:"message,omitempty"`

	// The user's id
	UserId int64 `json:"user_id,omitempty"`
}
