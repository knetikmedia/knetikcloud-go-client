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

type QuestionResource struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The list of available answers
	Answers []AnswerResource `json:"answers,omitempty"`

	// The category for the question
	Category *NestedCategory `json:"category"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The difficulty of the question
	Difficulty int32 `json:"difficulty"`

	// The unique ID for that resource
	Id string `json:"id,omitempty"`

	// The id of the import job that created the question, or null if not from an import
	ImportId int64 `json:"import_id,omitempty"`

	// When the question becomes available, null for never, in seconds since epoch
	PublishedDate int64 `json:"published_date,omitempty"`

	// The question. Different 'type' values indicate different structures as the question may be test, image, etc. See information on additional properties for the list and their structures
	Question *Property `json:"question"`

	// The first source of the question
	Source1 string `json:"source1,omitempty"`

	// The second source of the question
	Source2 string `json:"source2,omitempty"`

	// The list of tags
	Tags []string `json:"tags,omitempty"`

	// A question template this question is validated against (private). May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The supplier of the question
	Vendor string `json:"vendor,omitempty"`
}
