# \AuthRolesApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignClientRolesUsingPUT**](AuthRolesApi.md#AssignClientRolesUsingPUT) | **Put** /auth/clients/{client_key}/roles | Set roles for a client
[**AssignPermissionsUsingPUT**](AuthRolesApi.md#AssignPermissionsUsingPUT) | **Put** /auth/roles/{role}/permissions | Set permissions for a role
[**AssignUserRolesExternalUsingPUT**](AuthRolesApi.md#AssignUserRolesExternalUsingPUT) | **Put** /auth/users/{user_id}/roles | Set roles for a user
[**CreateRoleUsingPOST**](AuthRolesApi.md#CreateRoleUsingPOST) | **Post** /auth/roles | Create a new role
[**DeleteRoleUsingDELETE**](AuthRolesApi.md#DeleteRoleUsingDELETE) | **Delete** /auth/roles/{role} | Delete a role
[**GetClientRolesUsingGET**](AuthRolesApi.md#GetClientRolesUsingGET) | **Get** /auth/clients/{client_key}/roles | Get roles for a client
[**GetRoleUsingGET**](AuthRolesApi.md#GetRoleUsingGET) | **Get** /auth/roles/{role} | Get a single role
[**GetRolesUsingGET**](AuthRolesApi.md#GetRolesUsingGET) | **Get** /auth/roles | List and search roles
[**GetUserRolesUsingGET**](AuthRolesApi.md#GetUserRolesUsingGET) | **Get** /auth/users/{user_id}/roles | Get roles for a user
[**UpdateRoleUsingPUT**](AuthRolesApi.md#UpdateRoleUsingPUT) | **Put** /auth/roles/{role} | Update a role


# **AssignClientRolesUsingPUT**
> AssignClientRolesUsingPUT($clientKey, $rolesList)

Set roles for a client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The client key | 
 **rolesList** | **[]string**| The list of unique roles | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignPermissionsUsingPUT**
> AssignPermissionsUsingPUT($role, $permissionsList)

Set permissions for a role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string**| The role value | 
 **permissionsList** | **[]string**| The list of unique permissions | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignUserRolesExternalUsingPUT**
> AssignUserRolesExternalUsingPUT($userId, $rolesList)

Set roles for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user&#39;s id | 
 **rolesList** | **[]string**| The list of unique roles | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRoleUsingPOST**
> RoleResource CreateRoleUsingPOST($roleResource)

Create a new role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleResource** | [**RoleResource**](RoleResource.md)| The role resource object | [optional] 

### Return type

[**RoleResource**](RoleResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleUsingDELETE**
> DeleteRoleUsingDELETE($role, $force)

Delete a role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string**| The role value | 
 **force** | **bool**| If true, removes role from users/clients | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientRolesUsingGET**
> []RoleResource GetClientRolesUsingGET($clientKey)

Get roles for a client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| The client key | 

### Return type

[**[]RoleResource**](RoleResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleUsingGET**
> RoleResource GetRoleUsingGET($role)

Get a single role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string**| The role value | 

### Return type

[**RoleResource**](RoleResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRolesUsingGET**
> PageResourceRoleResource GetRolesUsingGET($size, $page, $order)

List and search roles


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageResourceRoleResource**](PageResource«RoleResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserRolesUsingGET**
> []RoleResource GetUserRolesUsingGET($userId)

Get roles for a user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user&#39;s id | 

### Return type

[**[]RoleResource**](RoleResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleUsingPUT**
> UpdateRoleUsingPUT($role, $roleResource)

Update a role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string**| The role value | 
 **roleResource** | [**RoleResource**](RoleResource.md)| The role resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

