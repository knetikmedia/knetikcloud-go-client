# \ContentArticlesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArticleTemplateUsingPOST**](ContentArticlesApi.md#CreateArticleTemplateUsingPOST) | **Post** /content/articles/templates | Create an article template
[**CreateArticleUsingPOST**](ContentArticlesApi.md#CreateArticleUsingPOST) | **Post** /content/articles | Create a new article
[**DeleteArticleTemplateUsingDELETE**](ContentArticlesApi.md#DeleteArticleTemplateUsingDELETE) | **Delete** /content/articles/templates/{id} | Delete an article template
[**DeleteArticleUsingDELETE**](ContentArticlesApi.md#DeleteArticleUsingDELETE) | **Delete** /content/articles/{id} | Delete an existing article
[**GetArticleTemplateUsingGET**](ContentArticlesApi.md#GetArticleTemplateUsingGET) | **Get** /content/articles/templates/{id} | Get a single article template
[**GetArticleTemplatesUsingGET**](ContentArticlesApi.md#GetArticleTemplatesUsingGET) | **Get** /content/articles/templates | List and search article templates
[**GetArticleUsingGET**](ContentArticlesApi.md#GetArticleUsingGET) | **Get** /content/articles/{id} | Get a single article
[**GetArticlesUsingGET**](ContentArticlesApi.md#GetArticlesUsingGET) | **Get** /content/articles | List and search articles
[**UpdateArticleTemplateUsingPUT**](ContentArticlesApi.md#UpdateArticleTemplateUsingPUT) | **Put** /content/articles/templates/{id} | Update an article template
[**UpdateArticleUsingPUT**](ContentArticlesApi.md#UpdateArticleUsingPUT) | **Put** /content/articles/{id} | Update an existing article


# **CreateArticleTemplateUsingPOST**
> TemplateResource CreateArticleTemplateUsingPOST($articleTemplateResource)

Create an article template

Article Templates define a type of article and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **articleTemplateResource** | [**TemplateResource**](TemplateResource.md)| The article template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateArticleUsingPOST**
> ArticleResource CreateArticleUsingPOST($articleResource)

Create a new article

Articles are blobs of text with titles, a category and assets. Formatting and display of the text is in the hands of the front end.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **articleResource** | [**ArticleResource**](ArticleResource.md)| The new article | [optional] 

### Return type

[**ArticleResource**](ArticleResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArticleTemplateUsingDELETE**
> DeleteArticleTemplateUsingDELETE($id, $cascade)

Delete an article template

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArticleUsingDELETE**
> DeleteArticleUsingDELETE($id)

Delete an existing article


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The article id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticleTemplateUsingGET**
> TemplateResource GetArticleTemplateUsingGET($id)

Get a single article template


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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticleTemplatesUsingGET**
> PageTemplateResource GetArticleTemplatesUsingGET($size, $page, $order)

List and search article templates


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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticleUsingGET**
> ArticleResource GetArticleUsingGET($id)

Get a single article


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The article id | 

### Return type

[**ArticleResource**](ArticleResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticlesUsingGET**
> PageArticleResource GetArticlesUsingGET($filterCategory, $filterTagset, $filterTitle, $size, $page, $order)

List and search articles

Get a list of articles with optional filtering. Assets will not be filled in on the resources returned. Use 'Get a single article' to retrieve the full resource with assets for a given item as needed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCategory** | **string**| Filter for articles from a specific category by id | [optional] 
 **filterTagset** | **string**| Filter for articles with specified tags (separated by comma) | [optional] 
 **filterTitle** | **string**| Filter for articles whose title contains a string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageArticleResource**](Page«ArticleResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArticleTemplateUsingPUT**
> UpdateArticleTemplateUsingPUT($id, $articleTemplateResource)

Update an article template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **articleTemplateResource** | [**TemplateResource**](TemplateResource.md)| The article template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArticleUsingPUT**
> UpdateArticleUsingPUT($id, $articleResource)

Update an existing article


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The article id | 
 **articleResource** | [**ArticleResource**](ArticleResource.md)| The article object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

