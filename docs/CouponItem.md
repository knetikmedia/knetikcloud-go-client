# CouponItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type, or be an extra not from the template | [optional] [default to null]
**Behaviors** | [**[]Behavior**](Behavior.md) | The behaviors linked to the item, describing various options and interactions. May not be included in item lists | [optional] [default to null]
**Category** | **string** | A category for filtering items | [optional] [default to null]
**CouponTypeHint** | **string** | The type of coupon | [default to null]
**CreatedDate** | **int64** | The date the item was created, unix timestamp in seconds | [optional] [default to null]
**DiscountMax** | **float64** | The amount this coupon is maxed out at.  Applies if coupon_type_hint is coupon_cart | [optional] [default to null]
**DiscountMinCartValue** | **float64** | The minimium amount needed in the cart for the coupon to apply.  Applies if coupon_type_hint is coupon_cart | [optional] [default to null]
**DiscountType** | **string** | The type of discount in terms of how it deducts price. Value based discount not available for coupon_cart type coupons | [default to null]
**DiscountValue** | **float64** | The amount the coupon will discount the item. If discount_type is &#39;value&#39; this will be a flat amount of currency. If discount type is &#39;percentage&#39; this will be a fraction (0.2 for 20% off) multiplied by the price of the matching item or items. | [default to null]
**Displayable** | **bool** | Whether or not the item is currently displayable.  Default &#x3D; true | [optional] [default to null]
**Exclusive** | **bool** | Whether this coupon is exclusive or not (true means cannot be in same cart as another).  Default &#x3D; false | [optional] [default to null]
**GeoCountryList** | **[]string** | A list of country ID to include in the blacklist/whitelist geo policy | [optional] [default to null]
**GeoPolicyType** | **string** | Whether to use the geo_country_list as a black list or white list for item geographical availability | [optional] [default to null]
**Id** | **int32** | The id of the item | [optional] [default to null]
**ItemId** | **int32** | The id of the item the coupon is applied to.  Applies if coupon_type_hint is coupon_single_item or coupon_voucher | [optional] [default to null]
**LongDescription** | **string** | A long description of the item | [optional] [default to null]
**MaxQuantity** | **int32** | The maximum quantity of items the coupon can apply to, null if no limit and minimum 1 otherwise.  Applies if coupon_type_hint is coupon_single_item or coupon_voucher | [optional] [default to null]
**Name** | **string** | The name of the item | [default to null]
**SelfExclusive** | **bool** | Whether this coupon is exclusive to itself or not (true means cannot add two of this same coupon to the same cart).  Default &#x3D; false | [optional] [default to null]
**ShippingTier** | **int32** | Provides the abstract shipping needs if this item is physical and can be shipped.  A value of zero means no shipping needed.  Default &#x3D; 0 | [optional] [default to null]
**ShortDescription** | **string** | A short description of the item, max 255 chars | [optional] [default to null]
**Skus** | [**[]Sku**](Sku.md) | The skus for the item. Each defines a unique configuration for the item to be purchased (Large-Blue, Small-Green, etc). These are what is ultimately selected in the store and added to the cart | [default to null]
**Sort** | **int32** | A number to use in sorting items.  Default 500 | [optional] [default to null]
**StoreEnd** | **int64** | The date the item will leave the store, unix timestamp in seconds.  If set to null, item will never leave the store | [optional] [default to null]
**StoreStart** | **int64** | The date the item will appear in the store, unix timestamp in seconds.  If set to null, item will appear in store immediately | [optional] [default to null]
**Tags** | **[]string** | List of tags used for filtering items | [optional] [default to null]
**Template** | **string** | An item template this item is validated against.  May be null and no validation of additional_properties will be done.  Default &#x3D; null | [optional] [default to null]
**TypeHint** | **string** | The type of the item | [default to null]
**UniqueKey** | **string** | The unique key for the item | [optional] [default to null]
**UpdatedDate** | **int64** | The date the item was last updated, unix timestamp in seconds | [optional] [default to null]
**ValidForTags** | **[]string** | A list of tags for a coupon.  The coupon can only apply to an item that has at least one of these tags.  Applies if coupon_type_hint is coupon_tag | [optional] [default to null]
**VendorId** | **int32** | The vendor who provides the item | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


