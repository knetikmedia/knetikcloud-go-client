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

type EntitlementItem struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type, or be an extra not from the template
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The behaviors linked to the item, describing various options and interactions. May not be included in item lists
	Behaviors []Behavior `json:"behaviors,omitempty"`

	// A category for filtering items
	Category string `json:"category,omitempty"`

	// The date the item was created, unix timestamp in seconds
	CreatedDate int64 `json:"created_date,omitempty"`

	// The id of the item
	Id int32 `json:"id,omitempty"`

	// A long description of the item
	LongDescription string `json:"long_description,omitempty"`

	// The name of the item
	Name string `json:"name"`

	// A short description of the item, max 255 chars
	ShortDescription string `json:"short_description,omitempty"`

	// A number to use in sorting items.  Default 500
	Sort int32 `json:"sort,omitempty"`

	// List of tags used for filtering items
	Tags []string `json:"tags,omitempty"`

	// An item template this item is validated against.  May be null and no validation of additional_properties will be done.  Default = null
	Template string `json:"template,omitempty"`

	// The type of the item
	TypeHint string `json:"type_hint"`

	// The unique key for the item
	UniqueKey string `json:"unique_key,omitempty"`

	// The date the item was last updated, unix timestamp in seconds
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
