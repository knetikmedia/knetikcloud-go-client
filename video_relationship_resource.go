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

type VideoRelationshipResource struct {

	// The owner of the relationship
	From SimpleReferenceResourcelong `json:"from,omitempty"`

	// The id of the relationship
	Id int64 `json:"id,omitempty"`

	// Details about the relationship such as type or other information. Max length 10 characters
	RelationshipDetails string `json:"relationship_details"`

	// The target of the relationship.
	To SimpleReferenceResourcelong `json:"to"`
}