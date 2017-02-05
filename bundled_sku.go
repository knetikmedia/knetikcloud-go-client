/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type BundledSku struct {

	// The amount this item will cost inside the bundle instead of its regular price
	PriceOverride float64 `json:"price_override,omitempty"`

	// The quantity of this item within the bundle
	Quantity int32 `json:"quantity"`

	// The stock keeping unit (SKU) for an item included in the bundle
	Sku string `json:"sku"`
}
