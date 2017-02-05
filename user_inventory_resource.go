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

type UserInventoryResource struct {

	// A map of data for behaviors
	BehaviorData interface{} `json:"behavior_data,omitempty"`

	// The date/time this resource was created in seconds since epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The date/time this resource exires in seconds since epoch. Null for no expiration. For subscriptions, this is the end of the 'grace period' if left unpaid
	Expires int64 `json:"expires,omitempty"`

	// The id of the inventory
	Id int32 `json:"id,omitempty"`

	// The id of the invoice that resulted in this inventory, if any
	InvoiceId int32 `json:"invoice_id,omitempty"`

	// The id of the item
	ItemId int32 `json:"item_id,omitempty"`

	// The name of the item
	ItemName string `json:"item_name,omitempty"`

	// The type hint of the item
	ItemTypeHint string `json:"item_type_hint,omitempty"`

	// The status of the inventory. Pending inventory is not yet ready for use. Inactive inventory has expired or been used up
	Status string `json:"status,omitempty"`

	// The date/time this resource was last updated in seconds since epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The id of the user this inventory belongs to
	User SimpleUserResource `json:"user,omitempty"`
}
