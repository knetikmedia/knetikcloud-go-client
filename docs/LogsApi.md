# \LogsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserLog**](LogsApi.md#AddUserLog) | **Post** /audit/logs | Add a user log entry
[**GetBREEventLog**](LogsApi.md#GetBREEventLog) | **Get** /bre/logs/event-log/{id} | Get an existing BRE event log entry by id
[**GetBREEventLogs**](LogsApi.md#GetBREEventLogs) | **Get** /bre/logs/event-log | Returns a list of BRE event log entries
[**GetBREForwardLog**](LogsApi.md#GetBREForwardLog) | **Get** /bre/logs/forward-log/{id} | Get an existing forward log entry by id
[**GetBREForwardLogs**](LogsApi.md#GetBREForwardLogs) | **Get** /bre/logs/forward-log | Returns a list of forward log entries
[**GetUserLog**](LogsApi.md#GetUserLog) | **Get** /audit/logs/{id} | Returns a user log entry by id
[**GetUserLogs**](LogsApi.md#GetUserLogs) | **Get** /audit/logs | Returns a page of user logs entries


# **AddUserLog**
> AddUserLog($logEntry)

Add a user log entry


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logEntry** | [**UserActionLog**](UserActionLog.md)| The user log entry to be added | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREEventLog**
> BreEventLog GetBREEventLog($id)

Get an existing BRE event log entry by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The BRE event log entry id | 

### Return type

[**BreEventLog**](BreEventLog.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREEventLogs**
> PageResourceBreEventLog GetBREEventLogs($filterStartDate, $filterEventName, $size, $page, $order)

Returns a list of BRE event log entries


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStartDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the event log start date, a unix timestamp in seconds.  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 
 **filterEventName** | **string**| Filter event logs by event name | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:DESC]

### Return type

[**PageResourceBreEventLog**](PageResource«BreEventLog».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREForwardLog**
> ForwardLog GetBREForwardLog($id)

Get an existing forward log entry by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The forward log entry id | 

### Return type

[**ForwardLog**](ForwardLog.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREForwardLogs**
> PageResourceForwardLog GetBREForwardLogs($filterStartDate, $filterEndDate, $filterStatusCode, $size, $page, $order)

Returns a list of forward log entries


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStartDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the log start date, a unix timestamp in seconds.  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 
 **filterEndDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the log end date, a unix timestamp in seconds.  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 
 **filterStatusCode** | **int32**| Filter forward logs by http status code | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:DESC]

### Return type

[**PageResourceForwardLog**](PageResource«ForwardLog».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLog**
> UserActionLog GetUserLog($id)

Returns a user log entry by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The user log entry id | 

### Return type

[**UserActionLog**](UserActionLog.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLogs**
> PageResourceUserActionLog GetUserLogs($filterUser, $filterActionName, $size, $page)

Returns a page of user logs entries


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterUser** | **int32**| Filter for actions taken by a specific user by id | [optional] 
 **filterActionName** | **string**| Filter for actions of a specific name | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUserActionLog**](PageResource«UserActionLog».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

