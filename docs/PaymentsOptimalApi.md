# \PaymentsOptimalApi

All URIs are relative to *https://devsandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SilentPostUsingPOST**](PaymentsOptimalApi.md#SilentPostUsingPOST) | **Post** /payment/provider/optimal/silent | Initiate silent post with Optimal


# **SilentPostUsingPOST**
> string SilentPostUsingPOST($request)

Initiate silent post with Optimal

Will return the url for a hosted payment endpoint to post to. See Optimal documentation for details.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**OptimalPaymentRequest**](OptimalPaymentRequest.md)| The payment request to initiate | [optional] 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

