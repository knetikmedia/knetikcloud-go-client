# \AuthPermissionsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermission**](AuthPermissionsApi.md#CreatePermission) | **Post** /auth/permissions | Create a new permission
[**DeletePermission**](AuthPermissionsApi.md#DeletePermission) | **Delete** /auth/permissions/{permission} | Delete a permission
[**GetPermission**](AuthPermissionsApi.md#GetPermission) | **Get** /auth/permissions/{permission} | Get a single permission
[**GetPermissions**](AuthPermissionsApi.md#GetPermissions) | **Get** /auth/permissions | List and search permissions
[**UpdatePermission**](AuthPermissionsApi.md#UpdatePermission) | **Put** /auth/permissions/{permission} | Update a permission


# **CreatePermission**
> PermissionResource CreatePermission(optional)
Create a new permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionResource** | [**PermissionResource**](PermissionResource.md)| The permission resource object | 

### Return type

[**PermissionResource**](PermissionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePermission**
> DeletePermission(permission, optional)
Delete a permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **permission** | **string**| The permission value | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| The permission value | 
 **force** | **bool**| If true, removes permission assigned to roles | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermission**
> PermissionResource GetPermission(permission)
Get a single permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **permission** | **string**| The permission value | 

### Return type

[**PermissionResource**](PermissionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissions**
> PageResourcePermissionResource GetPermissions(optional)
List and search permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to permission:ASC]

### Return type

[**PageResourcePermissionResource**](PageResource«PermissionResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePermission**
> PermissionResource UpdatePermission(permission, optional)
Update a permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **permission** | **string**| The permission value | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| The permission value | 
 **permissionResource** | [**PermissionResource**](PermissionResource.md)| The permission resource object | 

### Return type

[**PermissionResource**](PermissionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

