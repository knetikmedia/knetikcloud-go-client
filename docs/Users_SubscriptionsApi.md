# \Users_SubscriptionsApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSubscriptionDetails**](Users_SubscriptionsApi.md#GetUserSubscriptionDetails) | **Get** /users/{user_id}/subscriptions/{inventory_id} | Get details about a user&#39;s subscription
[**GetUsersSubscriptionDetails**](Users_SubscriptionsApi.md#GetUsersSubscriptionDetails) | **Get** /users/{user_id}/subscriptions | Get details about a user&#39;s subscriptions
[**ReactivateUserSubscription**](Users_SubscriptionsApi.md#ReactivateUserSubscription) | **Post** /users/{user_id}/subscriptions/{inventory_id}/reactivate | Reactivate a subscription and charge fee
[**SetSubscriptionBillDate**](Users_SubscriptionsApi.md#SetSubscriptionBillDate) | **Put** /users/{user_id}/subscriptions/{inventory_id}/bill-date | Set a new date to bill a subscription on
[**SetSubscriptionPaymentMethod**](Users_SubscriptionsApi.md#SetSubscriptionPaymentMethod) | **Put** /users/{user_id}/subscriptions/{inventory_id}/payment-method | Set the payment method to use for a subscription
[**SetSubscriptionStatus**](Users_SubscriptionsApi.md#SetSubscriptionStatus) | **Put** /users/{user_id}/subscriptions/{inventory_id}/status | Set the status of a subscription
[**SetUserSubscriptionPlan**](Users_SubscriptionsApi.md#SetUserSubscriptionPlan) | **Put** /users/{user_id}/subscriptions/{inventory_id}/plan | Set a new subscription plan for a user
[**SetUserSubscriptionPrice**](Users_SubscriptionsApi.md#SetUserSubscriptionPrice) | **Put** /users/{user_id}/subscriptions/{inventory_id}/price-override | Set a new subscription price for a user


# **GetUserSubscriptionDetails**
> InventorySubscriptionResource GetUserSubscriptionDetails(ctx, ctx, userId, inventoryId)
Get details about a user's subscription

<b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 

### Return type

[**InventorySubscriptionResource**](InventorySubscriptionResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersSubscriptionDetails**
> []InventorySubscriptionResource GetUsersSubscriptionDetails(ctx, ctx, userId)
Get details about a user's subscriptions

<b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| The id of the user | 

### Return type

[**[]InventorySubscriptionResource**](InventorySubscriptionResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactivateUserSubscription**
> InvoiceResource ReactivateUserSubscription(ctx, ctx, userId, inventoryId, optional)
Reactivate a subscription and charge fee

<b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSubscriptionBillDate**
> SetSubscriptionBillDate(ctx, ctx, userId, inventoryId, billDate)
Set a new date to bill a subscription on

<b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
  **billDate** | **int64**| The new bill date. Unix timestamp in seconds | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSubscriptionPaymentMethod**
> SetSubscriptionPaymentMethod(ctx, ctx, userId, inventoryId, optional)
Set the payment method to use for a subscription

May send null to use floating default. <br><br><b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSubscriptionStatus**
> SetSubscriptionStatus(ctx, ctx, userId, inventoryId, status)
Set the status of a subscription

Note that the new status may be blocked if the system is not configured to allow the current status to be changed to the new, to enforce proper flow. The default options for statuses are shown below but may be altered for special use cases. <br><br><b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| The id of the user | 
  **inventoryId** | **int32**| The id of the user&#39;s inventory | 
  **status** | [**SubscriptionStatusWrapper**](SubscriptionStatusWrapper.md)| The new status for the subscription | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserSubscriptionPlan**
> SetUserSubscriptionPlan(ctx, ctx, userId, inventoryId, optional)
Set a new subscription plan for a user

<b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserSubscriptionPrice**
> SetUserSubscriptionPrice(ctx, ctx, userId, inventoryId, optional)
Set a new subscription price for a user

This new price will be what the user is charged at the begining of each new period. This override is specific to the current subscription and will not carry over if they end and later re-subscribe. It will persist if the plan is changed using the setUserSubscriptionPlan endpoint. <br><br><b>Permissions Needed:</b> USERS_SUBSCRIPTIONS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

