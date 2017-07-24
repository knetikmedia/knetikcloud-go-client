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

type XsollaPaymentRequest struct {

	// The id of an invoice to pay
	InvoiceId int32 `json:"invoice_id"`

	// The endpoint URL xsolla should forward the user to after they pay
	ReturnUrl string `json:"return_url"`
}
