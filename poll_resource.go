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

type PollResource struct {

	// Whether the poll is active
	Active bool `json:"active,omitempty"`

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The answers to the poll
	Answers []PollAnswerResource `json:"answers,omitempty"`

	// The category for the poll
	Category NestedCategory `json:"category,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The id of the poll
	Id string `json:"id,omitempty"`

	// The tags for the poll
	Tags []string `json:"tags,omitempty"`

	// A poll template this poll is validated against (private). May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`

	// The text of the poll
	Text string `json:"text,omitempty"`

	// The media type of the poll
	Type_ string `json:"type,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
