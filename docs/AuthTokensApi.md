# \AuthTokensApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTokens**](AuthTokensApi.md#DeleteTokens) | **Delete** /auth/tokens | Delete tokens by username, client id, or both
[**GetToken**](AuthTokensApi.md#GetToken) | **Get** /auth/tokens/{username}/{client_id} | Get a single token by username and client id
[**GetTokens**](AuthTokensApi.md#GetTokens) | **Get** /auth/tokens | List usernames and client ids


# **DeleteTokens**
> DeleteTokens($username, $clientId)

Delete tokens by username, client id, or both


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| The username of the user | [optional] 
 **clientId** | **string**| The id of the client | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetToken**
> OauthAccessTokenResource GetToken($username, $clientId)

Get a single token by username and client id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| The username of the user | 
 **clientId** | **string**| The id of the client | 

### Return type

[**OauthAccessTokenResource**](OauthAccessTokenResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokens**
> PageResourceOauthAccessTokenResource GetTokens($filterClientId, $filterUsername, $size, $page, $order)

List usernames and client ids

Token value not shown


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterClientId** | **string**| Filters for token whose client id matches provided string | [optional] 
 **filterUsername** | **string**| Filters for token whose username matches provided string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] 

### Return type

[**PageResourceOauthAccessTokenResource**](PageResource«OauthAccessTokenResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

