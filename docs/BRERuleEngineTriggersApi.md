# \BRERuleEngineTriggersApi

All URIs are relative to *https://devsandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTriggerUsingPOST**](BRERuleEngineTriggersApi.md#CreateTriggerUsingPOST) | **Post** /bre/triggers | Create a trigger
[**DeleteTriggerUsingDELETE**](BRERuleEngineTriggersApi.md#DeleteTriggerUsingDELETE) | **Delete** /bre/triggers/{event_name} | Delete a trigger
[**GetTriggerUsingGET**](BRERuleEngineTriggersApi.md#GetTriggerUsingGET) | **Get** /bre/triggers/{event_name} | Get a single trigger
[**GetTriggersUsingGET**](BRERuleEngineTriggersApi.md#GetTriggersUsingGET) | **Get** /bre/triggers | List triggers
[**UpdateTriggerUsingPUT**](BRERuleEngineTriggersApi.md#UpdateTriggerUsingPUT) | **Put** /bre/triggers/{event_name} | Update a trigger


# **CreateTriggerUsingPOST**
> BreTriggerResource CreateTriggerUsingPOST($breTriggerResource)

Create a trigger

Customer added triggers will not be fired automatically or have rules associated with them by default. Custom rules must be added to get use from the trigger and it must then be fired from the outside. See the Bre Event services


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breTriggerResource** | [**BreTriggerResource**](BreTriggerResource.md)| The BRE trigger resource object | [optional] 

### Return type

[**BreTriggerResource**](BreTriggerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTriggerUsingDELETE**
> DeleteTriggerUsingDELETE($eventName)

Delete a trigger

May fail if there are existing rules against it. Cannot delete core triggers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventName** | **string**| The trigger event name | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerUsingGET**
> BreTriggerResource GetTriggerUsingGET($eventName)

Get a single trigger


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventName** | **string**| The trigger event name | 

### Return type

[**BreTriggerResource**](BreTriggerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggersUsingGET**
> PageBreTriggerResource GetTriggersUsingGET($filterSystem, $filterCategory, $filterName, $size, $page)

List triggers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSystem** | **bool**| Filter for triggers that are system triggers when true, or not when false. Leave off for both mixed | [optional] 
 **filterCategory** | **string**| Filter for triggers that are within a specific category | [optional] 
 **filterName** | **string**| Filter for triggers that have names containing the given string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageBreTriggerResource**](Page«BreTriggerResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTriggerUsingPUT**
> UpdateTriggerUsingPUT($eventName, $breTriggerResource)

Update a trigger

May fail if new parameters mismatch requirements of existing rules. Cannot update core triggers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventName** | **string**| The trigger event name | 
 **breTriggerResource** | [**BreTriggerResource**](BreTriggerResource.md)| The BRE trigger resource object | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

