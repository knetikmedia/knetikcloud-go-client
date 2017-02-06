# \PaymentsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentMethod**](PaymentsApi.md#CreatePaymentMethod) | **Post** /users/{user_id}/payment-methods | Create a new payment method for a user
[**DeletePaymentMethod**](PaymentsApi.md#DeletePaymentMethod) | **Delete** /users/{user_id}/payment-methods/{id} | Delete an existing payment method for a user
[**GetPaymentMethod**](PaymentsApi.md#GetPaymentMethod) | **Get** /users/{user_id}/payment-methods/{id} | Get a single payment method for a user
[**GetPaymentMethods**](PaymentsApi.md#GetPaymentMethods) | **Get** /users/{user_id}/payment-methods | Get all payment methods for a user
[**PaymentAuthorization**](PaymentsApi.md#PaymentAuthorization) | **Post** /payment/authorizations | Authorize payment of an invoice for later capture
[**PaymentCapture**](PaymentsApi.md#PaymentCapture) | **Post** /payment/authorizations/{id}/capture | Capture an existing invoice payment authorization
[**UpdatePaymentMethod**](PaymentsApi.md#UpdatePaymentMethod) | **Put** /users/{user_id}/payment-methods/{id} | Update an existing payment method for a user


# **CreatePaymentMethod**
> PaymentMethodResource CreatePaymentMethod($userId, $paymentMethod)

Create a new payment method for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment method is being created | 
 **paymentMethod** | [**PaymentMethodResource**](PaymentMethodResource.md)| Payment method being created | [optional] 

### Return type

[**PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePaymentMethod**
> DeletePaymentMethod($userId, $id)

Delete an existing payment method for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment method is being updated | 
 **id** | **int32**| ID of the payment method being deleted | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethod**
> PaymentMethodResource GetPaymentMethod($userId, $id)

Get a single payment method for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment method is being retrieved | 
 **id** | **int32**| ID of the payment method being retrieved | 

### Return type

[**PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethods**
> []PaymentMethodResource GetPaymentMethods($userId, $size, $page, $order)

Get all payment methods for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment methods are being retrieved | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| a comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**[]PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentAuthorization**
> PaymentAuthorizationResource PaymentAuthorization($request)

Authorize payment of an invoice for later capture


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PaymentAuthorizationResource**](PaymentAuthorizationResource.md)| Payment authorization request | [optional] 

### Return type

[**PaymentAuthorizationResource**](PaymentAuthorizationResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentCapture**
> PaymentCapture($id)

Capture an existing invoice payment authorization


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| ID of the payment authorization to capture | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePaymentMethod**
> UpdatePaymentMethod($userId, $id, $paymentMethod)

Update an existing payment method for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| ID of the user for whom the payment method is being updated | 
 **id** | **int32**| ID of the payment method being updated | 
 **paymentMethod** | [**PaymentMethodResource**](PaymentMethodResource.md)| The updated payment method data | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

