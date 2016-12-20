# \ReportingRevenueApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecutiveRevenueCountrySalesUsingGET**](ReportingRevenueApi.md#ExecutiveRevenueCountrySalesUsingGET) | **Get** /reporting/revenue/countries/{currency_code} | Get revenue info by country
[**ExecutiveRevenueItemSalesUsingGET**](ReportingRevenueApi.md#ExecutiveRevenueItemSalesUsingGET) | **Get** /reporting/revenue/item-sales/{currency_code} | Get item revenue info
[**ExecutiveRevenueProductSalesUsingGET**](ReportingRevenueApi.md#ExecutiveRevenueProductSalesUsingGET) | **Get** /reporting/revenue/products/{currency_code} | Get revenue info by item
[**ExecutiveRevenueRefundsUsingGET**](ReportingRevenueApi.md#ExecutiveRevenueRefundsUsingGET) | **Get** /reporting/revenue/refunds/{currency_code} | Get refund revenue info
[**ExecutiveRevenueSubscriptionSalesUsingGET**](ReportingRevenueApi.md#ExecutiveRevenueSubscriptionSalesUsingGET) | **Get** /reporting/revenue/subscription-sales/{currency_code} | Get subscription revenue info


# **ExecutiveRevenueCountrySalesUsingGET**
> PageResourceRevenueCountryReportResource ExecutiveRevenueCountrySalesUsingGET($currencyCode, $startDate, $endDate, $size, $page)

Get revenue info by country

Get basic info about revenue from sales of all types, summed up within a time range and split out by country. Sorted for largest revenue at the top


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceRevenueCountryReportResource**](PageResource«RevenueCountryReportResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutiveRevenueItemSalesUsingGET**
> RevenueReportResource ExecutiveRevenueItemSalesUsingGET($currencyCode, $startDate, $endDate)

Get item revenue info

Get basic info about revenue from sales of items and bundles (not subscriptions, shipping, etc), summed up within a time range


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | [optional] 

### Return type

[**RevenueReportResource**](RevenueReportResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutiveRevenueProductSalesUsingGET**
> PageResourceRevenueProductReportResource ExecutiveRevenueProductSalesUsingGET($currencyCode, $startDate, $endDate, $size, $page)

Get revenue info by item

Get basic info about revenue from sales of all types, summed up within a time range and split out by specific item. Sorted for largest revenue at the top


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceRevenueProductReportResource**](PageResource«RevenueProductReportResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutiveRevenueRefundsUsingGET**
> RevenueReportResource ExecutiveRevenueRefundsUsingGET($currencyCode, $startDate, $endDate)

Get refund revenue info

Get basic info about revenue loss from refunds (for all item types), summed up within a time range.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get refund data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | [optional] 

### Return type

[**RevenueReportResource**](RevenueReportResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutiveRevenueSubscriptionSalesUsingGET**
> RevenueReportResource ExecutiveRevenueSubscriptionSalesUsingGET($currencyCode, $startDate, $endDate)

Get subscription revenue info

Get basic info about revenue from sales of new subscriptions as well as recurring payemnts, summed up within a time range


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | [optional] 

### Return type

[**RevenueReportResource**](RevenueReportResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

