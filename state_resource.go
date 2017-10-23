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

type StateResource struct {

	// The code of the state
	Code string `json:"code,omitempty"`

	// The iso3 of the country this state is in
	CountryCodeIso3 string `json:"country_code_iso3,omitempty"`

	// The unique ID for the state
	Id int32 `json:"id,omitempty"`

	// The name of the state
	Name string `json:"name,omitempty"`
}
