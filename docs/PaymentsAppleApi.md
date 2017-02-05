# \PaymentsAppleApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyReceiptUsingPOST**](PaymentsAppleApi.md#VerifyReceiptUsingPOST) | **Post** /payment/provider/apple/receipt | Pay invoice with Apple receipt


# **VerifyReceiptUsingPOST**
> string VerifyReceiptUsingPOST($request)

Pay invoice with Apple receipt

Mark an invoice paid using Apple payment receipt. A receipt will only be accepted once and the details of the transaction must match the invoice, including the product_id matching the sku text of the item in the invoice. Returns the transaction ID if successful.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApplyPaymentRequest**](ApplyPaymentRequest.md)| The request for paying an invoice through an Apple receipt | [optional] 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

