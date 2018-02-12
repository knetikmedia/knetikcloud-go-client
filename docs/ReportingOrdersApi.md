# \ReportingOrdersApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoiceReports**](ReportingOrdersApi.md#GetInvoiceReports) | **Get** /reporting/orders/count/{currency_code} | Retrieve invoice counts aggregated by time ranges


# **GetInvoiceReports**
> PageResourceAggregateInvoiceReportResource GetInvoiceReports(ctx, ctx, currencyCode, optional)
Retrieve invoice counts aggregated by time ranges

<b>Permissions Needed:</b> REPORTING_ORDERS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **currencyCode** | **string**| The code for a currency to get sales data for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyCode** | **string**| The code for a currency to get sales data for | 
 **granularity** | **string**| The time duration to aggregate by | [default to day]
 **filterPaymentStatus** | **string**| A payment status to filter results by, can be a comma separated list | 
 **filterFulfillmentStatus** | **string**| An invoice fulfillment status to filter results by, can be a comma separated list | 
 **startDate** | **int64**| The start of the time range to return, unix timestamp in seconds. Default is beginning of time | 
 **endDate** | **int64**| The end of the time range to return, unix timestamp in seconds. Default is end of time | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned | [default to 1]

### Return type

[**PageResourceAggregateInvoiceReportResource**](PageResource«AggregateInvoiceReportResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

