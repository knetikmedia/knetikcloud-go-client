# \ReportingUsersApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecutiveRevenueItemSalesUsingGET1**](ReportingUsersApi.md#ExecutiveRevenueItemSalesUsingGET1) | **Get** /reporting/users/registrations | Get user registration info


# **ExecutiveRevenueItemSalesUsingGET1**
> PageResourceAggregateCountResource ExecutiveRevenueItemSalesUsingGET1($granularity, $startDate, $endDate)

Get user registration info

Get user registration counts grouped by time range


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **granularity** | **string**| The time duration to aggregate by | [optional] [default to day]
 **startDate** | **int64**| The start of the time range to aggregate, unix timestamp in seconds. Default is beginning of time | [optional] 
 **endDate** | **int64**| The end of the time range to aggregate, unix timestamp in seconds. Default is end of time | [optional] 

### Return type

[**PageResourceAggregateCountResource**](PageResource«AggregateCountResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

