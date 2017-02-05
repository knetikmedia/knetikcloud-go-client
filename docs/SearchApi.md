# \SearchApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalAddUsingPOST**](SearchApi.md#ExternalAddUsingPOST) | **Post** /search/index/{type}/{id} | Add a new object to an index
[**ExternalDeleteAllUsingDELETE**](SearchApi.md#ExternalDeleteAllUsingDELETE) | **Delete** /search/index/{type} | Delete all objects in an index
[**ExternalDeleteUsingDELETE**](SearchApi.md#ExternalDeleteUsingDELETE) | **Delete** /search/index/{type}/{id} | Delete an object
[**ExternalRegisterUsingPOST**](SearchApi.md#ExternalRegisterUsingPOST) | **Post** /search/mappings | Register reference mappings
[**SearchUsingPOST**](SearchApi.md#SearchUsingPOST) | **Post** /search/index/{type} | Search an index


# **ExternalAddUsingPOST**
> ExternalAddUsingPOST($type_, $id, $object)

Add a new object to an index

Mainly intended for internal use.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| The index type | 
 **id** | **string**| The ID of the object | 
 **object** | [**interface{}**](interface{}.md)| The object to add | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExternalDeleteAllUsingDELETE**
> ExternalDeleteAllUsingDELETE($type_)

Delete all objects in an index

Mainly intended for internal use


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| The index type | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExternalDeleteUsingDELETE**
> ExternalDeleteUsingDELETE($type_, $id)

Delete an object

Mainly intended for internal use. Requires SEARCH_ADMIN.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| The index type | 
 **id** | **string**| The ID of the object to delete | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExternalRegisterUsingPOST**
> ExternalRegisterUsingPOST($mappings)

Register reference mappings

Add a new type mapping to connect data from one index to another, or discover an exsting one. Mainly intended for internal use.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappings** | [**[]SearchReferenceMapping**](SearchReferenceMapping.md)| The mappings to add | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsingPOST**
> PageResourceMapstringobject SearchUsingPOST($type_, $query, $size, $page)

Search an index

The body is an ElasticSearch query in JSON format. Please see their <a href='https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html'>documentation</a> for details on the format and search options. The searchable object's format depends on on the type. See individual search endpoints on other resources for details on their format.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| The index type | 
 **query** | [**interface{}**](interface{}.md)| The query to be used for the search | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceMapstringobject**](PageResource«Map«string,object»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

