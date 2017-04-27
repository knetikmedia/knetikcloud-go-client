# \AuthClientsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClient**](AuthClientsApi.md#CreateClient) | **Post** /auth/clients | Create a new client
[**DeleteClient**](AuthClientsApi.md#DeleteClient) | **Delete** /auth/clients/{client_key} | Delete a client
[**GetClient**](AuthClientsApi.md#GetClient) | **Get** /auth/clients/{client_key} | Get a single client
[**GetClientGrantTypes**](AuthClientsApi.md#GetClientGrantTypes) | **Get** /auth/clients/grant-types | List available client grant types
[**GetClients**](AuthClientsApi.md#GetClients) | **Get** /auth/clients | List and search clients
[**SetClientGrantTypes**](AuthClientsApi.md#SetClientGrantTypes) | **Put** /auth/clients/{client_key}/grant-types | Set grant types for a client
[**SetClientRedirectUris**](AuthClientsApi.md#SetClientRedirectUris) | **Put** /auth/clients/{client_key}/redirect-uris | Set redirect uris for a client
[**UpdateClient**](AuthClientsApi.md#UpdateClient) | **Put** /auth/clients/{client_key} | Update a client


# **CreateClient**
> ClientResource CreateClient($clientResource)

Create a new client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientResource** | [**ClientResource**](ClientResource.md)| The client resource object | [optional] 

### Return type

[**ClientResource**](ClientResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClient**
> DeleteClient($clientKey)

Delete a client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The key of the client | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClient**
> ClientResource GetClient($clientKey)

Get a single client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The key of the client | 

### Return type

[**ClientResource**](ClientResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientGrantTypes**
> []GrantTypeResource GetClientGrantTypes()

List available client grant types


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]GrantTypeResource**](GrantTypeResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClients**
> PageResourceClientResource GetClients($size, $page, $order)

List and search clients


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceClientResource**](PageResource«ClientResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetClientGrantTypes**
> SetClientGrantTypes($clientKey, $grantList)

Set grant types for a client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The key of the client | 
 **grantList** | **[]string**| A list of unique grant types | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetClientRedirectUris**
> SetClientRedirectUris($clientKey, $redirectList)

Set redirect uris for a client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The key of the client | 
 **redirectList** | **[]string**| A list of unique redirect uris | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClient**
> ClientResource UpdateClient($clientKey, $clientResource)

Update a client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The key of the client | 
 **clientResource** | [**ClientResource**](ClientResource.md)| The client resource object | [optional] 

### Return type

[**ClientResource**](ClientResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

