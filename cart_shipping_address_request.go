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

type CartShippingAddressRequest struct {

	// The city of the user
	City string `json:"city,omitempty"`

	// The country code of the user
	CountryCodeIso3 string `json:"country_code_iso3,omitempty"`

	// The email of the user
	Email string `json:"email,omitempty"`

	// The first name of the user
	FirstName string `json:"first_name,omitempty"`

	// The last name of the user
	LastName string `json:"last_name,omitempty"`

	NamePrefix string `json:"name_prefix,omitempty"`

	// The order notes the user
	OrderNotes string `json:"order_notes,omitempty"`

	// The phone number of the user
	PhoneNumber string `json:"phone_number,omitempty"`

	// The postal state code of the user
	PostalStateCode string `json:"postal_state_code,omitempty"`

	// The shipping address of the user, first line
	ShippingAddressLine1 string `json:"shipping_address_line1,omitempty"`

	// The shipping address of the user, second line
	ShippingAddressLine2 string `json:"shipping_address_line2,omitempty"`

	// The zipcode of the user
	Zip string `json:"zip,omitempty"`
}
