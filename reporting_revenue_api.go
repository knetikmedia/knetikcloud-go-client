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
	"fmt"
)

type ReportingRevenueApi struct {
	Configuration *Configuration
}

func NewReportingRevenueApi() *ReportingRevenueApi {
	configuration := NewConfiguration()
	return &ReportingRevenueApi{
		Configuration: configuration,
	}
}

func NewReportingRevenueApiWithBasePath(basePath string) *ReportingRevenueApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ReportingRevenueApi{
		Configuration: configuration,
	}
}

/**
 * Get revenue info by country
 * Get basic info about revenue from sales of all types, summed up within a time range and split out by country. Sorted for largest revenue at the top
 *
 * @param currencyCode The code for a currency to get sales data for
 * @param startDate The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time
 * @param endDate The end of the time range to aggregate, unix timestamp in seconds. Default is end of time
 * @param size The number of objects returned per page
 * @param page The number of the page returned, starting with 1
 * @return *PageResourceRevenueCountryReportResource
 */
func (a ReportingRevenueApi) ExecutiveRevenueCountrySalesUsingGET(currencyCode string, startDate int64, endDate int64, size int32, page int32) (*PageResourceRevenueCountryReportResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/revenue/countries/{currency_code}"
	localVarPath = strings.Replace(localVarPath, "{"+"currency_code"+"}", fmt.Sprintf("%v", currencyCode), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(knetik_oauth)' required
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
	var successPayload = new(PageResourceRevenueCountryReportResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ExecutiveRevenueCountrySalesUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get item revenue info
 * Get basic info about revenue from sales of items and bundles (not subscriptions, shipping, etc), summed up within a time range
 *
 * @param currencyCode The code for a currency to get sales data for
 * @param startDate The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time
 * @param endDate The end of the time range to aggregate, unix timestamp in seconds. Default is end of time
 * @return *RevenueReportResource
 */
func (a ReportingRevenueApi) ExecutiveRevenueItemSalesUsingGET(currencyCode string, startDate int64, endDate int64) (*RevenueReportResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/revenue/item-sales/{currency_code}"
	localVarPath = strings.Replace(localVarPath, "{"+"currency_code"+"}", fmt.Sprintf("%v", currencyCode), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(knetik_oauth)' required
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
	var successPayload = new(RevenueReportResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ExecutiveRevenueItemSalesUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get revenue info by item
 * Get basic info about revenue from sales of all types, summed up within a time range and split out by specific item. Sorted for largest revenue at the top
 *
 * @param currencyCode The code for a currency to get sales data for
 * @param startDate The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time
 * @param endDate The end of the time range to aggregate, unix timestamp in seconds. Default is end of time
 * @param size The number of objects returned per page
 * @param page The number of the page returned, starting with 1
 * @return *PageResourceRevenueProductReportResource
 */
func (a ReportingRevenueApi) ExecutiveRevenueProductSalesUsingGET(currencyCode string, startDate int64, endDate int64, size int32, page int32) (*PageResourceRevenueProductReportResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/revenue/products/{currency_code}"
	localVarPath = strings.Replace(localVarPath, "{"+"currency_code"+"}", fmt.Sprintf("%v", currencyCode), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(knetik_oauth)' required
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
	var successPayload = new(PageResourceRevenueProductReportResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ExecutiveRevenueProductSalesUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get refund revenue info
 * Get basic info about revenue loss from refunds (for all item types), summed up within a time range.
 *
 * @param currencyCode The code for a currency to get refund data for
 * @param startDate The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time
 * @param endDate The end of the time range to aggregate, unix timestamp in seconds. Default is end of time
 * @return *RevenueReportResource
 */
func (a ReportingRevenueApi) ExecutiveRevenueRefundsUsingGET(currencyCode string, startDate int64, endDate int64) (*RevenueReportResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/revenue/refunds/{currency_code}"
	localVarPath = strings.Replace(localVarPath, "{"+"currency_code"+"}", fmt.Sprintf("%v", currencyCode), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(knetik_oauth)' required
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
	var successPayload = new(RevenueReportResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ExecutiveRevenueRefundsUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Get subscription revenue info
 * Get basic info about revenue from sales of new subscriptions as well as recurring payemnts, summed up within a time range
 *
 * @param currencyCode The code for a currency to get sales data for
 * @param startDate The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time
 * @param endDate The end of the time range to aggregate, unix timestamp in seconds. Default is end of time
 * @return *RevenueReportResource
 */
func (a ReportingRevenueApi) ExecutiveRevenueSubscriptionSalesUsingGET(currencyCode string, startDate int64, endDate int64) (*RevenueReportResource, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/reporting/revenue/subscription-sales/{currency_code}"
	localVarPath = strings.Replace(localVarPath, "{"+"currency_code"+"}", fmt.Sprintf("%v", currencyCode), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(knetik_oauth)' required
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
	var successPayload = new(RevenueReportResource)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ExecutiveRevenueSubscriptionSalesUsingGET", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

