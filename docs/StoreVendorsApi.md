# \StoreVendorsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVendorUsingPOST**](StoreVendorsApi.md#CreateVendorUsingPOST) | **Post** /vendors | Create a vendor
[**DeleteVendorUsingDELETE**](StoreVendorsApi.md#DeleteVendorUsingDELETE) | **Delete** /vendors/{id} | Delete a vendor
[**GetVendorUsingGET**](StoreVendorsApi.md#GetVendorUsingGET) | **Get** /vendors/{id} | Get a single vendor
[**GetVendorsUsingGET**](StoreVendorsApi.md#GetVendorsUsingGET) | **Get** /vendors | List and search vendors
[**UpdateVendorUsingPUT**](StoreVendorsApi.md#UpdateVendorUsingPUT) | **Put** /vendors/{id} | Update a vendor


# **CreateVendorUsingPOST**
> VendorResource CreateVendorUsingPOST($vendor)

Create a vendor


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendor** | [**VendorResource**](VendorResource.md)| The vendor | [optional] 

### Return type

[**VendorResource**](VendorResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVendorUsingDELETE**
> DeleteVendorUsingDELETE($id)

Delete a vendor


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the vendor | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVendorUsingGET**
> VendorResource GetVendorUsingGET($id)

Get a single vendor


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the vendor | 

### Return type

[**VendorResource**](VendorResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVendorsUsingGET**
> PageVendorResource GetVendorsUsingGET($filterName, $size, $page, $order)

List and search vendors


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **string**| Filters vendors by name starting with the text provided in the filter | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageVendorResource**](Page«VendorResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVendorUsingPUT**
> UpdateVendorUsingPUT($id, $vendor)

Update a vendor


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the vendor | 
 **vendor** | [**VendorResource**](VendorResource.md)| The vendor | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

