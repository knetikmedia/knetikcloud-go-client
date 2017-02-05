# \UsersSubscriptionsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptionDetailsUsingGET**](UsersSubscriptionsApi.md#GetSubscriptionDetailsUsingGET) | **Get** /users/{user_id}/subscriptions/{inventory_id} | Get details about a user&#39;s subscription
[**GetSubscriptionDetailsUsingGET1**](UsersSubscriptionsApi.md#GetSubscriptionDetailsUsingGET1) | **Get** /users/{user_id}/subscriptions | Get details about a user&#39;s subscriptions
[**ReactivateUsingPOST**](UsersSubscriptionsApi.md#ReactivateUsingPOST) | **Post** /users/{user_id}/subscriptions/{inventory_id}/reactivate | Reactivate a subscription and charge fee
[**SetBillDateUsingPUT**](UsersSubscriptionsApi.md#SetBillDateUsingPUT) | **Put** /users/{user_id}/subscriptions/{inventory_id}/bill-date | Set a new date to bill a subscription on
[**SetPaymentMethodUsingPUT**](UsersSubscriptionsApi.md#SetPaymentMethodUsingPUT) | **Put** /users/{user_id}/subscriptions/{inventory_id}/payment-method | Set the payment method to use for a subscription
[**SetStatusUsingPUT**](UsersSubscriptionsApi.md#SetStatusUsingPUT) | **Put** /users/{user_id}/subscriptions/{inventory_id}/status | Set the status of a subscription
[**SetUserPlanUsingPUT**](UsersSubscriptionsApi.md#SetUserPlanUsingPUT) | **Put** /users/{user_id}/subscriptions/{inventory_id}/plan | Set a new subscription plan for a user


# **GetSubscriptionDetailsUsingGET**
> InventorySubscriptionResource GetSubscriptionDetailsUsingGET($userId, $inventoryId)

Get details about a user's subscription


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 

### Return type

[**InventorySubscriptionResource**](InventorySubscriptionResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionDetailsUsingGET1**
> []InventorySubscriptionResource GetSubscriptionDetailsUsingGET1($userId)

Get details about a user's subscriptions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 

### Return type

[**[]InventorySubscriptionResource**](InventorySubscriptionResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactivateUsingPOST**
> InvoiceResource ReactivateUsingPOST($userId, $inventoryId, $reactivateSubscriptionRequest)

Reactivate a subscription and charge fee


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **reactivateSubscriptionRequest** | [**ReactivateSubscriptionRequest**](ReactivateSubscriptionRequest.md)| The reactivate subscription request object inventory | [optional] 

### Return type

[**InvoiceResource**](InvoiceResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBillDateUsingPUT**
> SetBillDateUsingPUT($userId, $inventoryId, $billDate)

Set a new date to bill a subscription on


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **billDate** | **int64**| The new bill date. Unix timestamp in seconds | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPaymentMethodUsingPUT**
> SetPaymentMethodUsingPUT($userId, $inventoryId, $paymentMethodId)

Set the payment method to use for a subscription

May send null to use floating default


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **paymentMethodId** | **int32**| The id of the payment method | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetStatusUsingPUT**
> SetStatusUsingPUT($userId, $inventoryId, $status)

Set the status of a subscription

The body is a json string (put in quotes) that should match a desired invoice status type. Note that the new status may be blocked if the system is not configured to allow the current status to be changed to the new, to enforce proper flow. The default options for statuses are shown below but may be altered for special use cases


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **status** | **string**| The new status for the subscription. Actual options may differ from the indicated set if the invoice status type data has been altered.  Allowable values: (&#39;current&#39;, &#39;canceled&#39;, &#39;stopped&#39;, &#39;payment_failed&#39;, &#39;suspended&#39;) | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserPlanUsingPUT**
> SetUserPlanUsingPUT($userId, $inventoryId, $planId)

Set a new subscription plan for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **planId** | **string**| The id of the new plan. Must be from the same subscription | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

