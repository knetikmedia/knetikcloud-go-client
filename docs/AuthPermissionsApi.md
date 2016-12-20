# \AuthPermissionsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermissionUsingPOST**](AuthPermissionsApi.md#CreatePermissionUsingPOST) | **Post** /auth/permissions | Create a new permission
[**DeletePermissionUsingDELETE**](AuthPermissionsApi.md#DeletePermissionUsingDELETE) | **Delete** /auth/permissions/{permission} | Delete a permission
[**GetPermissionUsingGET**](AuthPermissionsApi.md#GetPermissionUsingGET) | **Get** /auth/permissions/{permission} | Get a single permission
[**GetPermissionUsingGET1**](AuthPermissionsApi.md#GetPermissionUsingGET1) | **Get** /auth/permissions | List and search permissions
[**UpdatePermissionUsingPUT**](AuthPermissionsApi.md#UpdatePermissionUsingPUT) | **Put** /auth/permissions/{permission} | Update a permission


# **CreatePermissionUsingPOST**
> PermissionResource CreatePermissionUsingPOST($permissionResource)

Create a new permission


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionResource** | [**PermissionResource**](PermissionResource.md)| The permission resource object | [optional] 

### Return type

[**PermissionResource**](PermissionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePermissionUsingDELETE**
> DeletePermissionUsingDELETE($permission, $force)

Delete a permission


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| The permission value | 
 **force** | **bool**| If true, removes permission assigned to roles | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionUsingGET**
> PermissionResource GetPermissionUsingGET($permission)

Get a single permission


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| The permission value | 

### Return type

[**PermissionResource**](PermissionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionUsingGET1**
> PagePermissionResource GetPermissionUsingGET1($size, $page, $order)

List and search permissions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to permission:ASC]

### Return type

[**PagePermissionResource**](Page«PermissionResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePermissionUsingPUT**
> UpdatePermissionUsingPUT($permission, $permissionResource)

Update a permission


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| The permission value | 
 **permissionResource** | [**PermissionResource**](PermissionResource.md)| The permission resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

