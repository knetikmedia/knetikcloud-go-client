# \SearchApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSearchIndex**](SearchApi.md#AddSearchIndex) | **Post** /search/index/{type}/{id} | Add a new object to an index
[**AddSearchMappings**](SearchApi.md#AddSearchMappings) | **Post** /search/mappings | Register reference mappings
[**DeleteSearchIndex**](SearchApi.md#DeleteSearchIndex) | **Delete** /search/index/{type}/{id} | Delete an object
[**DeleteSearchIndexes**](SearchApi.md#DeleteSearchIndexes) | **Delete** /search/index/{type} | Delete all objects in an index
[**SearchIndex**](SearchApi.md#SearchIndex) | **Post** /search/index/{type} | Search an index


# **AddSearchIndex**
> AddSearchIndex(type_, id, optional)
Add a new object to an index

Mainly intended for internal use.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **type_** | **string**| The index type | 
  **id** | **string**| The ID of the object | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| The index type | 
 **id** | **string**| The ID of the object | 
 **object** | [**interface{}**](interface{}.md)| The object to add | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSearchMappings**
> AddSearchMappings(optional)
Register reference mappings

Add a new type mapping to connect data from one index to another, or discover an exsting one. Mainly intended for internal use.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappings** | [**[]SearchReferenceMapping**](SearchReferenceMapping.md)| The mappings to add | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSearchIndex**
> DeleteSearchIndex(type_, id)
Delete an object

Mainly intended for internal use. Requires SEARCH_ADMIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **type_** | **string**| The index type | 
  **id** | **string**| The ID of the object to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSearchIndexes**
> DeleteSearchIndexes(type_)
Delete all objects in an index

Mainly intended for internal use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **type_** | **string**| The index type | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchIndex**
> PageResourceMapstringobject SearchIndex(type_, optional)
Search an index

The body is an ElasticSearch query in JSON format. Please see their <a href='https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html'>documentation</a> for details on the format and search options. The searchable object's format depends on on the type. See individual search endpoints on other resources for details on their format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **type_** | **string**| The index type | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| The index type | 
 **query** | [**interface{}**](interface{}.md)| The query to be used for the search | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceMapstringobject**](PageResource«Map«string,object»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

