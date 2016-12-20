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

type Role struct {

	ClientCount int32 `json:"client_count,omitempty"`

	CreatedDate int64 `json:"created_date,omitempty"`

	Id int32 `json:"id,omitempty"`

	Locked bool `json:"locked,omitempty"`

	Name string `json:"name,omitempty"`

	Role string `json:"role,omitempty"`

	RolePermission []Permission `json:"role_permission,omitempty"`

	UserCount int32 `json:"user_count,omitempty"`
}
