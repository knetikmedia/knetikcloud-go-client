/* 
 * Knetik Platform API Documentation Latest
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type DispositionResource struct {

	// The context of that resource. Required when passed to /dispositions rather than context specific endpoint
	Context string `json:"context,omitempty"`

	// The context_id of that resource. Required when passed to /dispositions rather than context specific endpoint
	ContextId string `json:"context_id,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The unique ID for that resource
	Id int64 `json:"id,omitempty"`

	// The name of the disposition, 1-20 characters. (ex: like/dislike/favorite, etc)
	Name string `json:"name,omitempty"`

	// The user
	User SimpleUserResource `json:"user,omitempty"`
}
