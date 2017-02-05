# \StoreShippingApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShippingItemUsingPOST**](StoreShippingApi.md#CreateShippingItemUsingPOST) | **Post** /store/shipping | Create a shipping item
[**CreateShippingTemplateUsingPOST**](StoreShippingApi.md#CreateShippingTemplateUsingPOST) | **Post** /store/shipping/templates | Create a shipping template
[**DeleteShippingItemUsingDELETE**](StoreShippingApi.md#DeleteShippingItemUsingDELETE) | **Delete** /store/shipping/{id} | Delete a shipping item
[**DeleteShippingTemplateUsingDELETE**](StoreShippingApi.md#DeleteShippingTemplateUsingDELETE) | **Delete** /store/shipping/templates/{id} | Delete a shipping template
[**GetShippingItemUsingGET**](StoreShippingApi.md#GetShippingItemUsingGET) | **Get** /store/shipping/{id} | Get a single shipping item
[**GetShippingTemplateUsingGET**](StoreShippingApi.md#GetShippingTemplateUsingGET) | **Get** /store/shipping/templates/{id} | Get a single shipping template
[**GetShippingTemplatesUsingGET**](StoreShippingApi.md#GetShippingTemplatesUsingGET) | **Get** /store/shipping/templates | List and search shipping templates
[**UpdateShippingItemUsingPUT**](StoreShippingApi.md#UpdateShippingItemUsingPUT) | **Put** /store/shipping/{id} | Update a shipping item
[**UpdateShippingTemplateUsingPUT**](StoreShippingApi.md#UpdateShippingTemplateUsingPUT) | **Put** /store/shipping/templates/{id} | Update a shipping template


# **CreateShippingItemUsingPOST**
> ShippingItem CreateShippingItemUsingPOST($shippingItem)

Create a shipping item

A shipping item represents a shipping option and cost. SKUs have to be unique in the entire store.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingItem** | [**ShippingItem**](ShippingItem.md)| The shipping item object | [optional] 

### Return type

[**ShippingItem**](ShippingItem.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateShippingTemplateUsingPOST**
> ItemTemplateResource CreateShippingTemplateUsingPOST($shippingTemplateResource)

Create a shipping template

Shipping Templates define a type of shipping and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The new shipping template | [optional] 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteShippingItemUsingDELETE**
> DeleteShippingItemUsingDELETE($id)

Delete a shipping item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the shipping item | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteShippingTemplateUsingDELETE**
> DeleteShippingTemplateUsingDELETE($id, $cascade)

Delete a shipping template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| force deleting the template if it&#39;s attached to other objects, cascade &#x3D; detach | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingItemUsingGET**
> ShippingItem GetShippingItemUsingGET($id)

Get a single shipping item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the shipping item | 

### Return type

[**ShippingItem**](ShippingItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingTemplateUsingGET**
> ItemTemplateResource GetShippingTemplateUsingGET($id)

Get a single shipping template

Shipping Templates define a type of shipping and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingTemplatesUsingGET**
> PageResourceItemTemplateResource GetShippingTemplatesUsingGET($size, $page, $order)

List and search shipping templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceItemTemplateResource**](PageResource«ItemTemplateResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShippingItemUsingPUT**
> UpdateShippingItemUsingPUT($id, $shippingItem)

Update a shipping item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the shipping item | 
 **shippingItem** | [**ShippingItem**](ShippingItem.md)| The shipping item object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShippingTemplateUsingPUT**
> UpdateShippingTemplateUsingPUT($id, $shippingTemplateResource)

Update a shipping template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **shippingTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The shipping template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

