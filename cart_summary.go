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

type CartSummary struct {

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The unique id code for the currency used in the cart
	CurrencyCode string `json:"currency_code,omitempty"`

	// The grand total for the cart
	GrandTotal float64 `json:"grand_total,omitempty"`

	// The unique ID for the cart
	Id string `json:"id,omitempty"`

	// The ID of the invoice associated with this cart
	InvoiceId float64 `json:"invoice_id,omitempty"`

	// The number of items in the cart
	ItemsInCart int32 `json:"items_in_cart,omitempty"`

	// The status of the cart
	Status string `json:"status,omitempty"`

	// The subtotal of all items in the cart
	Subtotal float64 `json:"subtotal,omitempty"`
}
