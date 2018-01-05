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

type BreRule struct {

	// A list of actions to execute, and the mapping for their parameters, when the rule runs. Minimum 1
	Actions []ActionContextobject `json:"actions"`

	// A condition expression that must be met in a given event for the rule to run. Null to always run.
	Condition *PredicateResource `json:"condition,omitempty"`

	// The condition as a readable string. Filled in by the system from the condition
	ConditionText string `json:"condition_text,omitempty"`

	// The human readable description of the rule
	Description string `json:"description,omitempty"`

	// Whether the rule is enabled to run (in conjunction with dates). Default true
	Enabled bool `json:"enabled,omitempty"`

	// The date the rule ceases to take effect, or null if never. Unix timestamp in seconds
	EndDate int64 `json:"end_date,omitempty"`

	// How many times the rule has been evaluated (it's conditions checked, whether it then runs or not)
	EvaluationCount int64 `json:"evaluation_count,omitempty"`

	// The event name of the trigger this rule runs for. Affects which parameters are available
	EventName string `json:"event_name"`

	// The id of the rule for later references. If left null a random guid will be generated. Must be unique. Cannot be changed
	Id string `json:"id,omitempty"`

	// The human readable name of the rule
	Name string `json:"name"`

	// How many times the rule has run
	RunCount int64 `json:"run_count,omitempty"`

	// Used to sort rules to control the order they run in. Larger numbered sort values run first.  Default 500
	Sort int32 `json:"sort,omitempty"`

	// The date the rule begins to take effect, or null if always. Unix timestamp in seconds
	StartDate int64 `json:"start_date,omitempty"`

	// Whether the rule is a default part of the system. System rules cannot be edited or deleted, but may be disabled
	SystemRule bool `json:"system_rule,omitempty"`
}
