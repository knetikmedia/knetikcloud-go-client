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

type ContributionResource struct {

	// A reference to the contributing artist
	Artist SimpleReferenceResourcelong `json:"artist"`

	// A reference to the media being contributed to
	Media SimpleReferenceResourcelong `json:"media"`

	// The nature of the contribution (role of the artist as in 'producer', 'performer', etc)
	Role string `json:"role"`
}
