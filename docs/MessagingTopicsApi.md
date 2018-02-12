# \MessagingTopicsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableTopicSubscriber**](MessagingTopicsApi.md#DisableTopicSubscriber) | **Put** /messaging/topics/{id}/subscribers/{user_id}/disabled | Enable or disable messages for a user
[**GetTopicSubscriber**](MessagingTopicsApi.md#GetTopicSubscriber) | **Get** /messaging/topics/{id}/subscribers/{user_id} | Get a subscriber to a topic
[**GetTopicSubscribers**](MessagingTopicsApi.md#GetTopicSubscribers) | **Get** /messaging/topics/{id}/subscribers | Get all subscribers to a topic
[**GetUserTopics**](MessagingTopicsApi.md#GetUserTopics) | **Get** /users/{id}/topics | Get all messaging topics for a given user


# **DisableTopicSubscriber**
> DisableTopicSubscriber(ctx, ctx, id, userId, disabled)
Enable or disable messages for a user

Useful for opt-out options on a single topic. Consider multiple topics for multiple opt-out options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the topic | 
  **userId** | **string**| The id of the subscriber or &#39;me&#39; | 
  **disabled** | [**ValueWrapperboolean**](ValueWrapperboolean.md)| disabled | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTopicSubscriber**
> TopicSubscriberResource GetTopicSubscriber(ctx, ctx, id, userId)
Get a subscriber to a topic

<b>Permissions Needed:</b> TOPICS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the topic | 
  **userId** | **string**| The id of the subscriber or &#39;me&#39; | 

### Return type

[**TopicSubscriberResource**](TopicSubscriberResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTopicSubscribers**
> PageResourceTopicSubscriberResource GetTopicSubscribers(ctx, ctx, id)
Get all subscribers to a topic

<b>Permissions Needed:</b> TOPICS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the topic | 

### Return type

[**PageResourceTopicSubscriberResource**](PageResource«TopicSubscriberResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserTopics**
> PageResourceTopicResource GetUserTopics(ctx, ctx, id)
Get all messaging topics for a given user

<b>Permissions Needed:</b> TOPICS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the user or &#39;me&#39; | 

### Return type

[**PageResourceTopicResource**](PageResource«TopicResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

