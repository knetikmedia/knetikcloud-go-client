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

import (
	"net/url"
	"strings"
	"encoding/json"
)

type ReportingUsersApi struct {
	Configuration *Configuration
}

func NewReportingUsersApi() *ReportingUsersApi {
	configuration := NewConfiguration()
	return &ReportingUsersApi{
		Configuration: configuration,
	}
}

func NewReportingUsersApiWithBasePath(basePath string) *ReportingUsersApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ReportingUsersApi{
		Configuration: configuration,
	}
}

/**
 * Get user registration info
 * Get user registration counts grouped by time range
 *
 * @param granularity The time duration to aggregate by
 * @param startDate The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time
 * @param endDate The end of the time range to aggregate, unix timestamp in seconds. Default is end of time
 * @return *PageAggregateCountResource
 */
func (a ReportingUsersApi) ExecutiveRevenueItemSalesUsingGET1(granularity string, startDate int64, endDate int64) (*PageAggregateCountResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/users/registrations"

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
	localVarQueryParams.Add("granularity", a.Configuration.APIClient.ParameterToString(granularity, ""))
	localVarQueryParams.Add("start_date", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("end_date", a.Configuration.APIClient.ParameterToString(endDate, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(PageAggregateCountResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ExecutiveRevenueItemSalesUsingGET1", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}
