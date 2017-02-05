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

type CouponDefinition struct {

	Code string `json:"code,omitempty"`

	Description string `json:"description,omitempty"`

	DiscountType string `json:"discount_type,omitempty"`

	Exclusive bool `json:"exclusive,omitempty"`

	MaxDiscount float64 `json:"max_discount,omitempty"`

	MaxQuantity int32 `json:"max_quantity,omitempty"`

	MinCartTotal float64 `json:"min_cart_total,omitempty"`

	Name string `json:"name,omitempty"`

	SelfExclusive bool `json:"self_exclusive,omitempty"`

	TargetItemId int32 `json:"target_item_id,omitempty"`

	Type_ string `json:"type,omitempty"`

	ValidForTags []string `json:"valid_for_tags,omitempty"`

	Value float64 `json:"value,omitempty"`

	VendorId int32 `json:"vendor_id,omitempty"`
}
