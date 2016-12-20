# \ReportingUsageApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageByDayUsingGET**](ReportingUsageApi.md#GetUsageByDayUsingGET) | **Get** /reporting/usage/day | Returns aggregated endpoint usage information by the day
[**GetUsageByHourUsingGET**](ReportingUsageApi.md#GetUsageByHourUsingGET) | **Get** /reporting/usage/hour | Returns aggregated endpoint usage information by hour
[**GetUsageByMinuteUsingGET**](ReportingUsageApi.md#GetUsageByMinuteUsingGET) | **Get** /reporting/usage/minute | Returns aggregated endpoint usage information by minute
[**GetUsageByMonthUsingGET**](ReportingUsageApi.md#GetUsageByMonthUsingGET) | **Get** /reporting/usage/month | Returns aggregated endpoint usage information by month
[**GetUsageByYearUsingGET**](ReportingUsageApi.md#GetUsageByYearUsingGET) | **Get** /reporting/usage/year | Returns aggregated endpoint usage information by year


# **GetUsageByDayUsingGET**
> PageUsageInfo GetUsageByDayUsingGET($startDate, $endDate, $combineEndpoints, $size, $page)

Returns aggregated endpoint usage information by the day


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageUsageInfo**](Page«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByHourUsingGET**
> PageUsageInfo GetUsageByHourUsingGET($startDate, $endDate, $combineEndpoints, $size, $page)

Returns aggregated endpoint usage information by hour


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageUsageInfo**](Page«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMinuteUsingGET**
> PageUsageInfo GetUsageByMinuteUsingGET($startDate, $endDate, $combineEndpoints, $size, $page)

Returns aggregated endpoint usage information by minute


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageUsageInfo**](Page«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByMonthUsingGET**
> PageUsageInfo GetUsageByMonthUsingGET($startDate, $endDate, $combineEndpoints, $size, $page)

Returns aggregated endpoint usage information by month


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageUsageInfo**](Page«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByYearUsingGET**
> PageUsageInfo GetUsageByYearUsingGET($startDate, $endDate, $combineEndpoints, $size, $page)

Returns aggregated endpoint usage information by year


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64**| The beginning of the range being requested, unix timestamp in seconds | 
 **endDate** | **int64**| The ending of the range being requested, unix timestamp in seconds | 
 **combineEndpoints** | **bool**| Whether to combine counts from different endpoint. Removes the url and method from the result object | [optional] [default to false]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageUsageInfo**](Page«UsageInfo».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

