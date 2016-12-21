# \StoreBundlesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundleItemUsingPOST**](StoreBundlesApi.md#CreateBundleItemUsingPOST) | **Post** /store/bundles | Create a bundle item
[**CreateBundleTemplateUsingPOST**](StoreBundlesApi.md#CreateBundleTemplateUsingPOST) | **Post** /store/bundles/templates | Create a bundle template
[**DeleteBundleTemplateUsingDELETE**](StoreBundlesApi.md#DeleteBundleTemplateUsingDELETE) | **Delete** /store/bundles/templates/{id} | Delete a bundle template
[**DeleteStoreItemUsingDELETE**](StoreBundlesApi.md#DeleteStoreItemUsingDELETE) | **Delete** /store/bundles/{id} | Delete a bundle item
[**GetBundleTemplateUsingGET**](StoreBundlesApi.md#GetBundleTemplateUsingGET) | **Get** /store/bundles/templates/{id} | Get a single bundle template
[**GetBundleTemplatesUsingGET**](StoreBundlesApi.md#GetBundleTemplatesUsingGET) | **Get** /store/bundles/templates | List and search bundle templates
[**GetStoreItemUsingGET**](StoreBundlesApi.md#GetStoreItemUsingGET) | **Get** /store/bundles/{id} | Get a single bundle item
[**UpdateBundleItemUsingPUT**](StoreBundlesApi.md#UpdateBundleItemUsingPUT) | **Put** /store/bundles/{id} | Update a bundle item
[**UpdateBundleTemplateUsingPUT**](StoreBundlesApi.md#UpdateBundleTemplateUsingPUT) | **Put** /store/bundles/templates/{id} | Update a bundle template


# **CreateBundleItemUsingPOST**
> BundleItem CreateBundleItemUsingPOST($bundleItem)

Create a bundle item

The SKU for the bundle itself must be unique and there can only be one SKU.  Extra notes for price_override:  The price of all the items (multiplied by the quantity) must equal the price of the bundle.  With individual prices set, items will be processed individually and can be refunded as such.  However, if all prices are set to null, the price of the bundle will be used and will be treated as one item.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleItem** | [**BundleItem**](BundleItem.md)| The bundle item object | [optional] 

### Return type

[**BundleItem**](BundleItem.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBundleTemplateUsingPOST**
> ItemTemplateResource CreateBundleTemplateUsingPOST($bundleTemplateResource)

Create a bundle template

Bundle Templates define a type of bundle and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The new bundle template | [optional] 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBundleTemplateUsingDELETE**
> DeleteBundleTemplateUsingDELETE($id, $cascade)

Delete a bundle template


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

# **DeleteStoreItemUsingDELETE**
> DeleteStoreItemUsingDELETE($id)

Delete a bundle item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the bundle | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleTemplateUsingGET**
> ItemTemplateResource GetBundleTemplateUsingGET($id)

Get a single bundle template

Bundle Templates define a type of bundle and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleTemplatesUsingGET**
> PageResourceItemTemplateResource GetBundleTemplatesUsingGET($size, $page, $order)

List and search bundle templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceItemTemplateResource**](PageResource«ItemTemplateResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoreItemUsingGET**
> BundleItem GetStoreItemUsingGET($id)

Get a single bundle item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the bundle | 

### Return type

[**BundleItem**](BundleItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBundleItemUsingPUT**
> UpdateBundleItemUsingPUT($id, $bundleItem)

Update a bundle item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the bundle | 
 **bundleItem** | [**BundleItem**](BundleItem.md)| The bundle item object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBundleTemplateUsingPUT**
> UpdateBundleTemplateUsingPUT($id, $bundleTemplateResource)

Update a bundle template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **bundleTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The bundle template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

