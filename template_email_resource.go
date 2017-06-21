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

type TemplateEmailResource struct {

	// Address to attribute the outgoing message to. Optional if the config email.out_address is set.
	From string `json:"from,omitempty"`

	// A list of user ids to send the message to.
	Recipients []int32 `json:"recipients"`

	// The key for the template
	TemplateKey string `json:"template_key"`

	// A list of variables to fill in the template
	TemplateVars []KeyValuePairstringstring `json:"template_vars,omitempty"`
}
