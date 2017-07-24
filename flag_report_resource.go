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

type FlagReportResource struct {

	// The context of that resource 
	Context string `json:"context,omitempty"`

	// The context ID of that resource
	ContextId string `json:"context_id,omitempty"`

	// The date/time this resource was created in seconds since epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The unique ID for that resource
	Id int64 `json:"id,omitempty"`

	// The reason of that resource required only in case of active resolution
	Reason string `json:"reason,omitempty"`

	// The resolution of that resource
	Resolution string `json:"resolution"`

	// The date/time this report was resolved in seconds since epoch. Null if not resolved yet
	Resolved int64 `json:"resolved,omitempty"`

	// The date/time this resource was last updated in seconds since epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
