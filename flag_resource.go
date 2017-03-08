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

type FlagResource struct {

	// The context of that resource
	Context string `json:"context"`

	// The context_id of that resource
	ContextId string `json:"context_id"`

	// The date/time this resource was created in seconds since epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The unique ID for that resource
	Id int64 `json:"id,omitempty"`

	// The flag reason of that resource
	Reason string `json:"reason,omitempty"`

	// The date/time this resource was last updated in seconds since epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The basic user resource
	User SimpleUserResource `json:"user,omitempty"`
}
