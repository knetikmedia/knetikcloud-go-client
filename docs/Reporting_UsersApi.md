# \Reporting_UsersApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserRegistrations**](Reporting_UsersApi.md#GetUserRegistrations) | **Get** /reporting/users/registrations | Get user registration info


# **GetUserRegistrations**
> PageResourceAggregateCountResource GetUserRegistrations(ctx, ctx, optional)
Get user registration info

Get user registration counts grouped by time range. <br><br><b>Permissions Needed:</b> REPORTING_USER_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **granularity** | **string**| The time duration to aggregate by | [default to day]
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceAggregateCountResource**](PageResource«AggregateCountResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

