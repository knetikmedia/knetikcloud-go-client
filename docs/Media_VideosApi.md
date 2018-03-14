# \Media_VideosApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToVideoWhitelist**](Media_VideosApi.md#AddUserToVideoWhitelist) | **Post** /media/videos/{id}/whitelist | Adds a user to a video&#39;s whitelist
[**AddVideo**](Media_VideosApi.md#AddVideo) | **Post** /media/videos | Adds a new video in the system
[**AddVideoComment**](Media_VideosApi.md#AddVideoComment) | **Post** /media/videos/{video_id}/comments | Add a new video comment
[**AddVideoContributor**](Media_VideosApi.md#AddVideoContributor) | **Post** /media/videos/{video_id}/contributors | Adds a contributor to a video
[**AddVideoFlag**](Media_VideosApi.md#AddVideoFlag) | **Post** /media/videos/{video_id}/moderation | Add a new flag
[**AddVideoRelationships**](Media_VideosApi.md#AddVideoRelationships) | **Post** /media/videos/{video_id}/related | Adds one or more existing videos as related to this one
[**CreateVideoDisposition**](Media_VideosApi.md#CreateVideoDisposition) | **Post** /media/videos/{video_id}/dispositions | Create a video disposition
[**CreateVideoTemplate**](Media_VideosApi.md#CreateVideoTemplate) | **Post** /media/videos/templates | Create a video template
[**DeleteVideo**](Media_VideosApi.md#DeleteVideo) | **Delete** /media/videos/{id} | Deletes a video from the system if no resources are attached to it
[**DeleteVideoComment**](Media_VideosApi.md#DeleteVideoComment) | **Delete** /media/videos/{video_id}/comments/{id} | Delete a video comment
[**DeleteVideoDisposition**](Media_VideosApi.md#DeleteVideoDisposition) | **Delete** /media/videos/{video_id}/dispositions/{disposition_id} | Delete a video disposition
[**DeleteVideoFlag**](Media_VideosApi.md#DeleteVideoFlag) | **Delete** /media/videos/{video_id}/moderation | Delete a flag
[**DeleteVideoRelationship**](Media_VideosApi.md#DeleteVideoRelationship) | **Delete** /media/videos/{video_id}/related/{id} | Delete a video&#39;s relationship
[**DeleteVideoTemplate**](Media_VideosApi.md#DeleteVideoTemplate) | **Delete** /media/videos/templates/{id} | Delete a video template
[**GetUserVideos**](Media_VideosApi.md#GetUserVideos) | **Get** /users/{user_id}/videos | Get user videos
[**GetVideo**](Media_VideosApi.md#GetVideo) | **Get** /media/videos/{id} | Loads a specific video details
[**GetVideoComments**](Media_VideosApi.md#GetVideoComments) | **Get** /media/videos/{video_id}/comments | Returns a page of comments for a video
[**GetVideoDispositions**](Media_VideosApi.md#GetVideoDispositions) | **Get** /media/videos/{video_id}/dispositions | Returns a page of dispositions for a video
[**GetVideoRelationships**](Media_VideosApi.md#GetVideoRelationships) | **Get** /media/videos/{video_id}/related | Returns a page of video relationships
[**GetVideoTemplate**](Media_VideosApi.md#GetVideoTemplate) | **Get** /media/videos/templates/{id} | Get a single video template
[**GetVideoTemplates**](Media_VideosApi.md#GetVideoTemplates) | **Get** /media/videos/templates | List and search video templates
[**GetVideos**](Media_VideosApi.md#GetVideos) | **Get** /media/videos | Search videos using the documented filters
[**RemoveUserFromVideoWhitelist**](Media_VideosApi.md#RemoveUserFromVideoWhitelist) | **Delete** /media/videos/{video_id}/whitelist/{id} | Removes a user from a video&#39;s whitelist
[**RemoveVideoContributor**](Media_VideosApi.md#RemoveVideoContributor) | **Delete** /media/videos/{video_id}/contributors/{id} | Removes a contributor from a video
[**UpdateVideo**](Media_VideosApi.md#UpdateVideo) | **Put** /media/videos/{id} | Modifies a video&#39;s details
[**UpdateVideoComment**](Media_VideosApi.md#UpdateVideoComment) | **Put** /media/videos/{video_id}/comments/{id}/content | Update a video comment
[**UpdateVideoRelationship**](Media_VideosApi.md#UpdateVideoRelationship) | **Put** /media/videos/{video_id}/related/{id}/relationship_details | Update a video&#39;s relationship details
[**UpdateVideoTemplate**](Media_VideosApi.md#UpdateVideoTemplate) | **Put** /media/videos/templates/{id} | Update a video template
[**ViewVideo**](Media_VideosApi.md#ViewVideo) | **Post** /media/videos/{id}/views | Increment a video&#39;s view count


# **AddUserToVideoWhitelist**
> AddUserToVideoWhitelist(ctx, ctx, id, optional)
Adds a user to a video's whitelist

Whitelisted users can view video regardless of privacy setting. <br><br><b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The video id | 
 **userId** | [**IntWrapper**](IntWrapper.md)| The user id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideo**
> VideoResource AddVideo(ctx, ctx, optional)
Adds a new video in the system

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoResource** | [**VideoResource**](VideoResource.md)| The video object | 

### Return type

[**VideoResource**](VideoResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoComment**
> CommentResource AddVideoComment(ctx, ctx, videoId, optional)
Add a new video comment

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int32**| The video id  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id  | 
 **commentResource** | [**CommentResource**](CommentResource.md)| The comment object | 

### Return type

[**CommentResource**](CommentResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoContributor**
> AddVideoContributor(ctx, ctx, videoId, optional)
Adds a contributor to a video

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **contributionResource** | [**ContributionResource**](ContributionResource.md)| The contribution object | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoFlag**
> FlagResource AddVideoFlag(ctx, ctx, videoId, optional)
Add a new flag

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **reason** | [**StringWrapper**](StringWrapper.md)| The flag reason | 

### Return type

[**FlagResource**](FlagResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoRelationships**
> VideoRelationshipResource AddVideoRelationships(ctx, ctx, videoId, optional)
Adds one or more existing videos as related to this one

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **videoRelationshipResource** | [**VideoRelationshipResource**](VideoRelationshipResource.md)| The video relationship object  | 

### Return type

[**VideoRelationshipResource**](VideoRelationshipResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVideoDisposition**
> DispositionResource CreateVideoDisposition(ctx, ctx, videoId, optional)
Create a video disposition

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int32**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id | 
 **dispositionResource** | [**DispositionResource**](DispositionResource.md)| The disposition object | 

### Return type

[**DispositionResource**](DispositionResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVideoTemplate**
> TemplateResource CreateVideoTemplate(ctx, ctx, optional)
Create a video template

Video Templates define a type of video and the properties they have. <br><br><b>Permissions Needed:</b> TEMPLATE_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoTemplateResource** | [**TemplateResource**](TemplateResource.md)| The video template resource object | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideo**
> DeleteVideo(ctx, ctx, id)
Deletes a video from the system if no resources are attached to it

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The video id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoComment**
> DeleteVideoComment(ctx, ctx, videoId, id)
Delete a video comment

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int64**| The comment id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoDisposition**
> DeleteVideoDisposition(ctx, ctx, dispositionId)
Delete a video disposition

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **dispositionId** | **int64**| The disposition id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoFlag**
> DeleteVideoFlag(ctx, ctx, videoId)
Delete a flag

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoRelationship**
> DeleteVideoRelationship(ctx, ctx, videoId, id)
Delete a video's relationship

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int64**| The relationship id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoTemplate**
> DeleteVideoTemplate(ctx, ctx, id, optional)
Delete a video template

If cascade = 'detach', it will force delete the template even if it's attached to other objects. <br><br><b>Permissions Needed:</b> TEMPLATE_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserVideos**
> PageResourceVideoResource GetUserVideos(ctx, ctx, userId, optional)
Get user videos

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **userId** | **int32**| The user id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user id | 
 **excludeFlagged** | **bool**| Skip videos that have been flagged by the current user | [default to true]
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceVideoResource**](PageResource«VideoResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideo**
> VideoResource GetVideo(ctx, ctx, id)
Loads a specific video details

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The video id | 

### Return type

[**VideoResource**](VideoResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoComments**
> PageResourceCommentResource GetVideoComments(ctx, ctx, videoId, optional)
Returns a page of comments for a video

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int32**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceCommentResource**](PageResource«CommentResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoDispositions**
> PageResourceDispositionResource GetVideoDispositions(ctx, ctx, videoId, optional)
Returns a page of dispositions for a video

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int32**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceDispositionResource**](PageResource«DispositionResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoRelationships**
> PageResourceVideoRelationshipResource GetVideoRelationships(ctx, ctx, videoId, optional)
Returns a page of video relationships

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceVideoRelationshipResource**](PageResource«VideoRelationshipResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoTemplate**
> TemplateResource GetVideoTemplate(ctx, ctx, id)
Get a single video template

<b>Permissions Needed:</b> TEMPLATE_ADMIN or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoTemplates**
> PageResourceTemplateResource GetVideoTemplates(ctx, ctx, optional)
List and search video templates

<b>Permissions Needed:</b> TEMPLATE_ADMIN or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideos**
> PageResourceVideoResource GetVideos(ctx, ctx, optional)
Search videos using the documented filters

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeFlagged** | **bool**| Skip videos that have been flagged by the current user | [default to true]
 **filterVideosByUploader** | **int32**| Filter for videos by uploader id | 
 **filterCategory** | **string**| Filter for videos from a specific category by id | 
 **filterTagset** | **string**| Filter for videos with specified tags (separated by comma) | 
 **filterVideosByName** | **string**| Filter for videos which name *STARTS* with the given string | 
 **filterVideosByContributor** | **int32**| Filter for videos with contribution from the artist specified by ID | 
 **filterVideosByAuthor** | **int32**| Filter for videos with an artist as author specified by ID | 
 **filterHasAuthor** | **bool**| Filter for videos that have an author set if true, or that have no author if false | 
 **filterHasUploader** | **bool**| Filter for videos that have an uploader set if true, or that have no uploader if false | 
 **filterRelatedTo** | **string**| Filter for videos that have designated a particular video as the TO of a relationship. Pattern should match VIDEO_ID or VIDEO_ID:DETAILS to match with a specific details string as well | 
 **filterFriends** | **bool**| Filter for videos uploaded by friends. &#39;true&#39; for friends of the caller (requires user token) or a user id for a specific user&#39;s friends (requires VIDEOS_ADMIN permission) | 
 **filterDisposition** | **string**| Filter for videos a given user has a given disposition towards. USER_ID:DISPOSITION where USER_ID is the id of the user who has this disposition or &#39;me&#39; for the caller (requires user token for &#39;me&#39;) and DISPOSITION is the name of the disposition. E.G. filter_disposition&#x3D;123:like or filter_disposition&#x3D;me:favorite | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to author:ASC]

### Return type

[**PageResourceVideoResource**](PageResource«VideoResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserFromVideoWhitelist**
> RemoveUserFromVideoWhitelist(ctx, ctx, videoId, id)
Removes a user from a video's whitelist

Remove the user with the id given in the path from the whitelist of users that can view this video regardless of privacy setting. <br><br><b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int32**| The user id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveVideoContributor**
> RemoveVideoContributor(ctx, ctx, videoId, id)
Removes a contributor from a video

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int32**| The contributor id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideo**
> UpdateVideo(ctx, ctx, id, optional)
Modifies a video's details

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The video id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The video id | 
 **videoResource** | [**VideoResource**](VideoResource.md)| The video object | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideoComment**
> UpdateVideoComment(ctx, ctx, videoId, id, optional)
Update a video comment

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int64**| The comment id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **id** | **int64**| The comment id | 
 **content** | [**StringWrapper**](StringWrapper.md)| The comment content | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideoRelationship**
> UpdateVideoRelationship(ctx, ctx, videoId, relationshipId, optional)
Update a video's relationship details

<b>Permissions Needed:</b> VIDEOS_USER or VIDEOS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **relationshipId** | **int64**| The relationship id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **relationshipId** | **int64**| The relationship id | 
 **details** | [**StringWrapper**](StringWrapper.md)| The video relationship details | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideoTemplate**
> TemplateResource UpdateVideoTemplate(ctx, ctx, id, optional)
Update a video template

<b>Permissions Needed:</b> TEMPLATE_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **videoTemplateResource** | [**TemplateResource**](TemplateResource.md)| The video template resource object | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewVideo**
> ViewVideo(ctx, ctx, id)
Increment a video's view count

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The video id | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

