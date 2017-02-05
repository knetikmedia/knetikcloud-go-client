/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type GroupResource struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// A description of the group. Max 250 characters
	Description string `json:"description,omitempty"`

	// The number of users in the group
	MemberCount int32 `json:"member_count,omitempty"`

	// A message of the day for members of the group
	MessageOfTheDay string `json:"message_of_the_day,omitempty"`

	// The name of the group. Max 50 characters
	Name string `json:"name"`

	// The unique name of another group that this group is a subset of
	Parent string `json:"parent,omitempty"`

	// The status which describes whether other users can freely join the group or not
	Status string `json:"status"`

	// The number of users in child groups
	SubMemberCount int32 `json:"sub_member_count,omitempty"`

	// A group template this group is validated against. May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`

	// Unique name used in url and references. Uppercase, lowercase, numbers and hyphens only. Max 50 characters. Cannot be altered once created
	UniqueName string `json:"unique_name"`
}
