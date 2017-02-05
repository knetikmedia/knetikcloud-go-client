# \UsersRelationshipsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRelationshipUsingPOST**](UsersRelationshipsApi.md#AddRelationshipUsingPOST) | **Post** /users/relationships | Create a user relationship
[**DeleteRelationshipUsingDELETE**](UsersRelationshipsApi.md#DeleteRelationshipUsingDELETE) | **Delete** /users/relationships/{id} | Delete a user relationship
[**GetRelationshipUsingGET**](UsersRelationshipsApi.md#GetRelationshipUsingGET) | **Get** /users/relationships/{id} | Get a user relationship
[**GetRelationshipsUsingGET**](UsersRelationshipsApi.md#GetRelationshipsUsingGET) | **Get** /users/relationships | Get a list of user relationships
[**UpdateRelationshipUsingPUT**](UsersRelationshipsApi.md#UpdateRelationshipUsingPUT) | **Put** /users/relationships/{id} | Update a user relationship


# **AddRelationshipUsingPOST**
> UserRelationshipResource AddRelationshipUsingPOST($relationship)

Create a user relationship


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relationship** | [**UserRelationshipResource**](UserRelationshipResource.md)| The new relationship | [optional] 

### Return type

[**UserRelationshipResource**](UserRelationshipResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRelationshipUsingDELETE**
> DeleteRelationshipUsingDELETE($id)

Delete a user relationship


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the relationship | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelationshipUsingGET**
> UserRelationshipResource GetRelationshipUsingGET($id)

Get a user relationship


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the relationship | 

### Return type

[**UserRelationshipResource**](UserRelationshipResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelationshipsUsingGET**
> PageResourceUserRelationshipResource GetRelationshipsUsingGET()

Get a list of user relationships


### Parameters
This endpoint does not need any parameter.

### Return type

[**PageResourceUserRelationshipResource**](PageResource«UserRelationshipResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRelationshipUsingPUT**
> UserRelationshipResource UpdateRelationshipUsingPUT($id, $relationship)

Update a user relationship


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the relationship | 
 **relationship** | [**UserRelationshipResource**](UserRelationshipResource.md)| The new relationship | [optional] 

### Return type

[**UserRelationshipResource**](UserRelationshipResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

