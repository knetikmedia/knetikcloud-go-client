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

type ChatThreadResource struct {

	// The number of active users in the thread
	ActiveUsers int32 `json:"active_users,omitempty"`

	// The number of messages in the thread
	Count int32 `json:"count,omitempty"`

	// The date the thread was created
	CreatedDate int64 `json:"created_date,omitempty"`

	// The id for this thread
	Id string `json:"id,omitempty"`

	// The id of the recipient
	RecipientId string `json:"recipient_id"`

	// The recipient type of the thread
	RecipientType string `json:"recipient_type"`

	// The subject of the thread
	Subject string `json:"subject,omitempty"`

	// The date the thread was updated
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
