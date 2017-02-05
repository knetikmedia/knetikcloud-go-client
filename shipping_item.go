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

type ShippingItem struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type, or be an extra not from the template
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The behaviors linked to the item, describing various options and interactions. May not be included in item lists
	Behaviors []Behavior `json:"behaviors,omitempty"`

	// A category for filtering items
	Category string `json:"category,omitempty"`

	// A unique list of country iso3 codes that allow the shipping option
	Countries []string `json:"countries,omitempty"`

	// The date the item was created, unix timestamp in seconds
	CreatedDate int64 `json:"created_date,omitempty"`

	// Whether or not the item is currently displayable.  Default = true
	Displayable bool `json:"displayable,omitempty"`

	// A list of country ID to include in the blacklist/whitelist geo policy
	GeoCountryList []string `json:"geo_country_list,omitempty"`

	// Whether to use the geo_country_list as a black list or white list for item geographical availability
	GeoPolicyType string `json:"geo_policy_type,omitempty"`

	// The id of the item
	Id int32 `json:"id,omitempty"`

	// A long description of the item
	LongDescription string `json:"long_description,omitempty"`

	// An abstract max value that the values of item's shipping_tier work against to decide whether an order can be fulfilled
	MaxTierTotal int32 `json:"max_tier_total"`

	// The name of the item
	Name string `json:"name"`

	// Provides the abstract shipping needs if this item is physical and can be shipped.  A value of zero means no shipping needed.  Default = 0
	ShippingTier int32 `json:"shipping_tier,omitempty"`

	// A short description of the item, max 255 chars
	ShortDescription string `json:"short_description,omitempty"`

	// The skus for the item. Each defines a unique configuration for the item to be purchased (Large-Blue, Small-Green, etc). These are what is ultimately selected in the store and added to the cart
	Skus []Sku `json:"skus"`

	// A number to use in sorting items.  Default 500
	Sort int32 `json:"sort,omitempty"`

	// The date the item will leave the store, unix timestamp in seconds.  If set to null, item will never leave the store
	StoreEnd int64 `json:"store_end,omitempty"`

	// The date the item will appear in the store, unix timestamp in seconds.  If set to null, item will appear in store immediately
	StoreStart int64 `json:"store_start,omitempty"`

	// List of tags used for filtering items
	Tags []string `json:"tags,omitempty"`

	// Whether tax should be applied to the shipping price.  Default = false
	Taxable bool `json:"taxable,omitempty"`

	// An item template this item is validated against.  May be null and no validation of additional_properties will be done.  Default = null
	Template string `json:"template,omitempty"`

	// The type of the item
	TypeHint string `json:"type_hint"`

	// The unique key for the item
	UniqueKey string `json:"unique_key,omitempty"`

	// The date the item was last updated, unix timestamp in seconds
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The vendor who provides the item
	VendorId int32 `json:"vendor_id"`
}
