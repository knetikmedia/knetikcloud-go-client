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

type ItemBehaviorDefinitionResource struct {

	// The default version of the behavior
	Behavior *Behavior `json:"behavior"`

	// Whether the behavior's values can be modified
	Modifiable bool `json:"modifiable"`

	// Whether the behavior can be removed
	Required bool `json:"required"`
}
