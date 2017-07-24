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

type SubscriptionPlanResource struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this subscription
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The length of the billing cycle in number of billing cycle unit
	BillingCycleLength int32 `json:"billing_cycle_length"`

	// The time period unit to apply to the length of billing cycles
	BillingCycleUnit string `json:"billing_cycle_unit"`

	// Whether this plan will be renewed on the consolidated billing cycle
	Consolidated bool `json:"consolidated"`

	// The ISO3 currency code to use for the fees
	CurrencyCode string `json:"currency_code"`

	// Used to schedule plan availability end date
	EndDate int64 `json:"end_date,omitempty"`

	// Optional override for the length of the first billing cycle before the first recurring billing
	FirstBillingCycleLength int32 `json:"first_billing_cycle_length,omitempty"`

	// The time period unit to apply to the length of the first billing cycle. Required when first_billing_cycle_length is specified
	FirstBillingCycleUnit string `json:"first_billing_cycle_unit,omitempty"`

	// The number of late payment days before a subscription is canceled
	GracePeriod int32 `json:"grace_period"`

	// The id of the plan used to generate the SKUs
	Id string `json:"id,omitempty"`

	// The fee charged when the subscription is purchased
	InitialFee float64 `json:"initial_fee"`

	// The SKU to be used when purchasing the subscription through the cart
	InitialSku string `json:"initial_sku,omitempty"`

	// The fee to add to the bill when an invoice has gone unpaid passed the grace period
	LatePaymentFee float64 `json:"late_payment_fee"`

	// The SKU that will show on the invoice when the subscription is delinquent
	LatePaymentSku string `json:"late_payment_sku,omitempty"`

	// Whether this plan is locked because it has been purchased by at least one user.  When locked, a number of properties can no longer be changed
	Locked bool `json:"locked,omitempty"`

	// The number of charge attempts before the subscription becomes delinquent
	MaxBillAttempts int32 `json:"max_bill_attempts"`

	// Maximum number of renewals. If a migration plan is provided, the subscription will automatically switch to it when this limit is reached
	MaxCycles int32 `json:"max_cycles,omitempty"`

	// Automatically migrate to the specified plan when the subscription is first renewed
	MigrateToPlan string `json:"migrate_to_plan,omitempty"`

	// The minimum number of renewals to charge for
	MinCycles int32 `json:"min_cycles,omitempty"`

	// The name of the plan used to generate the SKUs
	Name string `json:"name"`

	// Whether this plan is currently available
	Published bool `json:"published"`

	// The fee to charge when a suspended subscription is to be re-activated
	ReactivationFee float64 `json:"reactivation_fee"`

	// The SKU that will show on the invoice when the subscription is re-activated after a suspension
	ReactivationSku string `json:"reactivation_sku,omitempty"`

	// The recurring fee to charge for each renewal
	RecurringFee float64 `json:"recurring_fee"`

	// The SKU that will show on the invoice when the subscription is activated
	RecurringSku string `json:"recurring_sku,omitempty"`

	// Used to schedule plan availability start date
	StartDate int64 `json:"start_date,omitempty"`
}
