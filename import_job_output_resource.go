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

type ImportJobOutputResource struct {

	// The description of the import job
	Description string `json:"description,omitempty"`

	// The line number of the import job
	LineNumber int64 `json:"line_number,omitempty"`
}