# \LogsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserLogUsingPOST**](LogsApi.md#AddUserLogUsingPOST) | **Post** /audit/logs | Add a new user log entry
[**GetEventLogUsingGET**](LogsApi.md#GetEventLogUsingGET) | **Get** /bre/logs/event-log/{id} | Get an existing BRE event log entry by id
[**GetEventsLogsUsingGET**](LogsApi.md#GetEventsLogsUsingGET) | **Get** /bre/logs/event-log | Returns a list of BRE event log entries
[**GetForwardLogUsingGET**](LogsApi.md#GetForwardLogUsingGET) | **Get** /bre/logs/forward-log/{id} | Get an existing forward log entry by id
[**GetForwardLogsUsingGET**](LogsApi.md#GetForwardLogsUsingGET) | **Get** /bre/logs/forward-log | Returns a list of forward log entries
[**GetUserLogsUsingGET**](LogsApi.md#GetUserLogsUsingGET) | **Get** /audit/logs/{id} | Returns a user log entry by id
[**GetUserLogsUsingGET1**](LogsApi.md#GetUserLogsUsingGET1) | **Get** /audit/logs | Returns a page of user logs entries


# **AddUserLogUsingPOST**
> AddUserLogUsingPOST($logEntry)

Add a new user log entry


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

# **GetEventLogUsingGET**
> BreEventLog GetEventLogUsingGET($id)

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

# **GetEventsLogsUsingGET**
> PageBreEventLog GetEventsLogsUsingGET($filterStartDate, $filterEventName, $size, $page, $order)

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

[**PageBreEventLog**](Page«BreEventLog».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForwardLogUsingGET**
> ForwardLog GetForwardLogUsingGET($id)

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

# **GetForwardLogsUsingGET**
> PageForwardLog GetForwardLogsUsingGET($filterStartDate, $filterEndDate, $filterStatusCode, $size, $page, $order)

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

[**PageForwardLog**](Page«ForwardLog».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLogsUsingGET**
> UserActionLog GetUserLogsUsingGET($id)

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

# **GetUserLogsUsingGET1**
> PageUserActionLog GetUserLogsUsingGET1($filterUser, $filterActionName, $size, $page)

Returns a page of user logs entries


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterUser** | **int32**| Filter for actions taken by a specific user by id | [optional] 
 **filterActionName** | **string**| Filter for actions of a specific name | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageUserActionLog**](Page«UserActionLog».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

