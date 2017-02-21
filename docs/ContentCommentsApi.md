# \ContentCommentsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddComment**](ContentCommentsApi.md#AddComment) | **Post** /comments | Add a new comment
[**DeleteComment**](ContentCommentsApi.md#DeleteComment) | **Delete** /comments/{id} | Delete a comment
[**GetComment**](ContentCommentsApi.md#GetComment) | **Get** /comments/{id} | Return a comment
[**GetComments**](ContentCommentsApi.md#GetComments) | **Get** /comments | Returns a page of comments
[**SearchComments**](ContentCommentsApi.md#SearchComments) | **Post** /comments/search | Search the comment index
[**UpdateComment**](ContentCommentsApi.md#UpdateComment) | **Put** /comments/{id}/content | Update a comment


# **AddComment**
> CommentResource AddComment($commentResource)

Add a new comment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentResource** | [**CommentResource**](CommentResource.md)| The comment to be added | [optional] 

### Return type

[**CommentResource**](CommentResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComment**
> DeleteComment($id)

Delete a comment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The comment id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComment**
> CommentResource GetComment($id)

Return a comment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The comment id | 

### Return type

[**CommentResource**](CommentResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComments**
> PageResourceCommentResource GetComments($context, $contextId, $size, $page)

Returns a page of comments


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string**| Get comments by context type | 
 **contextId** | **int32**| Get comments by context id | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceCommentResource**](PageResource«CommentResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchComments**
> CommentSearch SearchComments($query, $size, $page)

Search the comment index

The body is an ElasticSearch query json. Please see their <a href='https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html'>documentation</a> for details on the format and search options


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**interface{}**](interface{}.md)| The search query | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**CommentSearch**](CommentSearch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateComment**
> UpdateComment($id, $content)

Update a comment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The comment id | 
 **content** | **string**| The comment content | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

