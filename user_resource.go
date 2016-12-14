/* 
 * Knetik Platform API Documentation Latest
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type UserResource struct {

	// A map of additional properties, keyed on the property name (private). Must match the names and types defined in the template for this user type, or be an extra not from the template
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The first line of the user's address (private)
	Address string `json:"address,omitempty"`

	// The second line of user's address (private)
	Address2 string `json:"address2,omitempty"`

	// The url of the user's avatar image
	AvatarUrl string `json:"avatar_url,omitempty"`

	// The user's city (private)
	City string `json:"city,omitempty"`

	// The ISO3 code for the country from the user's address (private). Will be filled in based on GeoIP country at registration if not provided.
	CountryCode string `json:"country_code,omitempty"`

	// The code for the user's real money currency (private)
	CurrencyCode string `json:"currency_code,omitempty"`

	// The user's date of birth (private) as a unix timestamp
	DateOfBirth int64 `json:"date_of_birth,omitempty"`

	// The user's self description (private)
	Description string `json:"description,omitempty"`

	// The chosen display name of the user, defaults to username if not present
	DisplayName string `json:"display_name,omitempty"`

	// The user's email address (private). May be required and/or unique depending on system configuration (both on by default). Must match standard email requirements if provided (RFC 2822)
	Email string `json:"email,omitempty"`

	// The user's first name (private)
	FirstName string `json:"first_name,omitempty"`

	// The user's full name (private)
	Fullname string `json:"fullname,omitempty"`

	// The user's gender (private)
	Gender string `json:"gender,omitempty"`

	// The id of the user
	Id int32 `json:"id,omitempty"`

	// The ISO3 code for the user's currency (private)
	LanguageCode string `json:"language_code,omitempty"`

	// The user's last name (private)
	LastName string `json:"last_name,omitempty"`

	// The user's mobile phone number (private)
	MobileNumber string `json:"mobile_number,omitempty"`

	// The plain text password for the new user account. Required for registration; ignored on profile update.  Use password specific endpoints for editing
	Password string `json:"password,omitempty"`

	// The user's postal code (private)
	PostalCode string `json:"postal_code,omitempty"`

	// The user's state (private)
	State string `json:"state,omitempty"`

	// A user template this user is validated against (private). May be null and no validation of properties will be done
	Template string `json:"template,omitempty"`

	// The code for the user's timezone (private)
	TimezoneCode string `json:"timezone_code,omitempty"`

	// The login username for the user (private). May be set to match email if system does not require usernames separately.
	Username string `json:"username,omitempty"`
}
