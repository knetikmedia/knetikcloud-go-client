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

type CountryResource struct {

	// The iso2 of the country
	Iso2 string `json:"iso2,omitempty"`

	// The iso3 of the country
	Iso3 string `json:"iso3,omitempty"`

	// The name of the country resource
	Name string `json:"name,omitempty"`
}
