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

type RoleResource struct {

	// The number of clients this role is assigned to
	ClientCount int32 `json:"client_count,omitempty"`

	// The date the role was added. Unix timestamp in seconds
	CreatedDate int64 `json:"created_date,omitempty"`

	// Whether a role is locked from being deleted
	Locked bool `json:"locked,omitempty"`

	// The name of the role used for display purposes
	Name string `json:"name,omitempty"`

	// The keyword that defines the role
	Role string `json:"role,omitempty"`

	// The list of permissions this role has
	RolePermission []PermissionResource `json:"role_permission,omitempty"`

	// The number of users this role is assigned to
	UserCount int32 `json:"user_count,omitempty"`
}
