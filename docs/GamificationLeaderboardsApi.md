# \GamificationLeaderboardsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLeaderboard**](GamificationLeaderboardsApi.md#GetLeaderboard) | **Get** /leaderboards/{context_type}/{context_id} | Retrieves leaderboard details and paginated entries
[**GetLeaderboardRank**](GamificationLeaderboardsApi.md#GetLeaderboardRank) | **Get** /leaderboards/{context_type}/{context_id}/users/{id}/rank | Retrieves a specific user entry with rank
[**GetLeaderboardStrategies**](GamificationLeaderboardsApi.md#GetLeaderboardStrategies) | **Get** /leaderboards/strategies | Get a list of available leaderboard strategy names


# **GetLeaderboard**
> LeaderboardResource GetLeaderboard($contextType, $contextId, $size, $page)

Retrieves leaderboard details and paginated entries

The context type identifies the type of entity (i.e., 'activity') being tracked on the leaderboard. The context ID is the unique ID of the actual entity tracked by the leaderboard.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextType** | **string**| The context type for the leaderboard | 
 **contextId** | **string**| The context id for the leaderboard | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**LeaderboardResource**](LeaderboardResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeaderboardRank**
> LeaderboardEntryResource GetLeaderboardRank($contextType, $contextId, $id)

Retrieves a specific user entry with rank

The context type identifies the type of entity (i.e., 'activity') being tracked on the leaderboard. The context ID is the unique ID of the actual entity tracked by the leaderboard


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextType** | **string**| The context type for the leaderboard | 
 **contextId** | **string**| The context id for the leaderboard | 
 **id** | **string**| The id of a user | 

### Return type

[**LeaderboardEntryResource**](LeaderboardEntryResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeaderboardStrategies**
> []string GetLeaderboardStrategies()

Get a list of available leaderboard strategy names


### Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

