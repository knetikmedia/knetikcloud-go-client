# \CurrenciesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCurrencyUsingPOST**](CurrenciesApi.md#CreateCurrencyUsingPOST) | **Post** /currencies | Create a currency
[**DeleteCurrencyUsingDELETE**](CurrenciesApi.md#DeleteCurrencyUsingDELETE) | **Delete** /currencies/{code} | Delete a currency
[**GetCurrenciesUsingGET**](CurrenciesApi.md#GetCurrenciesUsingGET) | **Get** /currencies | List and search currencies
[**GetCurrencyUsingGET**](CurrenciesApi.md#GetCurrencyUsingGET) | **Get** /currencies/{code} | Get a single currency
[**UpdateCurrencyUsingPUT**](CurrenciesApi.md#UpdateCurrencyUsingPUT) | **Put** /currencies/{code} | Update a currency


# **CreateCurrencyUsingPOST**
> CurrencyResource CreateCurrencyUsingPOST($currency)

Create a currency


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**CurrencyResource**](CurrencyResource.md)| The currency object | [optional] 

### Return type

[**CurrencyResource**](CurrencyResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCurrencyUsingDELETE**
> DeleteCurrencyUsingDELETE($code)

Delete a currency


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string**| The currency code | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrenciesUsingGET**
> PageCurrencyResource GetCurrenciesUsingGET($filterEnabledCurrencies, $filterType, $size, $page, $order)

List and search currencies


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEnabledCurrencies** | **bool**| Filter for alternate currencies setup explicitely in system config | [optional] 
 **filterType** | **string**| Filter currencies by type.  Allowable values: (&#39;virtual&#39;, &#39;real&#39;) | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageCurrencyResource**](Page«CurrencyResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencyUsingGET**
> CurrencyResource GetCurrencyUsingGET($code)

Get a single currency


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string**| The currency code | 

### Return type

[**CurrencyResource**](CurrencyResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCurrencyUsingPUT**
> UpdateCurrencyUsingPUT($code, $currency)

Update a currency


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string**| The currency code | 
 **currency** | [**CurrencyResource**](CurrencyResource.md)| The currency object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

