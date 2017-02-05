# \BRERuleEngineEventsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FireEventUsingPOST**](BRERuleEngineEventsApi.md#FireEventUsingPOST) | **Post** /bre/events | Fire a new event, based on an existing trigger


# **FireEventUsingPOST**
> FireEventUsingPOST($breEvent)

Fire a new event, based on an existing trigger

Parameters within the event must match names and types from the trigger. Actual rule execution is asynchornous.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breEvent** | [**BreEvent**](BreEvent.md)| The BRE event object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

