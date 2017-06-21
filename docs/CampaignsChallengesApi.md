# \CampaignsChallengesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChallenge**](CampaignsChallengesApi.md#CreateChallenge) | **Post** /challenges | Create a challenge
[**CreateChallengeActivity**](CampaignsChallengesApi.md#CreateChallengeActivity) | **Post** /challenges/{challenge_id}/activities | Create a challenge activity
[**CreateChallengeActivityTemplate**](CampaignsChallengesApi.md#CreateChallengeActivityTemplate) | **Post** /challenge-activities/templates | Create a challenge activity template
[**CreateChallengeTemplate**](CampaignsChallengesApi.md#CreateChallengeTemplate) | **Post** /challenges/templates | Create a challenge template
[**DeleteChallenge**](CampaignsChallengesApi.md#DeleteChallenge) | **Delete** /challenges/{id} | Delete a challenge
[**DeleteChallengeActivity**](CampaignsChallengesApi.md#DeleteChallengeActivity) | **Delete** /challenges/{challenge_id}/activities/{activity_id} | Delete a challenge activity
[**DeleteChallengeActivityTemplate**](CampaignsChallengesApi.md#DeleteChallengeActivityTemplate) | **Delete** /challenge-activities/templates/{id} | Delete a challenge activity template
[**DeleteChallengeEvent**](CampaignsChallengesApi.md#DeleteChallengeEvent) | **Delete** /challenges/events/{id} | Delete a challenge event
[**DeleteChallengeTemplate**](CampaignsChallengesApi.md#DeleteChallengeTemplate) | **Delete** /challenges/templates/{id} | Delete a challenge template
[**GetChallenge**](CampaignsChallengesApi.md#GetChallenge) | **Get** /challenges/{id} | Retrieve a challenge
[**GetChallengeActivities**](CampaignsChallengesApi.md#GetChallengeActivities) | **Get** /challenges/{challenge_id}/activities | List and search challenge activities
[**GetChallengeActivity**](CampaignsChallengesApi.md#GetChallengeActivity) | **Get** /challenges/{challenge_id}/activities/{activity_id} | Get a single challenge activity
[**GetChallengeActivityTemplate**](CampaignsChallengesApi.md#GetChallengeActivityTemplate) | **Get** /challenge-activities/templates/{id} | Get a single challenge activity template
[**GetChallengeActivityTemplates**](CampaignsChallengesApi.md#GetChallengeActivityTemplates) | **Get** /challenge-activities/templates | List and search challenge activity templates
[**GetChallengeEvent**](CampaignsChallengesApi.md#GetChallengeEvent) | **Get** /challenges/events/{id} | Retrieve a single challenge event details
[**GetChallengeEvents**](CampaignsChallengesApi.md#GetChallengeEvents) | **Get** /challenges/events | Retrieve a list of challenge events
[**GetChallengeTemplate**](CampaignsChallengesApi.md#GetChallengeTemplate) | **Get** /challenges/templates/{id} | Get a single challenge template
[**GetChallengeTemplates**](CampaignsChallengesApi.md#GetChallengeTemplates) | **Get** /challenges/templates | List and search challenge templates
[**GetChallenges**](CampaignsChallengesApi.md#GetChallenges) | **Get** /challenges | Retrieve a list of challenges
[**UpdateChallenge**](CampaignsChallengesApi.md#UpdateChallenge) | **Put** /challenges/{id} | Update a challenge
[**UpdateChallengeActivity**](CampaignsChallengesApi.md#UpdateChallengeActivity) | **Put** /challenges/{challenge_id}/activities/{activity_id} | Update a challenge activity
[**UpdateChallengeActivityTemplate**](CampaignsChallengesApi.md#UpdateChallengeActivityTemplate) | **Put** /challenge-activities/templates/{id} | Update an challenge activity template
[**UpdateChallengeTemplate**](CampaignsChallengesApi.md#UpdateChallengeTemplate) | **Put** /challenges/templates/{id} | Update a challenge template


# **CreateChallenge**
> ChallengeResource CreateChallenge($challengeResource)

Create a challenge

Challenges do not run on their own.  They must be added to a campaign before events will spawn.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeResource** | [**ChallengeResource**](ChallengeResource.md)| The challenge resource object | [optional] 

### Return type

[**ChallengeResource**](ChallengeResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChallengeActivity**
> ChallengeActivityResource CreateChallengeActivity($challengeId, $challengeActivityResource, $validateSettings)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChallengeActivityTemplate**
> TemplateResource CreateChallengeActivityTemplate($challengeActivityTemplateResource)

Create a challenge activity template

Challenge Activity Templates define a type of challenge activity and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeActivityTemplateResource** | [**TemplateResource**](TemplateResource.md)| The challengeActivity template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChallengeTemplate**
> TemplateResource CreateChallengeTemplate($challengeTemplateResource)

Create a challenge template

Challenge Templates define a type of challenge and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challengeTemplateResource** | [**TemplateResource**](TemplateResource.md)| The challenge template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChallenge**
> DeleteChallenge($id)

Delete a challenge


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChallengeActivity**
> DeleteChallengeActivity($activityId, $challengeId)

Delete a challenge activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **int64**| The activity id | 
 **challengeId** | **int64**| The challenge id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChallengeActivityTemplate**
> DeleteChallengeActivityTemplate($id, $cascade)

Delete a challenge activity template

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

# **DeleteChallengeEvent**
> DeleteChallengeEvent($id)

Delete a challenge event


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge event id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChallengeTemplate**
> DeleteChallengeTemplate($id, $cascade)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallenge**
> ChallengeResource GetChallenge($id)

Retrieve a challenge


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

# **GetChallengeActivities**
> PageResourceBareChallengeActivityResource GetChallengeActivities($challengeId, $size, $page, $order)

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

# **GetChallengeActivity**
> ChallengeActivityResource GetChallengeActivity($activityId)

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

# **GetChallengeActivityTemplate**
> TemplateResource GetChallengeActivityTemplate($id)

Get a single challenge activity template


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

# **GetChallengeActivityTemplates**
> PageResourceTemplateResource GetChallengeActivityTemplates($size, $page, $order)

List and search challenge activity templates


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

# **GetChallengeEvent**
> ChallengeEventResource GetChallengeEvent($id)

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

# **GetChallengeEvents**
> PageResourceChallengeEventResource GetChallengeEvents($filterStartDate, $filterEndDate, $filterCampaigns, $filterChallenge, $size, $page, $order)

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

# **GetChallengeTemplate**
> TemplateResource GetChallengeTemplate($id)

Get a single challenge template


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

# **GetChallengeTemplates**
> PageResourceTemplateResource GetChallengeTemplates($size, $page, $order)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallenges**
> PageResourceChallengeResource GetChallenges($filterTemplate, $filterActiveCampaign, $filterStartDate, $filterEndDate, $size, $page, $order)

Retrieve a list of challenges


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterTemplate** | **bool**| Filter for challenges that are not tied to campaigns (templates) | [optional] 
 **filterActiveCampaign** | **bool**| Filter for challenges that are tied to active campaigns | [optional] 
 **filterStartDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the challenge start date, a unix timestamp in seconds.  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 
 **filterEndDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the challenge end date, a unix timestamp in seconds.  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceChallengeResource**](PageResource«ChallengeResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallenge**
> ChallengeResource UpdateChallenge($id, $challengeResource)

Update a challenge

If the challenge is a copy, changes will propagate to all the related challenges


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The challenge id | 
 **challengeResource** | [**ChallengeResource**](ChallengeResource.md)| The challenge resource object | [optional] 

### Return type

[**ChallengeResource**](ChallengeResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallengeActivity**
> ChallengeActivityResource UpdateChallengeActivity($activityId, $challengeId, $challengeActivityResource)

Update a challenge activity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **int64**| The activity id | 
 **challengeId** | **int64**| The challenge id | 
 **challengeActivityResource** | [**ChallengeActivityResource**](ChallengeActivityResource.md)| The challenge activity resource object | [optional] 

### Return type

[**ChallengeActivityResource**](ChallengeActivityResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallengeActivityTemplate**
> TemplateResource UpdateChallengeActivityTemplate($id, $challengeActivityTemplateResource)

Update an challenge activity template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **challengeActivityTemplateResource** | [**TemplateResource**](TemplateResource.md)| The challengeActivity template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallengeTemplate**
> TemplateResource UpdateChallengeTemplate($id, $challengeTemplateResource)

Update a challenge template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **challengeTemplateResource** | [**TemplateResource**](TemplateResource.md)| The challenge template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

