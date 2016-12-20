/* 
 * Knetik Platform API Documentation Latest
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type CartItemRequest struct {

	// The affiliate key of the item
	AffiliateKey string `json:"affiliate_key,omitempty"`

	// The catalog SKU of the item
	CatalogSku string `json:"catalog_sku,omitempty"`

	// The quantity of the item
	Quantity int32 `json:"quantity,omitempty"`
}