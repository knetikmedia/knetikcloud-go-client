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

type PaymentAuthorizationResource struct {

	// Whether this authorization has been captured
	Captured bool `json:"captured,omitempty"`

	// The date this authorization was received, unix timestamp in seconds
	Created int64 `json:"created,omitempty"`

	// The details for this authorization. Format dependent on payment provider
	Details *interface{} `json:"details,omitempty"`

	// The id of the authorization
	Id int32 `json:"id,omitempty"`

	// The invoice this authorization is intended to pay
	Invoice int32 `json:"invoice,omitempty"`

	// The payment type (which provider) this payment is through
	PaymentType *SimpleReferenceResourceint `json:"payment_type"`
}
