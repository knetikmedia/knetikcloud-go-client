# \FulfillmentApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFulfillmentTypeUsingPOST**](FulfillmentApi.md#CreateFulfillmentTypeUsingPOST) | **Post** /store/fulfillment/types | Create a fulfillment type
[**DeleteFulfillmentTypeUsingDELETE**](FulfillmentApi.md#DeleteFulfillmentTypeUsingDELETE) | **Delete** /store/fulfillment/types/{id} | Delete a fulfillment type
[**GetFulfillmentTypeUsingGET**](FulfillmentApi.md#GetFulfillmentTypeUsingGET) | **Get** /store/fulfillment/types/{id} | Get a single fulfillment type
[**GetFulfillmentsUsingGET**](FulfillmentApi.md#GetFulfillmentsUsingGET) | **Get** /store/fulfillment/types | List and search fulfillment types
[**UpdateFulfillmentTypeUsingPUT**](FulfillmentApi.md#UpdateFulfillmentTypeUsingPUT) | **Put** /store/fulfillment/types/{id} | Update a fulfillment type


# **CreateFulfillmentTypeUsingPOST**
> FulfillmentType CreateFulfillmentTypeUsingPOST($type_)

Create a fulfillment type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**FulfillmentType**](FulfillmentType.md)| The fulfillment type | [optional] 

### Return type

[**FulfillmentType**](FulfillmentType.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFulfillmentTypeUsingDELETE**
> DeleteFulfillmentTypeUsingDELETE($id)

Delete a fulfillment type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFulfillmentTypeUsingGET**
> FulfillmentType GetFulfillmentTypeUsingGET($id)

Get a single fulfillment type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id | 

### Return type

[**FulfillmentType**](FulfillmentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFulfillmentsUsingGET**
> PageFulfillmentType GetFulfillmentsUsingGET($size, $page, $order)

List and search fulfillment types


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to 1]

### Return type

[**PageFulfillmentType**](Page«FulfillmentType».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFulfillmentTypeUsingPUT**
> UpdateFulfillmentTypeUsingPUT($id, $fulfillmentType)

Update a fulfillment type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id | 
 **fulfillmentType** | [**FulfillmentType**](FulfillmentType.md)| The fulfillment type | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

