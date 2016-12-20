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

type ReportingUsageApi struct {
	Configuration *Configuration
}

func NewReportingUsageApi() *ReportingUsageApi {
	configuration := NewConfiguration()
	return &ReportingUsageApi{
		Configuration: configuration,
	}
}

func NewReportingUsageApiWithBasePath(basePath string) *ReportingUsageApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ReportingUsageApi{
		Configuration: configuration,
	}
}

/**
 * Returns aggregated endpoint usage information by the day
 *
 * @param startDate The beginning of the range being requested, unix timestamp in seconds
 * @param endDate The ending of the range being requested, unix timestamp in seconds
 * @param combineEndpoints Whether to combine counts from different endpoint. Removes the url and method from the result object
 * @param size The number of objects returned per page
 * @param page The number of the page returned, starting with 1
 * @return *PageResourceUsageInfo
 */
func (a ReportingUsageApi) GetUsageByDayUsingGET(startDate int64, endDate int64, combineEndpoints bool, size int32, page int32) (*PageResourceUsageInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/usage/day"

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
	localVarQueryParams.Add("start_date", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("end_date", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("combine_endpoints", a.Configuration.APIClient.ParameterToString(combineEndpoints, ""))
	localVarQueryParams.Add("size", a.Configuration.APIClient.ParameterToString(size, ""))
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))

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
	var successPayload = new(PageResourceUsageInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetUsageByDayUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

/**
 * Returns aggregated endpoint usage information by hour
 *
 * @param startDate The beginning of the range being requested, unix timestamp in seconds
 * @param endDate The ending of the range being requested, unix timestamp in seconds
 * @param combineEndpoints Whether to combine counts from different endpoint. Removes the url and method from the result object
 * @param size The number of objects returned per page
 * @param page The number of the page returned, starting with 1
 * @return *PageResourceUsageInfo
 */
func (a ReportingUsageApi) GetUsageByHourUsingGET(startDate int64, endDate int64, combineEndpoints bool, size int32, page int32) (*PageResourceUsageInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/usage/hour"

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
	localVarQueryParams.Add("start_date", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("end_date", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("combine_endpoints", a.Configuration.APIClient.ParameterToString(combineEndpoints, ""))
	localVarQueryParams.Add("size", a.Configuration.APIClient.ParameterToString(size, ""))
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))

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
	var successPayload = new(PageResourceUsageInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetUsageByHourUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

/**
 * Returns aggregated endpoint usage information by minute
 *
 * @param startDate The beginning of the range being requested, unix timestamp in seconds
 * @param endDate The ending of the range being requested, unix timestamp in seconds
 * @param combineEndpoints Whether to combine counts from different endpoint. Removes the url and method from the result object
 * @param size The number of objects returned per page
 * @param page The number of the page returned, starting with 1
 * @return *PageResourceUsageInfo
 */
func (a ReportingUsageApi) GetUsageByMinuteUsingGET(startDate int64, endDate int64, combineEndpoints bool, size int32, page int32) (*PageResourceUsageInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/usage/minute"

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
	localVarQueryParams.Add("start_date", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("end_date", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("combine_endpoints", a.Configuration.APIClient.ParameterToString(combineEndpoints, ""))
	localVarQueryParams.Add("size", a.Configuration.APIClient.ParameterToString(size, ""))
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))

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
	var successPayload = new(PageResourceUsageInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetUsageByMinuteUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

/**
 * Returns aggregated endpoint usage information by month
 *
 * @param startDate The beginning of the range being requested, unix timestamp in seconds
 * @param endDate The ending of the range being requested, unix timestamp in seconds
 * @param combineEndpoints Whether to combine counts from different endpoint. Removes the url and method from the result object
 * @param size The number of objects returned per page
 * @param page The number of the page returned, starting with 1
 * @return *PageResourceUsageInfo
 */
func (a ReportingUsageApi) GetUsageByMonthUsingGET(startDate int64, endDate int64, combineEndpoints bool, size int32, page int32) (*PageResourceUsageInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/usage/month"

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
	localVarQueryParams.Add("start_date", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("end_date", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("combine_endpoints", a.Configuration.APIClient.ParameterToString(combineEndpoints, ""))
	localVarQueryParams.Add("size", a.Configuration.APIClient.ParameterToString(size, ""))
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))

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
	var successPayload = new(PageResourceUsageInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetUsageByMonthUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

/**
 * Returns aggregated endpoint usage information by year
 *
 * @param startDate The beginning of the range being requested, unix timestamp in seconds
 * @param endDate The ending of the range being requested, unix timestamp in seconds
 * @param combineEndpoints Whether to combine counts from different endpoint. Removes the url and method from the result object
 * @param size The number of objects returned per page
 * @param page The number of the page returned, starting with 1
 * @return *PageResourceUsageInfo
 */
func (a ReportingUsageApi) GetUsageByYearUsingGET(startDate int64, endDate int64, combineEndpoints bool, size int32, page int32) (*PageResourceUsageInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/usage/year"

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
	localVarQueryParams.Add("start_date", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("end_date", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("combine_endpoints", a.Configuration.APIClient.ParameterToString(combineEndpoints, ""))
	localVarQueryParams.Add("size", a.Configuration.APIClient.ParameterToString(size, ""))
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))

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
	var successPayload = new(PageResourceUsageInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetUsageByYearUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

