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

type WalletTotalResponse struct {

	// The currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// The sum of all wallets in the system for this currency
	Total float64 `json:"total,omitempty"`
}