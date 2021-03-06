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

type DispositionCount struct {

	// The number of users that have expressed this disposition
	Count int64 `json:"count,omitempty"`

	// The name of the disposition this count is for
	Name string `json:"name,omitempty"`
}
