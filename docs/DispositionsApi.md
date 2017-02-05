# \DispositionsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDispositionUsingPOST**](DispositionsApi.md#AddDispositionUsingPOST) | **Post** /dispositions | Add a new disposition. 
[**DeleteDispositionUsingDELETE**](DispositionsApi.md#DeleteDispositionUsingDELETE) | **Delete** /dispositions/{id} | Delete a disposition
[**GetDispositionCountUsingGET**](DispositionsApi.md#GetDispositionCountUsingGET) | **Get** /dispositions/count | Returns a list of disposition counts
[**GetDispositionUsingGET**](DispositionsApi.md#GetDispositionUsingGET) | **Get** /dispositions/{id} | Returns a disposition
[**GetDispositionsUsingGET**](DispositionsApi.md#GetDispositionsUsingGET) | **Get** /dispositions | Returns a page of dispositions


# **AddDispositionUsingPOST**
> DispositionResource AddDispositionUsingPOST($disposition)

Add a new disposition. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disposition** | [**DispositionResource**](DispositionResource.md)| The new disposition record | [optional] 

### Return type

[**DispositionResource**](DispositionResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDispositionUsingDELETE**
> DeleteDispositionUsingDELETE($id)

Delete a disposition


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the disposition record | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDispositionCountUsingGET**
> []DispositionCount GetDispositionCountUsingGET($filterContext, $filterOwner)

Returns a list of disposition counts


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterContext** | **string**| Filter for dispositions within a context type (games, articles, polls, etc). Optionally with a specific id like filter_context&#x3D;video:47 | [optional] 
 **filterOwner** | **string**| Filter for dispositions from a specific user by id or &#39;me&#39; | [optional] 

### Return type

[**[]DispositionCount**](DispositionCount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDispositionUsingGET**
> DispositionResource GetDispositionUsingGET($id)

Returns a disposition


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the disposition record | 

### Return type

[**DispositionResource**](DispositionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDispositionsUsingGET**
> PageResourceDispositionResource GetDispositionsUsingGET($filterContext, $filterOwner, $size, $page, $order)

Returns a page of dispositions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterContext** | **string**| Filter for dispositions within a context type (games, articles, polls, etc). Optionally with a specific id like filter_context&#x3D;video:47 | [optional] 
 **filterOwner** | **string**| Filter for dispositions from a specific user by id or &#39;me&#39; | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceDispositionResource**](PageResource«DispositionResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

