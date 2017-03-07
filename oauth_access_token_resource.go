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

type OauthAccessTokenResource struct {

	// The key of the client assosciated with the token
	ClientId string `json:"client_id,omitempty"`

	// The token.  Not shown in list view
	Token string `json:"token,omitempty"`

	// The username of the user associated with the token
	Username string `json:"username,omitempty"`
}