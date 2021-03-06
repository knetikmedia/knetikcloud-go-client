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

type CreatePayPalPaymentRequest struct {

	// The endpoint URL to which PayPal should forward the user to if they cancel the checkout process
	CancelUrl string `json:"cancel_url"`

	// The ID of an invoice to pay
	InvoiceId int32 `json:"invoice_id"`

	// The endpoint URL to which PayPal should forward the user after they accept. This endpoint will receive information needed for the next step
	ReturnUrl string `json:"return_url"`
}
