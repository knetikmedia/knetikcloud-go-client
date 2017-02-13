# \UtilSecurityApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserLocationLog**](UtilSecurityApi.md#GetUserLocationLog) | **Get** /security/country-log | Returns the authentication log for a user
[**GetUserTokenDetails**](UtilSecurityApi.md#GetUserTokenDetails) | **Get** /me | Returns the authentication token details. Use /users endpoint for detailed user&#39;s info


# **GetUserLocationLog**
> PageResourceLocationLogResource GetUserLocationLog($userId)

Returns the authentication log for a user

A log entry is recorded everytime a user requests a new token. Standard pagination available


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user id | [optional] 

### Return type

[**PageResourceLocationLogResource**](PageResource«LocationLogResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserTokenDetails**
> TokenDetailsResource GetUserTokenDetails()

Returns the authentication token details. Use /users endpoint for detailed user's info


### Parameters
This endpoint does not need any parameter.

### Return type

[**TokenDetailsResource**](TokenDetailsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

