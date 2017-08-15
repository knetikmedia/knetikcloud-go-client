# \PaymentsWalletsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserWallet**](PaymentsWalletsApi.md#GetUserWallet) | **Get** /users/{user_id}/wallets/{currency_code} | Returns the user&#39;s wallet for the given currency code
[**GetUserWalletTransactions**](PaymentsWalletsApi.md#GetUserWalletTransactions) | **Get** /users/{user_id}/wallets/{currency_code}/transactions | Retrieve a user&#39;s wallet transactions
[**GetUserWallets**](PaymentsWalletsApi.md#GetUserWallets) | **Get** /users/{user_id}/wallets | List all of a user&#39;s wallets
[**GetWalletBalances**](PaymentsWalletsApi.md#GetWalletBalances) | **Get** /wallets/totals | Retrieves a summation of wallet balances by currency code
[**GetWalletTransactions**](PaymentsWalletsApi.md#GetWalletTransactions) | **Get** /wallets/transactions | Retrieve wallet transactions across the system
[**GetWallets**](PaymentsWalletsApi.md#GetWallets) | **Get** /wallets | Retrieve a list of wallets across the system
[**UpdateWalletBalance**](PaymentsWalletsApi.md#UpdateWalletBalance) | **Put** /users/{user_id}/wallets/{currency_code}/balance | Updates the balance for a user&#39;s wallet


# **GetUserWallet**
> SimpleWallet GetUserWallet(userId, currencyCode)
Returns the user's wallet for the given currency code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The ID of the user for whom wallet is being retrieved | 
  **currencyCode** | **string**| Currency code of the user&#39;s wallet | 

### Return type

[**SimpleWallet**](SimpleWallet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletTransactions**
> PageResourceWalletTransactionResource GetUserWalletTransactions(userId, currencyCode, optional)
Retrieve a user's wallet transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The ID of the user for whom wallet transactions are being retrieved | 
  **currencyCode** | **string**| Currency code of the user&#39;s wallet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The ID of the user for whom wallet transactions are being retrieved | 
 **currencyCode** | **string**| Currency code of the user&#39;s wallet | 
 **filterType** | **string**| Filter for transactions with specified type | 
 **filterMaxDate** | **int64**| Filter for transactions from no earlier than the specified date as a unix timestamp in seconds | 
 **filterMinDate** | **int64**| Filter for transactions from no later than the specified date as a unix timestamp in seconds | 
 **filterSign** | **string**| Filter for transactions with amount with the given sign.  Allowable values: (&#39;positive&#39;, &#39;negative&#39;) | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceWalletTransactionResource**](PageResource«WalletTransactionResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWallets**
> []SimpleWallet GetUserWallets(userId)
List all of a user's wallets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The ID of the user for whom wallets are being retrieved | 

### Return type

[**[]SimpleWallet**](SimpleWallet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWalletBalances**
> PageResourceWalletTotalResponse GetWalletBalances()
Retrieves a summation of wallet balances by currency code

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PageResourceWalletTotalResponse**](PageResource«WalletTotalResponse».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWalletTransactions**
> PageResourceWalletTransactionResource GetWalletTransactions(optional)
Retrieve wallet transactions across the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterInvoice** | **int32**| Filter for transactions from a specific invoice | 
 **filterType** | **string**| Filter for transactions with specified type | 
 **filterDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the log start date, a unix timestamp in seconds. Can be repeated for a range, eg: GT,123,LT,456  Allowed operators: (GT, LT, EQ, GOE, LOE). | 
 **filterSign** | **string**| Filter for transactions with amount with the given sign | 
 **filterUserId** | **int32**| Filter for transactions for specific userId | 
 **filterUsername** | **string**| Filter for transactions for specific username that start with the given string | 
 **filterDetails** | **string**| Filter for transactions for specific details that start with the given string | 
 **filterCurrencyCode** | **string**| Filter for transactions for specific currency code | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceWalletTransactionResource**](PageResource«WalletTransactionResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWallets**
> PageResourceSimpleWallet GetWallets(optional)
Retrieve a list of wallets across the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceSimpleWallet**](PageResource«SimpleWallet».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWalletBalance**
> WalletTransactionResource UpdateWalletBalance(userId, currencyCode, optional)
Updates the balance for a user's wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The ID of the user for whom wallet is being modified | 
  **currencyCode** | **string**| Currency code of the user&#39;s wallet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The ID of the user for whom wallet is being modified | 
 **currencyCode** | **string**| Currency code of the user&#39;s wallet | 
 **request** | [**WalletAlterRequest**](WalletAlterRequest.md)| The requested balance modification to be made to the user&#39;s wallet | 

### Return type

[**WalletTransactionResource**](WalletTransactionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

