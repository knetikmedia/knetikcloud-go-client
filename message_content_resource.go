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

type MessageContentResource struct {

	// The content of the email
	Email string `json:"email,omitempty"`

	// The content of the mobile app push notification
	Push string `json:"push,omitempty"`

	// The content of the sms
	Sms string `json:"sms,omitempty"`

	// The content of the templated email
	TemplatedEmail *TemplatedEmail `json:"templated_email,omitempty"`

	// The content of the websocket message
	Websocket *interface{} `json:"websocket,omitempty"`
}
