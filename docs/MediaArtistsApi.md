# \MediaArtistsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArtistUsingPOST**](MediaArtistsApi.md#AddArtistUsingPOST) | **Post** /media/artists | Adds a new artist in the system
[**CreateArtistTemplateUsingPOST**](MediaArtistsApi.md#CreateArtistTemplateUsingPOST) | **Post** /media/artists/templates | Create an artist template
[**DeleteArtistTemplateUsingDELETE**](MediaArtistsApi.md#DeleteArtistTemplateUsingDELETE) | **Delete** /media/artists/templates/{id} | Delete an artist template
[**DeleteArtistUsingDELETE**](MediaArtistsApi.md#DeleteArtistUsingDELETE) | **Delete** /media/artists/{id} | Removes an artist from the system IF no resources are attached to it
[**GetArtistTemplateUsingGET**](MediaArtistsApi.md#GetArtistTemplateUsingGET) | **Get** /media/artists/templates/{id} | Get a single artist template
[**GetArtistTemplatesUsingGET**](MediaArtistsApi.md#GetArtistTemplatesUsingGET) | **Get** /media/artists/templates | List and search artist templates
[**GetArtistUsingGET**](MediaArtistsApi.md#GetArtistUsingGET) | **Get** /media/artists/{id} | Loads a specific artist details
[**SearchArtistsUsingGET**](MediaArtistsApi.md#SearchArtistsUsingGET) | **Get** /media/artists | Search for artists
[**UpdateArtistTemplateUsingPUT**](MediaArtistsApi.md#UpdateArtistTemplateUsingPUT) | **Put** /media/artists/templates/{id} | Update an artist template
[**UpdateArtistUsingPUT**](MediaArtistsApi.md#UpdateArtistUsingPUT) | **Put** /media/artists/{id} | Modifies an artist details


# **AddArtistUsingPOST**
> ArtistResource AddArtistUsingPOST($artistResource)

Adds a new artist in the system

Adds a new artist in the system. Use specific media contributions endpoint to add contributions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistResource** | [**ArtistResource**](ArtistResource.md)| The new artist | [optional] 

### Return type

[**ArtistResource**](ArtistResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateArtistTemplateUsingPOST**
> TemplateResource CreateArtistTemplateUsingPOST($artistTemplateResource)

Create an artist template

Artist Templates define a type of artist and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistTemplateResource** | [**TemplateResource**](TemplateResource.md)| The artist template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArtistTemplateUsingDELETE**
> DeleteArtistTemplateUsingDELETE($id, $cascade)

Delete an artist template

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

# **DeleteArtistUsingDELETE**
> DeleteArtistUsingDELETE($id)

Removes an artist from the system IF no resources are attached to it


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The artist id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArtistTemplateUsingGET**
> TemplateResource GetArtistTemplateUsingGET($id)

Get a single artist template


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

# **GetArtistTemplatesUsingGET**
> PageResourceTemplateResource GetArtistTemplatesUsingGET($size, $page, $order)

List and search artist templates


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

# **GetArtistUsingGET**
> ArtistResource GetArtistUsingGET($id, $showContributions)

Loads a specific artist details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The artist id | 
 **showContributions** | **int32**| The number of contributions to show fetch | [optional] 

### Return type

[**ArtistResource**](ArtistResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchArtistsUsingGET**
> PageResourceArtistResource SearchArtistsUsingGET($filterArtistsByName, $size, $page, $order)

Search for artists


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterArtistsByName** | **string**| Filter for artists which name *STARTS* with the given string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceArtistResource**](PageResource«ArtistResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArtistTemplateUsingPUT**
> UpdateArtistTemplateUsingPUT($id, $artistTemplateResource)

Update an artist template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **artistTemplateResource** | [**TemplateResource**](TemplateResource.md)| The artist template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateArtistUsingPUT**
> UpdateArtistUsingPUT($id, $artistResource)

Modifies an artist details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The artist id | 
 **artistResource** | [**ArtistResource**](ArtistResource.md)| The new artist | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

