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

type UserBaseResource struct {

	// The url of the user's avatar image
	AvatarUrl string `json:"avatar_url,omitempty"`

	// The chosen display name of the user, defaults to username if not present
	DisplayName string `json:"display_name,omitempty"`

	// The user's email address (private). May be required and/or unique depending on system configuration (both on by default). Must match standard email requirements if provided (RFC 2822)
	Email string `json:"email,omitempty"`

	// The user's full name (private)
	Fullname string `json:"fullname,omitempty"`

	// The id of the user
	Id int32 `json:"id,omitempty"`

	// The login username for the user (private). May be set to match email if system does not require usernames separately.
	Username string `json:"username,omitempty"`
}