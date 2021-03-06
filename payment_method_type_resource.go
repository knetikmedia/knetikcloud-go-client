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

type PaymentMethodTypeResource struct {

	// The id of the payment method type
	Id int32 `json:"id"`

	// The maximum timelimit in hours for an invoice in the processing status while waiting on this payment method type. Defaults to the global config invoice.processing_expiration_hours if null
	InvoiceProcessingHours int32 `json:"invoice_processing_hours,omitempty"`

	// The name of the payment method type
	Name string `json:"name"`

	// Whether the payment handler supports the authorize and capture flow
	SupportsCapture bool `json:"supports_capture,omitempty"`

	// Whether the payment handler supports paying for part of an invoice, rather than the full grand_total
	SupportsPartial bool `json:"supports_partial,omitempty"`

	// Whether the payment handler supports rebilling the method later (for saved payments or subscriptions)
	SupportsRebill bool `json:"supports_rebill,omitempty"`

	// Whether the payment handler supports refunding
	SupportsRefunds bool `json:"supports_refunds,omitempty"`
}
