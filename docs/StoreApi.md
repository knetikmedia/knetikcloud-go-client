# \StoreApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemTemplateUsingPOST**](StoreApi.md#CreateItemTemplateUsingPOST) | **Post** /store/items/templates | Create an item template
[**CreateStoreItemExternalUsingPOST**](StoreApi.md#CreateStoreItemExternalUsingPOST) | **Post** /store/items | Create a store item
[**DeleteItemTemplateUsingDELETE**](StoreApi.md#DeleteItemTemplateUsingDELETE) | **Delete** /store/items/templates/{id} | Delete an item template
[**DeleteStoreItemUsingDELETE1**](StoreApi.md#DeleteStoreItemUsingDELETE1) | **Delete** /store/items/{id} | Delete a store item
[**GetItemTemplateUsingGET**](StoreApi.md#GetItemTemplateUsingGET) | **Get** /store/items/templates/{id} | Get a single item template
[**GetItemTemplatesUsingGET**](StoreApi.md#GetItemTemplatesUsingGET) | **Get** /store/items/templates | List and search item templates
[**GetStoreItemUsingGET1**](StoreApi.md#GetStoreItemUsingGET1) | **Get** /store/items/{id} | Get a single store item
[**GetStoreItemsUsingGET**](StoreApi.md#GetStoreItemsUsingGET) | **Get** /store/items | List and search store items
[**GetUsingGET2**](StoreApi.md#GetUsingGET2) | **Get** /store | Get a listing of store items
[**UpdateItemTemplateUsingPUT**](StoreApi.md#UpdateItemTemplateUsingPUT) | **Put** /store/items/templates/{id} | Update an item template
[**UpdateStoreItemExternalUsingPUT**](StoreApi.md#UpdateStoreItemExternalUsingPUT) | **Put** /store/items/{id} | Update a store item


# **CreateItemTemplateUsingPOST**
> StoreItemTemplateResource CreateItemTemplateUsingPOST($itemTemplateResource)

Create an item template

Item Templates define a type of item and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemTemplateResource** | [**StoreItemTemplateResource**](StoreItemTemplateResource.md)| The new item template | [optional] 

### Return type

[**StoreItemTemplateResource**](StoreItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoreItemExternalUsingPOST**
> StoreItem CreateStoreItemExternalUsingPOST($storeItem)

Create a store item

SKUs have to be unique in the entire store. If a duplicate SKU is found, a 400 error is generated and the response will have a \"parameters\" field that is a list of duplicates. A duplicate is an object like {item_id, offending_sku_list}. Ex:<br /> {..., parameters: [[{item: 1, skus: [\"SKU-1\"]}]]}<br /> If an item is brand new and has duplicate SKUs within itself, the item ID will be 0.  Item subclasses are not allowed here, you will have to use their respective endpoints.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeItem** | [**StoreItem**](StoreItem.md)| The store item object | [optional] 

### Return type

[**StoreItem**](StoreItem.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteItemTemplateUsingDELETE**
> DeleteItemTemplateUsingDELETE($id, $cascade)

Delete an item template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| force deleting the template if it&#39;s attached to other objects, cascade &#x3D; detach | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoreItemUsingDELETE1**
> DeleteStoreItemUsingDELETE1($id)

Delete a store item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the item | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItemTemplateUsingGET**
> StoreItemTemplateResource GetItemTemplateUsingGET($id)

Get a single item template

Item Templates define a type of item and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**StoreItemTemplateResource**](StoreItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItemTemplatesUsingGET**
> PageResourceStoreItemTemplateResource GetItemTemplatesUsingGET($size, $page, $order)

List and search item templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceStoreItemTemplateResource**](PageResource«StoreItemTemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoreItemUsingGET1**
> StoreItem GetStoreItemUsingGET1($id)

Get a single store item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the item | 

### Return type

[**StoreItem**](StoreItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoreItemsUsingGET**
> PageResourceStoreItem GetStoreItemsUsingGET($filterNameSearch, $filterUniqueKey, $filterPublished, $filterDisplayable, $filterStart, $filterEnd, $filterStartDate, $filterStopDate, $filterSku, $filterPrice, $filterTag, $filterItemsByType, $filterBundledSkus, $size, $page, $order)

List and search store items


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterNameSearch** | **string**| Filter for items whose name starts with a given string. | [optional] 
 **filterUniqueKey** | **string**| Filter for items whose unique_key is a given string. | [optional] 
 **filterPublished** | **bool**| Filter for skus that have been published. | [optional] 
 **filterDisplayable** | **bool**| Filter for items that are displayable. | [optional] 
 **filterStart** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the store start date, a unix timestamp in seconds.  Allowed operators: (LT, GT, LTE, GTE, EQ). | [optional] 
 **filterEnd** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the store end date, a unix timestamp in seconds.  Allowed operators: (LT, GT, LTE, GTE, EQ). | [optional] 
 **filterStartDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the sku start date, a unix timestamp in seconds.  Allowed operators: (LT, GT, LTE, GTE, EQ). | [optional] 
 **filterStopDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the sku end date, a unix timestamp in seconds.  Allowed operators: (LT, GT, LTE, GTE, EQ). | [optional] 
 **filterSku** | **string**| Filter for skus whose name starts with a given string. | [optional] 
 **filterPrice** | **string**| A colon separated string without spaces.  First value is the operator to search on, second value is the price of a sku.  Allowed operators: (LT, GT, LTE, GTE, EQ). | [optional] 
 **filterTag** | **string**| A comma separated list without spaces of the names of tags. Will only return items with at least one of the tags. | [optional] 
 **filterItemsByType** | **string**| Filter for item type based on its type hint. | [optional] 
 **filterBundledSkus** | **string**| Filter for skus inside bundles whose name starts with a given string.  Used only when type hint is &#39;bundle_item&#39; | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceStoreItem**](PageResource«StoreItem».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET2**
> PageResourceStoreItem GetUsingGET2($storeListRequest)

Get a listing of store items

The exact structure of each items may differ to include fields specific to the type. The same is true for behaviors.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeListRequest** | [**StoreListRequest**](StoreListRequest.md)| The store list request object | [optional] 

### Return type

[**PageResourceStoreItem**](PageResource«StoreItem».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItemTemplateUsingPUT**
> UpdateItemTemplateUsingPUT($id, $itemTemplateResource)

Update an item template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **itemTemplateResource** | [**StoreItemTemplateResource**](StoreItemTemplateResource.md)| The item template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStoreItemExternalUsingPUT**
> UpdateStoreItemExternalUsingPUT($id, $storeItem)

Update a store item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the item | 
 **storeItem** | [**StoreItem**](StoreItem.md)| The store item object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
