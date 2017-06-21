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

type ImageGroupPropertyDefinitionResource struct {

	// A list of the fields on both the property definition and property of this type
	FieldList PropertyFieldListResource `json:"field_list,omitempty"`

	// The name of the property
	Name string `json:"name"`

	// Whether the property is required
	Required bool `json:"required"`

	// The type of the property. Used for polymorphic type recognition and thus must match an expected type with additional properties.
	Type_ string `json:"type"`

	// If provided, a file type that the property must match
	FileType string `json:"file_type,omitempty"`

	// If provided, the maximum number of files in the group
	MaxCount int32 `json:"max_count,omitempty"`

	// If provided, the maximum allowed size per file in bytes
	MaxFileSize int64 `json:"max_file_size,omitempty"`

	// If provided, the minimum number of files in the group
	MinCount int32 `json:"min_count,omitempty"`

	// If provided, the maximum height of each image
	MaxHeight int32 `json:"max_height,omitempty"`

	// If provided, the maximum width of each image
	MaxWidth int32 `json:"max_width,omitempty"`

	// If provided, the minimum height of each image
	MinHeight int32 `json:"min_height,omitempty"`

	// If provided, the minumum width of each image
	MinWidth int32 `json:"min_width,omitempty"`
}
