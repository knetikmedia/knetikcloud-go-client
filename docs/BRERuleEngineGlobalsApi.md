# \BRERuleEngineGlobalsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBREGlobal**](BRERuleEngineGlobalsApi.md#CreateBREGlobal) | **Post** /bre/globals/definitions | Create a global definition
[**DeleteBREGlobal**](BRERuleEngineGlobalsApi.md#DeleteBREGlobal) | **Delete** /bre/globals/definitions/{id} | Delete a global
[**GetBREGlobal**](BRERuleEngineGlobalsApi.md#GetBREGlobal) | **Get** /bre/globals/definitions/{id} | Get a single global definition
[**GetBREGlobals**](BRERuleEngineGlobalsApi.md#GetBREGlobals) | **Get** /bre/globals/definitions | List global definitions
[**UpdateBREGlobal**](BRERuleEngineGlobalsApi.md#UpdateBREGlobal) | **Put** /bre/globals/definitions/{id} | Update a global definition


# **CreateBREGlobal**
> BreGlobalResource CreateBREGlobal($breGlobalResource)

Create a global definition

Once created you can then use in a custom rule. Note that global definitions cannot be modified or deleted if in use.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breGlobalResource** | [**BreGlobalResource**](BreGlobalResource.md)| The BRE global resource object | [optional] 

### Return type

[**BreGlobalResource**](BreGlobalResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBREGlobal**
> DeleteBREGlobal($id)

Delete a global

May fail if there are existing rules against it. Cannot delete core globals


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the global definition | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREGlobal**
> BreGlobalResource GetBREGlobal($id)

Get a single global definition


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the global definition | 

### Return type

[**BreGlobalResource**](BreGlobalResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREGlobals**
> PageResourceBreGlobalResource GetBREGlobals($filterSystem, $size, $page)

List global definitions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSystem** | **bool**| Filter for globals that are system globals when true, or not when false. Leave off for both mixed | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceBreGlobalResource**](PageResource«BreGlobalResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBREGlobal**
> BreGlobalResource UpdateBREGlobal($id, $breGlobalResource)

Update a global definition

May fail if new parameters mismatch requirements of existing rules. Cannot update core globals


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the global definition | 
 **breGlobalResource** | [**BreGlobalResource**](BreGlobalResource.md)| The BRE global resource object | [optional] 

### Return type

[**BreGlobalResource**](BreGlobalResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

