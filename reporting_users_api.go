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
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type ReportingUsersApiService service


/* ReportingUsersApiService Get user registration info
 Get user registration counts grouped by time range

 @param optional (nil or map[string]interface{}) with one or more of:
     @param "granularity" (string) The time duration to aggregate by
     @param "startDate" (int64) The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time
     @param "endDate" (int64) The end of the time range to aggregate, unix timestamp in seconds. Default is end of time
     @param "size" (int32) The number of objects returned per page
     @param "page" (int32) The number of the page returned, starting with 1
 @return PageResourceAggregateCountResource*/
func (a *ReportingUsersApiService) GetUserRegistrations(localVarOptionals map[string]interface{}) (PageResourceAggregateCountResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PageResourceAggregateCountResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/reporting/users/registrations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["granularity"], "string", "granularity"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startDate"], "int64", "startDate"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endDate"], "int64", "endDate"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["size"], "int32", "size"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "int32", "page"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["granularity"].(string); localVarOk {
		localVarQueryParams.Add("granularity", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startDate"].(int64); localVarOk {
		localVarQueryParams.Add("start_date", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endDate"].(int64); localVarOk {
		localVarQueryParams.Add("end_date", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["size"].(int32); localVarOk {
		localVarQueryParams.Add("size", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(int32); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

