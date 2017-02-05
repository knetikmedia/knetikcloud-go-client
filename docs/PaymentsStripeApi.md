# \PaymentsStripeApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerUsingPOST**](PaymentsStripeApi.md#CreateCustomerUsingPOST) | **Post** /payment/provider/stripe/payment-methods | Create a Stripe payment method for a user
[**PayInvoiceUsingPOST1**](PaymentsStripeApi.md#PayInvoiceUsingPOST1) | **Post** /payment/provider/stripe/payments | Pay with a single use token


# **CreateCustomerUsingPOST**
> PaymentMethodResource CreateCustomerUsingPOST($request)

Create a Stripe payment method for a user

Stores customer information and creates a payment method that can be used to pay invoices through the payments endpoints.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**StripeCreatePaymentMethod**](StripeCreatePaymentMethod.md)| The request to create a Stripe customer with payment info | [optional] 

### Return type

[**PaymentMethodResource**](PaymentMethodResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayInvoiceUsingPOST1**
> PayInvoiceUsingPOST1($request)

Pay with a single use token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**StripePaymentRequest**](StripePaymentRequest.md)| The request to pay an invoice | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

