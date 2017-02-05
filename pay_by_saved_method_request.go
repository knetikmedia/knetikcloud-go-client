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

type PayBySavedMethodRequest struct {

	// The id of the payment method to use. Must belong to the caller, be public or have PAYMENTS_ADMIN permission
	PaymentMethod int32 `json:"payment_method"`
}
