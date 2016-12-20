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

type Timezone struct {

	Code string `json:"code,omitempty"`

	Id int32 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Offset float64 `json:"offset,omitempty"`
}
