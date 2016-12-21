# \CampaignsChallengesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChallengeActivityUsingPOST**](CampaignsChallengesApi.md#CreateChallengeActivityUsingPOST) | **Post** /challenges/{challenge_id}/activities | Create a challenge activity
[**CreateChallengeTemplateUsingPOST**](CampaignsChallengesApi.md#CreateChallengeTemplateUsingPOST) | **Post** /challenges/templates | Create a challenge template
[**CreateChallengeUsingPOST**](CampaignsChallengesApi.md#CreateChallengeUsingPOST) | **Post** /challenges | Create a challenge
[**DeleteChallengeActivityUsingDELETE**](CampaignsChallengesApi.md#DeleteChallengeActivityUsingDELETE) | **Delete** /challenges/{challenge_id}/activities/{activity_id} | Delete a challenge activity
[**DeleteChallengeEventUsingDELETE**](CampaignsChallengesApi.md#DeleteChallengeEventUsingDELETE) | **Delete** /challenges/events/{id} | Delete a challenge event
[**DeleteChallengeTemplateUsingDELETE**](CampaignsChallengesApi.md#DeleteChallengeTemplateUsingDELETE) | **Delete** /challenges/templates/{id} | Delete a challenge template
[**DeleteChallengeUsingDELETE**](CampaignsChallengesApi.md#DeleteChallengeUsingDELETE) | **Delete** /challenges/{id} | Delete a challenge
[**GetActivitiesUsingGET1**](CampaignsChallengesApi.md#GetActivitiesUsingGET1) | **Get** /challenges/{challenge_id}/activities | List and search challenge activities
[**GetChallengeActivityUsingGET**](CampaignsChallengesApi.md#GetChallengeActivityUsingGET) | **Get** /challenges/{challenge_id}/activities/{activity_id} | Get a single challenge activity
[**GetChallengeEventUsingGET**](CampaignsChallengesApi.md#GetChallengeEventUsingGET) | **Get** /challenges/events/{id} | Retrieve a single challenge event details
[**GetChallengeEventssUsingGET**](CampaignsChallengesApi.md#GetChallengeEventssUsingGET) | **Get** /challenges/events | Retrieve a list of challenge events
[**GetChallengeTemplateUsingGET**](CampaignsChallengesApi.md#GetChallengeTemplateUsingGET) | **Get** /challenges/templates/{id} | Get a single challenge template
[**GetChallengeTemplatesUsingGET**](CampaignsChallengesApi.md#GetChallengeTemplatesUsingGET) | **Get** /challenges/templates | List and search challenge templates
[**GetChallengeUsingGET**](CampaignsChallengesApi.md#GetChallengeUsingGET) | **Get** /challenges/{id} | Retrieve a single challenge details
[**GetChallengesUsingGET1**](CampaignsChallengesApi.md#GetChallengesUsingGET1) | **Get** /challenges | Retrieve a list of challenges
[**UpdateChallengeActivityUsingPUT**](CampaignsChallengesApi.md#UpdateChallengeActivityUsingPUT) | **Put** /challenges/{challenge_id}/activities/{activity_id} | Update a challenge activity
[**UpdateChallengeTemplateUsingPUT**](CampaignsChallengesApi.md#UpdateChallengeTemplateUsingPUT) | **Put** /challenges/templates/{id} | Update a challenge template
[**UpdateChallengeUsingPUT**](CampaignsChallengesApi.md#UpdateChallengeUsingPUT) | **Put** /challenges/{id} | Update a challenge&#39;s information


# **CreateChallengeActivityUsingPOST**
> ChallengeActivityResource CreateChallengeActivityUsingPOST($challengeId, $challengeActivityResource, $validateSettings)

Create a challenge activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeId** | **int64**| The challenge id | 
 **challengeActivityResource** | [**ChallengeActivityResource**](ChallengeActivityResource.md)| The challenge activity resource object | [optional] 
 **validateSettings** | **bool**| Whether to validate the settings being sent against the available settings on the base activity. | [optional] [default to false]

### Return type

[**ChallengeActivityResource**](ChallengeActivityResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChallengeTemplateUsingPOST**
> TemplateResource CreateChallengeTemplateUsingPOST($challengeTemplateResource)

Create a challenge template

Challenge Templates define a type of challenge and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeTemplateResource** | [**TemplateResource**](TemplateResource.md)| The challenge template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChallengeUsingPOST**
> ChallengeResource CreateChallengeUsingPOST($challengeResource)

Create a challenge

Challenges do not run on their own.  They must be added to a campaign before events will spawn.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeResource** | [**ChallengeResource**](ChallengeResource.md)| The challenge resource object | [optional] 

### Return type

[**ChallengeResource**](ChallengeResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChallengeActivityUsingDELETE**
> DeleteChallengeActivityUsingDELETE($activityId, $challengeId)

Delete a challenge activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **int64**| The activity id | 
 **challengeId** | **int64**| The challenge id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChallengeEventUsingDELETE**
> DeleteChallengeEventUsingDELETE($id)

Delete a challenge event


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge event id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChallengeTemplateUsingDELETE**
> DeleteChallengeTemplateUsingDELETE($id, $cascade)

Delete a challenge template

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

# **DeleteChallengeUsingDELETE**
> DeleteChallengeUsingDELETE($id)

Delete a challenge


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivitiesUsingGET1**
> PageResourceBareChallengeActivityResource GetActivitiesUsingGET1($challengeId, $size, $page, $order)

List and search challenge activities


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeId** | **int64**| The challenge id | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceBareChallengeActivityResource**](PageResource«BareChallengeActivityResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengeActivityUsingGET**
> ChallengeActivityResource GetChallengeActivityUsingGET($activityId)

Get a single challenge activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **int64**| The activity id | 

### Return type

[**ChallengeActivityResource**](ChallengeActivityResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengeEventUsingGET**
> ChallengeEventResource GetChallengeEventUsingGET($id)

Retrieve a single challenge event details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge event id | 

### Return type

[**ChallengeEventResource**](ChallengeEventResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengeEventssUsingGET**
> PageResourceChallengeEventResource GetChallengeEventssUsingGET($filterStartDate, $filterEndDate, $filterCampaigns, $filterChallenge, $size, $page, $order)

Retrieve a list of challenge events


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStartDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the event start date, a unix timestamp in seconds.  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 
 **filterEndDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the event end date, a unix timestamp in seconds.  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 
 **filterCampaigns** | **bool**| check only for events from currently running campaigns | [optional] 
 **filterChallenge** | **int64**| check only for events from the challenge specified by id | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceChallengeEventResource**](PageResource«ChallengeEventResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengeTemplateUsingGET**
> TemplateResource GetChallengeTemplateUsingGET($id)

Get a single challenge template


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

# **GetChallengeTemplatesUsingGET**
> PageResourceTemplateResource GetChallengeTemplatesUsingGET($size, $page, $order)

List and search challenge templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengeUsingGET**
> ChallengeResource GetChallengeUsingGET($id)

Retrieve a single challenge details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge id | 

### Return type

[**ChallengeResource**](ChallengeResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengesUsingGET1**
> PageResourceChallengeResource GetChallengesUsingGET1($filterTemplate, $filterActiveCampaign)

Retrieve a list of challenges


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTemplate** | **bool**| Filter for challenges that are not tied to campaigns (templates) | [optional] 
 **filterActiveCampaign** | **bool**| Filter for challenges that are tied to active campaigns | [optional] 

### Return type

[**PageResourceChallengeResource**](PageResource«ChallengeResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallengeActivityUsingPUT**
> UpdateChallengeActivityUsingPUT($activityId, $challengeId, $challengeActivityResource)

Update a challenge activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **int64**| The activity id | 
 **challengeId** | **int64**| The challenge id | 
 **challengeActivityResource** | [**ChallengeActivityResource**](ChallengeActivityResource.md)| The challenge activity resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallengeTemplateUsingPUT**
> UpdateChallengeTemplateUsingPUT($id, $challengeTemplateResource)

Update a challenge template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **challengeTemplateResource** | [**TemplateResource**](TemplateResource.md)| The challenge template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallengeUsingPUT**
> ChallengeResource UpdateChallengeUsingPUT($id, $challengeResource)

Update a challenge's information

If the challenge is a copy, changes will propagate to all the related challenges


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge id | 
 **challengeResource** | [**ChallengeResource**](ChallengeResource.md)| The challenge resource object | [optional] 

### Return type

[**ChallengeResource**](ChallengeResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

