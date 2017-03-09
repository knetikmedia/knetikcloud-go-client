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

type SearchReferenceMapping struct {

	// Unique identifier for the mapping to protect against duplicates
	Id string `json:"id"`

	// The field within the type that contains the id from the refType
	RefIdField string `json:"ref_id_field"`

	// The index type that the mapping pulls data from
	RefType string `json:"ref_type"`

	// A map whose keys are the field names in the refType and values are the field name in the type
	SourceFieldToDestinationField map[string]string `json:"source_field_to_destination_field"`

	// The index type that the mapping is for
	Type_ string `json:"type"`
}
