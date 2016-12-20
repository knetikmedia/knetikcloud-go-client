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

type SelectedSettingResource struct {

	// The unique ID for the setting
	Key string `json:"key,omitempty"`

	// The textual name of the setting
	KeyName string `json:"key_name,omitempty"`

	// The unique ID for the option
	Value string `json:"value,omitempty"`

	// The textual name of the option
	ValueName string `json:"value_name,omitempty"`
}
