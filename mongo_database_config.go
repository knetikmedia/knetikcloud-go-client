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

type MongoDatabaseConfig struct {

	DbName string `json:"db_name,omitempty"`

	Options string `json:"options,omitempty"`

	Password string `json:"password,omitempty"`

	Servers string `json:"servers,omitempty"`

	Username string `json:"username,omitempty"`
}
