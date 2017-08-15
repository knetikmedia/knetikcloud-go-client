# \UsersSubscriptionsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSubscriptionDetails**](UsersSubscriptionsApi.md#GetUserSubscriptionDetails) | **Get** /users/{user_id}/subscriptions/{inventory_id} | Get details about a user&#39;s subscription
[**GetUsersSubscriptionDetails**](UsersSubscriptionsApi.md#GetUsersSubscriptionDetails) | **Get** /users/{user_id}/subscriptions | Get details about a user&#39;s subscriptions
[**ReactivateUserSubscription**](UsersSubscriptionsApi.md#ReactivateUserSubscription) | **Post** /users/{user_id}/subscriptions/{inventory_id}/reactivate | Reactivate a subscription and charge fee
[**SetSubscriptionBillDate**](UsersSubscriptionsApi.md#SetSubscriptionBillDate) | **Put** /users/{user_id}/subscriptions/{inventory_id}/bill-date | Set a new date to bill a subscription on
[**SetSubscriptionPaymentMethod**](UsersSubscriptionsApi.md#SetSubscriptionPaymentMethod) | **Put** /users/{user_id}/subscriptions/{inventory_id}/payment-method | Set the payment method to use for a subscription
[**SetSubscriptionStatus**](UsersSubscriptionsApi.md#SetSubscriptionStatus) | **Put** /users/{user_id}/subscriptions/{inventory_id}/status | Set the status of a subscription
[**SetUserSubscriptionPlan**](UsersSubscriptionsApi.md#SetUserSubscriptionPlan) | **Put** /users/{user_id}/subscriptions/{inventory_id}/plan | Set a new subscription plan for a user
[**SetUserSubscriptionPrice**](UsersSubscriptionsApi.md#SetUserSubscriptionPrice) | **Put** /users/{user_id}/subscriptions/{inventory_id}/price-override | Set a new subscription price for a user


# **GetUserSubscriptionDetails**
> InventorySubscriptionResource GetUserSubscriptionDetails(userId, inventoryId)
Get details about a user's subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 

### Return type

[**InventorySubscriptionResource**](InventorySubscriptionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersSubscriptionDetails**
> []InventorySubscriptionResource GetUsersSubscriptionDetails(userId)
Get details about a user's subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 

### Return type

[**[]InventorySubscriptionResource**](InventorySubscriptionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactivateUserSubscription**
> InvoiceResource ReactivateUserSubscription(userId, inventoryId, optional)
Reactivate a subscription and charge fee

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **reactivateSubscriptionRequest** | [**ReactivateSubscriptionRequest**](ReactivateSubscriptionRequest.md)| The reactivate subscription request object inventory | 

### Return type

[**InvoiceResource**](InvoiceResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSubscriptionBillDate**
> SetSubscriptionBillDate(userId, inventoryId, billDate)
Set a new date to bill a subscription on

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
  **billDate** | **int64**| The new bill date. Unix timestamp in seconds | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSubscriptionPaymentMethod**
> SetSubscriptionPaymentMethod(userId, inventoryId, optional)
Set the payment method to use for a subscription

May send null to use floating default

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **paymentMethodId** | [**IntWrapper**](IntWrapper.md)| The id of the payment method | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSubscriptionStatus**
> SetSubscriptionStatus(userId, inventoryId, status)
Set the status of a subscription

Note that the new status may be blocked if the system is not configured to allow the current status to be changed to the new, to enforce proper flow. The default options for statuses are shown below but may be altered for special use cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
  **status** | [**StringWrapper**](StringWrapper.md)| The new status for the subscription. Actual options may differ from the indicated set if the invoice status type data has been altered.  Allowable values: (&#39;current&#39;, &#39;canceled&#39;, &#39;stopped&#39;, &#39;payment_failed&#39;, &#39;suspended&#39;) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserSubscriptionPlan**
> SetUserSubscriptionPlan(userId, inventoryId, optional)
Set a new subscription plan for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **planId** | [**StringWrapper**](StringWrapper.md)| The id of the new plan. Must be from the same subscription | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserSubscriptionPrice**
> SetUserSubscriptionPrice(userId, inventoryId, optional)
Set a new subscription price for a user

This new price will be what the user is charged at the begining of each new period. This override is specific to the current subscription and will not carry over if they end and later re-subscribe. It will persist if the plan is changed using the setUserSubscriptionPlan endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **inventoryId** | **int32**| The id of the user&#39;s inventory | 
 **theOverrideDetails** | [**SubscriptionPriceOverrideRequest**](SubscriptionPriceOverrideRequest.md)| override | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

