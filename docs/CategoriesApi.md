# \CategoriesApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategoryUsingPOST**](CategoriesApi.md#CreateCategoryUsingPOST) | **Post** /categories | Create a new category
[**CreateTemplateUsingPOST2**](CategoriesApi.md#CreateTemplateUsingPOST2) | **Post** /categories/templates | Create a category template
[**DeleteCategoryUsingDELETE**](CategoriesApi.md#DeleteCategoryUsingDELETE) | **Delete** /categories/{id} | Delete an existing category
[**DeleteTemplateUsingDELETE1**](CategoriesApi.md#DeleteTemplateUsingDELETE1) | **Delete** /categories/templates/{id} | Delete a category template
[**GetArticleTemplatesUsingGET1**](CategoriesApi.md#GetArticleTemplatesUsingGET1) | **Get** /categories/templates | List and search category templates
[**GetCategoriesUsingGET1**](CategoriesApi.md#GetCategoriesUsingGET1) | **Get** /categories | List and search categories with optional filters
[**GetCategoryUsingGET1**](CategoriesApi.md#GetCategoryUsingGET1) | **Get** /categories/{id} | Get a single category
[**GetTagsUsingGET**](CategoriesApi.md#GetTagsUsingGET) | **Get** /tags | List all trivia tags in the system
[**GetTemplateUsingGET1**](CategoriesApi.md#GetTemplateUsingGET1) | **Get** /categories/templates/{id} | Get a single category template
[**UpdateCategoryUsingPUT1**](CategoriesApi.md#UpdateCategoryUsingPUT1) | **Put** /categories/{id} | Update an existing category
[**UpdateTemplateUsingPUT2**](CategoriesApi.md#UpdateTemplateUsingPUT2) | **Put** /categories/templates/{id} | Update a category template


# **CreateCategoryUsingPOST**
> CategoryResource CreateCategoryUsingPOST($category)

Create a new category


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**CategoryResource**](CategoryResource.md)| The category to create | [optional] 

### Return type

[**CategoryResource**](CategoryResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTemplateUsingPOST2**
> TemplateResource CreateTemplateUsingPOST2($template)

Create a category template

Templates define a type of category and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | [**TemplateResource**](TemplateResource.md)| The template to create | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCategoryUsingDELETE**
> DeleteCategoryUsingDELETE($id)

Delete an existing category


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the category to be deleted | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTemplateUsingDELETE1**
> DeleteTemplateUsingDELETE1($id, $cascade)

Delete a category template

If cascade = 'detach', it will force delete the template even if it's attached to other objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticleTemplatesUsingGET1**
> PageResourceTemplateResource GetArticleTemplatesUsingGET1($size, $page, $order)

List and search category templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoriesUsingGET1**
> PageResourceCategoryResource GetCategoriesUsingGET1($filterSearch, $filterActive, $size, $page, $order)

List and search categories with optional filters


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSearch** | **string**| Filter for categories whose names begin with provided string | [optional] 
 **filterActive** | **bool**| Filter for categories that are specifically active or inactive | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceCategoryResource**](PageResource«CategoryResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoryUsingGET1**
> CategoryResource GetCategoryUsingGET1($id)

Get a single category


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the category to retrieve | 

### Return type

[**CategoryResource**](CategoryResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagsUsingGET**
> PageResourcestring GetTagsUsingGET($size, $page)

List all trivia tags in the system


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourcestring**](PageResource«string».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateUsingGET1**
> TemplateResource GetTemplateUsingGET1($id)

Get a single category template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCategoryUsingPUT1**
> UpdateCategoryUsingPUT1($id, $category)

Update an existing category


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the category | 
 **category** | [**CategoryResource**](CategoryResource.md)| The category to update | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTemplateUsingPUT2**
> UpdateTemplateUsingPUT2($id, $template)

Update a category template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **template** | [**TemplateResource**](TemplateResource.md)| The updated template information | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

