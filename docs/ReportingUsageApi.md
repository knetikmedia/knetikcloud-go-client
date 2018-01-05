# \ReportingUsageApi

All URIs are relative to *https://devsandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageByDay**](ReportingUsageApi.md#GetUsageByDay) | **Get** /reporting/usage/day | Returns aggregated endpoint usage information by day
[**GetUsageByHour**](ReportingUsageApi.md#GetUsageByHour) | **Get** /reporting/usage/hour | Returns aggregated endpoint usage information by hour
[**GetUsageByMinute**](ReportingUsageApi.md#GetUsageByMinute) | **Get** /reporting/usage/minute | Returns aggregated endpoint usage information by minute
[**GetUsageByMonth**](ReportingUsageApi.md#GetUsageByMonth) | **Get** /reporting/usage/month | Returns aggregated endpoint usage information by month
[**GetUsageByYear**](ReportingUsageApi.md#GetUsageByYear) | **Get** /reporting/usage/year | Returns aggregated endpoint usage information by year
[**GetUsageEndpoints**](ReportingUsageApi.md#GetUsageEndpoints) | **Get** /reporting/usage/endpoints | Returns list of endpoints called (method and url)


# **GetUsageByDay**
> PageResourceUsageInfo GetUsageByDay(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
  **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByHour**
> PageResourceUsageInfo GetUsageByHour(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by hour

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
  **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMinute**
> PageResourceUsageInfo GetUsageByMinute(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by minute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
  **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMonth**
> PageResourceUsageInfo GetUsageByMonth(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by month

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
  **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByYear**
> PageResourceUsageInfo GetUsageByYear(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by year

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
  **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoints. Removes the url and method from the result object | [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageEndpoints**
> []string GetUsageEndpoints(ctx, ctx, startDate, endDate)
Returns list of endpoints called (method and url)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
  **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 

### Return type

**[]string**

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

