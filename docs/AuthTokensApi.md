# \AuthTokensApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTokenUsingDELETE**](AuthTokensApi.md#DeleteTokenUsingDELETE) | **Delete** /auth/tokens/{username} | Delete all tokens by username
[**DeleteTokenWithClientIdUsingDELETE**](AuthTokensApi.md#DeleteTokenWithClientIdUsingDELETE) | **Delete** /auth/tokens/{username}/{client_id} | Delete a token by username and client id
[**GetTokenByUserUsingGET**](AuthTokensApi.md#GetTokenByUserUsingGET) | **Get** /auth/tokens/{username}/{client_id} | Get a single token by username and client id
[**GetTokensUsingGET**](AuthTokensApi.md#GetTokensUsingGET) | **Get** /auth/tokens | List usernames and client ids


# **DeleteTokenUsingDELETE**
> DeleteTokenUsingDELETE($username)

Delete all tokens by username


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| The username of the user | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTokenWithClientIdUsingDELETE**
> DeleteTokenWithClientIdUsingDELETE($username, $clientId)

Delete a token by username and client id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| The username of the user | 
 **clientId** | **string**| The id of the client | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenByUserUsingGET**
> OauthAccessTokenResource GetTokenByUserUsingGET($username, $clientId)

Get a single token by username and client id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| The username of the user | 
 **clientId** | **string**| The id of the client | 

### Return type

[**OauthAccessTokenResource**](OauthAccessTokenResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokensUsingGET**
> PageResourceOauthAccessTokenResource GetTokensUsingGET($filterClientId, $filterUsername, $size, $page, $order)

List usernames and client ids

Token value not shown


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterClientId** | **string**| Filters for token whose client id matches provided string | [optional] 
 **filterUsername** | **string**| Filters for token whose username matches provided string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to username:ASC]

### Return type

[**PageResourceOauthAccessTokenResource**](PageResource«OauthAccessTokenResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

