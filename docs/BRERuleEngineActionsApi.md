# \BRERuleEngineActionsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActionsUsingGET**](BRERuleEngineActionsApi.md#GetActionsUsingGET) | **Get** /bre/actions | Get a list of available actions


# **GetActionsUsingGET**
> []ActionResource GetActionsUsingGET($filterCategory, $filterName)

Get a list of available actions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCategory** | **string**| Filter for actions that are within a specific category | [optional] 
 **filterName** | **string**| Filter for actions that have names containing the given string | [optional] 

### Return type

[**[]ActionResource**](ActionResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

