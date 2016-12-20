# \BRERuleEngineCategoriesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateUsingPOST1**](BRERuleEngineCategoriesApi.md#CreateTemplateUsingPOST1) | **Post** /bre/categories/templates | Create a BRE category template
[**DeleteTemplateUsingDELETE**](BRERuleEngineCategoriesApi.md#DeleteTemplateUsingDELETE) | **Delete** /bre/categories/templates/{id} | Delete a BRE category template
[**GetCategoriesUsingGET**](BRERuleEngineCategoriesApi.md#GetCategoriesUsingGET) | **Get** /bre/categories | List categories
[**GetCategoryUsingGET**](BRERuleEngineCategoriesApi.md#GetCategoryUsingGET) | **Get** /bre/categories/{name} | Get a single category
[**GetTemplateUsingGET**](BRERuleEngineCategoriesApi.md#GetTemplateUsingGET) | **Get** /bre/categories/templates/{id} | Get a single BRE category template
[**GetTemplatesUsingGET**](BRERuleEngineCategoriesApi.md#GetTemplatesUsingGET) | **Get** /bre/categories/templates | List and search BRE category templates
[**UpdateCategoryUsingPUT**](BRERuleEngineCategoriesApi.md#UpdateCategoryUsingPUT) | **Put** /bre/categories/{name} | Update a category
[**UpdateTemplateUsingPUT1**](BRERuleEngineCategoriesApi.md#UpdateTemplateUsingPUT1) | **Put** /bre/categories/templates/{id} | Update a BRE category template


# **CreateTemplateUsingPOST1**
> TemplateResource CreateTemplateUsingPOST1($template)

Create a BRE category template

Templates define a type of BRE category and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | [**TemplateResource**](TemplateResource.md)| The category template to create | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTemplateUsingDELETE**
> DeleteTemplateUsingDELETE($id, $cascade)

Delete a BRE category template

If cascade = 'detach', it will force delete the template even if it's attached to other objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoriesUsingGET**
> PageBreCategoryResource GetCategoriesUsingGET($size, $page)

List categories


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageBreCategoryResource**](Page«BreCategoryResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoryUsingGET**
> BreCategoryResource GetCategoryUsingGET($name)

Get a single category


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The category name | 

### Return type

[**BreCategoryResource**](BreCategoryResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateUsingGET**
> TemplateResource GetTemplateUsingGET($id)

Get a single BRE category template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplatesUsingGET**
> PageTemplateResource GetTemplatesUsingGET($size, $page, $order)

List and search BRE category templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageTemplateResource**](Page«TemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCategoryUsingPUT**
> UpdateCategoryUsingPUT($name, $category)

Update a category


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The category name | 
 **category** | [**BreCategoryResource**](BreCategoryResource.md)| The updated BRE category information | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTemplateUsingPUT1**
> UpdateTemplateUsingPUT1($id, $template)

Update a BRE category template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **template** | [**TemplateResource**](TemplateResource.md)| The updated category template definition | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

