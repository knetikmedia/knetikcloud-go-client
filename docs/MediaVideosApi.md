# \MediaVideosApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommentUsingPOST1**](MediaVideosApi.md#AddCommentUsingPOST1) | **Post** /media/videos/{video_id}/comments | Add a new video comment
[**AddContributorUsingPOST**](MediaVideosApi.md#AddContributorUsingPOST) | **Post** /media/videos/{video_id}/contributors | Adds a contributor to a video
[**AddDispositionUsingPOST1**](MediaVideosApi.md#AddDispositionUsingPOST1) | **Post** /media/videos/{video_id}/dispositions | Add a new Video disposition
[**AddFlagUsingPOST**](MediaVideosApi.md#AddFlagUsingPOST) | **Post** /media/videos/{video_id}/moderation | Add a new flag
[**AddRelatedUsingPOST**](MediaVideosApi.md#AddRelatedUsingPOST) | **Post** /media/videos/{video_id}/related | Adds one or more existing videos as related to this one
[**AddVideoUsingPOST**](MediaVideosApi.md#AddVideoUsingPOST) | **Post** /media/videos | Adds a new video in the system
[**AddWhitelistUsingPOST**](MediaVideosApi.md#AddWhitelistUsingPOST) | **Post** /media/videos/{id}/whitelist | Adds a user to a video&#39;s whitelist
[**DeleteCommentUsingDELETE1**](MediaVideosApi.md#DeleteCommentUsingDELETE1) | **Delete** /media/videos/{video_id}/comments/{id} | Delete a video comment
[**DeleteDispositionUsingDELETE1**](MediaVideosApi.md#DeleteDispositionUsingDELETE1) | **Delete** /media/videos/{video_id}/dispositions/{disposition_id} | Delete a video comment
[**DeleteFlagUsingDELETE**](MediaVideosApi.md#DeleteFlagUsingDELETE) | **Delete** /media/videos/{video_id}/moderation | Delete a flag
[**DeleteRelationshipUsingDELETE1**](MediaVideosApi.md#DeleteRelationshipUsingDELETE1) | **Delete** /media/videos/{video_id}/related/{id} | Delete a video&#39;s relationship
[**DeleteVideoUsingDELETE**](MediaVideosApi.md#DeleteVideoUsingDELETE) | **Delete** /media/videos/{id} | Removes a video from the system if no resources are attached to it
[**GetCommentsUsingGET1**](MediaVideosApi.md#GetCommentsUsingGET1) | **Get** /media/videos/{video_id}/comments | Returns a page of comments for a video
[**GetDispositionsUsingGET1**](MediaVideosApi.md#GetDispositionsUsingGET1) | **Get** /media/videos/{video_id}/dispositions | Returns a page of dispositions for a video
[**GetRelatedUsingGET**](MediaVideosApi.md#GetRelatedUsingGET) | **Get** /media/videos/{video_id}/related | Returns a page of video relationships
[**GetUserVideosUsingGET**](MediaVideosApi.md#GetUserVideosUsingGET) | **Get** /users/{user_id}/videos | Get user videos
[**GetVideoUsingGET**](MediaVideosApi.md#GetVideoUsingGET) | **Get** /media/videos/{id} | Loads a specific video details
[**RemoveContributorUsingDELETE**](MediaVideosApi.md#RemoveContributorUsingDELETE) | **Delete** /media/videos/{video_id}/contributors/{id} | Removes a contributor from a video
[**RemoveWhitelistUsingDELETE**](MediaVideosApi.md#RemoveWhitelistUsingDELETE) | **Delete** /media/videos/{video_id}/whitelist/{id} | Removes a user from a video&#39;s whitelist
[**SearchVideosUsingGET**](MediaVideosApi.md#SearchVideosUsingGET) | **Get** /media/videos | Search videos using the documented filters
[**UpdateCommentUsingPUT1**](MediaVideosApi.md#UpdateCommentUsingPUT1) | **Put** /media/videos/{video_id}/comments/{id}/content | Update video comment content 
[**UpdateRelationshipUsingPUT1**](MediaVideosApi.md#UpdateRelationshipUsingPUT1) | **Put** /media/videos/{video_id}/related/{id}/relationship_details | Update a video&#39;s relationship details
[**UpdateVideoUsingPUT**](MediaVideosApi.md#UpdateVideoUsingPUT) | **Put** /media/videos/{id} | Modifies a video&#39;s details
[**ViewVideoUsingPOST**](MediaVideosApi.md#ViewVideoUsingPOST) | **Post** /media/videos/{id}/views | Increment a video&#39;s view count


# **AddCommentUsingPOST1**
> CommentResource AddCommentUsingPOST1($videoId, $commentResource)

Add a new video comment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id  | 
 **commentResource** | [**CommentResource**](CommentResource.md)| The comment object | [optional] 

### Return type

[**CommentResource**](CommentResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddContributorUsingPOST**
> AddContributorUsingPOST($videoId, $contributionResource)

Adds a contributor to a video


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **contributionResource** | [**ContributionResource**](ContributionResource.md)| The contribution object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddDispositionUsingPOST1**
> DispositionResource AddDispositionUsingPOST1($videoId, $dispositionResource)

Add a new Video disposition


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id | 
 **dispositionResource** | [**DispositionResource**](DispositionResource.md)| The disposition object | [optional] 

### Return type

[**DispositionResource**](DispositionResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFlagUsingPOST**
> AddFlagUsingPOST($videoId, $reason)

Add a new flag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **reason** | **string**| The flag reason | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRelatedUsingPOST**
> VideoRelationshipResource AddRelatedUsingPOST($videoId, $videoRelationshipResource)

Adds one or more existing videos as related to this one


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **videoRelationshipResource** | [**VideoRelationshipResource**](VideoRelationshipResource.md)| The video relationship object  | [optional] 

### Return type

[**VideoRelationshipResource**](VideoRelationshipResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVideoUsingPOST**
> VideoResource AddVideoUsingPOST($videoResource)

Adds a new video in the system


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoResource** | [**VideoResource**](VideoResource.md)| The video object | [optional] 

### Return type

[**VideoResource**](VideoResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddWhitelistUsingPOST**
> AddWhitelistUsingPOST($id, $userId)

Adds a user to a video's whitelist

Whitelisted users can view video regardless of privacy setting.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The video id | 
 **userId** | **int32**| The user id | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommentUsingDELETE1**
> DeleteCommentUsingDELETE1($videoId, $id)

Delete a video comment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **id** | **int64**| The comment id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDispositionUsingDELETE1**
> DeleteDispositionUsingDELETE1($dispositionId)

Delete a video comment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dispositionId** | **int64**| The disposition id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFlagUsingDELETE**
> DeleteFlagUsingDELETE($videoId)

Delete a flag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRelationshipUsingDELETE1**
> DeleteRelationshipUsingDELETE1($videoId, $id)

Delete a video's relationship


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **id** | **int64**| The relationship id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVideoUsingDELETE**
> DeleteVideoUsingDELETE($id)

Removes a video from the system if no resources are attached to it


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The video id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommentsUsingGET1**
> PageResourceCommentResource GetCommentsUsingGET1($videoId, $size, $page)

Returns a page of comments for a video


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceCommentResource**](PageResource«CommentResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDispositionsUsingGET1**
> PageResourceDispositionResource GetDispositionsUsingGET1($videoId, $size, $page)

Returns a page of dispositions for a video


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int32**| The video id | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceDispositionResource**](PageResource«DispositionResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelatedUsingGET**
> PageResourceVideoRelationshipResource GetRelatedUsingGET($videoId, $size, $page)

Returns a page of video relationships


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceVideoRelationshipResource**](PageResource«VideoRelationshipResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserVideosUsingGET**
> PageResourceVideoResource GetUserVideosUsingGET($userId, $excludeFlagged, $size, $page)

Get user videos


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user id | 
 **excludeFlagged** | **bool**| Skip videos that have been flagged by the current user | [optional] [default to true]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceVideoResource**](PageResource«VideoResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoUsingGET**
> VideoResource GetVideoUsingGET($id)

Loads a specific video details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The video id | 

### Return type

[**VideoResource**](VideoResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContributorUsingDELETE**
> RemoveContributorUsingDELETE($videoId, $id)

Removes a contributor from a video


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **id** | **int32**| The contributor id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveWhitelistUsingDELETE**
> RemoveWhitelistUsingDELETE($videoId, $id)

Removes a user from a video's whitelist

Remove the user with the id given in the path from the whitelist of users that can view this video regardless of privacy setting.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **id** | **int32**| The user id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchVideosUsingGET**
> PageResourceVideoResource SearchVideosUsingGET($excludeFlagged, $filterVideosByUploader, $filterCategory, $filterTagset, $filterVideosByName, $filterVideosByContributor, $filterVideosByAuthor, $filterHasAuthor, $filterHasUploader, $filterRelatedTo, $filterFriends, $filterDisposition, $size, $page, $order)

Search videos using the documented filters


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeFlagged** | **bool**| Skip videos that have been flagged by the current user | [optional] [default to true]
 **filterVideosByUploader** | **string**| Filter for videos by uploader id | [optional] 
 **filterCategory** | **string**| Filter for videos from a specific category by id | [optional] 
 **filterTagset** | **string**| Filter for videos with specified tags (separated by comma) | [optional] 
 **filterVideosByName** | **string**| Filter for videos which name *STARTS* with the given string | [optional] 
 **filterVideosByContributor** | **string**| Filter for videos with contribution from the artist specified by ID | [optional] 
 **filterVideosByAuthor** | **string**| Filter for videos with an artist as author specified by ID | [optional] 
 **filterHasAuthor** | **bool**| Filter for videos that have an author set if true, or that have no author if false | [optional] 
 **filterHasUploader** | **bool**| Filter for videos that have an uploader set if true, or that have no uploader if false | [optional] 
 **filterRelatedTo** | **string**| Filter for videos that have designated a particular video as the TO of a relationship. Pattern should match VIDEO_ID or VIDEO_ID:DETAILS to match with a specific details string as well | [optional] 
 **filterFriends** | **bool**| Filter for videos uploaded by friends. &#39;true&#39; for friends of the caller (requires user token) or a user id for a specific user&#39;s friends (requires VIDEOS_ADMIN permission) | [optional] 
 **filterDisposition** | **string**| Filter for videos a given user has a given disposition towards. USER_ID:DISPOSITION where USER_ID is the id of the user who has this disposition or &#39;me&#39; for the caller (requires user token for &#39;me&#39;) and DISPOSITION is the name of the disposition. E.G. filter_disposition&#x3D;123:like or filter_disposition&#x3D;me:favorite | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to author:ASC]

### Return type

[**PageResourceVideoResource**](PageResource«VideoResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommentUsingPUT1**
> UpdateCommentUsingPUT1($videoId, $id, $content)

Update video comment content 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **id** | **int64**| The comment id | 
 **content** | **string**| The comment content | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRelationshipUsingPUT1**
> UpdateRelationshipUsingPUT1($videoId, $relationshipId, $details)

Update a video's relationship details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **int64**| The video id | 
 **relationshipId** | **int64**| The relationship id | 
 **details** | **string**| The video relationship details | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideoUsingPUT**
> UpdateVideoUsingPUT($id, $videoResource)

Modifies a video's details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The video id | 
 **videoResource** | [**VideoResource**](VideoResource.md)| The video object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewVideoUsingPOST**
> ViewVideoUsingPOST($id)

Increment a video's view count


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The video id | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

