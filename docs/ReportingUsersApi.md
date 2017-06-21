# \ReportingUsersApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserRegistrations**](ReportingUsersApi.md#GetUserRegistrations) | **Get** /reporting/users/registrations | Get user registration info


# **GetUserRegistrations**
> PageResourceAggregateCountResource GetUserRegistrations($granularity, $startDate, $endDate, $size, $page)

Get user registration info

Get user registration counts grouped by time range


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **granularity** | **string**| The time duration to aggregate by | [optional] [default to day]
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceAggregateCountResource**](PageResource«AggregateCountResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

