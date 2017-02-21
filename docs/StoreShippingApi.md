# \StoreShippingApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShippingItem**](StoreShippingApi.md#CreateShippingItem) | **Post** /store/shipping | Create a shipping item
[**CreateShippingTemplate**](StoreShippingApi.md#CreateShippingTemplate) | **Post** /store/shipping/templates | Create a shipping template
[**DeleteShippingItem**](StoreShippingApi.md#DeleteShippingItem) | **Delete** /store/shipping/{id} | Delete a shipping item
[**DeleteShippingTemplate**](StoreShippingApi.md#DeleteShippingTemplate) | **Delete** /store/shipping/templates/{id} | Delete a shipping template
[**GetShippingItem**](StoreShippingApi.md#GetShippingItem) | **Get** /store/shipping/{id} | Get a single shipping item
[**GetShippingTemplate**](StoreShippingApi.md#GetShippingTemplate) | **Get** /store/shipping/templates/{id} | Get a single shipping template
[**GetShippingTemplates**](StoreShippingApi.md#GetShippingTemplates) | **Get** /store/shipping/templates | List and search shipping templates
[**UpdateShippingItem**](StoreShippingApi.md#UpdateShippingItem) | **Put** /store/shipping/{id} | Update a shipping item
[**UpdateShippingTemplate**](StoreShippingApi.md#UpdateShippingTemplate) | **Put** /store/shipping/templates/{id} | Update a shipping template


# **CreateShippingItem**
> ShippingItem CreateShippingItem($cascade, $shippingItem)

Create a shipping item

A shipping item represents a shipping option and cost. SKUs have to be unique in the entire store.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cascade** | **bool**| Whether to cascade group changes, such as in the limited gettable behavior. A 400 error will return otherwise if the group is already in use with different values. | [optional] [default to false]
 **shippingItem** | [**ShippingItem**](ShippingItem.md)| The shipping item object | [optional] 

### Return type

[**ShippingItem**](ShippingItem.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateShippingTemplate**
> ItemTemplateResource CreateShippingTemplate($shippingTemplateResource)

Create a shipping template

Shipping Templates define a type of shipping and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The new shipping template | [optional] 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteShippingItem**
> DeleteShippingItem($id)

Delete a shipping item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the shipping item | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteShippingTemplate**
> DeleteShippingTemplate($id, $cascade)

Delete a shipping template


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

# **GetShippingItem**
> ShippingItem GetShippingItem($id)

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

# **GetShippingTemplate**
> ItemTemplateResource GetShippingTemplate($id)

Get a single shipping template

Shipping Templates define a type of shipping and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingTemplates**
> PageResourceItemTemplateResource GetShippingTemplates($size, $page, $order)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShippingItem**
> ShippingItem UpdateShippingItem($id, $cascade, $shippingItem)

Update a shipping item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the shipping item | 
 **cascade** | **bool**| Whether to cascade group changes, such as in the limited gettable behavior. A 400 error will return otherwise if the group is already in use with different values. | [optional] [default to false]
 **shippingItem** | [**ShippingItem**](ShippingItem.md)| The shipping item object | [optional] 

### Return type

[**ShippingItem**](ShippingItem.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShippingTemplate**
> ItemTemplateResource UpdateShippingTemplate($id, $shippingTemplateResource)

Update a shipping template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **shippingTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The shipping template resource object | [optional] 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

