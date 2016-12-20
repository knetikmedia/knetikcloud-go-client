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

type Group struct {

	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	Description string `json:"description,omitempty"`

	Id int32 `json:"id,omitempty"`

	MemberCount int32 `json:"member_count,omitempty"`

	MessageOfTheDay string `json:"message_of_the_day,omitempty"`

	Name string `json:"name,omitempty"`

	Parent Group `json:"parent,omitempty"`

	PropertiesString string `json:"properties_string,omitempty"`

	Status string `json:"status,omitempty"`

	SubMemberCount int32 `json:"sub_member_count,omitempty"`

	Template string `json:"template,omitempty"`

	UniqueName string `json:"unique_name,omitempty"`
}
