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

type ImportJobResource struct {

	// The id of the category to assign all questions in the import to
	CategoryId string `json:"category_id"`

	// The date the job was created in seconds since unix epoc
	CreatedDate int64 `json:"created_date,omitempty"`

	// The id of the job
	Id int64 `json:"id,omitempty"`

	// A name for this import for later reference
	Name string `json:"name"`

	// Error information from validation
	Output []ImportJobOutputResource `json:"output,omitempty"`

	// The number of questions form the CSV file. Filled in after validation
	RecordCount int64 `json:"record_count,omitempty"`

	// The status of the job
	Status string `json:"status,omitempty"`

	// The date the job was last updated in seconds since unix epoc
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The url of a CSV file to pull trivia questions from. Cannot be changed after initial POST
	Url string `json:"url"`

	// The vendor who supplied this set of questions
	Vendor string `json:"vendor"`
}
