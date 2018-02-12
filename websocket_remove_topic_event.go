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

type WebsocketRemoveTopicEvent struct {

	Client string `json:"client,omitempty"`

	Customer string `json:"customer,omitempty"`

	DoNotBroadcast bool `json:"do_not_broadcast,omitempty"`

	Section string `json:"section,omitempty"`

	Source *interface{} `json:"source,omitempty"`

	Specifics string `json:"specifics,omitempty"`

	Synchronous bool `json:"synchronous,omitempty"`

	Timestamp int64 `json:"timestamp,omitempty"`

	// The type of the event. Used for polymorphic type recognition and thus must match an expected type
	Type_ string `json:"type"`

	Topic *Topic `json:"topic,omitempty"`
}
