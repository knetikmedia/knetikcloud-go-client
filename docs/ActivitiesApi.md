# \ActivitiesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActivity**](ActivitiesApi.md#CreateActivity) | **Post** /activities | Create an activity
[**CreateActivityOccurrence**](ActivitiesApi.md#CreateActivityOccurrence) | **Post** /activity-occurrences | Create a new activity occurrence
[**DeleteActivity**](ActivitiesApi.md#DeleteActivity) | **Delete** /activities/{id} | Delete an activity
[**GetActivities**](ActivitiesApi.md#GetActivities) | **Get** /activities | List activity definitions
[**GetActivity**](ActivitiesApi.md#GetActivity) | **Get** /activities/{id} | Get a single activity
[**SetActivityOccurrenceResults**](ActivitiesApi.md#SetActivityOccurrenceResults) | **Post** /activity-occurrences/{activity_occurrence_id}/results | Sets the status of an activity occurrence to FINISHED and logs metrics
[**UpdateActivity**](ActivitiesApi.md#UpdateActivity) | **Put** /activities/{id} | Update an activity
[**UpdateActivityOccurrence**](ActivitiesApi.md#UpdateActivityOccurrence) | **Put** /activity-occurrences/{activity_occurrence_id}/status | Updated the status of an activity occurrence


# **CreateActivity**
> ActivityResource CreateActivity($activityResource)

Create an activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityResource** | [**ActivityResource**](ActivityResource.md)| The activity resource object | [optional] 

### Return type

[**ActivityResource**](ActivityResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateActivityOccurrence**
> ActivityOccurrenceResource CreateActivityOccurrence($test, $activityOccurrenceResource)

Create a new activity occurrence

Has to enforce extra rules if not used as an admin


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **test** | **bool**| if true, indicates that the occurrence should NOT be created. This can be used to test for eligibility and valid settings | [optional] [default to false]
 **activityOccurrenceResource** | [**ActivityOccurrenceResource**](ActivityOccurrenceResource.md)| The activity occurrence object | [optional] 

### Return type

[**ActivityOccurrenceResource**](ActivityOccurrenceResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteActivity**
> DeleteActivity($id)

Delete an activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the activity | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivities**
> PageResourceBareActivityResource GetActivities($filterTemplate, $filterName, $filterId, $size, $page, $order)

List activity definitions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTemplate** | **bool**| Filter for activities that are templates, or specifically not if false | [optional] 
 **filterName** | **string**| Filter for activities that have a name starting with specified string | [optional] 
 **filterId** | **string**| Filter for activities with an id in the given comma separated list of ids | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceBareActivityResource**](PageResource«BareActivityResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivity**
> ActivityResource GetActivity($id)

Get a single activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the activity | 

### Return type

[**ActivityResource**](ActivityResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetActivityOccurrenceResults**
> ActivityOccurrenceResults SetActivityOccurrenceResults($activityOccurrenceId, $activityOccurrenceResults)

Sets the status of an activity occurrence to FINISHED and logs metrics


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityOccurrenceId** | **int64**| The id of the activity occurrence | 
 **activityOccurrenceResults** | [**ActivityOccurrenceResults**](ActivityOccurrenceResults.md)| The activity occurrence object | [optional] 

### Return type

[**ActivityOccurrenceResults**](ActivityOccurrenceResults.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateActivity**
> ActivityResource UpdateActivity($id, $activityResource)

Update an activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the activity | 
 **activityResource** | [**ActivityResource**](ActivityResource.md)| The activity resource object | [optional] 

### Return type

[**ActivityResource**](ActivityResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateActivityOccurrence**
> UpdateActivityOccurrence($activityOccurrenceId, $activityCccurrenceStatus)

Updated the status of an activity occurrence

If setting to 'FINISHED' you must POST to /results instead to record the metrics and get synchronous reward results


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityOccurrenceId** | **int64**| The id of the activity occurrence | 
 **activityCccurrenceStatus** | **string**| The activity occurrence status object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

