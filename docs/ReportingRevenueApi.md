# \ReportingRevenueApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemRevenue**](ReportingRevenueApi.md#GetItemRevenue) | **Get** /reporting/revenue/item-sales/{currency_code} | Get item revenue info
[**GetRefundRevenue**](ReportingRevenueApi.md#GetRefundRevenue) | **Get** /reporting/revenue/refunds/{currency_code} | Get refund revenue info
[**GetRevenueByCountry**](ReportingRevenueApi.md#GetRevenueByCountry) | **Get** /reporting/revenue/countries/{currency_code} | Get revenue info by country
[**GetRevenueByItem**](ReportingRevenueApi.md#GetRevenueByItem) | **Get** /reporting/revenue/products/{currency_code} | Get revenue info by item
[**GetSubscriptionRevenue**](ReportingRevenueApi.md#GetSubscriptionRevenue) | **Get** /reporting/revenue/subscription-sales/{currency_code} | Get subscription revenue info


# **GetItemRevenue**
> RevenueReportResource GetItemRevenue(currencyCode, optional)
Get item revenue info

Get basic info about revenue from sales of items and bundles (not subscriptions, shipping, etc), summed up within a time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **currencyCode** | **string**| The code for a currency to get sales data for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | 

### Return type

[**RevenueReportResource**](RevenueReportResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRefundRevenue**
> RevenueReportResource GetRefundRevenue(currencyCode, optional)
Get refund revenue info

Get basic info about revenue loss from refunds (for all item types), summed up within a time range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **currencyCode** | **string**| The code for a currency to get refund data for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get refund data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | 

### Return type

[**RevenueReportResource**](RevenueReportResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRevenueByCountry**
> PageResourceRevenueCountryReportResource GetRevenueByCountry(currencyCode, optional)
Get revenue info by country

Get basic info about revenue from sales of all types, summed up within a time range and split out by country. Sorted for largest revenue at the top

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **currencyCode** | **string**| The code for a currency to get sales data for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceRevenueCountryReportResource**](PageResource«RevenueCountryReportResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRevenueByItem**
> PageResourceRevenueProductReportResource GetRevenueByItem(currencyCode, optional)
Get revenue info by item

Get basic info about revenue from sales of all types, summed up within a time range and split out by specific item. Sorted for largest revenue at the top

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **currencyCode** | **string**| The code for a currency to get sales data for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceRevenueProductReportResource**](PageResource«RevenueProductReportResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionRevenue**
> RevenueReportResource GetSubscriptionRevenue(currencyCode, optional)
Get subscription revenue info

Get basic info about revenue from sales of new subscriptions as well as recurring payemnts, summed up within a time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **currencyCode** | **string**| The code for a currency to get sales data for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | 

### Return type

[**RevenueReportResource**](RevenueReportResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

