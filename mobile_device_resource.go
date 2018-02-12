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

type MobileDeviceResource struct {

	// A map of additional properties, keyed on the property name.  Must match the names and types defined in the template if one is specified
	AdditionalProperties map[string]Property `json:"additional_properties,omitempty"`

	// The date the device log was created
	CreatedDate int64 `json:"created_date,omitempty"`

	// The description of the device
	Description string `json:"description,omitempty"`

	// The type of device. Use mobile_device to specifically register mobile devices. This particular type will be used to send and receive notifications
	DeviceType string `json:"device_type,omitempty"`

	// The unique ID for this device. Cannot be changed after creation. Default: random
	Id string `json:"id,omitempty"`

	// The physical location of the device, coordinates or named place (office, living room, etc)
	Location string `json:"location,omitempty"`

	// The MAC (media access control) address of the device
	MacAddress string `json:"mac_address,omitempty"`

	// The make of the device
	Make string `json:"make,omitempty"`

	// The model of the device
	Model string `json:"model,omitempty"`

	// The name of the device
	Name string `json:"name,omitempty"`

	// The OS (operating system) on the device
	Os string `json:"os,omitempty"`

	// The user that owns the device
	Owner *SimpleUserResource `json:"owner,omitempty"`

	// The serial number of the device
	Serial string `json:"serial,omitempty"`

	// Random tags to facilitate search
	Tags []string `json:"tags,omitempty"`

	// Use to describe and validate custom properties (custom schema). May be null and no validation of additional_properties will be done
	Template string `json:"template,omitempty"`

	// The date the device log was updated
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The users currently using the device
	Users []SimpleUserResource `json:"users,omitempty"`

	// The authorization code for push notifications provided by the provider platform (APNS, GCM, etc).
	Authorization string `json:"authorization,omitempty"`

	Imei string `json:"imei,omitempty"`

	// The platform used for push notifications. Only Apple and Android are supported at the moment.
	NotificationPlatform string `json:"notification_platform,omitempty"`

	// The phone number associated with this device if applicable, in international format
	Number string `json:"number,omitempty"`
}
