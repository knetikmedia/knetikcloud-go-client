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

type FileProperty struct {

	// The type of the property. Used for polymorphic type recognition and thus must match an expected type with additional properties.
	Type_ string `json:"type"`

	// A crc value for file integrity verification
	Crc string `json:"crc,omitempty"`

	// A description of the file
	Description string `json:"description,omitempty"`

	// The type of file such as txt, mp3, mov or csv
	FileType string `json:"file_type,omitempty"`

	// The url of the file
	Url string `json:"url,omitempty"`
}