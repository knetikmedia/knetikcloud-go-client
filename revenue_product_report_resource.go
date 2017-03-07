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

type RevenueProductReportResource struct {

	ItemId int32 `json:"item_id,omitempty"`

	ItemName string `json:"item_name,omitempty"`

	Revenue float64 `json:"revenue,omitempty"`

	Volume int64 `json:"volume,omitempty"`
}