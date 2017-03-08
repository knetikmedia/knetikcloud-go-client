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

type GroupMemberResource struct {

	// The url of the user's avatar image
	AvatarUrl string `json:"avatar_url,omitempty"`

	// The public username of the user
	DisplayName string `json:"display_name,omitempty"`

	// The id of the user
	Id int32 `json:"id"`

	// The member's access level. Default: member
	Status string `json:"status"`

	// The username of the user
	Username string `json:"username,omitempty"`
}
