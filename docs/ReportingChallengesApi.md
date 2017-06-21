# \ReportingChallengesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChallengeEventLeaderboard**](ReportingChallengesApi.md#GetChallengeEventLeaderboard) | **Get** /reporting/events/leaderboard | Retrieve a challenge event leaderboard details
[**GetChallengeEventParticipants**](ReportingChallengesApi.md#GetChallengeEventParticipants) | **Get** /reporting/events/participants | Retrieve a challenge event participant details


# **GetChallengeEventLeaderboard**
> PageResourceChallengeEventParticipantResource GetChallengeEventLeaderboard($filterEvent, $size, $page, $order)

Retrieve a challenge event leaderboard details

Lists all leaderboard entries with additional user details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEvent** | **int64**| A sepecific challenge event id | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] 

### Return type

[**PageResourceChallengeEventParticipantResource**](PageResource«ChallengeEventParticipantResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengeEventParticipants**
> PageResourceChallengeEventParticipantResource GetChallengeEventParticipants($filterEvent, $size, $page, $order)

Retrieve a challenge event participant details

Lists all user submitted scores sorted by value, including those that do not apear in the leaderboard due to value or aggregation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEvent** | **int64**| A sepecific challenge event id | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] 

### Return type

[**PageResourceChallengeEventParticipantResource**](PageResource«ChallengeEventParticipantResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

