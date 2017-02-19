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

type RevenueReportResource struct {

	CustomerCount int64 `json:"customer_count,omitempty"`

	SaleCount int64 `json:"sale_count,omitempty"`

	SalesAverage float64 `json:"sales_average,omitempty"`

	SalesTotal float64 `json:"sales_total,omitempty"`
}
