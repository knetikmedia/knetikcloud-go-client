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

type UsageInfo struct {

	// The number of requests within the range
	Count int64 `json:"count,omitempty"`

	// The date at the start of the range (see granularity)
	Date int64 `json:"date,omitempty"`

	// The http method
	Method string `json:"method,omitempty"`

	// The url path
	Url string `json:"url,omitempty"`
}