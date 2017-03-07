# \AccessTokenApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOAuthToken**](AccessTokenApi.md#GetOAuthToken) | **Post** /oauth/token | Get access token


# **GetOAuthToken**
> OAuth2Resource GetOAuthToken($grantType, $clientId, $clientSecret, $username, $password)

Get access token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string**| Grant type | [default to client_credentials]
 **clientId** | **string**| The id of the client | [default to knetik]
 **clientSecret** | **string**| The secret key of the client.  Used only with a grant_type of client_credentials | [optional] 
 **username** | **string**| The username of the client.  Used only with a grant_type of password | [optional] 
 **password** | **string**| The password of the client.  Used only with a grant_type of password | [optional] 

### Return type

[**OAuth2Resource**](OAuth2Resource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

