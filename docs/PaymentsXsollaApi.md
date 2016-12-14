# \PaymentsXsollaApi

All URIs are relative to *https://devsandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenUrlUsingPOST**](PaymentsXsollaApi.md#CreateTokenUrlUsingPOST) | **Post** /payment/provider/xsolla/payment | Create a payment token that should be used to forward the user to Xsolla so they can complete payment
[**ReceiveNotificationUsingPOST**](PaymentsXsollaApi.md#ReceiveNotificationUsingPOST) | **Post** /payment/provider/xsolla/notifications | Receives payment response from Xsolla


# **CreateTokenUrlUsingPOST**
> string CreateTokenUrlUsingPOST($request)

Create a payment token that should be used to forward the user to Xsolla so they can complete payment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**XsollaPaymentRequest**](XsollaPaymentRequest.md)| The payment request to be sent to XSolla | [optional] 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReceiveNotificationUsingPOST**
> ReceiveNotificationUsingPOST()

Receives payment response from Xsolla

Only used by XSolla to call back to JSAPI after processing user payment action


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

