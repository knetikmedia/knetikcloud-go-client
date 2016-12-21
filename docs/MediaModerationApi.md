# \MediaModerationApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlagReportUsingGET**](MediaModerationApi.md#GetFlagReportUsingGET) | **Get** /moderation/reports/{id} | Get a flag report
[**GetFlagsReportUsingGET**](MediaModerationApi.md#GetFlagsReportUsingGET) | **Get** /moderation/reports | Returns a page of flag reports
[**SetFlagResolutionUsingPUT**](MediaModerationApi.md#SetFlagResolutionUsingPUT) | **Put** /moderation/reports/{id} | Update a flag report


# **GetFlagReportUsingGET**
> GetFlagReportUsingGET($id)

Get a flag report


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The flag report id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlagsReportUsingGET**
> PageResourceFlagReportResource GetFlagsReportUsingGET($excludeResolved, $filterContext, $size, $page)

Returns a page of flag reports

Context can be either a free-form string or a pre-defined context name


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeResolved** | **bool**| Ignore resolved context | [optional] [default to true]
 **filterContext** | **string**| Filter by moderation context | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceFlagReportResource**](PageResource«FlagReportResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetFlagResolutionUsingPUT**
> SetFlagResolutionUsingPUT($id, $flagReportResource)

Update a flag report

Lets you set the resolution of a report. Resolution types is {banned,ignore} in case of 'banned' you will need to pass the reason.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The flag report id | 
 **flagReportResource** | [**FlagReportResource**](FlagReportResource.md)| The new flag report | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

