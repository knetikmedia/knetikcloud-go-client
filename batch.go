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

type Batch struct {

	// The list of batch requests
	Batch []BatchRequest `json:"batch"`

	// The amount of time before a request token is returned instead of the batch result.  Default is 60.  Range is 0-300
	Timeout int32 `json:"timeout"`
}
