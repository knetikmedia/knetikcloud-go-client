/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Country struct {

	Id int32 `json:"id,omitempty"`

	Iso2 string `json:"iso2,omitempty"`

	Iso3 string `json:"iso3,omitempty"`

	Name string `json:"name,omitempty"`
}
