# \PaymentsPayPalClassicApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingAgreementUrlUsingPOST**](PaymentsPayPalClassicApi.md#CreateBillingAgreementUrlUsingPOST) | **Post** /payment/provider/paypal/classic/agreements/start | Create a PayPal Classic billing agreement for the user
[**ExpressCheckoutUsingPOST**](PaymentsPayPalClassicApi.md#ExpressCheckoutUsingPOST) | **Post** /payment/provider/paypal/classic/checkout/start | Create a payment token for PayPal express checkout
[**FinalizeBillingAgreementUsingPOST**](PaymentsPayPalClassicApi.md#FinalizeBillingAgreementUsingPOST) | **Post** /payment/provider/paypal/classic/agreements/finish | Finalizes a billing agreement after the user has accepted through PayPal
[**FinalizeCheckoutUsingPOST**](PaymentsPayPalClassicApi.md#FinalizeCheckoutUsingPOST) | **Post** /payment/provider/paypal/classic/checkout/finish | Finalizes a payment after the user has completed checkout with PayPal


# **CreateBillingAgreementUrlUsingPOST**
> string CreateBillingAgreementUrlUsingPOST($request)

Create a PayPal Classic billing agreement for the user

Returns the token that should be used to forward the user to PayPal so they can accept the agreement.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateBillingAgreementRequest**](CreateBillingAgreementRequest.md)| The request to create a PayPal billing agreement | [optional] 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpressCheckoutUsingPOST**
> string ExpressCheckoutUsingPOST($request)

Create a payment token for PayPal express checkout

Returns the token that should be used to forward the user to PayPal so they can complete the checkout.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreatePayPalPaymentRequest**](CreatePayPalPaymentRequest.md)| The request to create a PayPal payment token | [optional] 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FinalizeBillingAgreementUsingPOST**
> int32 FinalizeBillingAgreementUsingPOST($request)

Finalizes a billing agreement after the user has accepted through PayPal

Returns the ID of the new payment method created for the user for the billing agreement.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**FinalizeBillingAgreementRequest**](FinalizeBillingAgreementRequest.md)| The request to finalize a PayPal billing agreement | [optional] 

### Return type

**int32**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FinalizeCheckoutUsingPOST**
> FinalizeCheckoutUsingPOST($request)

Finalizes a payment after the user has completed checkout with PayPal

The invoice will be marked paid/failed by asynchronous IPN callback.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**FinalizePayPalPaymentRequest**](FinalizePayPalPaymentRequest.md)| The request to finalize the payment | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

