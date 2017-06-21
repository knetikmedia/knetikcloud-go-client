/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://knetikcloud.com
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
)

type BRERuleEngineActionsApi struct {
	Configuration *Configuration
}

func NewBRERuleEngineActionsApi() *BRERuleEngineActionsApi {
	configuration := NewConfiguration()
	return &BRERuleEngineActionsApi{
		Configuration: configuration,
	}
}

func NewBRERuleEngineActionsApiWithBasePath(basePath string) *BRERuleEngineActionsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &BRERuleEngineActionsApi{
		Configuration: configuration,
	}
}

/**
 * Get a list of available actions
 *
 * @param filterCategory Filter for actions that are within a specific category
 * @param filterName Filter for actions that have names containing the given string
 * @param filterTags Filter for actions that have all of the given tags (comma separated list)
 * @param filterSearch Filter for actions containing the given words somewhere within name, description and tags
 * @return []ActionResource
 */
func (a BRERuleEngineActionsApi) GetBREActions(filterCategory string, filterName string, filterTags string, filterSearch string) ([]ActionResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/bre/actions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(OAuth2)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("filter_category", a.Configuration.APIClient.ParameterToString(filterCategory, ""))
	localVarQueryParams.Add("filter_name", a.Configuration.APIClient.ParameterToString(filterName, ""))
	localVarQueryParams.Add("filter_tags", a.Configuration.APIClient.ParameterToString(filterTags, ""))
	localVarQueryParams.Add("filter_search", a.Configuration.APIClient.ParameterToString(filterSearch, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]ActionResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetBREActions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

