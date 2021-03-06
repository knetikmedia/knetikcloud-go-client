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

type TokenDetailsResource struct {

	ClientId string `json:"client_id,omitempty"`

	Roles []string `json:"roles,omitempty"`

	UserId int32 `json:"user_id,omitempty"`
}
