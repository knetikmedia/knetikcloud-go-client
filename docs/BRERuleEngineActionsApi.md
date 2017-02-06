# \BRERuleEngineActionsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBREActions**](BRERuleEngineActionsApi.md#GetBREActions) | **Get** /bre/actions | Get a list of available actions


# **GetBREActions**
> []ActionResource GetBREActions($filterCategory, $filterName)

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

