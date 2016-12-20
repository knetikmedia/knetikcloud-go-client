# \CampaignsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChallengesUsingPOST**](CampaignsApi.md#AddChallengesUsingPOST) | **Post** /campaigns/{id}/challenges | Add a challenge to a campaign
[**CreateCampaignTemplateUsingPOST**](CampaignsApi.md#CreateCampaignTemplateUsingPOST) | **Post** /campaigns/templates | Create a campaign template
[**CreateCampaignUsingPOST**](CampaignsApi.md#CreateCampaignUsingPOST) | **Post** /campaigns | Create a campaign
[**DeleteCampaignTemplateUsingDELETE**](CampaignsApi.md#DeleteCampaignTemplateUsingDELETE) | **Delete** /campaigns/templates/{id} | Delete a campaign template
[**DeleteCampaignUsingDELETE**](CampaignsApi.md#DeleteCampaignUsingDELETE) | **Delete** /campaigns/{id} | Delete a campaign
[**GetCampaignTemplateUsingGET**](CampaignsApi.md#GetCampaignTemplateUsingGET) | **Get** /campaigns/templates/{id} | Get a single campaign template
[**GetCampaignTemplatesUsingGET**](CampaignsApi.md#GetCampaignTemplatesUsingGET) | **Get** /campaigns/templates | List and search campaign templates
[**GetCampaignUsingGET**](CampaignsApi.md#GetCampaignUsingGET) | **Get** /campaigns/{id} | Returns a single campaign
[**GetCampaignsUsingGET**](CampaignsApi.md#GetCampaignsUsingGET) | **Get** /campaigns | List and search campaigns
[**GetChallengesUsingGET**](CampaignsApi.md#GetChallengesUsingGET) | **Get** /campaigns/{id}/challenges | List the challenges associated with a campaign
[**RemoveChallengeUsingDELETE**](CampaignsApi.md#RemoveChallengeUsingDELETE) | **Delete** /campaigns/{campaign_id}/challenges/{id} | Remove a challenge from a campaign
[**UpdateCampaignTemplateUsingPUT**](CampaignsApi.md#UpdateCampaignTemplateUsingPUT) | **Put** /campaigns/templates/{id} | Update an campaign template
[**UpdateCampaignUsingPUT**](CampaignsApi.md#UpdateCampaignUsingPUT) | **Put** /campaigns/{id} | Update a campaign


# **AddChallengesUsingPOST**
> AddChallengesUsingPOST($id, $challengeId)

Add a challenge to a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the campaign | 
 **challengeId** | **int64**| The id of the challenge | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCampaignTemplateUsingPOST**
> TemplateResource CreateCampaignTemplateUsingPOST($campaignTemplateResource)

Create a campaign template

Campaign Templates define a type of campaign and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignTemplateResource** | [**TemplateResource**](TemplateResource.md)| The campaign template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCampaignUsingPOST**
> CampaignResource CreateCampaignUsingPOST($campaignResource)

Create a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignResource** | [**CampaignResource**](CampaignResource.md)| The campaign resource object | [optional] 

### Return type

[**CampaignResource**](CampaignResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCampaignTemplateUsingDELETE**
> DeleteCampaignTemplateUsingDELETE($id, $cascade)

Delete a campaign template

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

# **DeleteCampaignUsingDELETE**
> DeleteCampaignUsingDELETE($id)

Delete a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The campaign id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignTemplateUsingGET**
> TemplateResource GetCampaignTemplateUsingGET($id)

Get a single campaign template


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

# **GetCampaignTemplatesUsingGET**
> PageResourceTemplateResource GetCampaignTemplatesUsingGET($size, $page, $order)

List and search campaign templates


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

# **GetCampaignUsingGET**
> CampaignResource GetCampaignUsingGET($id)

Returns a single campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The campaign id | 

### Return type

[**CampaignResource**](CampaignResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsUsingGET**
> PageResourceCampaignResource GetCampaignsUsingGET($filterActive, $size, $page, $order)

List and search campaigns


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterActive** | **bool**| Filter for campaigns that are active | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceCampaignResource**](PageResource«CampaignResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChallengesUsingGET**
> PageResourceChallengeResource GetChallengesUsingGET($id)

List the challenges associated with a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The campaign id | 

### Return type

[**PageResourceChallengeResource**](PageResource«ChallengeResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveChallengeUsingDELETE**
> RemoveChallengeUsingDELETE($campaignId, $id)

Remove a challenge from a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **int64**| The campaign id | 
 **id** | **int64**| The challenge id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaignTemplateUsingPUT**
> UpdateCampaignTemplateUsingPUT($id, $campaignTemplateResource)

Update an campaign template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **campaignTemplateResource** | [**TemplateResource**](TemplateResource.md)| The campaign template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaignUsingPUT**
> UpdateCampaignUsingPUT($id, $campaignResource)

Update a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The campaign id | 
 **campaignResource** | [**CampaignResource**](CampaignResource.md)| The campaign resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
