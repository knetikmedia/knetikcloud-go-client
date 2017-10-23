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

type InvoiceResource struct {

	// Line one of the customer's billing address
	BillingAddress1 string `json:"billing_address1,omitempty"`

	// Line two of the customer's billing address
	BillingAddress2 string `json:"billing_address2,omitempty"`

	// The city for the customer's billing address
	BillingCityName string `json:"billing_city_name,omitempty"`

	// The country for the customer's billing address
	BillingCountryName string `json:"billing_country_name,omitempty"`

	// The customer's name for the billing address
	BillingFullName string `json:"billing_full_name,omitempty"`

	// The postal code for the customer's billing address
	BillingPostalCode string `json:"billing_postal_code,omitempty"`

	// The state for the customer's billing address
	BillingStateName string `json:"billing_state_name,omitempty"`

	// The guid of the cart this invoice came from
	CartId string `json:"cart_id,omitempty"`

	// The date the invoice was created, unix timestamp in seconds
	CreatedDate int64 `json:"created_date,omitempty"`

	// The code for the currency invoice prices are in
	Currency string `json:"currency,omitempty"`

	// The fulfillment status of the invoice
	CurrentFulfillmentStatus string `json:"current_fulfillment_status,omitempty"`

	// The payment status of the invoice
	CurrentPaymentStatus string `json:"current_payment_status,omitempty"`

	// The amount of money saved through coupons
	Discount float32 `json:"discount,omitempty"`

	// The customer's email address
	Email string `json:"email,omitempty"`

	// An external reference to filter on
	ExternalRef string `json:"external_ref,omitempty"`

	// The amount of federal tax added
	FedTax float32 `json:"fed_tax,omitempty"`

	// The final price of the invoice
	GrandTotal float32 `json:"grand_total,omitempty"`

	// The id of the invoice
	Id int32 `json:"id,omitempty"`

	// A reference number for the invoice
	InvoiceNumber string `json:"invoice_number,omitempty"`

	// A list of items within the invoice
	Items []InvoiceItemResource `json:"items,omitempty"`

	// The customer's name prefix
	NamePrefix string `json:"name_prefix,omitempty"`

	// Notes about the order
	OrderNotes string `json:"order_notes,omitempty"`

	// The id of an invoice this is a child of
	ParentInvoiceId int32 `json:"parent_invoice_id,omitempty"`

	// The id of a saved payment method used to pay for the invoice
	PaymentMethodId int32 `json:"payment_method_id,omitempty"`

	// The customer's phone number
	Phone string `json:"phone,omitempty"`

	// The customer's phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// The remaining price of the invoice (after any payments made so far)
	RemainingBalance float32 `json:"remaining_balance,omitempty"`

	// The shipping cost
	Shipping float32 `json:"shipping,omitempty"`

	// Line one of the customer's shipping address
	ShippingAddress1 string `json:"shipping_address1,omitempty"`

	// Line two of the customer's shipping address
	ShippingAddress2 string `json:"shipping_address2,omitempty"`

	// The city for the customer's shipping address
	ShippingCityName string `json:"shipping_city_name,omitempty"`

	// The country for the customer's shipping address
	ShippingCountryName string `json:"shipping_country_name,omitempty"`

	// The customer's name for the shipping address
	ShippingFullName string `json:"shipping_full_name,omitempty"`

	// The postal code for the customer's shipping address
	ShippingPostalCode string `json:"shipping_postal_code,omitempty"`

	// The state for the customer's shipping address
	ShippingStateName string `json:"shipping_state_name,omitempty"`

	// A number to use in sorting items. default 500.
	Sort int32 `json:"sort,omitempty"`

	// The amount of state tax added
	StateTax float32 `json:"state_tax,omitempty"`

	// The sum price of all items before shipping, coupons and tax
	Subtotal float32 `json:"subtotal,omitempty"`

	// The date the invoice was last updated, unix timestamp in seconds
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The owner of the invoice
	User *SimpleUserResource `json:"user,omitempty"`

	// The id of the vendor
	VendorId int32 `json:"vendor_id,omitempty"`

	// The name of the invoice
	VendorName string `json:"vendor_name,omitempty"`
}
