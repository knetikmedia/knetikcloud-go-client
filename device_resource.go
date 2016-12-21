/* 
 * Knetik Platform API Documentation Latest
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://demo.sandbox.knetikcloud.com
 *
 * OpenAPI spec version: Latest
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type DeviceResource struct {

	// The authorization code for the device
	Authorization string `json:"authorization,omitempty"`

	// The current condition of the device (New, Defective, Reconditioned)
	Condition string `json:"condition,omitempty"`

	// The date the device log was created
	CreatedDate int64 `json:"created_date,omitempty"`

	// The key/value pairs for extended data
	Data map[string]string `json:"data,omitempty"`

	// The description of the device
	Description string `json:"description,omitempty"`

	// The type of the device
	DeviceType string `json:"device_type,omitempty"`

	// The unique ID for this device. Cannot be changed once created
	Id int32 `json:"id,omitempty"`

	// The location of the device
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

	// The serial number of the device
	Serial string `json:"serial,omitempty"`

	// The current status the device (Active, Pending Active, Inactive, Repair
	Status string `json:"status,omitempty"`

	// The date the device log was updated
	UpdatedDate int64 `json:"updated_date,omitempty"`

	// The user that owns the device
	User SimpleUserResource `json:"user,omitempty"`
}
