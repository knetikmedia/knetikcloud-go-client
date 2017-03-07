# \ReportingUsageApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageByDay**](ReportingUsageApi.md#GetUsageByDay) | **Get** /reporting/usage/day | Returns aggregated endpoint usage information by day
[**GetUsageByHour**](ReportingUsageApi.md#GetUsageByHour) | **Get** /reporting/usage/hour | Returns aggregated endpoint usage information by hour
[**GetUsageByMinute**](ReportingUsageApi.md#GetUsageByMinute) | **Get** /reporting/usage/minute | Returns aggregated endpoint usage information by minute
[**GetUsageByMonth**](ReportingUsageApi.md#GetUsageByMonth) | **Get** /reporting/usage/month | Returns aggregated endpoint usage information by month
[**GetUsageByYear**](ReportingUsageApi.md#GetUsageByYear) | **Get** /reporting/usage/year | Returns aggregated endpoint usage information by year
[**GetUsageEndpoints**](ReportingUsageApi.md#GetUsageEndpoints) | **Get** /reporting/usage/endpoints | Returns list of endpoints called (method and url)


# **GetUsageByDay**
> PageResourceUsageInfo GetUsageByDay($startDate, $endDate, $combineEndpoints, $method, $url, $size, $page)

Returns aggregated endpoint usage information by day


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | [optional] 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByHour**
> PageResourceUsageInfo GetUsageByHour($startDate, $endDate, $combineEndpoints, $method, $url, $size, $page)

Returns aggregated endpoint usage information by hour


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | [optional] 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMinute**
> PageResourceUsageInfo GetUsageByMinute($startDate, $endDate, $combineEndpoints, $method, $url, $size, $page)

Returns aggregated endpoint usage information by minute


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | [optional] 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMonth**
> PageResourceUsageInfo GetUsageByMonth($startDate, $endDate, $combineEndpoints, $method, $url, $size, $page)

Returns aggregated endpoint usage information by month


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | [optional] 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByYear**
> PageResourceUsageInfo GetUsageByYear($startDate, $endDate, $combineEndpoints, $method, $url, $size, $page)

Returns aggregated endpoint usage information by year


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoints. Removes the url and method from the result object | [optional] [default to false]
 **method** | **string**| Filter for a certain endpoint method.  Must include url as well to work | [optional] 
 **url** | **string**| Filter for a certain endpoint.  Must include method as well to work | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUsageInfo**](PageResource«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageEndpoints**
> []string GetUsageEndpoints($startDate, $endDate)

Returns list of endpoints called (method and url)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 

### Return type

**[]string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

