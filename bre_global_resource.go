/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type BreGlobalResource struct {

	// A human readable description for display in admin pages
	Description string `json:"description,omitempty"`

	// The id of the global definition. Default is a random guid. Cannot be updated
	Id string `json:"id,omitempty"`

	// The key for the global. Must be unique when combined with scope names. Usually a single descriptive word like 'purchases' or 'logins'
	Key string `json:"key"`

	// A human readable name for display in admin pages
	Name string `json:"name,omitempty"`

	// A list of scoping parameters. Allows the global to have a different value in different context such as a count of purchases for each user (by putting a 'user' scope in this list). When using this global in a rule these scopes will need to be mapped with an expression to provide a value, similar to the parameters in an action
	Scopes []BreGlobalScopeDefinition `json:"scopes,omitempty"`

	// Where this global came from. System globals cannot be removed or updated
	SystemGlobal bool `json:"system_global,omitempty"`

	// The variable type the global stores. See the See Bre Variables enpoint for list
	Type_ string `json:"type"`
}
