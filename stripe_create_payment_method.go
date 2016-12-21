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

type StripeCreatePaymentMethod struct {

	// A token from Stripe to identify payment info to be tied to the customer
	Token string `json:"token,omitempty"`

	// The id of the user, if null the logged in user is used. Admin privilege need to specify other users
	UserId int32 `json:"user_id,omitempty"`
}
