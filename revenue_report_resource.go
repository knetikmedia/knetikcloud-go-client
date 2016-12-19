/* 
 * Knetik Platform API Documentation Latest
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: Latest
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
