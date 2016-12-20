# \CampaignsRewardsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRewardSetUsingPOST**](CampaignsRewardsApi.md#CreateRewardSetUsingPOST) | **Post** /rewards | Create a reward set
[**DeleteRewardSetUsingDELETE**](CampaignsRewardsApi.md#DeleteRewardSetUsingDELETE) | **Delete** /rewards/{id} | Delete a reward set
[**GetRewardSetUsingGET**](CampaignsRewardsApi.md#GetRewardSetUsingGET) | **Get** /rewards/{id} | Get a single reward set
[**GetRewardSetsUsingGET**](CampaignsRewardsApi.md#GetRewardSetsUsingGET) | **Get** /rewards | List and search reward sets
[**UpdateRewardSetUsingPUT**](CampaignsRewardsApi.md#UpdateRewardSetUsingPUT) | **Put** /rewards/{id} | Update a reward set


# **CreateRewardSetUsingPOST**
> RewardSetResource CreateRewardSetUsingPOST($rewardSetResource)

Create a reward set


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewardSetResource** | [**RewardSetResource**](RewardSetResource.md)| The reward set resource object | [optional] 

### Return type

[**RewardSetResource**](RewardSetResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRewardSetUsingDELETE**
> DeleteRewardSetUsingDELETE($id)

Delete a reward set


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The reward id | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRewardSetUsingGET**
> RewardSetResource GetRewardSetUsingGET($id)

Get a single reward set


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The reward id | 

### Return type

[**RewardSetResource**](RewardSetResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRewardSetsUsingGET**
> PageRewardSetResource GetRewardSetsUsingGET($size, $page, $order)

List and search reward sets


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageRewardSetResource**](Page«RewardSetResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRewardSetUsingPUT**
> UpdateRewardSetUsingPUT($id, $rewardSetResource)

Update a reward set


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The reward id | 
 **rewardSetResource** | [**RewardSetResource**](RewardSetResource.md)| The reward set resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

