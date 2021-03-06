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

type Sku struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type, or be an extra not from the template
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The currency code for the SKU, a three letter string (ISO3)
	CurrencyCode string `json:"currency_code"`

	// The friendly name of the SKU as it will appear on invoices and reports. Typically represents the option name like red, large, etc
	Description string `json:"description"`

	// The number of SKUs currently in stock
	Inventory int32 `json:"inventory,omitempty"`

	// Alerts vendor when SKU inventory drops below this value
	MinInventoryThreshold int32 `json:"min_inventory_threshold,omitempty"`

	NotAvailable bool `json:"not_available,omitempty"`

	NotDisplayable bool `json:"not_displayable,omitempty"`

	// The base price before any sale
	OriginalPrice float32 `json:"original_price"`

	// The current price of the SKU with sales, if any. Set original_price for the base
	Price float32 `json:"price,omitempty"`

	// Whether or not the SKU is currently visible to users. This will not block users from purchase. Use start_date or stop_date to prevent purchase. Default: true
	Published bool `json:"published,omitempty"`

	// The id of a sale affecting the price, if any
	SaleId int32 `json:"sale_id,omitempty"`

	// The name of a sale affecting the price, if any
	SaleName string `json:"sale_name,omitempty"`

	// The stock keeping unit (SKU), a unique identifier for a given product.  Max 40 characters
	Sku string `json:"sku"`

	// The date the sku becomes visible (if published) and available for purchase, unix timestamp in seconds.  If set to null, sku will become available immediately
	StartDate int64 `json:"start_date,omitempty"`

	// The date the sku becomes hidden and unavailable for purchase, unix timestamp in seconds.  If set to null, sku is always available
	StopDate int64 `json:"stop_date,omitempty"`
}
