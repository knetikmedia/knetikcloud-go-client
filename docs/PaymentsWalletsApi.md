# \PaymentsWalletsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWalletTotalsUsingGET**](PaymentsWalletsApi.md#GetWalletTotalsUsingGET) | **Get** /wallets/totals | Retrieves a summation of wallet balances by currency code
[**GetWalletUsingGET**](PaymentsWalletsApi.md#GetWalletUsingGET) | **Get** /users/{user_id}/wallets/{currency_code} | Returns the user&#39;s wallet for the given currency code
[**GetWalletsUsingGET**](PaymentsWalletsApi.md#GetWalletsUsingGET) | **Get** /users/{user_id}/wallets | List all of a user&#39;s wallets
[**GetWalletsUsingGET1**](PaymentsWalletsApi.md#GetWalletsUsingGET1) | **Get** /wallets | Retrieve a list of wallets across the system
[**TransactionHistoryUsingGET**](PaymentsWalletsApi.md#TransactionHistoryUsingGET) | **Get** /wallets/transactions | Retrieve wallet transactions across the system
[**UpdateBalanceUsingPUT**](PaymentsWalletsApi.md#UpdateBalanceUsingPUT) | **Put** /users/{user_id}/wallets/{currency_code}/balance | Updates the balance for a user&#39;s wallet
[**UserTransactionHistoryUsingGET**](PaymentsWalletsApi.md#UserTransactionHistoryUsingGET) | **Get** /users/{user_id}/wallets/{currency_code}/transactions | Retrieve a user&#39;s wallet transactions


# **GetWalletTotalsUsingGET**
> PageResourceWalletTotalResponse GetWalletTotalsUsingGET()

Retrieves a summation of wallet balances by currency code


### Parameters
This endpoint does not need any parameter.

### Return type

[**PageResourceWalletTotalResponse**](PageResource«WalletTotalResponse».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWalletUsingGET**
> SimpleWallet GetWalletUsingGET($userId, $currencyCode)

Returns the user's wallet for the given currency code


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The ID of the user for whom wallet is being retrieved | 
 **currencyCode** | **string**| Currency code of the user&#39;s wallet | 

### Return type

[**SimpleWallet**](SimpleWallet.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWalletsUsingGET**
> []SimpleWallet GetWalletsUsingGET($userId)

List all of a user's wallets


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The ID of the user for whom wallets are being retrieved | 

### Return type

[**[]SimpleWallet**](SimpleWallet.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWalletsUsingGET1**
> PageResourceSimpleWallet GetWalletsUsingGET1($size, $page, $order)

Retrieve a list of wallets across the system


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceSimpleWallet**](PageResource«SimpleWallet».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionHistoryUsingGET**
> PageResourceWalletTransactionResource TransactionHistoryUsingGET($filterInvoice, $filterType, $filterMaxDate, $filterMinDate, $filterSign, $filterUserId, $filterUsername, $filterDetails, $filterCurrencyCode, $size, $page, $order)

Retrieve wallet transactions across the system


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterInvoice** | **int32**| Filter for transactions from a specific invoice | [optional] 
 **filterType** | **string**| Filter for transactions with specified type | [optional] 
 **filterMaxDate** | **int64**| Filter for transactions from no earlier than the specified date as a unix timestamp in seconds | [optional] 
 **filterMinDate** | **int64**| Filter for transactions from no later than the specified date as a unix timestamp in seconds | [optional] 
 **filterSign** | **string**| Filter for transactions with amount with the given sign | [optional] 
 **filterUserId** | **int32**| Filter for transactions for specific userId | [optional] 
 **filterUsername** | **string**| Filter for transactions for specific username that start with the given string | [optional] 
 **filterDetails** | **string**| Filter for transactions for specific details that start with the given string | [optional] 
 **filterCurrencyCode** | **string**| Filter for transactions for specific currency code | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceWalletTransactionResource**](PageResource«WalletTransactionResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBalanceUsingPUT**
> WalletTransactionResource UpdateBalanceUsingPUT($userId, $currencyCode, $request)

Updates the balance for a user's wallet


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The ID of the user for whom wallet is being modified | 
 **currencyCode** | **string**| Currency code of the user&#39;s wallet | 
 **request** | [**WalletAlterRequest**](WalletAlterRequest.md)| The requested balance modification to be made to the user&#39;s wallet | [optional] 

### Return type

[**WalletTransactionResource**](WalletTransactionResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserTransactionHistoryUsingGET**
> PageResourceWalletTransactionResource UserTransactionHistoryUsingGET($userId, $currencyCode, $filterType, $filterMaxDate, $filterMinDate, $filterSign, $size, $page, $order)

Retrieve a user's wallet transactions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The ID of the user for whom wallet transactions are being retrieved | 
 **currencyCode** | **string**| Currency code of the user&#39;s wallet | 
 **filterType** | **string**| Filter for transactions with specified type | [optional] 
 **filterMaxDate** | **int64**| Filter for transactions from no earlier than the specified date as a unix timestamp in seconds | [optional] 
 **filterMinDate** | **int64**| Filter for transactions from no later than the specified date as a unix timestamp in seconds | [optional] 
 **filterSign** | **string**| Filter for transactions with amount with the given sign.  Allowable values: (&#39;positive&#39;, &#39;negative&#39;) | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceWalletTransactionResource**](PageResource«WalletTransactionResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

