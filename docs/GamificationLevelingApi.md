# \GamificationLevelingApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserLevelExperienceUsingPUT**](GamificationLevelingApi.md#ChangeUserLevelExperienceUsingPUT) | **Put** /users/{user_id}/leveling/{name} | Update or create a leveling progress record for a user
[**CreateLevelUsingPOST**](GamificationLevelingApi.md#CreateLevelUsingPOST) | **Post** /leveling | Create a level schema
[**DeleteLevelUsingDELETE**](GamificationLevelingApi.md#DeleteLevelUsingDELETE) | **Delete** /leveling/{name} | Delete a level
[**GetAvailableTriggersUsingGET1**](GamificationLevelingApi.md#GetAvailableTriggersUsingGET1) | **Get** /leveling/triggers | Get the list of triggers that can be used to trigger a leveling progress update
[**GetLevelUsingGET**](GamificationLevelingApi.md#GetLevelUsingGET) | **Get** /leveling/{name} | Retrieve a particular level
[**GetLevelsUsingGET**](GamificationLevelingApi.md#GetLevelsUsingGET) | **Get** /leveling | List and search levels
[**GetUserLevelUsingGET**](GamificationLevelingApi.md#GetUserLevelUsingGET) | **Get** /users/{user_id}/leveling/{name} | Get a user&#39;s progress for a given level schema
[**GetUserLevelsUsingGET**](GamificationLevelingApi.md#GetUserLevelsUsingGET) | **Get** /users/{user_id}/leveling | Get a user&#39;s progress for all level schemas
[**UpdateLevelUsingPUT**](GamificationLevelingApi.md#UpdateLevelUsingPUT) | **Put** /leveling/{name} | Update a level


# **ChangeUserLevelExperienceUsingPUT**
> ChangeUserLevelExperienceUsingPUT($userId, $name, $progress)

Update or create a leveling progress record for a user

If no progress record yet exists for the user, it will be created. Otherwise it will be updated. If progress meets or exceeds the level's max_value it will be marked as earned and a BRE event will be triggered for the <code>BreAchievementEarnedTrigger</code>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **name** | **string**| The level schema name | 
 **progress** | **int32**| The current progress amount | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLevelUsingPOST**
> LevelingResource CreateLevelUsingPOST($level)

Create a level schema


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | [**LevelingResource**](LevelingResource.md)| The level schema definition | [optional] 

### Return type

[**LevelingResource**](LevelingResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLevelUsingDELETE**
> DeleteLevelUsingDELETE($name)

Delete a level


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The level schema name | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableTriggersUsingGET1**
> []BreTriggerResource GetAvailableTriggersUsingGET1()

Get the list of triggers that can be used to trigger a leveling progress update


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]BreTriggerResource**](BreTriggerResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLevelUsingGET**
> LevelingResource GetLevelUsingGET($name)

Retrieve a particular level


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The level schema name | 

### Return type

[**LevelingResource**](LevelingResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLevelsUsingGET**
> PageResourceLevelingResource GetLevelsUsingGET($filterName, $size, $page, $order)

List and search levels

Get a list of levels schemas with optional filtering


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **string**| Filter for level schemas whose name contains a given string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageResourceLevelingResource**](PageResource«LevelingResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLevelUsingGET**
> UserLevelingResource GetUserLevelUsingGET($userId, $name)

Get a user's progress for a given level schema


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **name** | **string**| The level schema name | 

### Return type

[**UserLevelingResource**](UserLevelingResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLevelsUsingGET**
> PageResourceUserLevelingResource GetUserLevelsUsingGET($userId)

Get a user's progress for all level schemas


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 

### Return type

[**PageResourceUserLevelingResource**](PageResource«UserLevelingResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLevelUsingPUT**
> UpdateLevelUsingPUT($name, $newLevel)

Update a level


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The level schema name | 
 **newLevel** | [**LevelingResource**](LevelingResource.md)| The level schema definition | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

