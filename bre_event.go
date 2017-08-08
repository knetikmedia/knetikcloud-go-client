/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://knetikcloud.com.
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type BreEvent struct {

	// The event name of the trigger to be fired
	EventName string `json:"event_name"`

	// The parameters to the event. A Map (assosiative array) with a key for each trigger parameter name and a corrosponding value.
	Params *interface{} `json:"params"`
}
