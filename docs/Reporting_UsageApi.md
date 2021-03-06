# \Reporting_UsageApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageByDay**](Reporting_UsageApi.md#GetUsageByDay) | **Get** /reporting/usage/day | Returns aggregated endpoint usage information by day
[**GetUsageByHour**](Reporting_UsageApi.md#GetUsageByHour) | **Get** /reporting/usage/hour | Returns aggregated endpoint usage information by hour
[**GetUsageByMinute**](Reporting_UsageApi.md#GetUsageByMinute) | **Get** /reporting/usage/minute | Returns aggregated endpoint usage information by minute
[**GetUsageByMonth**](Reporting_UsageApi.md#GetUsageByMonth) | **Get** /reporting/usage/month | Returns aggregated endpoint usage information by month
[**GetUsageByYear**](Reporting_UsageApi.md#GetUsageByYear) | **Get** /reporting/usage/year | Returns aggregated endpoint usage information by year
[**GetUsageEndpoints**](Reporting_UsageApi.md#GetUsageEndpoints) | **Get** /reporting/usage/endpoints | Returns list of endpoints called (method and url)


# **GetUsageByDay**
> PageResourceUsageInfo GetUsageByDay(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by day

<b>Permissions Needed:</b> USAGE_ADMIN

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByHour**
> PageResourceUsageInfo GetUsageByHour(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by hour

<b>Permissions Needed:</b> USAGE_ADMIN

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMinute**
> PageResourceUsageInfo GetUsageByMinute(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by minute

<b>Permissions Needed:</b> USAGE_ADMIN

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMonth**
> PageResourceUsageInfo GetUsageByMonth(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by month

<b>Permissions Needed:</b> USAGE_ADMIN

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByYear**
> PageResourceUsageInfo GetUsageByYear(ctx, ctx, startDate, endDate, optional)
Returns aggregated endpoint usage information by year

<b>Permissions Needed:</b> USAGE_ADMIN

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageEndpoints**
> []string GetUsageEndpoints(ctx, ctx, startDate, endDate)
Returns list of endpoints called (method and url)

<b>Permissions Needed:</b> USAGE_ADMIN

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

