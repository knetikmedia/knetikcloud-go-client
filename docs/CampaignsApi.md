# \CampaignsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChallengeToCampaign**](CampaignsApi.md#AddChallengeToCampaign) | **Post** /campaigns/{id}/challenges | Add a challenge to a campaign
[**CreateCampaign**](CampaignsApi.md#CreateCampaign) | **Post** /campaigns | Create a campaign
[**CreateCampaignTemplate**](CampaignsApi.md#CreateCampaignTemplate) | **Post** /campaigns/templates | Create a campaign template
[**DeleteCampaign**](CampaignsApi.md#DeleteCampaign) | **Delete** /campaigns/{id} | Delete a campaign
[**DeleteCampaignTemplate**](CampaignsApi.md#DeleteCampaignTemplate) | **Delete** /campaigns/templates/{id} | Delete a campaign template
[**GetCampaign**](CampaignsApi.md#GetCampaign) | **Get** /campaigns/{id} | Returns a single campaign
[**GetCampaignChallenges**](CampaignsApi.md#GetCampaignChallenges) | **Get** /campaigns/{id}/challenges | List the challenges associated with a campaign
[**GetCampaignTemplate**](CampaignsApi.md#GetCampaignTemplate) | **Get** /campaigns/templates/{id} | Get a single campaign template
[**GetCampaignTemplates**](CampaignsApi.md#GetCampaignTemplates) | **Get** /campaigns/templates | List and search campaign templates
[**GetCampaigns**](CampaignsApi.md#GetCampaigns) | **Get** /campaigns | List and search campaigns
[**RemoveChallengeFromCampaign**](CampaignsApi.md#RemoveChallengeFromCampaign) | **Delete** /campaigns/{campaign_id}/challenges/{id} | Remove a challenge from a campaign
[**UpdateCampaign**](CampaignsApi.md#UpdateCampaign) | **Put** /campaigns/{id} | Update a campaign
[**UpdateCampaignTemplate**](CampaignsApi.md#UpdateCampaignTemplate) | **Put** /campaigns/templates/{id} | Update an campaign template


# **AddChallengeToCampaign**
> AddChallengeToCampaign($id, $challengeId)

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

# **CreateCampaign**
> CampaignResource CreateCampaign($campaignResource)

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

# **CreateCampaignTemplate**
> TemplateResource CreateCampaignTemplate($campaignTemplateResource)

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

# **DeleteCampaign**
> DeleteCampaign($id)

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

# **DeleteCampaignTemplate**
> DeleteCampaignTemplate($id, $cascade)

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

# **GetCampaign**
> CampaignResource GetCampaign($id)

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

# **GetCampaignChallenges**
> PageResourceChallengeResource GetCampaignChallenges($id)

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

# **GetCampaignTemplate**
> TemplateResource GetCampaignTemplate($id)

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

# **GetCampaignTemplates**
> PageResourceTemplateResource GetCampaignTemplates($size, $page, $order)

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

# **GetCampaigns**
> PageResourceCampaignResource GetCampaigns($filterActive, $size, $page, $order)

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

# **RemoveChallengeFromCampaign**
> RemoveChallengeFromCampaign($campaignId, $id)

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

# **UpdateCampaign**
> CampaignResource UpdateCampaign($id, $campaignResource)

Update a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The campaign id | 
 **campaignResource** | [**CampaignResource**](CampaignResource.md)| The campaign resource object | [optional] 

### Return type

[**CampaignResource**](CampaignResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaignTemplate**
> TemplateResource UpdateCampaignTemplate($id, $campaignTemplateResource)

Update an campaign template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **campaignTemplateResource** | [**TemplateResource**](TemplateResource.md)| The campaign template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

