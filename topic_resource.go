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

type TopicResource struct {

	// The created date of the topic
	CreatedDate int64 `json:"created_date,omitempty"`

	// The display name of the topic
	DisplayName string `json:"display_name,omitempty"`

	// The unique ID for this topic
	Id string `json:"id,omitempty"`

	// Whether this topic is locked or not.
	Locked bool `json:"locked,omitempty"`

	// Random tags to facilitate search
	Tags []string `json:"tags,omitempty"`

	// The last time the topic was updated
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The subscribed user count of the topic
	UserCount int32 `json:"user_count,omitempty"`
}
