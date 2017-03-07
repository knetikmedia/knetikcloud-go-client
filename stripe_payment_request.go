/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type StripePaymentRequest struct {

	// The id of the invoice to pay
	InvoiceId int32 `json:"invoice_id"`

	// A token from Stripe to identify payment info to be tied to the customer
	Token string `json:"token"`
}