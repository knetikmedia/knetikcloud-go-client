# \StoreSubscriptionsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscriptionTemplateUsingPOST**](StoreSubscriptionsApi.md#CreateSubscriptionTemplateUsingPOST) | **Post** /subscriptions/templates | Create a subscription template
[**CreateSubscriptionUsingPOST**](StoreSubscriptionsApi.md#CreateSubscriptionUsingPOST) | **Post** /subscriptions | Creates a subscription item and associated plans
[**DeletePlanUsingDELETE**](StoreSubscriptionsApi.md#DeletePlanUsingDELETE) | **Delete** /subscriptions/{id}/plans/{plan_id} | Delete a subscription plan
[**DeleteSubscriptionTemplateUsingDELETE**](StoreSubscriptionsApi.md#DeleteSubscriptionTemplateUsingDELETE) | **Delete** /subscriptions/templates/{id} | Delete a subscription template
[**GetSubscriptionTemplateUsingGET**](StoreSubscriptionsApi.md#GetSubscriptionTemplateUsingGET) | **Get** /subscriptions/templates/{id} | Get a single subscription template
[**GetSubscriptionTemplatesUsingGET**](StoreSubscriptionsApi.md#GetSubscriptionTemplatesUsingGET) | **Get** /subscriptions/templates | List and search subscription templates
[**GetSubscriptionUsingGET**](StoreSubscriptionsApi.md#GetSubscriptionUsingGET) | **Get** /subscriptions/{id} | Retrieve a single subscription item and associated plans
[**ListSubscriptionsUsingGET**](StoreSubscriptionsApi.md#ListSubscriptionsUsingGET) | **Get** /subscriptions | List available subscription items and associated plans
[**ProcessUsingPOST**](StoreSubscriptionsApi.md#ProcessUsingPOST) | **Post** /subscriptions/process | Processes subscriptions and charge dues
[**UpdateSubscriptionTemplateUsingPUT**](StoreSubscriptionsApi.md#UpdateSubscriptionTemplateUsingPUT) | **Put** /subscriptions/templates/{id} | Update a subscription template
[**UpdateSubscriptionUsingPUT**](StoreSubscriptionsApi.md#UpdateSubscriptionUsingPUT) | **Put** /subscriptions/{id} | Updates a subscription item and associated plans


# **CreateSubscriptionTemplateUsingPOST**
> SubscriptionTemplateResource CreateSubscriptionTemplateUsingPOST($subscriptionTemplateResource)

Create a subscription template

Subscription Templates define a type of subscription and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionTemplateResource** | [**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)| The new subscription template | [optional] 

### Return type

[**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubscriptionUsingPOST**
> SubscriptionResource CreateSubscriptionUsingPOST($subscriptionResource)

Creates a subscription item and associated plans


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionResource** | [**SubscriptionResource**](SubscriptionResource.md)| The subscription to be created | [optional] 

### Return type

[**SubscriptionResource**](SubscriptionResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePlanUsingDELETE**
> DeletePlanUsingDELETE($id, $planId)

Delete a subscription plan

Must not be locked or a migration target


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the subscription | 
 **planId** | **string**| The id of the plan | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionTemplateUsingDELETE**
> DeleteSubscriptionTemplateUsingDELETE($id, $cascade)

Delete a subscription template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| force deleting the template if it&#39;s attached to other objects, cascade &#x3D; detach | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionTemplateUsingGET**
> SubscriptionTemplateResource GetSubscriptionTemplateUsingGET($id)

Get a single subscription template

Subscription Templates define a type of subscription and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionTemplatesUsingGET**
> PageResourceSubscriptionTemplateResource GetSubscriptionTemplatesUsingGET($size, $page, $order)

List and search subscription templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceSubscriptionTemplateResource**](PageResource«SubscriptionTemplateResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionUsingGET**
> SubscriptionResource GetSubscriptionUsingGET($id)

Retrieve a single subscription item and associated plans


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the subscription | 

### Return type

[**SubscriptionResource**](SubscriptionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubscriptionsUsingGET**
> PageResourceSubscriptionResource ListSubscriptionsUsingGET($size, $page, $order)

List available subscription items and associated plans


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceSubscriptionResource**](PageResource«SubscriptionResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProcessUsingPOST**
> ProcessUsingPOST()

Processes subscriptions and charge dues


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscriptionTemplateUsingPUT**
> UpdateSubscriptionTemplateUsingPUT($id, $subscriptionTemplateResource)

Update a subscription template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **subscriptionTemplateResource** | [**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)| The subscription template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscriptionUsingPUT**
> UpdateSubscriptionUsingPUT($id, $subscriptionResource)

Updates a subscription item and associated plans

Will not remove plans left out


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the subscription | 
 **subscriptionResource** | [**SubscriptionResource**](SubscriptionResource.md)| The subscription resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

