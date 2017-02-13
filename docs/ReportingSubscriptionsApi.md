# \ReportingSubscriptionsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptionReports**](ReportingSubscriptionsApi.md#GetSubscriptionReports) | **Get** /reporting/subscription | Get a list of available subscription reports in most recent first order


# **GetSubscriptionReports**
> PageResourceBillingReport GetSubscriptionReports($size, $page)

Get a list of available subscription reports in most recent first order


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceBillingReport**](PageResource«BillingReport».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

