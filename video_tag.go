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

type VideoTag struct {

	Id int64 `json:"id,omitempty"`

	Tag string `json:"tag,omitempty"`

	Video Video `json:"video,omitempty"`
}
