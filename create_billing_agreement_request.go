/* 
 * Knetik Platform API Documentation Latest
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type CreateBillingAgreementRequest struct {

	// The endpoint URL to which PayPal should forward the user if they cancel (do not accept) the agreement
	CancelUrl string `json:"cancel_url,omitempty"`

	// The endpoint URL to which PayPal should forward the user after they accept the agreement. This endpoint will receive information needed for the next step
	ReturnUrl string `json:"return_url,omitempty"`

	// The ID of the user. Defaults to the logged in user
	UserId int32 `json:"user_id,omitempty"`
}
