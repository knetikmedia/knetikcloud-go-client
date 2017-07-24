# \ContentArticlesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArticle**](ContentArticlesApi.md#CreateArticle) | **Post** /content/articles | Create a new article
[**CreateArticleTemplate**](ContentArticlesApi.md#CreateArticleTemplate) | **Post** /content/articles/templates | Create an article template
[**DeleteArticle**](ContentArticlesApi.md#DeleteArticle) | **Delete** /content/articles/{id} | Delete an existing article
[**DeleteArticleTemplate**](ContentArticlesApi.md#DeleteArticleTemplate) | **Delete** /content/articles/templates/{id} | Delete an article template
[**GetArticle**](ContentArticlesApi.md#GetArticle) | **Get** /content/articles/{id} | Get a single article
[**GetArticleTemplate**](ContentArticlesApi.md#GetArticleTemplate) | **Get** /content/articles/templates/{id} | Get a single article template
[**GetArticleTemplates**](ContentArticlesApi.md#GetArticleTemplates) | **Get** /content/articles/templates | List and search article templates
[**GetArticles**](ContentArticlesApi.md#GetArticles) | **Get** /content/articles | List and search articles
[**UpdateArticle**](ContentArticlesApi.md#UpdateArticle) | **Put** /content/articles/{id} | Update an existing article
[**UpdateArticleTemplate**](ContentArticlesApi.md#UpdateArticleTemplate) | **Put** /content/articles/templates/{id} | Update an article template


# **CreateArticle**
> ArticleResource CreateArticle($articleResource)

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

# **CreateArticleTemplate**
> TemplateResource CreateArticleTemplate($articleTemplateResource)

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

# **DeleteArticle**
> DeleteArticle($id)

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

# **DeleteArticleTemplate**
> DeleteArticleTemplate($id, $cascade)

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

# **GetArticle**
> ArticleResource GetArticle($id)

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

# **GetArticleTemplate**
> TemplateResource GetArticleTemplate($id)

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

# **GetArticleTemplates**
> PageResourceTemplateResource GetArticleTemplates($size, $page, $order)

List and search article templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArticles**
> PageResourceArticleResource GetArticles($filterCategory, $filterTagset, $filterTagIntersection, $filterTagExclusion, $filterTitle, $size, $page, $order)

List and search articles

Get a list of articles with optional filtering. Assets will not be filled in on the resources returned. Use 'Get a single article' to retrieve the full resource with assets for a given item as needed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCategory** | **string**| Filter for articles from a specific category by id | [optional] 
 **filterTagset** | **string**| Filter for articles with at least one of a specified set of tags (separated by comma) | [optional] 
 **filterTagIntersection** | **string**| Filter for articles with all of a specified set of tags (separated by comma) | [optional] 
 **filterTagExclusion** | **string**| Filter for articles with none of a specified set of tags (separated by comma) | [optional] 
 **filterTitle** | **string**| Filter for articles whose title contains a string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceArticleResource**](PageResource«ArticleResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArticle**
> ArticleResource UpdateArticle($id, $articleResource)

Update an existing article


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The article id | 
 **articleResource** | [**ArticleResource**](ArticleResource.md)| The article object | [optional] 

### Return type

[**ArticleResource**](ArticleResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArticleTemplate**
> TemplateResource UpdateArticleTemplate($id, $articleTemplateResource)

Update an article template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **articleTemplateResource** | [**TemplateResource**](TemplateResource.md)| The article template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

