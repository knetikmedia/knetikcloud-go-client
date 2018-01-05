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

import (
	"net/http"
)

const ContextOAuth2 		int = 1
const ContextBasicAuth 		int = 2
const ContextAccessToken 	int = 3
const ContextAPIKey 		int = 4 

type BasicAuth struct {
	UserName      string            `json:"userName,omitempty"`
	Password      string            `json:"password,omitempty"`	
}

type APIKey struct {
	Key 	string
	Prefix	string
}

type Configuration struct {
	BasePath      string            	`json:"basePath,omitempty"`
	Host          string            	`json:"host,omitempty"`
	Scheme        string            	`json:"scheme,omitempty"`
	DefaultHeader map[string]string 	`json:"defaultHeader,omitempty"`
	UserAgent     string            	`json:"userAgent,omitempty"`
	HTTPClient 	  *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://devsandbox.knetikcloud.com",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}