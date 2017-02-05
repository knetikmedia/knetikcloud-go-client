# \UsersGroupsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupUsingPOST**](UsersGroupsApi.md#AddGroupUsingPOST) | **Post** /users/groups | Adds a new group in the system
[**AddMemberUsingPOST**](UsersGroupsApi.md#AddMemberUsingPOST) | **Post** /users/groups/{unique_name}/members | Adds a new member to the group
[**CreateGroupTemplateUsingPOST**](UsersGroupsApi.md#CreateGroupTemplateUsingPOST) | **Post** /users/groups/templates | Create a group template
[**DeleteGroupMemberUsingDELETE**](UsersGroupsApi.md#DeleteGroupMemberUsingDELETE) | **Delete** /users/groups/{unique_name}/members/{user_id} | Removes a user from a group
[**DeleteGroupTemplateUsingDELETE**](UsersGroupsApi.md#DeleteGroupTemplateUsingDELETE) | **Delete** /users/groups/templates/{id} | Delete a group template
[**DeleteGroupUsingDELETE**](UsersGroupsApi.md#DeleteGroupUsingDELETE) | **Delete** /users/groups/{unique_name} | Removes a group from the system IF no resources are attached to it
[**GetGroupListUsingGET**](UsersGroupsApi.md#GetGroupListUsingGET) | **Get** /users/{user_id}/groups | List groups a user is in
[**GetGroupMemberUsingGET**](UsersGroupsApi.md#GetGroupMemberUsingGET) | **Get** /users/groups/{unique_name}/members/{user_id} | Get a user from a group
[**GetGroupTemplateUsingGET**](UsersGroupsApi.md#GetGroupTemplateUsingGET) | **Get** /users/groups/templates/{id} | Get a single group template
[**GetGroupTemplatesUsingGET**](UsersGroupsApi.md#GetGroupTemplatesUsingGET) | **Get** /users/groups/templates | List and search group templates
[**GetGroupUsingGET**](UsersGroupsApi.md#GetGroupUsingGET) | **Get** /users/groups/{unique_name} | Loads a specific group&#39;s details
[**SearchGroupsUsingGET**](UsersGroupsApi.md#SearchGroupsUsingGET) | **Get** /users/groups/{unique_name}/members | Lists members of the group
[**SearchGroupsUsingGET1**](UsersGroupsApi.md#SearchGroupsUsingGET1) | **Get** /users/groups | List and search groups
[**UpdateGroupTemplateUsingPUT**](UsersGroupsApi.md#UpdateGroupTemplateUsingPUT) | **Put** /users/groups/templates/{id} | Update a group template
[**UpdateGroupUsingPUT**](UsersGroupsApi.md#UpdateGroupUsingPUT) | **Put** /users/groups/{unique_name} | Modifies a group&#39;s details
[**UpdateMemberStatusUsingPUT**](UsersGroupsApi.md#UpdateMemberStatusUsingPUT) | **Put** /users/groups/{unique_name}/members/{user_id}/status | Change a user&#39;s status


# **AddGroupUsingPOST**
> GroupResource AddGroupUsingPOST($groupResource)

Adds a new group in the system


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupResource** | [**GroupResource**](GroupResource.md)| The new group | [optional] 

### Return type

[**GroupResource**](GroupResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMemberUsingPOST**
> GroupMemberResource AddMemberUsingPOST($uniqueName, $username)

Adds a new member to the group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **username** | [**GroupMemberResource**](GroupMemberResource.md)| The username of a user to add to the group | 

### Return type

[**GroupMemberResource**](GroupMemberResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroupTemplateUsingPOST**
> TemplateResource CreateGroupTemplateUsingPOST($groupTemplateResource)

Create a group template

Group Templates define a type of group and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupMemberUsingDELETE**
> DeleteGroupMemberUsingDELETE($uniqueName, $userId)

Removes a user from a group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **userId** | **int32**| The id of the user to remove | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupTemplateUsingDELETE**
> DeleteGroupTemplateUsingDELETE($id, $cascade)

Delete a group template

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

# **DeleteGroupUsingDELETE**
> DeleteGroupUsingDELETE($uniqueName)

Removes a group from the system IF no resources are attached to it


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupListUsingGET**
> []string GetGroupListUsingGET($userId)

List groups a user is in


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupMemberUsingGET**
> GroupMemberResource GetGroupMemberUsingGET($uniqueName, $userId)

Get a user from a group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **userId** | **int32**| The id of the user | 

### Return type

[**GroupMemberResource**](GroupMemberResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupTemplateUsingGET**
> TemplateResource GetGroupTemplateUsingGET($id)

Get a single group template


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

# **GetGroupTemplatesUsingGET**
> PageResourceTemplateResource GetGroupTemplatesUsingGET($size, $page, $order)

List and search group templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| a comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupUsingGET**
> GroupResource GetGroupUsingGET($uniqueName)

Loads a specific group's details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 

### Return type

[**GroupResource**](GroupResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGroupsUsingGET**
> PageResourceGroupMemberResource SearchGroupsUsingGET($uniqueName, $size, $page, $order)

Lists members of the group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceGroupMemberResource**](PageResource«GroupMemberResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGroupsUsingGET1**
> PageResourceGroupResource SearchGroupsUsingGET1($size, $page, $order)

List and search groups


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageResourceGroupResource**](PageResource«GroupResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupTemplateUsingPUT**
> UpdateGroupTemplateUsingPUT($id, $groupTemplateResource)

Update a group template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **groupTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupUsingPUT**
> UpdateGroupUsingPUT($uniqueName, $groupResource)

Modifies a group's details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **groupResource** | [**GroupResource**](GroupResource.md)| The updated group | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMemberStatusUsingPUT**
> UpdateMemberStatusUsingPUT($uniqueName, $userId, $status)

Change a user's status


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **userId** | **int32**| The user id of the member to modify | 
 **status** | **string**| The new status for the user | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

