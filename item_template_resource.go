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

type ItemTemplateResource struct {

	// Whether to allow additional properties beyond those specified or not
	AllowAdditional bool `json:"allow_additional,omitempty"`

	// The customized behaviors that are required or default for this type of item
	Behaviors []ItemBehaviorDefinitionResource `json:"behaviors,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The id of the template
	Id string `json:"id,omitempty"`

	// The name of the template
	Name string `json:"name"`

	// The customized properties that are present
	Properties []PropertyDefinitionResource `json:"properties,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
