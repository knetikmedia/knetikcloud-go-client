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

type InvoiceItemResource struct {

	AffiliateId int32 `json:"affiliate_id,omitempty"`

	BundleSku string `json:"bundle_sku,omitempty"`

	CurrentFulfillmentStatus string `json:"current_fulfillment_status,omitempty"`

	Id int32 `json:"id,omitempty"`

	InvoiceId int32 `json:"invoice_id,omitempty"`

	ItemId int32 `json:"item_id,omitempty"`

	ItemName string `json:"item_name,omitempty"`

	OriginalTotalPrice float64 `json:"original_total_price,omitempty"`

	OriginalUnitPrice float64 `json:"original_unit_price,omitempty"`

	Qty int32 `json:"qty,omitempty"`

	SaleName string `json:"sale_name,omitempty"`

	Sku string `json:"sku,omitempty"`

	SkuDescription string `json:"sku_description,omitempty"`

	SystemPrice float64 `json:"system_price,omitempty"`

	TotalPrice float64 `json:"total_price,omitempty"`

	TypeHint string `json:"type_hint,omitempty"`

	UnitPrice float64 `json:"unit_price,omitempty"`
}
