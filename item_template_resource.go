/* 
 * Knetik Platform API Documentation Latest
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ItemTemplateResource struct {

	// The customized behaviors that are required or default for this type of item
	Behaviors []ItemBehaviorDefinitionResource `json:"behaviors,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The id of the template
	Id string `json:"id,omitempty"`

	// The name of the template
	Name string `json:"name,omitempty"`

	// The customized properties that are present
	Properties []PropertyDefinitionResource `json:"properties,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
