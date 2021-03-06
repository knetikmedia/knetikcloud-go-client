# \PaymentsApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentMethod**](PaymentsApi.md#CreatePaymentMethod) | **Post** /users/{user_id}/payment-methods | Create a new payment method for a user
[**DeletePaymentMethod**](PaymentsApi.md#DeletePaymentMethod) | **Delete** /users/{user_id}/payment-methods/{id} | Delete an existing payment method for a user
[**GetPaymentMethod**](PaymentsApi.md#GetPaymentMethod) | **Get** /users/{user_id}/payment-methods/{id} | Get a single payment method for a user
[**GetPaymentMethodType**](PaymentsApi.md#GetPaymentMethodType) | **Get** /payment/types/{id} | Get a single payment method type
[**GetPaymentMethodTypes**](PaymentsApi.md#GetPaymentMethodTypes) | **Get** /payment/types | Get all payment method types
[**GetPaymentMethods**](PaymentsApi.md#GetPaymentMethods) | **Get** /users/{user_id}/payment-methods | Get all payment methods for a user
[**PaymentAuthorization**](PaymentsApi.md#PaymentAuthorization) | **Post** /payment/authorizations | Authorize payment of an invoice for later capture
[**PaymentCapture**](PaymentsApi.md#PaymentCapture) | **Post** /payment/authorizations/{id}/capture | Capture an existing invoice payment authorization
[**UpdatePaymentMethod**](PaymentsApi.md#UpdatePaymentMethod) | **Put** /users/{user_id}/payment-methods/{id} | Update an existing payment method for a user


# **CreatePaymentMethod**
> PaymentMethodResource CreatePaymentMethod(ctx, ctx, userId, optional)
Create a new payment method for a user

<b>Permissions Needed:</b> PAYMENTS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| ID of the user for whom the payment method is being created | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment method is being created | 
 **paymentMethod** | [**PaymentMethodResource**](PaymentMethodResource.md)| Payment method being created | 

### Return type

[**PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePaymentMethod**
> DeletePaymentMethod(ctx, ctx, userId, id)
Delete an existing payment method for a user

<b>Permissions Needed:</b> PAYMENTS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| ID of the user for whom the payment method is being updated | 
  **id** | **int32**| ID of the payment method being deleted | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethod**
> PaymentMethodResource GetPaymentMethod(ctx, ctx, userId, id)
Get a single payment method for a user

<b>Permissions Needed:</b> PAYMENTS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| ID of the user for whom the payment method is being retrieved | 
  **id** | **int32**| ID of the payment method being retrieved | 

### Return type

[**PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethodType**
> PaymentMethodTypeResource GetPaymentMethodType(ctx, ctx, id)
Get a single payment method type

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int32**| ID of the payment method type being retrieved | 

### Return type

[**PaymentMethodTypeResource**](PaymentMethodTypeResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethodTypes**
> PageResourcePaymentMethodTypeResource GetPaymentMethodTypes(ctx, ctx, optional)
Get all payment method types

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **string**| Filter for payment method types whose name matches a given string | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| a comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourcePaymentMethodTypeResource**](PageResource«PaymentMethodTypeResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethods**
> []PaymentMethodResource GetPaymentMethods(ctx, ctx, userId, optional)
Get all payment methods for a user

<b>Permissions Needed:</b> PAYMENTS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| ID of the user for whom the payment methods are being retrieved | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment methods are being retrieved | 
 **filterName** | **string**| Filter for payment methods whose name starts with a given string | 
 **filterPaymentType** | **string**| Filter for payment methods with a specific payment type | 
 **filterPaymentMethodTypeId** | **int32**| Filter for payment methods with a specific payment method type by id | 
 **filterPaymentMethodTypeName** | **string**| Filter for payment methods whose payment method type name starts with a given string | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| a comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**[]PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentAuthorization**
> PaymentAuthorizationResource PaymentAuthorization(ctx, ctx, optional)
Authorize payment of an invoice for later capture

<b>Permissions Needed:</b> PAYMENTS_ADMIN or PAYMENTS_USER

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PaymentAuthorizationResource**](PaymentAuthorizationResource.md)| Payment authorization request | 

### Return type

[**PaymentAuthorizationResource**](PaymentAuthorizationResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentCapture**
> PaymentCapture(ctx, ctx, id)
Capture an existing invoice payment authorization

<b>Permissions Needed:</b> PAYMENTS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int32**| ID of the payment authorization to capture | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePaymentMethod**
> PaymentMethodResource UpdatePaymentMethod(ctx, ctx, userId, id, optional)
Update an existing payment method for a user

<b>Permissions Needed:</b> PAYMENTS_ADMIN or owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| ID of the user for whom the payment method is being updated | 
  **id** | **int32**| ID of the payment method being updated | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment method is being updated | 
 **id** | **int32**| ID of the payment method being updated | 
 **paymentMethod** | [**PaymentMethodResource**](PaymentMethodResource.md)| The updated payment method data | 

### Return type

[**PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

