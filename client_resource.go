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

type ClientResource struct {

	// The time limit of the token in seconds
	AccessTokenValiditySeconds int32 `json:"access_token_validity_seconds,omitempty"`

	// The client key, cannot be edited once created
	ClientKey string `json:"client_key,omitempty"`

	// The grant types of the client
	GrantTypes []string `json:"grant_types,omitempty"`

	// The id of the client
	Id int32 `json:"id,omitempty"`

	// Whether the client is public or not
	IsPublic bool `json:"is_public,omitempty"`

	// Whether a client is locked from being deleted
	Locked bool `json:"locked,omitempty"`

	// The name of the client
	Name string `json:"name,omitempty"`

	// The redirect uris of the client
	RedirectUris []string `json:"redirect_uris,omitempty"`

	// The time limit of the refresh token in seconds
	RefreshTokenValiditySeconds int32 `json:"refresh_token_validity_seconds,omitempty"`

	// The secret key of the client
	Secret string `json:"secret,omitempty"`
}