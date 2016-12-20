# \ActivitiesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteActivityOccurrenceUsingPUT**](ActivitiesApi.md#CompleteActivityOccurrenceUsingPUT) | **Put** /activity-occurrences/{activity_occurrence_id}/status | Updated the status of an activity occurrence
[**CreateActivityOccurrenceUsingPOST**](ActivitiesApi.md#CreateActivityOccurrenceUsingPOST) | **Post** /activity-occurrences | Create a new activity occurrence
[**CreateActivityUsingPOST**](ActivitiesApi.md#CreateActivityUsingPOST) | **Post** /activities | Create an activity
[**DeleteActivityUsingDELETE**](ActivitiesApi.md#DeleteActivityUsingDELETE) | **Delete** /activities/{id} | Delete an activity
[**GetActivitiesUsingGET**](ActivitiesApi.md#GetActivitiesUsingGET) | **Get** /activities | List activity definitions
[**GetActivityUsingGET**](ActivitiesApi.md#GetActivityUsingGET) | **Get** /activities/{id} | Get a single activity
[**PostResultsUsingPOST**](ActivitiesApi.md#PostResultsUsingPOST) | **Post** /activity-occurrences/{activity_occurrence_id}/results | Sets the status of an activity occurrence to FINISHED and logs metrics
[**UpdateActivityUsingPUT**](ActivitiesApi.md#UpdateActivityUsingPUT) | **Put** /activities/{id} | Update an activity


# **CompleteActivityOccurrenceUsingPUT**
> CompleteActivityOccurrenceUsingPUT($activityOccurrenceId, $activityCccurrenceStatus)

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateActivityOccurrenceUsingPOST**
> ActivityOccurrenceResource CreateActivityOccurrenceUsingPOST($test, $activityOccurrenceResource)

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateActivityUsingPOST**
> ActivityResource CreateActivityUsingPOST($activityResource)

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteActivityUsingDELETE**
> DeleteActivityUsingDELETE($id)

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivitiesUsingGET**
> PageBareActivityResource GetActivitiesUsingGET($filterTemplate, $size, $page, $order)

List activity definitions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTemplate** | **bool**| Filter for activities that are templates, or specifically not if false | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageBareActivityResource**](Page«BareActivityResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivityUsingGET**
> ActivityResource GetActivityUsingGET($id)

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostResultsUsingPOST**
> ActivityOccurrenceResults PostResultsUsingPOST($activityOccurrenceId, $activityOccurrenceResults)

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateActivityUsingPUT**
> UpdateActivityUsingPUT($id, $activityResource)

Update an activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the activity | 
 **activityResource** | [**ActivityResource**](ActivityResource.md)| The activity resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

