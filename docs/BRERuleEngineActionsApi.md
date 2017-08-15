# \BRERuleEngineActionsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBREActions**](BRERuleEngineActionsApi.md#GetBREActions) | **Get** /bre/actions | Get a list of available actions


# **GetBREActions**
> []ActionResource GetBREActions(optional)
Get a list of available actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCategory** | **string**| Filter for actions that are within a specific category | 
 **filterName** | **string**| Filter for actions that have names containing the given string | 
 **filterTags** | **string**| Filter for actions that have all of the given tags (comma separated list) | 
 **filterSearch** | **string**| Filter for actions containing the given words somewhere within name, description and tags | 

### Return type

[**[]ActionResource**](ActionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

