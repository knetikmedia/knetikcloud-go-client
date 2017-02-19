# \GamificationAchievementsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAchievement**](GamificationAchievementsApi.md#CreateAchievement) | **Post** /achievements | Create a new achievement definition
[**CreateAchievementTemplate**](GamificationAchievementsApi.md#CreateAchievementTemplate) | **Post** /achievements/templates | Create an achievement template
[**DeleteAchievement**](GamificationAchievementsApi.md#DeleteAchievement) | **Delete** /achievements/{name} | Delete an achievement definition
[**DeleteAchievementTemplate**](GamificationAchievementsApi.md#DeleteAchievementTemplate) | **Delete** /achievements/templates/{id} | Delete an achievement template
[**GetAchievement**](GamificationAchievementsApi.md#GetAchievement) | **Get** /achievements/{name} | Get a single achievement definition
[**GetAchievementTemplate**](GamificationAchievementsApi.md#GetAchievementTemplate) | **Get** /achievements/templates/{id} | Get a single achievement template
[**GetAchievementTemplates**](GamificationAchievementsApi.md#GetAchievementTemplates) | **Get** /achievements/templates | List and search achievement templates
[**GetAchievementTriggers**](GamificationAchievementsApi.md#GetAchievementTriggers) | **Get** /achievements/triggers | Get the list of triggers that can be used to trigger an achievement progress update
[**GetAchievements**](GamificationAchievementsApi.md#GetAchievements) | **Get** /achievements | Get all achievement definitions in the system
[**GetDerivedAchievements**](GamificationAchievementsApi.md#GetDerivedAchievements) | **Get** /achievements/derived/{name} | Get a list of derived achievements
[**GetUserAchievementProgress**](GamificationAchievementsApi.md#GetUserAchievementProgress) | **Get** /users/{user_id}/achievements/{achievement_name} | Retrieve progress on a given achievement for a given user
[**GetUserAchievementsProgress**](GamificationAchievementsApi.md#GetUserAchievementsProgress) | **Get** /users/{user_id}/achievements | Retrieve progress on achievements for a given user
[**GetUsersAchievementProgress**](GamificationAchievementsApi.md#GetUsersAchievementProgress) | **Get** /users/achievements/{achievement_name} | Retrieve progress on a given achievement for all users
[**GetUsersAchievementsProgress**](GamificationAchievementsApi.md#GetUsersAchievementsProgress) | **Get** /users/achievements | Retrieve progress on achievements for all users
[**UpdateAchievement**](GamificationAchievementsApi.md#UpdateAchievement) | **Put** /achievements/{name} | Update an achievement definition
[**UpdateAchievementProgress**](GamificationAchievementsApi.md#UpdateAchievementProgress) | **Put** /users/{user_id}/achievements/{achievement_name} | Update or create an achievement progress record for a user
[**UpdateAchievementTemplate**](GamificationAchievementsApi.md#UpdateAchievementTemplate) | **Put** /achievements/templates/{id} | Update an achievement template


# **CreateAchievement**
> AchievementDefinitionResource CreateAchievement($achievement)

Create a new achievement definition

If the definition contains a trigger event name, a BRE rule is created, so that tracking logic is executed when the triggering event occurs. If no trigger event name is specified, the user's achievement status must manually be updated via the API.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achievement** | [**AchievementDefinitionResource**](AchievementDefinitionResource.md)| The achievement definition | [optional] 

### Return type

[**AchievementDefinitionResource**](AchievementDefinitionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAchievementTemplate**
> TemplateResource CreateAchievementTemplate($template)

Create an achievement template

Achievement templates define a type of achievement and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | [**TemplateResource**](TemplateResource.md)| The achievement template to be created | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAchievement**
> DeleteAchievement($name)

Delete an achievement definition

Will also disable the associated generated rule, if any.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the achievement | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAchievementTemplate**
> DeleteAchievementTemplate($id, $cascade)

Delete an achievement template

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

# **GetAchievement**
> AchievementDefinitionResource GetAchievement($name)

Get a single achievement definition


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the achievement | 

### Return type

[**AchievementDefinitionResource**](AchievementDefinitionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAchievementTemplate**
> TemplateResource GetAchievementTemplate($id)

Get a single achievement template


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

# **GetAchievementTemplates**
> PageResourceTemplateResource GetAchievementTemplates($size, $page, $order)

List and search achievement templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAchievementTriggers**
> []BreTriggerResource GetAchievementTriggers()

Get the list of triggers that can be used to trigger an achievement progress update


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]BreTriggerResource**](BreTriggerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAchievements**
> PageResourceAchievementDefinitionResource GetAchievements($filterTagset, $filterName, $filterHidden, $filterDerived, $size, $page, $order)

Get all achievement definitions in the system


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTagset** | **string**| Filter for achievements with specified tags (separated by comma) | [optional] 
 **filterName** | **string**| Filter for achievements whose name contains a string | [optional] 
 **filterHidden** | **bool**| Filter for achievements that are hidden or not | [optional] 
 **filterDerived** | **bool**| Filter for achievements that are derived from other services | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageResourceAchievementDefinitionResource**](PageResource«AchievementDefinitionResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDerivedAchievements**
> []AchievementDefinitionResource GetDerivedAchievements($name)

Get a list of derived achievements

Used by other services that depend on achievements


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the derived achievement | 

### Return type

[**[]AchievementDefinitionResource**](AchievementDefinitionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserAchievementProgress**
> UserAchievementGroupResource GetUserAchievementProgress($userId, $achievementName)

Retrieve progress on a given achievement for a given user

Assets will not be filled in on the resources returned. Use 'Get a single poll' to retrieve the full resource with assets for a given item as needed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user&#39;s id | 
 **achievementName** | **string**| The achievement&#39;s name | 

### Return type

[**UserAchievementGroupResource**](UserAchievementGroupResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserAchievementsProgress**
> PageResourceUserAchievementGroupResource GetUserAchievementsProgress($userId, $filterAchievementDerived, $filterAchievementTagset, $filterAchievementName, $size, $page)

Retrieve progress on achievements for a given user

Assets will not be filled in on the resources returned. Use 'Get a single poll' to retrieve the full resource with assets for a given item as needed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user&#39;s id | 
 **filterAchievementDerived** | **bool**| Filter for achievements that are derived from other services | [optional] 
 **filterAchievementTagset** | **string**| Filter for achievements with specified tags (separated by comma) | [optional] 
 **filterAchievementName** | **string**| Filter for achievements whose name contains a string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUserAchievementGroupResource**](PageResource«UserAchievementGroupResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersAchievementProgress**
> PageResourceUserAchievementGroupResource GetUsersAchievementProgress($achievementName, $filterAchievementDerived, $filterAchievementTagset, $filterAchievementName, $size, $page)

Retrieve progress on a given achievement for all users

Assets will not be filled in on the resources returned. Use 'Get single achievement progress for user' to retrieve the full resource with assets for a given user as needed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achievementName** | **string**| The achievement&#39;s name | 
 **filterAchievementDerived** | **bool**| Filter for achievements that are derived from other services | [optional] 
 **filterAchievementTagset** | **string**| Filter for achievements with specified tags (separated by comma) | [optional] 
 **filterAchievementName** | **string**| Filter for achievements whose name contains a string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUserAchievementGroupResource**](PageResource«UserAchievementGroupResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersAchievementsProgress**
> PageResourceUserAchievementGroupResource GetUsersAchievementsProgress($filterAchievementDerived, $filterAchievementTagset, $filterAchievementName, $size, $page)

Retrieve progress on achievements for all users

Assets will not be filled in on the resources returned. Use 'Get single achievement progress for user' to retrieve the full resource with assets for a given user as needed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAchievementDerived** | **bool**| Filter for achievements that are derived from other services | [optional] 
 **filterAchievementTagset** | **string**| Filter for achievements with specified tags (separated by comma) | [optional] 
 **filterAchievementName** | **string**| Filter for achievements whose name contains a string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUserAchievementGroupResource**](PageResource«UserAchievementGroupResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAchievement**
> AchievementDefinitionResource UpdateAchievement($name, $achievement)

Update an achievement definition

The existing generated rule, if any, will be deleted. A new rule will be created if a trigger event name is specified in the new version.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the achievement | 
 **achievement** | [**AchievementDefinitionResource**](AchievementDefinitionResource.md)| The achievement definition | [optional] 

### Return type

[**AchievementDefinitionResource**](AchievementDefinitionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAchievementProgress**
> UserAchievementGroupResource UpdateAchievementProgress($userId, $achievementName, $request)

Update or create an achievement progress record for a user

If no progress record yet exists for the user, it will be created. Otherwise it will be updated. If progress meets or exceeds the achievement's max_value it will be marked as earned and a BRE event will be triggered for the <code>BreAchievementEarnedTrigger</code>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user&#39;s id | 
 **achievementName** | **string**| The achievement&#39;s name | 
 **request** | [**AchievementProgressUpdateRequest**](AchievementProgressUpdateRequest.md)| The progress update details | [optional] 

### Return type

[**UserAchievementGroupResource**](UserAchievementGroupResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAchievementTemplate**
> TemplateResource UpdateAchievementTemplate($id, $template)

Update an achievement template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **template** | [**TemplateResource**](TemplateResource.md)| The updated template | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

