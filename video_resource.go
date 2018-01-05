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

type VideoResource struct {

	// Whether the video is available, based on various factors
	Active bool `json:"active,omitempty"`

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The original artist of the media
	Author *SimpleReferenceResourcelong `json:"author,omitempty"`

	// The date the media was created as a unix timestamp in seconds
	Authored int64 `json:"authored,omitempty"`

	// Whether the video has been banned or not
	Banned bool `json:"banned,omitempty"`

	// The category of the video
	Category *SimpleReferenceResourcestring `json:"category"`

	// The comments of the video
	Comments []CommentResource `json:"comments,omitempty"`

	// Artists that contributed to the creation. See separate endpoint to add to list
	Contributors []ContributionResource `json:"contributors,omitempty"`

	// The date/time this resource was created in seconds since unix epoch
	CreatedDate int64 `json:"created_date,omitempty"`

	// The country of an embedable version
	Embed string `json:"embed,omitempty"`

	// The file extension of the media file. 1-5 characters
	Extension string `json:"extension"`

	// The height of the video in px
	Height int32 `json:"height"`

	// The unique ID for that resource
	Id int64 `json:"id,omitempty"`

	// The length of the video in seconds
	Length int32 `json:"length"`

	// The country of the media. Typically a url. Cannot be blank
	Location string `json:"location"`

	// The user friendly name of that resource. Defaults to blank string
	LongDescription string `json:"long_description,omitempty"`

	// The mime-type of the media
	MimeType string `json:"mime_type,omitempty"`

	// The user friendly name of that resource
	Name string `json:"name"`

	// The sort order of the video. default: 100
	Priority int32 `json:"priority,omitempty"`

	// The privacy setting. default: private
	Privacy string `json:"privacy,omitempty"`

	// Whether the video has been made public. Default true
	Published bool `json:"published,omitempty"`

	// The user friendly name of that resource. Defaults to blank string
	ShortDescription string `json:"short_description,omitempty"`

	// The size of the media. Minimum 0 if supplied
	Size int64 `json:"size,omitempty"`

	// The tags for the video
	Tags []string `json:"tags,omitempty"`

	// A video template this video is validated against (private). May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`

	// The country of a thumbnail version. Typically a url
	Thumbnail string `json:"thumbnail,omitempty"`

	// The date/time this resource was last updated in seconds since unix epoch
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The user the media was uploaded by. May be null for system uploaded media. May only be set to a user other than the current caller if VIDEOS_ADMIN permission. Null will mean the caller is the uploader unless the caller has VIDEOS_ADMIN permission, in which case it will be set to null
	Uploader *SimpleUserResource `json:"uploader,omitempty"`

	// The view count of the video
	Views int64 `json:"views,omitempty"`

	// The width of the video in px
	Width int32 `json:"width"`
}
