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

type RepresentsAnActivityThatCanBeParameterizedAndTrackedThroughMetricsScoresEtc struct {

	// A map of additional properties keyed on the property name. Used to further describe an activity. While settings will vary from one activity occurrence (a game) to another, additional properties are shared by all the occurrences of this activity. Ex: Activity Logo, Disclaimer, Greeting, etc. Validated against template if one exists for activities
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The list of items that can be used for entitlement (wager amounts/etc)
	Entitlements []ActivityEntitlementResource `json:"entitlements,omitempty"`

	// The unique ID for that resource
	Id int64 `json:"id,omitempty"`

	// Details about how to launch the activity
	Launch string `json:"launch,omitempty"`

	// The user friendly name of that resource. Defaults to blank string
	LongDescription string `json:"long_description,omitempty"`

	// The user friendly name of that resource
	Name string `json:"name"`

	// The rewards to give at the end of each occurence of the activity. When creating/updating only id is used. Reward set must be pre-existing
	RewardSet RewardSetResource `json:"reward_set,omitempty"`

	// Define what parameters are required/available to start and run an activity. For example: Difficulty, Number of Questions, Character name, Avatar, Duration, etc. Not populated when getting listing
	Settings []TheDefinitionOfAnActivityParametersExDifficultyLevel `json:"settings,omitempty"`

	// The user friendly name of that resource. Defaults to blank string
	ShortDescription string `json:"short_description,omitempty"`

	// Whether this activity is a template for other activities. Default: false
	Template bool `json:"template,omitempty"`

	// An activity template this activity is validated against (private). May be null and no validation of additional_properties will be done
	TemplateId string `json:"template_id,omitempty"`

	// The type of the activity
	Type_ string `json:"type"`

	// The unique key (for static reference in code) of the activity
	UniqueKey string `json:"unique_key,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`
}
