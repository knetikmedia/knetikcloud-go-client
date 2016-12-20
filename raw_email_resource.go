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

type RawEmailResource struct {

	// The body of the outgoing message.
	Body string `json:"body,omitempty"`

	// Address to attribute the outgoing message to. Optional if the config email.out_address is set.
	From string `json:"from,omitempty"`

	// Whether the body is to be treated as html. Default false.
	Html bool `json:"html,omitempty"`

	// A list of user ids to send the message to.
	Recipients []int32 `json:"recipients,omitempty"`

	// The subject of the outgoing message.
	Subject string `json:"subject,omitempty"`
}
