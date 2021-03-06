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

type AddressResource struct {

	// The first line of the address
	Address1 string `json:"address1"`

	// A second line of the address
	Address2 string `json:"address2,omitempty"`

	// The city
	City string `json:"city"`

	// The iso3 code for the country
	CountryCode string `json:"country_code"`

	// The postal code
	PostalCode string `json:"postal_code,omitempty"`

	// The code for the state. Required if the country has states/provinces/equivalent
	StateCode string `json:"state_code,omitempty"`
}
