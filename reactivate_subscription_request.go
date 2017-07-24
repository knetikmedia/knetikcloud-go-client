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

type ReactivateSubscriptionRequest struct {

	// The inventory to reactivate. Only required if using the deprecated subscriptions service
	InventoryId int32 `json:"inventory_id,omitempty"`

	// Whether to add the additional reactivation fee in addition to the recurring fee
	ReactivationFee bool `json:"reactivation_fee,omitempty"`
}
