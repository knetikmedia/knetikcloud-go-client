/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type CouponDefinition struct {

	// A unique identifier for the discount. Can be used to remove the discount, and uniqueness within the cart will be enforced.
	Code string `json:"code"`

	// A description for the discount.
	Description string `json:"description,omitempty"`

	// The type of discount in terms of how it deducts price.
	DiscountType string `json:"discount_type"`

	// Whether this discount is exclusive and cannot be used in conjunction with other discounts/coupons. default=false
	Exclusive bool `json:"exclusive,omitempty"`

	// For coupon_cart, a minimum total price that the cart must meet to be valid.
	MaxDiscount float64 `json:"max_discount,omitempty"`

	// The maximum number of items to count this discount for (not for cart_coupon).
	MaxQuantity int32 `json:"max_quantity,omitempty"`

	// For coupon_cart, a minimum total price that the cart must meet to be valid.
	MinCartTotal float64 `json:"min_cart_total,omitempty"`

	// A name for the discount.
	Name string `json:"name"`

	// Whether this coupon is exclusive to itself or not (true means cannot add two of this same coupon to the same cart).  Default = false
	SelfExclusive bool `json:"self_exclusive,omitempty"`

	// The id of the item this discount applies to, which must be present in the cart. Applies if coupon_type_hint is coupon_single_item or coupon_voucher.
	TargetItemId int32 `json:"target_item_id,omitempty"`

	// The type of discount in terms of what it applies to. coupon_cart applies to the cart as a whole, other types apply to specific items based on different criteria.
	Type_ string `json:"type"`

	// Which tags this applies for (item must have at least one of them), if coupon_type is coupon_tag.
	ValidForTags []string `json:"valid_for_tags,omitempty"`

	// The amount of the discount. If discount_type is value then this is the raw currency amount to remove. If discount_type is percentage then this will be multiplied by the cart total or item price to get the discount amount (0.5 is half price).
	Value float64 `json:"value"`

	// Which vendor this applies for, if coupon_type is coupon_vendor.
	VendorId int32 `json:"vendor_id,omitempty"`
}