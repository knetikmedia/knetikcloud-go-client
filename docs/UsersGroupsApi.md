# \UsersGroupsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemberToGroup**](UsersGroupsApi.md#AddMemberToGroup) | **Post** /users/groups/{unique_name}/members | Adds a new member to the group
[**AddMembersToGroup**](UsersGroupsApi.md#AddMembersToGroup) | **Post** /users/groups/{unique_name}/members/batch-add | Adds multiple members to the group
[**CreateGroup**](UsersGroupsApi.md#CreateGroup) | **Post** /users/groups | Create a group
[**CreateGroupMemberTemplate**](UsersGroupsApi.md#CreateGroupMemberTemplate) | **Post** /users/groups/members/templates | Create an group member template
[**CreateGroupTemplate**](UsersGroupsApi.md#CreateGroupTemplate) | **Post** /users/groups/templates | Create a group template
[**DeleteGroup**](UsersGroupsApi.md#DeleteGroup) | **Delete** /users/groups/{unique_name} | Removes a group from the system IF no resources are attached to it
[**DeleteGroupMemberTemplate**](UsersGroupsApi.md#DeleteGroupMemberTemplate) | **Delete** /users/groups/members/templates/{id} | Delete an group member template
[**DeleteGroupTemplate**](UsersGroupsApi.md#DeleteGroupTemplate) | **Delete** /users/groups/templates/{id} | Delete a group template
[**GetGroup**](UsersGroupsApi.md#GetGroup) | **Get** /users/groups/{unique_name} | Loads a specific group&#39;s details
[**GetGroupMember**](UsersGroupsApi.md#GetGroupMember) | **Get** /users/groups/{unique_name}/members/{user_id} | Get a user from a group
[**GetGroupMemberTemplate**](UsersGroupsApi.md#GetGroupMemberTemplate) | **Get** /users/groups/members/templates/{id} | Get a single group member template
[**GetGroupMemberTemplates**](UsersGroupsApi.md#GetGroupMemberTemplates) | **Get** /users/groups/members/templates | List and search group member templates
[**GetGroupMembers**](UsersGroupsApi.md#GetGroupMembers) | **Get** /users/groups/{unique_name}/members | Lists members of the group
[**GetGroupTemplate**](UsersGroupsApi.md#GetGroupTemplate) | **Get** /users/groups/templates/{id} | Get a single group template
[**GetGroupTemplates**](UsersGroupsApi.md#GetGroupTemplates) | **Get** /users/groups/templates | List and search group templates
[**GetGroupsForUser**](UsersGroupsApi.md#GetGroupsForUser) | **Get** /users/{user_id}/groups | List groups a user is in
[**ListGroups**](UsersGroupsApi.md#ListGroups) | **Get** /users/groups | List and search groups
[**RemoveGroupMember**](UsersGroupsApi.md#RemoveGroupMember) | **Delete** /users/groups/{unique_name}/members/{user_id} | Removes a user from a group
[**UpdateGroup**](UsersGroupsApi.md#UpdateGroup) | **Put** /users/groups/{unique_name} | Update a group
[**UpdateGroupMemberProperties**](UsersGroupsApi.md#UpdateGroupMemberProperties) | **Put** /users/groups/{unique_name}/members/{user_id}/order | Change a user&#39;s order
[**UpdateGroupMemberProperties1**](UsersGroupsApi.md#UpdateGroupMemberProperties1) | **Put** /users/groups/{unique_name}/members/{user_id}/properties | Change a user&#39;s membership properties
[**UpdateGroupMemberStatus**](UsersGroupsApi.md#UpdateGroupMemberStatus) | **Put** /users/groups/{unique_name}/members/{user_id}/status | Change a user&#39;s status
[**UpdateGroupMemberTemplate**](UsersGroupsApi.md#UpdateGroupMemberTemplate) | **Put** /users/groups/members/templates/{id} | Update an group member template
[**UpdateGroupTemplate**](UsersGroupsApi.md#UpdateGroupTemplate) | **Put** /users/groups/templates/{id} | Update a group template


# **AddMemberToGroup**
> GroupMemberResource AddMemberToGroup(ctx, ctx, uniqueName, user)
Adds a new member to the group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 
  **user** | [**GroupMemberResource**](GroupMemberResource.md)| The id and status for a user to add to the group | 

### Return type

[**GroupMemberResource**](GroupMemberResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMembersToGroup**
> []GroupMemberResource AddMembersToGroup(ctx, ctx, uniqueName, users)
Adds multiple members to the group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 
  **users** | [**[]GroupMemberResource**](GroupMemberResource.md)| The id and status for a list of users to add to the group | 

### Return type

[**[]GroupMemberResource**](GroupMemberResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> GroupResource CreateGroup(ctx, ctx, optional)
Create a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupResource** | [**GroupResource**](GroupResource.md)| The new group | 

### Return type

[**GroupResource**](GroupResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroupMemberTemplate**
> TemplateResource CreateGroupMemberTemplate(ctx, ctx, optional)
Create an group member template

GroupMember Templates define a type of group member and the properties they have

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupMemberTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group member template resource object | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroupTemplate**
> TemplateResource CreateGroupTemplate(ctx, ctx, optional)
Create a group template

Group Templates define a type of group and the properties they have

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group template resource object | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, ctx, uniqueName)
Removes a group from the system IF no resources are attached to it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupMemberTemplate**
> DeleteGroupMemberTemplate(ctx, ctx, id, optional)
Delete an group member template

If cascade = 'detach', it will force delete the template even if it's attached to other objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupTemplate**
> DeleteGroupTemplate(ctx, ctx, id, optional)
Delete a group template

If cascade = 'detach', it will force delete the template even if it's attached to other objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> GroupResource GetGroup(uniqueName)
Loads a specific group's details

### Required Parameters

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
> GroupMemberResource GetGroupMember(uniqueName, userId)
Get a user from a group

### Required Parameters

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

# **GetGroupMemberTemplate**
> TemplateResource GetGroupMemberTemplate(ctx, ctx, id)
Get a single group member template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupMemberTemplates**
> PageResourceTemplateResource GetGroupMemberTemplates(ctx, ctx, optional)
List and search group member templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupMembers**
> PageResourceGroupMemberResource GetGroupMembers(uniqueName, optional)
Lists members of the group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **uniqueName** | **string**| The group unique name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to order:ASC]

### Return type

[**PageResourceGroupMemberResource**](PageResource«GroupMemberResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupTemplate**
> TemplateResource GetGroupTemplate(ctx, ctx, id)
Get a single group template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupTemplates**
> PageResourceTemplateResource GetGroupTemplates(ctx, ctx, optional)
List and search group templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| a comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForUser**
> []string GetGroupsForUser(userId, optional)
List groups a user is in

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userId** | **int32**| The id of the user | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **filterChildren** | **bool**| Whether to limit group list to children of groups only. If true, shows only groups with parents. If false, shows only groups with no parent. | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroups**
> PageResourceGroupResource ListGroups(optional)
List and search groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTemplate** | **string**| Filter for groups using a specific template, by id | 
 **filterMemberCount** | **string**| Filters groups by member count. Multiple values possible for range search. Format: filter_member_count&#x3D;OP,ts&amp;... where OP in (GT, LT, GOE, LOE, EQ). Ex: filter_member_count&#x3D;GT,14,LT,17 | 
 **filterName** | **string**| Filter for groups with names starting with the given string | 
 **filterUniqueName** | **string**| Filter for groups whose unique_name starts with provided string | 
 **filterParent** | **string**| Filter for groups with a specific parent, by unique name | 
 **filterStatus** | **string**| Filter for groups with a certain status | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to name:ASC]

### Return type

[**PageResourceGroupResource**](PageResource«GroupResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGroupMember**
> RemoveGroupMember(ctx, ctx, uniqueName, userId)
Removes a user from a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 
  **userId** | **int32**| The id of the user to remove | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> UpdateGroup(ctx, ctx, uniqueName, optional)
Update a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueName** | **string**| The group unique name | 
 **groupResource** | [**GroupResource**](GroupResource.md)| The updated group | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupMemberProperties**
> UpdateGroupMemberProperties(ctx, ctx, uniqueName, userId, order)
Change a user's order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 
  **userId** | **int32**| The user id of the member to modify | 
  **order** | [**StringWrapper**](StringWrapper.md)| The new order for the membership | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupMemberProperties1**
> UpdateGroupMemberProperties1(ctx, ctx, uniqueName, userId, properties)
Change a user's membership properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 
  **userId** | **int32**| The user id of the member to modify | 
  **properties** | [**interface{}**](interface{}.md)| The new properties for the membership | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupMemberStatus**
> UpdateGroupMemberStatus(ctx, ctx, uniqueName, userId, status)
Change a user's status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **uniqueName** | **string**| The group unique name | 
  **userId** | **int32**| The user id of the member to modify | 
  **status** | **string**| The new status for the user | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupMemberTemplate**
> TemplateResource UpdateGroupMemberTemplate(ctx, ctx, id, optional)
Update an group member template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **groupMemberTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group member template resource object | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupTemplate**
> TemplateResource UpdateGroupTemplate(ctx, ctx, id, optional)
Update a group template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **groupTemplateResource** | [**TemplateResource**](TemplateResource.md)| The group template resource object | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

