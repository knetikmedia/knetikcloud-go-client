# \BRERuleEngineTriggersApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBRETrigger**](BRERuleEngineTriggersApi.md#CreateBRETrigger) | **Post** /bre/triggers | Create a trigger
[**DeleteBRETrigger**](BRERuleEngineTriggersApi.md#DeleteBRETrigger) | **Delete** /bre/triggers/{event_name} | Delete a trigger
[**GetBRETrigger**](BRERuleEngineTriggersApi.md#GetBRETrigger) | **Get** /bre/triggers/{event_name} | Get a single trigger
[**GetBRETriggers**](BRERuleEngineTriggersApi.md#GetBRETriggers) | **Get** /bre/triggers | List triggers
[**UpdateBRETrigger**](BRERuleEngineTriggersApi.md#UpdateBRETrigger) | **Put** /bre/triggers/{event_name} | Update a trigger


# **CreateBRETrigger**
> BreTriggerResource CreateBRETrigger(optional)
Create a trigger

Customer added triggers will not be fired automatically or have rules associated with them by default. Custom rules must be added to get use from the trigger and it must then be fired from the outside. See the Bre Event services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breTriggerResource** | [**BreTriggerResource**](BreTriggerResource.md)| The BRE trigger resource object | 

### Return type

[**BreTriggerResource**](BreTriggerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBRETrigger**
> DeleteBRETrigger(eventName)
Delete a trigger

May fail if there are existing rules against it. Cannot delete core triggers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **eventName** | **string**| The trigger event name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBRETrigger**
> BreTriggerResource GetBRETrigger(eventName)
Get a single trigger

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **eventName** | **string**| The trigger event name | 

### Return type

[**BreTriggerResource**](BreTriggerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBRETriggers**
> PageResourceBreTriggerResource GetBRETriggers(optional)
List triggers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSystem** | **bool**| Filter for triggers that are system triggers when true, or not when false. Leave off for both mixed | 
 **filterCategory** | **string**| Filter for triggers that are within a specific category | 
 **filterTags** | **string**| Filter for triggers that have all of the given tags (comma separated list) | 
 **filterName** | **string**| Filter for triggers that have names containing the given string | 
 **filterSearch** | **string**| Filter for triggers containing the given words somewhere within name, description and tags | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceBreTriggerResource**](PageResource«BreTriggerResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBRETrigger**
> BreTriggerResource UpdateBRETrigger(eventName, optional)
Update a trigger

May fail if new parameters mismatch requirements of existing rules. Cannot update core triggers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **eventName** | **string**| The trigger event name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventName** | **string**| The trigger event name | 
 **breTriggerResource** | [**BreTriggerResource**](BreTriggerResource.md)| The BRE trigger resource object | 

### Return type

[**BreTriggerResource**](BreTriggerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

