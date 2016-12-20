# \UsersAddressesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddressUsingPOST**](UsersAddressesApi.md#CreateAddressUsingPOST) | **Post** /users/{user_id}/addresses | Save a new address
[**DeleteAddressUsingDELETE**](UsersAddressesApi.md#DeleteAddressUsingDELETE) | **Delete** /users/{user_id}/addresses/{id} | Delete an address
[**GetAddressUsingGET**](UsersAddressesApi.md#GetAddressUsingGET) | **Get** /users/{user_id}/addresses/{id} | Get a single address
[**GetAddressesUsingGET**](UsersAddressesApi.md#GetAddressesUsingGET) | **Get** /users/{user_id}/addresses | List and search addresses
[**UpdateAddressUsingPUT**](UsersAddressesApi.md#UpdateAddressUsingPUT) | **Put** /users/{user_id}/addresses/{id} | Update an address


# **CreateAddressUsingPOST**
> SavedAddressResource CreateAddressUsingPOST($userId, $savedAddressResource)

Save a new address


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user | 
 **savedAddressResource** | [**SavedAddressResource**](SavedAddressResource.md)| The new address | [optional] 

### Return type

[**SavedAddressResource**](SavedAddressResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAddressUsingDELETE**
> DeleteAddressUsingDELETE($userId, $id)

Delete an address


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user | 
 **id** | **int32**| The id of the address | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressUsingGET**
> SavedAddressResource GetAddressUsingGET($userId, $id)

Get a single address


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user | 
 **id** | **int32**| The id of the address | 

### Return type

[**SavedAddressResource**](SavedAddressResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressesUsingGET**
> PageResourceSavedAddressResource GetAddressesUsingGET($userId, $size, $page, $order)

List and search addresses


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceSavedAddressResource**](PageResource«SavedAddressResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAddressUsingPUT**
> UpdateAddressUsingPUT($userId, $id, $savedAddressResource)

Update an address


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user | 
 **id** | **int32**| The id of the address | 
 **savedAddressResource** | [**SavedAddressResource**](SavedAddressResource.md)| The saved address resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
