# \MediaVideosApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToVideoWhitelist**](MediaVideosApi.md#AddUserToVideoWhitelist) | **Post** /media/videos/{id}/whitelist | Adds a user to a video&#39;s whitelist
[**AddVideo**](MediaVideosApi.md#AddVideo) | **Post** /media/videos | Adds a new video in the system
[**AddVideoComment**](MediaVideosApi.md#AddVideoComment) | **Post** /media/videos/{video_id}/comments | Add a new video comment
[**AddVideoContributor**](MediaVideosApi.md#AddVideoContributor) | **Post** /media/videos/{video_id}/contributors | Adds a contributor to a video
[**AddVideoFlag**](MediaVideosApi.md#AddVideoFlag) | **Post** /media/videos/{video_id}/moderation | Add a new flag
[**AddVideoRelationships**](MediaVideosApi.md#AddVideoRelationships) | **Post** /media/videos/{video_id}/related | Adds one or more existing videos as related to this one
[**CreateVideoDisposition**](MediaVideosApi.md#CreateVideoDisposition) | **Post** /media/videos/{video_id}/dispositions | Create a video disposition
[**DeleteVideo**](MediaVideosApi.md#DeleteVideo) | **Delete** /media/videos/{id} | Deletes a video from the system if no resources are attached to it
[**DeleteVideoComment**](MediaVideosApi.md#DeleteVideoComment) | **Delete** /media/videos/{video_id}/comments/{id} | Delete a video comment
[**DeleteVideoDisposition**](MediaVideosApi.md#DeleteVideoDisposition) | **Delete** /media/videos/{video_id}/dispositions/{disposition_id} | Delete a video disposition
[**DeleteVideoFlag**](MediaVideosApi.md#DeleteVideoFlag) | **Delete** /media/videos/{video_id}/moderation | Delete a flag
[**DeleteVideoRelationship**](MediaVideosApi.md#DeleteVideoRelationship) | **Delete** /media/videos/{video_id}/related/{id} | Delete a video&#39;s relationship
[**GetUserVideos**](MediaVideosApi.md#GetUserVideos) | **Get** /users/{user_id}/videos | Get user videos
[**GetVideo**](MediaVideosApi.md#GetVideo) | **Get** /media/videos/{id} | Loads a specific video details
[**GetVideoComments**](MediaVideosApi.md#GetVideoComments) | **Get** /media/videos/{video_id}/comments | Returns a page of comments for a video
[**GetVideoDispositions**](MediaVideosApi.md#GetVideoDispositions) | **Get** /media/videos/{video_id}/dispositions | Returns a page of dispositions for a video
[**GetVideoRelationships**](MediaVideosApi.md#GetVideoRelationships) | **Get** /media/videos/{video_id}/related | Returns a page of video relationships
[**GetVideos**](MediaVideosApi.md#GetVideos) | **Get** /media/videos | Search videos using the documented filters
[**RemoveUserFromVideoWhitelist**](MediaVideosApi.md#RemoveUserFromVideoWhitelist) | **Delete** /media/videos/{video_id}/whitelist/{id} | Removes a user from a video&#39;s whitelist
[**RemoveVideoContributor**](MediaVideosApi.md#RemoveVideoContributor) | **Delete** /media/videos/{video_id}/contributors/{id} | Removes a contributor from a video
[**UpdateVideo**](MediaVideosApi.md#UpdateVideo) | **Put** /media/videos/{id} | Modifies a video&#39;s details
[**UpdateVideoComment**](MediaVideosApi.md#UpdateVideoComment) | **Put** /media/videos/{video_id}/comments/{id}/content | Update a video comment
[**UpdateVideoRelationship**](MediaVideosApi.md#UpdateVideoRelationship) | **Put** /media/videos/{video_id}/related/{id}/relationship_details | Update a video&#39;s relationship details
[**ViewVideo**](MediaVideosApi.md#ViewVideo) | **Post** /media/videos/{id}/views | Increment a video&#39;s view count


# **AddUserToVideoWhitelist**
> AddUserToVideoWhitelist(ctx, id, optional)
Adds a user to a video's whitelist

Whitelisted users can view video regardless of privacy setting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideo**
> VideoResource AddVideo(ctx, optional)
Adds a new video in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoComment**
> CommentResource AddVideoComment(ctx, videoId, optional)
Add a new video comment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoContributor**
> AddVideoContributor(ctx, videoId, optional)
Adds a contributor to a video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoFlag**
> FlagResource AddVideoFlag(ctx, videoId, optional)
Add a new flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoRelationships**
> VideoRelationshipResource AddVideoRelationships(ctx, videoId, optional)
Adds one or more existing videos as related to this one

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVideoDisposition**
> DispositionResource CreateVideoDisposition(ctx, videoId, optional)
Create a video disposition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideo**
> DeleteVideo(ctx, id)
Deletes a video from the system if no resources are attached to it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The video id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoComment**
> DeleteVideoComment(ctx, videoId, id)
Delete a video comment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int64**| The comment id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoDisposition**
> DeleteVideoDisposition(ctx, dispositionId)
Delete a video disposition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **dispositionId** | **int64**| The disposition id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoFlag**
> DeleteVideoFlag(ctx, videoId)
Delete a flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoRelationship**
> DeleteVideoRelationship(ctx, videoId, id)
Delete a video's relationship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int64**| The relationship id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserVideos**
> PageResourceVideoResource GetUserVideos(ctx, userId, optional)
Get user videos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideo**
> VideoResource GetVideo(ctx, id)
Loads a specific video details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The video id | 

### Return type

[**VideoResource**](VideoResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoComments**
> PageResourceCommentResource GetVideoComments(videoId, optional)
Returns a page of comments for a video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoDispositions**
> PageResourceDispositionResource GetVideoDispositions(videoId, optional)
Returns a page of dispositions for a video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoRelationships**
> PageResourceVideoRelationshipResource GetVideoRelationships(videoId, optional)
Returns a page of video relationships

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideos**
> PageResourceVideoResource GetVideos(optional)
Search videos using the documented filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserFromVideoWhitelist**
> RemoveUserFromVideoWhitelist(ctx, videoId, id)
Removes a user from a video's whitelist

Remove the user with the id given in the path from the whitelist of users that can view this video regardless of privacy setting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int32**| The user id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveVideoContributor**
> RemoveVideoContributor(ctx, videoId, id)
Removes a contributor from a video

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **videoId** | **int64**| The video id | 
  **id** | **int32**| The contributor id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideo**
> UpdateVideo(ctx, id, optional)
Modifies a video's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideoComment**
> UpdateVideoComment(ctx, videoId, id, optional)
Update a video comment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideoRelationship**
> UpdateVideoRelationship(ctx, videoId, relationshipId, optional)
Update a video's relationship details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewVideo**
> ViewVideo(id)
Increment a video's view count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **id** | **int64**| The video id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

