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

type UserRelationshipResource struct {

	// The child in the relationship
	Child *SimpleUserResource `json:"child"`

	// Context about the relationship or its type
	Context string `json:"context,omitempty"`

	// A generated unique id. Read-Only
	Id int64 `json:"id,omitempty"`

	// The parent in the relationship
	Parent *SimpleUserResource `json:"parent"`
}
