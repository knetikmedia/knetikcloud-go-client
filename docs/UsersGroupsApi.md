# \UsersGroupsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemberToGroup**](UsersGroupsApi.md#AddMemberToGroup) | **Post** /users/groups/{unique_name}/members | Adds a new member to the group
[**AddMembersToGroup**](UsersGroupsApi.md#AddMembersToGroup) | **Post** /users/groups/{unique_name}/members/batch-add | Adds multiple members to the group
[**CreateGroup**](UsersGroupsApi.md#CreateGroup) | **Post** /users/groups | Create a group
[**CreateGroupTemplate**](UsersGroupsApi.md#CreateGroupTemplate) | **Post** /users/groups/templates | Create a group template
[**DeleteGroup**](UsersGroupsApi.md#DeleteGroup) | **Delete** /users/groups/{unique_name} | Removes a group from the system IF no resources are attached to it
[**DeleteGroupTemplate**](UsersGroupsApi.md#DeleteGroupTemplate) | **Delete** /users/groups/templates/{id} | Delete a group template
[**GetGroup**](UsersGroupsApi.md#GetGroup) | **Get** /users/groups/{unique_name} | Loads a specific group&#39;s details
[**GetGroupMember**](UsersGroupsApi.md#GetGroupMember) | **Get** /users/groups/{unique_name}/members/{user_id} | Get a user from a group
[**GetGroupMembers**](UsersGroupsApi.md#GetGroupMembers) | **Get** /users/groups/{unique_name}/members | Lists members of the group
[**GetGroupTemplate**](UsersGroupsApi.md#GetGroupTemplate) | **Get** /users/groups/templates/{id} | Get a single group template
[**GetGroupTemplates**](UsersGroupsApi.md#GetGroupTemplates) | **Get** /users/groups/templates | List and search group templates
[**GetGroupsForUser**](UsersGroupsApi.md#GetGroupsForUser) | **Get** /users/{user_id}/groups | List groups a user is in
[**RemoveGroupMember**](UsersGroupsApi.md#RemoveGroupMember) | **Delete** /users/groups/{unique_name}/members/{user_id} | Removes a user from a group
[**UpdateGroup**](UsersGroupsApi.md#UpdateGroup) | **Put** /users/groups/{unique_name} | Update a group
[**UpdateGroupMemberStatus**](UsersGroupsApi.md#UpdateGroupMemberStatus) | **Put** /users/groups/{unique_name}/members/{user_id}/status | Change a user&#39;s status
[**UpdateGroupTemplate**](UsersGroupsApi.md#UpdateGroupTemplate) | **Put** /users/groups/templates/{id} | Update a group template
[**UpdateGroups**](UsersGroupsApi.md#UpdateGroups) | **Get** /users/groups | List and search groups


# **AddMemberToGroup**
> GroupMemberResource AddMemberToGroup($uniqueName, $user)

Adds a new member to the group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **user** | [**GroupMemberResource**](GroupMemberResource.md)| The id and status for a user to add to the group | 

### Return type

[**GroupMemberResource**](GroupMemberResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMembersToGroup**
> []GroupMemberResource AddMembersToGroup($uniqueName, $users)

Adds multiple members to the group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **users** | [**[]GroupMemberResource**](GroupMemberResource.md)| The id and status for a list of users to add to the group | 

### Return type

[**[]GroupMemberResource**](GroupMemberResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> GroupResource CreateGroup($groupResource)

Create a group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupResource** | [**GroupResource**](GroupResource.md)| The new group | [optional] 

### Return type

[**GroupResource**](GroupResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroupTemplate**
> TemplateResource CreateGroupTemplate($groupTemplateResource)

Create a group template

Group Templates define a type of group and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup($uniqueName)

Removes a group from the system IF no resources are attached to it


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupTemplate**
> DeleteGroupTemplate($id, $cascade)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> GroupResource GetGroup($uniqueName)

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

# **GetGroupMember**
> GroupMemberResource GetGroupMember($uniqueName, $userId)

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

# **GetGroupMembers**
> PageResourceGroupMemberResource GetGroupMembers($uniqueName, $size, $page, $order)

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

# **GetGroupTemplate**
> TemplateResource GetGroupTemplate($id)

Get a single group template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupTemplates**
> PageResourceTemplateResource GetGroupTemplates($size, $page, $order)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForUser**
> []string GetGroupsForUser($userId)

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

# **RemoveGroupMember**
> RemoveGroupMember($uniqueName, $userId)

Removes a user from a group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **userId** | **int32**| The id of the user to remove | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> UpdateGroup($uniqueName, $groupResource)

Update a group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **groupResource** | [**GroupResource**](GroupResource.md)| The updated group | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupMemberStatus**
> UpdateGroupMemberStatus($uniqueName, $userId, $status)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupTemplate**
> TemplateResource UpdateGroupTemplate($id, $groupTemplateResource)

Update a group template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **groupTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroups**
> PageResourceGroupResource UpdateGroups($filterTemplate, $filterMemberCount, $filterName, $filterUniqueName, $filterParent, $filterStatus, $size, $page, $order)

List and search groups


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTemplate** | **string**| Filter for groups using a specific template, by id | [optional] 
 **filterMemberCount** | **string**| Filters groups by member count. Multiple values possible for range search. Format: filter_member_count&#x3D;OP,ts&amp;... where OP in (GT, LT, GOE, LOE, EQ). Ex: filter_member_count&#x3D;GT,14,LT,17 | [optional] 
 **filterName** | **string**| Filter for groups with names starting with the given string | [optional] 
 **filterUniqueName** | **string**| Filter for groups whose unique_name starts with provided string | [optional] 
 **filterParent** | **string**| Filter for groups with a specific parent, by unique name | [optional] 
 **filterStatus** | **string**| Filter for groups with a certain status | [optional] 
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

