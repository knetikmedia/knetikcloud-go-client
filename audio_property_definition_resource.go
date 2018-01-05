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

type AudioPropertyDefinitionResource struct {

	// The description of the property
	Description string `json:"description,omitempty"`

	// A list of the fields on both the property definition and property of this type
	FieldList *PropertyFieldListResource `json:"field_list,omitempty"`

	// The friendly front-facing name of the property
	FriendlyName string `json:"friendly_name,omitempty"`

	// The name of the property
	Name string `json:"name"`

	// The JSON path to the option label
	OptionLabelPath string `json:"option_label_path,omitempty"`

	// The JSON path to the option value
	OptionValuePath string `json:"option_value_path,omitempty"`

	// URL of service containing the property options (assumed JSON array)
	OptionsUrl string `json:"options_url,omitempty"`

	// Whether the property is required
	Required bool `json:"required"`

	// The type of the property. Used for polymorphic type recognition and thus must match an expected type with additional properties.
	Type_ string `json:"type"`

	// If provided, a file type the property must match
	FileType string `json:"file_type,omitempty"`

	// If provided, the maximum length of the audio
	MaxLength int32 `json:"max_length,omitempty"`

	// If provided, the minimum length of the audio
	MinLength int32 `json:"min_length,omitempty"`
}
