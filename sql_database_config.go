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

type SqlDatabaseConfig struct {

	ConnectionPoolSize int32 `json:"connection_pool_size,omitempty"`

	DbName string `json:"db_name,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	Password string `json:"password,omitempty"`

	Port int32 `json:"port,omitempty"`

	Username string `json:"username,omitempty"`
}
