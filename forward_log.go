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

type ForwardLog struct {

	// The end date of the forward log entry
	EndDate int64 `json:"end_date,omitempty"`

	ErrorMsg string `json:"error_msg,omitempty"`

	// The http status code the forward log entry
	HttpStatusCode int32 `json:"http_status_code,omitempty"`

	// The id of the forward log entry
	Id string `json:"id,omitempty"`

	// The payload of the forward log entry
	Payload interface{} `json:"payload,omitempty"`

	// The response string of the forward log entry
	Response string `json:"response,omitempty"`

	// The retry count of the forward log entry
	RetryCount int32 `json:"retry_count,omitempty"`

	// The start date of the forward log entry
	StartDate int64 `json:"start_date,omitempty"`

	// The endpoint url of the forward log entry
	Url string `json:"url,omitempty"`
}