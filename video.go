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

type Video struct {

	Active bool `json:"active,omitempty"`

	Author Artist `json:"author,omitempty"`

	Authored int64 `json:"authored,omitempty"`

	Banned bool `json:"banned,omitempty"`

	CategoryId string `json:"category_id,omitempty"`

	CategoryName string `json:"category_name,omitempty"`

	Contributors CollectionVideoContribution `json:"contributors,omitempty"`

	Created int64 `json:"created,omitempty"`

	Embed string `json:"embed,omitempty"`

	Extension string `json:"extension,omitempty"`

	Height int32 `json:"height,omitempty"`

	Id int64 `json:"id,omitempty"`

	Length int32 `json:"length,omitempty"`

	Location string `json:"location,omitempty"`

	LongDescription string `json:"long_description,omitempty"`

	MimeType string `json:"mime_type,omitempty"`

	Name string `json:"name,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	Privacy string `json:"privacy,omitempty"`

	Published bool `json:"published,omitempty"`

	ShortDescription string `json:"short_description,omitempty"`

	Size int64 `json:"size,omitempty"`

	Tags []VideoTag `json:"tags,omitempty"`

	Thumbnail string `json:"thumbnail,omitempty"`

	Updated int64 `json:"updated,omitempty"`

	Uploader User `json:"uploader,omitempty"`

	Views int64 `json:"views,omitempty"`

	Whitelist []User `json:"whitelist,omitempty"`

	Width int32 `json:"width,omitempty"`
}
