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

type OAuth2Resource struct {

	// The access token issued by the authorization server
	AccessToken string `json:"access_token,omitempty"`

	// The lifetime in seconds of the access token
	ExpiresIn string `json:"expires_in,omitempty"`

	// The scope of the access token. Currently these values can be ignored, as security defaults to roles and permissions
	Scope string `json:"scope,omitempty"`

	// The type of the token issued
	TokenType string `json:"token_type,omitempty"`
}
