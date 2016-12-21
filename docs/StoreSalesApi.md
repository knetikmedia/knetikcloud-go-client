# \StoreSalesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCatalogSaleUsingPOST**](StoreSalesApi.md#CreateCatalogSaleUsingPOST) | **Post** /store/sales | Create a sale
[**DeleteCatalogSaleUsingDELETE**](StoreSalesApi.md#DeleteCatalogSaleUsingDELETE) | **Delete** /store/sales/{id} | Delete a sale
[**GetCatalogSaleUsingGET**](StoreSalesApi.md#GetCatalogSaleUsingGET) | **Get** /store/sales/{id} | Get a single sale
[**GetCatalogSalesUsingGET**](StoreSalesApi.md#GetCatalogSalesUsingGET) | **Get** /store/sales | List and search sales
[**UpdateCatalogSaleUsingPUT**](StoreSalesApi.md#UpdateCatalogSaleUsingPUT) | **Put** /store/sales/{id} | Update a sale


# **CreateCatalogSaleUsingPOST**
> CatalogSale CreateCatalogSaleUsingPOST($catalogSale)

Create a sale


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **catalogSale** | [**CatalogSale**](CatalogSale.md)| The catalog sale object | [optional] 

### Return type

[**CatalogSale**](CatalogSale.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCatalogSaleUsingDELETE**
> DeleteCatalogSaleUsingDELETE($id)

Delete a sale


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the sale | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCatalogSaleUsingGET**
> CatalogSale GetCatalogSaleUsingGET($id)

Get a single sale


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the sale | 

### Return type

[**CatalogSale**](CatalogSale.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCatalogSalesUsingGET**
> PageResourceCatalogSale GetCatalogSalesUsingGET($size, $page, $order)

List and search sales


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceCatalogSale**](PageResource«CatalogSale».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCatalogSaleUsingPUT**
> UpdateCatalogSaleUsingPUT($id, $catalogSale)

Update a sale


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the sale | 
 **catalogSale** | [**CatalogSale**](CatalogSale.md)| The catalog sale object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

