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

type Currency struct {

	Active bool `json:"active,omitempty"`

	Code string `json:"code,omitempty"`

	DateCreated int64 `json:"date_created,omitempty"`

	DateUpdated int64 `json:"date_updated,omitempty"`

	Factor float64 `json:"factor,omitempty"`

	Icon string `json:"icon,omitempty"`

	Id int32 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Type_ string `json:"type,omitempty"`

	Virtual bool `json:"virtual,omitempty"`
}
