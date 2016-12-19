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

type StoreListRequest struct {

	// Whether the location is ignored
	IgnoreLocation bool `json:"ignore_location,omitempty"`

	// Whether the item is in stock
	InStockOnly bool `json:"in_stock_only,omitempty"`

	// The amount of items returned
	Limit int32 `json:"limit,omitempty"`

	// The page of the request
	Page int32 `json:"page,omitempty"`

	// Whether the catalog is used
	UseCatalog bool `json:"use_catalog,omitempty"`
}
