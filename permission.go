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

type Permission struct {

	CreatedDate int64 `json:"created_date,omitempty"`

	Description string `json:"description,omitempty"`

	Id int32 `json:"id,omitempty"`

	Locked bool `json:"locked,omitempty"`

	Name string `json:"name,omitempty"`

	Parent string `json:"parent,omitempty"`

	Permission string `json:"permission,omitempty"`

	UpdatedDate int64 `json:"updated_date,omitempty"`
}
