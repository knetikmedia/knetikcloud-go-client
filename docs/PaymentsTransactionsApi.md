# \PaymentsTransactionsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionUsingGET**](PaymentsTransactionsApi.md#GetTransactionUsingGET) | **Get** /transactions/{id} | Get the details for a single transaction
[**GetTransactionsUsingGET**](PaymentsTransactionsApi.md#GetTransactionsUsingGET) | **Get** /transactions | List and search transactions
[**RefundTransactionUsingPOST**](PaymentsTransactionsApi.md#RefundTransactionUsingPOST) | **Post** /transactions/{id}/refunds | Refund a payment transaction, in full or in part


# **GetTransactionUsingGET**
> TransactionResource GetTransactionUsingGET($id)

Get the details for a single transaction


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| id | 

### Return type

[**TransactionResource**](TransactionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsUsingGET**
> PageResourceTransactionResource GetTransactionsUsingGET($filterInvoice, $size, $page, $order)

List and search transactions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterInvoice** | **int32**| Filter for transactions from a specific invoice | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceTransactionResource**](PageResource«TransactionResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundTransactionUsingPOST**
> RefundResource RefundTransactionUsingPOST($id, $request)

Refund a payment transaction, in full or in part

Will not allow for refunding more than the full amount even with multiple partial refunds. Money is refunded to the payment method used to make the original payment. Payment method must support refunds.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the transaction to refund | 
 **request** | [**RefundRequest**](RefundRequest.md)| Request containing refund details | [optional] 

### Return type

[**RefundResource**](RefundResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
