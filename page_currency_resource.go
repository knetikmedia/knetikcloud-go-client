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

type PageCurrencyResource struct {

	Content []CurrencyResource `json:"content,omitempty"`

	First bool `json:"first,omitempty"`

	Last bool `json:"last,omitempty"`

	Number int32 `json:"number,omitempty"`

	NumberOfElements int32 `json:"number_of_elements,omitempty"`

	Size int32 `json:"size,omitempty"`

	Sort Sort `json:"sort,omitempty"`

	TotalElements int64 `json:"total_elements,omitempty"`

	TotalPages int32 `json:"total_pages,omitempty"`
}
