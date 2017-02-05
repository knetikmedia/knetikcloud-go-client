# \UtilSecurityApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserLocationLogUsingGET**](UtilSecurityApi.md#GetUserLocationLogUsingGET) | **Get** /security/country-log | Returns the authentication log for a user
[**UserUsingGET**](UtilSecurityApi.md#UserUsingGET) | **Get** /me | Returns the authentication token details. Use /users endpoint for detailed user&#39;s info


# **GetUserLocationLogUsingGET**
> PageResourceLocationLogResource GetUserLocationLogUsingGET($userId)

Returns the authentication log for a user

A log entry is recorded everytime a user requests a new token. Standard pagination available


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user id | [optional] 

### Return type

[**PageResourceLocationLogResource**](PageResource«LocationLogResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUsingGET**
> TokenDetailsResource UserUsingGET()

Returns the authentication token details. Use /users endpoint for detailed user's info


### Parameters
This endpoint does not need any parameter.

### Return type

[**TokenDetailsResource**](TokenDetailsResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

