# \AuthClientsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignClientGrantTypesUsingPUT**](AuthClientsApi.md#AssignClientGrantTypesUsingPUT) | **Put** /auth/clients/{client_key}/grant-types | Set grant types for a client
[**AssignClientRedirectUrisUsingPUT**](AuthClientsApi.md#AssignClientRedirectUrisUsingPUT) | **Put** /auth/clients/{client_key}/redirect-uris | Set redirect uris for a client
[**CreateClientUsingPOST**](AuthClientsApi.md#CreateClientUsingPOST) | **Post** /auth/clients | Create a new client
[**DeleteClientUsingDELETE**](AuthClientsApi.md#DeleteClientUsingDELETE) | **Delete** /auth/clients/{client_key} | Delete a client
[**GetClientUsingGET**](AuthClientsApi.md#GetClientUsingGET) | **Get** /auth/clients/{client_key} | Get a single client
[**GetClientsUsingGET**](AuthClientsApi.md#GetClientsUsingGET) | **Get** /auth/clients | List and search clients
[**ListAvailableGrantTypesUsingGET**](AuthClientsApi.md#ListAvailableGrantTypesUsingGET) | **Get** /auth/clients/grant-types | List available client grant types
[**UpdateClientUsingPUT**](AuthClientsApi.md#UpdateClientUsingPUT) | **Put** /auth/clients/{client_key} | Update a client


# **AssignClientGrantTypesUsingPUT**
> AssignClientGrantTypesUsingPUT($clientKey, $grantList)

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

# **AssignClientRedirectUrisUsingPUT**
> AssignClientRedirectUrisUsingPUT($clientKey, $redirectList)

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

# **CreateClientUsingPOST**
> ClientResource CreateClientUsingPOST($clientResource)

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

# **DeleteClientUsingDELETE**
> DeleteClientUsingDELETE($clientKey)

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

# **GetClientUsingGET**
> ClientResource GetClientUsingGET($clientKey)

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

# **GetClientsUsingGET**
> PageClientResource GetClientsUsingGET($size, $page, $order)

List and search clients


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageClientResource**](Page«ClientResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAvailableGrantTypesUsingGET**
> []GrantTypeResource ListAvailableGrantTypesUsingGET()

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

# **UpdateClientUsingPUT**
> UpdateClientUsingPUT($clientKey, $clientResource)

Update a client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The key of the client | 
 **clientResource** | [**ClientResource**](ClientResource.md)| The client resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

