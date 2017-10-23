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

type CartLineItem struct {

	CurrencyCode string `json:"currency_code,omitempty"`

	Description string `json:"description,omitempty"`

	Discount *Discount `json:"discount,omitempty"`

	LineTotal float32 `json:"line_total,omitempty"`

	Name string `json:"name,omitempty"`

	OriginalLineTotal float32 `json:"original_line_total,omitempty"`

	OriginalUnitPrice float32 `json:"original_unit_price,omitempty"`

	Qty int32 `json:"qty,omitempty"`

	SaleName string `json:"sale_name,omitempty"`

	Sku string `json:"sku,omitempty"`

	SkuDescription string `json:"sku_description,omitempty"`

	StoreItemId int32 `json:"store_item_id,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ThumbUrl string `json:"thumb_url,omitempty"`

	UniqueKey string `json:"unique_key,omitempty"`

	UnitPrice float32 `json:"unit_price,omitempty"`

	VendorId int32 `json:"vendor_id,omitempty"`

	VendorName string `json:"vendor_name,omitempty"`
}
