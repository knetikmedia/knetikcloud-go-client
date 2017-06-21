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

type CouponItem struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type, or be an extra not from the template
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The behaviors linked to the item, describing various options and interactions. May not be included in item lists
	Behaviors []Behavior `json:"behaviors,omitempty"`

	// A category for filtering items
	Category string `json:"category,omitempty"`

	// The date the item was created, unix timestamp in seconds
	CreatedDate int64 `json:"created_date,omitempty"`

	// The id of the item
	Id int32 `json:"id,omitempty"`

	// A long description of the item
	LongDescription string `json:"long_description,omitempty"`

	// The name of the item
	Name string `json:"name"`

	// A short description of the item, max 255 chars
	ShortDescription string `json:"short_description,omitempty"`

	// A number to use in sorting items.  Default 500
	Sort int32 `json:"sort,omitempty"`

	// List of tags used for filtering items
	Tags []string `json:"tags,omitempty"`

	// An item template this item is validated against.  May be null and no validation of additional_properties will be done.  Default = null
	Template string `json:"template,omitempty"`

	// The type of the item
	TypeHint string `json:"type_hint"`

	// The unique key for the item
	UniqueKey string `json:"unique_key,omitempty"`

	// The date the item was last updated, unix timestamp in seconds
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// Whether or not the item is currently displayable.  Default = true
	Displayable bool `json:"displayable,omitempty"`

	// A list of country ID to include in the blacklist/whitelist geo policy
	GeoCountryList []string `json:"geo_country_list,omitempty"`

	// Whether to use the geo_country_list as a black list or white list for item geographical availability
	GeoPolicyType string `json:"geo_policy_type,omitempty"`

	// Provides the abstract shipping needs if this item is physical and can be shipped.  A value of zero means no shipping needed.  Default = 0
	ShippingTier int32 `json:"shipping_tier,omitempty"`

	// The skus for the item. Each defines a unique configuration for the item to be purchased (Large-Blue, Small-Green, etc). These are what is ultimately selected in the store and added to the cart
	Skus []Sku `json:"skus"`

	// The date the item will leave the store, unix timestamp in seconds.  If set to null, item will never leave the store
	StoreEnd int64 `json:"store_end,omitempty"`

	// The date the item will appear in the store, unix timestamp in seconds.  If set to null, item will appear in store immediately
	StoreStart int64 `json:"store_start,omitempty"`

	// The vendor who provides the item
	VendorId int32 `json:"vendor_id"`

	// The type of coupon
	CouponTypeHint string `json:"coupon_type_hint"`

	// The amount this coupon is maxed out at.  Applies if coupon_type_hint is coupon_cart
	DiscountMax float64 `json:"discount_max,omitempty"`

	// The minimium amount needed in the cart for the coupon to apply.  Applies if coupon_type_hint is coupon_cart
	DiscountMinCartValue float64 `json:"discount_min_cart_value,omitempty"`

	// The type of discount in terms of how it deducts price. Value based discount not available for coupon_cart type coupons
	DiscountType string `json:"discount_type"`

	// The amount the coupon will discount the item. If discount_type is 'value' this will be a flat amount of currency. If discount type is 'percentage' this will be a fraction (0.2 for 20% off) multiplied by the price of the matching item or items.
	DiscountValue float64 `json:"discount_value"`

	// Whether this coupon is exclusive or not (true means cannot be in same cart as another).  Default = false
	Exclusive bool `json:"exclusive,omitempty"`

	// The id of the item the coupon is applied to.  Applies if coupon_type_hint is coupon_single_item or coupon_voucher
	ItemId int32 `json:"item_id,omitempty"`

	// The maximum quantity of items the coupon can apply to, null if no limit and minimum 1 otherwise.  Applies if coupon_type_hint is coupon_single_item or coupon_voucher
	MaxQuantity int32 `json:"max_quantity,omitempty"`

	// Whether this coupon is exclusive to itself or not (true means cannot add two of this same coupon to the same cart).  Default = false
	SelfExclusive bool `json:"self_exclusive,omitempty"`

	// A list of tags for a coupon.  The coupon can only apply to an item that has at least one of these tags.  Applies if coupon_type_hint is coupon_tag
	ValidForTags []string `json:"valid_for_tags,omitempty"`
}
