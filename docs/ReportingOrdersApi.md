# \ReportingOrdersApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyInvoicesUsingGET**](ReportingOrdersApi.md#GetDailyInvoicesUsingGET) | **Get** /reporting/orders/count/{currency_code} | Retrieve invoice counts aggregated by time ranges


# **GetDailyInvoicesUsingGET**
> PageAggregateInvoiceReportResource GetDailyInvoicesUsingGET($currencyCode, $granularity, $filterPaymentStatus, $filterFulfillmentStatus, $startDate, $endDate)

Retrieve invoice counts aggregated by time ranges


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **granularity** | **string**| The time duration to aggregate by | [optional] [default to day]
 **filterPaymentStatus** | **string**| A payment status to filter results by, can be a comma separated list | [optional] 
 **filterFulfillmentStatus** | **string**| An invoice fulfillment status to filter results by, can be a comma separated list | [optional] 
 **startDate** | **int64**| The start of the time range to return, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to return, unix timestamp in seconds. Default is end of time | [optional] 

### Return type

[**PageAggregateInvoiceReportResource**](Page«AggregateInvoiceReportResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

