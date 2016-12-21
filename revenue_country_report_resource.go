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

type RevenueCountryReportResource struct {

	Country string `json:"country,omitempty"`

	Revenue float64 `json:"revenue,omitempty"`

	Volume int64 `json:"volume,omitempty"`
}
