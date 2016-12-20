# \ConfigsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigUsingPOST**](ConfigsApi.md#CreateConfigUsingPOST) | **Post** /configs | Create a new config
[**DeleteConfigUsingDELETE**](ConfigsApi.md#DeleteConfigUsingDELETE) | **Delete** /configs/{name} | Delete an existing config
[**GetConfigUsingGET**](ConfigsApi.md#GetConfigUsingGET) | **Get** /configs/{name} | Get a single config
[**GetConfigsUsingGET**](ConfigsApi.md#GetConfigsUsingGET) | **Get** /configs | List and search configs
[**UpdateConfigUsingPUT**](ConfigsApi.md#UpdateConfigUsingPUT) | **Put** /configs/{name} | Update an existing config


# **CreateConfigUsingPOST**
> Config CreateConfigUsingPOST($config)

Create a new config


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | [**Config**](Config.md)| The config object | [optional] 

### Return type

[**Config**](Config.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigUsingDELETE**
> DeleteConfigUsingDELETE($name)

Delete an existing config


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The config name | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigUsingGET**
> Config GetConfigUsingGET($name)

Get a single config

Only configs that are public readable will be shown without admin access


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The config name | 

### Return type

[**Config**](Config.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigsUsingGET**
> PageConfig GetConfigsUsingGET($filterSearch, $size, $page, $order)

List and search configs


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSearch** | **string**| Filter for configs whose name contains the given string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageConfig**](Page«Config».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfigUsingPUT**
> UpdateConfigUsingPUT($name, $config)

Update an existing config


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The config name | 
 **config** | [**Config**](Config.md)| The config object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

