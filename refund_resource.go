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

type RefundResource struct {

	// The amount refunded
	Amount float64 `json:"amount,omitempty"`

	// The id of the refund transaction
	RefundTransactionId int32 `json:"refund_transaction_id,omitempty"`

	// The id of the original transaction
	TransactionId int32 `json:"transaction_id,omitempty"`
}
