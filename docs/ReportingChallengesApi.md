# \ReportingChallengesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChallengeEventLeaderboard**](ReportingChallengesApi.md#GetChallengeEventLeaderboard) | **Get** /reporting/events/leaderboard | Retrieve a challenge event leaderboard details
[**GetChallengeEventParticipants**](ReportingChallengesApi.md#GetChallengeEventParticipants) | **Get** /reporting/events/participants | Retrieve a challenge event participant details


# **GetChallengeEventLeaderboard**
> PageResourceChallengeEventParticipantResource GetChallengeEventLeaderboard($filterEvent)

Retrieve a challenge event leaderboard details

Lists all leaderboard entries with additional user details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEvent** | **int64**| A sepecific challenge event id | [optional] 

### Return type

[**PageResourceChallengeEventParticipantResource**](PageResource«ChallengeEventParticipantResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengeEventParticipants**
> PageResourceChallengeEventParticipantResource GetChallengeEventParticipants($filterEvent)

Retrieve a challenge event participant details

Lists all user submitted scores sorted by value, including those that do not apear in the leaderboard due to value or aggregation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterEvent** | **int64**| A sepecific challenge event id | [optional] 

### Return type

[**PageResourceChallengeEventParticipantResource**](PageResource«ChallengeEventParticipantResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

