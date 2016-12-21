# \PaymentsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentMethodUsingPOST**](PaymentsApi.md#CreatePaymentMethodUsingPOST) | **Post** /users/{user_id}/payment-methods | Create a new payment method for a user
[**DeletePaymentMethodUsingDELETE**](PaymentsApi.md#DeletePaymentMethodUsingDELETE) | **Delete** /users/{user_id}/payment-methods/{id} | Delete an existing payment method for a user
[**GetPaymentMethodUsingGET**](PaymentsApi.md#GetPaymentMethodUsingGET) | **Get** /users/{user_id}/payment-methods/{id} | Get a single payment method for a user
[**GetPaymentMethodsUsingGET**](PaymentsApi.md#GetPaymentMethodsUsingGET) | **Get** /users/{user_id}/payment-methods | Get all payment methods for a user
[**PaymentAuthorizationUsingPOST**](PaymentsApi.md#PaymentAuthorizationUsingPOST) | **Post** /payment/authorizations | Authorize payment of an invoice for later capture
[**PaymentCaptureUsingPOST**](PaymentsApi.md#PaymentCaptureUsingPOST) | **Post** /payment/authorizations/{id}/capture | Capture an existing invoice payment authorization
[**UpdatePaymentMethodUsingPUT**](PaymentsApi.md#UpdatePaymentMethodUsingPUT) | **Put** /users/{user_id}/payment-methods/{id} | Update an existing payment method for a user


# **CreatePaymentMethodUsingPOST**
> PaymentMethodResource CreatePaymentMethodUsingPOST($userId, $paymentMethod)

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

# **DeletePaymentMethodUsingDELETE**
> DeletePaymentMethodUsingDELETE($userId, $id)

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

# **GetPaymentMethodUsingGET**
> PaymentMethodResource GetPaymentMethodUsingGET($userId, $id)

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

# **GetPaymentMethodsUsingGET**
> []PaymentMethodResource GetPaymentMethodsUsingGET($userId, $size, $page, $order)

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

# **PaymentAuthorizationUsingPOST**
> PaymentAuthorizationResource PaymentAuthorizationUsingPOST($request)

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

# **PaymentCaptureUsingPOST**
> PaymentCaptureUsingPOST($id)

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

# **UpdatePaymentMethodUsingPUT**
> UpdatePaymentMethodUsingPUT($userId, $id, $paymentMethod)

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

