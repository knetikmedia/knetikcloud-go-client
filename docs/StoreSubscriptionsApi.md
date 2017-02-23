# \StoreSubscriptionsApi

All URIs are relative to *https://sandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](StoreSubscriptionsApi.md#CreateSubscription) | **Post** /subscriptions | Creates a subscription item and associated plans
[**CreateSubscriptionTemplate**](StoreSubscriptionsApi.md#CreateSubscriptionTemplate) | **Post** /subscriptions/templates | Create a subscription template
[**DeleteSubscription**](StoreSubscriptionsApi.md#DeleteSubscription) | **Delete** /subscriptions/{id}/plans/{plan_id} | Delete a subscription plan
[**DeleteSubscriptionTemplate**](StoreSubscriptionsApi.md#DeleteSubscriptionTemplate) | **Delete** /subscriptions/templates/{id} | Delete a subscription template
[**GetSubscription**](StoreSubscriptionsApi.md#GetSubscription) | **Get** /subscriptions/{id} | Retrieve a single subscription item and associated plans
[**GetSubscriptionTemplate**](StoreSubscriptionsApi.md#GetSubscriptionTemplate) | **Get** /subscriptions/templates/{id} | Get a single subscription template
[**GetSubscriptionTemplates**](StoreSubscriptionsApi.md#GetSubscriptionTemplates) | **Get** /subscriptions/templates | List and search subscription templates
[**GetSubscriptions**](StoreSubscriptionsApi.md#GetSubscriptions) | **Get** /subscriptions | List available subscription items and associated plans
[**ProcessSubscriptions**](StoreSubscriptionsApi.md#ProcessSubscriptions) | **Post** /subscriptions/process | Processes subscriptions and charge dues
[**UpdateSubscription**](StoreSubscriptionsApi.md#UpdateSubscription) | **Put** /subscriptions/{id} | Updates a subscription item and associated plans
[**UpdateSubscriptionTemplate**](StoreSubscriptionsApi.md#UpdateSubscriptionTemplate) | **Put** /subscriptions/templates/{id} | Update a subscription template


# **CreateSubscription**
> SubscriptionResource CreateSubscription($subscriptionResource)

Creates a subscription item and associated plans


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionResource** | [**SubscriptionResource**](SubscriptionResource.md)| The subscription to be created | [optional] 

### Return type

[**SubscriptionResource**](SubscriptionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubscriptionTemplate**
> SubscriptionTemplateResource CreateSubscriptionTemplate($subscriptionTemplateResource)

Create a subscription template

Subscription Templates define a type of subscription and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionTemplateResource** | [**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)| The new subscription template | [optional] 

### Return type

[**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscription**
> DeleteSubscription($id, $planId)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionTemplate**
> DeleteSubscriptionTemplate($id, $cascade)

Delete a subscription template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| force deleting the template if it&#39;s attached to other objects, cascade &#x3D; detach | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscription**
> SubscriptionResource GetSubscription($id)

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

# **GetSubscriptionTemplate**
> SubscriptionTemplateResource GetSubscriptionTemplate($id)

Get a single subscription template

Subscription Templates define a type of subscription and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionTemplates**
> PageResourceSubscriptionTemplateResource GetSubscriptionTemplates($size, $page, $order)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptions**
> PageResourceSubscriptionResource GetSubscriptions($size, $page, $order)

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

# **ProcessSubscriptions**
> ProcessSubscriptions()

Processes subscriptions and charge dues


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscription**
> UpdateSubscription($id, $subscriptionResource)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscriptionTemplate**
> SubscriptionTemplateResource UpdateSubscriptionTemplate($id, $subscriptionTemplateResource)

Update a subscription template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **subscriptionTemplateResource** | [**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)| The subscription template resource object | [optional] 

### Return type

[**SubscriptionTemplateResource**](SubscriptionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

