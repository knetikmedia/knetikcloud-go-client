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

type PageResourceAggregateInvoiceReportResource struct {

	Content []AggregateInvoiceReportResource `json:"content,omitempty"`

	First bool `json:"first,omitempty"`

	Last bool `json:"last,omitempty"`

	Number int32 `json:"number,omitempty"`

	NumberOfElements int32 `json:"number_of_elements,omitempty"`

	Size int32 `json:"size,omitempty"`

	Sort []Order `json:"sort,omitempty"`

	TotalElements int64 `json:"total_elements,omitempty"`

	TotalPages int32 `json:"total_pages,omitempty"`
}
