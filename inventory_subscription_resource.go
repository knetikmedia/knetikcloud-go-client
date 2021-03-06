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

type InventorySubscriptionResource struct {

	// The date the subscription will be billed
	BillDate int64 `json:"bill_date,omitempty"`

	// A credit of money already applied to a subscription for the next bill, or a debt if negative
	Credit float32 `json:"credit,omitempty"`

	// A record of past and present credit/debt changes
	CreditLog []SubscriptionCreditResource `json:"credit_log,omitempty"`

	// The date the grace period ends
	GraceEnd int64 `json:"grace_end,omitempty"`

	// The id of the inventory
	InventoryId int32 `json:"inventory_id,omitempty"`

	// The inventory status object
	InventoryStatus string `json:"inventory_status,omitempty"`

	// The id of the item
	ItemId int32 `json:"item_id,omitempty"`

	// The payment method object
	PaymentMethod *PaymentMethodResource `json:"payment_method,omitempty"`

	// The recurring price that has been set to override the base price. Null if not overriding
	PriceOverride float32 `json:"price_override,omitempty"`

	// An explanation for the reason the price is being overridden
	PriceOverrideReason string `json:"price_override_reason,omitempty"`

	// The default recurring price
	RecurringPrice float32 `json:"recurring_price,omitempty"`

	// The recurring sku of the subscription
	Sku string `json:"sku,omitempty"`

	// The date the subscription will start
	StartDate int64 `json:"start_date,omitempty"`

	// The status of the subscription
	SubscriptionStatus int32 `json:"subscription_status,omitempty"`

	// The user
	User *SimpleUserResource `json:"user,omitempty"`
}
